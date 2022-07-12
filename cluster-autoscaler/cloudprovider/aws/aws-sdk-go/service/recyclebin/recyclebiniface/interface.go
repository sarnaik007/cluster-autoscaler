// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package recyclebiniface provides an interface to enable mocking the Amazon Recycle Bin service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package recyclebiniface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/recyclebin"
)

// RecycleBinAPI provides an interface to enable mocking the
// recyclebin.RecycleBin service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // Amazon Recycle Bin.
//    func myFunc(svc recyclebiniface.RecycleBinAPI) bool {
//        // Make svc.CreateRule request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := recyclebin.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockRecycleBinClient struct {
//        recyclebiniface.RecycleBinAPI
//    }
//    func (m *mockRecycleBinClient) CreateRule(input *recyclebin.CreateRuleInput) (*recyclebin.CreateRuleOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockRecycleBinClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type RecycleBinAPI interface {
	CreateRule(*recyclebin.CreateRuleInput) (*recyclebin.CreateRuleOutput, error)
	CreateRuleWithContext(aws.Context, *recyclebin.CreateRuleInput, ...request.Option) (*recyclebin.CreateRuleOutput, error)
	CreateRuleRequest(*recyclebin.CreateRuleInput) (*request.Request, *recyclebin.CreateRuleOutput)

	DeleteRule(*recyclebin.DeleteRuleInput) (*recyclebin.DeleteRuleOutput, error)
	DeleteRuleWithContext(aws.Context, *recyclebin.DeleteRuleInput, ...request.Option) (*recyclebin.DeleteRuleOutput, error)
	DeleteRuleRequest(*recyclebin.DeleteRuleInput) (*request.Request, *recyclebin.DeleteRuleOutput)

	GetRule(*recyclebin.GetRuleInput) (*recyclebin.GetRuleOutput, error)
	GetRuleWithContext(aws.Context, *recyclebin.GetRuleInput, ...request.Option) (*recyclebin.GetRuleOutput, error)
	GetRuleRequest(*recyclebin.GetRuleInput) (*request.Request, *recyclebin.GetRuleOutput)

	ListRules(*recyclebin.ListRulesInput) (*recyclebin.ListRulesOutput, error)
	ListRulesWithContext(aws.Context, *recyclebin.ListRulesInput, ...request.Option) (*recyclebin.ListRulesOutput, error)
	ListRulesRequest(*recyclebin.ListRulesInput) (*request.Request, *recyclebin.ListRulesOutput)

	ListRulesPages(*recyclebin.ListRulesInput, func(*recyclebin.ListRulesOutput, bool) bool) error
	ListRulesPagesWithContext(aws.Context, *recyclebin.ListRulesInput, func(*recyclebin.ListRulesOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*recyclebin.ListTagsForResourceInput) (*recyclebin.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *recyclebin.ListTagsForResourceInput, ...request.Option) (*recyclebin.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*recyclebin.ListTagsForResourceInput) (*request.Request, *recyclebin.ListTagsForResourceOutput)

	TagResource(*recyclebin.TagResourceInput) (*recyclebin.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *recyclebin.TagResourceInput, ...request.Option) (*recyclebin.TagResourceOutput, error)
	TagResourceRequest(*recyclebin.TagResourceInput) (*request.Request, *recyclebin.TagResourceOutput)

	UntagResource(*recyclebin.UntagResourceInput) (*recyclebin.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *recyclebin.UntagResourceInput, ...request.Option) (*recyclebin.UntagResourceOutput, error)
	UntagResourceRequest(*recyclebin.UntagResourceInput) (*request.Request, *recyclebin.UntagResourceOutput)

	UpdateRule(*recyclebin.UpdateRuleInput) (*recyclebin.UpdateRuleOutput, error)
	UpdateRuleWithContext(aws.Context, *recyclebin.UpdateRuleInput, ...request.Option) (*recyclebin.UpdateRuleOutput, error)
	UpdateRuleRequest(*recyclebin.UpdateRuleInput) (*request.Request, *recyclebin.UpdateRuleOutput)
}

var _ RecycleBinAPI = (*recyclebin.RecycleBin)(nil)
