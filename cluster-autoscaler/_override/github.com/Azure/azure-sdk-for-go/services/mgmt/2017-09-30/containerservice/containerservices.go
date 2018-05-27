package containerservice

// Copyright (c) Microsoft and contributors.  All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"context"
	"github.com/Azure/go-autorest/autorest"
	"github.com/Azure/go-autorest/autorest/azure"
	"github.com/Azure/go-autorest/autorest/validation"
	"net/http"
)

// ContainerServicesClient is the the Container Service Client.
type ContainerServicesClient struct {
	BaseClient
}

// NewContainerServicesClient creates an instance of the ContainerServicesClient client.
func NewContainerServicesClient(subscriptionID string) ContainerServicesClient {
	return NewContainerServicesClientWithBaseURI(DefaultBaseURI, subscriptionID)
}

// NewContainerServicesClientWithBaseURI creates an instance of the ContainerServicesClient client.
func NewContainerServicesClientWithBaseURI(baseURI string, subscriptionID string) ContainerServicesClient {
	return ContainerServicesClient{NewWithBaseURI(baseURI, subscriptionID)}
}

// CreateOrUpdate creates or updates a container service with the specified configuration of orchestrator, masters, and
// agents.
//
// resourceGroupName is the name of the resource group. containerServiceName is the name of the container service in
// the specified subscription and resource group. parameters is parameters supplied to the Create or Update a Container
// Service operation.
func (client ContainerServicesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService) (result ContainerServicesCreateOrUpdateFuture, err error) {
	if err := validation.Validate([]validation.Validation{
		{TargetValue: parameters,
			Constraints: []validation.Constraint{{Target: "parameters.Properties", Name: validation.Null, Rule: false,
				Chain: []validation.Constraint{{Target: "parameters.Properties.OrchestratorProfile", Name: validation.Null, Rule: true, Chain: nil},
					{Target: "parameters.Properties.CustomProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.CustomProfile.Orchestrator", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.Properties.ServicePrincipalProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.ServicePrincipalProfile.ClientID", Name: validation.Null, Rule: true, Chain: nil},
							{Target: "parameters.Properties.ServicePrincipalProfile.KeyVaultSecretRef", Name: validation.Null, Rule: false,
								Chain: []validation.Constraint{{Target: "parameters.Properties.ServicePrincipalProfile.KeyVaultSecretRef.VaultID", Name: validation.Null, Rule: true, Chain: nil},
									{Target: "parameters.Properties.ServicePrincipalProfile.KeyVaultSecretRef.SecretName", Name: validation.Null, Rule: true, Chain: nil},
								}},
						}},
					{Target: "parameters.Properties.MasterProfile", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.Properties.MasterProfile.DNSPrefix", Name: validation.Null, Rule: true, Chain: nil}}},
					{Target: "parameters.Properties.WindowsProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.WindowsProfile.AdminUsername", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.Properties.WindowsProfile.AdminUsername", Name: validation.Pattern, Rule: `^[a-zA-Z0-9]+([._]?[a-zA-Z0-9]+)*$`, Chain: nil}}},
							{Target: "parameters.Properties.WindowsProfile.AdminPassword", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.Properties.WindowsProfile.AdminPassword", Name: validation.Pattern, Rule: `^(?=.*[a-z])(?=.*[A-Z])(?=.*[!@#$%\^&\*\(\)])[a-zA-Z\d!@#$%\^&\*\(\)]{12,123}$`, Chain: nil}}},
						}},
					{Target: "parameters.Properties.LinuxProfile", Name: validation.Null, Rule: true,
						Chain: []validation.Constraint{{Target: "parameters.Properties.LinuxProfile.AdminUsername", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.Properties.LinuxProfile.AdminUsername", Name: validation.Pattern, Rule: `^[a-z][a-z0-9_-]*$`, Chain: nil}}},
							{Target: "parameters.Properties.LinuxProfile.SSH", Name: validation.Null, Rule: true,
								Chain: []validation.Constraint{{Target: "parameters.Properties.LinuxProfile.SSH.PublicKeys", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
					{Target: "parameters.Properties.DiagnosticsProfile", Name: validation.Null, Rule: false,
						Chain: []validation.Constraint{{Target: "parameters.Properties.DiagnosticsProfile.VMDiagnostics", Name: validation.Null, Rule: true,
							Chain: []validation.Constraint{{Target: "parameters.Properties.DiagnosticsProfile.VMDiagnostics.Enabled", Name: validation.Null, Rule: true, Chain: nil}}},
						}},
				}}}}}); err != nil {
		return result, validation.NewErrorWithValidationError(err, "containerservice.ContainerServicesClient", "CreateOrUpdate")
	}

	req, err := client.CreateOrUpdatePreparer(ctx, resourceGroupName, containerServiceName, parameters)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "CreateOrUpdate", nil, "Failure preparing request")
		return
	}

	result, err = client.CreateOrUpdateSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "CreateOrUpdate", result.Response(), "Failure sending request")
		return
	}

	return
}

// CreateOrUpdatePreparer prepares the CreateOrUpdate request.
func (client ContainerServicesClient) CreateOrUpdatePreparer(ctx context.Context, resourceGroupName string, containerServiceName string, parameters ContainerService) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerServiceName": autorest.Encode("path", containerServiceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsJSON(),
		autorest.AsPut(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", pathParameters),
		autorest.WithJSON(parameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// CreateOrUpdateSender sends the CreateOrUpdate request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) CreateOrUpdateSender(req *http.Request) (future ContainerServicesCreateOrUpdateFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted))
	return
}

// CreateOrUpdateResponder handles the response to the CreateOrUpdate request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) CreateOrUpdateResponder(resp *http.Response) (result ContainerService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusCreated, http.StatusAccepted),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// Delete deletes the specified container service in the specified subscription and resource group. The operation does
// not delete other resources created as part of creating a container service, including storage accounts, VMs, and
// availability sets. All the other resources created with the container service are part of the same resource group
// and can be deleted individually.
//
// resourceGroupName is the name of the resource group. containerServiceName is the name of the container service in
// the specified subscription and resource group.
func (client ContainerServicesClient) Delete(ctx context.Context, resourceGroupName string, containerServiceName string) (result ContainerServicesDeleteFuture, err error) {
	req, err := client.DeletePreparer(ctx, resourceGroupName, containerServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "Delete", nil, "Failure preparing request")
		return
	}

	result, err = client.DeleteSender(req)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "Delete", result.Response(), "Failure sending request")
		return
	}

	return
}

