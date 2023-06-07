// Code generated by volcengine with private/model/cli/gen-api/main.go. DO NOT EDIT.

package autoscaling

import (
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/request"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/response"
	"k8s.io/autoscaler/cluster-autoscaler/cloudprovider/volcengine/volcengine-go-sdk/volcengine/volcengineutil"
)

const opDeleteLifecycleHookCommon = "DeleteLifecycleHook"

// DeleteLifecycleHookCommonRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteLifecycleHookCommon operation. The "output" return
// value will be populated with the DeleteLifecycleHookCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteLifecycleHookCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteLifecycleHookCommon Send returns without error.
//
// See DeleteLifecycleHookCommon for more information on using the DeleteLifecycleHookCommon
// API call, and error handling.
//
//	// Example sending a request using the DeleteLifecycleHookCommonRequest method.
//	req, resp := client.DeleteLifecycleHookCommonRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DeleteLifecycleHookCommonRequest(input *map[string]interface{}) (req *request.Request, output *map[string]interface{}) {
	op := &request.Operation{
		Name:       opDeleteLifecycleHookCommon,
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

// DeleteLifecycleHookCommon API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteLifecycleHookCommon for usage and error information.
func (c *AUTOSCALING) DeleteLifecycleHookCommon(input *map[string]interface{}) (*map[string]interface{}, error) {
	req, out := c.DeleteLifecycleHookCommonRequest(input)
	return out, req.Send()
}

// DeleteLifecycleHookCommonWithContext is the same as DeleteLifecycleHookCommon with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteLifecycleHookCommon for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. If the context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteLifecycleHookCommonWithContext(ctx volcengine.Context, input *map[string]interface{}, opts ...request.Option) (*map[string]interface{}, error) {
	req, out := c.DeleteLifecycleHookCommonRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

const opDeleteLifecycleHook = "DeleteLifecycleHook"

// DeleteLifecycleHookRequest generates a "volcengine/request.Request" representing the
// client's request for the DeleteLifecycleHook operation. The "output" return
// value will be populated with the DeleteLifecycleHookCommon request's response once the request completes
// successfully.
//
// Use "Send" method on the returned DeleteLifecycleHookCommon Request to send the API call to the service.
// the "output" return value is not valid until after DeleteLifecycleHookCommon Send returns without error.
//
// See DeleteLifecycleHook for more information on using the DeleteLifecycleHook
// API call, and error handling.
//
//	// Example sending a request using the DeleteLifecycleHookRequest method.
//	req, resp := client.DeleteLifecycleHookRequest(params)
//
//	err := req.Send()
//	if err == nil { // resp is now filled
//	    fmt.Println(resp)
//	}
func (c *AUTOSCALING) DeleteLifecycleHookRequest(input *DeleteLifecycleHookInput) (req *request.Request, output *DeleteLifecycleHookOutput) {
	op := &request.Operation{
		Name:       opDeleteLifecycleHook,
		HTTPMethod: "GET",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &DeleteLifecycleHookInput{}
	}

	output = &DeleteLifecycleHookOutput{}
	req = c.newRequest(op, input, output)

	return
}

// DeleteLifecycleHook API operation for AUTO_SCALING.
//
// Returns volcengineerr.Error for service API and SDK errors. Use runtime type assertions
// with volcengineerr.Error's Code and Message methods to get detailed information about
// the error.
//
// See the VOLCENGINE API reference guide for AUTO_SCALING's
// API operation DeleteLifecycleHook for usage and error information.
func (c *AUTOSCALING) DeleteLifecycleHook(input *DeleteLifecycleHookInput) (*DeleteLifecycleHookOutput, error) {
	req, out := c.DeleteLifecycleHookRequest(input)
	return out, req.Send()
}

// DeleteLifecycleHookWithContext is the same as DeleteLifecycleHook with the addition of
// the ability to pass a context and additional request options.
//
// See DeleteLifecycleHook for details on how to use this API operation.
//
// The context must be non-nil and will be used for request cancellation. Ifthe context is nil a panic will occur.
// In the future the SDK may create sub-contexts for http.Requests. See https://golang.org/pkg/context/
// for more information on using Contexts.
func (c *AUTOSCALING) DeleteLifecycleHookWithContext(ctx volcengine.Context, input *DeleteLifecycleHookInput, opts ...request.Option) (*DeleteLifecycleHookOutput, error) {
	req, out := c.DeleteLifecycleHookRequest(input)
	req.SetContext(ctx)
	req.ApplyOptions(opts...)
	return out, req.Send()
}

type DeleteLifecycleHookInput struct {
	_ struct{} `type:"structure"`

	LifecycleHookId *string `type:"string"`
}

// String returns the string representation
func (s DeleteLifecycleHookInput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLifecycleHookInput) GoString() string {
	return s.String()
}

// SetLifecycleHookId sets the LifecycleHookId field's value.
func (s *DeleteLifecycleHookInput) SetLifecycleHookId(v string) *DeleteLifecycleHookInput {
	s.LifecycleHookId = &v
	return s
}

type DeleteLifecycleHookOutput struct {
	_ struct{} `type:"structure"`

	Metadata *response.ResponseMetadata

	LifecycleHookId *string `type:"string"`
}

// String returns the string representation
func (s DeleteLifecycleHookOutput) String() string {
	return volcengineutil.Prettify(s)
}

// GoString returns the string representation
func (s DeleteLifecycleHookOutput) GoString() string {
	return s.String()
}

// SetLifecycleHookId sets the LifecycleHookId field's value.
func (s *DeleteLifecycleHookOutput) SetLifecycleHookId(v string) *DeleteLifecycleHookOutput {
	s.LifecycleHookId = &v
	return s
}
