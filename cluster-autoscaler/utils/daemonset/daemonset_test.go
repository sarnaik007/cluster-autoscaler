/*
Copyright 2017 The Kubernetes Authors.

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

package daemonset

import (
	"strings"
	"testing"
	"time"

	"k8s.io/autoscaler/cluster-autoscaler/simulator"
	. "k8s.io/autoscaler/cluster-autoscaler/utils/test"

	appsv1 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	schedulerframework "k8s.io/kubernetes/pkg/scheduler/framework"

	"github.com/stretchr/testify/assert"
)

func TestGetDaemonSetPodsForNode(t *testing.T) {
	node := BuildTestNode("node", 1000, 1000)
	SetNodeReadyState(node, true, time.Now())
	nodeInfo := schedulerframework.NewNodeInfo()
	nodeInfo.SetNode(node)

	predicateChecker, err := simulator.NewTestPredicateChecker()
	assert.NoError(t, err)
	ds1 := newDaemonSet("ds1")
	ds2 := newDaemonSet("ds2")
	ds2.Spec.Template.Spec.NodeSelector = map[string]string{"foo": "bar"}

	{
		daemonSets, err := GetDaemonSetPodsForNode(nodeInfo, []*appsv1.DaemonSet{ds1, ds2}, predicateChecker)

		assert.NoError(t, err)
		assert.Equal(t, 1, len(daemonSets))
		assert.True(t, strings.HasPrefix(daemonSets[0].Name, "ds1"))
	}
	{
		daemonSets, err := GetDaemonSetPodsForNode(nodeInfo, []*appsv1.DaemonSet{ds1}, predicateChecker)
		assert.NoError(t, err)
		assert.Equal(t, 1, len(daemonSets))
	}
	{
		daemonSets, err := GetDaemonSetPodsForNode(nodeInfo, []*appsv1.DaemonSet{ds2}, predicateChecker)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(daemonSets))
	}
	{
		daemonSets, err := GetDaemonSetPodsForNode(nodeInfo, []*appsv1.DaemonSet{}, predicateChecker)
		assert.NoError(t, err)
		assert.Equal(t, 0, len(daemonSets))
	}
}

func TestEvictedPodsFilter(t *testing.T) {
	testCases := []struct {
		name            string
		pods            map[string]string
		evictionDefault bool
		expectedPods    []string
	}{
		{
			name: "all pods evicted by default",
			pods: map[string]string{
				"p1": "",
				"p2": "",
				"p3": "",
			},
			evictionDefault: true,
			expectedPods:    []string{"p1", "p2", "p3"},
		},
		{
			name: "no pods evicted by default",
			pods: map[string]string{
				"p1": "",
				"p2": "",
				"p3": "",
			},
			evictionDefault: false,
			expectedPods:    []string{},
		},
		{
			name: "all pods evicted by default, one opt-out",
			pods: map[string]string{
				"p1": "",
				"p2": "false",
				"p3": "",
			},
			evictionDefault: true,
			expectedPods:    []string{"p1", "p3"},
		},
		{
			name: "no pods evicted by default, one opt-in",
			pods: map[string]string{
				"p1": "",
				"p2": "true",
				"p3": "",
			},
			evictionDefault: false,
			expectedPods:    []string{"p2"},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()
			var dsPods []*apiv1.Pod
			for n, av := range tc.pods {
				p := BuildTestPod(n, 100, 0)
				if av != "" {
					p.Annotations[EnableDsEvictionKey] = av
				}
				dsPods = append(dsPods, p)
			}
			pte := PodsToEvict(dsPods, tc.evictionDefault)
			got := make([]string, len(pte))
			for i, p := range pte {
				got[i] = p.Name
			}
			assert.ElementsMatch(t, got, tc.expectedPods)
		})
	}
}

func newDaemonSet(name string) *appsv1.DaemonSet {
	return &appsv1.DaemonSet{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: metav1.NamespaceDefault,
		},
		Spec: appsv1.DaemonSetSpec{
			Selector: &metav1.LabelSelector{MatchLabels: map[string]string{"name": "simple-daemon", "type": "production"}},
			Template: apiv1.PodTemplateSpec{
				ObjectMeta: metav1.ObjectMeta{
					Labels: map[string]string{"name": "simple-daemon", "type": "production"},
				},
				Spec: apiv1.PodSpec{
					Containers: []apiv1.Container{
						{
							Image: "foo/bar",
						},
					},
				},
			},
		},
	}
}
