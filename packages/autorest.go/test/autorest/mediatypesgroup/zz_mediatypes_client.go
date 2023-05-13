//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package mediatypesgroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/streaming"
	"io"
	"net/http"
	"strings"
)

// MediaTypesClient contains the methods for the MediaTypesClient group.
// Don't use this type directly, use a constructor function instead.
type MediaTypesClient struct {
	internal *azcore.Client
}

// AnalyzeBody - Analyze body, that could be different media types.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - options - MediaTypesClientAnalyzeBodyOptions contains the optional parameters for the MediaTypesClient.AnalyzeBody method.
func (client *MediaTypesClient) AnalyzeBody(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyOptions) (MediaTypesClientAnalyzeBodyResponse, error) {
	req, err := client.analyzeBodyCreateRequest(ctx, contentType, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientAnalyzeBodyResponse{}, runtime.NewResponseError(resp)
	}
	return client.analyzeBodyHandleResponse(resp)
}

// analyzeBodyCreateRequest creates the AnalyzeBody request.
func (client *MediaTypesClient) analyzeBodyCreateRequest(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Input != nil {
		if err := req.SetBody(options.Input, string(contentType)); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// analyzeBodyHandleResponse handles the AnalyzeBody response.
func (client *MediaTypesClient) analyzeBodyHandleResponse(resp *http.Response) (MediaTypesClientAnalyzeBodyResponse, error) {
	result := MediaTypesClientAnalyzeBodyResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyResponse{}, err
	}
	return result, nil
}

// AnalyzeBodyNoAcceptHeader - Analyze body, that could be different media types. Adds to AnalyzeBody by not having an accept
// type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - options - MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyNoAcceptHeader
//     method.
func (client *MediaTypesClient) AnalyzeBodyNoAcceptHeader(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse, error) {
	req, err := client.analyzeBodyNoAcceptHeaderCreateRequest(ctx, contentType, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, runtime.NewResponseError(resp)
	}
	return MediaTypesClientAnalyzeBodyNoAcceptHeaderResponse{}, nil
}

// analyzeBodyNoAcceptHeaderCreateRequest creates the AnalyzeBodyNoAcceptHeader request.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderCreateRequest(ctx context.Context, contentType ContentType, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyzeNoAccept"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	if options != nil && options.Input != nil {
		if err := req.SetBody(options.Input, string(contentType)); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// AnalyzeBodyNoAcceptHeaderWithJSON - Analyze body, that could be different media types. Adds to AnalyzeBody by not having
// an accept type.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyNoAcceptHeaderWithJSON
//     method.
func (client *MediaTypesClient) AnalyzeBodyNoAcceptHeaderWithJSON(ctx context.Context, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONOptions) (MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONResponse, error) {
	req, err := client.analyzeBodyNoAcceptHeaderWithJSONCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusAccepted) {
		return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONResponse{}, runtime.NewResponseError(resp)
	}
	return MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONResponse{}, nil
}

// analyzeBodyNoAcceptHeaderWithJSONCreateRequest creates the AnalyzeBodyNoAcceptHeaderWithJSON request.
func (client *MediaTypesClient) analyzeBodyNoAcceptHeaderWithJSONCreateRequest(ctx context.Context, options *MediaTypesClientAnalyzeBodyNoAcceptHeaderWithJSONOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyzeNoAccept"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	if options != nil && options.Input != nil {
		if err := runtime.MarshalAsJSON(req, *options.Input); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// AnalyzeBodyWithJSON - Analyze body, that could be different media types.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - MediaTypesClientAnalyzeBodyWithJSONOptions contains the optional parameters for the MediaTypesClient.AnalyzeBodyWithJSON
//     method.
func (client *MediaTypesClient) AnalyzeBodyWithJSON(ctx context.Context, options *MediaTypesClientAnalyzeBodyWithJSONOptions) (MediaTypesClientAnalyzeBodyWithJSONResponse, error) {
	req, err := client.analyzeBodyWithJSONCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithJSONResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientAnalyzeBodyWithJSONResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientAnalyzeBodyWithJSONResponse{}, runtime.NewResponseError(resp)
	}
	return client.analyzeBodyWithJSONHandleResponse(resp)
}

