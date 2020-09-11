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

// HTTPSuccessOperations contains the methods for the HTTPSuccess group.
type HTTPSuccessOperations interface {
	// Delete200 - Delete simple boolean value true returns 200
	Delete200(ctx context.Context) (*http.Response, error)
	// Delete202 - Delete true Boolean value in request returns 202 (accepted)
	Delete202(ctx context.Context) (*http.Response, error)
	// Delete204 - Delete true Boolean value in request returns 204 (no content)
	Delete204(ctx context.Context) (*http.Response, error)
	// Get200 - Get 200 success
	Get200(ctx context.Context) (*BoolResponse, error)
	// Head200 - Return 200 status code if successful
	Head200(ctx context.Context) (*http.Response, error)
	// Head204 - Return 204 status code if successful
	Head204(ctx context.Context) (*http.Response, error)
	// Head404 - Return 404 status code
	Head404(ctx context.Context) (*http.Response, error)
	// Options200 - Options 200 success
	Options200(ctx context.Context) (*BoolResponse, error)
	// Patch200 - Patch true Boolean value in request returning 200
	Patch200(ctx context.Context) (*http.Response, error)
	// Patch202 - Patch true Boolean value in request returns 202
	Patch202(ctx context.Context) (*http.Response, error)
	// Patch204 - Patch true Boolean value in request returns 204 (no content)
	Patch204(ctx context.Context) (*http.Response, error)
	// Post200 - Post bollean value true in request that returns a 200
	Post200(ctx context.Context) (*http.Response, error)
	// Post201 - Post true Boolean value in request returns 201 (Created)
	Post201(ctx context.Context) (*http.Response, error)
	// Post202 - Post true Boolean value in request returns 202 (Accepted)
	Post202(ctx context.Context) (*http.Response, error)
	// Post204 - Post true Boolean value in request returns 204 (no content)
	Post204(ctx context.Context) (*http.Response, error)
	// Put200 - Put boolean value true returning 200 success
	Put200(ctx context.Context) (*http.Response, error)
	// Put201 - Put true Boolean value in request returns 201
	Put201(ctx context.Context) (*http.Response, error)
	// Put202 - Put true Boolean value in request returns 202 (Accepted)
	Put202(ctx context.Context) (*http.Response, error)
	// Put204 - Put true Boolean value in request returns 204 (no content)
	Put204(ctx context.Context) (*http.Response, error)
}

// HTTPSuccessClient implements the HTTPSuccessOperations interface.
// Don't use this type directly, use NewHTTPSuccessClient() instead.
type HTTPSuccessClient struct {
	*Client
}

// NewHTTPSuccessClient creates a new instance of HTTPSuccessClient with the specified values.
func NewHTTPSuccessClient(c *Client) HTTPSuccessOperations {
	return &HTTPSuccessClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *HTTPSuccessClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// Delete200 - Delete simple boolean value true returns 200
func (client *HTTPSuccessClient) Delete200(ctx context.Context) (*http.Response, error) {
	req, err := client.Delete200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Delete200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete200CreateRequest creates the Delete200 request.
func (client *HTTPSuccessClient) Delete200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Delete200HandleResponse handles the Delete200 response.
func (client *HTTPSuccessClient) Delete200HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Delete200HandleError(resp)
	}
	return resp.Response, nil
}

// Delete200HandleError handles the Delete200 error response.
func (client *HTTPSuccessClient) Delete200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete202 - Delete true Boolean value in request returns 202 (accepted)
func (client *HTTPSuccessClient) Delete202(ctx context.Context) (*http.Response, error) {
	req, err := client.Delete202CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Delete202HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete202CreateRequest creates the Delete202 request.
func (client *HTTPSuccessClient) Delete202CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Delete202HandleResponse handles the Delete202 response.
func (client *HTTPSuccessClient) Delete202HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.Delete202HandleError(resp)
	}
	return resp.Response, nil
}

// Delete202HandleError handles the Delete202 error response.
func (client *HTTPSuccessClient) Delete202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Delete204 - Delete true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Delete204(ctx context.Context) (*http.Response, error) {
	req, err := client.Delete204CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Delete204HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Delete204CreateRequest creates the Delete204 request.
func (client *HTTPSuccessClient) Delete204CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodDelete, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Delete204HandleResponse handles the Delete204 response.
func (client *HTTPSuccessClient) Delete204HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.Delete204HandleError(resp)
	}
	return resp.Response, nil
}

