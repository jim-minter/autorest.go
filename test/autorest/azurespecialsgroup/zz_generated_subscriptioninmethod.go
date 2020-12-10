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
	"net/url"
	"strings"
)

// SubscriptionInMethodClient contains the methods for the SubscriptionInMethod group.
// Don't use this type directly, use NewSubscriptionInMethodClient() instead.
type SubscriptionInMethodClient struct {
	con *Connection
}

// NewSubscriptionInMethodClient creates a new instance of SubscriptionInMethodClient with the specified values.
func NewSubscriptionInMethodClient(con *Connection) *SubscriptionInMethodClient {
	return &SubscriptionInMethodClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *SubscriptionInMethodClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// PostMethodLocalNull - POST method with subscriptionId modeled in the method. pass in subscription id = null, client-side validation should prevent you
// from making this call
func (client *SubscriptionInMethodClient) PostMethodLocalNull(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalNullOptions) (*http.Response, error) {
	req, err := client.postMethodLocalNullCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postMethodLocalNullHandleError(resp)
	}
	return resp.Response, nil
}

// postMethodLocalNullCreateRequest creates the PostMethodLocalNull request.
func (client *SubscriptionInMethodClient) postMethodLocalNullCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalNullOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/null/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postMethodLocalNullHandleError handles the PostMethodLocalNull error response.
func (client *SubscriptionInMethodClient) postMethodLocalNullHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostMethodLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *SubscriptionInMethodClient) PostMethodLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalValidOptions) (*http.Response, error) {
	req, err := client.postMethodLocalValidCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postMethodLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// postMethodLocalValidCreateRequest creates the PostMethodLocalValid request.
func (client *SubscriptionInMethodClient) postMethodLocalValidCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostMethodLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/method/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postMethodLocalValidHandleError handles the PostMethodLocalValid error response.
func (client *SubscriptionInMethodClient) postMethodLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostPathLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *SubscriptionInMethodClient) PostPathLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostPathLocalValidOptions) (*http.Response, error) {
	req, err := client.postPathLocalValidCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postPathLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// postPathLocalValidCreateRequest creates the PostPathLocalValid request.
func (client *SubscriptionInMethodClient) postPathLocalValidCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostPathLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/path/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postPathLocalValidHandleError handles the PostPathLocalValid error response.
func (client *SubscriptionInMethodClient) postPathLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostSwaggerLocalValid - POST method with subscriptionId modeled in the method. pass in subscription id = '1234-5678-9012-3456' to succeed
func (client *SubscriptionInMethodClient) PostSwaggerLocalValid(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostSwaggerLocalValidOptions) (*http.Response, error) {
	req, err := client.postSwaggerLocalValidCreateRequest(ctx, subscriptionId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postSwaggerLocalValidHandleError(resp)
	}
	return resp.Response, nil
}

// postSwaggerLocalValidCreateRequest creates the PostSwaggerLocalValid request.
func (client *SubscriptionInMethodClient) postSwaggerLocalValidCreateRequest(ctx context.Context, subscriptionId string, options *SubscriptionInMethodPostSwaggerLocalValidOptions) (*azcore.Request, error) {
	urlPath := "/azurespecials/subscriptionId/swagger/string/none/path/local/1234-5678-9012-3456/{subscriptionId}"
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(subscriptionId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postSwaggerLocalValidHandleError handles the PostSwaggerLocalValid error response.
func (client *SubscriptionInMethodClient) postSwaggerLocalValidHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
