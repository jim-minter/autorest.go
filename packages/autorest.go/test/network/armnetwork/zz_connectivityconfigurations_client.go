//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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
	"strconv"
	"strings"
)

// ConnectivityConfigurationsClient contains the methods for the ConnectivityConfigurations group.
// Don't use this type directly, use NewConnectivityConfigurationsClient() instead.
type ConnectivityConfigurationsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewConnectivityConfigurationsClient creates a new instance of ConnectivityConfigurationsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewConnectivityConfigurationsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*ConnectivityConfigurationsClient, error) {
	cl, err := arm.NewClient(moduleName+".ConnectivityConfigurationsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &ConnectivityConfigurationsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// CreateOrUpdate - Creates/Updates a new network manager connectivity configuration
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - configurationName - The name of the network manager connectivity configuration.
//   - connectivityConfiguration - Parameters supplied to create/update a network manager connectivity configuration
//   - options - ConnectivityConfigurationsClientCreateOrUpdateOptions contains the optional parameters for the ConnectivityConfigurationsClient.CreateOrUpdate
//     method.
func (client *ConnectivityConfigurationsClient) CreateOrUpdate(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, connectivityConfiguration ConnectivityConfiguration, options *ConnectivityConfigurationsClientCreateOrUpdateOptions) (ConnectivityConfigurationsClientCreateOrUpdateResponse, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, connectivityConfiguration, options)
	if err != nil {
		return ConnectivityConfigurationsClientCreateOrUpdateResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectivityConfigurationsClientCreateOrUpdateResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated) {
		return ConnectivityConfigurationsClientCreateOrUpdateResponse{}, runtime.NewResponseError(resp)
	}
	return client.createOrUpdateHandleResponse(resp)
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *ConnectivityConfigurationsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, connectivityConfiguration ConnectivityConfiguration, options *ConnectivityConfigurationsClientCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/connectivityConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, connectivityConfiguration); err != nil {
		return nil, err
	}
	return req, nil
}

// createOrUpdateHandleResponse handles the CreateOrUpdate response.
func (client *ConnectivityConfigurationsClient) createOrUpdateHandleResponse(resp *http.Response) (ConnectivityConfigurationsClientCreateOrUpdateResponse, error) {
	result := ConnectivityConfigurationsClientCreateOrUpdateResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectivityConfiguration); err != nil {
		return ConnectivityConfigurationsClientCreateOrUpdateResponse{}, err
	}
	return result, nil
}

// BeginDelete - Deletes a network manager connectivity configuration, specified by the resource group, network manager name,
// and connectivity configuration name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - configurationName - The name of the network manager connectivity configuration.
//   - options - ConnectivityConfigurationsClientBeginDeleteOptions contains the optional parameters for the ConnectivityConfigurationsClient.BeginDelete
//     method.
func (client *ConnectivityConfigurationsClient) BeginDelete(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *ConnectivityConfigurationsClientBeginDeleteOptions) (*runtime.Poller[ConnectivityConfigurationsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, networkManagerName, configurationName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[ConnectivityConfigurationsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
		})
	} else {
		return runtime.NewPollerFromResumeToken[ConnectivityConfigurationsClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes a network manager connectivity configuration, specified by the resource group, network manager name, and
// connectivity configuration name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *ConnectivityConfigurationsClient) deleteOperation(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *ConnectivityConfigurationsClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *ConnectivityConfigurationsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *ConnectivityConfigurationsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/connectivityConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	if options != nil && options.Force != nil {
		reqQP.Set("force", strconv.FormatBool(*options.Force))
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get - Gets a Network Connectivity Configuration, specified by the resource group, network manager name, and connectivity
// Configuration name
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - configurationName - The name of the network manager connectivity configuration.
//   - options - ConnectivityConfigurationsClientGetOptions contains the optional parameters for the ConnectivityConfigurationsClient.Get
//     method.
func (client *ConnectivityConfigurationsClient) Get(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *ConnectivityConfigurationsClientGetOptions) (ConnectivityConfigurationsClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkManagerName, configurationName, options)
	if err != nil {
		return ConnectivityConfigurationsClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ConnectivityConfigurationsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return ConnectivityConfigurationsClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *ConnectivityConfigurationsClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, configurationName string, options *ConnectivityConfigurationsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/connectivityConfigurations/{configurationName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	if configurationName == "" {
		return nil, errors.New("parameter configurationName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{configurationName}", url.PathEscape(configurationName))
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

// getHandleResponse handles the Get response.
func (client *ConnectivityConfigurationsClient) getHandleResponse(resp *http.Response) (ConnectivityConfigurationsClientGetResponse, error) {
	result := ConnectivityConfigurationsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectivityConfiguration); err != nil {
		return ConnectivityConfigurationsClientGetResponse{}, err
	}
	return result, nil
}

// NewListPager - Lists all the network manager connectivity configuration in a specified network manager.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - networkManagerName - The name of the network manager.
//   - options - ConnectivityConfigurationsClientListOptions contains the optional parameters for the ConnectivityConfigurationsClient.NewListPager
//     method.
func (client *ConnectivityConfigurationsClient) NewListPager(resourceGroupName string, networkManagerName string, options *ConnectivityConfigurationsClientListOptions) *runtime.Pager[ConnectivityConfigurationsClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ConnectivityConfigurationsClientListResponse]{
		More: func(page ConnectivityConfigurationsClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ConnectivityConfigurationsClientListResponse) (ConnectivityConfigurationsClientListResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, resourceGroupName, networkManagerName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ConnectivityConfigurationsClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ConnectivityConfigurationsClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ConnectivityConfigurationsClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
	})
}

// listCreateRequest creates the List request.
func (client *ConnectivityConfigurationsClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkManagerName string, options *ConnectivityConfigurationsClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkManagers/{networkManagerName}/connectivityConfigurations"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkManagerName == "" {
		return nil, errors.New("parameter networkManagerName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkManagerName}", url.PathEscape(networkManagerName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	if options != nil && options.SkipToken != nil {
		reqQP.Set("$skipToken", *options.SkipToken)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *ConnectivityConfigurationsClient) listHandleResponse(resp *http.Response) (ConnectivityConfigurationsClientListResponse, error) {
	result := ConnectivityConfigurationsClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ConnectivityConfigurationListResult); err != nil {
		return ConnectivityConfigurationsClientListResponse{}, err
	}
	return result, nil
}
