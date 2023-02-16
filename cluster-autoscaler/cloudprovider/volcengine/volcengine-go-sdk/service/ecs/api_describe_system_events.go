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

// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package ecs

import (
	"encoding/json"

	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/response"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDescribeSystemEventsCommon = "DescribeSystemEvents"

// DescribeSystemEventsCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSystemEventsCommon operation. The "output" return
// value will be populated with the DescribeSystemEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSystemEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSystemEventsCommon Send returns without error.
//
// See DescribeSystemEventsCommon for more information on using the DescribeSystemEventsCommon
// API call, and error handling.
//
//    // Example sending a request using the DescribeSystemEventsCommonRequest method.
//    req, resp := client.DescribeSystemEventsCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeSystemEventsCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDescribeSystemEventsCommon,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &map[string]interface{}{}
	}

	output = &map[string]interface{}{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSystemEventsCommon API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeSystemEventsCommon for usage and error information.
func (c *ECS) DescribeSystemEventsCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DescribeSystemEventsCommonRequest(input)
	return out, req.Send()
}

// DescribeSystemEventsCommonWithContext is the same as DescribeSystemEventsCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSystemEventsCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeSystemEventsCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DescribeSystemEventsCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDescribeSystemEvents = "DescribeSystemEvents"

// DescribeSystemEventsRequest generates a "volcengine/request.Request" representing the
// client's request for the DescribeSystemEvents operation. The "output" return
// value will be populated with the DescribeSystemEventsCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DescribeSystemEventsCommon Request to send the API call to the service.
// the "output" return value is not valid until after DescribeSystemEventsCommon Send returns without error.
//
// See DescribeSystemEvents for more information on using the DescribeSystemEvents
// API call, and error handling.
//
//    // Example sending a request using the DescribeSystemEventsRequest method.
//    req, resp := client.DescribeSystemEventsRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *ECS) DescribeSystemEventsRequest(input *DescribeSystemEventsInput) (req *request.Request, output *DescribeSystemEventsOutput) {
	op := &request.Operation{
		Name:       opDescribeSystemEvents,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DescribeSystemEventsInput{}
	}

	output = &DescribeSystemEventsOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DescribeSystemEvents API operation for ECS.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for ECS's
// API operation DescribeSystemEvents for usage and error information.
func (c *ECS) DescribeSystemEvents(input *DescribeSystemEventsInput) (*DescribeSystemEventsOutput, error) {
	req, out := c.DescribeSystemEventsRequest(input)
	return out, req.Send()
}

// DescribeSystemEventsWithContext is the same as DescribeSystemEvents with the addition of
// the ability to pass a context and additional request options.
//
// See DescribeSystemEvents for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *ECS) DescribeSystemEventsWithContext(ctx volcengine.Context, input *DescribeSystemEventsInput, opts ...request.Option) (*DescribeSystemEventsOutput, error) {
	req, out := c.DescribeSystemEventsRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DescribeSystemEventsInput struct {
	_ struct{} `type:"structure"`

	CreatedAtEnd *string `type:"string"`

	CreatedAtStart *string `type:"string"`

	EventIds []*string `type:"list"`

	MaxResults *json.Number `type:"json_number"`

	NextToken *string `type:"string"`

	ResourceIds []*string `type:"list"`

	Status []*string `type:"list"`

	Types []*string `type:"list"`
}

// String returns the string representation
func (s DescribeSystemEventsInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSystemEventsInput) GoString() string {
	return s.String()
}

// SetCreatedAtEnd sets the CreatedAtEnd field's value.
func (s *DescribeSystemEventsInput) SetCreatedAtEnd(v string) *DescribeSystemEventsInput {
	s.CreatedAtEnd = &v
	return s
}

// SetCreatedAtStart sets the CreatedAtStart field's value.
func (s *DescribeSystemEventsInput) SetCreatedAtStart(v string) *DescribeSystemEventsInput {
	s.CreatedAtStart = &v
	return s
}

// SetEventIds sets the EventIds field's value.
func (s *DescribeSystemEventsInput) SetEventIds(v []*string) *DescribeSystemEventsInput {
	s.EventIds = v
	return s
}

