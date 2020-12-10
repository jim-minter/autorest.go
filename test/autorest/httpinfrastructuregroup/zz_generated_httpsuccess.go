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

// HTTPSuccessClient contains the methods for the HTTPSuccess group.
// Don't use this type directly, use NewHTTPSuccessClient() instead.
type HTTPSuccessClient struct {
	con *Connection
}

// NewHTTPSuccessClient creates a new instance of HTTPSuccessClient with the specified values.
func NewHTTPSuccessClient(con *Connection) *HTTPSuccessClient {
	return &HTTPSuccessClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *HTTPSuccessClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// Delete200 - Delete simple boolean value true returns 200
func (client *HTTPSuccessClient) Delete200(ctx context.Context, options *HTTPSuccessDelete200Options) (*http.Response, error) {
	req, err := client.delete200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.delete200HandleError(resp)
	}
	return resp.Response, nil
}

// delete200CreateRequest creates the Delete200 request.
func (client *HTTPSuccessClient) delete200CreateRequest(ctx context.Context, options *HTTPSuccessDelete200Options) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete200HandleError handles the Delete200 error response.
func (client *HTTPSuccessClient) delete200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete202 - Delete true Boolean value in request returns 202 (accepted)
func (client *HTTPSuccessClient) Delete202(ctx context.Context, options *HTTPSuccessDelete202Options) (*http.Response, error) {
	req, err := client.delete202CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.delete202HandleError(resp)
	}
	return resp.Response, nil
}

// delete202CreateRequest creates the Delete202 request.
func (client *HTTPSuccessClient) delete202CreateRequest(ctx context.Context, options *HTTPSuccessDelete202Options) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete202HandleError handles the Delete202 error response.
func (client *HTTPSuccessClient) delete202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Delete204 - Delete true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Delete204(ctx context.Context, options *HTTPSuccessDelete204Options) (*http.Response, error) {
	req, err := client.delete204CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.delete204HandleError(resp)
	}
	return resp.Response, nil
}

// delete204CreateRequest creates the Delete204 request.
func (client *HTTPSuccessClient) delete204CreateRequest(ctx context.Context, options *HTTPSuccessDelete204Options) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// delete204HandleError handles the Delete204 error response.
func (client *HTTPSuccessClient) delete204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Get200 - Get 200 success
func (client *HTTPSuccessClient) Get200(ctx context.Context, options *HTTPSuccessGet200Options) (BoolResponse, error) {
	req, err := client.get200CreateRequest(ctx, options)
	if err != nil {
		return BoolResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolResponse{}, client.get200HandleError(resp)
	}
	return client.get200HandleResponse(resp)
}

// get200CreateRequest creates the Get200 request.
func (client *HTTPSuccessClient) get200CreateRequest(ctx context.Context, options *HTTPSuccessGet200Options) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// get200HandleResponse handles the Get200 response.
func (client *HTTPSuccessClient) get200HandleResponse(resp *azcore.Response) (BoolResponse, error) {
	var val *bool
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BoolResponse{}, err
	}
	return BoolResponse{RawResponse: resp.Response, Value: val}, nil
}

// get200HandleError handles the Get200 error response.
func (client *HTTPSuccessClient) get200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head200 - Return 200 status code if successful
func (client *HTTPSuccessClient) Head200(ctx context.Context, options *HTTPSuccessHead200Options) (BooleanResponse, error) {
	req, err := client.head200CreateRequest(ctx, options)
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
		return BooleanResponse{}, client.head200HandleError(resp)
	}
}

// head200CreateRequest creates the Head200 request.
func (client *HTTPSuccessClient) head200CreateRequest(ctx context.Context, options *HTTPSuccessHead200Options) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head200HandleError handles the Head200 error response.
func (client *HTTPSuccessClient) head200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head204 - Return 204 status code if successful
func (client *HTTPSuccessClient) Head204(ctx context.Context, options *HTTPSuccessHead204Options) (BooleanResponse, error) {
	req, err := client.head204CreateRequest(ctx, options)
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
		return BooleanResponse{}, client.head204HandleError(resp)
	}
}

