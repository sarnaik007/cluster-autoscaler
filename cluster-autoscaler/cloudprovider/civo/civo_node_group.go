/*
Copyright 2022 The Kubernetes Authors.

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

package civo

import (
	"errors"
	"fmt"

	"github.com/civo/civogo"

	apiv1 "k8s.io/api/core/v1"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
	autoscaler "k8s.io/autoscaler/cluster-autoscaler/config"
	"k8s.io/klog/v2"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"
)

// NodeGroup implements cloudprovider.NodeGroup interface. NodeGroup contains
// configuration info and functions to control a set of nodes that have the
// same capacity and set of labels.
type NodeGroup struct {
	id         string
	clusterID  string
	client     nodeGroupClient
	nodePool   *civogo.KubernetesPool
	minSize    int
	maxSize    int
	getOptions *autoscaler.NodeGroupAutoscalingOptions
}

// MaxSize returns maximum size of the node group.
func (n *NodeGroup) MaxSize() int {
	return n.maxSize
}

// MinSize returns minimum size of the node group.
func (n *NodeGroup) MinSize() int {
	return n.minSize
}

// GetOptions returns the options used to create this node group.
func (n *NodeGroup) GetOptions(autoscaler.NodeGroupAutoscalingOptions) (*autoscaler.NodeGroupAutoscalingOptions, error) {
	return n.getOptions, nil
}

// TargetSize returns the current target size of the node group. It is possible
// that the number of nodes in Kubernetes is different at the moment but should
// be equal to Size() once everything stabilizes (new nodes finish startup and
// registration or removed nodes are deleted completely). Implementation
// required.
func (n *NodeGroup) TargetSize() (int, error) {
	return n.nodePool.Count, nil
}

// IncreaseSize increases the size of the node group. To delete a node you need
// to explicitly name it and use DeleteNode. This function should wait until
// node group size is updated. Implementation required.
func (n *NodeGroup) IncreaseSize(delta int) error {
	if delta <= 0 {
		return fmt.Errorf("delta must be positive, have: %d", delta)
	}

	targetSize := n.nodePool.Count + delta

	if targetSize > n.MaxSize() {
		return fmt.Errorf("size increase is too large. current: %d desired: %d max: %d",
			n.nodePool.Count, targetSize, n.MaxSize())
	}

	req := &civogo.KubernetesClusterPoolUpdateConfig{
		Count:  targetSize,
		Region: Region,
	}
	updatedNodePool, err := n.client.UpdateKubernetesClusterPool(n.clusterID, n.id, req)
	if err != nil {
		return err
	}

	if updatedNodePool.Count != targetSize {
		return fmt.Errorf("couldn't increase size to %d (delta: %d). Current size is: %d",
			targetSize, delta, updatedNodePool.Count)
	}

	// update internal cache
	n.nodePool.Count = targetSize
	return nil
}

// DeleteNodes deletes nodes from this node group (and also increasing the size
// of the node group with that). Error is returned either on failure or if the
// given node doesn't belong to this node group. This function should wait
// until node group size is updated. Implementation required.
func (n *NodeGroup) DeleteNodes(nodes []*apiv1.Node) error {
	for _, node := range nodes {
		instanceID := toNodeID(node.Spec.ProviderID)
		klog.V(4).Info("deleteing node: %q", instanceID)
		_, err := n.client.DeleteKubernetesClusterPoolInstance(n.clusterID, n.id, instanceID)
		if err != nil {
			return fmt.Errorf("deleting node failed for cluster: %q node pool: %q node: %q: %s",
				n.clusterID, n.id, node.Name, err)
		}

		// decrement the count by one  after a successful delete
		n.nodePool.Count--
	}

	return nil
}

// DecreaseTargetSize decreases the target size of the node group. This function
// doesn't permit to delete any existing node and can be used only to reduce the
// request for new nodes that have not been yet fulfilled. Delta should be negative.
// It is assumed that cloud provider will not delete the existing nodes when there
// is an option to just decrease the target. Implementation required.
func (n *NodeGroup) DecreaseTargetSize(delta int) error {
	if delta >= 0 {
		return fmt.Errorf("delta must be negative, have: %d", delta)
	}

	targetSize := n.nodePool.Count + delta
	if targetSize < n.MinSize() {
		return fmt.Errorf("size decrease is too small. current: %d desired: %d min: %d",
			n.nodePool.Count, targetSize, n.MinSize())
	}

	req := &civogo.KubernetesClusterPoolUpdateConfig{
		Count:  targetSize,
		Region: Region,
	}

	updatedNodePool, err := n.client.UpdateKubernetesClusterPool(n.clusterID, n.id, req)
	if err != nil {
		return err
	}

	if updatedNodePool.Count != targetSize {
		return fmt.Errorf("couldn't increase size to %d (delta: %d). Current size is: %d",
			targetSize, delta, updatedNodePool.Count)
	}

	// update internal cache
	n.nodePool.Count = targetSize
	return nil
}

// Id returns an unique identifier of the node group.
func (n *NodeGroup) Id() string {
	return n.id
}

// Debug returns a string containing all information regarding this node group.
func (n *NodeGroup) Debug() string {
	return fmt.Sprintf("cluster ID: %s (min:%d max:%d)", n.Id(), n.MinSize(), n.MaxSize())
}

// Nodes returns a list of all nodes that belong to this node group.  It is
// required that Instance objects returned by this method have Id field set.
// Other fields are optional.
func (n *NodeGroup) Nodes() ([]cloudprovider.Instance, error) {
	if n.nodePool == nil {
		return nil, errors.New("node pool instance is not created")
	}

	return toInstances(n.nodePool.Instances), nil
}

// TemplateNodeInfo returns a schedulernodeinfo.NodeInfo structure of an empty
// (as if just started) node. This will be used in scale-up simulations to
// predict what would a new node look like if a node group was expanded. The
// returned NodeInfo is expected to have a fully populated Node object, with
// all of the labels, capacity and allocatable information as well as all pods
// that are started on the node by default, using manifest (most likely only
// kube-proxy). Implementation optional.
func (n *NodeGroup) TemplateNodeInfo() (*schedulerframework.NodeInfo, error) {
	return nil, cloudprovider.ErrNotImplemented
}

// Exist checks if the node group really exists on the cloud provider side.
// Allows to tell the theoretical node group from the real one. Implementation
// required.
func (n *NodeGroup) Exist() bool {
	return n.nodePool != nil
}

// Create creates the node group on the cloud provider side. Implementation
// optional.
func (n *NodeGroup) Create() (cloudprovider.NodeGroup, error) {
	return nil, cloudprovider.ErrNotImplemented
}

// Delete deletes the node group on the cloud provider side.  This will be
// executed only for autoprovisioned node groups, once their size drops to 0.
// Implementation optional.
func (n *NodeGroup) Delete() error {
	return cloudprovider.ErrNotImplemented
}

// Autoprovisioned returns true if the node group is autoprovisioned. An
// autoprovisioned group was created by CA and can be deleted when scaled to 0.
func (n *NodeGroup) Autoprovisioned() bool {
	return false
}

// toInstances converts a slice of civogo.KubernetesInstance to
// cloudprovider.Instance
func toInstances(nodes []civogo.KubernetesInstance) []cloudprovider.Instance {
	instances := make([]cloudprovider.Instance, 0, len(nodes))
	for _, nd := range nodes {
		instances = append(instances, toInstance(nd))
	}
	return instances
}

// toInstance converts the given civogo.KubernetesInstance to a
// cloudprovider.Instance
func toInstance(node civogo.KubernetesInstance) cloudprovider.Instance {
	return cloudprovider.Instance{
		Id:     toProviderID(node.ID),
		Status: toInstanceStatus(node.Status),
	}
}

// toInstanceStatus converts the given civo instance status to a
// cloudprovider.InstanceStatus
func toInstanceStatus(nodeState string) *cloudprovider.InstanceStatus {
	if nodeState == "" {
		return nil
	}

	st := &cloudprovider.InstanceStatus{}
	switch nodeState {
	case "BUILDING":
		st.State = cloudprovider.InstanceCreating
	case "ACTIVE":
		st.State = cloudprovider.InstanceRunning
	case "DELETING":
		st.State = cloudprovider.InstanceDeleting
	default:
		st.ErrorInfo = &cloudprovider.InstanceErrorInfo{
			ErrorClass:   cloudprovider.OtherErrorClass,
			ErrorCode:    "no-code-civo",
			ErrorMessage: nodeState,
		}
	}

	return st
}
