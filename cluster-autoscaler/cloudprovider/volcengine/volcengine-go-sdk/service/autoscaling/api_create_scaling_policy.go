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

const opCreateScalingPolicyCommon = "CreateScalingPolicy"

// CreateScalingPolicyCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateScalingPolicyCommon operation. The "output" return
// value will be populated with the CreateScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateScalingPolicyCommon Send returns without error.
//
// See CreateScalingPolicyCommon for more information on using the CreateScalingPolicyCommon
// API call, and error handling.
//
//    // Example sending a request using the CreateScalingPolicyCommonRequest method.
//    req, resp := client.CreateScalingPolicyCommonRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) CreateScalingPolicyCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opCreateScalingPolicyCommon,
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

// CreateScalingPolicyCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation CreateScalingPolicyCommon for usage and error information.
func (c *AUTOSCALING) CreateScalingPolicyCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.CreateScalingPolicyCommonRequest(input)
	return out, req.Send()
}

// CreateScalingPolicyCommonWithContext is the same as CreateScalingPolicyCommon with the addition of
// the ability to pass a context and additional request options.
//
// See CreateScalingPolicyCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) CreateScalingPolicyCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.CreateScalingPolicyCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opCreateScalingPolicy = "CreateScalingPolicy"

// CreateScalingPolicyRequest generates a "volcengine/request.Request" representing the
// client's request for the CreateScalingPolicy operation. The "output" return
// value will be populated with the CreateScalingPolicyCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned CreateScalingPolicyCommon Request to send the API call to the service.
// the "output" return value is not valid until after CreateScalingPolicyCommon Send returns without error.
//
// See CreateScalingPolicy for more information on using the CreateScalingPolicy
// API call, and error handling.
//
//    // Example sending a request using the CreateScalingPolicyRequest method.
//    req, resp := client.CreateScalingPolicyRequest(params)
//
//    err := req.Send()
//    if err == nil { // resp is now filled
//        fmt.Println(resp)
//    }
func (c *AUTOSCALING) CreateScalingPolicyRequest(input *CreateScalingPolicyInput) (req *request.Request, output *CreateScalingPolicyOutput) {
	op := &request.Operation{
		Name:       opCreateScalingPolicy,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &CreateScalingPolicyInput{}
	}

	output = &CreateScalingPolicyOutput{}
	req = c.newRequest(op, input, output)

	return
}

// CreateScalingPolicy API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation CreateScalingPolicy for usage and error information.
func (c *AUTOSCALING) CreateScalingPolicy(input *CreateScalingPolicyInput) (*CreateScalingPolicyOutput, error) {
	req, out := c.CreateScalingPolicyRequest(input)
	return out, req.Send()
}

// CreateScalingPolicyWithContext is the same as CreateScalingPolicy with the addition of
// the ability to pass a context and additional request options.
//
// See CreateScalingPolicy for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) CreateScalingPolicyWithContext(ctx volcengine.Context, input *CreateScalingPolicyInput, opts ...request.Option) (*CreateScalingPolicyOutput, error) {
	req, out := c.CreateScalingPolicyRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type AlarmPolicyConditionForCreateScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	ComparisonOperator *string `type:"string"`

	MetricName *string `type:"string"`

	MetricUnit *string `type:"string"`

	Threshold *string `type:"string"`
}

// String returns the string representation
func (s AlarmPolicyConditionForCreateScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AlarmPolicyConditionForCreateScalingPolicyInput) GoString() string {
	return s.String()
}

// SetComparisonOperator sets the ComparisonOperator field's value.
func (s *AlarmPolicyConditionForCreateScalingPolicyInput) SetComparisonOperator(v string) *AlarmPolicyConditionForCreateScalingPolicyInput {
	s.ComparisonOperator = &v
	return s
}

// SetMetricName sets the MetricName field's value.
func (s *AlarmPolicyConditionForCreateScalingPolicyInput) SetMetricName(v string) *AlarmPolicyConditionForCreateScalingPolicyInput {
	s.MetricName = &v
	return s
}

// SetMetricUnit sets the MetricUnit field's value.
func (s *AlarmPolicyConditionForCreateScalingPolicyInput) SetMetricUnit(v string) *AlarmPolicyConditionForCreateScalingPolicyInput {
	s.MetricUnit = &v
	return s
}

// SetThreshold sets the Threshold field's value.
func (s *AlarmPolicyConditionForCreateScalingPolicyInput) SetThreshold(v string) *AlarmPolicyConditionForCreateScalingPolicyInput {
	s.Threshold = &v
	return s
}

type AlarmPolicyForCreateScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	Condition *AlarmPolicyConditionForCreateScalingPolicyInput `type:"structure"`

	EvaluationCount *int32 `type:"int32"`

	RuleType *string `type:"string"`
}

// String returns the string representation
func (s AlarmPolicyForCreateScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s AlarmPolicyForCreateScalingPolicyInput) GoString() string {
	return s.String()
}

// SetCondition sets the Condition field's value.
func (s *AlarmPolicyForCreateScalingPolicyInput) SetCondition(v *AlarmPolicyConditionForCreateScalingPolicyInput) *AlarmPolicyForCreateScalingPolicyInput {
	s.Condition = v
	return s
}

