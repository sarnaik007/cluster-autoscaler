/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package simulator

import (
	"fmt"
	"sync"

	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/informers"
	kube_client "k8s.io/client-go/kubernetes"
	"k8s.io/klog"

	"k8s.io/kubernetes/pkg/scheduler"
	"k8s.io/kubernetes/pkg/scheduler/algorithm/predicates"
	schedulernodeinfo "k8s.io/kubernetes/pkg/scheduler/nodeinfo"
	// We need to import provider to initialize default scheduler.
	"k8s.io/kubernetes/pkg/scheduler/algorithmprovider"
)

const (
	// We want to disable affinity predicate for performance reasons if no pod
	// requires it.
	affinityPredicateName = "MatchInterPodAffinity"
)

var (
	// initMutex is used for guarding static initialization.
	staticInitMutex sync.Mutex
	// statiInitHappened denotes if static initialization happened.
	staticInitDone bool
)

// PredicateInfo assigns a name to a predicate
type PredicateInfo struct {
	Name      string
	Predicate predicates.FitPredicate
}

// PredicateChecker checks whether all required predicates pass for given Pod and Node.
type PredicateChecker struct {
	predicates              []PredicateInfo
	enableAffinityPredicate bool
	scheduler               *scheduler.Scheduler
}

// We run some predicates first as they are cheap to check and they should be enough
// to fail predicates in most of our simulations (especially binpacking).
// There are no const arrays in Go, this is meant to be used as a const.
var priorityPredicates = []string{"PodFitsResources", "PodToleratesNodeTaints", "GeneralPredicates", "ready"}

func staticInitIfNeeded() {
	staticInitMutex.Lock()
	defer staticInitMutex.Unlock()

	if staticInitDone {
		return
	}

	// This results in filtering out some predicate functions registered by defaults.init() method.
	// In scheduler this method is run from app.runCommand().
	// We also need to call it in CA to have simulation behaviour consistent with scheduler.
	// Note: the logic of method is conditional and depends on feature gates enabled. To have same
	//       behaviour in CA and scheduler both need to be run with same set of feature gates.
	algorithmprovider.ApplyFeatureGates()
	staticInitDone = true
}

// NoOpEventRecorder is a noop implementation of EventRecorder
type NoOpEventRecorder struct{}

// Event is a noop method implementation
func (NoOpEventRecorder) Event(object runtime.Object, eventtype, reason, message string) {
}

// Eventf is a noop method implementation
func (NoOpEventRecorder) Eventf(regarding runtime.Object, related runtime.Object, eventtype, reason, action, note string, args ...interface{}) {
}

// PastEventf is a noop method implementation
func (NoOpEventRecorder) PastEventf(object runtime.Object, timestamp metav1.Time, eventtype, reason, messageFmt string, args ...interface{}) {
}

// AnnotatedEventf is a noop method implementation
func (NoOpEventRecorder) AnnotatedEventf(object runtime.Object, annotations map[string]string, eventtype, reason, messageFmt string, args ...interface{}) {
}

// NewPredicateChecker builds PredicateChecker.
func NewPredicateChecker(kubeClient kube_client.Interface, stop <-chan struct{}) (*PredicateChecker, error) {
	staticInitIfNeeded()

	informerFactory := informers.NewSharedInformerFactory(kubeClient, 0)
	podInformer := informerFactory.Core().V1().Pods()

	sched, err := scheduler.New(
		kubeClient,
		informerFactory,
		podInformer,
		NoOpEventRecorder{},
		stop,
		scheduler.WithFrameworkConfigProducerRegistry(nil),
	)
	if err != nil {
		return nil, fmt.Errorf("couldn't create scheduler; %v", err)
	}

	predicateMap := map[string]predicates.FitPredicate{}
	for predicateName, predicateFunc := range sched.Algorithm.Predicates() {
		predicateMap[predicateName] = predicateFunc
	}
	// We want to make sure that some predicates are present to run them first
	// as they are cheap to check and they should be enough to fail predicates
	// in most of our simulations (especially binpacking).
	if _, found := predicateMap["PodFitsResources"]; !found {
		predicateMap["PodFitsResources"] = predicates.PodFitsResources
	}
	if _, found := predicateMap["PodToleratesNodeTaints"]; !found {
		predicateMap["PodToleratesNodeTaints"] = predicates.PodToleratesNodeTaints
	}

	predicateList := make([]PredicateInfo, 0, len(predicateMap))
	for _, predicateName := range priorityPredicates {
		if predicate, found := predicateMap[predicateName]; found {
			predicateList = append(predicateList, PredicateInfo{Name: predicateName, Predicate: predicate})
			delete(predicateMap, predicateName)
		}
	}

	for predicateName, predicate := range predicateMap {
		predicateList = append(predicateList, PredicateInfo{Name: predicateName, Predicate: predicate})
	}

	for _, predInfo := range predicateList {
		klog.V(1).Infof("Using predicate %s", predInfo.Name)
	}

	informerFactory.Start(stop)

	return &PredicateChecker{
		predicates:              predicateList,
		enableAffinityPredicate: true,
		scheduler:               sched,
	}, nil
}

