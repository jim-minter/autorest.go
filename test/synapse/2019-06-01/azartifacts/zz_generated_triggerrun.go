// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azartifacts

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"strings"
)

type triggerRunClient struct {
	con *connection
}

// Pipeline returns the pipeline associated with this client.
func (client *triggerRunClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// CancelTriggerInstance - Cancel single trigger instance by runId.
func (client *triggerRunClient) CancelTriggerInstance(ctx context.Context, triggerName string, runId string, options *TriggerRunCancelTriggerInstanceOptions) (*http.Response, error) {
	req, err := client.cancelTriggerInstanceCreateRequest(ctx, triggerName, runId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.cancelTriggerInstanceHandleError(resp)
	}
	return resp.Response, nil
}

// cancelTriggerInstanceCreateRequest creates the CancelTriggerInstance request.
func (client *triggerRunClient) cancelTriggerInstanceCreateRequest(ctx context.Context, triggerName string, runId string, options *TriggerRunCancelTriggerInstanceOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/triggerRuns/{runId}/cancel"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// cancelTriggerInstanceHandleError handles the CancelTriggerInstance error response.
func (client *triggerRunClient) cancelTriggerInstanceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// QueryTriggerRunsByWorkspace - Query trigger runs.
func (client *triggerRunClient) QueryTriggerRunsByWorkspace(ctx context.Context, filterParameters RunFilterParameters, options *TriggerRunQueryTriggerRunsByWorkspaceOptions) (TriggerRunsQueryResponseResponse, error) {
	req, err := client.queryTriggerRunsByWorkspaceCreateRequest(ctx, filterParameters, options)
	if err != nil {
		return TriggerRunsQueryResponseResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return TriggerRunsQueryResponseResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return TriggerRunsQueryResponseResponse{}, client.queryTriggerRunsByWorkspaceHandleError(resp)
	}
	return client.queryTriggerRunsByWorkspaceHandleResponse(resp)
}

// queryTriggerRunsByWorkspaceCreateRequest creates the QueryTriggerRunsByWorkspace request.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceCreateRequest(ctx context.Context, filterParameters RunFilterParameters, options *TriggerRunQueryTriggerRunsByWorkspaceOptions) (*azcore.Request, error) {
	urlPath := "/queryTriggerRuns"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(filterParameters)
}

// queryTriggerRunsByWorkspaceHandleResponse handles the QueryTriggerRunsByWorkspace response.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceHandleResponse(resp *azcore.Response) (TriggerRunsQueryResponseResponse, error) {
	var val *TriggerRunsQueryResponse
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return TriggerRunsQueryResponseResponse{}, err
	}
	return TriggerRunsQueryResponseResponse{RawResponse: resp.Response, TriggerRunsQueryResponse: val}, nil
}

// queryTriggerRunsByWorkspaceHandleError handles the QueryTriggerRunsByWorkspace error response.
func (client *triggerRunClient) queryTriggerRunsByWorkspaceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// RerunTriggerInstance - Rerun single trigger instance by runId.
func (client *triggerRunClient) RerunTriggerInstance(ctx context.Context, triggerName string, runId string, options *TriggerRunRerunTriggerInstanceOptions) (*http.Response, error) {
	req, err := client.rerunTriggerInstanceCreateRequest(ctx, triggerName, runId, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.rerunTriggerInstanceHandleError(resp)
	}
	return resp.Response, nil
}

// rerunTriggerInstanceCreateRequest creates the RerunTriggerInstance request.
func (client *triggerRunClient) rerunTriggerInstanceCreateRequest(ctx context.Context, triggerName string, runId string, options *TriggerRunRerunTriggerInstanceOptions) (*azcore.Request, error) {
	urlPath := "/triggers/{triggerName}/triggerRuns/{runId}/rerun"
	urlPath = strings.ReplaceAll(urlPath, "{triggerName}", url.PathEscape(triggerName))
	urlPath = strings.ReplaceAll(urlPath, "{runId}", url.PathEscape(runId))
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	query := req.URL.Query()
	query.Set("api-version", "2019-06-01-preview")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// rerunTriggerInstanceHandleError handles the RerunTriggerInstance error response.
func (client *triggerRunClient) rerunTriggerInstanceHandleError(resp *azcore.Response) error {
	var err CloudError
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
