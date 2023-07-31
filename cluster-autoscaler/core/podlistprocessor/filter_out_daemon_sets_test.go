/*
Copyright 2023 The Kubernetes Authors.

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

package podlistprocessor

import (
	"testing"

	"github.com/stretchr/testify/assert"
	apiv1 "k8s.io/api/core/v1"
	po "k8s.io/autoscaler/cluster-autoscaler/utils/test/pod"
)

func TestFilterOutDaemonSetPodListProcessor(t *testing.T) {
	testCases := []struct {
		name     string
		pods     []*apiv1.Pod
		wantPods []*apiv1.Pod
	}{
		{
			name: "no pods",
		},
		{
			name: "single non-DS pod",
			pods: []*apiv1.Pod{
				po.BuildTestPod("p", 1000, 1),
			},
			wantPods: []*apiv1.Pod{
				po.BuildTestPod("p", 1000, 1),
			},
		},
		{
			name: "single DS pod",
			pods: []*apiv1.Pod{
				po.SetDSPodSpec(po.BuildTestPod("p", 1000, 1)),
			},
		},
		{
			name: "mixed DS and non-DS pods",
			pods: []*apiv1.Pod{
				po.BuildTestPod("p1", 1000, 1),
				po.SetDSPodSpec(po.BuildTestPod("p2", 1000, 1)),
				po.SetDSPodSpec(po.BuildTestPod("p3", 1000, 1)),
				po.BuildTestPod("p4", 1000, 1),
				po.BuildTestPod("p5", 1000, 1),
				po.SetDSPodSpec(po.BuildTestPod("p6", 1000, 1)),
			},
			wantPods: []*apiv1.Pod{
				po.BuildTestPod("p1", 1000, 1),
				po.BuildTestPod("p4", 1000, 1),
				po.BuildTestPod("p5", 1000, 1),
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			processor := NewFilterOutDaemonSetPodListProcessor()
			pods, err := processor.Process(nil, tc.pods)
			assert.NoError(t, err)
			assert.Equal(t, tc.wantPods, pods)
		})
	}
}