// DeletePreparer prepares the Delete request.
func (client ContainerServicesClient) DeletePreparer(ctx context.Context, resourceGroupName string, containerServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerServiceName": autorest.Encode("path", containerServiceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsDelete(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// DeleteSender sends the Delete request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) DeleteSender(req *http.Request) (future ContainerServicesDeleteFuture, err error) {
	sender := autorest.DecorateSender(client, azure.DoRetryWithRegistration(client.Client))
	future.Future = azure.NewFuture(req)
	future.req = req
	_, err = future.Done(sender)
	if err != nil {
		return
	}
	err = autorest.Respond(future.Response(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent))
	return
}

// DeleteResponder handles the response to the Delete request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) DeleteResponder(resp *http.Response) (result autorest.Response, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent),
		autorest.ByClosing())
	result.Response = resp
	return
}

// Get gets the properties of the specified container service in the specified subscription and resource group. The
// operation returns the properties including state, orchestrator, number of masters and agents, and FQDNs of masters
// and agents.
//
// resourceGroupName is the name of the resource group. containerServiceName is the name of the container service in
// the specified subscription and resource group.
func (client ContainerServicesClient) Get(ctx context.Context, resourceGroupName string, containerServiceName string) (result ContainerService, err error) {
	req, err := client.GetPreparer(ctx, resourceGroupName, containerServiceName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "Get", nil, "Failure preparing request")
		return
	}

	resp, err := client.GetSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "Get", resp, "Failure sending request")
		return
	}

	result, err = client.GetResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "Get", resp, "Failure responding to request")
	}

	return
}

// GetPreparer prepares the Get request.
func (client ContainerServicesClient) GetPreparer(ctx context.Context, resourceGroupName string, containerServiceName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"containerServiceName": autorest.Encode("path", containerServiceName),
		"resourceGroupName":    autorest.Encode("path", resourceGroupName),
		"subscriptionId":       autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices/{containerServiceName}", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// GetSender sends the Get request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) GetSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// GetResponder handles the response to the Get request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) GetResponder(resp *http.Response) (result ContainerService, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// List gets a list of container services in the specified subscription. The operation returns properties of each
// container service including state, orchestrator, number of masters and agents, and FQDNs of masters and agents.
func (client ContainerServicesClient) List(ctx context.Context) (result ListResultPage, err error) {
	result.fn = client.listNextResults
	req, err := client.ListPreparer(ctx)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "List", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "List", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "List", resp, "Failure responding to request")
	}

	return
}

// ListPreparer prepares the List request.
func (client ContainerServicesClient) ListPreparer(ctx context.Context) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/containerServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListSender sends the List request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) ListSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListResponder handles the response to the List request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) ListResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listNextResults retrieves the next set of results, if any.
func (client ContainerServicesClient) listNextResults(lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "listNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "listNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "listNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContainerServicesClient) ListComplete(ctx context.Context) (result ListResultIterator, err error) {
	result.page, err = client.List(ctx)
	return
}

