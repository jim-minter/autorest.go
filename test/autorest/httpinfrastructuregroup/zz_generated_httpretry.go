// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// HTTPRetryClient contains the methods for the HTTPRetry group.
// Don't use this type directly, use NewHTTPRetryClient() instead.
type HTTPRetryClient struct {
	con *Connection
}

// NewHTTPRetryClient creates a new instance of HTTPRetryClient with the specified values.
func NewHTTPRetryClient(con *Connection) *HTTPRetryClient {
	return &HTTPRetryClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *HTTPRetryClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Delete503 - Return 503 status code, then 200 after retry
func (client *HTTPRetryClient) Delete503(ctx context.Context, options *HTTPRetryDelete503Options) (*http.Response, error) {
	req, err := client.delete503CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.delete503HandleError(resp)
	}
	return resp.Response, nil
}

// delete503CreateRequest creates the Delete503 request.
func (client *HTTPRetryClient) delete503CreateRequest(ctx context.Context, options *HTTPRetryDelete503Options) (*azcore.Request, error) {
	urlPath := "/http/retry/503"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete503HandleError handles the Delete503 error response.
func (client *HTTPRetryClient) delete503HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get502 - Return 502 status code, then 200 after retry
func (client *HTTPRetryClient) Get502(ctx context.Context, options *HTTPRetryGet502Options) (*http.Response, error) {
	req, err := client.get502CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.get502HandleError(resp)
	}
	return resp.Response, nil
}

// get502CreateRequest creates the Get502 request.
func (client *HTTPRetryClient) get502CreateRequest(ctx context.Context, options *HTTPRetryGet502Options) (*azcore.Request, error) {
	urlPath := "/http/retry/502"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get502HandleError handles the Get502 error response.
func (client *HTTPRetryClient) get502HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head408 - Return 408 status code, then 200 after retry
func (client *HTTPRetryClient) Head408(ctx context.Context, options *HTTPRetryHead408Options) (BooleanResponse, error) {
	req, err := client.head408CreateRequest(ctx, options)
	if err != nil {
		return BooleanResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BooleanResponse{}, err
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		return BooleanResponse{RawResponse: resp.Response, Success: true}, nil
	} else if resp.StatusCode >= 400 && resp.StatusCode < 500 {
		return BooleanResponse{RawResponse: resp.Response, Success: false}, nil
	} else {
		return BooleanResponse{}, client.head408HandleError(resp)
	}
}

// head408CreateRequest creates the Head408 request.
func (client *HTTPRetryClient) head408CreateRequest(ctx context.Context, options *HTTPRetryHead408Options) (*azcore.Request, error) {
	urlPath := "/http/retry/408"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head408HandleError handles the Head408 error response.
func (client *HTTPRetryClient) head408HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Options502 - Return 502 status code, then 200 after retry
func (client *HTTPRetryClient) Options502(ctx context.Context, options *HTTPRetryOptions502Options) (BoolResponse, error) {
	req, err := client.options502CreateRequest(ctx, options)
	if err != nil {
		return BoolResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolResponse{}, client.options502HandleError(resp)
	}
	return client.options502HandleResponse(resp)
}

// options502CreateRequest creates the Options502 request.
func (client *HTTPRetryClient) options502CreateRequest(ctx context.Context, options *HTTPRetryOptions502Options) (*azcore.Request, error) {
	urlPath := "/http/retry/502"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// options502HandleResponse handles the Options502 response.
func (client *HTTPRetryClient) options502HandleResponse(resp *azcore.Response) (BoolResponse, error) {
	var val *bool
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BoolResponse{}, err
	}
	return BoolResponse{RawResponse: resp.Response, Value: val}, nil
}

// options502HandleError handles the Options502 error response.
func (client *HTTPRetryClient) options502HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch500 - Return 500 status code, then 200 after retry
func (client *HTTPRetryClient) Patch500(ctx context.Context, options *HTTPRetryPatch500Options) (*http.Response, error) {
	req, err := client.patch500CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.patch500HandleError(resp)
	}
	return resp.Response, nil
}

// patch500CreateRequest creates the Patch500 request.
func (client *HTTPRetryClient) patch500CreateRequest(ctx context.Context, options *HTTPRetryPatch500Options) (*azcore.Request, error) {
	urlPath := "/http/retry/500"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch500HandleError handles the Patch500 error response.
func (client *HTTPRetryClient) patch500HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch504 - Return 504 status code, then 200 after retry
func (client *HTTPRetryClient) Patch504(ctx context.Context, options *HTTPRetryPatch504Options) (*http.Response, error) {
	req, err := client.patch504CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.patch504HandleError(resp)
	}
	return resp.Response, nil
}

// patch504CreateRequest creates the Patch504 request.
func (client *HTTPRetryClient) patch504CreateRequest(ctx context.Context, options *HTTPRetryPatch504Options) (*azcore.Request, error) {
	urlPath := "/http/retry/504"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch504HandleError handles the Patch504 error response.
func (client *HTTPRetryClient) patch504HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post503 - Return 503 status code, then 200 after retry
func (client *HTTPRetryClient) Post503(ctx context.Context, options *HTTPRetryPost503Options) (*http.Response, error) {
	req, err := client.post503CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.post503HandleError(resp)
	}
	return resp.Response, nil
}

// post503CreateRequest creates the Post503 request.
func (client *HTTPRetryClient) post503CreateRequest(ctx context.Context, options *HTTPRetryPost503Options) (*azcore.Request, error) {
	urlPath := "/http/retry/503"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post503HandleError handles the Post503 error response.
func (client *HTTPRetryClient) post503HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put500 - Return 500 status code, then 200 after retry
func (client *HTTPRetryClient) Put500(ctx context.Context, options *HTTPRetryPut500Options) (*http.Response, error) {
	req, err := client.put500CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.put500HandleError(resp)
	}
	return resp.Response, nil
}

// put500CreateRequest creates the Put500 request.
func (client *HTTPRetryClient) put500CreateRequest(ctx context.Context, options *HTTPRetryPut500Options) (*azcore.Request, error) {
	urlPath := "/http/retry/500"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put500HandleError handles the Put500 error response.
func (client *HTTPRetryClient) put500HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put504 - Return 504 status code, then 200 after retry
func (client *HTTPRetryClient) Put504(ctx context.Context, options *HTTPRetryPut504Options) (*http.Response, error) {
	req, err := client.put504CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.put504HandleError(resp)
	}
	return resp.Response, nil
}

// put504CreateRequest creates the Put504 request.
func (client *HTTPRetryClient) put504CreateRequest(ctx context.Context, options *HTTPRetryPut504Options) (*azcore.Request, error) {
	urlPath := "/http/retry/504"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put504HandleError handles the Put504 error response.
func (client *HTTPRetryClient) put504HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
