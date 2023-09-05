//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azalias

import (
	"context"
	"fmt"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"strconv"
	"strings"
)

// Client contains the methods for the Alias group.
// Don't use this type directly, use a constructor function instead.
type Client struct {
	internal            *azcore.Client
	endpoint            string
	clientGroup         ClientGroup
	clientOptionalGroup *ClientOptionalGroup
}

// Create - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to create an alias. You can also assign the alias during the create request. An alias can reference
// an ID generated by a creator service, but cannot reference another alias
// ID.
// SUBMIT CREATE REQUEST To create your alias, you will use a POST request. If you would like to assign the alias during the
// creation, you will pass the resourceId query parameter.
// CREATE ALIAS RESPONSE The Create API returns a HTTP 201 Created response with the alias resource in the body.
// A sample response from creating an alias:
// { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0
//   - options - ClientCreateOptions contains the optional parameters for the Client.Create method.
func (client *Client) Create(ctx context.Context, options *ClientCreateOptions) (ClientCreateResponse, error) {
	var err error
	const operationName = "Client.Create"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.createCreateRequest(ctx, options)
	if err != nil {
		return ClientCreateResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientCreateResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusCreated) {
		err = runtime.NewResponseError(httpResp)
		return ClientCreateResponse{}, err
	}
	resp, err := client.createHandleResponse(httpResp)
	return resp, err
}

// createCreateRequest creates the Create request.
func (client *Client) createCreateRequest(ctx context.Context, options *ClientCreateOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("client-version", client.clientGroup.ClientVersion)
	creatorIDDefault := int32(123)
	if options != nil && options.CreatorID != nil {
		creatorIDDefault = *options.CreatorID
	}
	reqQP.Set("creator-id", strconv.FormatInt(int64(creatorIDDefault), 10))
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", fmt.Sprintf("%d", qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["client-index"] = []string{strconv.FormatInt(int64(client.clientGroup.ClientIndex), 10)}
	assignedIDDefault := float32(8989)
	if options != nil && options.AssignedID != nil {
		assignedIDDefault = *options.AssignedID
	}
	req.Raw().Header["assigned-id"] = []string{strconv.FormatFloat(float64(assignedIDDefault), 'f', -1, 32)}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// createHandleResponse handles the Create response.
func (client *Client) createHandleResponse(resp *http.Response) (ClientCreateResponse, error) {
	result := ClientCreateResponse{}
	if val := resp.Header.Get("Access-Control-Expose-Headers"); val != "" {
		result.AccessControlExposeHeaders = &val
	}
	if err := runtime.UnmarshalAsJSON(resp, &result.AliasesCreateResponse); err != nil {
		return ClientCreateResponse{}, err
	}
	return result, nil
}

// GetScript - Retrieve the configuration script identified by configuration name.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0
//   - options - ClientGetScriptOptions contains the optional parameters for the Client.GetScript method.
func (client *Client) GetScript(ctx context.Context, props GeoJSONObjectNamedCollection, options *ClientGetScriptOptions) (ClientGetScriptResponse, error) {
	var err error
	const operationName = "Client.GetScript"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.getScriptCreateRequest(ctx, props, options)
	if err != nil {
		return ClientGetScriptResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientGetScriptResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientGetScriptResponse{}, err
	}
	resp, err := client.getScriptHandleResponse(httpResp)
	return resp, err
}

// getScriptCreateRequest creates the GetScript request.
func (client *Client) getScriptCreateRequest(ctx context.Context, props GeoJSONObjectNamedCollection, options *ClientGetScriptOptions) (*policy.Request, error) {
	urlPath := "/scripts"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"text/powershell"}
	if err := runtime.MarshalAsJSON(req, props); err != nil {
		return nil, err
	}
	return req, nil
}

