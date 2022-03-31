//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strings"
)

// TagsClient contains the methods for the Tags group.
// Don't use this type directly, use NewTagsClient() instead.
type TagsClient struct {
	host string
	pl   runtime.Pipeline
}

// NewTagsClient creates a new instance of TagsClient with the specified values.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewTagsClient(credential azcore.TokenCredential, options *arm.ClientOptions) *TagsClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &TagsClient{
		host: string(ep),
		pl:   armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Get all available tag keys for the defined scope
// If the operation fails it returns an *azcore.ResponseError type.
// scope - The scope associated with tags operations. This includes '/subscriptions/{subscriptionId}/' for subscription scope,
// '/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}' for
// resourceGroup scope, '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}' for Billing Account scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/departments/{departmentId}' for Department scope,
// '/providers/Microsoft.Billing/billingAccounts/{billingAccountId}/enrollmentAccounts/{enrollmentAccountId}' for EnrollmentAccount
// scope and
// '/providers/Microsoft.Management/managementGroups/{managementGroupId}' for Management Group scope..
// options - TagsClientGetOptions contains the optional parameters for the TagsClient.Get method.
func (client *TagsClient) Get(ctx context.Context, scope string, options *TagsClientGetOptions) (TagsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, scope, options)
	if err != nil {
		return TagsClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return TagsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusNoContent) {
		return TagsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *TagsClient) getCreateRequest(ctx context.Context, scope string, options *TagsClientGetOptions) (*policy.Request, error) {
	urlPath := "/{scope}/providers/Microsoft.Consumption/tags"
	urlPath = strings.ReplaceAll(urlPath, "{scope}", scope)
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *TagsClient) getHandleResponse(resp *http.Response) (TagsClientGetResponse, error) {
	result := TagsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.TagsResult); err != nil {
		return TagsClientGetResponse{}, err
	}
	return result, nil
}
