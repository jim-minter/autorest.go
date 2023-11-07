//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armdataboxedge

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

// MonitoringConfigClient contains the methods for the MonitoringConfig group.
// Don't use this type directly, use NewMonitoringConfigClient() instead.
type MonitoringConfigClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewMonitoringConfigClient creates a new instance of MonitoringConfigClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewMonitoringConfigClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*MonitoringConfigClient, error) {
	cl, err := arm.NewClient(moduleName, moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &MonitoringConfigClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a new metric configuration or updates an existing one for a role.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - resourceGroupName - The resource group name.
//   - monitoringMetricConfiguration - The metric configuration.
//   - options - MonitoringConfigClientBeginCreateOrUpdateOptions contains the optional parameters for the MonitoringConfigClient.BeginCreateOrUpdate
//     method.
func (client *MonitoringConfigClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*runtime.Poller[MonitoringConfigClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, roleName, resourceGroupName, monitoringMetricConfiguration, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[MonitoringConfigClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[MonitoringConfigClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates a new metric configuration or updates an existing one for a role.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *MonitoringConfigClient) createOrUpdate(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MonitoringConfigClient.BeginCreateOrUpdate")
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, roleName, resourceGroupName, monitoringMetricConfiguration, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *MonitoringConfigClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, monitoringMetricConfiguration MonitoringMetricConfiguration, options *MonitoringConfigClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, monitoringMetricConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - deletes a new metric configuration for a role.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - resourceGroupName - The resource group name.
//   - options - MonitoringConfigClientBeginDeleteOptions contains the optional parameters for the MonitoringConfigClient.BeginDelete
//     method.
func (client *MonitoringConfigClient) BeginDelete(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*runtime.Poller[MonitoringConfigClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, roleName, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller[MonitoringConfigClientDeleteResponse](resp, client.internal.Pipeline(), nil)
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken[MonitoringConfigClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - deletes a new metric configuration for a role.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *MonitoringConfigClient) deleteOperation(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MonitoringConfigClient.BeginDelete")
	req, err := client.deleteCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
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
func (client *MonitoringConfigClient) deleteCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a metric configuration of a role.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - resourceGroupName - The resource group name.
//   - options - MonitoringConfigClientGetOptions contains the optional parameters for the MonitoringConfigClient.Get method.
func (client *MonitoringConfigClient) Get(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientGetOptions) (MonitoringConfigClientGetResponse, error) {
	var err error
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MonitoringConfigClient.Get")
	req, err := client.getCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
	if err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return MonitoringConfigClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *MonitoringConfigClient) getCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig/default"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *MonitoringConfigClient) getHandleResponse(resp *http.Response) (MonitoringConfigClientGetResponse, error) {
	result := MonitoringConfigClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringMetricConfiguration); err != nil {
		return MonitoringConfigClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists metric configurations in a role.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - roleName - The role name.
//   - resourceGroupName - The resource group name.
//   - options - MonitoringConfigClientListOptions contains the optional parameters for the MonitoringConfigClient.NewListPager
//     method.
func (client *MonitoringConfigClient) NewListPager(deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientListOptions) *runtime.Pager[MonitoringConfigClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[MonitoringConfigClientListResponse]{
		More: func(page MonitoringConfigClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *MonitoringConfigClientListResponse) (MonitoringConfigClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "MonitoringConfigClient.NewListPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listCreateRequest(ctx, deviceName, roleName, resourceGroupName, options)
			}, nil)
			if err != nil {
				return MonitoringConfigClientListResponse{}, err
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *MonitoringConfigClient) listCreateRequest(ctx context.Context, deviceName string, roleName string, resourceGroupName string, options *MonitoringConfigClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/roles/{roleName}/monitoringConfig"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if roleName == "" {
		return nil, errors.New("parameter roleName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{roleName}", url.PathEscape(roleName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2021-02-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *MonitoringConfigClient) listHandleResponse(resp *http.Response) (MonitoringConfigClientListResponse, error) {
	result := MonitoringConfigClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.MonitoringMetricConfigurationList); err != nil {
		return MonitoringConfigClientListResponse{}, err
	}
	return result, nil
}
