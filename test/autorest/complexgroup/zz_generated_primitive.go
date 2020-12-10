// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// PrimitiveClient contains the methods for the Primitive group.
// Don't use this type directly, use NewPrimitiveClient() instead.
type PrimitiveClient struct {
	con *Connection
}

// NewPrimitiveClient creates a new instance of PrimitiveClient with the specified values.
func NewPrimitiveClient(con *Connection) *PrimitiveClient {
	return &PrimitiveClient{con: con}
}

// Pipeline returns the pipeline associated with this client.
func (client *PrimitiveClient) Pipeline() azcore.Pipeline {
	return client.con.Pipeline()
}

// GetBool - Get complex types with bool properties
func (client *PrimitiveClient) GetBool(ctx context.Context, options *PrimitiveGetBoolOptions) (BooleanWrapperResponse, error) {
	req, err := client.getBoolCreateRequest(ctx, options)
	if err != nil {
		return BooleanWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return BooleanWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return BooleanWrapperResponse{}, client.getBoolHandleError(resp)
	}
	return client.getBoolHandleResponse(resp)
}

// getBoolCreateRequest creates the GetBool request.
func (client *PrimitiveClient) getBoolCreateRequest(ctx context.Context, options *PrimitiveGetBoolOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/bool"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getBoolHandleResponse handles the GetBool response.
func (client *PrimitiveClient) getBoolHandleResponse(resp *azcore.Response) (BooleanWrapperResponse, error) {
	var val *BooleanWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return BooleanWrapperResponse{}, err
	}
	return BooleanWrapperResponse{RawResponse: resp.Response, BooleanWrapper: val}, nil
}

// getBoolHandleError handles the GetBool error response.
func (client *PrimitiveClient) getBoolHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetByte - Get complex types with byte properties
func (client *PrimitiveClient) GetByte(ctx context.Context, options *PrimitiveGetByteOptions) (ByteWrapperResponse, error) {
	req, err := client.getByteCreateRequest(ctx, options)
	if err != nil {
		return ByteWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return ByteWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return ByteWrapperResponse{}, client.getByteHandleError(resp)
	}
	return client.getByteHandleResponse(resp)
}

// getByteCreateRequest creates the GetByte request.
func (client *PrimitiveClient) getByteCreateRequest(ctx context.Context, options *PrimitiveGetByteOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/byte"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getByteHandleResponse handles the GetByte response.
func (client *PrimitiveClient) getByteHandleResponse(resp *azcore.Response) (ByteWrapperResponse, error) {
	var val *ByteWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return ByteWrapperResponse{}, err
	}
	return ByteWrapperResponse{RawResponse: resp.Response, ByteWrapper: val}, nil
}

// getByteHandleError handles the GetByte error response.
func (client *PrimitiveClient) getByteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDate - Get complex types with date properties
func (client *PrimitiveClient) GetDate(ctx context.Context, options *PrimitiveGetDateOptions) (DateWrapperResponse, error) {
	req, err := client.getDateCreateRequest(ctx, options)
	if err != nil {
		return DateWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return DateWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DateWrapperResponse{}, client.getDateHandleError(resp)
	}
	return client.getDateHandleResponse(resp)
}

// getDateCreateRequest creates the GetDate request.
func (client *PrimitiveClient) getDateCreateRequest(ctx context.Context, options *PrimitiveGetDateOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/date"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDateHandleResponse handles the GetDate response.
func (client *PrimitiveClient) getDateHandleResponse(resp *azcore.Response) (DateWrapperResponse, error) {
	var val *DateWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DateWrapperResponse{}, err
	}
	return DateWrapperResponse{RawResponse: resp.Response, DateWrapper: val}, nil
}

// getDateHandleError handles the GetDate error response.
func (client *PrimitiveClient) getDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDateTime - Get complex types with datetime properties
func (client *PrimitiveClient) GetDateTime(ctx context.Context, options *PrimitiveGetDateTimeOptions) (DatetimeWrapperResponse, error) {
	req, err := client.getDateTimeCreateRequest(ctx, options)
	if err != nil {
		return DatetimeWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return DatetimeWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DatetimeWrapperResponse{}, client.getDateTimeHandleError(resp)
	}
	return client.getDateTimeHandleResponse(resp)
}

