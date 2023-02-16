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

const opDeleteScalingPolicyCommon = "DeleteScalingPolicy"

// DeleteScalingPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteScalingPolicyCommon operation. The "output" return
// value will be populated with the DeleteScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteScalingPolicyCommon Send returns without error.
//
// See DeleteScalingPolicyCommon for more information on using the DeleteScalingPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the DeleteScalingPolicyCommonRequest method.
//    req, resp := client.DeleteScalingPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DeleteScalingPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteScalingPolicyCommon,
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

// DeleteScalingPolicyCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteScalingPolicyCommon for usage and error information.
func (c *AUTOSCALING) DeleteScalingPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteScalingPolicyCommonRequest(input)
	return out, req.Send()
}

// DeleteScalingPolicyCommonWithContext is the same as DeleteScalingPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteScalingPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteScalingPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteScalingPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteScalingPolicy = "DeleteScalingPolicy"

// DeleteScalingPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteScalingPolicy operation. The "output" return
// value will be populated with the DeleteScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteScalingPolicyCommon Send returns without error.
//
// See DeleteScalingPolicy for more information on using the DeleteScalingPolicy
// API call, and error handling.
//
//    // Example sending a request using the DeleteScalingPolicyRequest method.
//    req, resp := client.DeleteScalingPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) DeleteScalingPolicyRequest(input *DeleteScalingPolicyInput) (req *request.Request, output *DeleteScalingPolicyOutput) {
	op := &request.Operation{
		Name:       opDeleteScalingPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteScalingPolicyInput{}
	}

	output = &DeleteScalingPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteScalingPolicy API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteScalingPolicy for usage and error information.
func (c *AUTOSCALING) DeleteScalingPolicy(input *DeleteScalingPolicyInput) (*DeleteScalingPolicyOutput, error) {
	req, out := c.DeleteScalingPolicyRequest(input)
	return out, req.Send()
}

// DeleteScalingPolicyWithContext is the same as DeleteScalingPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteScalingPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteScalingPolicyWithContext(ctx volcengine.Context, input *DeleteScalingPolicyInput, opts ...request.Option) (*DeleteScalingPolicyOutput, error) {
	req, out := c.DeleteScalingPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	ScalingPolicyId *string `type:"string"`
}

// String returns the string representation
func (s DeleteScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteScalingPolicyInput) GoString() string {
	return s.String()
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *DeleteScalingPolicyInput) SetScalingPolicyId(v string) *DeleteScalingPolicyInput {
	s.ScalingPolicyId = &v
	return s
}

type DeleteScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingPolicyId *string `type:"string"`
}

// String returns the string representation
func (s DeleteScalingPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteScalingPolicyOutput) GoString() string {
	return s.String()
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *DeleteScalingPolicyOutput) SetScalingPolicyId(v string) *DeleteScalingPolicyOutput {
	s.ScalingPolicyId = &v
	return s
}
