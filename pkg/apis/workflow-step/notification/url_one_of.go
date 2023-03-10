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

// checks if the UrlOneOf type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &UrlOneOf{}

// UrlOneOf struct for UrlOneOf
type UrlOneOf struct {
	// the url address content in string
	Value *string `json:"value"`
}

// NewUrlOneOfWith instantiates a new UrlOneOf object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewUrlOneOfWith(value string) *UrlOneOf {
	this := UrlOneOf{}
	this.Value = &value
	return &this
}

// NewUrlOneOfWithDefault instantiates a new UrlOneOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlOneOfWithDefault() *UrlOneOf {
	this := UrlOneOf{}
	return &this
}

// NewUrlOneOf is short for NewUrlOneOfWithDefault which instantiates a new UrlOneOf object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUrlOneOf() *UrlOneOf {
	return NewUrlOneOfWithDefault()
}

// NewUrlOneOfEmpty instantiates a new UrlOneOf object with no properties set.
// This constructor will not assign any default values to properties.
func NewUrlOneOfEmpty() *UrlOneOf {
	this := UrlOneOf{}
	return &this
}

// NewUrlOneOfs converts a list UrlOneOf pointers to objects.
// This is helpful when the SetUrlOneOf requires a list of objects
func NewUrlOneOfList(ps ...*UrlOneOf) []UrlOneOf {
	objs := []UrlOneOf{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this UrlOneOf
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *UrlOneOf) Validate() error {
	if o.Value == nil {
		return errors.New("Value in UrlOneOf must be set")
	}
	// validate all nested properties
	return nil
}

// GetValue returns the Value field value
func (o *UrlOneOf) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *UrlOneOf) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *UrlOneOf) SetValue(v string) *UrlOneOf {
	o.Value = &v
	return o
}

func (o UrlOneOf) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UrlOneOf) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullableUrlOneOf struct {
	value *UrlOneOf
	isSet bool
}

func (v *NullableUrlOneOf) Get() *UrlOneOf {
	return v.value
}

func (v *NullableUrlOneOf) Set(val *UrlOneOf) {
	v.value = val
	v.isSet = true
}

func (v *NullableUrlOneOf) IsSet() bool {
	return v.isSet
}

func (v *NullableUrlOneOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUrlOneOf(val *UrlOneOf) *NullableUrlOneOf {
	return &NullableUrlOneOf{value: val, isSet: true}
}

func (v NullableUrlOneOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUrlOneOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}