// getDateTimeCreateRequest creates the GetDateTime request.
func (client *PrimitiveClient) getDateTimeCreateRequest(ctx context.Context, options *PrimitiveGetDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetime"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDateTimeHandleResponse handles the GetDateTime response.
func (client *PrimitiveClient) getDateTimeHandleResponse(resp *azcore.Response) (DatetimeWrapperResponse, error) {
	var val *DatetimeWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DatetimeWrapperResponse{}, err
	}
	return DatetimeWrapperResponse{RawResponse: resp.Response, DatetimeWrapper: val}, nil
}

// getDateTimeHandleError handles the GetDateTime error response.
func (client *PrimitiveClient) getDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDateTimeRFC1123 - Get complex types with datetimeRfc1123 properties
func (client *PrimitiveClient) GetDateTimeRFC1123(ctx context.Context, options *PrimitiveGetDateTimeRFC1123Options) (Datetimerfc1123WrapperResponse, error) {
	req, err := client.getDateTimeRfc1123CreateRequest(ctx, options)
	if err != nil {
		return Datetimerfc1123WrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return Datetimerfc1123WrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return Datetimerfc1123WrapperResponse{}, client.getDateTimeRfc1123HandleError(resp)
	}
	return client.getDateTimeRfc1123HandleResponse(resp)
}

// getDateTimeRfc1123CreateRequest creates the GetDateTimeRFC1123 request.
func (client *PrimitiveClient) getDateTimeRfc1123CreateRequest(ctx context.Context, options *PrimitiveGetDateTimeRFC1123Options) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDateTimeRfc1123HandleResponse handles the GetDateTimeRFC1123 response.
func (client *PrimitiveClient) getDateTimeRfc1123HandleResponse(resp *azcore.Response) (Datetimerfc1123WrapperResponse, error) {
	var val *Datetimerfc1123Wrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return Datetimerfc1123WrapperResponse{}, err
	}
	return Datetimerfc1123WrapperResponse{RawResponse: resp.Response, Datetimerfc1123Wrapper: val}, nil
}

// getDateTimeRfc1123HandleError handles the GetDateTimeRFC1123 error response.
func (client *PrimitiveClient) getDateTimeRfc1123HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDouble - Get complex types with double properties
func (client *PrimitiveClient) GetDouble(ctx context.Context, options *PrimitiveGetDoubleOptions) (DoubleWrapperResponse, error) {
	req, err := client.getDoubleCreateRequest(ctx, options)
	if err != nil {
		return DoubleWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return DoubleWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DoubleWrapperResponse{}, client.getDoubleHandleError(resp)
	}
	return client.getDoubleHandleResponse(resp)
}

// getDoubleCreateRequest creates the GetDouble request.
func (client *PrimitiveClient) getDoubleCreateRequest(ctx context.Context, options *PrimitiveGetDoubleOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/double"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDoubleHandleResponse handles the GetDouble response.
func (client *PrimitiveClient) getDoubleHandleResponse(resp *azcore.Response) (DoubleWrapperResponse, error) {
	var val *DoubleWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DoubleWrapperResponse{}, err
	}
	return DoubleWrapperResponse{RawResponse: resp.Response, DoubleWrapper: val}, nil
}

// getDoubleHandleError handles the GetDouble error response.
func (client *PrimitiveClient) getDoubleHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetDuration - Get complex types with duration properties
func (client *PrimitiveClient) GetDuration(ctx context.Context, options *PrimitiveGetDurationOptions) (DurationWrapperResponse, error) {
	req, err := client.getDurationCreateRequest(ctx, options)
	if err != nil {
		return DurationWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return DurationWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return DurationWrapperResponse{}, client.getDurationHandleError(resp)
	}
	return client.getDurationHandleResponse(resp)
}

