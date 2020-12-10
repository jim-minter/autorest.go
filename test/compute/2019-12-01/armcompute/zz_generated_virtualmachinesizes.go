// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armcompute

import (
	"context"
	"errors"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/armcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// VirtualMachineSizesClient contains the methods for the VirtualMachineSizes group.
// Don't use this type directly, use NewVirtualMachineSizesClient() instead.
type VirtualMachineSizesClient struct {
	con            *armcore.Connection
	subscriptionID string
}

// NewVirtualMachineSizesClient creates a new instance of VirtualMachineSizesClient with the specified values.
func NewVirtualMachineSizesClient(con *armcore.Connection, subscriptionID string) *VirtualMachineSizesClient {
	return &VirtualMachineSizesClient{con: con, subscriptionID: subscriptionID}
}

// Pipeline returns the pipeline associated with this client.
func (client *VirtualMachineSizesClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// List - This API is deprecated. Use Resources Skus [https://docs.microsoft.com/en-us/rest/api/compute/resourceskus/list]
func (client *VirtualMachineSizesClient) List(ctx context.Context, location string, options *VirtualMachineSizesListOptions) (VirtualMachineSizeListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, location, options)
	if err != nil {
		return VirtualMachineSizeListResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return VirtualMachineSizeListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return VirtualMachineSizeListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *VirtualMachineSizesClient) listCreateRequest(ctx context.Context, location string, options *VirtualMachineSizesListOptions) (*azcore.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Compute/locations/{location}/vmSizes"
	urlPath = strings.ReplaceAll(urlPath, "{location}", url.PathEscape(location))
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-12-01")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *VirtualMachineSizesClient) listHandleResponse(resp *azcore.Response) (VirtualMachineSizeListResultResponse, error) {
	var val *VirtualMachineSizeListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return VirtualMachineSizeListResultResponse{}, err
	}
	return VirtualMachineSizeListResultResponse{RawResponse: resp.Response, VirtualMachineSizeListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *VirtualMachineSizesClient) listHandleError(resp *azcore.Response) error {
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("%s; failed to read response body: %w", resp.Status, err)
	}
	if len(body) == 0 {
		return azcore.NewResponseError(errors.New(resp.Status), resp.Response)
	}
	return azcore.NewResponseError(errors.New(string(body)), resp.Response)
}
