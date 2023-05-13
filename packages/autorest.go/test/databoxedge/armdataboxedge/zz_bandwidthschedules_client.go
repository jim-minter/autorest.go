//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

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

// BandwidthSchedulesClient contains the methods for the BandwidthSchedules group.
// Don't use this type directly, use NewBandwidthSchedulesClient() instead.
type BandwidthSchedulesClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewBandwidthSchedulesClient creates a new instance of BandwidthSchedulesClient with the specified values.
//   - subscriptionID - The subscription ID.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewBandwidthSchedulesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*BandwidthSchedulesClient, error) {
	cl, err := arm.NewClient(moduleName+".BandwidthSchedulesClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &BandwidthSchedulesClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates or updates a bandwidth schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The bandwidth schedule name which needs to be added/updated.
//   - resourceGroupName - The resource group name.
//   - parameters - The bandwidth schedule to be added or updated.
//   - options - BandwidthSchedulesClientBeginCreateOrUpdateOptions contains the optional parameters for the BandwidthSchedulesClient.BeginCreateOrUpdate
//     method.
func (client *BandwidthSchedulesClient) BeginCreateOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, parameters BandwidthSchedule, options *BandwidthSchedulesClientBeginCreateOrUpdateOptions) (*runtime.Poller[BandwidthSchedulesClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, deviceName, name, resourceGroupName, parameters, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BandwidthSchedulesClientCreateOrUpdateResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[BandwidthSchedulesClientCreateOrUpdateResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// CreateOrUpdate - Creates or updates a bandwidth schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *BandwidthSchedulesClient) createOrUpdate(ctx context.Context, deviceName string, name string, resourceGroupName string, parameters BandwidthSchedule, options *BandwidthSchedulesClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	req, err := client.createOrUpdateCreateRequest(ctx, deviceName, name, resourceGroupName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createOrUpdateCreateRequest creates the CreateOrUpdate request.
func (client *BandwidthSchedulesClient) createOrUpdateCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, parameters BandwidthSchedule, options *BandwidthSchedulesClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
	if err := runtime.MarshalAsJSON(req, parameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes the specified bandwidth schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The bandwidth schedule name.
//   - resourceGroupName - The resource group name.
//   - options - BandwidthSchedulesClientBeginDeleteOptions contains the optional parameters for the BandwidthSchedulesClient.BeginDelete
//     method.
func (client *BandwidthSchedulesClient) BeginDelete(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesClientBeginDeleteOptions) (*runtime.Poller[BandwidthSchedulesClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, deviceName, name, resourceGroupName, options)
		if err != nil {
			return nil, err
		}
		return runtime.NewPoller[BandwidthSchedulesClientDeleteResponse](resp, client.internal.Pipeline(), nil)
	} else {
		return runtime.NewPollerFromResumeToken[BandwidthSchedulesClientDeleteResponse](options.ResumeToken, client.internal.Pipeline(), nil)
	}
}

// Delete - Deletes the specified bandwidth schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
func (client *BandwidthSchedulesClient) deleteOperation(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, deviceName, name, resourceGroupName, options)
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
func (client *BandwidthSchedulesClient) deleteCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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

// Get - Gets the properties of the specified bandwidth schedule.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - name - The bandwidth schedule name.
//   - resourceGroupName - The resource group name.
//   - options - BandwidthSchedulesClientGetOptions contains the optional parameters for the BandwidthSchedulesClient.Get method.
func (client *BandwidthSchedulesClient) Get(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesClientGetOptions) (BandwidthSchedulesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, deviceName, name, resourceGroupName, options)
	if err != nil {
		return BandwidthSchedulesClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return BandwidthSchedulesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return BandwidthSchedulesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *BandwidthSchedulesClient) getCreateRequest(ctx context.Context, deviceName string, name string, resourceGroupName string, options *BandwidthSchedulesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules/{name}"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
	if name == "" {
		return nil, errors.New("parameter name cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{name}", url.PathEscape(name))
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
func (client *BandwidthSchedulesClient) getHandleResponse(resp *http.Response) (BandwidthSchedulesClientGetResponse, error) {
	result := BandwidthSchedulesClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BandwidthSchedule); err != nil {
		return BandwidthSchedulesClientGetResponse{}, err
	}
	return result, nil
}

// NewListByDataBoxEdgeDevicePager - Gets all the bandwidth schedules for a Data Box Edge/Data Box Gateway device.
//
// Generated from API version 2021-02-01
//   - deviceName - The device name.
//   - resourceGroupName - The resource group name.
//   - options - BandwidthSchedulesClientListByDataBoxEdgeDeviceOptions contains the optional parameters for the BandwidthSchedulesClient.NewListByDataBoxEdgeDevicePager
//     method.
func (client *BandwidthSchedulesClient) NewListByDataBoxEdgeDevicePager(deviceName string, resourceGroupName string, options *BandwidthSchedulesClientListByDataBoxEdgeDeviceOptions) *runtime.Pager[BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse] {
	return runtime.NewPager(runtime.PagingHandler[BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse]{
		More: func(page BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse) (BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse, error) {
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listByDataBoxEdgeDeviceCreateRequest(ctx, deviceName, resourceGroupName, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse{}, runtime.NewResponseError(resp)
			}
			return client.listByDataBoxEdgeDeviceHandleResponse(resp)
		},
	})
}

// listByDataBoxEdgeDeviceCreateRequest creates the ListByDataBoxEdgeDevice request.
func (client *BandwidthSchedulesClient) listByDataBoxEdgeDeviceCreateRequest(ctx context.Context, deviceName string, resourceGroupName string, options *BandwidthSchedulesClientListByDataBoxEdgeDeviceOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/{deviceName}/bandwidthSchedules"
	urlPath = strings.ReplaceAll(urlPath, "{deviceName}", url.PathEscape(deviceName))
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

// listByDataBoxEdgeDeviceHandleResponse handles the ListByDataBoxEdgeDevice response.
func (client *BandwidthSchedulesClient) listByDataBoxEdgeDeviceHandleResponse(resp *http.Response) (BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse, error) {
	result := BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.BandwidthSchedulesList); err != nil {
		return BandwidthSchedulesClientListByDataBoxEdgeDeviceResponse{}, err
	}
	return result, nil
}
