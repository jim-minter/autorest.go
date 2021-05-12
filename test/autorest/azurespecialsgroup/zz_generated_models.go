// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azurespecialsgroup

import "net/http"

// APIVersionDefaultGetMethodGlobalNotProvidedValidOptions contains the optional parameters for the APIVersionDefault.GetMethodGlobalNotProvidedValid method.
type APIVersionDefaultGetMethodGlobalNotProvidedValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionDefaultGetMethodGlobalValidOptions contains the optional parameters for the APIVersionDefault.GetMethodGlobalValid method.
type APIVersionDefaultGetMethodGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionDefaultGetPathGlobalValidOptions contains the optional parameters for the APIVersionDefault.GetPathGlobalValid method.
type APIVersionDefaultGetPathGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionDefaultGetSwaggerGlobalValidOptions contains the optional parameters for the APIVersionDefault.GetSwaggerGlobalValid method.
type APIVersionDefaultGetSwaggerGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionLocalGetMethodLocalNullOptions contains the optional parameters for the APIVersionLocal.GetMethodLocalNull method.
type APIVersionLocalGetMethodLocalNullOptions struct {
	// This should appear as a method parameter, use value null, this should result in no serialized parameter
	APIVersion *string
}

// APIVersionLocalGetMethodLocalValidOptions contains the optional parameters for the APIVersionLocal.GetMethodLocalValid method.
type APIVersionLocalGetMethodLocalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionLocalGetPathLocalValidOptions contains the optional parameters for the APIVersionLocal.GetPathLocalValid method.
type APIVersionLocalGetPathLocalValidOptions struct {
	// placeholder for future optional parameters
}

// APIVersionLocalGetSwaggerLocalValidOptions contains the optional parameters for the APIVersionLocal.GetSwaggerLocalValid method.
type APIVersionLocalGetSwaggerLocalValidOptions struct {
	// placeholder for future optional parameters
}

// Implements the error and azcore.HTTPResponse interfaces.
type Error struct {
	raw string
	// REQUIRED
	ConstantID *int32  `json:"constantId,omitempty"`
	Message    *string `json:"message,omitempty"`
	Status     *int32  `json:"status,omitempty"`
}

// Error implements the error interface for type Error.
// The contents of the error text are not contractual and subject to change.
func (e Error) Error() string {
	return e.raw
}

// HeaderCustomNamedRequestIDHeadOptions contains the optional parameters for the Header.CustomNamedRequestIDHead method.
type HeaderCustomNamedRequestIDHeadOptions struct {
	// placeholder for future optional parameters
}

// HeaderCustomNamedRequestIDHeadResponse contains the response from method Header.CustomNamedRequestIDHead.
type HeaderCustomNamedRequestIDHeadResponse struct {
	// FooRequestID contains the information returned from the foo-request-id header response.
	FooRequestID *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response

	// Success indicates if the operation succeeded or failed.
	Success bool
}

// HeaderCustomNamedRequestIDOptions contains the optional parameters for the Header.CustomNamedRequestID method.
type HeaderCustomNamedRequestIDOptions struct {
	// placeholder for future optional parameters
}

// HeaderCustomNamedRequestIDParamGroupingOptions contains the optional parameters for the Header.CustomNamedRequestIDParamGrouping method.
type HeaderCustomNamedRequestIDParamGroupingOptions struct {
	// placeholder for future optional parameters
}

// HeaderCustomNamedRequestIDParamGroupingParameters contains a group of parameters for the Header.CustomNamedRequestIDParamGrouping method.
type HeaderCustomNamedRequestIDParamGroupingParameters struct {
	// The fooRequestId
	FooClientRequestID string
}

