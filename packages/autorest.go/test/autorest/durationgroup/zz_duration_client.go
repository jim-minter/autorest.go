//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package durationgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// DurationClient contains the methods for the Duration group.
// Don't use this type directly, use a constructor function instead.
type DurationClient struct {
	internal *azcore.Client
}

// GetInvalid - Get an invalid duration value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DurationClientGetInvalidOptions contains the optional parameters for the DurationClient.GetInvalid method.
func (client *DurationClient) GetInvalid(ctx context.Context, options *DurationClientGetInvalidOptions) (DurationClientGetInvalidResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return DurationClientGetInvalidResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationClientGetInvalidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DurationClientGetInvalidResponse{}, runtime.NewResponseError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *DurationClient) getInvalidCreateRequest(ctx context.Context, options *DurationClientGetInvalidOptions) (*policy.Request, error) {
	urlPath := "/duration/invalid"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *DurationClient) getInvalidHandleResponse(resp *http.Response) (DurationClientGetInvalidResponse, error) {
	result := DurationClientGetInvalidResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DurationClientGetInvalidResponse{}, err
	}
	return result, nil
}

// GetNull - Get null duration value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DurationClientGetNullOptions contains the optional parameters for the DurationClient.GetNull method.
func (client *DurationClient) GetNull(ctx context.Context, options *DurationClientGetNullOptions) (DurationClientGetNullResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return DurationClientGetNullResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationClientGetNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DurationClientGetNullResponse{}, runtime.NewResponseError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *DurationClient) getNullCreateRequest(ctx context.Context, options *DurationClientGetNullOptions) (*policy.Request, error) {
	urlPath := "/duration/null"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *DurationClient) getNullHandleResponse(resp *http.Response) (DurationClientGetNullResponse, error) {
	result := DurationClientGetNullResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DurationClientGetNullResponse{}, err
	}
	return result, nil
}

// GetPositiveDuration - Get a positive duration value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - DurationClientGetPositiveDurationOptions contains the optional parameters for the DurationClient.GetPositiveDuration
//     method.
func (client *DurationClient) GetPositiveDuration(ctx context.Context, options *DurationClientGetPositiveDurationOptions) (DurationClientGetPositiveDurationResponse, error) {
	req, err := client.getPositiveDurationCreateRequest(ctx, options)
	if err != nil {
		return DurationClientGetPositiveDurationResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationClientGetPositiveDurationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DurationClientGetPositiveDurationResponse{}, runtime.NewResponseError(resp)
	}
	return client.getPositiveDurationHandleResponse(resp)
}

// getPositiveDurationCreateRequest creates the GetPositiveDuration request.
func (client *DurationClient) getPositiveDurationCreateRequest(ctx context.Context, options *DurationClientGetPositiveDurationOptions) (*policy.Request, error) {
	urlPath := "/duration/positiveduration"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// getPositiveDurationHandleResponse handles the GetPositiveDuration response.
func (client *DurationClient) getPositiveDurationHandleResponse(resp *http.Response) (DurationClientGetPositiveDurationResponse, error) {
	result := DurationClientGetPositiveDurationResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return DurationClientGetPositiveDurationResponse{}, err
	}
	return result, nil
}

// PutPositiveDuration - Put a positive duration value
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - durationBody - duration body
//   - options - DurationClientPutPositiveDurationOptions contains the optional parameters for the DurationClient.PutPositiveDuration
//     method.
func (client *DurationClient) PutPositiveDuration(ctx context.Context, durationBody string, options *DurationClientPutPositiveDurationOptions) (DurationClientPutPositiveDurationResponse, error) {
	req, err := client.putPositiveDurationCreateRequest(ctx, durationBody, options)
	if err != nil {
		return DurationClientPutPositiveDurationResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return DurationClientPutPositiveDurationResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return DurationClientPutPositiveDurationResponse{}, runtime.NewResponseError(resp)
	}
	return DurationClientPutPositiveDurationResponse{}, nil
}

// putPositiveDurationCreateRequest creates the PutPositiveDuration request.
func (client *DurationClient) putPositiveDurationCreateRequest(ctx context.Context, durationBody string, options *DurationClientPutPositiveDurationOptions) (*policy.Request, error) {
	urlPath := "/duration/positiveduration"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, durationBody); err != nil {
		return nil, err
	}
	return req, nil
}
