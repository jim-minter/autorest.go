//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package nonstringenumgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// FloatClient contains the methods for the Float group.
// Don't use this type directly, use a constructor function instead.
type FloatClient struct {
	internal *azcore.Client
}

// Get - Get a float enum
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - FloatClientGetOptions contains the optional parameters for the FloatClient.Get method.
func (client *FloatClient) Get(ctx context.Context, options *FloatClientGetOptions) (FloatClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return FloatClientGetResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FloatClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FloatClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *FloatClient) getCreateRequest(ctx context.Context, options *FloatClientGetOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/float/get"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *FloatClient) getHandleResponse(resp *http.Response) (FloatClientGetResponse, error) {
	result := FloatClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return FloatClientGetResponse{}, err
	}
	return result, nil
}

// Put - Put a float enum
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - input - Input float enum.
//   - options - FloatClientPutOptions contains the optional parameters for the FloatClient.Put method.
func (client *FloatClient) Put(ctx context.Context, input FloatEnum, options *FloatClientPutOptions) (FloatClientPutResponse, error) {
	req, err := client.putCreateRequest(ctx, input, options)
	if err != nil {
		return FloatClientPutResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return FloatClientPutResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return FloatClientPutResponse{}, runtime.NewResponseError(resp)
	}
	return client.putHandleResponse(resp)
}

// putCreateRequest creates the Put request.
func (client *FloatClient) putCreateRequest(ctx context.Context, input FloatEnum, options *FloatClientPutOptions) (*policy.Request, error) {
	urlPath := "/nonStringEnums/float/put"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, input); err != nil {
		return nil, err
	}
	return req, nil
}

// putHandleResponse handles the Put response.
func (client *FloatClient) putHandleResponse(resp *http.Response) (FloatClientPutResponse, error) {
	result := FloatClientPutResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return FloatClientPutResponse{}, err
	}
	return result, nil
}
