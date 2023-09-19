//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package fake

import (
	"armdataboxedge/v2"
	"context"
	"errors"
	"fmt"
	azfake "github.com/Azure/azure-sdk-for-go/sdk/azcore/fake"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/fake/server"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore/runtime"
	"net/http"
	"net/url"
	"regexp"
)

// SupportPackagesServer is a fake server for instances of the armdataboxedge.SupportPackagesClient type.
type SupportPackagesServer struct {
	// BeginTriggerSupportPackage is the fake for method SupportPackagesClient.BeginTriggerSupportPackage
	// HTTP status codes to indicate success: http.StatusOK, http.StatusAccepted
	BeginTriggerSupportPackage func(ctx context.Context, deviceName string, resourceGroupName string, triggerSupportPackageRequest armdataboxedge.TriggerSupportPackageRequest, options *armdataboxedge.SupportPackagesClientBeginTriggerSupportPackageOptions) (resp azfake.PollerResponder[armdataboxedge.SupportPackagesClientTriggerSupportPackageResponse], errResp azfake.ErrorResponder)
}

// NewSupportPackagesServerTransport creates a new instance of SupportPackagesServerTransport with the provided implementation.
// The returned SupportPackagesServerTransport instance is connected to an instance of armdataboxedge.SupportPackagesClient via the
// azcore.ClientOptions.Transporter field in the client's constructor parameters.
func NewSupportPackagesServerTransport(srv *SupportPackagesServer) *SupportPackagesServerTransport {
	return &SupportPackagesServerTransport{
		srv:                        srv,
		beginTriggerSupportPackage: newTracker[azfake.PollerResponder[armdataboxedge.SupportPackagesClientTriggerSupportPackageResponse]](),
	}
}

// SupportPackagesServerTransport connects instances of armdataboxedge.SupportPackagesClient to instances of SupportPackagesServer.
// Don't use this type directly, use NewSupportPackagesServerTransport instead.
type SupportPackagesServerTransport struct {
	srv                        *SupportPackagesServer
	beginTriggerSupportPackage *tracker[azfake.PollerResponder[armdataboxedge.SupportPackagesClientTriggerSupportPackageResponse]]
}

// Do implements the policy.Transporter interface for SupportPackagesServerTransport.
func (s *SupportPackagesServerTransport) Do(req *http.Request) (*http.Response, error) {
	rawMethod := req.Context().Value(runtime.CtxAPINameKey{})
	method, ok := rawMethod.(string)
	if !ok {
		return nil, nonRetriableError{errors.New("unable to dispatch request, missing value for CtxAPINameKey")}
	}

	var resp *http.Response
	var err error

	switch method {
	case "SupportPackagesClient.BeginTriggerSupportPackage":
		resp, err = s.dispatchBeginTriggerSupportPackage(req)
	default:
		err = fmt.Errorf("unhandled API %s", method)
	}

	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *SupportPackagesServerTransport) dispatchBeginTriggerSupportPackage(req *http.Request) (*http.Response, error) {
	if s.srv.BeginTriggerSupportPackage == nil {
		return nil, &nonRetriableError{errors.New("fake for method BeginTriggerSupportPackage not implemented")}
	}
	beginTriggerSupportPackage := s.beginTriggerSupportPackage.get(req)
	if beginTriggerSupportPackage == nil {
		const regexStr = `/subscriptions/(?P<subscriptionId>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/resourceGroups/(?P<resourceGroupName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/providers/Microsoft\.DataBoxEdge/dataBoxEdgeDevices/(?P<deviceName>[!#&$-;=?-\[\]_a-zA-Z0-9~%@]+)/triggerSupportPackage`
		regex := regexp.MustCompile(regexStr)
		matches := regex.FindStringSubmatch(req.URL.EscapedPath())
		if matches == nil || len(matches) < 3 {
			return nil, fmt.Errorf("failed to parse path %s", req.URL.Path)
		}
		body, err := server.UnmarshalRequestAsJSON[armdataboxedge.TriggerSupportPackageRequest](req)
		if err != nil {
			return nil, err
		}
		deviceNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("deviceName")])
		if err != nil {
			return nil, err
		}
		resourceGroupNameUnescaped, err := url.PathUnescape(matches[regex.SubexpIndex("resourceGroupName")])
		if err != nil {
			return nil, err
		}
		respr, errRespr := s.srv.BeginTriggerSupportPackage(req.Context(), deviceNameUnescaped, resourceGroupNameUnescaped, body, nil)
		if respErr := server.GetError(errRespr, req); respErr != nil {
			return nil, respErr
		}
		beginTriggerSupportPackage = &respr
		s.beginTriggerSupportPackage.add(req, beginTriggerSupportPackage)
	}

	resp, err := server.PollerResponderNext(beginTriggerSupportPackage, req)
	if err != nil {
		return nil, err
	}

	if !contains([]int{http.StatusOK, http.StatusAccepted}, resp.StatusCode) {
		s.beginTriggerSupportPackage.remove(req)
		return nil, &nonRetriableError{fmt.Errorf("unexpected status code %d. acceptable values are http.StatusOK, http.StatusAccepted", resp.StatusCode)}
	}
	if !server.PollerResponderMore(beginTriggerSupportPackage) {
		s.beginTriggerSupportPackage.remove(req)
	}

	return resp, nil
}