// SetEvaluationCount sets the EvaluationCount field's value.
func (s *AlarmPolicyForCreateScalingPolicyInput) SetEvaluationCount(v int32) *AlarmPolicyForCreateScalingPolicyInput {
	s.EvaluationCount = &v
	return s
}

// SetRuleType sets the RuleType field's value.
func (s *AlarmPolicyForCreateScalingPolicyInput) SetRuleType(v string) *AlarmPolicyForCreateScalingPolicyInput {
	s.RuleType = &v
	return s
}

type CreateScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	AdjustmentType *string `type:"string"`

	AdjustmentValue *int32 `type:"int32"`

	AlarmPolicy *AlarmPolicyForCreateScalingPolicyInput `type:"structure"`

	Cooldown *int32 `type:"int32"`

	ScalingGroupId *string `type:"string"`

	ScalingPolicyName *string `type:"string"`

	ScalingPolicyType *string `type:"string"`

	ScheduledPolicy *ScheduledPolicyForCreateScalingPolicyInput `type:"structure"`
}

// String returns the string representation
func (s CreateScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateScalingPolicyInput) GoString() string {
	return s.String()
}

// SetAdjustmentType sets the AdjustmentType field's value.
func (s *CreateScalingPolicyInput) SetAdjustmentType(v string) *CreateScalingPolicyInput {
	s.AdjustmentType = &v
	return s
}

// SetAdjustmentValue sets the AdjustmentValue field's value.
func (s *CreateScalingPolicyInput) SetAdjustmentValue(v int32) *CreateScalingPolicyInput {
	s.AdjustmentValue = &v
	return s
}

// SetAlarmPolicy sets the AlarmPolicy field's value.
func (s *CreateScalingPolicyInput) SetAlarmPolicy(v *AlarmPolicyForCreateScalingPolicyInput) *CreateScalingPolicyInput {
	s.AlarmPolicy = v
	return s
}

// SetCooldown sets the Cooldown field's value.
func (s *CreateScalingPolicyInput) SetCooldown(v int32) *CreateScalingPolicyInput {
	s.Cooldown = &v
	return s
}

// SetScalingGroupId sets the ScalingGroupId field's value.
func (s *CreateScalingPolicyInput) SetScalingGroupId(v string) *CreateScalingPolicyInput {
	s.ScalingGroupId = &v
	return s
}

// SetScalingPolicyName sets the ScalingPolicyName field's value.
func (s *CreateScalingPolicyInput) SetScalingPolicyName(v string) *CreateScalingPolicyInput {
	s.ScalingPolicyName = &v
	return s
}

// SetScalingPolicyType sets the ScalingPolicyType field's value.
func (s *CreateScalingPolicyInput) SetScalingPolicyType(v string) *CreateScalingPolicyInput {
	s.ScalingPolicyType = &v
	return s
}

// SetScheduledPolicy sets the ScheduledPolicy field's value.
func (s *CreateScalingPolicyInput) SetScheduledPolicy(v *ScheduledPolicyForCreateScalingPolicyInput) *CreateScalingPolicyInput {
	s.ScheduledPolicy = v
	return s
}

type CreateScalingPolicyOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	ScalingPolicyId *string `type:"string"`
}

// String returns the string representation
func (s CreateScalingPolicyOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s CreateScalingPolicyOutput) GoString() string {
	return s.String()
}

// SetScalingPolicyId sets the ScalingPolicyId field's value.
func (s *CreateScalingPolicyOutput) SetScalingPolicyId(v string) *CreateScalingPolicyOutput {
	s.ScalingPolicyId = &v
	return s
}

type ScheduledPolicyForCreateScalingPolicyInput struct {
	_ struct{} `type:"structure"`

	LaunchTime *string `type:"string"`

	RecurrenceEndTime *string `type:"string"`

	RecurrenceType *string `type:"string"`

	RecurrenceValue *string `type:"string"`
}

// String returns the string representation
func (s ScheduledPolicyForCreateScalingPolicyInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s ScheduledPolicyForCreateScalingPolicyInput) GoString() string {
	return s.String()
}

// SetLaunchTime sets the LaunchTime field's value.
func (s *ScheduledPolicyForCreateScalingPolicyInput) SetLaunchTime(v string) *ScheduledPolicyForCreateScalingPolicyInput {
	s.LaunchTime = &v
	return s
}

// SetRecurrenceEndTime sets the RecurrenceEndTime field's value.
func (s *ScheduledPolicyForCreateScalingPolicyInput) SetRecurrenceEndTime(v string) *ScheduledPolicyForCreateScalingPolicyInput {
	s.RecurrenceEndTime = &v
	return s
}

// SetRecurrenceType sets the RecurrenceType field's value.
func (s *ScheduledPolicyForCreateScalingPolicyInput) SetRecurrenceType(v string) *ScheduledPolicyForCreateScalingPolicyInput {
	s.RecurrenceType = &v
	return s
}

// SetRecurrenceValue sets the RecurrenceValue field's value.
func (s *ScheduledPolicyForCreateScalingPolicyInput) SetRecurrenceValue(v string) *ScheduledPolicyForCreateScalingPolicyInput {
	s.RecurrenceValue = &v
	return s
}
