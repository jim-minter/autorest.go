//go:build go1.18
// +build go1.18

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See License.txt in the project root for license information.
// Code generated by Microsoft (R) AutoRest Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

package complexgroup

import "time"

// ArrayClientGetEmptyOptions contains the optional parameters for the ArrayClient.GetEmpty method.
type ArrayClientGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// ArrayClientGetNotProvidedOptions contains the optional parameters for the ArrayClient.GetNotProvided method.
type ArrayClientGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// ArrayClientGetValidOptions contains the optional parameters for the ArrayClient.GetValid method.
type ArrayClientGetValidOptions struct {
	// placeholder for future optional parameters
}

// ArrayClientPutEmptyOptions contains the optional parameters for the ArrayClient.PutEmpty method.
type ArrayClientPutEmptyOptions struct {
	// placeholder for future optional parameters
}

// ArrayClientPutValidOptions contains the optional parameters for the ArrayClient.PutValid method.
type ArrayClientPutValidOptions struct {
	// placeholder for future optional parameters
}

type ArrayWrapper struct {
	Array []*string `json:"array,omitempty"`
}

type Basic struct {
	Color *CMYKColors `json:"color,omitempty"`

	// Basic Id
	ID *int32 `json:"id,omitempty"`

	// Name property with a very long description that does not fit on a single line and a line break.
	Name *string `json:"name,omitempty"`
}

// BasicClientGetEmptyOptions contains the optional parameters for the BasicClient.GetEmpty method.
type BasicClientGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// BasicClientGetInvalidOptions contains the optional parameters for the BasicClient.GetInvalid method.
type BasicClientGetInvalidOptions struct {
	// placeholder for future optional parameters
}

// BasicClientGetNotProvidedOptions contains the optional parameters for the BasicClient.GetNotProvided method.
type BasicClientGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// BasicClientGetNullOptions contains the optional parameters for the BasicClient.GetNull method.
type BasicClientGetNullOptions struct {
	// placeholder for future optional parameters
}

// BasicClientGetValidOptions contains the optional parameters for the BasicClient.GetValid method.
type BasicClientGetValidOptions struct {
	// placeholder for future optional parameters
}

// BasicClientPutValidOptions contains the optional parameters for the BasicClient.PutValid method.
type BasicClientPutValidOptions struct {
	// placeholder for future optional parameters
}

type BooleanWrapper struct {
	FieldFalse *bool `json:"field_false,omitempty"`
	FieldTrue  *bool `json:"field_true,omitempty"`
}

type ByteWrapper struct {
	Field []byte `json:"field,omitempty"`
}

type Cat struct {
	Color *string `json:"color,omitempty"`
	Hates []*Dog  `json:"hates,omitempty"`
	ID    *int32  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type Cookiecuttershark struct {
	// REQUIRED
	Birthday *time.Time `json:"birthday,omitempty"`

	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length   *float32             `json:"length,omitempty"`
	Age      *int32               `json:"age,omitempty"`
	Siblings []FishClassification `json:"siblings,omitempty"`
	Species  *string              `json:"species,omitempty"`
}

type DateWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Leap  *time.Time `json:"leap,omitempty"`
}

type DatetimeWrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

type Datetimerfc1123Wrapper struct {
	Field *time.Time `json:"field,omitempty"`
	Now   *time.Time `json:"now,omitempty"`
}

// DictionaryClientGetEmptyOptions contains the optional parameters for the DictionaryClient.GetEmpty method.
type DictionaryClientGetEmptyOptions struct {
	// placeholder for future optional parameters
}

// DictionaryClientGetNotProvidedOptions contains the optional parameters for the DictionaryClient.GetNotProvided method.
type DictionaryClientGetNotProvidedOptions struct {
	// placeholder for future optional parameters
}

// DictionaryClientGetNullOptions contains the optional parameters for the DictionaryClient.GetNull method.
type DictionaryClientGetNullOptions struct {
	// placeholder for future optional parameters
}

// DictionaryClientGetValidOptions contains the optional parameters for the DictionaryClient.GetValid method.
type DictionaryClientGetValidOptions struct {
	// placeholder for future optional parameters
}

// DictionaryClientPutEmptyOptions contains the optional parameters for the DictionaryClient.PutEmpty method.
type DictionaryClientPutEmptyOptions struct {
	// placeholder for future optional parameters
}

