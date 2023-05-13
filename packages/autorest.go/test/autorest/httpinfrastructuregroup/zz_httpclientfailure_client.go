//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.
// DO NOT EDIT.

package httpinfrastructuregroup

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/policy"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
)

// HTTPClientFailureClient contains the methods for the HTTPClientFailure group.
// Don't use this type directly, use a constructor function instead.
type HTTPClientFailureClient struct {
	internal *azcore.Client
}

// Delete400 - Return 400 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientDelete400Options contains the optional parameters for the HTTPClientFailureClient.Delete400
//     method.
func (client *HTTPClientFailureClient) Delete400(ctx context.Context, options *HTTPClientFailureClientDelete400Options) (HTTPClientFailureClientDelete400Response, error) {
	req, err := client.delete400CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientDelete400Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientDelete400Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientDelete400Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientDelete400Response{}, nil
}

// delete400CreateRequest creates the Delete400 request.
func (client *HTTPClientFailureClient) delete400CreateRequest(ctx context.Context, options *HTTPClientFailureClientDelete400Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete407 - Return 407 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientDelete407Options contains the optional parameters for the HTTPClientFailureClient.Delete407
//     method.
func (client *HTTPClientFailureClient) Delete407(ctx context.Context, options *HTTPClientFailureClientDelete407Options) (HTTPClientFailureClientDelete407Response, error) {
	req, err := client.delete407CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientDelete407Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientDelete407Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientDelete407Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientDelete407Response{}, nil
}

// delete407CreateRequest creates the Delete407 request.
func (client *HTTPClientFailureClient) delete407CreateRequest(ctx context.Context, options *HTTPClientFailureClientDelete407Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/407"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Delete417 - Return 417 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientDelete417Options contains the optional parameters for the HTTPClientFailureClient.Delete417
//     method.
func (client *HTTPClientFailureClient) Delete417(ctx context.Context, options *HTTPClientFailureClientDelete417Options) (HTTPClientFailureClientDelete417Response, error) {
	req, err := client.delete417CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientDelete417Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientDelete417Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientDelete417Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientDelete417Response{}, nil
}

// delete417CreateRequest creates the Delete417 request.
func (client *HTTPClientFailureClient) delete417CreateRequest(ctx context.Context, options *HTTPClientFailureClientDelete417Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/417"
	req, err := runtime.NewRequest(ctx, http.MethodDelete, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Get400 - Return 400 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientGet400Options contains the optional parameters for the HTTPClientFailureClient.Get400
//     method.
func (client *HTTPClientFailureClient) Get400(ctx context.Context, options *HTTPClientFailureClientGet400Options) (HTTPClientFailureClientGet400Response, error) {
	req, err := client.get400CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientGet400Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientGet400Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientGet400Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientGet400Response{}, nil
}

// get400CreateRequest creates the Get400 request.
func (client *HTTPClientFailureClient) get400CreateRequest(ctx context.Context, options *HTTPClientFailureClientGet400Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get402 - Return 402 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientGet402Options contains the optional parameters for the HTTPClientFailureClient.Get402
//     method.
func (client *HTTPClientFailureClient) Get402(ctx context.Context, options *HTTPClientFailureClientGet402Options) (HTTPClientFailureClientGet402Response, error) {
	req, err := client.get402CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientGet402Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientGet402Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientGet402Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientGet402Response{}, nil
}

// get402CreateRequest creates the Get402 request.
func (client *HTTPClientFailureClient) get402CreateRequest(ctx context.Context, options *HTTPClientFailureClientGet402Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/402"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get403 - Return 403 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientGet403Options contains the optional parameters for the HTTPClientFailureClient.Get403
//     method.
func (client *HTTPClientFailureClient) Get403(ctx context.Context, options *HTTPClientFailureClientGet403Options) (HTTPClientFailureClientGet403Response, error) {
	req, err := client.get403CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientGet403Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientGet403Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientGet403Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientGet403Response{}, nil
}

// get403CreateRequest creates the Get403 request.
func (client *HTTPClientFailureClient) get403CreateRequest(ctx context.Context, options *HTTPClientFailureClientGet403Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/403"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get411 - Return 411 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientGet411Options contains the optional parameters for the HTTPClientFailureClient.Get411
//     method.
func (client *HTTPClientFailureClient) Get411(ctx context.Context, options *HTTPClientFailureClientGet411Options) (HTTPClientFailureClientGet411Response, error) {
	req, err := client.get411CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientGet411Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientGet411Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientGet411Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientGet411Response{}, nil
}

// get411CreateRequest creates the Get411 request.
func (client *HTTPClientFailureClient) get411CreateRequest(ctx context.Context, options *HTTPClientFailureClientGet411Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/411"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get412 - Return 412 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientGet412Options contains the optional parameters for the HTTPClientFailureClient.Get412
//     method.
func (client *HTTPClientFailureClient) Get412(ctx context.Context, options *HTTPClientFailureClientGet412Options) (HTTPClientFailureClientGet412Response, error) {
	req, err := client.get412CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientGet412Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientGet412Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientGet412Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientGet412Response{}, nil
}

// get412CreateRequest creates the Get412 request.
func (client *HTTPClientFailureClient) get412CreateRequest(ctx context.Context, options *HTTPClientFailureClientGet412Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/412"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Get416 - Return 416 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientGet416Options contains the optional parameters for the HTTPClientFailureClient.Get416
//     method.
func (client *HTTPClientFailureClient) Get416(ctx context.Context, options *HTTPClientFailureClientGet416Options) (HTTPClientFailureClientGet416Response, error) {
	req, err := client.get416CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientGet416Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientGet416Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientGet416Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientGet416Response{}, nil
}

// get416CreateRequest creates the Get416 request.
func (client *HTTPClientFailureClient) get416CreateRequest(ctx context.Context, options *HTTPClientFailureClientGet416Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/416"
	req, err := runtime.NewRequest(ctx, http.MethodGet, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head400 - Return 400 status code - should be represented in the client as an error
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientHead400Options contains the optional parameters for the HTTPClientFailureClient.Head400
//     method.
func (client *HTTPClientFailureClient) Head400(ctx context.Context, options *HTTPClientFailureClientHead400Options) (HTTPClientFailureClientHead400Response, error) {
	req, err := client.head400CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientHead400Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientHead400Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientHead400Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientHead400Response{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// head400CreateRequest creates the Head400 request.
func (client *HTTPClientFailureClient) head400CreateRequest(ctx context.Context, options *HTTPClientFailureClientHead400Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head401 - Return 401 status code - should be represented in the client as an error
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientHead401Options contains the optional parameters for the HTTPClientFailureClient.Head401
//     method.
func (client *HTTPClientFailureClient) Head401(ctx context.Context, options *HTTPClientFailureClientHead401Options) (HTTPClientFailureClientHead401Response, error) {
	req, err := client.head401CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientHead401Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientHead401Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientHead401Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientHead401Response{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// head401CreateRequest creates the Head401 request.
func (client *HTTPClientFailureClient) head401CreateRequest(ctx context.Context, options *HTTPClientFailureClientHead401Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/401"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head410 - Return 410 status code - should be represented in the client as an error
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientHead410Options contains the optional parameters for the HTTPClientFailureClient.Head410
//     method.
func (client *HTTPClientFailureClient) Head410(ctx context.Context, options *HTTPClientFailureClientHead410Options) (HTTPClientFailureClientHead410Response, error) {
	req, err := client.head410CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientHead410Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientHead410Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientHead410Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientHead410Response{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// head410CreateRequest creates the Head410 request.
func (client *HTTPClientFailureClient) head410CreateRequest(ctx context.Context, options *HTTPClientFailureClientHead410Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/410"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Head429 - Return 429 status code - should be represented in the client as an error
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientHead429Options contains the optional parameters for the HTTPClientFailureClient.Head429
//     method.
func (client *HTTPClientFailureClient) Head429(ctx context.Context, options *HTTPClientFailureClientHead429Options) (HTTPClientFailureClientHead429Response, error) {
	req, err := client.head429CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientHead429Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientHead429Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientHead429Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientHead429Response{Success: resp.StatusCode >= 200 && resp.StatusCode < 300}, nil
}

// head429CreateRequest creates the Head429 request.
func (client *HTTPClientFailureClient) head429CreateRequest(ctx context.Context, options *HTTPClientFailureClientHead429Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/429"
	req, err := runtime.NewRequest(ctx, http.MethodHead, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Options400 - Return 400 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientOptions400Options contains the optional parameters for the HTTPClientFailureClient.Options400
//     method.
func (client *HTTPClientFailureClient) Options400(ctx context.Context, options *HTTPClientFailureClientOptions400Options) (HTTPClientFailureClientOptions400Response, error) {
	req, err := client.options400CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientOptions400Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientOptions400Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientOptions400Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientOptions400Response{}, nil
}

// options400CreateRequest creates the Options400 request.
func (client *HTTPClientFailureClient) options400CreateRequest(ctx context.Context, options *HTTPClientFailureClientOptions400Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Options403 - Return 403 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientOptions403Options contains the optional parameters for the HTTPClientFailureClient.Options403
//     method.
func (client *HTTPClientFailureClient) Options403(ctx context.Context, options *HTTPClientFailureClientOptions403Options) (HTTPClientFailureClientOptions403Response, error) {
	req, err := client.options403CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientOptions403Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientOptions403Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientOptions403Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientOptions403Response{}, nil
}

// options403CreateRequest creates the Options403 request.
func (client *HTTPClientFailureClient) options403CreateRequest(ctx context.Context, options *HTTPClientFailureClientOptions403Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/403"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Options412 - Return 412 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientOptions412Options contains the optional parameters for the HTTPClientFailureClient.Options412
//     method.
func (client *HTTPClientFailureClient) Options412(ctx context.Context, options *HTTPClientFailureClientOptions412Options) (HTTPClientFailureClientOptions412Response, error) {
	req, err := client.options412CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientOptions412Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientOptions412Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientOptions412Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientOptions412Response{}, nil
}

// options412CreateRequest creates the Options412 request.
func (client *HTTPClientFailureClient) options412CreateRequest(ctx context.Context, options *HTTPClientFailureClientOptions412Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/412"
	req, err := runtime.NewRequest(ctx, http.MethodOptions, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	return req, nil
}

// Patch400 - Return 400 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPatch400Options contains the optional parameters for the HTTPClientFailureClient.Patch400
//     method.
func (client *HTTPClientFailureClient) Patch400(ctx context.Context, options *HTTPClientFailureClientPatch400Options) (HTTPClientFailureClientPatch400Response, error) {
	req, err := client.patch400CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPatch400Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPatch400Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPatch400Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPatch400Response{}, nil
}

// patch400CreateRequest creates the Patch400 request.
func (client *HTTPClientFailureClient) patch400CreateRequest(ctx context.Context, options *HTTPClientFailureClientPatch400Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Patch405 - Return 405 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPatch405Options contains the optional parameters for the HTTPClientFailureClient.Patch405
//     method.
func (client *HTTPClientFailureClient) Patch405(ctx context.Context, options *HTTPClientFailureClientPatch405Options) (HTTPClientFailureClientPatch405Response, error) {
	req, err := client.patch405CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPatch405Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPatch405Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPatch405Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPatch405Response{}, nil
}

// patch405CreateRequest creates the Patch405 request.
func (client *HTTPClientFailureClient) patch405CreateRequest(ctx context.Context, options *HTTPClientFailureClientPatch405Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/405"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Patch414 - Return 414 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPatch414Options contains the optional parameters for the HTTPClientFailureClient.Patch414
//     method.
func (client *HTTPClientFailureClient) Patch414(ctx context.Context, options *HTTPClientFailureClientPatch414Options) (HTTPClientFailureClientPatch414Response, error) {
	req, err := client.patch414CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPatch414Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPatch414Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPatch414Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPatch414Response{}, nil
}

// patch414CreateRequest creates the Patch414 request.
func (client *HTTPClientFailureClient) patch414CreateRequest(ctx context.Context, options *HTTPClientFailureClientPatch414Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/414"
	req, err := runtime.NewRequest(ctx, http.MethodPatch, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post400 - Return 400 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPost400Options contains the optional parameters for the HTTPClientFailureClient.Post400
//     method.
func (client *HTTPClientFailureClient) Post400(ctx context.Context, options *HTTPClientFailureClientPost400Options) (HTTPClientFailureClientPost400Response, error) {
	req, err := client.post400CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPost400Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPost400Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPost400Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPost400Response{}, nil
}

// post400CreateRequest creates the Post400 request.
func (client *HTTPClientFailureClient) post400CreateRequest(ctx context.Context, options *HTTPClientFailureClientPost400Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post406 - Return 406 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPost406Options contains the optional parameters for the HTTPClientFailureClient.Post406
//     method.
func (client *HTTPClientFailureClient) Post406(ctx context.Context, options *HTTPClientFailureClientPost406Options) (HTTPClientFailureClientPost406Response, error) {
	req, err := client.post406CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPost406Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPost406Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPost406Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPost406Response{}, nil
}

// post406CreateRequest creates the Post406 request.
func (client *HTTPClientFailureClient) post406CreateRequest(ctx context.Context, options *HTTPClientFailureClientPost406Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/406"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Post415 - Return 415 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPost415Options contains the optional parameters for the HTTPClientFailureClient.Post415
//     method.
func (client *HTTPClientFailureClient) Post415(ctx context.Context, options *HTTPClientFailureClientPost415Options) (HTTPClientFailureClientPost415Response, error) {
	req, err := client.post415CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPost415Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPost415Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPost415Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPost415Response{}, nil
}

// post415CreateRequest creates the Post415 request.
func (client *HTTPClientFailureClient) post415CreateRequest(ctx context.Context, options *HTTPClientFailureClientPost415Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/415"
	req, err := runtime.NewRequest(ctx, http.MethodPost, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put400 - Return 400 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPut400Options contains the optional parameters for the HTTPClientFailureClient.Put400
//     method.
func (client *HTTPClientFailureClient) Put400(ctx context.Context, options *HTTPClientFailureClientPut400Options) (HTTPClientFailureClientPut400Response, error) {
	req, err := client.put400CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPut400Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPut400Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPut400Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPut400Response{}, nil
}

// put400CreateRequest creates the Put400 request.
func (client *HTTPClientFailureClient) put400CreateRequest(ctx context.Context, options *HTTPClientFailureClientPut400Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/400"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put404 - Return 404 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPut404Options contains the optional parameters for the HTTPClientFailureClient.Put404
//     method.
func (client *HTTPClientFailureClient) Put404(ctx context.Context, options *HTTPClientFailureClientPut404Options) (HTTPClientFailureClientPut404Response, error) {
	req, err := client.put404CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPut404Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPut404Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPut404Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPut404Response{}, nil
}

// put404CreateRequest creates the Put404 request.
func (client *HTTPClientFailureClient) put404CreateRequest(ctx context.Context, options *HTTPClientFailureClientPut404Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/404"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put409 - Return 409 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPut409Options contains the optional parameters for the HTTPClientFailureClient.Put409
//     method.
func (client *HTTPClientFailureClient) Put409(ctx context.Context, options *HTTPClientFailureClientPut409Options) (HTTPClientFailureClientPut409Response, error) {
	req, err := client.put409CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPut409Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPut409Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPut409Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPut409Response{}, nil
}

// put409CreateRequest creates the Put409 request.
func (client *HTTPClientFailureClient) put409CreateRequest(ctx context.Context, options *HTTPClientFailureClientPut409Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/409"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}

// Put413 - Return 413 status code - should be represented in the client as an error
// If the operation fails it returns an *azcore.ResponseError type.
//
// Generated from API version 1.0.0
//   - options - HTTPClientFailureClientPut413Options contains the optional parameters for the HTTPClientFailureClient.Put413
//     method.
func (client *HTTPClientFailureClient) Put413(ctx context.Context, options *HTTPClientFailureClientPut413Options) (HTTPClientFailureClientPut413Response, error) {
	req, err := client.put413CreateRequest(ctx, options)
	if err != nil {
		return HTTPClientFailureClientPut413Response{}, err
	}
	resp, err := client.internal.Pipeline().Do(req)
	if err != nil {
		return HTTPClientFailureClientPut413Response{}, err
	}
	if !runtime.HasStatusCode(resp, http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent) {
		return HTTPClientFailureClientPut413Response{}, runtime.NewResponseError(resp)
	}
	return HTTPClientFailureClientPut413Response{}, nil
}

// put413CreateRequest creates the Put413 request.
func (client *HTTPClientFailureClient) put413CreateRequest(ctx context.Context, options *HTTPClientFailureClientPut413Options) (*policy.Request, error) {
	urlPath := "/http/failure/client/413"
	req, err := runtime.NewRequest(ctx, http.MethodPut, runtime.JoinPaths(host, urlPath))
	if err != nil {
		return nil, err
	}
	req.Raw().Header["Accept"] = []string{"application/json"}
	if err := runtime.MarshalAsJSON(req, true); err != nil {
		return nil, err
	}
	return req, nil
}