// SetMaxResults sets the MaxResults field's value.
func (s *DescribeSystemEventsInput) SetMaxResults(v json.Number) *DescribeSystemEventsInput {
	s.MaxResults = &v
	return s
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeSystemEventsInput) SetNextToken(v string) *DescribeSystemEventsInput {
	s.NextToken = &v
	return s
}

// SetResourceIds sets the ResourceIds field's value.
func (s *DescribeSystemEventsInput) SetResourceIds(v []*string) *DescribeSystemEventsInput {
	s.ResourceIds = v
	return s
}

// SetStatus sets the Status field's value.
func (s *DescribeSystemEventsInput) SetStatus(v []*string) *DescribeSystemEventsInput {
	s.Status = v
	return s
}

// SetTypes sets the Types field's value.
func (s *DescribeSystemEventsInput) SetTypes(v []*string) *DescribeSystemEventsInput {
	s.Types = v
	return s
}

type DescribeSystemEventsOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	NextToken *string `type:"string"`

	SystemEvents []*SystemEventForDescribeSystemEventsOutput `type:"list"`
}

// String returns the string representation
func (s DescribeSystemEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DescribeSystemEventsOutput) GoString() string {
	return s.String()
}

// SetNextToken sets the NextToken field's value.
func (s *DescribeSystemEventsOutput) SetNextToken(v string) *DescribeSystemEventsOutput {
	s.NextToken = &v
	return s
}

// SetSystemEvents sets the SystemEvents field's value.
func (s *DescribeSystemEventsOutput) SetSystemEvents(v []*SystemEventForDescribeSystemEventsOutput) *DescribeSystemEventsOutput {
	s.SystemEvents = v
	return s
}

type SystemEventForDescribeSystemEventsOutput struct {
	_ struct{} `type:"structure"`

	CreatedAt *string `type:"string"`

	Id *string `type:"string"`

	OperatedEndAt *string `type:"string"`

	OperatedStartAt *string `type:"string"`

	ResourceId *string `type:"string"`

	Status *string `type:"string" enum:"StatusForDescribeSystemEventsOutput"`

	Type *string `type:"string" enum:"TypeForDescribeSystemEventsOutput"`

	UpdatedAt *string `type:"string"`
}

// String returns the string representation
func (s SystemEventForDescribeSystemEventsOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SystemEventForDescribeSystemEventsOutput) GoString() string {
	return s.String()
}

// SetCreatedAt sets the CreatedAt field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetCreatedAt(v string) *SystemEventForDescribeSystemEventsOutput {
	s.CreatedAt = &v
	return s
}

// SetId sets the Id field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetId(v string) *SystemEventForDescribeSystemEventsOutput {
	s.Id = &v
	return s
}

// SetOperatedEndAt sets the OperatedEndAt field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetOperatedEndAt(v string) *SystemEventForDescribeSystemEventsOutput {
	s.OperatedEndAt = &v
	return s
}

// SetOperatedStartAt sets the OperatedStartAt field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetOperatedStartAt(v string) *SystemEventForDescribeSystemEventsOutput {
	s.OperatedStartAt = &v
	return s
}

// SetResourceId sets the ResourceId field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetResourceId(v string) *SystemEventForDescribeSystemEventsOutput {
	s.ResourceId = &v
	return s
}

// SetStatus sets the Status field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetStatus(v string) *SystemEventForDescribeSystemEventsOutput {
	s.Status = &v
	return s
}

// SetType sets the Type field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetType(v string) *SystemEventForDescribeSystemEventsOutput {
	s.Type = &v
	return s
}

// SetUpdatedAt sets the UpdatedAt field's value.
func (s *SystemEventForDescribeSystemEventsOutput) SetUpdatedAt(v string) *SystemEventForDescribeSystemEventsOutput {
	s.UpdatedAt = &v
	return s
}