// Delete204HandleError handles the Delete204 error response.
func (client *HTTPSuccessClient) Delete204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Get200 - Get 200 success
func (client *HTTPSuccessClient) Get200(ctx context.Context) (*BoolResponse, error) {
	req, err := client.Get200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Get200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Get200CreateRequest creates the Get200 request.
func (client *HTTPSuccessClient) Get200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Get200HandleResponse handles the Get200 response.
func (client *HTTPSuccessClient) Get200HandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Get200HandleError(resp)
	}
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// Get200HandleError handles the Get200 error response.
func (client *HTTPSuccessClient) Get200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head200 - Return 200 status code if successful
func (client *HTTPSuccessClient) Head200(ctx context.Context) (*http.Response, error) {
	req, err := client.Head200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Head200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Head200CreateRequest creates the Head200 request.
func (client *HTTPSuccessClient) Head200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head200HandleResponse handles the Head200 response.
func (client *HTTPSuccessClient) Head200HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Head200HandleError(resp)
	}
	return resp.Response, nil
}

// Head200HandleError handles the Head200 error response.
func (client *HTTPSuccessClient) Head200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head204 - Return 204 status code if successful
func (client *HTTPSuccessClient) Head204(ctx context.Context) (*http.Response, error) {
	req, err := client.Head204CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Head204HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Head204CreateRequest creates the Head204 request.
func (client *HTTPSuccessClient) Head204CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head204HandleResponse handles the Head204 response.
func (client *HTTPSuccessClient) Head204HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.Head204HandleError(resp)
	}
	return resp.Response, nil
}

// Head204HandleError handles the Head204 error response.
func (client *HTTPSuccessClient) Head204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Head404 - Return 404 status code
func (client *HTTPSuccessClient) Head404(ctx context.Context) (*http.Response, error) {
	req, err := client.Head404CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Head404HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Head404CreateRequest creates the Head404 request.
func (client *HTTPSuccessClient) Head404CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/404"
	req, err := azcore.NewRequest(ctx, http.MethodHead, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Head404HandleResponse handles the Head404 response.
func (client *HTTPSuccessClient) Head404HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusNoContent, http.StatusNotFound) {
		return nil, client.Head404HandleError(resp)
	}
	return resp.Response, nil
}

// Head404HandleError handles the Head404 error response.
func (client *HTTPSuccessClient) Head404HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Options200 - Options 200 success
func (client *HTTPSuccessClient) Options200(ctx context.Context) (*BoolResponse, error) {
	req, err := client.Options200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Options200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Options200CreateRequest creates the Options200 request.
func (client *HTTPSuccessClient) Options200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodOptions, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// Options200HandleResponse handles the Options200 response.
func (client *HTTPSuccessClient) Options200HandleResponse(resp *azcore.Response) (*BoolResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Options200HandleError(resp)
	}
	result := BoolResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// Options200HandleError handles the Options200 error response.
func (client *HTTPSuccessClient) Options200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch200 - Patch true Boolean value in request returning 200
func (client *HTTPSuccessClient) Patch200(ctx context.Context) (*http.Response, error) {
	req, err := client.Patch200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Patch200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Patch200CreateRequest creates the Patch200 request.
func (client *HTTPSuccessClient) Patch200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Patch200HandleResponse handles the Patch200 response.
func (client *HTTPSuccessClient) Patch200HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Patch200HandleError(resp)
	}
	return resp.Response, nil
}

// Patch200HandleError handles the Patch200 error response.
func (client *HTTPSuccessClient) Patch200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch202 - Patch true Boolean value in request returns 202
func (client *HTTPSuccessClient) Patch202(ctx context.Context) (*http.Response, error) {
	req, err := client.Patch202CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Patch202HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Patch202CreateRequest creates the Patch202 request.
func (client *HTTPSuccessClient) Patch202CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Patch202HandleResponse handles the Patch202 response.
func (client *HTTPSuccessClient) Patch202HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.Patch202HandleError(resp)
	}
	return resp.Response, nil
}

// Patch202HandleError handles the Patch202 error response.
func (client *HTTPSuccessClient) Patch202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Patch204 - Patch true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Patch204(ctx context.Context) (*http.Response, error) {
	req, err := client.Patch204CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Patch204HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Patch204CreateRequest creates the Patch204 request.
func (client *HTTPSuccessClient) Patch204CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodPatch, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Patch204HandleResponse handles the Patch204 response.
func (client *HTTPSuccessClient) Patch204HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.Patch204HandleError(resp)
	}
	return resp.Response, nil
}

