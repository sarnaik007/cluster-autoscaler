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
	"testing"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/predicatechecker"
	"k8s.io/autoscaler/cluster-autoscaler/utils/drain"
	kube_util "k8s.io/autoscaler/cluster-autoscaler/utils/kubernetes"
	. "k8s.io/autoscaler/cluster-autoscaler/utils/test"
	no "k8s.io/autoscaler/cluster-autoscaler/utils/test/node"
	po "k8s.io/autoscaler/cluster-autoscaler/utils/test/pod"

	"github.com/stretchr/testify/assert"
	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	policyv1 "k8s.io/api/policy/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/kubernetes/pkg/kubelet/types"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

func TestFindEmptyNodes(t *testing.T) {
	nodes := []*apiv1.Node{}
	nodeNames := []string{}
	for i := 0; i < 4; i++ {
		nodeName := fmt.Sprintf("n%d", i)
		node := no.BuildTestNode(nodeName, 1000, 2000000)
		no.SetNodeReadyState(node, true, time.Time{})
		nodes = append(nodes, node)
		nodeNames = append(nodeNames, nodeName)
	}

	pod1 := po.BuildTestPod("p1", 300, 500000)
	pod1.Spec.NodeName = "n1"

	pod2 := po.BuildTestPod("p2", 300, 500000)
	pod2.Spec.NodeName = "n2"
	pod2.Annotations = map[string]string{
		types.ConfigMirrorAnnotationKey: "",
	}

	clusterSnapshot := clustersnapshot.NewBasicClusterSnapshot()
	clustersnapshot.InitializeClusterSnapshotOrDie(t, clusterSnapshot, []*apiv1.Node{nodes[0], nodes[1], nodes[2], nodes[3]}, []*apiv1.Pod{pod1, pod2})
	testTime := time.Date(2020, time.December, 18, 17, 0, 0, 0, time.UTC)
	r := NewRemovalSimulator(nil, clusterSnapshot, nil, nil, testDeleteOptions(), false)
	emptyNodes := r.FindEmptyNodesToRemove(nodeNames, testTime)
	assert.Equal(t, []string{nodeNames[0], nodeNames[2], nodeNames[3]}, emptyNodes)
}

type findNodesToRemoveTestConfig struct {
	name        string
	pods        []*apiv1.Pod
	allNodes    []*apiv1.Node
	candidates  []string
	toRemove    []NodeToBeRemoved
	unremovable []*UnremovableNode
}

