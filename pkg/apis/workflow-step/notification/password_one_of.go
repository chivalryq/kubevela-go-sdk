/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the PasswordOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PasswordOneOf{}

// PasswordOneOf struct for PasswordOneOf
type PasswordOneOf struct {
	// the password content in string
	Value *string `json:"value"`
}

// NewPasswordOneOfWith instantiates a new PasswordOneOf object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPasswordOneOfWith(value string) *PasswordOneOf {
	this := PasswordOneOf{}
	this.Value = &value
	return &this
}

// NewPasswordOneOfWithDefault instantiates a new PasswordOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordOneOfWithDefault() *PasswordOneOf {
	this := PasswordOneOf{}
	return &this
}

// NewPasswordOneOf is short for NewPasswordOneOfWithDefault which instantiates a new PasswordOneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPasswordOneOf() *PasswordOneOf {
	return NewPasswordOneOfWithDefault()
}

// NewPasswordOneOfEmpty instantiates a new PasswordOneOf object with no properties set.
// This constructor will not assign any default values to properties.
func NewPasswordOneOfEmpty() *PasswordOneOf {
	this := PasswordOneOf{}
	return &this
}

// NewPasswordOneOfs converts a list PasswordOneOf pointers to objects.
// This is helpful when the SetPasswordOneOf requires a list of objects
func NewPasswordOneOfList(ps ...*PasswordOneOf) []PasswordOneOf {
	objs := []PasswordOneOf{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this PasswordOneOf
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *PasswordOneOf) Validate() error {
	if o.Value == nil {
		return errors.New("Value in PasswordOneOf must be set")
	}
	// validate all nested properties
	return nil
}

// GetValue returns the Value field value
func (o *PasswordOneOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PasswordOneOf) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *PasswordOneOf) SetValue(v string) *PasswordOneOf {
	o.Value = &v
	return o
}

func (o PasswordOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PasswordOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullablePasswordOneOf struct {
	value *PasswordOneOf
	isSet bool
}

func (v *NullablePasswordOneOf) Get() *PasswordOneOf {
	return v.value
}

func (v *NullablePasswordOneOf) Set(val *PasswordOneOf) {
	v.value = val
	v.isSet = true
}

func (v *NullablePasswordOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullablePasswordOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePasswordOneOf(val *PasswordOneOf) *NullablePasswordOneOf {
	return &NullablePasswordOneOf{value: val, isSet: true}
}

func (v NullablePasswordOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePasswordOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
