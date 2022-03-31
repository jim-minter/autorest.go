//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// FirewallPoliciesClient contains the methods for the FirewallPolicies group.
// Don't use this type directly, use NewFirewallPoliciesClient() instead.
type FirewallPoliciesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewFirewallPoliciesClient creates a new instance of FirewallPoliciesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewFirewallPoliciesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *FirewallPoliciesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &FirewallPoliciesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreateOrUpdate - Creates or updates the specified Firewall Policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// parameters - Parameters supplied to the create or update Firewall Policy operation.
// options - FirewallPoliciesClientBeginCreateOrUpdateOptions contains the optional parameters for the FirewallPoliciesClient.BeginCreateOrUpdate
// method.
func (client *FirewallPoliciesClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy, options *FirewallPoliciesClientBeginCreateOrUpdateOptions) (*armruntime.Poller[FirewallPoliciesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, firewallPolicyName, parameters, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[FirewallPoliciesClientCreateOrUpdateResponse]("FirewallPoliciesClient.CreateOrUpdate", "azure-async-operation", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[FirewallPoliciesClientCreateOrUpdateResponse]("FirewallPoliciesClient.CreateOrUpdate", options.ResumeToken, client.pl, nil)
	}
}

// CreateOrUpdate - Creates or updates the specified Firewall Policy.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FirewallPoliciesClient) createOrUpdate(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy, options *FirewallPoliciesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, firewallPolicyName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *FirewallPoliciesClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, parameters FirewallPolicy, options *FirewallPoliciesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, runtime.MarshalAsJSON(req, parameters)
}

// BeginDelete - Deletes the specified Firewall Policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// options - FirewallPoliciesClientBeginDeleteOptions contains the optional parameters for the FirewallPoliciesClient.BeginDelete
// method.
func (client *FirewallPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesClientBeginDeleteOptions) (*armruntime.Poller[FirewallPoliciesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, firewallPolicyName, options)
		if err != nil {
			return nil, err
		}
		return armruntime.NewPoller[FirewallPoliciesClientDeleteResponse]("FirewallPoliciesClient.Delete", "location", resp, client.pl, nil)
	} else {
		return armruntime.NewPollerFromResumeToken[FirewallPoliciesClientDeleteResponse]("FirewallPoliciesClient.Delete", options.ResumeToken, client.pl, nil)
	}
}

// Delete - Deletes the specified Firewall Policy.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *FirewallPoliciesClient) deleteOperation(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *FirewallPoliciesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// Get - Gets the specified Firewall Policy.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// firewallPolicyName - The name of the Firewall Policy.
// options - FirewallPoliciesClientGetOptions contains the optional parameters for the FirewallPoliciesClient.Get method.
func (client *FirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesClientGetOptions) (FirewallPoliciesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, firewallPolicyName, options)
	if err != nil {
		return FirewallPoliciesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return FirewallPoliciesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FirewallPoliciesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FirewallPoliciesClient) getCreateRequest(ctx context.Context, resourceGroupName string, firewallPolicyName string, options *FirewallPoliciesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies/{firewallPolicyName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if firewallPolicyName == "" {
		return nil, errors.New("parameter firewallPolicyName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{firewallPolicyName}", url.PathEscape(firewallPolicyName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FirewallPoliciesClient) getHandleResponse(resp *http.Response) (FirewallPoliciesClientGetResponse, error) {
	result := FirewallPoliciesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicy); err != nil {
		return FirewallPoliciesClientGetResponse{}, err
	}
	return result, nil
}

// List - Lists all Firewall Policies in a resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// options - FirewallPoliciesClientListOptions contains the optional parameters for the FirewallPoliciesClient.List method.
func (client *FirewallPoliciesClient) List(resourceGroupName string, options *FirewallPoliciesClientListOptions) *runtime.Pager[FirewallPoliciesClientListResponse] {
	return runtime.NewPager(runtime.PageProcessor[FirewallPoliciesClientListResponse]{
		More: func(page FirewallPoliciesClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallPoliciesClientListResponse) (FirewallPoliciesClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FirewallPoliciesClientListResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FirewallPoliciesClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FirewallPoliciesClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *FirewallPoliciesClient) listCreateRequest(ctx context.Context, resourceGroupName string, options *FirewallPoliciesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/firewallPolicies"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *FirewallPoliciesClient) listHandleResponse(resp *http.Response) (FirewallPoliciesClientListResponse, error) {
	result := FirewallPoliciesClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyListResult); err != nil {
		return FirewallPoliciesClientListResponse{}, err
	}
	return result, nil
}

// ListAll - Gets all the Firewall Policies in a subscription.
// If the operation fails it returns an *azcore.ResponseError type.
// options - FirewallPoliciesClientListAllOptions contains the optional parameters for the FirewallPoliciesClient.ListAll
// method.
func (client *FirewallPoliciesClient) ListAll(options *FirewallPoliciesClientListAllOptions) *runtime.Pager[FirewallPoliciesClientListAllResponse] {
	return runtime.NewPager(runtime.PageProcessor[FirewallPoliciesClientListAllResponse]{
		More: func(page FirewallPoliciesClientListAllResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *FirewallPoliciesClientListAllResponse) (FirewallPoliciesClientListAllResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listAllCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return FirewallPoliciesClientListAllResponse{}, err
			}
			resp, err := client.pl.Do(req)
			if err != nil {
				return FirewallPoliciesClientListAllResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return FirewallPoliciesClientListAllResponse{}, runtime.NewResponseError(resp)
			}
			return client.listAllHandleResponse(resp)
		},
	})
}

// listAllCreateRequest creates the ListAll request.
func (client *FirewallPoliciesClient) listAllCreateRequest(ctx context.Context, options *FirewallPoliciesClientListAllOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/firewallPolicies"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// listAllHandleResponse handles the ListAll response.
func (client *FirewallPoliciesClient) listAllHandleResponse(resp *http.Response) (FirewallPoliciesClientListAllResponse, error) {
	result := FirewallPoliciesClientListAllResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.FirewallPolicyListResult); err != nil {
		return FirewallPoliciesClientListAllResponse{}, err
	}
	return result, nil
}