const (
	// StatusForDescribeSystemEventsOutputUnknownStatus is a StatusForDescribeSystemEventsOutput enum value
	StatusForDescribeSystemEventsOutputUnknownStatus = "UnknownStatus"

	// StatusForDescribeSystemEventsOutputExecuting is a StatusForDescribeSystemEventsOutput enum value
	StatusForDescribeSystemEventsOutputExecuting = "Executing"

	// StatusForDescribeSystemEventsOutputSucceeded is a StatusForDescribeSystemEventsOutput enum value
	StatusForDescribeSystemEventsOutputSucceeded = "Succeeded"

	// StatusForDescribeSystemEventsOutputFailed is a StatusForDescribeSystemEventsOutput enum value
	StatusForDescribeSystemEventsOutputFailed = "Failed"

	// StatusForDescribeSystemEventsOutputInquiring is a StatusForDescribeSystemEventsOutput enum value
	StatusForDescribeSystemEventsOutputInquiring = "Inquiring"

	// StatusForDescribeSystemEventsOutputScheduled is a StatusForDescribeSystemEventsOutput enum value
	StatusForDescribeSystemEventsOutputScheduled = "Scheduled"

	// StatusForDescribeSystemEventsOutputRejected is a StatusForDescribeSystemEventsOutput enum value
	StatusForDescribeSystemEventsOutputRejected = "Rejected"
)

const (
	// TypeForDescribeSystemEventsOutputUnknownType is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputUnknownType = "UnknownType"

	// TypeForDescribeSystemEventsOutputSystemFailureStop is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputSystemFailureStop = "SystemFailure_Stop"

	// TypeForDescribeSystemEventsOutputSystemFailureReboot is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputSystemFailureReboot = "SystemFailure_Reboot"

	// TypeForDescribeSystemEventsOutputSystemFailurePleaseCheck is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputSystemFailurePleaseCheck = "SystemFailure_PleaseCheck"

	// TypeForDescribeSystemEventsOutputDiskErrorRedeploy is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputDiskErrorRedeploy = "DiskError_Redeploy"

	// TypeForDescribeSystemEventsOutputHddbadSectorRedeploy is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputHddbadSectorRedeploy = "HDDBadSector_Redeploy"

	// TypeForDescribeSystemEventsOutputGpuErrorRedeploy is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputGpuErrorRedeploy = "GpuError_Redeploy"

	// TypeForDescribeSystemEventsOutputSystemMaintenanceRedeploy is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputSystemMaintenanceRedeploy = "SystemMaintenance_Redeploy"

	// TypeForDescribeSystemEventsOutputSystemFailureRedeploy is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputSystemFailureRedeploy = "SystemFailure_Redeploy"

	// TypeForDescribeSystemEventsOutputCreateInstance is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputCreateInstance = "CreateInstance"

	// TypeForDescribeSystemEventsOutputRunInstance is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputRunInstance = "RunInstance"

	// TypeForDescribeSystemEventsOutputStopInstance is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputStopInstance = "StopInstance"

	// TypeForDescribeSystemEventsOutputDeleteInstance is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputDeleteInstance = "DeleteInstance"

	// TypeForDescribeSystemEventsOutputSpotInstanceInterruptionDelete is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputSpotInstanceInterruptionDelete = "SpotInstanceInterruption_Delete"

	// TypeForDescribeSystemEventsOutputAccountUnbalancedStop is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputAccountUnbalancedStop = "AccountUnbalanced_Stop"

	// TypeForDescribeSystemEventsOutputAccountUnbalancedDelete is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputAccountUnbalancedDelete = "AccountUnbalanced_Delete"

	// TypeForDescribeSystemEventsOutputInstanceChargeTypeChange is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputInstanceChargeTypeChange = "InstanceChargeType_Change"

	// TypeForDescribeSystemEventsOutputInstanceConfigurationChange is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputInstanceConfigurationChange = "InstanceConfiguration_Change"

	// TypeForDescribeSystemEventsOutputFileSystemReadOnlyChange is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputFileSystemReadOnlyChange = "FileSystemReadOnly_Change"

	// TypeForDescribeSystemEventsOutputRebootInstance is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputRebootInstance = "RebootInstance"

	// TypeForDescribeSystemEventsOutputInstanceFailure is a TypeForDescribeSystemEventsOutput enum value
	TypeForDescribeSystemEventsOutputInstanceFailure = "InstanceFailure"
)
