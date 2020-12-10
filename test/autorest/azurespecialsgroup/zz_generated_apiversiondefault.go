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

// APIVersionDefaultClient contains the methods for the APIVersionDefault group.
// Don't use this type directly, use NewAPIVersionDefaultClient() instead.
type APIVersionDefaultClient struct {
	con *Connection
}

// NewAPIVersionDefaultClient creates a new instance of APIVersionDefaultClient with the specified values.
func NewAPIVersionDefaultClient(con *Connection) *APIVersionDefaultClient {
	return &APIVersionDefaultClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *APIVersionDefaultClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetMethodGlobalNotProvidedValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetMethodGlobalNotProvidedValid(ctx context.Context, options *APIVersionDefaultGetMethodGlobalNotProvidedValidOptions) (*http.Response, error) {
	req, err := client.getMethodGlobalNotProvidedValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodGlobalNotProvidedValidHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodGlobalNotProvidedValidCreateRequest creates the GetMethodGlobalNotProvidedValid request.
func (client *APIVersionDefaultClient) getMethodGlobalNotProvidedValidCreateRequest(ctx context.Context, options *APIVersionDefaultGetMethodGlobalNotProvidedValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/globalNotProvided/2015-07-01-preview"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2015-07-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodGlobalNotProvidedValidHandleError handles the GetMethodGlobalNotProvidedValid error response.
func (client *APIVersionDefaultClient) getMethodGlobalNotProvidedValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetMethodGlobalValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetMethodGlobalValid(ctx context.Context, options *APIVersionDefaultGetMethodGlobalValidOptions) (*http.Response, error) {
	req, err := client.getMethodGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getMethodGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getMethodGlobalValidCreateRequest creates the GetMethodGlobalValid request.
func (client *APIVersionDefaultClient) getMethodGlobalValidCreateRequest(ctx context.Context, options *APIVersionDefaultGetMethodGlobalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/method/string/none/query/global/2015-07-01-preview"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2015-07-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getMethodGlobalValidHandleError handles the GetMethodGlobalValid error response.
func (client *APIVersionDefaultClient) getMethodGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPathGlobalValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetPathGlobalValid(ctx context.Context, options *APIVersionDefaultGetPathGlobalValidOptions) (*http.Response, error) {
	req, err := client.getPathGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getPathGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getPathGlobalValidCreateRequest creates the GetPathGlobalValid request.
func (client *APIVersionDefaultClient) getPathGlobalValidCreateRequest(ctx context.Context, options *APIVersionDefaultGetPathGlobalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/path/string/none/query/global/2015-07-01-preview"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2015-07-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPathGlobalValidHandleError handles the GetPathGlobalValid error response.
func (client *APIVersionDefaultClient) getPathGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetSwaggerGlobalValid - GET method with api-version modeled in global settings.
func (client *APIVersionDefaultClient) GetSwaggerGlobalValid(ctx context.Context, options *APIVersionDefaultGetSwaggerGlobalValidOptions) (*http.Response, error) {
	req, err := client.getSwaggerGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.getSwaggerGlobalValidHandleError(resp)
	}
	return resp.Response, nil
}

// getSwaggerGlobalValidCreateRequest creates the GetSwaggerGlobalValid request.
func (client *APIVersionDefaultClient) getSwaggerGlobalValidCreateRequest(ctx context.Context, options *APIVersionDefaultGetSwaggerGlobalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/apiVersion/swagger/string/none/query/global/2015-07-01-preview"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2015-07-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getSwaggerGlobalValidHandleError handles the GetSwaggerGlobalValid error response.
func (client *APIVersionDefaultClient) getSwaggerGlobalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
