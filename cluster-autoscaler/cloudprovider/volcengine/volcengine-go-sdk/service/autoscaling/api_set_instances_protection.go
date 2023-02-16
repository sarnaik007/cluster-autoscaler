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

package autoscaling

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/response"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opSetInstancesProtectionCommon = "SetInstancesProtection"

// SetInstancesProtectionCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the SetInstancesProtectionCommon operation. The "output" return
// value will be populated with the SetInstancesProtectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetInstancesProtectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetInstancesProtectionCommon Send returns without error.
//
// See SetInstancesProtectionCommon for more information on using the SetInstancesProtectionCommon
// API call, and error handling.
//
//    // Example sending a request using the SetInstancesProtectionCommonRequest method.
//    req, resp := client.SetInstancesProtectionCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) SetInstancesProtectionCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opSetInstancesProtectionCommon,
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

// SetInstancesProtectionCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation SetInstancesProtectionCommon for usage and error information.
func (c *AUTOSCALING) SetInstancesProtectionCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.SetInstancesProtectionCommonRequest(input)
	return out, req.Send()
}

// SetInstancesProtectionCommonWithContext is the same as SetInstancesProtectionCommon with the addition of
// the ability to pass a context and additional request options.
//
// See SetInstancesProtectionCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) SetInstancesProtectionCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.SetInstancesProtectionCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opSetInstancesProtection = "SetInstancesProtection"

// SetInstancesProtectionRequest generates a "volcengine/request.Request" representing the
// client's request for the SetInstancesProtection operation. The "output" return
// value will be populated with the SetInstancesProtectionCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned SetInstancesProtectionCommon Request to send the API call to the service.
// the "output" return value is not valid until after SetInstancesProtectionCommon Send returns without error.
//
// See SetInstancesProtection for more information on using the SetInstancesProtection
// API call, and error handling.
//
//    // Example sending a request using the SetInstancesProtectionRequest method.
//    req, resp := client.SetInstancesProtectionRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) SetInstancesProtectionRequest(input *SetInstancesProtectionInput) (req *request.Request, output *SetInstancesProtectionOutput) {
	op := &request.Operation{
		Name:       opSetInstancesProtection,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &SetInstancesProtectionInput{}
	}

	output = &SetInstancesProtectionOutput{}
	req = c.newRequest(op, input, output)

	return
}

// SetInstancesProtection API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation SetInstancesProtection for usage and error information.
func (c *AUTOSCALING) SetInstancesProtection(input *SetInstancesProtectionInput) (*SetInstancesProtectionOutput, error) {
	req, out := c.SetInstancesProtectionRequest(input)
	return out, req.Send()
}

// SetInstancesProtectionWithContext is the same as SetInstancesProtection with the addition of
// the ability to pass a context and additional request options.
//
// See SetInstancesProtection for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) SetInstancesProtectionWithContext(ctx volcengine.Context, input *SetInstancesProtectionInput, opts ...request.Option) (*SetInstancesProtectionOutput, error) {
	req, out := c.SetInstancesProtectionRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type InstanceProtectionResultForSetInstancesProtectionOutput struct {
	_ struct{} `type:"structure"`

	Code *string `type:"string"`

	InstanceId *string `type:"string"`

	Message *string `type:"string"`

	Result *string `type:"string"`
}

// String returns the string representation
func (s InstanceProtectionResultForSetInstancesProtectionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s InstanceProtectionResultForSetInstancesProtectionOutput) GoString() string {
	return s.String()
}

// SetCode sets the Code field's value.
func (s *InstanceProtectionResultForSetInstancesProtectionOutput) SetCode(v string) *InstanceProtectionResultForSetInstancesProtectionOutput {
	s.Code = &v
	return s
}

// SetInstanceId sets the InstanceId field's value.
func (s *InstanceProtectionResultForSetInstancesProtectionOutput) SetInstanceId(v string) *InstanceProtectionResultForSetInstancesProtectionOutput {
	s.InstanceId = &v
	return s
}

// SetMessage sets the Message field's value.
func (s *InstanceProtectionResultForSetInstancesProtectionOutput) SetMessage(v string) *InstanceProtectionResultForSetInstancesProtectionOutput {
	s.Message = &v
	return s
}

// SetResult sets the Result field's value.
func (s *InstanceProtectionResultForSetInstancesProtectionOutput) SetResult(v string) *InstanceProtectionResultForSetInstancesProtectionOutput {
	s.Result = &v
	return s
}

type SetInstancesProtectionInput struct {
	_ struct{} `type:"structure"`

	InstanceIds []*string `type:"list"`

	ProtectedFromScaleIn *bool `type:"boolean"`

	ScalingGroupId *string `type:"string"`
}

// String returns the string representation
func (s SetInstancesProtectionInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetInstancesProtectionInput) GoString() string {
	return s.String()
}

// SetInstanceIds sets the InstanceIds field's value.
func (s *SetInstancesProtectionInput) SetInstanceIds(v []*string) *SetInstancesProtectionInput {
	s.InstanceIds = v
	return s
}

// SetProtectedFromScaleIn sets the ProtectedFromScaleIn field's value.
func (s *SetInstancesProtectionInput) SetProtectedFromScaleIn(v bool) *SetInstancesProtectionInput {
	s.ProtectedFromScaleIn = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *SetInstancesProtectionInput) SetScalingGroupId(v string) *SetInstancesProtectionInput {
	s.ScalingGroupId = &v
	return s
}

type SetInstancesProtectionOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	InstanceProtectionResults []*InstanceProtectionResultForSetInstancesProtectionOutput `type:"list"`
}

// String returns the string representation
func (s SetInstancesProtectionOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s SetInstancesProtectionOutput) GoString() string {
	return s.String()
}

// SetInstanceProtectionResults sets the InstanceProtectionResults field's value.
func (s *SetInstancesProtectionOutput) SetInstanceProtectionResults(v []*InstanceProtectionResultForSetInstancesProtectionOutput) *SetInstancesProtectionOutput {
	s.InstanceProtectionResults = v
	return s
}