// analyzeBodyWithJSONCreateRequest creates the AnalyzeBodyWithJSON request.
func (client *MediaTypesClient) analyzeBodyWithJSONCreateRequest(ctx context.Context, options *MediaTypesClientAnalyzeBodyWithJSONOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/analyze"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Input != nil {
		if err := runtime.MarshalAsJSON(req, *options.Input); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// analyzeBodyWithJSONHandleResponse handles the AnalyzeBodyWithJSON response.
func (client *MediaTypesClient) analyzeBodyWithJSONHandleResponse(resp *http.Response) (MediaTypesClientAnalyzeBodyWithJSONResponse, error) {
	result := MediaTypesClientAnalyzeBodyWithJSONResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientAnalyzeBodyWithJSONResponse{}, err
	}
	return result, nil
}

// BinaryBodyWithThreeContentTypes - Binary body with three content types. Pass in string 'hello, world' with content type
// 'text/plain', {'hello': world'} with content type 'application/json' and a byte string for
// 'application/octet-stream'.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - message - The payload body.
//   - options - MediaTypesClientBinaryBodyWithThreeContentTypesOptions contains the optional parameters for the MediaTypesClient.BinaryBodyWithThreeContentTypes
//     method.
func (client *MediaTypesClient) BinaryBodyWithThreeContentTypes(ctx context.Context, contentType ContentType1AutoGenerated, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithThreeContentTypesOptions) (MediaTypesClientBinaryBodyWithThreeContentTypesResponse, error) {
	req, err := client.binaryBodyWithThreeContentTypesCreateRequest(ctx, contentType, message, options)
	if err != nil {
		return MediaTypesClientBinaryBodyWithThreeContentTypesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientBinaryBodyWithThreeContentTypesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientBinaryBodyWithThreeContentTypesResponse{}, runtime.NewResponseError(resp)
	}
	return client.binaryBodyWithThreeContentTypesHandleResponse(resp)
}

// binaryBodyWithThreeContentTypesCreateRequest creates the BinaryBodyWithThreeContentTypes request.
func (client *MediaTypesClient) binaryBodyWithThreeContentTypesCreateRequest(ctx context.Context, contentType ContentType1AutoGenerated, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithThreeContentTypesOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/binaryBodyThreeContentTypes"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"text/plain"}
	if err := req.SetBody(message, string(contentType)); err != nil {
		return nil, err
	}
	return req, nil
}

// binaryBodyWithThreeContentTypesHandleResponse handles the BinaryBodyWithThreeContentTypes response.
func (client *MediaTypesClient) binaryBodyWithThreeContentTypesHandleResponse(resp *http.Response) (MediaTypesClientBinaryBodyWithThreeContentTypesResponse, error) {
	result := MediaTypesClientBinaryBodyWithThreeContentTypesResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return MediaTypesClientBinaryBodyWithThreeContentTypesResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// BinaryBodyWithTwoContentTypes - Binary body with two content types. Pass in of {'hello': 'world'} for the application/json
// content type, and a byte stream of 'hello, world!' for application/octet-stream.
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - message - The payload body.
//   - options - MediaTypesClientBinaryBodyWithTwoContentTypesOptions contains the optional parameters for the MediaTypesClient.BinaryBodyWithTwoContentTypes
//     method.
func (client *MediaTypesClient) BinaryBodyWithTwoContentTypes(ctx context.Context, contentType ContentType1, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithTwoContentTypesOptions) (MediaTypesClientBinaryBodyWithTwoContentTypesResponse, error) {
	req, err := client.binaryBodyWithTwoContentTypesCreateRequest(ctx, contentType, message, options)
	if err != nil {
		return MediaTypesClientBinaryBodyWithTwoContentTypesResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientBinaryBodyWithTwoContentTypesResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientBinaryBodyWithTwoContentTypesResponse{}, runtime.NewResponseError(resp)
	}
	return client.binaryBodyWithTwoContentTypesHandleResponse(resp)
}

// binaryBodyWithTwoContentTypesCreateRequest creates the BinaryBodyWithTwoContentTypes request.
func (client *MediaTypesClient) binaryBodyWithTwoContentTypesCreateRequest(ctx context.Context, contentType ContentType1, message io.ReadSeekCloser, options *MediaTypesClientBinaryBodyWithTwoContentTypesOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/binaryBodyTwoContentTypes"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"text/plain"}
	if err := req.SetBody(message, string(contentType)); err != nil {
		return nil, err
	}
	return req, nil
}

// binaryBodyWithTwoContentTypesHandleResponse handles the BinaryBodyWithTwoContentTypes response.
func (client *MediaTypesClient) binaryBodyWithTwoContentTypesHandleResponse(resp *http.Response) (MediaTypesClientBinaryBodyWithTwoContentTypesResponse, error) {
	result := MediaTypesClientBinaryBodyWithTwoContentTypesResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return MediaTypesClientBinaryBodyWithTwoContentTypesResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}

// ContentTypeWithEncoding - Pass in contentType 'text/plain; charset=UTF-8' to pass test. Value for input does not matter
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - options - MediaTypesClientContentTypeWithEncodingOptions contains the optional parameters for the MediaTypesClient.ContentTypeWithEncoding
//     method.
func (client *MediaTypesClient) ContentTypeWithEncoding(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingOptions) (MediaTypesClientContentTypeWithEncodingResponse, error) {
	req, err := client.contentTypeWithEncodingCreateRequest(ctx, options)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientContentTypeWithEncodingResponse{}, runtime.NewResponseError(resp)
	}
	return client.contentTypeWithEncodingHandleResponse(resp)
}