// getDurationCreateRequest creates the GetDuration request.
func (client *PrimitiveClient) getDurationCreateRequest(ctx context.Context, options *PrimitiveGetDurationOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/duration"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getDurationHandleResponse handles the GetDuration response.
func (client *PrimitiveClient) getDurationHandleResponse(resp *azcore.Response) (DurationWrapperResponse, error) {
	var val *DurationWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return DurationWrapperResponse{}, err
	}
	return DurationWrapperResponse{RawResponse: resp.Response, DurationWrapper: val}, nil
}

// getDurationHandleError handles the GetDuration error response.
func (client *PrimitiveClient) getDurationHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetFloat - Get complex types with float properties
func (client *PrimitiveClient) GetFloat(ctx context.Context, options *PrimitiveGetFloatOptions) (FloatWrapperResponse, error) {
	req, err := client.getFloatCreateRequest(ctx, options)
	if err != nil {
		return FloatWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return FloatWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return FloatWrapperResponse{}, client.getFloatHandleError(resp)
	}
	return client.getFloatHandleResponse(resp)
}

// getFloatCreateRequest creates the GetFloat request.
func (client *PrimitiveClient) getFloatCreateRequest(ctx context.Context, options *PrimitiveGetFloatOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/float"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getFloatHandleResponse handles the GetFloat response.
func (client *PrimitiveClient) getFloatHandleResponse(resp *azcore.Response) (FloatWrapperResponse, error) {
	var val *FloatWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return FloatWrapperResponse{}, err
	}
	return FloatWrapperResponse{RawResponse: resp.Response, FloatWrapper: val}, nil
}

// getFloatHandleError handles the GetFloat error response.
func (client *PrimitiveClient) getFloatHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetInt - Get complex types with integer properties
func (client *PrimitiveClient) GetInt(ctx context.Context, options *PrimitiveGetIntOptions) (IntWrapperResponse, error) {
	req, err := client.getIntCreateRequest(ctx, options)
	if err != nil {
		return IntWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return IntWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return IntWrapperResponse{}, client.getIntHandleError(resp)
	}
	return client.getIntHandleResponse(resp)
}

// getIntCreateRequest creates the GetInt request.
func (client *PrimitiveClient) getIntCreateRequest(ctx context.Context, options *PrimitiveGetIntOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/integer"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getIntHandleResponse handles the GetInt response.
func (client *PrimitiveClient) getIntHandleResponse(resp *azcore.Response) (IntWrapperResponse, error) {
	var val *IntWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return IntWrapperResponse{}, err
	}
	return IntWrapperResponse{RawResponse: resp.Response, IntWrapper: val}, nil
}

// getIntHandleError handles the GetInt error response.
func (client *PrimitiveClient) getIntHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetLong - Get complex types with long properties
func (client *PrimitiveClient) GetLong(ctx context.Context, options *PrimitiveGetLongOptions) (LongWrapperResponse, error) {
	req, err := client.getLongCreateRequest(ctx, options)
	if err != nil {
		return LongWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return LongWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return LongWrapperResponse{}, client.getLongHandleError(resp)
	}
	return client.getLongHandleResponse(resp)
}

// getLongCreateRequest creates the GetLong request.
func (client *PrimitiveClient) getLongCreateRequest(ctx context.Context, options *PrimitiveGetLongOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/long"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getLongHandleResponse handles the GetLong response.
func (client *PrimitiveClient) getLongHandleResponse(resp *azcore.Response) (LongWrapperResponse, error) {
	var val *LongWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return LongWrapperResponse{}, err
	}
	return LongWrapperResponse{RawResponse: resp.Response, LongWrapper: val}, nil
}

// getLongHandleError handles the GetLong error response.
func (client *PrimitiveClient) getLongHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// GetString - Get complex types with string properties
func (client *PrimitiveClient) GetString(ctx context.Context, options *PrimitiveGetStringOptions) (StringWrapperResponse, error) {
	req, err := client.getStringCreateRequest(ctx, options)
	if err != nil {
		return StringWrapperResponse{}, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return StringWrapperResponse{}, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return StringWrapperResponse{}, client.getStringHandleError(resp)
	}
	return client.getStringHandleResponse(resp)
}

