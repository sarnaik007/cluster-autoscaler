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

package actuation

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	apiv1 "k8s.io/api/core/v1"

	"k8s.io/autoscaler/cluster-autoscaler/utils/kubernetes"
	. "k8s.io/autoscaler/cluster-autoscaler/utils/test"
	no "k8s.io/autoscaler/cluster-autoscaler/utils/test/node"
)

func TestWaitForDelayDeletion(t *testing.T) {
	type testcase struct {
		name                 string
		timeout              time.Duration
		addAnnotation        bool
		removeAnnotation     bool
		expectCallingGetNode bool
	}
	tests := []testcase{
		{
			name:             "annotation not set",
			timeout:          6 * time.Second,
			addAnnotation:    false,
			removeAnnotation: false,
		},
		{
			name:             "annotation set and removed",
			timeout:          6 * time.Second,
			addAnnotation:    true,
			removeAnnotation: true,
		},
		{
			name:             "annotation set but not removed",
			timeout:          6 * time.Second,
			addAnnotation:    true,
			removeAnnotation: false,
		},
		{
			name:             "timeout is 0 - mechanism disable",
			timeout:          0 * time.Second,
			addAnnotation:    true,
			removeAnnotation: false,
		},
	}

	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			node := no.BuildTestNode("n1", 1000, 10)
			nodeWithAnnotation := no.BuildTestNode("n1", 1000, 10)
			nodeWithAnnotation.Annotations = map[string]string{DelayDeletionAnnotationPrefix + "ingress": "true"}
			allNodeLister := kubernetes.NewTestNodeLister(nil)
			if test.addAnnotation {
				if test.removeAnnotation {
					allNodeLister.SetNodes([]*apiv1.Node{node})
				} else {
					allNodeLister.SetNodes([]*apiv1.Node{nodeWithAnnotation})
				}
			}
			var err error
			if test.addAnnotation {
				err = WaitForDelayDeletion(nodeWithAnnotation, allNodeLister, test.timeout)
			} else {
				err = WaitForDelayDeletion(node, allNodeLister, test.timeout)
			}
			assert.NoError(t, err)
		})
	}
}
