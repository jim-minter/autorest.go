// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azblob

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"strconv"
	"time"
)

// ServiceOperations contains the methods for the Service group.
type ServiceOperations interface {
	// GetAccountInfo - Returns the sku name and account kind
	GetAccountInfo(ctx context.Context) (*ServiceGetAccountInfoResponse, error)
	// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
	GetProperties(ctx context.Context, serviceGetPropertiesOptions *ServiceGetPropertiesOptions) (*StorageServicePropertiesResponse, error)
	// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary location endpoint when read-access geo-redundant replication is enabled for the storage account.
	GetStatistics(ctx context.Context, serviceGetStatisticsOptions *ServiceGetStatisticsOptions) (*StorageServiceStatsResponse, error)
	// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using bearer token authentication.
	GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, serviceGetUserDelegationKeyOptions *ServiceGetUserDelegationKeyOptions) (*UserDelegationKeyResponse, error)
	// ListContainersSegment - The List Containers Segment operation returns a list of the containers under the specified account
	ListContainersSegment(serviceListContainersSegmentOptions *ServiceListContainersSegmentOptions) ListContainersSegmentResponsePager
	// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules
	SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, serviceSetPropertiesOptions *ServiceSetPropertiesOptions) (*ServiceSetPropertiesResponse, error)
	// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
	SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, serviceSubmitBatchOptions *ServiceSubmitBatchOptions) (*ServiceSubmitBatchResponse, error)
}

// serviceClient implements the ServiceOperations interface.
// Don't use this type directly, use newServiceClient() instead.
type serviceClient struct {
	*client
}

// newServiceClient creates a new instance of serviceClient with the specified values.
func newServiceClient(c *client) ServiceOperations {
	return &serviceClient{client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *serviceClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetAccountInfo - Returns the sku name and account kind
func (client *serviceClient) GetAccountInfo(ctx context.Context) (*ServiceGetAccountInfoResponse, error) {
	req, err := client.GetAccountInfoCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetAccountInfoHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetAccountInfoCreateRequest creates the GetAccountInfo request.
func (client *serviceClient) GetAccountInfoCreateRequest(ctx context.Context) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.u)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "account")
	query.Set("comp", "properties")
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// GetAccountInfoHandleResponse handles the GetAccountInfo response.
func (client *serviceClient) GetAccountInfoHandleResponse(resp *azcore.Response) (*ServiceGetAccountInfoResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetAccountInfoHandleError(resp)
	}
	result := ServiceGetAccountInfoResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	if val := resp.Header.Get("x-ms-sku-name"); val != "" {
		result.SKUName = (*SKUName)(&val)
	}
	if val := resp.Header.Get("x-ms-account-kind"); val != "" {
		result.AccountKind = (*AccountKind)(&val)
	}
	return &result, nil
}

