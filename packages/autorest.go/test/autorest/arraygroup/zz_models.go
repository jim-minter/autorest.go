//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package arraygroup

type Product struct {
	Integer *int32
	String  *string
}

func (p *Product) GetInteger() (rv int32) {
	if p != nil && p.Integer != nil {
		return *p.Integer
	}
	return
}

func (p *Product) GetString() (rv string) {
	if p != nil && p.String != nil {
		return *p.String
	}
	return
}
