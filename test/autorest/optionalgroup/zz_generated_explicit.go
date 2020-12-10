// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"strings"
)

// ExplicitClient contains the methods for the Explicit group.
// Don't use this type directly, use NewExplicitClient() instead.
type ExplicitClient struct {
	con *Connection
}

// NewExplicitClient creates a new instance of ExplicitClient with the specified values.
func NewExplicitClient(con *Connection) *ExplicitClient {
	return &ExplicitClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *ExplicitClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// PostOptionalArrayHeader - Test explicitly optional integer. Please put a header 'headerParameter' => null.
func (client *ExplicitClient) PostOptionalArrayHeader(ctx context.Context, options *ExplicitPostOptionalArrayHeaderOptions) (*http.Response, error) {
	req, err := client.postOptionalArrayHeaderCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalArrayHeaderHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalArrayHeaderCreateRequest creates the PostOptionalArrayHeader request.
func (client *ExplicitClient) postOptionalArrayHeaderCreateRequest(ctx context.Context, options *ExplicitPostOptionalArrayHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/array/header"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	if options != nil && options.HeaderParameter != nil {
		req.Header.Set("headerParameter", strings.Join(*options.HeaderParameter, ","))
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postOptionalArrayHeaderHandleError handles the PostOptionalArrayHeader error response.
func (client *ExplicitClient) postOptionalArrayHeaderHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalArrayParameter - Test explicitly optional array. Please put null.
func (client *ExplicitClient) PostOptionalArrayParameter(ctx context.Context, options *ExplicitPostOptionalArrayParameterOptions) (*http.Response, error) {
	req, err := client.postOptionalArrayParameterCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalArrayParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalArrayParameterCreateRequest creates the PostOptionalArrayParameter request.
func (client *ExplicitClient) postOptionalArrayParameterCreateRequest(ctx context.Context, options *ExplicitPostOptionalArrayParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/array/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalArrayParameterHandleError handles the PostOptionalArrayParameter error response.
func (client *ExplicitClient) postOptionalArrayParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalArrayProperty - Test explicitly optional array. Please put a valid array-wrapper with 'value' = null.
func (client *ExplicitClient) PostOptionalArrayProperty(ctx context.Context, options *ExplicitPostOptionalArrayPropertyOptions) (*http.Response, error) {
	req, err := client.postOptionalArrayPropertyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalArrayPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalArrayPropertyCreateRequest creates the PostOptionalArrayProperty request.
func (client *ExplicitClient) postOptionalArrayPropertyCreateRequest(ctx context.Context, options *ExplicitPostOptionalArrayPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/array/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalArrayPropertyHandleError handles the PostOptionalArrayProperty error response.
func (client *ExplicitClient) postOptionalArrayPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalClassParameter - Test explicitly optional complex object. Please put null.
func (client *ExplicitClient) PostOptionalClassParameter(ctx context.Context, options *ExplicitPostOptionalClassParameterOptions) (*http.Response, error) {
	req, err := client.postOptionalClassParameterCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalClassParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalClassParameterCreateRequest creates the PostOptionalClassParameter request.
func (client *ExplicitClient) postOptionalClassParameterCreateRequest(ctx context.Context, options *ExplicitPostOptionalClassParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/class/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalClassParameterHandleError handles the PostOptionalClassParameter error response.
func (client *ExplicitClient) postOptionalClassParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalClassProperty - Test explicitly optional complex object. Please put a valid class-wrapper with 'value' = null.
func (client *ExplicitClient) PostOptionalClassProperty(ctx context.Context, options *ExplicitPostOptionalClassPropertyOptions) (*http.Response, error) {
	req, err := client.postOptionalClassPropertyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalClassPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalClassPropertyCreateRequest creates the PostOptionalClassProperty request.
func (client *ExplicitClient) postOptionalClassPropertyCreateRequest(ctx context.Context, options *ExplicitPostOptionalClassPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/class/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalClassPropertyHandleError handles the PostOptionalClassProperty error response.
func (client *ExplicitClient) postOptionalClassPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalIntegerHeader - Test explicitly optional integer. Please put a header 'headerParameter' => null.
func (client *ExplicitClient) PostOptionalIntegerHeader(ctx context.Context, options *ExplicitPostOptionalIntegerHeaderOptions) (*http.Response, error) {
	req, err := client.postOptionalIntegerHeaderCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalIntegerHeaderHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalIntegerHeaderCreateRequest creates the PostOptionalIntegerHeader request.
func (client *ExplicitClient) postOptionalIntegerHeaderCreateRequest(ctx context.Context, options *ExplicitPostOptionalIntegerHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/integer/header"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	if options != nil && options.HeaderParameter != nil {
		req.Header.Set("headerParameter", strconv.FormatInt(int64(*options.HeaderParameter), 10))
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postOptionalIntegerHeaderHandleError handles the PostOptionalIntegerHeader error response.
func (client *ExplicitClient) postOptionalIntegerHeaderHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalIntegerParameter - Test explicitly optional integer. Please put null.
func (client *ExplicitClient) PostOptionalIntegerParameter(ctx context.Context, options *ExplicitPostOptionalIntegerParameterOptions) (*http.Response, error) {
	req, err := client.postOptionalIntegerParameterCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalIntegerParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalIntegerParameterCreateRequest creates the PostOptionalIntegerParameter request.
func (client *ExplicitClient) postOptionalIntegerParameterCreateRequest(ctx context.Context, options *ExplicitPostOptionalIntegerParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/integer/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalIntegerParameterHandleError handles the PostOptionalIntegerParameter error response.
func (client *ExplicitClient) postOptionalIntegerParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalIntegerProperty - Test explicitly optional integer. Please put a valid int-wrapper with 'value' = null.
func (client *ExplicitClient) PostOptionalIntegerProperty(ctx context.Context, options *ExplicitPostOptionalIntegerPropertyOptions) (*http.Response, error) {
	req, err := client.postOptionalIntegerPropertyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalIntegerPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalIntegerPropertyCreateRequest creates the PostOptionalIntegerProperty request.
func (client *ExplicitClient) postOptionalIntegerPropertyCreateRequest(ctx context.Context, options *ExplicitPostOptionalIntegerPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/integer/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalIntegerPropertyHandleError handles the PostOptionalIntegerProperty error response.
func (client *ExplicitClient) postOptionalIntegerPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalStringHeader - Test explicitly optional string. Please put a header 'headerParameter' => null.
func (client *ExplicitClient) PostOptionalStringHeader(ctx context.Context, options *ExplicitPostOptionalStringHeaderOptions) (*http.Response, error) {
	req, err := client.postOptionalStringHeaderCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalStringHeaderHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalStringHeaderCreateRequest creates the PostOptionalStringHeader request.
func (client *ExplicitClient) postOptionalStringHeaderCreateRequest(ctx context.Context, options *ExplicitPostOptionalStringHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/string/header"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	if options != nil && options.BodyParameter != nil {
		req.Header.Set("bodyParameter", *options.BodyParameter)
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postOptionalStringHeaderHandleError handles the PostOptionalStringHeader error response.
func (client *ExplicitClient) postOptionalStringHeaderHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalStringParameter - Test explicitly optional string. Please put null.
func (client *ExplicitClient) PostOptionalStringParameter(ctx context.Context, options *ExplicitPostOptionalStringParameterOptions) (*http.Response, error) {
	req, err := client.postOptionalStringParameterCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalStringParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalStringParameterCreateRequest creates the PostOptionalStringParameter request.
func (client *ExplicitClient) postOptionalStringParameterCreateRequest(ctx context.Context, options *ExplicitPostOptionalStringParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/string/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalStringParameterHandleError handles the PostOptionalStringParameter error response.
func (client *ExplicitClient) postOptionalStringParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostOptionalStringProperty - Test explicitly optional integer. Please put a valid string-wrapper with 'value' = null.
func (client *ExplicitClient) PostOptionalStringProperty(ctx context.Context, options *ExplicitPostOptionalStringPropertyOptions) (*http.Response, error) {
	req, err := client.postOptionalStringPropertyCreateRequest(ctx, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postOptionalStringPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postOptionalStringPropertyCreateRequest creates the PostOptionalStringProperty request.
func (client *ExplicitClient) postOptionalStringPropertyCreateRequest(ctx context.Context, options *ExplicitPostOptionalStringPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/optional/string/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	if options != nil {
		return req, req.MarshalAsJSON(options.BodyParameter)
	}
	return req, nil
}

// postOptionalStringPropertyHandleError handles the PostOptionalStringProperty error response.
func (client *ExplicitClient) postOptionalStringPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredArrayHeader - Test explicitly required array. Please put a header 'headerParameter' => null and the client library should throw before the
// request is sent.
func (client *ExplicitClient) PostRequiredArrayHeader(ctx context.Context, headerParameter []string, options *ExplicitPostRequiredArrayHeaderOptions) (*http.Response, error) {
	req, err := client.postRequiredArrayHeaderCreateRequest(ctx, headerParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredArrayHeaderHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredArrayHeaderCreateRequest creates the PostRequiredArrayHeader request.
func (client *ExplicitClient) postRequiredArrayHeaderCreateRequest(ctx context.Context, headerParameter []string, options *ExplicitPostRequiredArrayHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/array/header"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("headerParameter", strings.Join(headerParameter, ","))
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postRequiredArrayHeaderHandleError handles the PostRequiredArrayHeader error response.
func (client *ExplicitClient) postRequiredArrayHeaderHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredArrayParameter - Test explicitly required array. Please put null and the client library should throw before the request is sent.
func (client *ExplicitClient) PostRequiredArrayParameter(ctx context.Context, bodyParameter []string, options *ExplicitPostRequiredArrayParameterOptions) (*http.Response, error) {
	req, err := client.postRequiredArrayParameterCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredArrayParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredArrayParameterCreateRequest creates the PostRequiredArrayParameter request.
func (client *ExplicitClient) postRequiredArrayParameterCreateRequest(ctx context.Context, bodyParameter []string, options *ExplicitPostRequiredArrayParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/array/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredArrayParameterHandleError handles the PostRequiredArrayParameter error response.
func (client *ExplicitClient) postRequiredArrayParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredArrayProperty - Test explicitly required array. Please put a valid array-wrapper with 'value' = null and the client library should throw
// before the request is sent.
func (client *ExplicitClient) PostRequiredArrayProperty(ctx context.Context, bodyParameter ArrayWrapper, options *ExplicitPostRequiredArrayPropertyOptions) (*http.Response, error) {
	req, err := client.postRequiredArrayPropertyCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredArrayPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredArrayPropertyCreateRequest creates the PostRequiredArrayProperty request.
func (client *ExplicitClient) postRequiredArrayPropertyCreateRequest(ctx context.Context, bodyParameter ArrayWrapper, options *ExplicitPostRequiredArrayPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/array/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredArrayPropertyHandleError handles the PostRequiredArrayProperty error response.
func (client *ExplicitClient) postRequiredArrayPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredClassParameter - Test explicitly required complex object. Please put null and the client library should throw before the request is sent.
func (client *ExplicitClient) PostRequiredClassParameter(ctx context.Context, bodyParameter Product, options *ExplicitPostRequiredClassParameterOptions) (*http.Response, error) {
	req, err := client.postRequiredClassParameterCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredClassParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredClassParameterCreateRequest creates the PostRequiredClassParameter request.
func (client *ExplicitClient) postRequiredClassParameterCreateRequest(ctx context.Context, bodyParameter Product, options *ExplicitPostRequiredClassParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/class/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredClassParameterHandleError handles the PostRequiredClassParameter error response.
func (client *ExplicitClient) postRequiredClassParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredClassProperty - Test explicitly required complex object. Please put a valid class-wrapper with 'value' = null and the client library should
// throw before the request is sent.
func (client *ExplicitClient) PostRequiredClassProperty(ctx context.Context, bodyParameter ClassWrapper, options *ExplicitPostRequiredClassPropertyOptions) (*http.Response, error) {
	req, err := client.postRequiredClassPropertyCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredClassPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredClassPropertyCreateRequest creates the PostRequiredClassProperty request.
func (client *ExplicitClient) postRequiredClassPropertyCreateRequest(ctx context.Context, bodyParameter ClassWrapper, options *ExplicitPostRequiredClassPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/class/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredClassPropertyHandleError handles the PostRequiredClassProperty error response.
func (client *ExplicitClient) postRequiredClassPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredIntegerHeader - Test explicitly required integer. Please put a header 'headerParameter' => null and the client library should throw before
// the request is sent.
func (client *ExplicitClient) PostRequiredIntegerHeader(ctx context.Context, headerParameter int32, options *ExplicitPostRequiredIntegerHeaderOptions) (*http.Response, error) {
	req, err := client.postRequiredIntegerHeaderCreateRequest(ctx, headerParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredIntegerHeaderHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredIntegerHeaderCreateRequest creates the PostRequiredIntegerHeader request.
func (client *ExplicitClient) postRequiredIntegerHeaderCreateRequest(ctx context.Context, headerParameter int32, options *ExplicitPostRequiredIntegerHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/integer/header"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("headerParameter", strconv.FormatInt(int64(headerParameter), 10))
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postRequiredIntegerHeaderHandleError handles the PostRequiredIntegerHeader error response.
func (client *ExplicitClient) postRequiredIntegerHeaderHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredIntegerParameter - Test explicitly required integer. Please put null and the client library should throw before the request is sent.
func (client *ExplicitClient) PostRequiredIntegerParameter(ctx context.Context, bodyParameter int32, options *ExplicitPostRequiredIntegerParameterOptions) (*http.Response, error) {
	req, err := client.postRequiredIntegerParameterCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredIntegerParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredIntegerParameterCreateRequest creates the PostRequiredIntegerParameter request.
func (client *ExplicitClient) postRequiredIntegerParameterCreateRequest(ctx context.Context, bodyParameter int32, options *ExplicitPostRequiredIntegerParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/integer/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredIntegerParameterHandleError handles the PostRequiredIntegerParameter error response.
func (client *ExplicitClient) postRequiredIntegerParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredIntegerProperty - Test explicitly required integer. Please put a valid int-wrapper with 'value' = null and the client library should throw
// before the request is sent.
func (client *ExplicitClient) PostRequiredIntegerProperty(ctx context.Context, bodyParameter IntWrapper, options *ExplicitPostRequiredIntegerPropertyOptions) (*http.Response, error) {
	req, err := client.postRequiredIntegerPropertyCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredIntegerPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredIntegerPropertyCreateRequest creates the PostRequiredIntegerProperty request.
func (client *ExplicitClient) postRequiredIntegerPropertyCreateRequest(ctx context.Context, bodyParameter IntWrapper, options *ExplicitPostRequiredIntegerPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/integer/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredIntegerPropertyHandleError handles the PostRequiredIntegerProperty error response.
func (client *ExplicitClient) postRequiredIntegerPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredStringHeader - Test explicitly required string. Please put a header 'headerParameter' => null and the client library should throw before
// the request is sent.
func (client *ExplicitClient) PostRequiredStringHeader(ctx context.Context, headerParameter string, options *ExplicitPostRequiredStringHeaderOptions) (*http.Response, error) {
	req, err := client.postRequiredStringHeaderCreateRequest(ctx, headerParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredStringHeaderHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredStringHeaderCreateRequest creates the PostRequiredStringHeader request.
func (client *ExplicitClient) postRequiredStringHeaderCreateRequest(ctx context.Context, headerParameter string, options *ExplicitPostRequiredStringHeaderOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/string/header"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("headerParameter", headerParameter)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// postRequiredStringHeaderHandleError handles the PostRequiredStringHeader error response.
func (client *ExplicitClient) postRequiredStringHeaderHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredStringParameter - Test explicitly required string. Please put null and the client library should throw before the request is sent.
func (client *ExplicitClient) PostRequiredStringParameter(ctx context.Context, bodyParameter string, options *ExplicitPostRequiredStringParameterOptions) (*http.Response, error) {
	req, err := client.postRequiredStringParameterCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredStringParameterHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredStringParameterCreateRequest creates the PostRequiredStringParameter request.
func (client *ExplicitClient) postRequiredStringParameterCreateRequest(ctx context.Context, bodyParameter string, options *ExplicitPostRequiredStringParameterOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/string/parameter"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredStringParameterHandleError handles the PostRequiredStringParameter error response.
func (client *ExplicitClient) postRequiredStringParameterHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PostRequiredStringProperty - Test explicitly required string. Please put a valid string-wrapper with 'value' = null and the client library should throw
// before the request is sent.
func (client *ExplicitClient) PostRequiredStringProperty(ctx context.Context, bodyParameter StringWrapper, options *ExplicitPostRequiredStringPropertyOptions) (*http.Response, error) {
	req, err := client.postRequiredStringPropertyCreateRequest(ctx, bodyParameter, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.postRequiredStringPropertyHandleError(resp)
	}
	return resp.Response, nil
}

// postRequiredStringPropertyCreateRequest creates the PostRequiredStringProperty request.
func (client *ExplicitClient) postRequiredStringPropertyCreateRequest(ctx context.Context, bodyParameter StringWrapper, options *ExplicitPostRequiredStringPropertyOptions) (*azcore.Request, error) {
	urlPath := "/reqopt/requied/string/property"
	req, err := azcore.NewRequest(ctx, http.MethodPost, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(bodyParameter)
}

// postRequiredStringPropertyHandleError handles the PostRequiredStringProperty error response.
func (client *ExplicitClient) postRequiredStringPropertyHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