// contentTypeWithEncodingCreateRequest creates the ContentTypeWithEncoding request.
func (client *MediaTypesClient) contentTypeWithEncodingCreateRequest(ctx context.Context, options *MediaTypesClientContentTypeWithEncodingOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/contentTypeWithEncoding"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if options != nil && options.Input != nil {
		body := streaming.NopCloser(strings.NewReader(*options.Input))
		if err := req.SetBody(body, "text/plain; charset=UTF-8"); err != nil {
			return nil, err
		}
		return req, nil
	}
	return req, nil
}

// contentTypeWithEncodingHandleResponse handles the ContentTypeWithEncoding response.
func (client *MediaTypesClient) contentTypeWithEncodingHandleResponse(resp *http.Response) (MediaTypesClientContentTypeWithEncodingResponse, error) {
	result := MediaTypesClientContentTypeWithEncodingResponse{}
	if err := runtime.UnmarshalAsJSON(resp, &result.Value); err != nil {
		return MediaTypesClientContentTypeWithEncodingResponse{}, err
	}
	return result, nil
}

// PutTextAndJSONBody - Body that's either text/plain or application/json
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 2.0-preview
//   - contentType - Upload file type
//   - message - The payload body.
//   - options - MediaTypesClientPutTextAndJSONBodyOptions contains the optional parameters for the MediaTypesClient.PutTextAndJSONBody
//     method.
func (client *MediaTypesClient) PutTextAndJSONBody(ctx context.Context, contentType ContentType1AutoGenerated2, message string, options *MediaTypesClientPutTextAndJSONBodyOptions) (MediaTypesClientPutTextAndJSONBodyResponse, error) {
	req, err := client.putTextAndJSONBodyCreateRequest(ctx, contentType, message, options)
	if err != nil {
		return MediaTypesClientPutTextAndJSONBodyResponse{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return MediaTypesClientPutTextAndJSONBodyResponse{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK) {
		return MediaTypesClientPutTextAndJSONBodyResponse{}, runtime.NewResponseError(resp)
	}
	return client.putTextAndJSONBodyHandleResponse(resp)
}

// putTextAndJSONBodyCreateRequest creates the PutTextAndJSONBody request.
func (client *MediaTypesClient) putTextAndJSONBodyCreateRequest(ctx context.Context, contentType ContentType1AutoGenerated2, message string, options *MediaTypesClientPutTextAndJSONBodyOptions) (*policy.Request, error) {
	urlPath := "/mediatypes/textAndJson"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Content-Type"] = []string{string(contentType)}
	req.Raw().Header["Accept"] = []string{"text/plain"}
	body := streaming.NopCloser(strings.NewReader(message))
	if err := req.SetBody(body, "application/json"); err != nil {
		return nil, err
	}

	return req, nil
}

// putTextAndJSONBodyHandleResponse handles the PutTextAndJSONBody response.
func (client *MediaTypesClient) putTextAndJSONBodyHandleResponse(resp *http.Response) (MediaTypesClientPutTextAndJSONBodyResponse, error) {
	result := MediaTypesClientPutTextAndJSONBodyResponse{}
	body, err := runtime.Payload(resp)
	if err != nil {
		return MediaTypesClientPutTextAndJSONBodyResponse{}, err
	}
	txt := string(body)
	result.Value = &txt
	return result, nil
}
