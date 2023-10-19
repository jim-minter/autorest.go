//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// RouteFiltersClient contains the methods for the RouteFilters group.
// Don't use this type directly, use NewRouteFiltersClient() instead.
type RouteFiltersClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewRouteFiltersClient creates a new instance of RouteFiltersClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewRouteFiltersClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*RouteFiltersClient, error) {
	cl, err := arm.NewClient(moduleName+".RouteFiltersClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &RouteFiltersClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a route filter in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - routeFilterParameters - Parameters supplied to the create or update route filter operation.
//   - options - RouteFiltersClientBeginCreateOrUpdateOptions contains the optional parameters for the RouteFiltersClient.BeginCreateOrUpdate
//     method.
func (client *RouteFiltersClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersClientBeginCreateOrUpdateOptions) (*runtime.Poller[RouteFiltersClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RouteFiltersClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RouteFiltersClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates or updates a route filter in a specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *RouteFiltersClient) createOrUpdate(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "RouteFiltersClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, routeFilterName, routeFilterParameters, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *RouteFiltersClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, routeFilterParameters RouteFilter, options *RouteFiltersClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, routeFilterParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - options - RouteFiltersClientBeginDeleteOptions contains the optional parameters for the RouteFiltersClient.BeginDelete
//     method.
func (client *RouteFiltersClient) BeginDelete(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientBeginDeleteOptions) (*runtime.Poller[RouteFiltersClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, routeFilterName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[RouteFiltersClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[RouteFiltersClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *RouteFiltersClient) deleteOperation(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "RouteFiltersClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *RouteFiltersClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets the specified route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - options - RouteFiltersClientGetOptions contains the optional parameters for the RouteFiltersClient.Get method.
func (client *RouteFiltersClient) Get(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientGetOptions) (RouteFiltersClientGetResponse, error) {
	var err error
	const operationName = "RouteFiltersClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, routeFilterName, options)
	if err != nil {
		return RouteFiltersClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RouteFiltersClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RouteFiltersClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *RouteFiltersClient) getCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, options *RouteFiltersClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *RouteFiltersClient) getHandleResponse(resp *http.Response) (RouteFiltersClientGetResponse, error) {
	result := RouteFiltersClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilter); err != nil {
		return RouteFiltersClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Gets all route filters in a subscription.
//
// Generated from API version 2022-09-01
//   - options - RouteFiltersClientListOptions contains the optional parameters for the RouteFiltersClient.NewListPager method.
func (client *RouteFiltersClient) NewListPager(options *RouteFiltersClientListOptions) *runtime.Pager[RouteFiltersClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[RouteFiltersClientListResponse]{
		More: func(page RouteFiltersClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RouteFiltersClientListResponse) (RouteFiltersClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RouteFiltersClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, options)
			}, nil)
			if err != nil {
				return RouteFiltersClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *RouteFiltersClient) listCreateRequest(ctx context.Context, options *RouteFiltersClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Network/routeFilters"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *RouteFiltersClient) listHandleResponse(resp *http.Response) (RouteFiltersClientListResponse, error) {
	result := RouteFiltersClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilterListResult); err != nil {
		return RouteFiltersClientListResponse{}, err
	}
	return result, nil
}

// NewListByResourceGroupPager - Gets all route filters in a resource group.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - options - RouteFiltersClientListByResourceGroupOptions contains the optional parameters for the RouteFiltersClient.NewListByResourceGroupPager
//     method.
func (client *RouteFiltersClient) NewListByResourceGroupPager(resourceGroupName string, options *RouteFiltersClientListByResourceGroupOptions) *runtime.Pager[RouteFiltersClientListByResourceGroupResponse] {
	return runtime.NewPager(runtime.PagingHandler[RouteFiltersClientListByResourceGroupResponse]{
		More: func(page RouteFiltersClientListByResourceGroupResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *RouteFiltersClientListByResourceGroupResponse) (RouteFiltersClientListByResourceGroupResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "RouteFiltersClient.NewListByResourceGroupPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByResourceGroupCreateRequest(ctx, resourceGroupName, options)
			}, nil)
			if err != nil {
				return RouteFiltersClientListByResourceGroupResponse{}, err
			}
			return client.listByResourceGroupHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByResourceGroupCreateRequest creates the ListByResourceGroup request.
func (client *RouteFiltersClient) listByResourceGroupCreateRequest(ctx context.Context, resourceGroupName string, options *RouteFiltersClientListByResourceGroupOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listByResourceGroupHandleResponse handles the ListByResourceGroup response.
func (client *RouteFiltersClient) listByResourceGroupHandleResponse(resp *http.Response) (RouteFiltersClientListByResourceGroupResponse, error) {
	result := RouteFiltersClientListByResourceGroupResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilterListResult); err != nil {
		return RouteFiltersClientListByResourceGroupResponse{}, err
	}
	return result, nil
}

// UpdateTags - Updates tags of a route filter.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - routeFilterName - The name of the route filter.
//   - parameters - Parameters supplied to update route filter tags.
//   - options - RouteFiltersClientUpdateTagsOptions contains the optional parameters for the RouteFiltersClient.UpdateTags method.
func (client *RouteFiltersClient) UpdateTags(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersClientUpdateTagsOptions) (RouteFiltersClientUpdateTagsResponse, error) {
	var err error
	const operationName = "RouteFiltersClient.UpdateTags"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.updateTagsCreateRequest(ctx, resourceGroupName, routeFilterName, parameters, options)
	if err != nil {
		return RouteFiltersClientUpdateTagsResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return RouteFiltersClientUpdateTagsResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return RouteFiltersClientUpdateTagsResponse{}, err
	}
	resp, err := client.updateTagsHandleResponse(httpResp)
	return resp, err
}

// updateTagsCreateRequest creates the UpdateTags request.
func (client *RouteFiltersClient) updateTagsCreateRequest(ctx context.Context, resourceGroupName string, routeFilterName string, parameters TagsObject, options *RouteFiltersClientUpdateTagsOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/routeFilters/{routeFilterName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if routeFilterName == "" {
		return nil, errors.New("parameter routeFilterName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{routeFilterName}", url.PathEscape(routeFilterName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// updateTagsHandleResponse handles the UpdateTags response.
func (client *RouteFiltersClient) updateTagsHandleResponse(resp *http.Response) (RouteFiltersClientUpdateTagsResponse, error) {
	result := RouteFiltersClientUpdateTagsResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.RouteFilter); err != nil {
		return RouteFiltersClientUpdateTagsResponse{}, err
	}
	return result, nil
}