// GetAccountInfoHandleError handles the GetAccountInfo error response.
func (client *serviceClient) GetAccountInfoHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// GetProperties - gets the properties of a storage account's Blob service, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules.
func (client *serviceClient) GetProperties(ctx context.Context, serviceGetPropertiesOptions *ServiceGetPropertiesOptions) (*StorageServicePropertiesResponse, error) {
	req, err := client.GetPropertiesCreateRequest(ctx, serviceGetPropertiesOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetPropertiesCreateRequest creates the GetProperties request.
func (client *serviceClient) GetPropertiesCreateRequest(ctx context.Context, serviceGetPropertiesOptions *ServiceGetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.u)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if serviceGetPropertiesOptions != nil && serviceGetPropertiesOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceGetPropertiesOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceGetPropertiesOptions != nil && serviceGetPropertiesOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceGetPropertiesOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// GetPropertiesHandleResponse handles the GetProperties response.
func (client *serviceClient) GetPropertiesHandleResponse(resp *azcore.Response) (*StorageServicePropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetPropertiesHandleError(resp)
	}
	result := StorageServicePropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, resp.UnmarshalAsXML(&result.StorageServiceProperties)
}

// GetPropertiesHandleError handles the GetProperties error response.
func (client *serviceClient) GetPropertiesHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// GetStatistics - Retrieves statistics related to replication for the Blob service. It is only available on the secondary location endpoint when read-access geo-redundant replication is enabled for the storage account.
func (client *serviceClient) GetStatistics(ctx context.Context, serviceGetStatisticsOptions *ServiceGetStatisticsOptions) (*StorageServiceStatsResponse, error) {
	req, err := client.GetStatisticsCreateRequest(ctx, serviceGetStatisticsOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetStatisticsHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetStatisticsCreateRequest creates the GetStatistics request.
func (client *serviceClient) GetStatisticsCreateRequest(ctx context.Context, serviceGetStatisticsOptions *ServiceGetStatisticsOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.u)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "stats")
	if serviceGetStatisticsOptions != nil && serviceGetStatisticsOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceGetStatisticsOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceGetStatisticsOptions != nil && serviceGetStatisticsOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceGetStatisticsOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// GetStatisticsHandleResponse handles the GetStatistics response.
func (client *serviceClient) GetStatisticsHandleResponse(resp *azcore.Response) (*StorageServiceStatsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetStatisticsHandleError(resp)
	}
	result := StorageServiceStatsResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, resp.UnmarshalAsXML(&result.StorageServiceStats)
}

// GetStatisticsHandleError handles the GetStatistics error response.
func (client *serviceClient) GetStatisticsHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// GetUserDelegationKey - Retrieves a user delegation key for the Blob service. This is only a valid operation when using bearer token authentication.
func (client *serviceClient) GetUserDelegationKey(ctx context.Context, keyInfo KeyInfo, serviceGetUserDelegationKeyOptions *ServiceGetUserDelegationKeyOptions) (*UserDelegationKeyResponse, error) {
	req, err := client.GetUserDelegationKeyCreateRequest(ctx, keyInfo, serviceGetUserDelegationKeyOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetUserDelegationKeyHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetUserDelegationKeyCreateRequest creates the GetUserDelegationKey request.
func (client *serviceClient) GetUserDelegationKeyCreateRequest(ctx context.Context, keyInfo KeyInfo, serviceGetUserDelegationKeyOptions *ServiceGetUserDelegationKeyOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.u)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "userdelegationkey")
	if serviceGetUserDelegationKeyOptions != nil && serviceGetUserDelegationKeyOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceGetUserDelegationKeyOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceGetUserDelegationKeyOptions != nil && serviceGetUserDelegationKeyOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceGetUserDelegationKeyOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(keyInfo)
}

// GetUserDelegationKeyHandleResponse handles the GetUserDelegationKey response.
func (client *serviceClient) GetUserDelegationKeyHandleResponse(resp *azcore.Response) (*UserDelegationKeyResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetUserDelegationKeyHandleError(resp)
	}
	result := UserDelegationKeyResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	if val := resp.Header.Get("Date"); val != "" {
		date, err := time.Parse(time.RFC1123, val)
		if err != nil {
			return nil, err
		}
		result.Date = &date
	}
	return &result, resp.UnmarshalAsXML(&result.UserDelegationKey)
}

// GetUserDelegationKeyHandleError handles the GetUserDelegationKey error response.
func (client *serviceClient) GetUserDelegationKeyHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// ListContainersSegment - The List Containers Segment operation returns a list of the containers under the specified account
func (client *serviceClient) ListContainersSegment(serviceListContainersSegmentOptions *ServiceListContainersSegmentOptions) ListContainersSegmentResponsePager {
	return &listContainersSegmentResponsePager{
		pipeline: client.p,
		requester: func(ctx context.Context) (*azcore.Request, error) {
			return client.ListContainersSegmentCreateRequest(ctx, serviceListContainersSegmentOptions)
		},
		responder: client.ListContainersSegmentHandleResponse,
		advancer: func(ctx context.Context, resp *ListContainersSegmentResponseResponse) (*azcore.Request, error) {
			return azcore.NewRequest(ctx, http.MethodGet, *resp.EnumerationResults.NextMarker)
		},
	}
}

// ListContainersSegmentCreateRequest creates the ListContainersSegment request.
func (client *serviceClient) ListContainersSegmentCreateRequest(ctx context.Context, serviceListContainersSegmentOptions *ServiceListContainersSegmentOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodGet, client.u)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("comp", "list")
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Prefix != nil {
		query.Set("prefix", *serviceListContainersSegmentOptions.Prefix)
	}
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Marker != nil {
		query.Set("marker", *serviceListContainersSegmentOptions.Marker)
	}
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Maxresults != nil {
		query.Set("maxresults", strconv.FormatInt(int64(*serviceListContainersSegmentOptions.Maxresults), 10))
	}
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceListContainersSegmentOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceListContainersSegmentOptions != nil && serviceListContainersSegmentOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceListContainersSegmentOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, nil
}

// ListContainersSegmentHandleResponse handles the ListContainersSegment response.
func (client *serviceClient) ListContainersSegmentHandleResponse(resp *azcore.Response) (*ListContainersSegmentResponseResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.ListContainersSegmentHandleError(resp)
	}
	result := ListContainersSegmentResponseResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, resp.UnmarshalAsXML(&result.EnumerationResults)
}

