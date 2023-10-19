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

// VPNConnectionsClient contains the methods for the VPNConnections group.
// Don't use this type directly, use NewVPNConnectionsClient() instead.
type VPNConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVPNConnectionsClient creates a new instance of VPNConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVPNConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".VPNConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VPNConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginCreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing
// connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the connection.
//   - vpnConnectionParameters - Parameters supplied to create or Update a VPN Connection.
//   - options - VPNConnectionsClientBeginCreateOrUpdateOptions contains the optional parameters for the VPNConnectionsClient.BeginCreateOrUpdate
//     method.
func (client *VPNConnectionsClient) BeginCreateOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*runtime.Poller[VPNConnectionsClientCreateOrUpdateResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.createOrUpdate(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientCreateOrUpdateResponse]{
			FinalStateVia: runtime.FinalStateViaAzureAsyncOp,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNConnectionsClientCreateOrUpdateResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// CreateOrUpdate - Creates a vpn connection to a scalable vpn gateway if it doesn't exist else updates the existing connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) createOrUpdate(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNConnectionsClient.BeginCreateOrUpdate"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createOrUpdateCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, vpnConnectionParameters, options)
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
func (client *VPNConnectionsClient) createOrUpdateCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, vpnConnectionParameters VPNConnection, options *VPNConnectionsClientBeginCreateOrUpdateOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, vpnConnectionParameters); err != nil {
		return nil, err
	}
	return req, nil
}

// BeginDelete - Deletes a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the connection.
//   - options - VPNConnectionsClientBeginDeleteOptions contains the optional parameters for the VPNConnectionsClient.BeginDelete
//     method.
func (client *VPNConnectionsClient) BeginDelete(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*runtime.Poller[VPNConnectionsClientDeleteResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.deleteOperation(ctx, resourceGroupName, gatewayName, connectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientDeleteResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNConnectionsClientDeleteResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// Delete - Deletes a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) deleteOperation(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNConnectionsClient.BeginDelete"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.deleteCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
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
func (client *VPNConnectionsClient) deleteCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientBeginDeleteOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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

// Get - Retrieves the details of a vpn connection.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the vpn connection.
//   - options - VPNConnectionsClientGetOptions contains the optional parameters for the VPNConnectionsClient.Get method.
func (client *VPNConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientGetOptions) (VPNConnectionsClientGetResponse, error) {
	var err error
	const operationName = "VPNConnectionsClient.Get"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
	if err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return VPNConnectionsClientGetResponse{}, err
	}
	resp, err := client.getHandleResponse(httpResp)
	return resp, err
}

// getCreateRequest creates the Get request.
func (client *VPNConnectionsClient) getCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNConnectionsClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if connectionName == "" {
		return nil, errors.New("parameter connectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
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
func (client *VPNConnectionsClient) getHandleResponse(resp *http.Response) (VPNConnectionsClientGetResponse, error) {
	result := VPNConnectionsClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.VPNConnection); err != nil {
		return VPNConnectionsClientGetResponse{}, err
	}
	return result, nil
}

// NewListByVPNGatewayPager - Retrieves all vpn connections for a particular virtual wan vpn gateway.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the VpnGateway.
//   - gatewayName - The name of the gateway.
//   - options - VPNConnectionsClientListByVPNGatewayOptions contains the optional parameters for the VPNConnectionsClient.NewListByVPNGatewayPager
//     method.
func (client *VPNConnectionsClient) NewListByVPNGatewayPager(resourceGroupName string, gatewayName string, options *VPNConnectionsClientListByVPNGatewayOptions) *runtime.Pager[VPNConnectionsClientListByVPNGatewayResponse] {
	return runtime.NewPager(runtime.PagingHandler[VPNConnectionsClientListByVPNGatewayResponse]{
		More: func(page VPNConnectionsClientListByVPNGatewayResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNConnectionsClientListByVPNGatewayResponse) (VPNConnectionsClientListByVPNGatewayResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VPNConnectionsClient.NewListByVPNGatewayPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByVPNGatewayCreateRequest(ctx, resourceGroupName, gatewayName, options)
			}, nil)
			if err != nil {
				return VPNConnectionsClientListByVPNGatewayResponse{}, err
			}
			return client.listByVPNGatewayHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVPNGatewayCreateRequest creates the ListByVPNGateway request.
func (client *VPNConnectionsClient) listByVPNGatewayCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, options *VPNConnectionsClientListByVPNGatewayOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
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

// listByVPNGatewayHandleResponse handles the ListByVPNGateway response.
func (client *VPNConnectionsClient) listByVPNGatewayHandleResponse(resp *http.Response) (VPNConnectionsClientListByVPNGatewayResponse, error) {
	result := VPNConnectionsClientListByVPNGatewayResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNConnectionsResult); err != nil {
		return VPNConnectionsClientListByVPNGatewayResponse{}, err
	}
	return result, nil
}

// BeginStartPacketCapture - Starts packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - gatewayName - The name of the gateway.
//   - vpnConnectionName - The name of the vpn connection.
//   - options - VPNConnectionsClientBeginStartPacketCaptureOptions contains the optional parameters for the VPNConnectionsClient.BeginStartPacketCapture
//     method.
func (client *VPNConnectionsClient) BeginStartPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStartPacketCaptureOptions) (*runtime.Poller[VPNConnectionsClientStartPacketCaptureResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.startPacketCapture(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientStartPacketCaptureResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNConnectionsClientStartPacketCaptureResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// StartPacketCapture - Starts packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) startPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStartPacketCaptureOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNConnectionsClient.BeginStartPacketCapture"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.startPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
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

// startPacketCaptureCreateRequest creates the StartPacketCapture request.
func (client *VPNConnectionsClient) startPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStartPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/startpacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if vpnConnectionName == "" {
		return nil, errors.New("parameter vpnConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnConnectionName}", url.PathEscape(vpnConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// BeginStopPacketCapture - Stops packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - gatewayName - The name of the gateway.
//   - vpnConnectionName - The name of the vpn connection.
//   - options - VPNConnectionsClientBeginStopPacketCaptureOptions contains the optional parameters for the VPNConnectionsClient.BeginStopPacketCapture
//     method.
func (client *VPNConnectionsClient) BeginStopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStopPacketCaptureOptions) (*runtime.Poller[VPNConnectionsClientStopPacketCaptureResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.stopPacketCapture(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNConnectionsClientStopPacketCaptureResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNConnectionsClientStopPacketCaptureResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// StopPacketCapture - Stops packet capture on Vpn connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNConnectionsClient) stopPacketCapture(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStopPacketCaptureOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNConnectionsClient.BeginStopPacketCapture"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.stopPacketCaptureCreateRequest(ctx, resourceGroupName, gatewayName, vpnConnectionName, options)
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

// stopPacketCaptureCreateRequest creates the StopPacketCapture request.
func (client *VPNConnectionsClient) stopPacketCaptureCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, vpnConnectionName string, options *VPNConnectionsClientBeginStopPacketCaptureOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{vpnConnectionName}/stoppacketcapture"
	if resourceGroupName == "" {
		return nil, errors.New("parameter resourceGroupName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	if gatewayName == "" {
		return nil, errors.New("parameter gatewayName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	if vpnConnectionName == "" {
		return nil, errors.New("parameter vpnConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{vpnConnectionName}", url.PathEscape(vpnConnectionName))
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Parameters != nil {
		if err := runtime.MarshalAsJSON(req, *options.Parameters); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}
