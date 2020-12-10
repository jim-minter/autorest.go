// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type bigDataPoolsClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *bigDataPoolsClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Get - Get Big Data Pool
func (client *bigDataPoolsClient) Get(ctx context.Context, bigDataPoolName string, options *BigDataPoolsGetOptions) (BigDataPoolResourceInfoResponse, error) {
	req, err := client.getCreateRequest(ctx, bigDataPoolName, options)
	if err != nil {
		return BigDataPoolResourceInfoResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BigDataPoolResourceInfoResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BigDataPoolResourceInfoResponse{}, client.getHandleError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *bigDataPoolsClient) getCreateRequest(ctx context.Context, bigDataPoolName string, options *BigDataPoolsGetOptions) (*azcore.Request, error) {
	urlPath := "/bigDataPools/{bigDataPoolName}"
	urlPath = strings.ReplaceAll(urlPath, "{bigDataPoolName}", url.PathEscape(bigDataPoolName))
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *bigDataPoolsClient) getHandleResponse(resp *azcore.Response) (BigDataPoolResourceInfoResponse, error) {
	var val *BigDataPoolResourceInfo
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BigDataPoolResourceInfoResponse{}, err
	}
	return BigDataPoolResourceInfoResponse{RawResponse: resp.Response, BigDataPoolResourceInfo: val}, nil
}

// getHandleError handles the Get error response.
func (client *bigDataPoolsClient) getHandleError(resp *azcore.Response) error {
	var err ErrorContract
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// List - List Big Data Pools
func (client *bigDataPoolsClient) List(ctx context.Context, options *BigDataPoolsListOptions) (BigDataPoolResourceInfoListResultResponse, error) {
	req, err := client.listCreateRequest(ctx, options)
	if err != nil {
		return BigDataPoolResourceInfoListResultResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BigDataPoolResourceInfoListResultResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BigDataPoolResourceInfoListResultResponse{}, client.listHandleError(resp)
	}
	return client.listHandleResponse(resp)
}

// listCreateRequest creates the List request.
func (client *bigDataPoolsClient) listCreateRequest(ctx context.Context, options *BigDataPoolsListOptions) (*azcore.Request, error) {
	urlPath := "/bigDataPools"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// listHandleResponse handles the List response.
func (client *bigDataPoolsClient) listHandleResponse(resp *azcore.Response) (BigDataPoolResourceInfoListResultResponse, error) {
	var val *BigDataPoolResourceInfoListResult
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BigDataPoolResourceInfoListResultResponse{}, err
	}
	return BigDataPoolResourceInfoListResultResponse{RawResponse: resp.Response, BigDataPoolResourceInfoListResult: val}, nil
}

// listHandleError handles the List error response.
func (client *bigDataPoolsClient) listHandleError(resp *azcore.Response) error {
	var err ErrorContract
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
