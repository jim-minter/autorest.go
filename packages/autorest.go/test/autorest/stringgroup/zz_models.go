//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package stringgroup

type RefColorConstant struct {
	// CONSTANT; Referenced Color Constant Description.
	// Field has constant value "green-color", any specified value is ignored.
	ColorConstant *string

	// Sample string.
	Field1 *string
}

func (r *RefColorConstant) GetColorConstant() (rv *string) {
	if r != nil {
		return r.ColorConstant
	}
	return
}

func (r *RefColorConstant) GetField1() (rv string) {
	if r != nil && r.Field1 != nil {
		return *r.Field1
	}
	return
}
