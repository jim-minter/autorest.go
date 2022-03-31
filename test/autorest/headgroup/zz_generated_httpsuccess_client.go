//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package headgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPSuccessClient contains the methods for the HTTPSuccess group.
// Don't use this type directly, use NewHTTPSuccessClient() instead.
type HTTPSuccessClient struct {
	pl runtime.Pipeline
}

// NewHTTPSuccessClient creates a new instance of HTTPSuccessClient with the specified values.
// options - pass nil to accept the default values.
func NewHTTPSuccessClient(options *azcore.ClientOptions) *HTTPSuccessClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &HTTPSuccessClient{
		pl: runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// Head200 - Return 200 status code if successful
// options - HTTPSuccessClientHead200Options contains the optional parameters for the HTTPSuccessClient.Head200 method.
func (client *HTTPSuccessClient) Head200(ctx context.Context, options *HTTPSuccessClientHead200Options) (HTTPSuccessClientHead200Response, error) {
	req, err := client.head200CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientHead200Response{}, err
	}
	result := HTTPSuccessClientHead200Response{}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head200CreateRequest creates the Head200 request.
func (client *HTTPSuccessClient) head200CreateRequest(ctx context.Context, options *HTTPSuccessClientHead200Options) (*policy.Request, error) {
	urlPath := "/http/success/200"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Head204 - Return 204 status code if successful
// options - HTTPSuccessClientHead204Options contains the optional parameters for the HTTPSuccessClient.Head204 method.
func (client *HTTPSuccessClient) Head204(ctx context.Context, options *HTTPSuccessClientHead204Options) (HTTPSuccessClientHead204Response, error) {
	req, err := client.head204CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientHead204Response{}, err
	}
	result := HTTPSuccessClientHead204Response{}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head204CreateRequest creates the Head204 request.
func (client *HTTPSuccessClient) head204CreateRequest(ctx context.Context, options *HTTPSuccessClientHead204Options) (*policy.Request, error) {
	urlPath := "/http/success/204"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}

// Head404 - Return 404 status code if successful
// options - HTTPSuccessClientHead404Options contains the optional parameters for the HTTPSuccessClient.Head404 method.
func (client *HTTPSuccessClient) Head404(ctx context.Context, options *HTTPSuccessClientHead404Options) (HTTPSuccessClientHead404Response, error) {
	req, err := client.head404CreateRequest(ctx, options)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return HTTPSuccessClientHead404Response{}, err
	}
	result := HTTPSuccessClientHead404Response{}
	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		result.Success = true
	}
	return result, nil
}

// head404CreateRequest creates the Head404 request.
func (client *HTTPSuccessClient) head404CreateRequest(ctx context.Context, options *HTTPSuccessClientHead404Options) (*policy.Request, error) {
	urlPath := "/http/success/404"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	return req, nil
}