// DictionaryClientPutValidOptions contains the optional parameters for the DictionaryClient.PutValid method.
type DictionaryClientPutValidOptions struct {
	// placeholder for future optional parameters
}

type DictionaryWrapper struct {
	// Dictionary of
	DefaultProgram map[string]*string `json:"defaultProgram,omitempty"`
}

type Dog struct {
	Food *string `json:"food,omitempty"`
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// DotFishClassification provides polymorphic access to related types.
// Call the interface's GetDotFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *DotFish, *DotSalmon
type DotFishClassification interface {
	// GetDotFish returns the DotFish content of the underlying type.
	GetDotFish() *DotFish
}

type DotFish struct {
	// REQUIRED
	FishType *string `json:"fish.type,omitempty"`
	Species  *string `json:"species,omitempty"`
}

type DotFishMarket struct {
	Fishes       []DotFishClassification `json:"fishes,omitempty"`
	Salmons      []*DotSalmon            `json:"salmons,omitempty"`
	SampleFish   DotFishClassification   `json:"sampleFish,omitempty"`
	SampleSalmon *DotSalmon              `json:"sampleSalmon,omitempty"`
}

type DotSalmon struct {
	// REQUIRED
	FishType *string `json:"fish.type,omitempty"`
	Iswild   *bool   `json:"iswild,omitempty"`
	Location *string `json:"location,omitempty"`
	Species  *string `json:"species,omitempty"`
}

type DoubleWrapper struct {
	Field1                                                                          *float64 `json:"field1,omitempty"`
	Field56ZerosAfterTheDotAndNegativeZeroBeforeDotAndThisIsALongFieldNameOnPurpose *float64 `json:"field_56_zeros_after_the_dot_and_negative_zero_before_dot_and_this_is_a_long_field_name_on_purpose,omitempty"`
}

type DurationWrapper struct {
	Field *string `json:"field,omitempty"`
}

type Error struct {
	Message *string `json:"message,omitempty"`
	Status  *int32  `json:"status,omitempty"`
}

// FishClassification provides polymorphic access to related types.
// Call the interface's GetFish() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Cookiecuttershark, *Fish, *Goblinshark, *Salmon, *Sawshark, *Shark, *SmartSalmon
type FishClassification interface {
	// GetFish returns the Fish content of the underlying type.
	GetFish() *Fish
}

type Fish struct {
	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length   *float32             `json:"length,omitempty"`
	Siblings []FishClassification `json:"siblings,omitempty"`
	Species  *string              `json:"species,omitempty"`
}

// FlattencomplexClientGetValidOptions contains the optional parameters for the FlattencomplexClient.GetValid method.
type FlattencomplexClientGetValidOptions struct {
	// placeholder for future optional parameters
}

type FloatWrapper struct {
	Field1 *float32 `json:"field1,omitempty"`
	Field2 *float32 `json:"field2,omitempty"`
}

type Goblinshark struct {
	// REQUIRED
	Birthday *time.Time `json:"birthday,omitempty"`

	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length *float32 `json:"length,omitempty"`
	Age    *int32   `json:"age,omitempty"`

	// Colors possible
	Color    *GoblinSharkColor    `json:"color,omitempty"`
	Jawsize  *int32               `json:"jawsize,omitempty"`
	Siblings []FishClassification `json:"siblings,omitempty"`
	Species  *string              `json:"species,omitempty"`
}

// InheritanceClientGetValidOptions contains the optional parameters for the InheritanceClient.GetValid method.
type InheritanceClientGetValidOptions struct {
	// placeholder for future optional parameters
}

// InheritanceClientPutValidOptions contains the optional parameters for the InheritanceClient.PutValid method.
type InheritanceClientPutValidOptions struct {
	// placeholder for future optional parameters
}

type IntWrapper struct {
	Field1 *int32 `json:"field1,omitempty"`
	Field2 *int32 `json:"field2,omitempty"`
}

type LongWrapper struct {
	Field1 *int64 `json:"field1,omitempty"`
	Field2 *int64 `json:"field2,omitempty"`
}

type MyBaseHelperType struct {
	PropBH1 *string `json:"propBH1,omitempty"`
}

// MyBaseTypeClassification provides polymorphic access to related types.
// Call the interface's GetMyBaseType() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *MyBaseType, *MyDerivedType
type MyBaseTypeClassification interface {
	// GetMyBaseType returns the MyBaseType content of the underlying type.
	GetMyBaseType() *MyBaseType
}

type MyBaseType struct {
	// REQUIRED
	Kind   *MyKind           `json:"kind,omitempty"`
	Helper *MyBaseHelperType `json:"helper,omitempty"`
	PropB1 *string           `json:"propB1,omitempty"`
}

type MyDerivedType struct {
	// REQUIRED
	Kind   *MyKind           `json:"kind,omitempty"`
	Helper *MyBaseHelperType `json:"helper,omitempty"`
	PropB1 *string           `json:"propB1,omitempty"`
	PropD1 *string           `json:"propD1,omitempty"`
}

type Pet struct {
	ID   *int32  `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// PolymorphicrecursiveClientGetValidOptions contains the optional parameters for the PolymorphicrecursiveClient.GetValid
// method.
type PolymorphicrecursiveClientGetValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphicrecursiveClientPutValidOptions contains the optional parameters for the PolymorphicrecursiveClient.PutValid
// method.
type PolymorphicrecursiveClientPutValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientGetComplicatedOptions contains the optional parameters for the PolymorphismClient.GetComplicated method.
type PolymorphismClientGetComplicatedOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientGetComposedWithDiscriminatorOptions contains the optional parameters for the PolymorphismClient.GetComposedWithDiscriminator
// method.
type PolymorphismClientGetComposedWithDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientGetComposedWithoutDiscriminatorOptions contains the optional parameters for the PolymorphismClient.GetComposedWithoutDiscriminator
// method.
type PolymorphismClientGetComposedWithoutDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientGetDotSyntaxOptions contains the optional parameters for the PolymorphismClient.GetDotSyntax method.
type PolymorphismClientGetDotSyntaxOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientGetValidOptions contains the optional parameters for the PolymorphismClient.GetValid method.
type PolymorphismClientGetValidOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientPutComplicatedOptions contains the optional parameters for the PolymorphismClient.PutComplicated method.
type PolymorphismClientPutComplicatedOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientPutMissingDiscriminatorOptions contains the optional parameters for the PolymorphismClient.PutMissingDiscriminator
// method.
type PolymorphismClientPutMissingDiscriminatorOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientPutValidMissingRequiredOptions contains the optional parameters for the PolymorphismClient.PutValidMissingRequired
// method.
type PolymorphismClientPutValidMissingRequiredOptions struct {
	// placeholder for future optional parameters
}

// PolymorphismClientPutValidOptions contains the optional parameters for the PolymorphismClient.PutValid method.
type PolymorphismClientPutValidOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetBoolOptions contains the optional parameters for the PrimitiveClient.GetBool method.
type PrimitiveClientGetBoolOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetByteOptions contains the optional parameters for the PrimitiveClient.GetByte method.
type PrimitiveClientGetByteOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetDateOptions contains the optional parameters for the PrimitiveClient.GetDate method.
type PrimitiveClientGetDateOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetDateTimeOptions contains the optional parameters for the PrimitiveClient.GetDateTime method.
type PrimitiveClientGetDateTimeOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetDateTimeRFC1123Options contains the optional parameters for the PrimitiveClient.GetDateTimeRFC1123 method.
type PrimitiveClientGetDateTimeRFC1123Options struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetDoubleOptions contains the optional parameters for the PrimitiveClient.GetDouble method.
type PrimitiveClientGetDoubleOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetDurationOptions contains the optional parameters for the PrimitiveClient.GetDuration method.
type PrimitiveClientGetDurationOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetFloatOptions contains the optional parameters for the PrimitiveClient.GetFloat method.
type PrimitiveClientGetFloatOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetIntOptions contains the optional parameters for the PrimitiveClient.GetInt method.
type PrimitiveClientGetIntOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetLongOptions contains the optional parameters for the PrimitiveClient.GetLong method.
type PrimitiveClientGetLongOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientGetStringOptions contains the optional parameters for the PrimitiveClient.GetString method.
type PrimitiveClientGetStringOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutBoolOptions contains the optional parameters for the PrimitiveClient.PutBool method.
type PrimitiveClientPutBoolOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutByteOptions contains the optional parameters for the PrimitiveClient.PutByte method.
type PrimitiveClientPutByteOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutDateOptions contains the optional parameters for the PrimitiveClient.PutDate method.
type PrimitiveClientPutDateOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutDateTimeOptions contains the optional parameters for the PrimitiveClient.PutDateTime method.
type PrimitiveClientPutDateTimeOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutDateTimeRFC1123Options contains the optional parameters for the PrimitiveClient.PutDateTimeRFC1123 method.
type PrimitiveClientPutDateTimeRFC1123Options struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutDoubleOptions contains the optional parameters for the PrimitiveClient.PutDouble method.
type PrimitiveClientPutDoubleOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutDurationOptions contains the optional parameters for the PrimitiveClient.PutDuration method.
type PrimitiveClientPutDurationOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutFloatOptions contains the optional parameters for the PrimitiveClient.PutFloat method.
type PrimitiveClientPutFloatOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutIntOptions contains the optional parameters for the PrimitiveClient.PutInt method.
type PrimitiveClientPutIntOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutLongOptions contains the optional parameters for the PrimitiveClient.PutLong method.
type PrimitiveClientPutLongOptions struct {
	// placeholder for future optional parameters
}

// PrimitiveClientPutStringOptions contains the optional parameters for the PrimitiveClient.PutString method.
type PrimitiveClientPutStringOptions struct {
	// placeholder for future optional parameters
}

type ReadonlyObj struct {
	Size *int32 `json:"size,omitempty"`

	// READ-ONLY
	ID *string `json:"id,omitempty" azure:"ro"`
}

// ReadonlypropertyClientGetValidOptions contains the optional parameters for the ReadonlypropertyClient.GetValid method.
type ReadonlypropertyClientGetValidOptions struct {
	// placeholder for future optional parameters
}

// ReadonlypropertyClientPutValidOptions contains the optional parameters for the ReadonlypropertyClient.PutValid method.
type ReadonlypropertyClientPutValidOptions struct {
	// placeholder for future optional parameters
}

// SalmonClassification provides polymorphic access to related types.
// Call the interface's GetSalmon() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Salmon, *SmartSalmon
type SalmonClassification interface {
	FishClassification
	// GetSalmon returns the Salmon content of the underlying type.
	GetSalmon() *Salmon
}

type Salmon struct {
	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length   *float32             `json:"length,omitempty"`
	Iswild   *bool                `json:"iswild,omitempty"`
	Location *string              `json:"location,omitempty"`
	Siblings []FishClassification `json:"siblings,omitempty"`
	Species  *string              `json:"species,omitempty"`
}

type Sawshark struct {
	// REQUIRED
	Birthday *time.Time `json:"birthday,omitempty"`

	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length   *float32             `json:"length,omitempty"`
	Age      *int32               `json:"age,omitempty"`
	Picture  []byte               `json:"picture,omitempty"`
	Siblings []FishClassification `json:"siblings,omitempty"`
	Species  *string              `json:"species,omitempty"`
}

// SharkClassification provides polymorphic access to related types.
// Call the interface's GetShark() method to access the common type.
// Use a type switch to determine the concrete type.  The possible types are:
// - *Cookiecuttershark, *Goblinshark, *Sawshark, *Shark
type SharkClassification interface {
	FishClassification
	// GetShark returns the Shark content of the underlying type.
	GetShark() *Shark
}

type Shark struct {
	// REQUIRED
	Birthday *time.Time `json:"birthday,omitempty"`

	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length   *float32             `json:"length,omitempty"`
	Age      *int32               `json:"age,omitempty"`
	Siblings []FishClassification `json:"siblings,omitempty"`
	Species  *string              `json:"species,omitempty"`
}

type Siamese struct {
	Breed *string `json:"breed,omitempty"`
	Color *string `json:"color,omitempty"`
	Hates []*Dog  `json:"hates,omitempty"`
	ID    *int32  `json:"id,omitempty"`
	Name  *string `json:"name,omitempty"`
}

type SmartSalmon struct {
	// REQUIRED
	Fishtype *string `json:"fishtype,omitempty"`

	// REQUIRED
	Length *float32 `json:"length,omitempty"`

	// OPTIONAL; Contains additional key/value pairs not defined in the schema.
	AdditionalProperties map[string]interface{}
	CollegeDegree        *string              `json:"college_degree,omitempty"`
	Iswild               *bool                `json:"iswild,omitempty"`
	Location             *string              `json:"location,omitempty"`
	Siblings             []FishClassification `json:"siblings,omitempty"`
	Species              *string              `json:"species,omitempty"`
}

type StringWrapper struct {
	Empty *string `json:"empty,omitempty"`
	Field *string `json:"field,omitempty"`
	Null  *string `json:"null,omitempty"`
}