// head204CreateRequest creates the Head204 request.
func (client *HTTPSuccessClient) head204CreateRequest(ctx context.Context, options *HTTPSuccessHead204Options) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head204HandleError handles the Head204 error response.
func (client *HTTPSuccessClient) head204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Head404 - Return 404 status code
func (client *HTTPSuccessClient) Head404(ctx context.Context, options *HTTPSuccessHead404Options) (BooleanResponse, error) {
	req, err := client.head404CreateRequest(ctx, options)
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
		return BooleanResponse{}, client.head404HandleError(resp)
	}
}

// head404CreateRequest creates the Head404 request.
func (client *HTTPSuccessClient) head404CreateRequest(ctx context.Context, options *HTTPSuccessHead404Options) (*azcore.Request, error) {
	urlPath := "/http/success/404"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// head404HandleError handles the Head404 error response.
func (client *HTTPSuccessClient) head404HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Options200 - Options 200 success
func (client *HTTPSuccessClient) Options200(ctx context.Context, options *HTTPSuccessOptions200Options) (BoolResponse, error) {
	req, err := client.options200CreateRequest(ctx, options)
	if err != nil {
		return BoolResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BoolResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BoolResponse{}, client.options200HandleError(resp)
	}
	return client.options200HandleResponse(resp)
}

// options200CreateRequest creates the Options200 request.
func (client *HTTPSuccessClient) options200CreateRequest(ctx context.Context, options *HTTPSuccessOptions200Options) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// options200HandleResponse handles the Options200 response.
func (client *HTTPSuccessClient) options200HandleResponse(resp *azcore.Response) (BoolResponse, error) {
	var val *bool
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BoolResponse{}, err
	}
	return BoolResponse{RawResponse: resp.Response, Value: val}, nil
}

// options200HandleError handles the Options200 error response.
func (client *HTTPSuccessClient) options200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch200 - Patch true Boolean value in request returning 200
func (client *HTTPSuccessClient) Patch200(ctx context.Context, options *HTTPSuccessPatch200Options) (*http.Response, error) {
	req, err := client.patch200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.patch200HandleError(resp)
	}
	return resp.Response, nil
}

// patch200CreateRequest creates the Patch200 request.
func (client *HTTPSuccessClient) patch200CreateRequest(ctx context.Context, options *HTTPSuccessPatch200Options) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch200HandleError handles the Patch200 error response.
func (client *HTTPSuccessClient) patch200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch202 - Patch true Boolean value in request returns 202
func (client *HTTPSuccessClient) Patch202(ctx context.Context, options *HTTPSuccessPatch202Options) (*http.Response, error) {
	req, err := client.patch202CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.patch202HandleError(resp)
	}
	return resp.Response, nil
}

// patch202CreateRequest creates the Patch202 request.
func (client *HTTPSuccessClient) patch202CreateRequest(ctx context.Context, options *HTTPSuccessPatch202Options) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch202HandleError handles the Patch202 error response.
func (client *HTTPSuccessClient) patch202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Patch204 - Patch true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Patch204(ctx context.Context, options *HTTPSuccessPatch204Options) (*http.Response, error) {
	req, err := client.patch204CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.patch204HandleError(resp)
	}
	return resp.Response, nil
}

// patch204CreateRequest creates the Patch204 request.
func (client *HTTPSuccessClient) patch204CreateRequest(ctx context.Context, options *HTTPSuccessPatch204Options) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// patch204HandleError handles the Patch204 error response.
func (client *HTTPSuccessClient) patch204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post200 - Post bollean value true in request that returns a 200
func (client *HTTPSuccessClient) Post200(ctx context.Context, options *HTTPSuccessPost200Options) (*http.Response, error) {
	req, err := client.post200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.post200HandleError(resp)
	}
	return resp.Response, nil
}

