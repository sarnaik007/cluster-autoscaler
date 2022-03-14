/*
Copyright 2019 The Kubernetes Authors.

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

package nodes

import (
	apiv1 "k8s.io/api/core/v1"
	kube_client "k8s.io/client-go/kubernetes"
	klog "k8s.io/klog/v2"
	"strings"

	"k8s.io/autoscaler/cluster-autoscaler/context"
	"k8s.io/autoscaler/cluster-autoscaler/utils"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
)

// PreFilteringScaleDownNodeProcessor filters out scale down candidates from nodegroup with
// size <= minimum number of nodes for that nodegroup and filters out node from non-autoscaled
// nodegroups
type PreFilteringScaleDownNodeProcessor struct {
}

// GetPodDestinationCandidates returns nodes that potentially could act as destinations for pods
// that would become unscheduled after a scale down.
func (n *PreFilteringScaleDownNodeProcessor) GetPodDestinationCandidates(ctx *context.AutoscalingContext,
	nodes []*apiv1.Node) ([]*apiv1.Node, errors.AutoscalerError) {
	return nodes, nil
}

// GetScaleDownCandidates returns nodes that potentially could be scaled down and
func (n *PreFilteringScaleDownNodeProcessor) GetScaleDownCandidates(ctx *context.AutoscalingContext,
	nodes []*apiv1.Node, kubeclient kube_client.Interface) ([]*apiv1.Node, errors.AutoscalerError) {
	result := make([]*apiv1.Node, 0, len(nodes))

	var numberWorkerNode int = 0
	for _, node := range nodes {
		if strings.Contains(node.Name, "worker") {
			numberWorkerNode += 1
		}
	}

	for _, node := range nodes {
		//nodeGroup, err := ctx.CloudProvider.NodeGroupForNode(node)
		//if err != nil {
		//	klog.Warningf("Error while checking node group for %s: %v", node.Name, err)
		//	continue
		//}
		//if nodeGroup == nil || reflect.ValueOf(nodeGroup).IsNil() {
		//	klog.V(4).Infof("Node %s should not be processed by cluster autoscaler (no node group config)", node.Name)
		//	continue
		//}
		size := numberWorkerNode
		//if !found {
		//	klog.Errorf("Error while checking node group size %s: group size not found", nodeGroup.Id())
		//	continue
		//}
		if size <= utils.GetMinSizeNodeGroup(kubeclient) {
			klog.V(1).Infof("Skipping %s - node group min size reached", node.Name)
			continue
		}
		result = append(result, node)
	}
	return result, nil
}

// CleanUp is called at CA termination.
func (n *PreFilteringScaleDownNodeProcessor) CleanUp() {
}

// NewPreFilteringScaleDownNodeProcessor returns a new PreFilteringScaleDownNodeProcessor.
func NewPreFilteringScaleDownNodeProcessor() *PreFilteringScaleDownNodeProcessor {
	return &PreFilteringScaleDownNodeProcessor{}
}