func TestFindNodesToRemove(t *testing.T) {
	emptyNode := no.BuildTestNode("n1", 1000, 2000000)
	emptyNodeInfo := schedulerframework.NewNodeInfo()
	emptyNodeInfo.SetNode(emptyNode)

	// two small pods backed by ReplicaSet
	drainableNode := no.BuildTestNode("n2", 1000, 2000000)
	drainableNodeInfo := schedulerframework.NewNodeInfo()
	drainableNodeInfo.SetNode(drainableNode)

	// one small pod, not backed by anything
	nonDrainableNode := no.BuildTestNode("n3", 1000, 2000000)
	nonDrainableNodeInfo := schedulerframework.NewNodeInfo()
	nonDrainableNodeInfo.SetNode(nonDrainableNode)

	// one very large pod
	fullNode := no.BuildTestNode("n4", 1000, 2000000)
	fullNodeInfo := schedulerframework.NewNodeInfo()
	fullNodeInfo.SetNode(fullNode)

	no.SetNodeReadyState(emptyNode, true, time.Time{})
	no.SetNodeReadyState(drainableNode, true, time.Time{})
	no.SetNodeReadyState(nonDrainableNode, true, time.Time{})
	no.SetNodeReadyState(fullNode, true, time.Time{})

	replicas := int32(5)
	replicaSets := []*appsv1.ReplicaSet{
		{
			ObjectMeta: metav1.ObjectMeta{
				Name:      "rs",
				Namespace: "default",
				SelfLink:  "api/v1/namespaces/default/replicasets/rs",
			},
			Spec: appsv1.ReplicaSetSpec{
				Replicas: &replicas,
			},
		},
	}
	rsLister, err := kube_util.NewTestReplicaSetLister(replicaSets)
	assert.NoError(t, err)
	registry := kube_util.NewListerRegistry(nil, nil, nil, nil, nil, nil, nil, nil, rsLister, nil)

	ownerRefs := GenerateOwnerReferences("rs", "ReplicaSet", "extensions/v1beta1", "")

	pod1 := po.BuildTestPod("p1", 100, 100000)
	pod1.OwnerReferences = ownerRefs
	pod1.Spec.NodeName = "n2"
	drainableNodeInfo.AddPod(pod1)

	pod2 := po.BuildTestPod("p2", 100, 100000)
	pod2.OwnerReferences = ownerRefs
	pod2.Spec.NodeName = "n2"
	drainableNodeInfo.AddPod(pod2)

	pod3 := po.BuildTestPod("p3", 100, 100000)
	pod3.Spec.NodeName = "n3"
	nonDrainableNodeInfo.AddPod(pod3)

	pod4 := po.BuildTestPod("p4", 1000, 100000)
	pod4.Spec.NodeName = "n4"
	fullNodeInfo.AddPod(pod4)

	emptyNodeToRemove := NodeToBeRemoved{
		Node:             emptyNode,
		PodsToReschedule: []*apiv1.Pod{},
		DaemonSetPods:    []*apiv1.Pod{},
	}
	drainableNodeToRemove := NodeToBeRemoved{
		Node:             drainableNode,
		PodsToReschedule: []*apiv1.Pod{pod1, pod2},
		DaemonSetPods:    []*apiv1.Pod{},
	}

	clusterSnapshot := clustersnapshot.NewBasicClusterSnapshot()
	predicateChecker, err := predicatechecker.NewTestPredicateChecker()
	assert.NoError(t, err)
	tracker := NewUsageTracker()

	tests := []findNodesToRemoveTestConfig{
		// just an empty node, should be removed
		{
			name:        "just an empty node, should be removed",
			pods:        []*apiv1.Pod{},
			candidates:  []string{emptyNode.Name},
			allNodes:    []*apiv1.Node{emptyNode},
			toRemove:    []NodeToBeRemoved{emptyNodeToRemove},
			unremovable: []*UnremovableNode{},
		},
		// just a drainable node, but nowhere for pods to go to
		{
			name:        "just a drainable node, but nowhere for pods to go to",
			pods:        []*apiv1.Pod{pod1, pod2},
			candidates:  []string{drainableNode.Name},
			allNodes:    []*apiv1.Node{drainableNode},
			toRemove:    []NodeToBeRemoved{},
			unremovable: []*UnremovableNode{{Node: drainableNode, Reason: NoPlaceToMovePods}},
		},
		// drainable node, and a mostly empty node that can take its pods
		{
			name:        "drainable node, and a mostly empty node that can take its pods",
			pods:        []*apiv1.Pod{pod1, pod2, pod3},
			candidates:  []string{drainableNode.Name, nonDrainableNode.Name},
			allNodes:    []*apiv1.Node{drainableNode, nonDrainableNode},
			toRemove:    []NodeToBeRemoved{drainableNodeToRemove},
			unremovable: []*UnremovableNode{{Node: nonDrainableNode, Reason: BlockedByPod, BlockingPod: &drain.BlockingPod{Pod: pod3, Reason: drain.NotReplicated}}},
		},
		// drainable node, and a full node that cannot fit anymore pods
		{
			name:        "drainable node, and a full node that cannot fit anymore pods",
			pods:        []*apiv1.Pod{pod1, pod2, pod4},
			candidates:  []string{drainableNode.Name},
			allNodes:    []*apiv1.Node{drainableNode, fullNode},
			toRemove:    []NodeToBeRemoved{},
			unremovable: []*UnremovableNode{{Node: drainableNode, Reason: NoPlaceToMovePods}},
		},
		// 4 nodes, 1 empty, 1 drainable
		{
			name:        "4 nodes, 1 empty, 1 drainable",
			pods:        []*apiv1.Pod{pod1, pod2, pod3, pod4},
			candidates:  []string{emptyNode.Name, drainableNode.Name},
			allNodes:    []*apiv1.Node{emptyNode, drainableNode, fullNode, nonDrainableNode},
			toRemove:    []NodeToBeRemoved{emptyNodeToRemove, drainableNodeToRemove},
			unremovable: []*UnremovableNode{},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			destinations := make([]string, 0, len(test.allNodes))
			for _, node := range test.allNodes {
				destinations = append(destinations, node.Name)
			}
			clustersnapshot.InitializeClusterSnapshotOrDie(t, clusterSnapshot, test.allNodes, test.pods)
			r := NewRemovalSimulator(registry, clusterSnapshot, predicateChecker, tracker, testDeleteOptions(), false)
			toRemove, unremovable := r.FindNodesToRemove(test.candidates, destinations, time.Now(), []*policyv1.PodDisruptionBudget{})
			fmt.Printf("Test scenario: %s, found len(toRemove)=%v, expected len(test.toRemove)=%v\n", test.name, len(toRemove), len(test.toRemove))
			assert.Equal(t, toRemove, test.toRemove)
			assert.Equal(t, unremovable, test.unremovable)
		})
	}
}

func testDeleteOptions() NodeDeleteOptions {
	return NodeDeleteOptions{
		SkipNodesWithSystemPods:           true,
		SkipNodesWithLocalStorage:         true,
		MinReplicaCount:                   0,
		SkipNodesWithCustomControllerPods: true,
	}
}