// ListContainersSegmentHandleError handles the ListContainersSegment error response.
func (client *serviceClient) ListContainersSegmentHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// SetProperties - Sets properties for a storage account's Blob service endpoint, including properties for Storage Analytics and CORS (Cross-Origin Resource Sharing) rules
func (client *serviceClient) SetProperties(ctx context.Context, storageServiceProperties StorageServiceProperties, serviceSetPropertiesOptions *ServiceSetPropertiesOptions) (*ServiceSetPropertiesResponse, error) {
	req, err := client.SetPropertiesCreateRequest(ctx, storageServiceProperties, serviceSetPropertiesOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.SetPropertiesHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetPropertiesCreateRequest creates the SetProperties request.
func (client *serviceClient) SetPropertiesCreateRequest(ctx context.Context, storageServiceProperties StorageServiceProperties, serviceSetPropertiesOptions *ServiceSetPropertiesOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPut, client.u)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("restype", "service")
	query.Set("comp", "properties")
	if serviceSetPropertiesOptions != nil && serviceSetPropertiesOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceSetPropertiesOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceSetPropertiesOptions != nil && serviceSetPropertiesOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceSetPropertiesOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(storageServiceProperties)
}

// SetPropertiesHandleResponse handles the SetProperties response.
func (client *serviceClient) SetPropertiesHandleResponse(resp *azcore.Response) (*ServiceSetPropertiesResponse, error) {
	if !resp.HasStatusCode(http.StatusAccepted) {
		return nil, client.SetPropertiesHandleError(resp)
	}
	result := ServiceSetPropertiesResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("x-ms-client-request-id"); val != "" {
		result.ClientRequestID = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// SetPropertiesHandleError handles the SetProperties error response.
func (client *serviceClient) SetPropertiesHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}

// SubmitBatch - The Batch operation allows multiple API calls to be embedded into a single HTTP request.
func (client *serviceClient) SubmitBatch(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, serviceSubmitBatchOptions *ServiceSubmitBatchOptions) (*ServiceSubmitBatchResponse, error) {
	req, err := client.SubmitBatchCreateRequest(ctx, contentLength, multipartContentType, body, serviceSubmitBatchOptions)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.SubmitBatchHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SubmitBatchCreateRequest creates the SubmitBatch request.
func (client *serviceClient) SubmitBatchCreateRequest(ctx context.Context, contentLength int64, multipartContentType string, body azcore.ReadSeekCloser, serviceSubmitBatchOptions *ServiceSubmitBatchOptions) (*azcore.Request, error) {
	req, err := azcore.NewRequest(ctx, http.MethodPost, client.u)
	if err != nil {
		return nil, err
	}
	query := req.URL.Query()
	query.Set("comp", "batch")
	if serviceSubmitBatchOptions != nil && serviceSubmitBatchOptions.Timeout != nil {
		query.Set("timeout", strconv.FormatInt(int64(*serviceSubmitBatchOptions.Timeout), 10))
	}
	req.URL.RawQuery = query.Encode()
	req.SkipBodyDownload()
	req.Header.Set("Content-Length", strconv.FormatInt(contentLength, 10))
	req.Header.Set("Content-Type", multipartContentType)
	req.Header.Set("x-ms-version", "2019-07-07")
	if serviceSubmitBatchOptions != nil && serviceSubmitBatchOptions.RequestId != nil {
		req.Header.Set("x-ms-client-request-id", *serviceSubmitBatchOptions.RequestId)
	}
	req.Header.Set("Accept", "application/xml")
	return req, req.MarshalAsXML(body)
}

// SubmitBatchHandleResponse handles the SubmitBatch response.
func (client *serviceClient) SubmitBatchHandleResponse(resp *azcore.Response) (*ServiceSubmitBatchResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.SubmitBatchHandleError(resp)
	}
	result := ServiceSubmitBatchResponse{RawResponse: resp.Response}
	if val := resp.Header.Get("Content-Type"); val != "" {
		result.ContentType = &val
	}
	if val := resp.Header.Get("x-ms-request-id"); val != "" {
		result.RequestID = &val
	}
	if val := resp.Header.Get("x-ms-version"); val != "" {
		result.Version = &val
	}
	return &result, nil
}

// SubmitBatchHandleError handles the SubmitBatch error response.
func (client *serviceClient) SubmitBatchHandleError(resp *azcore.Response) error {
	var err StorageError
	if err := resp.UnmarshalAsXML(&err); err != nil {
		return err
	}
	return err
}
