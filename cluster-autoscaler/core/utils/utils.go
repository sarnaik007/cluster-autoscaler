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

package utils

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	"k8s.io/autoscaler/cluster-autoscaler/clusterstate"
	"k8s.io/autoscaler/cluster-autoscaler/metrics"
	"k8s.io/autoscaler/cluster-autoscaler/simulator/clustersnapshot"
	"k8s.io/autoscaler/cluster-autoscaler/utils/daemonset"
	"k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	caerrors "k8s.io/autoscaler/cluster-autoscaler/utils/errors"
	"k8s.io/autoscaler/cluster-autoscaler/utils/gpu"
	"k8s.io/autoscaler/cluster-autoscaler/utils/labels"
	"k8s.io/autoscaler/cluster-autoscaler/utils/taints"
	klog "k8s.io/klog/v2"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

// GetNodeInfoFromTemplate returns NodeInfo object built base on TemplateNodeInfo returned by NodeGroup.TemplateNodeInfo().
func GetNodeInfoFromTemplate(nodeGroup cloudprovider.NodeGroup, daemonsets []*appsv1.DaemonSet, taintConfig taints.TaintConfig) (*schedulerframework.NodeInfo, errors.AutoscalerError) {
	id := nodeGroup.Id()
	baseNodeInfo, err := nodeGroup.TemplateNodeInfo()
	if err != nil {
		return nil, errors.ToAutoscalerError(errors.CloudProviderError, err)
	}

	labels.UpdateDeprecatedLabels(baseNodeInfo.Node().ObjectMeta.Labels)

	sanitizedNode, typedErr := SanitizeNode(baseNodeInfo.Node(), id, taintConfig)
	if err != nil {
		return nil, typedErr
	}
	baseNodeInfo.SetNode(sanitizedNode)

	pods, err := daemonset.GetDaemonSetPodsForNode(baseNodeInfo, daemonsets)
	if err != nil {
		return nil, errors.ToAutoscalerError(errors.InternalError, err)
	}
	for _, podInfo := range baseNodeInfo.Pods {
		pods = append(pods, podInfo.Pod)
	}

	sanitizedNodeInfo := schedulerframework.NewNodeInfo(SanitizePods(pods, sanitizedNode)...)
	sanitizedNodeInfo.SetNode(sanitizedNode)
	return sanitizedNodeInfo, nil
}

// isVirtualNode determines if the node is created by virtual kubelet
func isVirtualNode(node *apiv1.Node) bool {
	return node.ObjectMeta.Labels["type"] == "virtual-kubelet"
}

// FilterOutNodesFromNotAutoscaledGroups return subset of input nodes for which cloud provider does not
// return autoscaled node group.
func FilterOutNodesFromNotAutoscaledGroups(nodes []*apiv1.Node, cloudProvider cloudprovider.CloudProvider) ([]*apiv1.Node, errors.AutoscalerError) {
	result := make([]*apiv1.Node, 0)

	for _, node := range nodes {
		// Exclude the virtual node here since it may have lots of resource and exceed the total resource limit
		if isVirtualNode(node) {
			continue
		}
		nodeGroup, err := cloudProvider.NodeGroupForNode(node)
		if err != nil {
			return []*apiv1.Node{}, errors.ToAutoscalerError(errors.CloudProviderError, err)
		}
		if nodeGroup == nil || reflect.ValueOf(nodeGroup).IsNil() {
			result = append(result, node)
		}
	}
	return result, nil
}

// DeepCopyNodeInfo clones the provided nodeInfo
func DeepCopyNodeInfo(nodeInfo *schedulerframework.NodeInfo) *schedulerframework.NodeInfo {
	newPods := make([]*apiv1.Pod, 0)
	for _, podInfo := range nodeInfo.Pods {
		newPods = append(newPods, podInfo.Pod.DeepCopy())
	}

	// Build a new node info.
	newNodeInfo := schedulerframework.NewNodeInfo(newPods...)
	newNodeInfo.SetNode(nodeInfo.Node().DeepCopy())
	return newNodeInfo
}

// SanitizeNode cleans up nodes used for node group templates
func SanitizeNode(node *apiv1.Node, nodeGroup string, taintConfig taints.TaintConfig) (*apiv1.Node, errors.AutoscalerError) {
	newNode := node.DeepCopy()
	nodeName := fmt.Sprintf("template-node-for-%s-%d", nodeGroup, rand.Int63())
	newNode.Labels = make(map[string]string, len(node.Labels))
	for k, v := range node.Labels {
		if k != apiv1.LabelHostname {
			newNode.Labels[k] = v
		} else {
			newNode.Labels[k] = nodeName
		}
	}
	newNode.Name = nodeName
	newNode.Spec.Taints = taints.SanitizeTaints(newNode.Spec.Taints, taintConfig)
	return newNode, nil
}