// getStringCreateRequest creates the GetString request.
func (client *PrimitiveClient) getStringCreateRequest(ctx context.Context, options *PrimitiveGetStringOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/string"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// getStringHandleResponse handles the GetString response.
func (client *PrimitiveClient) getStringHandleResponse(resp *azcore.Response) (StringWrapperResponse, error) {
	var val *StringWrapper
	if err := resp.UnmarshalAsJSON(&val); err != nil {
		return StringWrapperResponse{}, err
	}
	return StringWrapperResponse{RawResponse: resp.Response, StringWrapper: val}, nil
}

// getStringHandleError handles the GetString error response.
func (client *PrimitiveClient) getStringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutBool - Put complex types with bool properties
func (client *PrimitiveClient) PutBool(ctx context.Context, complexBody BooleanWrapper, options *PrimitivePutBoolOptions) (*http.Response, error) {
	req, err := client.putBoolCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putBoolHandleError(resp)
	}
	return resp.Response, nil
}

// putBoolCreateRequest creates the PutBool request.
func (client *PrimitiveClient) putBoolCreateRequest(ctx context.Context, complexBody BooleanWrapper, options *PrimitivePutBoolOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/bool"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putBoolHandleError handles the PutBool error response.
func (client *PrimitiveClient) putBoolHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutByte - Put complex types with byte properties
func (client *PrimitiveClient) PutByte(ctx context.Context, complexBody ByteWrapper, options *PrimitivePutByteOptions) (*http.Response, error) {
	req, err := client.putByteCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putByteHandleError(resp)
	}
	return resp.Response, nil
}

// putByteCreateRequest creates the PutByte request.
func (client *PrimitiveClient) putByteCreateRequest(ctx context.Context, complexBody ByteWrapper, options *PrimitivePutByteOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/byte"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putByteHandleError handles the PutByte error response.
func (client *PrimitiveClient) putByteHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDate - Put complex types with date properties
func (client *PrimitiveClient) PutDate(ctx context.Context, complexBody DateWrapper, options *PrimitivePutDateOptions) (*http.Response, error) {
	req, err := client.putDateCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDateHandleError(resp)
	}
	return resp.Response, nil
}

// putDateCreateRequest creates the PutDate request.
func (client *PrimitiveClient) putDateCreateRequest(ctx context.Context, complexBody DateWrapper, options *PrimitivePutDateOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/date"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putDateHandleError handles the PutDate error response.
func (client *PrimitiveClient) putDateHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDateTime - Put complex types with datetime properties
func (client *PrimitiveClient) PutDateTime(ctx context.Context, complexBody DatetimeWrapper, options *PrimitivePutDateTimeOptions) (*http.Response, error) {
	req, err := client.putDateTimeCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDateTimeHandleError(resp)
	}
	return resp.Response, nil
}

// putDateTimeCreateRequest creates the PutDateTime request.
func (client *PrimitiveClient) putDateTimeCreateRequest(ctx context.Context, complexBody DatetimeWrapper, options *PrimitivePutDateTimeOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetime"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putDateTimeHandleError handles the PutDateTime error response.
func (client *PrimitiveClient) putDateTimeHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDateTimeRFC1123 - Put complex types with datetimeRfc1123 properties
func (client *PrimitiveClient) PutDateTimeRFC1123(ctx context.Context, complexBody Datetimerfc1123Wrapper, options *PrimitivePutDateTimeRFC1123Options) (*http.Response, error) {
	req, err := client.putDateTimeRfc1123CreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDateTimeRfc1123HandleError(resp)
	}
	return resp.Response, nil
}

// putDateTimeRfc1123CreateRequest creates the PutDateTimeRFC1123 request.
func (client *PrimitiveClient) putDateTimeRfc1123CreateRequest(ctx context.Context, complexBody Datetimerfc1123Wrapper, options *PrimitivePutDateTimeRFC1123Options) (*azcore.Request, error) {
	urlPath := "/complex/primitive/datetimerfc1123"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putDateTimeRfc1123HandleError handles the PutDateTimeRFC1123 error response.
func (client *PrimitiveClient) putDateTimeRfc1123HandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDouble - Put complex types with double properties
func (client *PrimitiveClient) PutDouble(ctx context.Context, complexBody DoubleWrapper, options *PrimitivePutDoubleOptions) (*http.Response, error) {
	req, err := client.putDoubleCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDoubleHandleError(resp)
	}
	return resp.Response, nil
}