// HeaderCustomNamedRequestIDParamGroupingResponse contains the response from method Header.CustomNamedRequestIDParamGrouping.
type HeaderCustomNamedRequestIDParamGroupingResponse struct {
	// FooRequestID contains the information returned from the foo-request-id header response.
	FooRequestID *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

// HeaderCustomNamedRequestIDResponse contains the response from method Header.CustomNamedRequestID.
type HeaderCustomNamedRequestIDResponse struct {
	// FooRequestID contains the information returned from the foo-request-id header response.
	FooRequestID *string

	// RawResponse contains the underlying HTTP response.
	RawResponse *http.Response
}

type OdataFilter struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// OdataGetWithFilterOptions contains the optional parameters for the Odata.GetWithFilter method.
type OdataGetWithFilterOptions struct {
	// The filter parameter with value '$filter=id gt 5 and name eq 'foo''.
	Filter *string
	// The orderby parameter with value id.
	Orderby *string
	// The top parameter with value 10.
	Top *int32
}

// SkipURLEncodingGetMethodPathValidOptions contains the optional parameters for the SkipURLEncoding.GetMethodPathValid method.
type SkipURLEncodingGetMethodPathValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingGetMethodQueryNullOptions contains the optional parameters for the SkipURLEncoding.GetMethodQueryNull method.
type SkipURLEncodingGetMethodQueryNullOptions struct {
	// Unencoded query parameter with value null
	Q1 *string
}

// SkipURLEncodingGetMethodQueryValidOptions contains the optional parameters for the SkipURLEncoding.GetMethodQueryValid method.
type SkipURLEncodingGetMethodQueryValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingGetPathQueryValidOptions contains the optional parameters for the SkipURLEncoding.GetPathQueryValid method.
type SkipURLEncodingGetPathQueryValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingGetPathValidOptions contains the optional parameters for the SkipURLEncoding.GetPathValid method.
type SkipURLEncodingGetPathValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingGetSwaggerPathValidOptions contains the optional parameters for the SkipURLEncoding.GetSwaggerPathValid method.
type SkipURLEncodingGetSwaggerPathValidOptions struct {
	// placeholder for future optional parameters
}

// SkipURLEncodingGetSwaggerQueryValidOptions contains the optional parameters for the SkipURLEncoding.GetSwaggerQueryValid method.
type SkipURLEncodingGetSwaggerQueryValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsPostMethodGlobalNotProvidedValidOptions contains the optional parameters for the SubscriptionInCredentials.PostMethodGlobalNotProvidedValid
// method.
type SubscriptionInCredentialsPostMethodGlobalNotProvidedValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsPostMethodGlobalNullOptions contains the optional parameters for the SubscriptionInCredentials.PostMethodGlobalNull method.
type SubscriptionInCredentialsPostMethodGlobalNullOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsPostMethodGlobalValidOptions contains the optional parameters for the SubscriptionInCredentials.PostMethodGlobalValid method.
type SubscriptionInCredentialsPostMethodGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsPostPathGlobalValidOptions contains the optional parameters for the SubscriptionInCredentials.PostPathGlobalValid method.
type SubscriptionInCredentialsPostPathGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInCredentialsPostSwaggerGlobalValidOptions contains the optional parameters for the SubscriptionInCredentials.PostSwaggerGlobalValid method.
type SubscriptionInCredentialsPostSwaggerGlobalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodPostMethodLocalNullOptions contains the optional parameters for the SubscriptionInMethod.PostMethodLocalNull method.
type SubscriptionInMethodPostMethodLocalNullOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodPostMethodLocalValidOptions contains the optional parameters for the SubscriptionInMethod.PostMethodLocalValid method.
type SubscriptionInMethodPostMethodLocalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodPostPathLocalValidOptions contains the optional parameters for the SubscriptionInMethod.PostPathLocalValid method.
type SubscriptionInMethodPostPathLocalValidOptions struct {
	// placeholder for future optional parameters
}

// SubscriptionInMethodPostSwaggerLocalValidOptions contains the optional parameters for the SubscriptionInMethod.PostSwaggerLocalValid method.
type SubscriptionInMethodPostSwaggerLocalValidOptions struct {
	// placeholder for future optional parameters
}

// XMSClientRequestIDGetOptions contains the optional parameters for the XMSClientRequestID.Get method.
type XMSClientRequestIDGetOptions struct {
	// placeholder for future optional parameters
}

// XMSClientRequestIDParamGetOptions contains the optional parameters for the XMSClientRequestID.ParamGet method.
type XMSClientRequestIDParamGetOptions struct {
	// placeholder for future optional parameters
}
