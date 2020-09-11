// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// WebApplicationFirewallPoliciesOperations contains the methods for the WebApplicationFirewallPolicies group.
type WebApplicationFirewallPoliciesOperations interface {
	// CreateOrUpdate - Creates or update policy with specified rule set name within a resource group.
	CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (*WebApplicationFirewallPolicyResponse, error)
	// BeginDelete - Deletes Policy.
	BeginDelete(ctx context.Context, resourceGroupName string, policyName string) (*HTTPPollerResponse, error)
	// ResumeDelete - Used to create a new instance of this poller from the resume token of a previous instance of this poller type.
	ResumeDelete(token string) (HTTPPoller, error)
	// Get - Retrieve protection policy with specified name within a resource group.
	Get(ctx context.Context, resourceGroupName string, policyName string) (*WebApplicationFirewallPolicyResponse, error)
	// List - Lists all of the protection policies within a resource group.
	List(resourceGroupName string) WebApplicationFirewallPolicyListResultPager
	// ListAll - Gets all the WAF policies in a subscription.
	ListAll() WebApplicationFirewallPolicyListResultPager
}

// WebApplicationFirewallPoliciesClient implements the WebApplicationFirewallPoliciesOperations interface.
// Don't use this type directly, use NewWebApplicationFirewallPoliciesClient() instead.
type WebApplicationFirewallPoliciesClient struct {
	*Client
	subscriptionID string
}

// NewWebApplicationFirewallPoliciesClient creates a new instance of WebApplicationFirewallPoliciesClient with the specified values.
func NewWebApplicationFirewallPoliciesClient(c *Client, subscriptionID string) WebApplicationFirewallPoliciesOperations {
	return &WebApplicationFirewallPoliciesClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *WebApplicationFirewallPoliciesClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// CreateOrUpdate - Creates or update policy with specified rule set name within a resource group.
func (client *WebApplicationFirewallPoliciesClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (*WebApplicationFirewallPolicyResponse, error) {
	req, err := client.CreateOrUpdateCreateRequest(ctx, resourceGroupName, policyName, parameters)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.CreateOrUpdateHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *WebApplicationFirewallPoliciesClient) CreateOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, policyName string, parameters WebApplicationFirewallPolicy) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(parameters)
}

// CreateOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *WebApplicationFirewallPoliciesClient) CreateOrUpdateHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusCreated) {
		return nil, client.CreateOrUpdateHandleError(resp)
	}
	result := WebApplicationFirewallPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicy)
}

// CreateOrUpdateHandleError handles the CreateOrUpdate error response.
func (client *WebApplicationFirewallPoliciesClient) CreateOrUpdateHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete - Deletes Policy.
func (client *WebApplicationFirewallPoliciesClient) BeginDelete(ctx context.Context, resourceGroupName string, policyName string) (*HTTPPollerResponse, error) {
	req, err := client.DeleteCreateRequest(ctx, resourceGroupName, policyName)
	if err != nil {
		return nil, err
	}
	// send the first request to initialize the poller
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.DeleteHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	pt, err := armcore.NewPoller("WebApplicationFirewallPoliciesClient.Delete", "location", resp, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	poller := &httpPoller{
		pt:       pt,
		pipeline: client.p,
	}
	result.Poller = poller
	result.PollUntilDone = func(ctx context.Context, frequency time.Duration) (*http.Response, error) {
		return poller.pollUntilDone(ctx, frequency)
	}
	return result, nil
}

func (client *WebApplicationFirewallPoliciesClient) ResumeDelete(token string) (HTTPPoller, error) {
	pt, err := armcore.NewPollerFromResumeToken("WebApplicationFirewallPoliciesClient.Delete", token, client.DeleteHandleError)
	if err != nil {
		return nil, err
	}
	return &httpPoller{
		pipeline: client.p,
		pt:       pt,
	}, nil
}

// DeleteCreateRequest creates the Delete request.
func (client *WebApplicationFirewallPoliciesClient) DeleteCreateRequest(ctx context.Context, resourceGroupName string, policyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// DeleteHandleResponse handles the Delete response.
func (client *WebApplicationFirewallPoliciesClient) DeleteHandleResponse(resp *azcore.Response) (*HTTPPollerResponse, error) {
	if !resp.HasStatusCode(http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, client.DeleteHandleError(resp)
	}
	return &HTTPPollerResponse{RawResponse: resp.Response}, nil
}

// DeleteHandleError handles the Delete error response.
func (client *WebApplicationFirewallPoliciesClient) DeleteHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get - Retrieve protection policy with specified name within a resource group.
func (client *WebApplicationFirewallPoliciesClient) Get(ctx context.Context, resourceGroupName string, policyName string) (*WebApplicationFirewallPolicyResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, policyName)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetCreateRequest creates the Get request.
func (client *WebApplicationFirewallPoliciesClient) GetCreateRequest(ctx context.Context, resourceGroupName string, policyName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies/{policyName}"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{policyName}", url.PathEscape(policyName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetHandleResponse handles the Get response.
func (client *WebApplicationFirewallPoliciesClient) GetHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := WebApplicationFirewallPolicyResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicy)
}

// GetHandleError handles the Get error response.
func (client *WebApplicationFirewallPoliciesClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// List - Lists all of the protection policies within a resource group.
func (client *WebApplicationFirewallPoliciesClient) List(resourceGroupName string) WebApplicationFirewallPolicyListResultPager {
	return &webApplicationFirewallPolicyListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListCreateRequest(ctx, resourceGroupName)
		},
		responder: client.ListHandleResponse,
		advancer: func(ctx context.Context, resp *WebApplicationFirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.WebApplicationFirewallPolicyListResult.NextLink)
		},
	}
}

// ListCreateRequest creates the List request.
func (client *WebApplicationFirewallPoliciesClient) ListCreateRequest(ctx context.Context, resourceGroupName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListHandleResponse handles the List response.
func (client *WebApplicationFirewallPoliciesClient) ListHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListHandleError(resp)
	}
	result := WebApplicationFirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicyListResult)
}

// ListHandleError handles the List error response.
func (client *WebApplicationFirewallPoliciesClient) ListHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// ListAll - Gets all the WAF policies in a subscription.
func (client *WebApplicationFirewallPoliciesClient) ListAll() WebApplicationFirewallPolicyListResultPager {
	return &webApplicationFirewallPolicyListResultPager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListAllCreateRequest(ctx)
		},
		responder: client.ListAllHandleResponse,
		advancer: func(ctx context.Context, resp *WebApplicationFirewallPolicyListResultResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.WebApplicationFirewallPolicyListResult.NextLink)
		},
	}
}

// ListAllCreateRequest creates the ListAll request.
func (client *WebApplicationFirewallPoliciesClient) ListAllCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/ApplicationGatewayWebApplicationFirewallPolicies"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("api-version", "2020-03-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// ListAllHandleResponse handles the ListAll response.
func (client *WebApplicationFirewallPoliciesClient) ListAllHandleResponse(resp *azcore.Response) (*WebApplicationFirewallPolicyListResultResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListAllHandleError(resp)
	}
	result := WebApplicationFirewallPolicyListResultResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.WebApplicationFirewallPolicyListResult)
}

// ListAllHandleError handles the ListAll error response.
func (client *WebApplicationFirewallPoliciesClient) ListAllHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
