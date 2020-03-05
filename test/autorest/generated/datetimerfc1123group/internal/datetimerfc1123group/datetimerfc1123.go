// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package datetimerfc1123group

import (
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"net/url"
	"path"
	"time"
)

// Datetimerfc1123Operations contains the methods for the Datetimerfc1123 group.
type Datetimerfc1123Operations struct{}

// GetInvalidCreateRequest creates the GetInvalid request.
func (Datetimerfc1123Operations) GetInvalidCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/invalid"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetInvalidHandleResponse handles the GetInvalid response.
func (Datetimerfc1123Operations) GetInvalidHandleResponse(resp *azcore.Response) (*Datetimerfc1123GetInvalidResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	result := Datetimerfc1123GetInvalidResponse{RawResponse: resp.Response}
	result.Value = aux.ToTime()
	return &result, err
}

// GetNullCreateRequest creates the GetNull request.
func (Datetimerfc1123Operations) GetNullCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/null"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetNullHandleResponse handles the GetNull response.
func (Datetimerfc1123Operations) GetNullHandleResponse(resp *azcore.Response) (*Datetimerfc1123GetNullResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	result := Datetimerfc1123GetNullResponse{RawResponse: resp.Response}
	result.Value = aux.ToTime()
	return &result, err
}

// GetOverflowCreateRequest creates the GetOverflow request.
func (Datetimerfc1123Operations) GetOverflowCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/overflow"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetOverflowHandleResponse handles the GetOverflow response.
func (Datetimerfc1123Operations) GetOverflowHandleResponse(resp *azcore.Response) (*Datetimerfc1123GetOverflowResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	result := Datetimerfc1123GetOverflowResponse{RawResponse: resp.Response}
	result.Value = aux.ToTime()
	return &result, err
}

// GetUTCLowercaseMaxDateTimeCreateRequest creates the GetUTCLowercaseMaxDateTime request.
func (Datetimerfc1123Operations) GetUTCLowercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/lowercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetUTCLowercaseMaxDateTimeHandleResponse handles the GetUTCLowercaseMaxDateTime response.
func (Datetimerfc1123Operations) GetUTCLowercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	result := Datetimerfc1123GetUTCLowercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = aux.ToTime()
	return &result, err
}

// GetUTCMinDateTimeCreateRequest creates the GetUTCMinDateTime request.
func (Datetimerfc1123Operations) GetUTCMinDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetUTCMinDateTimeHandleResponse handles the GetUTCMinDateTime response.
func (Datetimerfc1123Operations) GetUTCMinDateTimeHandleResponse(resp *azcore.Response) (*Datetimerfc1123GetUTCMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	result := Datetimerfc1123GetUTCMinDateTimeResponse{RawResponse: resp.Response}
	result.Value = aux.ToTime()
	return &result, err
}

// GetUTCUppercaseMaxDateTimeCreateRequest creates the GetUTCUppercaseMaxDateTime request.
func (Datetimerfc1123Operations) GetUTCUppercaseMaxDateTimeCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max/uppercase"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetUTCUppercaseMaxDateTimeHandleResponse handles the GetUTCUppercaseMaxDateTime response.
func (Datetimerfc1123Operations) GetUTCUppercaseMaxDateTimeHandleResponse(resp *azcore.Response) (*Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	result := Datetimerfc1123GetUTCUppercaseMaxDateTimeResponse{RawResponse: resp.Response}
	result.Value = aux.ToTime()
	return &result, err
}

// GetUnderflowCreateRequest creates the GetUnderflow request.
func (Datetimerfc1123Operations) GetUnderflowCreateRequest(u url.URL) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/underflow"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodGet, u)
	return req, nil
}

// GetUnderflowHandleResponse handles the GetUnderflow response.
func (Datetimerfc1123Operations) GetUnderflowHandleResponse(resp *azcore.Response) (*Datetimerfc1123GetUnderflowResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	var aux *timeRFC1123
	err := resp.UnmarshalAsJSON(&aux)
	result := Datetimerfc1123GetUnderflowResponse{RawResponse: resp.Response}
	result.Value = aux.ToTime()
	return &result, err
}

// PutUTCMaxDateTimeCreateRequest creates the PutUTCMaxDateTime request.
func (Datetimerfc1123Operations) PutUTCMaxDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/max"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	aux := timeRFC1123(datetimeBody)
	err := req.MarshalAsJSON(aux)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutUTCMaxDateTimeHandleResponse handles the PutUTCMaxDateTime response.
func (Datetimerfc1123Operations) PutUTCMaxDateTimeHandleResponse(resp *azcore.Response) (*Datetimerfc1123PutUTCMaxDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &Datetimerfc1123PutUTCMaxDateTimeResponse{RawResponse: resp.Response}, nil
}

// PutUTCMinDateTimeCreateRequest creates the PutUTCMinDateTime request.
func (Datetimerfc1123Operations) PutUTCMinDateTimeCreateRequest(u url.URL, datetimeBody time.Time) (*azcore.Request, error) {
	urlPath := "/datetimerfc1123/min"
	u.Path = path.Join(u.Path, urlPath)
	req := azcore.NewRequest(http.MethodPut, u)
	aux := timeRFC1123(datetimeBody)
	err := req.MarshalAsJSON(aux)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// PutUTCMinDateTimeHandleResponse handles the PutUTCMinDateTime response.
func (Datetimerfc1123Operations) PutUTCMinDateTimeHandleResponse(resp *azcore.Response) (*Datetimerfc1123PutUTCMinDateTimeResponse, error) {
	if !resp.HasStatusCode(http.StatusOK) {
		return nil, newError(resp)
	}
	return &Datetimerfc1123PutUTCMinDateTimeResponse{RawResponse: resp.Response}, nil
}
