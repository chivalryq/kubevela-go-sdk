/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kustomize

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Alphabetical type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Alphabetical{}

// Alphabetical Alphabetical set of rules to use for alphabetical ordering of the tags.
type Alphabetical struct {
	// Order specifies the sorting order of the tags. +usage=Given the letters of the alphabet as tags, ascending order would select Z, and descending order would select A.
	Order *string `json:"order,omitempty"`
}

// NewAlphabeticalWith instantiates a new Alphabetical object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewAlphabeticalWith() *Alphabetical {
	this := Alphabetical{}
	return &this
}

// NewAlphabeticalWithDefault instantiates a new Alphabetical object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlphabeticalWithDefault() *Alphabetical {
	this := Alphabetical{}
	return &this
}

// NewAlphabetical is short for NewAlphabeticalWithDefault which instantiates a new Alphabetical object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAlphabetical() *Alphabetical {
	return NewAlphabeticalWithDefault()
}

// NewAlphabeticalEmpty instantiates a new Alphabetical object with no properties set.
// This constructor will not assign any default values to properties.
func NewAlphabeticalEmpty() *Alphabetical {
	this := Alphabetical{}
	return &this
}

// NewAlphabeticals converts a list Alphabetical pointers to objects.
// This is helpful when the SetAlphabetical requires a list of objects
func NewAlphabeticalList(ps ...*Alphabetical) []Alphabetical {
	objs := []Alphabetical{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Alphabetical
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Alphabetical) Validate() error {
	// validate all nested properties
	return nil
}

// GetOrder returns the Order field value if set, zero value otherwise.
func (o *Alphabetical) GetOrder() string {
	if o == nil || utils.IsNil(o.Order) {
		var ret string
		return ret
	}
	return *o.Order
}

// GetOrderOk returns a tuple with the Order field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Alphabetical) GetOrderOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Order) {
		return nil, false
	}
	return o.Order, true
}

// HasOrder returns a boolean if a field has been set.
func (o *Alphabetical) HasOrder() bool {
	if o != nil && !utils.IsNil(o.Order) {
		return true
	}

	return false
}

// SetOrder gets a reference to the given string and assigns it to the order field.
// Order:  Order specifies the sorting order of the tags. +usage=Given the letters of the alphabet as tags, ascending order would select Z, and descending order would select A.
func (o *Alphabetical) SetOrder(v string) *Alphabetical {
	o.Order = &v
	return o
}

func (o Alphabetical) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Alphabetical) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Order) {
		toSerialize["order"] = o.Order
	}
	return toSerialize, nil
}

type NullableAlphabetical struct {
	value *Alphabetical
	isSet bool
}

func (v *NullableAlphabetical) Get() *Alphabetical {
	return v.value
}

func (v *NullableAlphabetical) Set(val *Alphabetical) {
	v.value = val
	v.isSet = true
}

func (v *NullableAlphabetical) IsSet() bool {
	return v.isSet
}

func (v *NullableAlphabetical) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAlphabetical(val *Alphabetical) *NullableAlphabetical {
	return &NullableAlphabetical{value: val, isSet: true}
}

func (v NullableAlphabetical) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAlphabetical) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
