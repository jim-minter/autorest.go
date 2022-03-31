//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strings"
)

// SubscriptionInCredentialsClient contains the methods for the SubscriptionInCredentials group.
// Don't use this type directly, use NewSubscriptionInCredentialsClient() instead.
type SubscriptionInCredentialsClient struct {
	subscriptionID string
	pl             runtime.Pipeline
}

// NewSubscriptionInCredentialsClient creates a new instance of SubscriptionInCredentialsClient with the specified values.
// subscriptionID - The subscription id, which appears in the path, always modeled in credentials. The value is always '1234-5678-9012-3456'
// options - pass nil to accept the default values.
func NewSubscriptionInCredentialsClient(subscriptionID string, options *azcore.ClientOptions) *SubscriptionInCredentialsClient {
	if options == nil {
		options = &azcore.ClientOptions{}
	}
	client := &SubscriptionInCredentialsClient{
		subscriptionID: subscriptionID,
		pl:             runtime.NewPipeline(moduleName, moduleVersion, runtime.PipelineOptions{}, options),
	}
	return client
}

// PostMethodGlobalNotProvidedValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId
// to '1234-5678-9012-3456' to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidOptions contains the optional parameters for the
// SubscriptionInCredentialsClient.PostMethodGlobalNotProvidedValid method.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNotProvidedValid(ctx context.Context, options *SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidOptions) (SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidResponse, error) {
	req, err := client.postMethodGlobalNotProvidedValidCreateRequest(ctx, options)
	if err != nil {
		return SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidResponse{}, nil
}

// postMethodGlobalNotProvidedValidCreateRequest creates the PostMethodGlobalNotProvidedValid request.
func (client *SubscriptionInCredentialsClient) postMethodGlobalNotProvidedValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsClientPostMethodGlobalNotProvidedValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/globalNotProvided/1234-5678-9012-3456/{subscriptionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("api-version", "2015-07-01-preview")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PostMethodGlobalNull - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to null,
// and client-side validation should prevent you from making this call
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionInCredentialsClientPostMethodGlobalNullOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostMethodGlobalNull
// method.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalNull(ctx context.Context, options *SubscriptionInCredentialsClientPostMethodGlobalNullOptions) (SubscriptionInCredentialsClientPostMethodGlobalNullResponse, error) {
	req, err := client.postMethodGlobalNullCreateRequest(ctx, options)
	if err != nil {
		return SubscriptionInCredentialsClientPostMethodGlobalNullResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionInCredentialsClientPostMethodGlobalNullResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInCredentialsClientPostMethodGlobalNullResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInCredentialsClientPostMethodGlobalNullResponse{}, nil
}

// postMethodGlobalNullCreateRequest creates the PostMethodGlobalNull request.
func (client *SubscriptionInCredentialsClient) postMethodGlobalNullCreateRequest(ctx context.Context, options *SubscriptionInCredentialsClientPostMethodGlobalNullOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/null/{subscriptionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PostMethodGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456'
// to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionInCredentialsClientPostMethodGlobalValidOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostMethodGlobalValid
// method.
func (client *SubscriptionInCredentialsClient) PostMethodGlobalValid(ctx context.Context, options *SubscriptionInCredentialsClientPostMethodGlobalValidOptions) (SubscriptionInCredentialsClientPostMethodGlobalValidResponse, error) {
	req, err := client.postMethodGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return SubscriptionInCredentialsClientPostMethodGlobalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionInCredentialsClientPostMethodGlobalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInCredentialsClientPostMethodGlobalValidResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInCredentialsClientPostMethodGlobalValidResponse{}, nil
}

// postMethodGlobalValidCreateRequest creates the PostMethodGlobalValid request.
func (client *SubscriptionInCredentialsClient) postMethodGlobalValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsClientPostMethodGlobalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PostPathGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456'
// to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionInCredentialsClientPostPathGlobalValidOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostPathGlobalValid
// method.
func (client *SubscriptionInCredentialsClient) PostPathGlobalValid(ctx context.Context, options *SubscriptionInCredentialsClientPostPathGlobalValidOptions) (SubscriptionInCredentialsClientPostPathGlobalValidResponse, error) {
	req, err := client.postPathGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return SubscriptionInCredentialsClientPostPathGlobalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionInCredentialsClientPostPathGlobalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInCredentialsClientPostPathGlobalValidResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInCredentialsClientPostPathGlobalValidResponse{}, nil
}

// postPathGlobalValidCreateRequest creates the PostPathGlobalValid request.
func (client *SubscriptionInCredentialsClient) postPathGlobalValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsClientPostPathGlobalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// PostSwaggerGlobalValid - POST method with subscriptionId modeled in credentials. Set the credential subscriptionId to '1234-5678-9012-3456'
// to succeed
// If the operation fails it returns an *azcore.ResponseError type.
// options - SubscriptionInCredentialsClientPostSwaggerGlobalValidOptions contains the optional parameters for the SubscriptionInCredentialsClient.PostSwaggerGlobalValid
// method.
func (client *SubscriptionInCredentialsClient) PostSwaggerGlobalValid(ctx context.Context, options *SubscriptionInCredentialsClientPostSwaggerGlobalValidOptions) (SubscriptionInCredentialsClientPostSwaggerGlobalValidResponse, error) {
	req, err := client.postSwaggerGlobalValidCreateRequest(ctx, options)
	if err != nil {
		return SubscriptionInCredentialsClientPostSwaggerGlobalValidResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return SubscriptionInCredentialsClientPostSwaggerGlobalValidResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return SubscriptionInCredentialsClientPostSwaggerGlobalValidResponse{}, runtime.NewResponseError(resp)
	}
	return SubscriptionInCredentialsClientPostSwaggerGlobalValidResponse{}, nil
}

// postSwaggerGlobalValidCreateRequest creates the PostSwaggerGlobalValid request.
func (client *SubscriptionInCredentialsClient) postSwaggerGlobalValidCreateRequest(ctx context.Context, options *SubscriptionInCredentialsClientPostSwaggerGlobalValidOptions) (*policy.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/global/1234-5678-9012-3456/{subscriptionId}"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}