// ListByResourceGroup gets a list of container services in the specified subscription and resource group. The
// operation returns properties of each container service including state, orchestrator, number of masters and agents,
// and FQDNs of masters and agents.
//
// resourceGroupName is the name of the resource group.
func (client ContainerServicesClient) ListByResourceGroup(ctx context.Context, resourceGroupName string) (result ListResultPage, err error) {
	result.fn = client.listByResourceGroupNextResults
	req, err := client.ListByResourceGroupPreparer(ctx, resourceGroupName)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "ListByResourceGroup", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.lr.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "ListByResourceGroup", resp, "Failure sending request")
		return
	}

	result.lr, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "ListByResourceGroup", resp, "Failure responding to request")
	}

	return
}

// ListByResourceGroupPreparer prepares the ListByResourceGroup request.
func (client ContainerServicesClient) ListByResourceGroupPreparer(ctx context.Context, resourceGroupName string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"resourceGroupName": autorest.Encode("path", resourceGroupName),
		"subscriptionId":    autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-07-01"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.ContainerService/containerServices", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListByResourceGroupSender sends the ListByResourceGroup request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) ListByResourceGroupSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListByResourceGroupResponder handles the response to the ListByResourceGroup request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) ListByResourceGroupResponder(resp *http.Response) (result ListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}

// listByResourceGroupNextResults retrieves the next set of results, if any.
func (client ContainerServicesClient) listByResourceGroupNextResults(lastResults ListResult) (result ListResult, err error) {
	req, err := lastResults.listResultPreparer()
	if err != nil {
		return result, autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "listByResourceGroupNextResults", nil, "Failure preparing next results request")
	}
	if req == nil {
		return
	}
	resp, err := client.ListByResourceGroupSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		return result, autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "listByResourceGroupNextResults", resp, "Failure sending next results request")
	}
	result, err = client.ListByResourceGroupResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "listByResourceGroupNextResults", resp, "Failure responding to next results request")
	}
	return
}

// ListByResourceGroupComplete enumerates all values, automatically crossing page boundaries as required.
func (client ContainerServicesClient) ListByResourceGroupComplete(ctx context.Context, resourceGroupName string) (result ListResultIterator, err error) {
	result.page, err = client.ListByResourceGroup(ctx, resourceGroupName)
	return
}

// ListOrchestrators gets a list of supported orchestrators in the specified subscription. The operation returns
// properties of each orchestrator including version and available upgrades.
//
// location is the name of a supported Azure region. resourceType is resource type for which the list of orchestrators
// needs to be returned
func (client ContainerServicesClient) ListOrchestrators(ctx context.Context, location string, resourceType string) (result OrchestratorVersionProfileListResult, err error) {
	req, err := client.ListOrchestratorsPreparer(ctx, location, resourceType)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "ListOrchestrators", nil, "Failure preparing request")
		return
	}

	resp, err := client.ListOrchestratorsSender(req)
	if err != nil {
		result.Response = autorest.Response{Response: resp}
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "ListOrchestrators", resp, "Failure sending request")
		return
	}

	result, err = client.ListOrchestratorsResponder(resp)
	if err != nil {
		err = autorest.NewErrorWithError(err, "containerservice.ContainerServicesClient", "ListOrchestrators", resp, "Failure responding to request")
	}

	return
}

// ListOrchestratorsPreparer prepares the ListOrchestrators request.
func (client ContainerServicesClient) ListOrchestratorsPreparer(ctx context.Context, location string, resourceType string) (*http.Request, error) {
	pathParameters := map[string]interface{}{
		"location":       autorest.Encode("path", location),
		"subscriptionId": autorest.Encode("path", client.SubscriptionID),
	}

	const APIVersion = "2017-09-30"
	queryParameters := map[string]interface{}{
		"api-version": APIVersion,
	}
	if len(resourceType) > 0 {
		queryParameters["resource-type"] = autorest.Encode("query", resourceType)
	}

	preparer := autorest.CreatePreparer(
		autorest.AsGet(),
		autorest.WithBaseURL(client.BaseURI),
		autorest.WithPathParameters("/subscriptions/{subscriptionId}/providers/Microsoft.ContainerService/locations/{location}/orchestrators", pathParameters),
		autorest.WithQueryParameters(queryParameters))
	return preparer.Prepare((&http.Request{}).WithContext(ctx))
}

// ListOrchestratorsSender sends the ListOrchestrators request. The method will close the
// http.Response Body if it receives an error.
func (client ContainerServicesClient) ListOrchestratorsSender(req *http.Request) (*http.Response, error) {
	return autorest.SendWithSender(client, req,
		azure.DoRetryWithRegistration(client.Client))
}

// ListOrchestratorsResponder handles the response to the ListOrchestrators request. The method always
// closes the http.Response Body.
func (client ContainerServicesClient) ListOrchestratorsResponder(resp *http.Response) (result OrchestratorVersionProfileListResult, err error) {
	err = autorest.Respond(
		resp,
		client.ByInspecting(),
		azure.WithErrorUnlessStatusCode(http.StatusOK),
		autorest.ByUnmarshallingJSON(&result),
		autorest.ByClosing())
	result.Response = autorest.Response{Response: resp}
	return
}
