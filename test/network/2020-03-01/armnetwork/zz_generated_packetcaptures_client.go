//go:build go1.16
// +build go1.16

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

// PacketCapturesClient contains the methods for the PacketCaptures group.
// Don't use this type directly, use NewPacketCapturesClient() instead.
type PacketCapturesClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPacketCapturesClient creates a new instance of PacketCapturesClient with the specified values.
// subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
// ID forms part of the URI for every service call.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPacketCapturesClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PacketCapturesClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PacketCapturesClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// BeginCreate - Create and start a packet capture on the specified VM.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// packetCaptureName - The name of the packet capture session.
// parameters - Parameters that define the create packet capture operation.
// options - PacketCapturesClientBeginCreateOptions contains the optional parameters for the PacketCapturesClient.BeginCreate
// method.
func (client *PacketCapturesClient) BeginCreate(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOptions) (PacketCapturesClientCreatePollerResponse, error) {
	resp, err := client.create(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
	if err != nil {
		return PacketCapturesClientCreatePollerResponse{}, err
	}
	result := PacketCapturesClientCreatePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.Create", "azure-async-operation", resp, client.pl)
	if err != nil {
		return PacketCapturesClientCreatePollerResponse{}, err
	}
	result.Poller = &PacketCapturesClientCreatePoller{
		pt: pt,
	}
	return result, nil
}

// Create - Create and start a packet capture on the specified VM.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PacketCapturesClient) create(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOptions) (*http.Response, error) {
	req, err := client.createCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, parameters, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusCreated) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// createCreateRequest creates the Create request.
func (client *PacketCapturesClient) createCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, parameters PacketCapture, options *PacketCapturesClientBeginCreateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
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

// BeginDelete - Deletes the specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// packetCaptureName - The name of the packet capture session.
// options - PacketCapturesClientBeginDeleteOptions contains the optional parameters for the PacketCapturesClient.BeginDelete
// method.
func (client *PacketCapturesClient) BeginDelete(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (PacketCapturesClientDeletePollerResponse, error) {
	resp, err := client.deleteOperation(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesClientDeletePollerResponse{}, err
	}
	result := PacketCapturesClientDeletePollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.Delete", "location", resp, client.pl)
	if err != nil {
		return PacketCapturesClientDeletePollerResponse{}, err
	}
	result.Poller = &PacketCapturesClientDeletePoller{
		pt: pt,
	}
	return result, nil
}

// Delete - Deletes the specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PacketCapturesClient) deleteOperation(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*http.Response, error) {
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted, http.StatusNoContent) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// deleteCreateRequest creates the Delete request.
func (client *PacketCapturesClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
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

// Get - Gets a packet capture session by name.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// packetCaptureName - The name of the packet capture session.
// options - PacketCapturesClientGetOptions contains the optional parameters for the PacketCapturesClient.Get method.
func (client *PacketCapturesClient) Get(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientGetOptions) (PacketCapturesClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PacketCapturesClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PacketCapturesClient) getCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
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

// getHandleResponse handles the Get response.
func (client *PacketCapturesClient) getHandleResponse(resp *http.Response) (PacketCapturesClientGetResponse, error) {
	result := PacketCapturesClientGetResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCaptureResult); err != nil {
		return PacketCapturesClientGetResponse{}, err
	}
	return result, nil
}

// BeginGetStatus - Query the status of a running packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the Network Watcher resource.
// packetCaptureName - The name given to the packet capture session.
// options - PacketCapturesClientBeginGetStatusOptions contains the optional parameters for the PacketCapturesClient.BeginGetStatus
// method.
func (client *PacketCapturesClient) BeginGetStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginGetStatusOptions) (PacketCapturesClientGetStatusPollerResponse, error) {
	resp, err := client.getStatus(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesClientGetStatusPollerResponse{}, err
	}
	result := PacketCapturesClientGetStatusPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.GetStatus", "location", resp, client.pl)
	if err != nil {
		return PacketCapturesClientGetStatusPollerResponse{}, err
	}
	result.Poller = &PacketCapturesClientGetStatusPoller{
		pt: pt,
	}
	return result, nil
}

// GetStatus - Query the status of a running packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PacketCapturesClient) getStatus(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginGetStatusOptions) (*http.Response, error) {
	req, err := client.getStatusCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// getStatusCreateRequest creates the GetStatus request.
func (client *PacketCapturesClient) getStatusCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginGetStatusOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/queryStatus"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// List - Lists all packet capture sessions within the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the Network Watcher resource.
// options - PacketCapturesClientListOptions contains the optional parameters for the PacketCapturesClient.List method.
func (client *PacketCapturesClient) List(resourceGroupName string, networkWatcherName string, options *PacketCapturesClientListOptions) *PacketCapturesClientListPager {
	return &PacketCapturesClientListPager{
		client: client,
		requester: func(ctx context.Context) (*policy.Request, error) {
			return client.listCreateRequest(ctx, resourceGroupName, networkWatcherName, options)
		},
	}
}

// listCreateRequest creates the List request.
func (client *PacketCapturesClient) listCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, options *PacketCapturesClientListOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
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
func (client *PacketCapturesClient) listHandleResponse(resp *http.Response) (PacketCapturesClientListResponse, error) {
	result := PacketCapturesClientListResponse{RawResponse: resp}
	if err := runtime.UnmarshalAsJSON(resp, &result.PacketCaptureListResult); err != nil {
		return PacketCapturesClientListResponse{}, err
	}
	return result, nil
}

// BeginStop - Stops a specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
// resourceGroupName - The name of the resource group.
// networkWatcherName - The name of the network watcher.
// packetCaptureName - The name of the packet capture session.
// options - PacketCapturesClientBeginStopOptions contains the optional parameters for the PacketCapturesClient.BeginStop
// method.
func (client *PacketCapturesClient) BeginStop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (PacketCapturesClientStopPollerResponse, error) {
	resp, err := client.stop(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return PacketCapturesClientStopPollerResponse{}, err
	}
	result := PacketCapturesClientStopPollerResponse{
		RawResponse: resp,
	}
	pt, err := armruntime.NewPoller("PacketCapturesClient.Stop", "location", resp, client.pl)
	if err != nil {
		return PacketCapturesClientStopPollerResponse{}, err
	}
	result.Poller = &PacketCapturesClientStopPoller{
		pt: pt,
	}
	return result, nil
}

// Stop - Stops a specified packet capture session.
// If the operation fails it returns an *azcore.ResponseError type.
func (client *PacketCapturesClient) stop(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*http.Response, error) {
	req, err := client.stopCreateRequest(ctx, resourceGroupName, networkWatcherName, packetCaptureName, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusAccepted) {
		return nil, runtime.NewResponseError(resp)
	}
	return resp, nil
}

// stopCreateRequest creates the Stop request.
func (client *PacketCapturesClient) stopCreateRequest(ctx context.Context, resourceGroupName string, networkWatcherName string, packetCaptureName string, options *PacketCapturesClientBeginStopOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/networkWatchers/{networkWatcherName}/packetCaptures/{packetCaptureName}/stop"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if networkWatcherName == "" {
		return nil, errors.New("parameter networkWatcherName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{networkWatcherName}", url.PathEscape(networkWatcherName))
	if packetCaptureName == "" {
		return nil, errors.New("parameter packetCaptureName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{packetCaptureName}", url.PathEscape(packetCaptureName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2020-03-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
