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

// VPNLinkConnectionsClient contains the methods for the VPNLinkConnections group.
// Don't use this type directly, use NewVPNLinkConnectionsClient() instead.
type VPNLinkConnectionsClient struct {
	internal       *arm.Client
	subscriptionID string
}

// NewVPNLinkConnectionsClient creates a new instance of VPNLinkConnectionsClient with the specified values.
//   - subscriptionID - The subscription credentials which uniquely identify the Microsoft Azure subscription. The subscription
//     ID forms part of the URI for every service call.
//   - credential - used to authorize requests. Usually a credential from azidentity.
//   - options - pass nil to accept the default values.
func NewVPNLinkConnectionsClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) (*VPNLinkConnectionsClient, error) {
	cl, err := arm.NewClient(moduleName+".VPNLinkConnectionsClient", moduleVersion, credential, options)
	if err != nil {
		return nil, err
	}
	client := &VPNLinkConnectionsClient{
		subscriptionID: subscriptionID,
		internal:       cl,
	}
	return client, nil
}

// BeginGetIkeSas - Lists IKE Security Associations for Vpn Site Link Connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the vpn connection.
//   - linkConnectionName - The name of the vpn link connection.
//   - options - VPNLinkConnectionsClientBeginGetIkeSasOptions contains the optional parameters for the VPNLinkConnectionsClient.BeginGetIkeSas
//     method.
func (client *VPNLinkConnectionsClient) BeginGetIkeSas(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNLinkConnectionsClientBeginGetIkeSasOptions) (*runtime.Poller[VPNLinkConnectionsClientGetIkeSasResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.getIkeSas(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNLinkConnectionsClientGetIkeSasResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNLinkConnectionsClientGetIkeSasResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// GetIkeSas - Lists IKE Security Associations for Vpn Site Link Connection in the specified resource group.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNLinkConnectionsClient) getIkeSas(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNLinkConnectionsClientBeginGetIkeSasOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNLinkConnectionsClient.BeginGetIkeSas"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getIkeSasCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName, options)
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

// getIkeSasCreateRequest creates the GetIkeSas request.
func (client *VPNLinkConnectionsClient) getIkeSasCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNLinkConnectionsClientBeginGetIkeSasOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}/getikesas"
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
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// NewListByVPNConnectionPager - Retrieves all vpn site link connections for a particular virtual wan vpn gateway vpn connection.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The resource group name of the vpn gateway.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the vpn connection.
//   - options - VPNLinkConnectionsClientListByVPNConnectionOptions contains the optional parameters for the VPNLinkConnectionsClient.NewListByVPNConnectionPager
//     method.
func (client *VPNLinkConnectionsClient) NewListByVPNConnectionPager(resourceGroupName string, gatewayName string, connectionName string, options *VPNLinkConnectionsClientListByVPNConnectionOptions) *runtime.Pager[VPNLinkConnectionsClientListByVPNConnectionResponse] {
	return runtime.NewPager(runtime.PagingHandler[VPNLinkConnectionsClientListByVPNConnectionResponse]{
		More: func(page VPNLinkConnectionsClientListByVPNConnectionResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *VPNLinkConnectionsClientListByVPNConnectionResponse) (VPNLinkConnectionsClientListByVPNConnectionResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "VPNLinkConnectionsClient.NewListByVPNConnectionPager")
			nextLink := ""
			if page != nil {
				nextLink = *page.NextLink
			}
			resp, err := runtime.FetcherForNextLink(ctx, client.internal.Pipeline(), nextLink, func(ctx context.Context) (*policy.Request, error) {
				return client.listByVPNConnectionCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, options)
			}, nil)
			if err != nil {
				return VPNLinkConnectionsClientListByVPNConnectionResponse{}, err
			}
			return client.listByVPNConnectionHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listByVPNConnectionCreateRequest creates the ListByVPNConnection request.
func (client *VPNLinkConnectionsClient) listByVPNConnectionCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, options *VPNLinkConnectionsClientListByVPNConnectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections"
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

// listByVPNConnectionHandleResponse handles the ListByVPNConnection response.
func (client *VPNLinkConnectionsClient) listByVPNConnectionHandleResponse(resp *http.Response) (VPNLinkConnectionsClientListByVPNConnectionResponse, error) {
	result := VPNLinkConnectionsClientListByVPNConnectionResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListVPNSiteLinkConnectionsResult); err != nil {
		return VPNLinkConnectionsClientListByVPNConnectionResponse{}, err
	}
	return result, nil
}

// BeginResetConnection - Resets the VpnLink connection specified.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
//   - resourceGroupName - The name of the resource group.
//   - gatewayName - The name of the gateway.
//   - connectionName - The name of the vpn connection.
//   - linkConnectionName - The name of the vpn link connection.
//   - options - VPNLinkConnectionsClientBeginResetConnectionOptions contains the optional parameters for the VPNLinkConnectionsClient.BeginResetConnection
//     method.
func (client *VPNLinkConnectionsClient) BeginResetConnection(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNLinkConnectionsClientBeginResetConnectionOptions) (*runtime.Poller[VPNLinkConnectionsClientResetConnectionResponse], error) {
	if options == nil || options.ResumeToken == "" {
		resp, err := client.resetConnection(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName, options)
		if err != nil {
			return nil, err
		}
		poller, err := runtime.NewPoller(resp, client.internal.Pipeline(), &runtime.NewPollerOptions[VPNLinkConnectionsClientResetConnectionResponse]{
			FinalStateVia: runtime.FinalStateViaLocation,
			Tracer:        client.internal.Tracer(),
		})
		return poller, err
	} else {
		return runtime.NewPollerFromResumeToken(options.ResumeToken, client.internal.Pipeline(), &runtime.NewPollerFromResumeTokenOptions[VPNLinkConnectionsClientResetConnectionResponse]{
			Tracer: client.internal.Tracer(),
		})
	}
}

// ResetConnection - Resets the VpnLink connection specified.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2022-09-01
func (client *VPNLinkConnectionsClient) resetConnection(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNLinkConnectionsClientBeginResetConnectionOptions) (*http.Response, error) {
	var err error
	const operationName = "VPNLinkConnectionsClient.BeginResetConnection"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.resetConnectionCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName, options)
	if err != nil {
		return nil, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusAccepted) {
		err = runtime.NewResponseError(httpResp)
		return nil, err
	}
	return httpResp, nil
}

// resetConnectionCreateRequest creates the ResetConnection request.
func (client *VPNLinkConnectionsClient) resetConnectionCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string, options *VPNLinkConnectionsClientBeginResetConnectionOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}/resetconnection"
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
	if linkConnectionName == "" {
		return nil, errors.New("parameter linkConnectionName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.internal.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2022-09-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}
