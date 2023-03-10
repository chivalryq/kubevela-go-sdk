/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sidecar

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ConfigMapKeyRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ConfigMapKeyRef{}

// ConfigMapKeyRef Selects a key of a config map in the pod's namespace
type ConfigMapKeyRef struct {
	// The key of the config map to select from. Must be a valid secret key
	Key *string `json:"key"`
	// The name of the config map in the pod's namespace to select from
	Name *string `json:"name"`
}

// NewConfigMapKeyRefWith instantiates a new ConfigMapKeyRef object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewConfigMapKeyRefWith(key string, name string) *ConfigMapKeyRef {
	this := ConfigMapKeyRef{}
	this.Key = &key
	this.Name = &name
	return &this
}

// NewConfigMapKeyRefWithDefault instantiates a new ConfigMapKeyRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMapKeyRefWithDefault() *ConfigMapKeyRef {
	this := ConfigMapKeyRef{}
	return &this
}

// NewConfigMapKeyRef is short for NewConfigMapKeyRefWithDefault which instantiates a new ConfigMapKeyRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConfigMapKeyRef() *ConfigMapKeyRef {
	return NewConfigMapKeyRefWithDefault()
}

// NewConfigMapKeyRefEmpty instantiates a new ConfigMapKeyRef object with no properties set.
// This constructor will not assign any default values to properties.
func NewConfigMapKeyRefEmpty() *ConfigMapKeyRef {
	this := ConfigMapKeyRef{}
	return &this
}

// NewConfigMapKeyRefs converts a list ConfigMapKeyRef pointers to objects.
// This is helpful when the SetConfigMapKeyRef requires a list of objects
func NewConfigMapKeyRefList(ps ...*ConfigMapKeyRef) []ConfigMapKeyRef {
	objs := []ConfigMapKeyRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ConfigMapKeyRef
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ConfigMapKeyRef) Validate() error {
	if o.Key == nil {
		return errors.New("Key in ConfigMapKeyRef must be set")
	}
	if o.Name == nil {
		return errors.New("Name in ConfigMapKeyRef must be set")
	}
	// validate all nested properties
	return nil
}

// GetKey returns the Key field value
func (o *ConfigMapKeyRef) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *ConfigMapKeyRef) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key, true
}

// SetKey sets field value
func (o *ConfigMapKeyRef) SetKey(v string) *ConfigMapKeyRef {
	o.Key = &v
	return o
}

// GetName returns the Name field value
func (o *ConfigMapKeyRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ConfigMapKeyRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *ConfigMapKeyRef) SetName(v string) *ConfigMapKeyRef {
	o.Name = &v
	return o
}

func (o ConfigMapKeyRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConfigMapKeyRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableConfigMapKeyRef struct {
	value *ConfigMapKeyRef
	isSet bool
}

func (v *NullableConfigMapKeyRef) Get() *ConfigMapKeyRef {
	return v.value
}

func (v *NullableConfigMapKeyRef) Set(val *ConfigMapKeyRef) {
	v.value = val
	v.isSet = true
}

func (v *NullableConfigMapKeyRef) IsSet() bool {
	return v.isSet
}

func (v *NullableConfigMapKeyRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConfigMapKeyRef(val *ConfigMapKeyRef) *NullableConfigMapKeyRef {
	return &NullableConfigMapKeyRef{value: val, isSet: true}
}

func (v NullableConfigMapKeyRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConfigMapKeyRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