// SanitizePods cleans up pods used for node group templates
func SanitizePods(pods []*apiv1.Pod, sanitizedNode *apiv1.Node) []*apiv1.Pod {
	// Update node name in pods.
	sanitizedPods := make([]*apiv1.Pod, 0)
	for _, pod := range pods {
		sanitizedPod := pod.DeepCopy()
		sanitizedPod.Spec.NodeName = sanitizedNode.Name
		sanitizedPods = append(sanitizedPods, sanitizedPod)
	}

	return sanitizedPods
}

func hasHardInterPodAffinity(affinity *apiv1.Affinity) bool {
	if affinity == nil {
		return false
	}
	if affinity.PodAffinity != nil {
		if len(affinity.PodAffinity.RequiredDuringSchedulingIgnoredDuringExecution) > 0 {
			return true
		}
	}
	if affinity.PodAntiAffinity != nil {
		if len(affinity.PodAntiAffinity.RequiredDuringSchedulingIgnoredDuringExecution) > 0 {
			return true
		}
	}
	return false
}

// GetNodeCoresAndMemory extracts cpu and memory resources out of Node object
func GetNodeCoresAndMemory(node *apiv1.Node) (int64, int64) {
	cores := getNodeResource(node, apiv1.ResourceCPU)
	memory := getNodeResource(node, apiv1.ResourceMemory)
	return cores, memory
}

func getNodeResource(node *apiv1.Node, resource apiv1.ResourceName) int64 {
	nodeCapacity, found := node.Status.Capacity[resource]
	if !found {
		return 0
	}

	nodeCapacityValue := nodeCapacity.Value()
	if nodeCapacityValue < 0 {
		nodeCapacityValue = 0
	}

	return nodeCapacityValue
}

// UpdateClusterStateMetrics updates metrics related to cluster state
func UpdateClusterStateMetrics(csr *clusterstate.ClusterStateRegistry) {
	if csr == nil || reflect.ValueOf(csr).IsNil() {
		return
	}
	metrics.UpdateClusterSafeToAutoscale(csr.IsClusterHealthy())
	readiness := csr.GetClusterReadiness()
	metrics.UpdateNodesCount(len(readiness.Ready), len(readiness.Unready), len(readiness.NotStarted), len(readiness.LongUnregistered), len(readiness.Unregistered))
}

// GetOldestCreateTime returns oldest creation time out of the pods in the set
func GetOldestCreateTime(pods []*apiv1.Pod) time.Time {
	oldest := time.Now()
	for _, pod := range pods {
		if oldest.After(pod.CreationTimestamp.Time) {
			oldest = pod.CreationTimestamp.Time
		}
	}
	return oldest
}

// GetOldestCreateTimeWithGpu returns oldest creation time out of pods with GPU in the set
func GetOldestCreateTimeWithGpu(pods []*apiv1.Pod) (bool, time.Time) {
	oldest := time.Now()
	gpuFound := false
	for _, pod := range pods {
		if gpu.PodRequestsGpu(pod) {
			gpuFound = true
			if oldest.After(pod.CreationTimestamp.Time) {
				oldest = pod.CreationTimestamp.Time
			}
		}
	}
	return gpuFound, oldest
}

func InitializeClusterSnapshot(c clustersnapshot.ClusterSnapshot, nodes []*apiv1.Node, scheduledPods []*apiv1.Pod) caerrors.AutoscalerError {
	c.Clear()

	knownNodes := make(map[string]bool)
	for _, node := range nodes {
		if err := c.AddNode(node); err != nil {
			klog.Errorf("Failed to add node %s to cluster snapshot: %v", node.Name, err)
			return caerrors.ToAutoscalerError(caerrors.InternalError, err)
		}
		knownNodes[node.Name] = true
	}
	for _, pod := range scheduledPods {
		if knownNodes[pod.Spec.NodeName] {
			if err := c.AddPod(pod, pod.Spec.NodeName); err != nil {
				klog.Errorf("Failed to add pod %s scheduled to node %s to cluster snapshot: %v", pod.Name, pod.Spec.NodeName, err)
				return caerrors.ToAutoscalerError(caerrors.InternalError, err)
			}
		}
	}
	return nil
}
