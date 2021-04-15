// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// APIVersionLocalClient contains the methods for the APIVersionLocal group.
// Don't use this type directly, use NewAPIVersionLocalClient() instead.
type APIVersionLocalClient struct {
	con *Connection
}

// NewAPIVersionLocalClient creates a new instance of APIVersionLocalClient with the specified values.
func NewAPIVersionLocalClient(con *Connection) *APIVersionLocalClient {
	return &APIVersionLocalClient{con: con}
}

// GetMethodLocalNull - Get method with api-version modeled in the method. pass in api-version = null to succeed
func (client *APIVersionLocalClient) GetMethodLocalNull(ctx context.Context, options *APIVersionLocalGetMethodLocalNullOptions) (*http.Response, error) {
	req, err := client.getMethodLocalNullCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodLocalNullHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodLocalNullCreateRequest creates the GetMethodLocalNull request.
func (client *APIVersionLocalClient) getMethodLocalNullCreateRequest(ctx context.Context, options *APIVersionLocalGetMethodLocalNullOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	if options != nil && options.APIVersion != nil {
		reqQP.Set("api-version", *options.APIVersion)
	}
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodLocalNullHandleError handles the GetMethodLocalNull error response.
func (client *APIVersionLocalClient) getMethodLocalNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetMethodLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
func (client *APIVersionLocalClient) GetMethodLocalValid(ctx context.Context, options *APIVersionLocalGetMethodLocalValidOptions) (*http.Response, error) {
	req, err := client.getMethodLocalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodLocalValidCreateRequest creates the GetMethodLocalValid request.
func (client *APIVersionLocalClient) getMethodLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalGetMethodLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/local/2.0"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodLocalValidHandleError handles the GetMethodLocalValid error response.
func (client *APIVersionLocalClient) getMethodLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPathLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
func (client *APIVersionLocalClient) GetPathLocalValid(ctx context.Context, options *APIVersionLocalGetPathLocalValidOptions) (*http.Response, error) {
	req, err := client.getPathLocalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getPathLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getPathLocalValidCreateRequest creates the GetPathLocalValid request.
func (client *APIVersionLocalClient) getPathLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalGetPathLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/path/string/none/query/local/2.0"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPathLocalValidHandleError handles the GetPathLocalValid error response.
func (client *APIVersionLocalClient) getPathLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSwaggerLocalValid - Get method with api-version modeled in the method. pass in api-version = '2.0' to succeed
func (client *APIVersionLocalClient) GetSwaggerLocalValid(ctx context.Context, options *APIVersionLocalGetSwaggerLocalValidOptions) (*http.Response, error) {
	req, err := client.getSwaggerLocalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.con.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getSwaggerLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getSwaggerLocalValidCreateRequest creates the GetSwaggerLocalValid request.
func (client *APIVersionLocalClient) getSwaggerLocalValidCreateRequest(ctx context.Context, options *APIVersionLocalGetSwaggerLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/local/2.0"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	reqQP := req.URL.Query()
	reqQP.Set("api-version", "2.0")
	req.URL.RawQuery = reqQP.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSwaggerLocalValidHandleError handles the GetSwaggerLocalValid error response.
func (client *APIVersionLocalClient) getSwaggerLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}