// post200CreateRequest creates the Post200 request.
func (client *HTTPSuccessClient) post200CreateRequest(ctx context.Context, options *HTTPSuccessPost200Options) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post200HandleError handles the Post200 error response.
func (client *HTTPSuccessClient) post200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post201 - Post true Boolean value in request returns 201 (Created)
func (client *HTTPSuccessClient) Post201(ctx context.Context, options *HTTPSuccessPost201Options) (*http.Response, error) {
	req, err := client.post201CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.post201HandleError(resp)
	}
	return resp.Response, nil
}

// post201CreateRequest creates the Post201 request.
func (client *HTTPSuccessClient) post201CreateRequest(ctx context.Context, options *HTTPSuccessPost201Options) (*azcore.Request, error) {
	urlPath := "/http/success/201"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post201HandleError handles the Post201 error response.
func (client *HTTPSuccessClient) post201HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post202 - Post true Boolean value in request returns 202 (Accepted)
func (client *HTTPSuccessClient) Post202(ctx context.Context, options *HTTPSuccessPost202Options) (*http.Response, error) {
	req, err := client.post202CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.post202HandleError(resp)
	}
	return resp.Response, nil
}

// post202CreateRequest creates the Post202 request.
func (client *HTTPSuccessClient) post202CreateRequest(ctx context.Context, options *HTTPSuccessPost202Options) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post202HandleError handles the Post202 error response.
func (client *HTTPSuccessClient) post202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Post204 - Post true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Post204(ctx context.Context, options *HTTPSuccessPost204Options) (*http.Response, error) {
	req, err := client.post204CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.post204HandleError(resp)
	}
	return resp.Response, nil
}

// post204CreateRequest creates the Post204 request.
func (client *HTTPSuccessClient) post204CreateRequest(ctx context.Context, options *HTTPSuccessPost204Options) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// post204HandleError handles the Post204 error response.
func (client *HTTPSuccessClient) post204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put200 - Put boolean value true returning 200 success
func (client *HTTPSuccessClient) Put200(ctx context.Context, options *HTTPSuccessPut200Options) (*http.Response, error) {
	req, err := client.put200CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.put200HandleError(resp)
	}
	return resp.Response, nil
}

// put200CreateRequest creates the Put200 request.
func (client *HTTPSuccessClient) put200CreateRequest(ctx context.Context, options *HTTPSuccessPut200Options) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put200HandleError handles the Put200 error response.
func (client *HTTPSuccessClient) put200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put201 - Put true Boolean value in request returns 201
func (client *HTTPSuccessClient) Put201(ctx context.Context, options *HTTPSuccessPut201Options) (*http.Response, error) {
	req, err := client.put201CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.put201HandleError(resp)
	}
	return resp.Response, nil
}

// put201CreateRequest creates the Put201 request.
func (client *HTTPSuccessClient) put201CreateRequest(ctx context.Context, options *HTTPSuccessPut201Options) (*azcore.Request, error) {
	urlPath := "/http/success/201"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put201HandleError handles the Put201 error response.
func (client *HTTPSuccessClient) put201HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put202 - Put true Boolean value in request returns 202 (Accepted)
func (client *HTTPSuccessClient) Put202(ctx context.Context, options *HTTPSuccessPut202Options) (*http.Response, error) {
	req, err := client.put202CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.put202HandleError(resp)
	}
	return resp.Response, nil
}

// put202CreateRequest creates the Put202 request.
func (client *HTTPSuccessClient) put202CreateRequest(ctx context.Context, options *HTTPSuccessPut202Options) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put202HandleError handles the Put202 error response.
func (client *HTTPSuccessClient) put202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// Put204 - Put true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Put204(ctx context.Context, options *HTTPSuccessPut204Options) (*http.Response, error) {
	req, err := client.put204CreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.put204HandleError(resp)
	}
	return resp.Response, nil
}

// put204CreateRequest creates the Put204 request.
func (client *HTTPSuccessClient) put204CreateRequest(ctx context.Context, options *HTTPSuccessPut204Options) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// put204HandleError handles the Put204 error response.
func (client *HTTPSuccessClient) put204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
