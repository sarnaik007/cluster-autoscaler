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

package mocks

import cloudprovider "k8s.io/autoscaler/cluster-autoscaler/cloudprovider"
import errors "k8s.io/autoscaler/cluster-autoscaler/utils/errors"
import mock "github.com/stretchr/testify/mock"
import resource "k8s.io/apimachinery/pkg/api/resource"
import v1 "k8s.io/api/core/v1"

// CloudProvider is an autogenerated mock type for the CloudProvider type
type CloudProvider struct {
	mock.Mock
}

// Cleanup provides a mock function with given fields:
func (_m *CloudProvider) Cleanup() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetAvailableMachineTypes provides a mock function with given fields:
func (_m *CloudProvider) GetAvailableMachineTypes() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetResourceLimiter provides a mock function with given fields:
func (_m *CloudProvider) GetResourceLimiter() (*cloudprovider.ResourceLimiter, error) {
	ret := _m.Called()

	var r0 *cloudprovider.ResourceLimiter
	if rf, ok := ret.Get(0).(func() *cloudprovider.ResourceLimiter); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*cloudprovider.ResourceLimiter)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Name provides a mock function with given fields:
func (_m *CloudProvider) Name() string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}

// NewNodeGroup provides a mock function with given fields: machineType, labels, systemLabels, taints, extraResources
func (_m *CloudProvider) NewNodeGroup(machineType string, labels map[string]string, systemLabels map[string]string, taints []v1.Taint, extraResources map[string]resource.Quantity) (cloudprovider.NodeGroup, error) {
	ret := _m.Called(machineType, labels, systemLabels, taints, extraResources)

	var r0 cloudprovider.NodeGroup
	if rf, ok := ret.Get(0).(func(string, map[string]string, map[string]string, []v1.Taint, map[string]resource.Quantity) cloudprovider.NodeGroup); ok {
		r0 = rf(machineType, labels, systemLabels, taints, extraResources)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cloudprovider.NodeGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, map[string]string, map[string]string, []v1.Taint, map[string]resource.Quantity) error); ok {
		r1 = rf(machineType, labels, systemLabels, taints, extraResources)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeGroupForNode provides a mock function with given fields: _a0
func (_m *CloudProvider) NodeGroupForNode(_a0 *v1.Node) (cloudprovider.NodeGroup, error) {
	ret := _m.Called(_a0)

	var r0 cloudprovider.NodeGroup
	if rf, ok := ret.Get(0).(func(*v1.Node) cloudprovider.NodeGroup); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cloudprovider.NodeGroup)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*v1.Node) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// NodeGroups provides a mock function with given fields:
func (_m *CloudProvider) NodeGroups() []cloudprovider.NodeGroup {
	ret := _m.Called()

	var r0 []cloudprovider.NodeGroup
	if rf, ok := ret.Get(0).(func() []cloudprovider.NodeGroup); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]cloudprovider.NodeGroup)
		}
	}

	return r0
}

// Pricing provides a mock function with given fields:
func (_m *CloudProvider) Pricing() (cloudprovider.PricingModel, errors.AutoscalerError) {
	ret := _m.Called()

	var r0 cloudprovider.PricingModel
	if rf, ok := ret.Get(0).(func() cloudprovider.PricingModel); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(cloudprovider.PricingModel)
		}
	}

	var r1 errors.AutoscalerError
	if rf, ok := ret.Get(1).(func() errors.AutoscalerError); ok {
		r1 = rf()
	} else {
		if ret.Get(1) != nil {
			r1 = ret.Get(1).(errors.AutoscalerError)
		}
	}

	return r0, r1
}

// Refresh provides a mock function with given fields:
func (_m *CloudProvider) Refresh() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetInstanceID gets the instance ID for the specified node.
func (_m *CloudProvider) GetInstanceID(node *v1.Node) string {
	ret := _m.Called()

	var r0 string
	if rf, ok := ret.Get(0).(func() string); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(string)
	}

	return r0
}
