/*
Copyright (c) 2023 SAP SE or an SAP affiliate company. All rights reserved.

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

package mcm

import (
	"context"
	"errors"
	"fmt"
	"github.com/gardener/machine-controller-manager/pkg/apis/machine/v1alpha1"
	. "github.com/onsi/gomega"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	customfake "k8s.io/autoscaler/cluster-autoscaler/cloudprovider/mcm/fakeclient"
	"math"
	"testing"
)

const (
	mdUpdateErrorMsg = "unable to update machine deployment"
	mcUpdateErrorMsg = "unable to update machine"
)

var (
	nodeGroup1 = "1:3:" + testNamespace + ".machinedeployment-1"
	nodeGroup2 = "0:1:" + testNamespace + ".machinedeployment-1"
	nodeGroup3 = "0:2:" + testNamespace + ".machinedeployment-2"
)

type setup struct {
	nodes                             []*corev1.Node
	machines                          []*v1alpha1.Machine
	machineSets                       []*v1alpha1.MachineSet
	machineDeployments                []*v1alpha1.MachineDeployment
	machineClasses                    []*v1alpha1.MachineClass
	nodeGroups                        []string
	targetCoreFakeResourceActions     *customfake.ResourceActions
	controlMachineFakeResourceActions *customfake.ResourceActions
}

func setupEnv(setup *setup) ([]runtime.Object, []runtime.Object) {
	var controlMachineObjects []runtime.Object
	for _, o := range setup.machines {
		controlMachineObjects = append(controlMachineObjects, o)
	}
	for _, o := range setup.machineSets {
		controlMachineObjects = append(controlMachineObjects, o)
	}
	for _, o := range setup.machineDeployments {
		controlMachineObjects = append(controlMachineObjects, o)
	}
	for _, o := range setup.machineClasses {
		controlMachineObjects = append(controlMachineObjects, o)
	}

	var targetCoreObjects []runtime.Object

	for _, o := range setup.nodes {
		targetCoreObjects = append(targetCoreObjects, o)
	}

	return controlMachineObjects, targetCoreObjects
}

func TestDeleteNodes(t *testing.T) {
	type action struct {
		node *corev1.Node
	}
	type expect struct {
		machines   []*v1alpha1.Machine
		mdName     string
		mdReplicas int32
		err        error
	}
	type data struct {
		name   string
		setup  setup
		action action
		expect expect
	}
	table := []data{
		{
			"should scale down machine deployment to remove a node",
			setup{
				nodes:              newNodes(2, "fakeID", []bool{true, false}),
				machines:           newMachines(2, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3", "3"}, []bool{false, false}),
				machineSets:        newMachineSets(1, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 2, nil, nil, nil),
				nodeGroups:         []string{nodeGroup1},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				mdName:     "machinedeployment-1",
				mdReplicas: 1,
				err:        nil,
			},
		},
		{
			"should scale down machine deployment to remove a placeholder node",
			setup{
				nodes:              nil,
				machines:           newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3"}, []bool{false}),
				machineSets:        newMachineSets(1, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 1, nil, nil, nil),
				nodeGroups:         []string{nodeGroup2},
			},
			action{node: newNode("node-1", "requested://machine-1", true)},
			expect{
				machines:   newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				mdName:     "machinedeployment-1",
				mdReplicas: 0,
				err:        nil,
			},
		},
		{
			"should not scale down a machine deployment when it is under rolling update",
			setup{
				nodes:       newNodes(2, "fakeID", []bool{true, false}),
				machines:    newMachines(2, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3", "3"}, []bool{false, false}),
				machineSets: newMachineSets(2, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 2, &v1alpha1.MachineDeploymentStatus{
					Conditions: []v1alpha1.MachineDeploymentCondition{
						{Type: "Progressing"},
					},
				}, nil, nil),
				nodeGroups: []string{nodeGroup1},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   nil,
				mdName:     "machinedeployment-1",
				mdReplicas: 2,
				err:        fmt.Errorf("MachineDeployment machinedeployment-1 is under rolling update , cannot reduce replica count"),
			},
		},
		{
			"should not scale down when machine deployment update call times out",
			setup{
				nodes:              newNodes(2, "fakeID", []bool{true, false}),
				machines:           newMachines(2, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3", "3"}, []bool{false, false}),
				machineSets:        newMachineSets(1, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 2, nil, nil, nil),
				nodeGroups:         []string{nodeGroup1},
				controlMachineFakeResourceActions: &customfake.ResourceActions{
					MachineDeployment: customfake.Actions{
						Update: customfake.CreateFakeResponse(math.MaxInt32, mdUpdateErrorMsg, 0),
					},
				},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				mdName:     "machinedeployment-1",
				mdReplicas: 2,
				err:        fmt.Errorf("unable to scale in machine deployment machinedeployment-1, Error: %v", mdUpdateErrorMsg),
			},
		},
		{
			"should scale down when machine deployment update call fails but passes within the timeout period",
			setup{
				nodes:              newNodes(2, "fakeID", []bool{true, false}),
				machines:           newMachines(2, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3", "3"}, []bool{false, false}),
				machineSets:        newMachineSets(1, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 2, nil, nil, nil),
				nodeGroups:         []string{nodeGroup1},
				controlMachineFakeResourceActions: &customfake.ResourceActions{
					MachineDeployment: customfake.Actions{
						Update: customfake.CreateFakeResponse(2, mdUpdateErrorMsg, 0),
					},
				},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				mdName:     "machinedeployment-1",
				mdReplicas: 1,
				err:        nil,
			},
		},
		{
			"should not scale down a machine deployment when the corresponding machine is already in terminating state",
			setup{
				nodes:              newNodes(2, "fakeID", []bool{true, false}),
				machines:           newMachines(2, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3", "3"}, []bool{true, false}),
				machineSets:        newMachineSets(1, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 2, nil, nil, nil),
				nodeGroups:         []string{nodeGroup1},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3"}, []bool{true}),
				mdName:     "machinedeployment-1",
				mdReplicas: 2,
				err:        nil,
			},
		},
		{
			"should not scale down a machine deployment below the minimum",
			setup{
				nodes:              newNodes(1, "fakeID", []bool{true}),
				machines:           newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3"}, []bool{false}),
				machineSets:        newMachineSets(1, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 1, nil, nil, nil),
				nodeGroups:         []string{nodeGroup1},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   nil,
				mdName:     "machinedeployment-1",
				mdReplicas: 1,
				err:        fmt.Errorf("min size reached, nodes will not be deleted"),
			},
		},
		{
			"no scale down of machine deployment if priority of the targeted machine cannot be updated to 1",
			setup{
				nodes:              newNodes(2, "fakeID", []bool{true, false}),
				machines:           newMachines(2, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3", "3"}, []bool{false, false}),
				machineSets:        newMachineSets(1, "machinedeployment-1"),
				machineDeployments: newMachineDeployments(1, 2, nil, nil, nil),
				nodeGroups:         []string{nodeGroup1},
				controlMachineFakeResourceActions: &customfake.ResourceActions{
					Machine: customfake.Actions{
						Update: customfake.CreateFakeResponse(math.MaxInt32, mcUpdateErrorMsg, 0),
					},
				},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   nil,
				mdName:     "machinedeployment-1",
				mdReplicas: 2,
				err:        fmt.Errorf("could not prioritize machine machine-1 for deletion, aborting scale in of machine deployment, Error: %s", mcUpdateErrorMsg),
			},
		},
		{
			"should not scale down machine deployment if the node belongs to another machine deployment",
			setup{
				nodes:              newNodes(2, "fakeID", []bool{true, false}),
				machines:           newMachines(2, "fakeID", nil, "machinedeployment-2", "machineset-1", []string{"3", "3"}, []bool{false, false}),
				machineSets:        newMachineSets(1, "machinedeployment-2"),
				machineDeployments: newMachineDeployments(2, 2, nil, nil, nil),
				nodeGroups:         []string{nodeGroup2, nodeGroup3},
			},
			action{node: newNodes(1, "fakeID", []bool{true})[0]},
			expect{
				machines:   nil,
				mdName:     "machinedeployment-2",
				mdReplicas: 2,
				err:        fmt.Errorf("node-1 belongs to a different machinedeployment than machinedeployment-1"),
			},
		},
	}

	for _, entry := range table {
		entry := entry // have a shallow copy of the entry for parallelization of tests
		t.Run(entry.name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			stop := make(chan struct{})
			defer close(stop)
			controlMachineObjects, targetCoreObjects := setupEnv(&entry.setup)
			m, trackers, hasSyncedCacheFns := createMcmManager(t, stop, testNamespace, nil, controlMachineObjects, targetCoreObjects)
			defer trackers.Stop()
			waitForCacheSync(t, stop, hasSyncedCacheFns)

			if entry.setup.targetCoreFakeResourceActions != nil {
				trackers.TargetCore.SetFailAtFakeResourceActions(entry.setup.targetCoreFakeResourceActions)
			}
			if entry.setup.controlMachineFakeResourceActions != nil {
				trackers.ControlMachine.SetFailAtFakeResourceActions(entry.setup.controlMachineFakeResourceActions)
			}

			md, err := buildMachineDeploymentFromSpec(entry.setup.nodeGroups[0], m)
			g.Expect(err).To(BeNil())

			err = md.DeleteNodes([]*corev1.Node{entry.action.node})

			if entry.expect.err != nil {
				g.Expect(err).To(Equal(entry.expect.err))
			} else {
				g.Expect(err).To(BeNil())
			}

			machineDeployment, err := m.machineClient.MachineDeployments(m.namespace).Get(context.TODO(), entry.expect.mdName, metav1.GetOptions{})
			g.Expect(err).ToNot(HaveOccurred())
			g.Expect(machineDeployment.Spec.Replicas).To(BeNumerically("==", entry.expect.mdReplicas))

			machines, err := m.machineClient.Machines(m.namespace).List(context.TODO(), metav1.ListOptions{
				LabelSelector: metav1.FormatLabelSelector(&metav1.LabelSelector{
					MatchLabels: map[string]string{"name": md.Name},
				}),
			})

			for _, machine := range machines.Items {
				flag := false
				for _, entryMachineItem := range entry.expect.machines {
					if entryMachineItem.Name == machine.Name {
						g.Expect(machine.Annotations[priorityAnnotationKey]).To(Equal(entryMachineItem.Annotations[priorityAnnotationKey]))
						flag = true
						break
					}
				}
				if !flag {
					g.Expect(machine.Annotations[priorityAnnotationKey]).To(Equal("3"))
				}
			}
		})
	}
}

func TestRefresh(t *testing.T) {
	type expect struct {
		machines []*v1alpha1.Machine
		err      error
	}
	type data struct {
		name   string
		setup  setup
		expect expect
	}
	table := []data{
		{
			"should reset priority of a machine with node without ToBeDeletedTaint to 3",
			setup{
				nodes:              newNodes(1, "fakeID", []bool{false}),
				machines:           newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				machineDeployments: newMachineDeployments(1, 1, nil, nil, nil),
				nodeGroups:         []string{nodeGroup2},
			},
			expect{
				machines: newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"3"}, []bool{false}),
				err:      nil,
			},
		},
		{
			"should not reset priority of a machine to 3 if the node has ToBeDeleted taint",
			setup{
				nodes:              newNodes(1, "fakeID", []bool{true}),
				machines:           newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				machineDeployments: newMachineDeployments(1, 1, nil, nil, nil),
				nodeGroups:         []string{nodeGroup2},
			},
			expect{
				machines: newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				err:      nil,
			},
		},
		{
			"priority reset of machine fails",
			setup{
				nodes:              newNodes(1, "fakeID", []bool{false}),
				machines:           newMachines(1, "fakeID", nil, "machinedeployment-1", "machineset-1", []string{"1"}, []bool{false}),
				machineDeployments: newMachineDeployments(1, 1, nil, nil, nil),
				controlMachineFakeResourceActions: &customfake.ResourceActions{
					Machine: customfake.Actions{
						Update: customfake.CreateFakeResponse(math.MaxInt32, mcUpdateErrorMsg, 0),
					},
				},
				nodeGroups: []string{nodeGroup2},
			},
			expect{
				machines: []*v1alpha1.Machine{newMachine("machine-1", "fakeID", nil, "machinedeployment-1", "machineset-1", "1", false)},
				err:      errors.Join(fmt.Errorf("could not reset priority annotation on machine machine-1, Error: %v", mcUpdateErrorMsg)),
			},
		},
	}
	for _, entry := range table {
		entry := entry
		t.Run(entry.name, func(t *testing.T) {
			t.Parallel()
			g := NewWithT(t)
			stop := make(chan struct{})
			defer close(stop)
			controlMachineObjects, targetCoreObjects := setupEnv(&entry.setup)
			m, trackers, hasSyncedCacheFns := createMcmManager(t, stop, testNamespace, entry.setup.nodeGroups, controlMachineObjects, targetCoreObjects)
			defer trackers.Stop()
			waitForCacheSync(t, stop, hasSyncedCacheFns)

			if entry.setup.targetCoreFakeResourceActions != nil {
				trackers.TargetCore.SetFailAtFakeResourceActions(entry.setup.targetCoreFakeResourceActions)
			}
			if entry.setup.controlMachineFakeResourceActions != nil {
				trackers.ControlMachine.SetFailAtFakeResourceActions(entry.setup.controlMachineFakeResourceActions)
			}
			mcmCloudProvider, err := BuildMcmCloudProvider(m, nil)
			g.Expect(err).To(BeNil())
			err = mcmCloudProvider.Refresh()
			if entry.expect.err != nil {
				g.Expect(err).To(Equal(entry.expect.err))
			} else {
				g.Expect(err).To(BeNil())
			}
			for _, mc := range entry.expect.machines {
				machine, err := m.machineClient.Machines(m.namespace).Get(context.TODO(), mc.Name, metav1.GetOptions{})
				g.Expect(err).To(BeNil())
				g.Expect(mc.Annotations[priorityAnnotationKey]).To(Equal(machine.Annotations[priorityAnnotationKey]))
			}
		})
	}
}