// getScriptHandleResponse handles the GetScript response.
func (client *Client) getScriptHandleResponse(resp *http.Response) (ClientGetScriptResponse, error) {
	result := ClientGetScriptResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return ClientGetScriptResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// NewListPager - Applies to: see pricing tiers [https://aka.ms/AzureMapsPricingTier].
// Creator makes it possible to develop applications based on your private indoor map data using Azure Maps API and SDK. This
// [https://docs.microsoft.com/azure/azure-maps/creator-indoor-maps] article
// introduces concepts and tools that apply to Azure Maps Creator.
// This API allows the caller to fetch a list of all previously successfully created aliases.
// SUBMIT LIST REQUEST To list all your aliases, you will issue a GET request with no additional parameters.
// LIST DATA RESPONSE The List API returns the complete list of all aliases in json format. The response contains the following
// details for each alias resource:
// > createdTimestamp - The timestamp that the alias was created. Format yyyy-MM-ddTHH:mm:ss.sssZ aliasId - The id for the
// alias. creatorDataItemId - The id for the creator data item that this alias
// references (could be null if the alias has not been assigned). lastUpdatedTimestamp - The last time the alias was assigned
// to a resource. Format yyyy-MM-ddTHH:mm:ss.sssZ
// A sample response returning 2 alias resources:
// { "aliases": [ { "createdTimestamp": "2020-02-13T21:19:11.123Z", "aliasId": "a8a4b8bb-ecf4-fb27-a618-f41721552766", "creatorDataItemId":
// "e89aebb9-70a3-8fe1-32bb-1fbd0c725f14", "lastUpdatedTimestamp":
// "2020-02-13T21:19:22.123Z" }, { "createdTimestamp": "2020-02-18T19:53:33.123Z", "aliasId": "1856dbfc-7a66-ee5a-bf8d-51dbfe1906f6",
// "creatorDataItemId": null, "lastUpdatedTimestamp":
// "2020-02-18T19:53:33.123Z" } ] }
//
// Generated from API version 2.0
//   - options - ClientListOptions contains the optional parameters for the Client.NewListPager method.
func (client *Client) NewListPager(options *ClientListOptions) *runtime.Pager[ClientListResponse] {
	return runtime.NewPager(runtime.PagingHandler[ClientListResponse]{
		More: func(page ClientListResponse) bool {
			return page.NextLink != nil && len(*page.NextLink) > 0
		},
		Fetcher: func(ctx context.Context, page *ClientListResponse) (ClientListResponse, error) {
			ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, "Client.NewListPager")
			var req *policy.Request
			var err error
			if page == nil {
				req, err = client.listCreateRequest(ctx, options)
			} else {
				req, err = runtime.NewRequest(ctx, http.MethodGet, *page.NextLink)
			}
			if err != nil {
				return ClientListResponse{}, err
			}
			resp, err := client.internal.Pipeline().Do(req)
			if err != nil {
				return ClientListResponse{}, err
			}
			if !runtime.HasStatusCode(resp, http.StatusOK) {
				return ClientListResponse{}, runtime.NewResponseError(resp)
			}
			return client.listHandleResponse(resp)
		},
		Tracer: client.internal.Tracer(),
	})
}

// listCreateRequest creates the List request.
func (client *Client) listCreateRequest(ctx context.Context, options *ClientListOptions) (*policy.Request, error) {
	urlPath := "/aliases"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	reqQP.Set("client-version", client.clientGroup.ClientVersion)
	if options != nil && options.GroupBy != nil {
		for _, qv := range options.GroupBy {
			reqQP.Add("groupBy", string(qv))
		}
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// listHandleResponse handles the List response.
func (client *Client) listHandleResponse(resp *http.Response) (ClientListResponse, error) {
	result := ClientListResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.ListResponse); err != nil {
		return ClientListResponse{}, err
	}
	return result, nil
}

// PolicyAssignment -
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0
//   - options - ClientPolicyAssignmentOptions contains the optional parameters for the Client.PolicyAssignment method.
func (client *Client) PolicyAssignment(ctx context.Context, things []Things, polymorphicParam GeoJSONObjectClassification, options *ClientPolicyAssignmentOptions) (ClientPolicyAssignmentResponse, error) {
	var err error
	const operationName = "Client.PolicyAssignment"
	ctx = context.WithValue(ctx, runtime.CtxAPINameKey{}, operationName)
	ctx, endSpan := runtime.StartSpan(ctx, operationName, client.internal.Tracer(), nil)
	defer func() { endSpan(err) }()
	req, err := client.policyAssignmentCreateRequest(ctx, things, polymorphicParam, options)
	if err != nil {
		return ClientPolicyAssignmentResponse{}, err
	}
	httpResp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return ClientPolicyAssignmentResponse{}, err
	}
	if !runtime.HasStatusCode(httpResp, http.StatusOK) {
		err = runtime.NewResponseError(httpResp)
		return ClientPolicyAssignmentResponse{}, err
	}
	resp, err := client.policyAssignmentHandleResponse(httpResp)
	return resp, err
}

// policyAssignmentCreateRequest creates the PolicyAssignment request.
func (client *Client) policyAssignmentCreateRequest(ctx context.Context, things []Things, polymorphicParam GeoJSONObjectClassification, options *ClientPolicyAssignmentOptions) (*policy.Request, error) {
	urlPath := "/policy"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(client.endpoint, urlPath))
	if err != nil {
		return nil, err
	}
	reqQP := req.Raw().URL.Query()
	if client.clientOptionalGroup != nil && client.clientOptionalGroup.OptionalVersion != nil {
		reqQP.Set("optional-version", *client.clientOptionalGroup.OptionalVersion)
	}
	reqQP.Set("things", strings.Join(strings.Fields(strings.Trim(fmt.Sprint(things), "[]")), ","))
	if options != nil && options.Interval != nil {
		reqQP.Set("interval", *options.Interval)
	}
	if options != nil && options.Unique != nil {
		reqQP.Set("unique", *options.Unique)
	}
	req.Raw().URL.RawQuery = reqQP.Encode()
	if client.clientOptionalGroup != nil && client.clientOptionalGroup.OptionalIndex != nil {
		req.Raw().Header["optional-index"] = []string{strconv.FormatInt(int64(*client.clientOptionalGroup.OptionalIndex), 10)}
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, polymorphicParam); err != nil {
		return nil, err
	}
	return req, nil
}

// policyAssignmentHandleResponse handles the PolicyAssignment response.
func (client *Client) policyAssignmentHandleResponse(resp *http.Response) (ClientPolicyAssignmentResponse, error) {
	result := ClientPolicyAssignmentResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.PolicyAssignmentProperties); err != nil {
		return ClientPolicyAssignmentResponse{}, err
	}
	return result, nil
}