// NewTestPredicateChecker builds test version of PredicateChecker.
func NewTestPredicateChecker() *PredicateChecker {
	return &PredicateChecker{
		predicates: []PredicateInfo{
			{Name: "default", Predicate: predicates.GeneralPredicates},
		},
	}
}

// NewCustomTestPredicateChecker builds test version of PredicateChecker with additional predicates.
// Helps with benchmarking different ordering of predicates.
func NewCustomTestPredicateChecker(predicateInfos []PredicateInfo) *PredicateChecker {
	return &PredicateChecker{
		predicates: predicateInfos,
	}
}

// SnapshotClusterState updates cluster snapshot used by the predicate checker.
// It should be called every CA loop iteration.
func (p *PredicateChecker) SnapshotClusterState() error {
	if p.scheduler != nil {
		return p.scheduler.Algorithm.Snapshot()
	}
	return nil
}

// SetAffinityPredicateEnabled can be used to enable or disable checking MatchInterPodAffinity
// predicate. This will cause incorrect CA behavior if there is at least a single pod in
// cluster using affinity/antiaffinity. However, checking affinity predicate is extremely
// costly even if no pod is using it, so it may be worth disabling it in such situation.
func (p *PredicateChecker) SetAffinityPredicateEnabled(enable bool) {
	p.enableAffinityPredicate = enable
}

// IsAffinityPredicateEnabled checks if affinity predicate is enabled.
func (p *PredicateChecker) IsAffinityPredicateEnabled() bool {
	return p.enableAffinityPredicate
}

// FitsAny checks if the given pod can be place on any of the given nodes.
func (p *PredicateChecker) FitsAny(pod *apiv1.Pod, nodeInfos map[string]*schedulernodeinfo.NodeInfo) (string, error) {
	for name, nodeInfo := range nodeInfos {
		// Be sure that the node is schedulable.
		if nodeInfo.Node().Spec.Unschedulable {
			continue
		}
		if err := p.CheckPredicates(pod, nodeInfo); err == nil {
			return name, nil
		}
	}
	return "", fmt.Errorf("cannot put pod %s on any node", pod.Name)
}

// CheckPredicates checks if the given pod can be placed on the given node.
func (p *PredicateChecker) CheckPredicates(pod *apiv1.Pod, nodeInfo *schedulernodeinfo.NodeInfo) PredicateError {
	for _, predInfo := range p.predicates {
		// Skip affinity predicate if it has been disabled.
		if !p.enableAffinityPredicate && predInfo.Name == affinityPredicateName {
			continue
		}

		match, failureReasons, err := predInfo.Predicate(pod, nil, nodeInfo)
		if err != nil || !match {
			return NewPredicateError(
				NotSchedulablePredicateError,
				predInfo.Name,
				"",
				nil, // dropping here due to upcoming rework for scheduler framework
				p.buildDebugInfo(predInfo, nodeInfo))
		}
	}
	return nil
}

func (p *PredicateChecker) buildDebugInfo(predInfo PredicateInfo, nodeInfo *schedulernodeinfo.NodeInfo) func() string {
	switch predInfo.Name {
	case "PodToleratesNodeTaints":
		taints := nodeInfo.Node().Spec.Taints
		return func() string {
			return fmt.Sprintf("taints on node: %#v", taints)
		}
	default:
		return emptyString
	}
}
