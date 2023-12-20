//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator. DO NOT EDIT.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package optionalgroup

type ArrayOptionalWrapper struct {
	Value []*string
}

func (a *ArrayOptionalWrapper) GetValue() (rv []*string) {
	if a != nil {
		return a.Value
	}
	return
}

type ArrayWrapper struct {
	// REQUIRED
	Value []*string
}

func (a *ArrayWrapper) GetValue() (rv []*string) {
	if a != nil {
		return a.Value
	}
	return
}

type ClassOptionalWrapper struct {
	Value *Product
}

func (c *ClassOptionalWrapper) GetValue() (rv *Product) {
	if c != nil {
		return c.Value
	}
	return
}

type ClassWrapper struct {
	// REQUIRED
	Value *Product
}

func (c *ClassWrapper) GetValue() (rv *Product) {
	if c != nil {
		return c.Value
	}
	return
}

type IntOptionalWrapper struct {
	Value *int32
}

func (i *IntOptionalWrapper) GetValue() (rv int32) {
	if i != nil && i.Value != nil {
		return *i.Value
	}
	return
}

type IntWrapper struct {
	// REQUIRED
	Value *int32
}

func (i *IntWrapper) GetValue() (rv int32) {
	if i != nil && i.Value != nil {
		return *i.Value
	}
	return
}

type Product struct {
	// REQUIRED
	ID   *int32
	Name *string
}

func (p *Product) GetID() (rv int32) {
	if p != nil && p.ID != nil {
		return *p.ID
	}
	return
}

func (p *Product) GetName() (rv string) {
	if p != nil && p.Name != nil {
		return *p.Name
	}
	return
}

type StringOptionalWrapper struct {
	Value *string
}

func (s *StringOptionalWrapper) GetValue() (rv string) {
	if s != nil && s.Value != nil {
		return *s.Value
	}
	return
}

type StringWrapper struct {
	// REQUIRED
	Value *string
}

func (s *StringWrapper) GetValue() (rv string) {
	if s != nil && s.Value != nil {
		return *s.Value
	}
	return
}
