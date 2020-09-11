// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
)

// EnumOperations contains the methods for the Enum group.
type EnumOperations interface {
	// GetNotExpandable - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
	GetNotExpandable(ctx context.Context) (*ColorsResponse, error)
	// GetReferenced - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
	GetReferenced(ctx context.Context) (*ColorsResponse, error)
	// GetReferencedConstant - Get value 'green-color' from the constant.
	GetReferencedConstant(ctx context.Context) (*RefColorConstantResponse, error)
	// PutNotExpandable - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
	PutNotExpandable(ctx context.Context, stringBody Colors) (*http.Response, error)
	// PutReferenced - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
	PutReferenced(ctx context.Context, enumStringBody Colors) (*http.Response, error)
	// PutReferencedConstant - Sends value 'green-color' from a constant
	PutReferencedConstant(ctx context.Context, enumStringBody RefColorConstant) (*http.Response, error)
}

// EnumClient implements the EnumOperations interface.
// Don't use this type directly, use NewEnumClient() instead.
type EnumClient struct {
	*Client
}

// NewEnumClient creates a new instance of EnumClient with the specified values.
func NewEnumClient(c *Client) EnumOperations {
	return &EnumClient{Client: c}
}

// Do invokes the Do() method on the pipeline associated with this client.
func (client *EnumClient) Do(req *azcore.Request) (*azcore.Response, error) {
	return client.p.Do(req)
}

// GetNotExpandable - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
func (client *EnumClient) GetNotExpandable(ctx context.Context) (*ColorsResponse, error) {
	req, err := client.GetNotExpandableCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetNotExpandableHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetNotExpandableCreateRequest creates the GetNotExpandable request.
func (client *EnumClient) GetNotExpandableCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetNotExpandableHandleResponse handles the GetNotExpandable response.
func (client *EnumClient) GetNotExpandableHandleResponse(resp *azcore.Response) (*ColorsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetNotExpandableHandleError(resp)
	}
	result := ColorsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetNotExpandableHandleError handles the GetNotExpandable error response.
func (client *EnumClient) GetNotExpandableHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetReferenced - Get enum value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'.
func (client *EnumClient) GetReferenced(ctx context.Context) (*ColorsResponse, error) {
	req, err := client.GetReferencedCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetReferencedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetReferencedCreateRequest creates the GetReferenced request.
func (client *EnumClient) GetReferencedCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetReferencedHandleResponse handles the GetReferenced response.
func (client *EnumClient) GetReferencedHandleResponse(resp *azcore.Response) (*ColorsResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetReferencedHandleError(resp)
	}
	result := ColorsResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.Value)
}

// GetReferencedHandleError handles the GetReferenced error response.
func (client *EnumClient) GetReferencedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// GetReferencedConstant - Get value 'green-color' from the constant.
func (client *EnumClient) GetReferencedConstant(ctx context.Context) (*RefColorConstantResponse, error) {
	req, err := client.GetReferencedConstantCreateRequest(ctx)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.GetReferencedConstantHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetReferencedConstantCreateRequest creates the GetReferencedConstant request.
func (client *EnumClient) GetReferencedConstantCreateRequest(ctx context.Context) (*azcore.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := azcore.NewRequest(ctx, http.MethodGet, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, nil
}

// GetReferencedConstantHandleResponse handles the GetReferencedConstant response.
func (client *EnumClient) GetReferencedConstantHandleResponse(resp *azcore.Response) (*RefColorConstantResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.GetReferencedConstantHandleError(resp)
	}
	result := RefColorConstantResponse{RawResponse: resp.Response}
	return &result, resp.UnmarshalAsJSON(&result.RefColorConstant)
}

// GetReferencedConstantHandleError handles the GetReferencedConstant error response.
func (client *EnumClient) GetReferencedConstantHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutNotExpandable - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
func (client *EnumClient) PutNotExpandable(ctx context.Context, stringBody Colors) (*http.Response, error) {
	req, err := client.PutNotExpandableCreateRequest(ctx, stringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutNotExpandableHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutNotExpandableCreateRequest creates the PutNotExpandable request.
func (client *EnumClient) PutNotExpandableCreateRequest(ctx context.Context, stringBody Colors) (*azcore.Request, error) {
	urlPath := "/string/enum/notExpandable"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(stringBody)
}

// PutNotExpandableHandleResponse handles the PutNotExpandable response.
func (client *EnumClient) PutNotExpandableHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutNotExpandableHandleError(resp)
	}
	return resp.Response, nil
}

// PutNotExpandableHandleError handles the PutNotExpandable error response.
func (client *EnumClient) PutNotExpandableHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutReferenced - Sends value 'red color' from enumeration of 'red color', 'green-color', 'blue_color'
func (client *EnumClient) PutReferenced(ctx context.Context, enumStringBody Colors) (*http.Response, error) {
	req, err := client.PutReferencedCreateRequest(ctx, enumStringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutReferencedHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutReferencedCreateRequest creates the PutReferenced request.
func (client *EnumClient) PutReferencedCreateRequest(ctx context.Context, enumStringBody Colors) (*azcore.Request, error) {
	urlPath := "/string/enum/Referenced"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(enumStringBody)
}

// PutReferencedHandleResponse handles the PutReferenced response.
func (client *EnumClient) PutReferencedHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutReferencedHandleError(resp)
	}
	return resp.Response, nil
}

// PutReferencedHandleError handles the PutReferenced error response.
func (client *EnumClient) PutReferencedHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}

// PutReferencedConstant - Sends value 'green-color' from a constant
func (client *EnumClient) PutReferencedConstant(ctx context.Context, enumStringBody RefColorConstant) (*http.Response, error) {
	req, err := client.PutReferencedConstantCreateRequest(ctx, enumStringBody)
	if err != nil {
		return nil, err
	}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	result, err := client.PutReferencedConstantHandleResponse(resp)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// PutReferencedConstantCreateRequest creates the PutReferencedConstant request.
func (client *EnumClient) PutReferencedConstantCreateRequest(ctx context.Context, enumStringBody RefColorConstant) (*azcore.Request, error) {
	urlPath := "/string/enum/ReferencedConstant"
	req, err := azcore.NewRequest(ctx, http.MethodPut, azcore.JoinPaths(client.u, urlPath))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	return req, req.MarshalAsJSON(enumStringBody)
}

// PutReferencedConstantHandleResponse handles the PutReferencedConstant response.
func (client *EnumClient) PutReferencedConstantHandleResponse(resp *azcore.Response) (*http.Response, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, client.PutReferencedConstantHandleError(resp)
	}
	return resp.Response, nil
}

// PutReferencedConstantHandleError handles the PutReferencedConstant error response.
func (client *EnumClient) PutReferencedConstantHandleError(resp *azcore.Response) error {
	var err Error
	if err := resp.UnmarshalAsJSON(&err); err != nil {
		return err
	}
	return err
}
