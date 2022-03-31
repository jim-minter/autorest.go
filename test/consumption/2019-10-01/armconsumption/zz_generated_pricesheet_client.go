//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package armconsumption

import (
	"context"
	"errors"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/arm"
	armruntime "github.com/Azure/azure-sdk-for-go/sdk/azcore/arm/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// PriceSheetClient contains the methods for the PriceSheet group.
// Don't use this type directly, use NewPriceSheetClient() instead.
type PriceSheetClient struct {
	host           string
	subscriptionID string
	pl             runtime.Pipeline
}

// NewPriceSheetClient creates a new instance of PriceSheetClient with the specified values.
// subscriptionID - Azure Subscription ID.
// credential - used to authorize requests. Usually a credential from azidentity.
// options - pass nil to accept the default values.
func NewPriceSheetClient(subscriptionID string, credential azcore.TokenCredential, options *arm.ClientOptions) *PriceSheetClient {
	if options == nil {
		options = &arm.ClientOptions{}
	}
	ep := options.Endpoint
	if len(ep) == 0 {
		ep = arm.AzurePublicCloud
	}
	client := &PriceSheetClient{
		subscriptionID: subscriptionID,
		host:           string(ep),
		pl:             armruntime.NewPipeline(moduleName, moduleVersion, credential, runtime.PipelineOptions{}, options),
	}
	return client
}

// Get - Gets the price sheet for a scope by subscriptionId. Price sheet is available via this API only for May 1, 2014 or
// later.
// If the operation fails it returns an *azcore.ResponseError type.
// options - PriceSheetClientGetOptions contains the optional parameters for the PriceSheetClient.Get method.
func (client *PriceSheetClient) Get(ctx context.Context, options *PriceSheetClientGetOptions) (PriceSheetClientGetResponse, error) {
	req, err := client.getCreateRequest(ctx, options)
	if err != nil {
		return PriceSheetClientGetResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PriceSheetClientGetResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PriceSheetClientGetResponse{}, runtime.NewResponseError(resp)
	}
	return client.getHandleResponse(resp)
}

// getCreateRequest creates the Get request.
func (client *PriceSheetClient) getCreateRequest(ctx context.Context, options *PriceSheetClientGetOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Consumption/pricesheets/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getHandleResponse handles the Get response.
func (client *PriceSheetClient) getHandleResponse(resp *http.Response) (PriceSheetClientGetResponse, error) {
	result := PriceSheetClientGetResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PriceSheetResult); err != nil {
		return PriceSheetClientGetResponse{}, err
	}
	return result, nil
}

// GetByBillingPeriod - Get the price sheet for a scope by subscriptionId and billing period. Price sheet is available via
// this API only for May 1, 2014 or later.
// If the operation fails it returns an *azcore.ResponseError type.
// billingPeriodName - Billing Period Name.
// options - PriceSheetClientGetByBillingPeriodOptions contains the optional parameters for the PriceSheetClient.GetByBillingPeriod
// method.
func (client *PriceSheetClient) GetByBillingPeriod(ctx context.Context, billingPeriodName string, options *PriceSheetClientGetByBillingPeriodOptions) (PriceSheetClientGetByBillingPeriodResponse, error) {
	req, err := client.getByBillingPeriodCreateRequest(ctx, billingPeriodName, options)
	if err != nil {
		return PriceSheetClientGetByBillingPeriodResponse{}, err
	}
	resp, err := client.pl.Do(req)
	if err != nil {
		return PriceSheetClientGetByBillingPeriodResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return PriceSheetClientGetByBillingPeriodResponse{}, runtime.NewResponseError(resp)
	}
	return client.getByBillingPeriodHandleResponse(resp)
}

// getByBillingPeriodCreateRequest creates the GetByBillingPeriod request.
func (client *PriceSheetClient) getByBillingPeriodCreateRequest(ctx context.Context, billingPeriodName string, options *PriceSheetClientGetByBillingPeriodOptions) (*policy.Request, error) {
	urlPath := "/subscriptions/{subscriptionId}/providers/Microsoft.Billing/billingPeriods/{billingPeriodName}/providers/Microsoft.Consumption/pricesheets/default"
	if client.subscriptionID == "" {
		return nil, errors.New("parameter client.subscriptionID cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{subscriptionId}", url.PathEscape(client.subscriptionID))
	if billingPeriodName == "" {
		return nil, errors.New("parameter billingPeriodName cannot be empty")
	}
	urlPath = strings.ReplaceAll(urlPath, "{billingPeriodName}", url.PathEscape(billingPeriodName))
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.host, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if options != nil && options.Expand != nil {
		reqQP.Set("$expand", *options.Expand)
	}
	if options != nil && options.Skiptoken != nil {
		reqQP.Set("$skiptoken", *options.Skiptoken)
	}
	if options != nil && options.Top != nil {
		reqQP.Set("$top", strconv.FormatInt(int64(*options.Top), 10))
	}
	reqQP.Set("api-version", "2019-10-01")
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header.Set("Accept", "application/json")
	return req, nil
}

// getByBillingPeriodHandleResponse handles the GetByBillingPeriod response.
func (client *PriceSheetClient) getByBillingPeriodHandleResponse(resp *http.Response) (PriceSheetClientGetByBillingPeriodResponse, error) {
	result := PriceSheetClientGetByBillingPeriodResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PriceSheetResult); err != nil {
		return PriceSheetClientGetByBillingPeriodResponse{}, err
	}
	return result, nil
}
