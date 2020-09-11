// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armnetwork

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

// VpnSiteLinkConnectionsOperations contains the methods for the VpnSiteLinkConnections group.
type VpnSiteLinkConnectionsOperations interface {
	// Get - Retrieves the details of a vpn site link connection.
	Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (*VpnSiteLinkConnectionResponse, error)
}

// VpnSiteLinkConnectionsClient implements the VpnSiteLinkConnectionsOperations interface.
// Don't use this type directly, use NewVpnSiteLinkConnectionsClient() instead.
type VpnSiteLinkConnectionsClient struct {
	*Client
	subscriptionID string
}

// NewVpnSiteLinkConnectionsClient creates a new instance of VpnSiteLinkConnectionsClient with the specified values.
func NewVpnSiteLinkConnectionsClient(c *Client, subscriptionID string) VpnSiteLinkConnectionsOperations {
	return &VpnSiteLinkConnectionsClient{Client: c, subscriptionID: subscriptionID}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *VpnSiteLinkConnectionsClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Get - Retrieves the details of a vpn site link connection.
func (client *VpnSiteLinkConnectionsClient) Get(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (*VpnSiteLinkConnectionResponse, error) {
	req, err := client.GetCreateRequest(ctx, resourceGroupName, gatewayName, connectionName, linkConnectionName)
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
func (client *VpnSiteLinkConnectionsClient) GetCreateRequest(ctx context.Context, resourceGroupName string, gatewayName string, connectionName string, linkConnectionName string) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Network/vpnGateways/{gatewayName}/vpnConnections/{connectionName}/vpnLinkConnections/{linkConnectionName}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	urlPath = strings.ReplaceAll(urlPath, "{resourceGroupName}", url.PathEscape(resourceGroupName))
	urlPath = strings.ReplaceAll(urlPath, "{gatewayName}", url.PathEscape(gatewayName))
	urlPath = strings.ReplaceAll(urlPath, "{connectionName}", url.PathEscape(connectionName))
	urlPath = strings.ReplaceAll(urlPath, "{linkConnectionName}", url.PathEscape(linkConnectionName))
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
func (client *VpnSiteLinkConnectionsClient) GetHandleResponse(resp *azcore.Response) (*VpnSiteLinkConnectionResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetHandleError(resp)
	}
	result := VpnSiteLinkConnectionResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.VpnSiteLinkConnection)
}

// GetHandleError handles the Get error response.
func (client *VpnSiteLinkConnectionsClient) GetHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
