// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// ArrayClient contains the methods for the Array group.
// Don't use this type directly, use NewArrayClient() instead.
type ArrayClient struct {
	con *Connection
}

// NewArrayClient creates a new instance of ArrayClient with the specified values.
func NewArrayClient(con *Connection) *ArrayClient {
	return &ArrayClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *ArrayClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetEmpty - Get complex types with array property which is empty
func (client *ArrayClient) GetEmpty(ctx context.Context, options *ArrayGetEmptyOptions) (ArrayWrapperResponse, error) {
	req, err := client.getEmptyCreateRequest(ctx, options)
	if err != nil {
		return ArrayWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ArrayWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ArrayWrapperResponse{}, client.getEmptyHandleError(resp)
	}
	return client.getEmptyHandleResponse(resp)
}

// getEmptyCreateRequest creates the GetEmpty request.
func (client *ArrayClient) getEmptyCreateRequest(ctx context.Context, options *ArrayGetEmptyOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getEmptyHandleResponse handles the GetEmpty response.
func (client *ArrayClient) getEmptyHandleResponse(resp *azcore.Response) (ArrayWrapperResponse, error) {
	var val *ArrayWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ArrayWrapperResponse{}, err
	}
	return ArrayWrapperResponse{RawResponse: resp.Response, ArrayWrapper: val}, nil
}

// getEmptyHandleError handles the GetEmpty error response.
func (client *ArrayClient) getEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNotProvided - Get complex types with array property while server doesn't provide a response payload
func (client *ArrayClient) GetNotProvided(ctx context.Context, options *ArrayGetNotProvidedOptions) (ArrayWrapperResponse, error) {
	req, err := client.getNotProvidedCreateRequest(ctx, options)
	if err != nil {
		return ArrayWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ArrayWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ArrayWrapperResponse{}, client.getNotProvidedHandleError(resp)
	}
	return client.getNotProvidedHandleResponse(resp)
}

// getNotProvidedCreateRequest creates the GetNotProvided request.
func (client *ArrayClient) getNotProvidedCreateRequest(ctx context.Context, options *ArrayGetNotProvidedOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/notprovided"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNotProvidedHandleResponse handles the GetNotProvided response.
func (client *ArrayClient) getNotProvidedHandleResponse(resp *azcore.Response) (ArrayWrapperResponse, error) {
	var val *ArrayWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ArrayWrapperResponse{}, err
	}
	return ArrayWrapperResponse{RawResponse: resp.Response, ArrayWrapper: val}, nil
}

// getNotProvidedHandleError handles the GetNotProvided error response.
func (client *ArrayClient) getNotProvidedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetValid - Get complex types with array property
func (client *ArrayClient) GetValid(ctx context.Context, options *ArrayGetValidOptions) (ArrayWrapperResponse, error) {
	req, err := client.getValidCreateRequest(ctx, options)
	if err != nil {
		return ArrayWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ArrayWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ArrayWrapperResponse{}, client.getValidHandleError(resp)
	}
	return client.getValidHandleResponse(resp)
}

// getValidCreateRequest creates the GetValid request.
func (client *ArrayClient) getValidCreateRequest(ctx context.Context, options *ArrayGetValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getValidHandleResponse handles the GetValid response.
func (client *ArrayClient) getValidHandleResponse(resp *azcore.Response) (ArrayWrapperResponse, error) {
	var val *ArrayWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ArrayWrapperResponse{}, err
	}
	return ArrayWrapperResponse{RawResponse: resp.Response, ArrayWrapper: val}, nil
}

// getValidHandleError handles the GetValid error response.
func (client *ArrayClient) getValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutEmpty - Put complex types with array property which is empty
func (client *ArrayClient) PutEmpty(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutEmptyOptions) (*http.Response, error) {
	req, err := client.putEmptyCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putEmptyHandleError(resp)
	}
	return resp.Response, nil
}

// putEmptyCreateRequest creates the PutEmpty request.
func (client *ArrayClient) putEmptyCreateRequest(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutEmptyOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/empty"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putEmptyHandleError handles the PutEmpty error response.
func (client *ArrayClient) putEmptyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutValid - Put complex types with array property
func (client *ArrayClient) PutValid(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutValidOptions) (*http.Response, error) {
	req, err := client.putValidCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putValidHandleError(resp)
	}
	return resp.Response, nil
}

// putValidCreateRequest creates the PutValid request.
func (client *ArrayClient) putValidCreateRequest(ctx context.Context, complexBody ArrayWrapper, options *ArrayPutValidOptions) (*azcore.Request, error) {
	urlPath := "/complex/array/valid"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putValidHandleError handles the PutValid error response.
func (client *ArrayClient) putValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
