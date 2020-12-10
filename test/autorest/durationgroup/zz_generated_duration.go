// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package durationgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// DurationClient contains the methods for the Duration group.
// Don't use this type directly, use NewDurationClient() instead.
type DurationClient struct {
	con *Connection
}

// NewDurationClient creates a new instance of DurationClient with the specified values.
func NewDurationClient(con *Connection) *DurationClient {
	return &DurationClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *DurationClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetInvalid - Get an invalid duration value
func (client *DurationClient) GetInvalid(ctx context.Context, options *DurationGetInvalidOptions) (StringResponse, error) {
	req, err := client.getInvalidCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getInvalidHandleError(resp)
	}
	return client.getInvalidHandleResponse(resp)
}

// getInvalidCreateRequest creates the GetInvalid request.
func (client *DurationClient) getInvalidCreateRequest(ctx context.Context, options *DurationGetInvalidOptions) (*azcore.Request, error) {
	urlPath := "/duration/invalid"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getInvalidHandleResponse handles the GetInvalid response.
func (client *DurationClient) getInvalidHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getInvalidHandleError handles the GetInvalid error response.
func (client *DurationClient) getInvalidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetNull - Get null duration value
func (client *DurationClient) GetNull(ctx context.Context, options *DurationGetNullOptions) (StringResponse, error) {
	req, err := client.getNullCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getNullHandleError(resp)
	}
	return client.getNullHandleResponse(resp)
}

// getNullCreateRequest creates the GetNull request.
func (client *DurationClient) getNullCreateRequest(ctx context.Context, options *DurationGetNullOptions) (*azcore.Request, error) {
	urlPath := "/duration/null"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getNullHandleResponse handles the GetNull response.
func (client *DurationClient) getNullHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getNullHandleError handles the GetNull error response.
func (client *DurationClient) getNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetPositiveDuration - Get a positive duration value
func (client *DurationClient) GetPositiveDuration(ctx context.Context, options *DurationGetPositiveDurationOptions) (StringResponse, error) {
	req, err := client.getPositiveDurationCreateRequest(ctx, options)
	if err != nil {
		return StringResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringResponse{}, client.getPositiveDurationHandleError(resp)
	}
	return client.getPositiveDurationHandleResponse(resp)
}

// getPositiveDurationCreateRequest creates the GetPositiveDuration request.
func (client *DurationClient) getPositiveDurationCreateRequest(ctx context.Context, options *DurationGetPositiveDurationOptions) (*azcore.Request, error) {
	urlPath := "/duration/positiveduration"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getPositiveDurationHandleResponse handles the GetPositiveDuration response.
func (client *DurationClient) getPositiveDurationHandleResponse(resp *azcore.Response) (StringResponse, error) {
	var val *string
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringResponse{}, err
	}
	return StringResponse{RawResponse: resp.Response, Value: val}, nil
}

// getPositiveDurationHandleError handles the GetPositiveDuration error response.
func (client *DurationClient) getPositiveDurationHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutPositiveDuration - Put a positive duration value
func (client *DurationClient) PutPositiveDuration(ctx context.Context, durationBody string, options *DurationPutPositiveDurationOptions) (*http.Response, error) {
	req, err := client.putPositiveDurationCreateRequest(ctx, durationBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putPositiveDurationHandleError(resp)
	}
	return resp.Response, nil
}

// putPositiveDurationCreateRequest creates the PutPositiveDuration request.
func (client *DurationClient) putPositiveDurationCreateRequest(ctx context.Context, durationBody string, options *DurationPutPositiveDurationOptions) (*azcore.Request, error) {
	urlPath := "/duration/positiveduration"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(durationBody)
}

// putPositiveDurationHandleError handles the PutPositiveDuration error response.
func (client *DurationClient) putPositiveDurationHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
