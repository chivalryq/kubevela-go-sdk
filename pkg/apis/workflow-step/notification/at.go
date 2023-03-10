/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the At type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &At{}

// At struct for At
type At struct {
	AtMobiles []string `json:"atMobiles,omitempty"`
	IsAtAll   *bool    `json:"isAtAll,omitempty"`
}

// NewAtWith instantiates a new At object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewAtWith() *At {
	this := At{}
	return &this
}

// NewAtWithDefault instantiates a new At object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAtWithDefault() *At {
	this := At{}
	return &this
}

// NewAt is short for NewAtWithDefault which instantiates a new At object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAt() *At {
	return NewAtWithDefault()
}

// NewAtEmpty instantiates a new At object with no properties set.
// This constructor will not assign any default values to properties.
func NewAtEmpty() *At {
	this := At{}
	return &this
}

// NewAts converts a list At pointers to objects.
// This is helpful when the SetAt requires a list of objects
func NewAtList(ps ...*At) []At {
	objs := []At{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this At
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *At) Validate() error {
	// validate all nested properties
	return nil
}

// GetAtMobiles returns the AtMobiles field value if set, zero value otherwise.
func (o *At) GetAtMobiles() []string {
	if o == nil || utils.IsNil(o.AtMobiles) {
		var ret []string
		return ret
	}
	return o.AtMobiles
}

// GetAtMobilesOk returns a tuple with the AtMobiles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *At) GetAtMobilesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.AtMobiles) {
		return nil, false
	}
	return o.AtMobiles, true
}

// HasAtMobiles returns a boolean if a field has been set.
func (o *At) HasAtMobiles() bool {
	if o != nil && !utils.IsNil(o.AtMobiles) {
		return true
	}

	return false
}

// SetAtMobiles gets a reference to the given []string and assigns it to the atMobiles field.
// AtMobiles:
func (o *At) SetAtMobiles(v []string) *At {
	o.AtMobiles = v
	return o
}

// GetIsAtAll returns the IsAtAll field value if set, zero value otherwise.
func (o *At) GetIsAtAll() bool {
	if o == nil || utils.IsNil(o.IsAtAll) {
		var ret bool
		return ret
	}
	return *o.IsAtAll
}

// GetIsAtAllOk returns a tuple with the IsAtAll field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *At) GetIsAtAllOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.IsAtAll) {
		return nil, false
	}
	return o.IsAtAll, true
}

// HasIsAtAll returns a boolean if a field has been set.
func (o *At) HasIsAtAll() bool {
	if o != nil && !utils.IsNil(o.IsAtAll) {
		return true
	}

	return false
}

// SetIsAtAll gets a reference to the given bool and assigns it to the isAtAll field.
// IsAtAll:
func (o *At) SetIsAtAll(v bool) *At {
	o.IsAtAll = &v
	return o
}

func (o At) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o At) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.AtMobiles) {
		toSerialize["atMobiles"] = o.AtMobiles
	}
	if !utils.IsNil(o.IsAtAll) {
		toSerialize["isAtAll"] = o.IsAtAll
	}
	return toSerialize, nil
}

type NullableAt struct {
	value *At
	isSet bool
}

func (v *NullableAt) Get() *At {
	return v.value
}

func (v *NullableAt) Set(val *At) {
	v.value = val
	v.isSet = true
}

func (v *NullableAt) IsSet() bool {
	return v.isSet
}

func (v *NullableAt) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAt(val *At) *NullableAt {
	return &NullableAt{value: val, isSet: true}
}

func (v NullableAt) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAt) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