// Patch204HandleError handles the Patch204 error response.
func (client *HTTPSuccessClient) Patch204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post200 - Post bollean value true in request that returns a 200
func (client *HTTPSuccessClient) Post200(ctx context.Context) (*http.Response, error) {
	req, err := client.Post200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Post200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Post200CreateRequest creates the Post200 request.
func (client *HTTPSuccessClient) Post200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Post200HandleResponse handles the Post200 response.
func (client *HTTPSuccessClient) Post200HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Post200HandleError(resp)
	}
	return resp.Response, nil
}

// Post200HandleError handles the Post200 error response.
func (client *HTTPSuccessClient) Post200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post201 - Post true Boolean value in request returns 201 (Created)
func (client *HTTPSuccessClient) Post201(ctx context.Context) (*http.Response, error) {
	req, err := client.Post201CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Post201HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Post201CreateRequest creates the Post201 request.
func (client *HTTPSuccessClient) Post201CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/201"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Post201HandleResponse handles the Post201 response.
func (client *HTTPSuccessClient) Post201HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.Post201HandleError(resp)
	}
	return resp.Response, nil
}

// Post201HandleError handles the Post201 error response.
func (client *HTTPSuccessClient) Post201HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post202 - Post true Boolean value in request returns 202 (Accepted)
func (client *HTTPSuccessClient) Post202(ctx context.Context) (*http.Response, error) {
	req, err := client.Post202CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Post202HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Post202CreateRequest creates the Post202 request.
func (client *HTTPSuccessClient) Post202CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Post202HandleResponse handles the Post202 response.
func (client *HTTPSuccessClient) Post202HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.Post202HandleError(resp)
	}
	return resp.Response, nil
}

// Post202HandleError handles the Post202 error response.
func (client *HTTPSuccessClient) Post202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Post204 - Post true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Post204(ctx context.Context) (*http.Response, error) {
	req, err := client.Post204CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Post204HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Post204CreateRequest creates the Post204 request.
func (client *HTTPSuccessClient) Post204CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Post204HandleResponse handles the Post204 response.
func (client *HTTPSuccessClient) Post204HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.Post204HandleError(resp)
	}
	return resp.Response, nil
}

// Post204HandleError handles the Post204 error response.
func (client *HTTPSuccessClient) Post204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put200 - Put boolean value true returning 200 success
func (client *HTTPSuccessClient) Put200(ctx context.Context) (*http.Response, error) {
	req, err := client.Put200CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Put200HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put200CreateRequest creates the Put200 request.
func (client *HTTPSuccessClient) Put200CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/200"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Put200HandleResponse handles the Put200 response.
func (client *HTTPSuccessClient) Put200HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.Put200HandleError(resp)
	}
	return resp.Response, nil
}

// Put200HandleError handles the Put200 error response.
func (client *HTTPSuccessClient) Put200HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put201 - Put true Boolean value in request returns 201
func (client *HTTPSuccessClient) Put201(ctx context.Context) (*http.Response, error) {
	req, err := client.Put201CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Put201HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put201CreateRequest creates the Put201 request.
func (client *HTTPSuccessClient) Put201CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/201"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Put201HandleResponse handles the Put201 response.
func (client *HTTPSuccessClient) Put201HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusCreated) {
		return nil, client.Put201HandleError(resp)
	}
	return resp.Response, nil
}

// Put201HandleError handles the Put201 error response.
func (client *HTTPSuccessClient) Put201HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put202 - Put true Boolean value in request returns 202 (Accepted)
func (client *HTTPSuccessClient) Put202(ctx context.Context) (*http.Response, error) {
	req, err := client.Put202CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Put202HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put202CreateRequest creates the Put202 request.
func (client *HTTPSuccessClient) Put202CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/202"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Put202HandleResponse handles the Put202 response.
func (client *HTTPSuccessClient) Put202HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.Put202HandleError(resp)
	}
	return resp.Response, nil
}

// Put202HandleError handles the Put202 error response.
func (client *HTTPSuccessClient) Put202HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// Put204 - Put true Boolean value in request returns 204 (no content)
func (client *HTTPSuccessClient) Put204(ctx context.Context) (*http.Response, error) {
	req, err := client.Put204CreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.Put204HandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// Put204CreateRequest creates the Put204 request.
func (client *HTTPSuccessClient) Put204CreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/http/success/204"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(true)
}

// Put204HandleResponse handles the Put204 response.
func (client *HTTPSuccessClient) Put204HandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusNoContent) {
		return nil, client.Put204HandleError(resp)
	}
	return resp.Response, nil
}

// Put204HandleError handles the Put204 error response.
func (client *HTTPSuccessClient) Put204HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
