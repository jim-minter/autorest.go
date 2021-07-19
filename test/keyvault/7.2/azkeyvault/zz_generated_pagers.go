// +build go1.13

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package azkeyvault

import (
	"context"
	"github.com/Azure/azure-sdk-for-go/sdk/azcore"
	"net/http"
	"reflect"
)

type KeyVaultClientGetCertificateIssuersPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetCertificateIssuersResponse.
	PageResponse() KeyVaultClientGetCertificateIssuersResponse
}

type keyVaultClientGetCertificateIssuersPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetCertificateIssuersResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetCertificateIssuersResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetCertificateIssuersPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetCertificateIssuersPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateIssuerListResult.NextLink == nil || len(*p.current.CertificateIssuerListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getCertificateIssuersHandleError(resp)
		return false
	}
	result, err := p.client.getCertificateIssuersHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetCertificateIssuersPager) PageResponse() KeyVaultClientGetCertificateIssuersResponse {
	return p.current
}

type KeyVaultClientGetCertificateVersionsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetCertificateVersionsResponse.
	PageResponse() KeyVaultClientGetCertificateVersionsResponse
}

type keyVaultClientGetCertificateVersionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetCertificateVersionsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetCertificateVersionsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetCertificateVersionsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetCertificateVersionsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getCertificateVersionsHandleError(resp)
		return false
	}
	result, err := p.client.getCertificateVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetCertificateVersionsPager) PageResponse() KeyVaultClientGetCertificateVersionsResponse {
	return p.current
}

type KeyVaultClientGetCertificatesPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetCertificatesResponse.
	PageResponse() KeyVaultClientGetCertificatesResponse
}

type keyVaultClientGetCertificatesPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetCertificatesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetCertificatesResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetCertificatesPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetCertificatesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.CertificateListResult.NextLink == nil || len(*p.current.CertificateListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getCertificatesHandleError(resp)
		return false
	}
	result, err := p.client.getCertificatesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetCertificatesPager) PageResponse() KeyVaultClientGetCertificatesResponse {
	return p.current
}

type KeyVaultClientGetDeletedCertificatesPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetDeletedCertificatesResponse.
	PageResponse() KeyVaultClientGetDeletedCertificatesResponse
}

type keyVaultClientGetDeletedCertificatesPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedCertificatesResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedCertificatesResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetDeletedCertificatesPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetDeletedCertificatesPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedCertificateListResult.NextLink == nil || len(*p.current.DeletedCertificateListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getDeletedCertificatesHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedCertificatesHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetDeletedCertificatesPager) PageResponse() KeyVaultClientGetDeletedCertificatesResponse {
	return p.current
}

type KeyVaultClientGetDeletedKeysPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetDeletedKeysResponse.
	PageResponse() KeyVaultClientGetDeletedKeysResponse
}

type keyVaultClientGetDeletedKeysPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedKeysResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedKeysResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetDeletedKeysPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetDeletedKeysPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedKeyListResult.NextLink == nil || len(*p.current.DeletedKeyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getDeletedKeysHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetDeletedKeysPager) PageResponse() KeyVaultClientGetDeletedKeysResponse {
	return p.current
}

type KeyVaultClientGetDeletedSasDefinitionsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetDeletedSasDefinitionsResponse.
	PageResponse() KeyVaultClientGetDeletedSasDefinitionsResponse
}

type keyVaultClientGetDeletedSasDefinitionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedSasDefinitionsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedSasDefinitionsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetDeletedSasDefinitionsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetDeletedSasDefinitionsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedSasDefinitionListResult.NextLink == nil || len(*p.current.DeletedSasDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getDeletedSasDefinitionsHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedSasDefinitionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetDeletedSasDefinitionsPager) PageResponse() KeyVaultClientGetDeletedSasDefinitionsResponse {
	return p.current
}

type KeyVaultClientGetDeletedSecretsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetDeletedSecretsResponse.
	PageResponse() KeyVaultClientGetDeletedSecretsResponse
}

type keyVaultClientGetDeletedSecretsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedSecretsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedSecretsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetDeletedSecretsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetDeletedSecretsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedSecretListResult.NextLink == nil || len(*p.current.DeletedSecretListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getDeletedSecretsHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedSecretsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetDeletedSecretsPager) PageResponse() KeyVaultClientGetDeletedSecretsResponse {
	return p.current
}

type KeyVaultClientGetDeletedStorageAccountsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetDeletedStorageAccountsResponse.
	PageResponse() KeyVaultClientGetDeletedStorageAccountsResponse
}

type keyVaultClientGetDeletedStorageAccountsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetDeletedStorageAccountsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetDeletedStorageAccountsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetDeletedStorageAccountsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetDeletedStorageAccountsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.DeletedStorageListResult.NextLink == nil || len(*p.current.DeletedStorageListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getDeletedStorageAccountsHandleError(resp)
		return false
	}
	result, err := p.client.getDeletedStorageAccountsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetDeletedStorageAccountsPager) PageResponse() KeyVaultClientGetDeletedStorageAccountsResponse {
	return p.current
}

type KeyVaultClientGetKeyVersionsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetKeyVersionsResponse.
	PageResponse() KeyVaultClientGetKeyVersionsResponse
}

type keyVaultClientGetKeyVersionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetKeyVersionsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetKeyVersionsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetKeyVersionsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetKeyVersionsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getKeyVersionsHandleError(resp)
		return false
	}
	result, err := p.client.getKeyVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetKeyVersionsPager) PageResponse() KeyVaultClientGetKeyVersionsResponse {
	return p.current
}

type KeyVaultClientGetKeysPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetKeysResponse.
	PageResponse() KeyVaultClientGetKeysResponse
}

type keyVaultClientGetKeysPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetKeysResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetKeysResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetKeysPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetKeysPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.KeyListResult.NextLink == nil || len(*p.current.KeyListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getKeysHandleError(resp)
		return false
	}
	result, err := p.client.getKeysHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetKeysPager) PageResponse() KeyVaultClientGetKeysResponse {
	return p.current
}

type KeyVaultClientGetSasDefinitionsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetSasDefinitionsResponse.
	PageResponse() KeyVaultClientGetSasDefinitionsResponse
}

type keyVaultClientGetSasDefinitionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetSasDefinitionsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetSasDefinitionsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetSasDefinitionsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetSasDefinitionsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SasDefinitionListResult.NextLink == nil || len(*p.current.SasDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getSasDefinitionsHandleError(resp)
		return false
	}
	result, err := p.client.getSasDefinitionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetSasDefinitionsPager) PageResponse() KeyVaultClientGetSasDefinitionsResponse {
	return p.current
}

type KeyVaultClientGetSecretVersionsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetSecretVersionsResponse.
	PageResponse() KeyVaultClientGetSecretVersionsResponse
}

type keyVaultClientGetSecretVersionsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetSecretVersionsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetSecretVersionsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetSecretVersionsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetSecretVersionsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getSecretVersionsHandleError(resp)
		return false
	}
	result, err := p.client.getSecretVersionsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetSecretVersionsPager) PageResponse() KeyVaultClientGetSecretVersionsResponse {
	return p.current
}

type KeyVaultClientGetSecretsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetSecretsResponse.
	PageResponse() KeyVaultClientGetSecretsResponse
}

type keyVaultClientGetSecretsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetSecretsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetSecretsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetSecretsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetSecretsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.SecretListResult.NextLink == nil || len(*p.current.SecretListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getSecretsHandleError(resp)
		return false
	}
	result, err := p.client.getSecretsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetSecretsPager) PageResponse() KeyVaultClientGetSecretsResponse {
	return p.current
}

type KeyVaultClientGetStorageAccountsPager interface {
	azcore.Pager
	// PageResponse returns the current KeyVaultClientGetStorageAccountsResponse.
	PageResponse() KeyVaultClientGetStorageAccountsResponse
}

type keyVaultClientGetStorageAccountsPager struct {
	client    *KeyVaultClient
	current   KeyVaultClientGetStorageAccountsResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, KeyVaultClientGetStorageAccountsResponse) (*azcore.Request, error)
}

func (p *keyVaultClientGetStorageAccountsPager) Err() error {
	return p.err
}

func (p *keyVaultClientGetStorageAccountsPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.StorageListResult.NextLink == nil || len(*p.current.StorageListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.getStorageAccountsHandleError(resp)
		return false
	}
	result, err := p.client.getStorageAccountsHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *keyVaultClientGetStorageAccountsPager) PageResponse() KeyVaultClientGetStorageAccountsResponse {
	return p.current
}

type RoleAssignmentsListForScopePager interface {
	azcore.Pager
	// PageResponse returns the current RoleAssignmentsListForScopeResponse.
	PageResponse() RoleAssignmentsListForScopeResponse
}

type roleAssignmentsListForScopePager struct {
	client    *RoleAssignmentsClient
	current   RoleAssignmentsListForScopeResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, RoleAssignmentsListForScopeResponse) (*azcore.Request, error)
}

func (p *roleAssignmentsListForScopePager) Err() error {
	return p.err
}

func (p *roleAssignmentsListForScopePager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleAssignmentListResult.NextLink == nil || len(*p.current.RoleAssignmentListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listForScopeHandleError(resp)
		return false
	}
	result, err := p.client.listForScopeHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleAssignmentsListForScopePager) PageResponse() RoleAssignmentsListForScopeResponse {
	return p.current
}

type RoleDefinitionsListPager interface {
	azcore.Pager
	// PageResponse returns the current RoleDefinitionsListResponse.
	PageResponse() RoleDefinitionsListResponse
}

type roleDefinitionsListPager struct {
	client    *RoleDefinitionsClient
	current   RoleDefinitionsListResponse
	err       error
	requester func(context.Context) (*azcore.Request, error)
	advancer  func(context.Context, RoleDefinitionsListResponse) (*azcore.Request, error)
}

func (p *roleDefinitionsListPager) Err() error {
	return p.err
}

func (p *roleDefinitionsListPager) NextPage(ctx context.Context) bool {
	var req *azcore.Request
	var err error
	if !reflect.ValueOf(p.current).IsZero() {
		if p.current.RoleDefinitionListResult.NextLink == nil || len(*p.current.RoleDefinitionListResult.NextLink) == 0 {
			return false
		}
		req, err = p.advancer(ctx, p.current)
	} else {
		req, err = p.requester(ctx)
	}
	if err != nil {
		p.err = err
		return false
	}
	resp, err := p.client.con.Pipeline().Do(req)
	if err != nil {
		p.err = err
		return false
	}
	if !resp.HasStatusCode(http.StatusOK) {
		p.err = p.client.listHandleError(resp)
		return false
	}
	result, err := p.client.listHandleResponse(resp)
	if err != nil {
		p.err = err
		return false
	}
	p.current = result
	return true
}

func (p *roleDefinitionsListPager) PageResponse() RoleDefinitionsListResponse {
	return p.current
}