// putDoubleCreateRequest creates the PutDouble request.
func (client *PrimitiveClient) putDoubleCreateRequest(ctx context.Context, complexBody DoubleWrapper, options *PrimitivePutDoubleOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/double"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putDoubleHandleError handles the PutDouble error response.
func (client *PrimitiveClient) putDoubleHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutDuration - Put complex types with duration properties
func (client *PrimitiveClient) PutDuration(ctx context.Context, complexBody DurationWrapper, options *PrimitivePutDurationOptions) (*http.Response, error) {
	req, err := client.putDurationCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putDurationHandleError(resp)
	}
	return resp.Response, nil
}

// putDurationCreateRequest creates the PutDuration request.
func (client *PrimitiveClient) putDurationCreateRequest(ctx context.Context, complexBody DurationWrapper, options *PrimitivePutDurationOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/duration"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putDurationHandleError handles the PutDuration error response.
func (client *PrimitiveClient) putDurationHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutFloat - Put complex types with float properties
func (client *PrimitiveClient) PutFloat(ctx context.Context, complexBody FloatWrapper, options *PrimitivePutFloatOptions) (*http.Response, error) {
	req, err := client.putFloatCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putFloatHandleError(resp)
	}
	return resp.Response, nil
}

// putFloatCreateRequest creates the PutFloat request.
func (client *PrimitiveClient) putFloatCreateRequest(ctx context.Context, complexBody FloatWrapper, options *PrimitivePutFloatOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/float"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putFloatHandleError handles the PutFloat error response.
func (client *PrimitiveClient) putFloatHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutInt - Put complex types with integer properties
func (client *PrimitiveClient) PutInt(ctx context.Context, complexBody IntWrapper, options *PrimitivePutIntOptions) (*http.Response, error) {
	req, err := client.putIntCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putIntHandleError(resp)
	}
	return resp.Response, nil
}

// putIntCreateRequest creates the PutInt request.
func (client *PrimitiveClient) putIntCreateRequest(ctx context.Context, complexBody IntWrapper, options *PrimitivePutIntOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/integer"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putIntHandleError handles the PutInt error response.
func (client *PrimitiveClient) putIntHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutLong - Put complex types with long properties
func (client *PrimitiveClient) PutLong(ctx context.Context, complexBody LongWrapper, options *PrimitivePutLongOptions) (*http.Response, error) {
	req, err := client.putLongCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putLongHandleError(resp)
	}
	return resp.Response, nil
}

// putLongCreateRequest creates the PutLong request.
func (client *PrimitiveClient) putLongCreateRequest(ctx context.Context, complexBody LongWrapper, options *PrimitivePutLongOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/long"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putLongHandleError handles the PutLong error response.
func (client *PrimitiveClient) putLongHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}

// PutString - Put complex types with string properties
func (client *PrimitiveClient) PutString(ctx context.Context, complexBody StringWrapper, options *PrimitivePutStringOptions) (*http.Response, error) {
	req, err := client.putStringCreateRequest(ctx, complexBody, options)
	if err != nil {
		return nil, err
	}
	resp, err := client.Pipeline().Do(req)
	if err != nil {
		return nil, err
	}
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.putStringHandleError(resp)
	}
	return resp.Response, nil
}

// putStringCreateRequest creates the PutString request.
func (client *PrimitiveClient) putStringCreateRequest(ctx context.Context, complexBody StringWrapper, options *PrimitivePutStringOptions) (*azcore.Request, error) {
	urlPath := "/complex/primitive/string"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.con.Endpoint(), urlPath))
	if err != nil {
		return nil, err
	}
	req.Telemetry(telemetryInfo)
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(complexBody)
}

// putStringHandleError handles the PutString error response.
func (client *PrimitiveClient) putStringHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return azcore.NewResponseError(&err, resp.Response)
}
