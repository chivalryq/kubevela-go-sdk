/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the MountToEnv type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MountToEnv{}

// MountToEnv struct for MountToEnv
type MountToEnv struct {
	ConfigMapKey *string `json:"configMapKey"`
	EnvName      *string `json:"envName"`
}

// NewMountToEnvWith instantiates a new MountToEnv object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewMountToEnvWith(configMapKey string, envName string) *MountToEnv {
	this := MountToEnv{}
	this.ConfigMapKey = &configMapKey
	this.EnvName = &envName
	return &this
}

// NewMountToEnvWithDefault instantiates a new MountToEnv object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountToEnvWithDefault() *MountToEnv {
	this := MountToEnv{}
	return &this
}

// NewMountToEnv is short for NewMountToEnvWithDefault which instantiates a new MountToEnv object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMountToEnv() *MountToEnv {
	return NewMountToEnvWithDefault()
}

// NewMountToEnvEmpty instantiates a new MountToEnv object with no properties set.
// This constructor will not assign any default values to properties.
func NewMountToEnvEmpty() *MountToEnv {
	this := MountToEnv{}
	return &this
}

// NewMountToEnvs converts a list MountToEnv pointers to objects.
// This is helpful when the SetMountToEnv requires a list of objects
func NewMountToEnvList(ps ...*MountToEnv) []MountToEnv {
	objs := []MountToEnv{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this MountToEnv
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *MountToEnv) Validate() error {
	if o.ConfigMapKey == nil {
		return errors.New("ConfigMapKey in MountToEnv must be set")
	}
	if o.EnvName == nil {
		return errors.New("EnvName in MountToEnv must be set")
	}
	// validate all nested properties
	return nil
}

// GetConfigMapKey returns the ConfigMapKey field value
func (o *MountToEnv) GetConfigMapKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.ConfigMapKey
}

// GetConfigMapKeyOk returns a tuple with the ConfigMapKey field value
// and a boolean to check if the value has been set.
func (o *MountToEnv) GetConfigMapKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ConfigMapKey, true
}

// SetConfigMapKey sets field value
func (o *MountToEnv) SetConfigMapKey(v string) *MountToEnv {
	o.ConfigMapKey = &v
	return o
}

// GetEnvName returns the EnvName field value
func (o *MountToEnv) GetEnvName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.EnvName
}

// GetEnvNameOk returns a tuple with the EnvName field value
// and a boolean to check if the value has been set.
func (o *MountToEnv) GetEnvNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.EnvName, true
}

// SetEnvName sets field value
func (o *MountToEnv) SetEnvName(v string) *MountToEnv {
	o.EnvName = &v
	return o
}

func (o MountToEnv) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MountToEnv) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["configMapKey"] = o.ConfigMapKey
	toSerialize["envName"] = o.EnvName
	return toSerialize, nil
}

type NullableMountToEnv struct {
	value *MountToEnv
	isSet bool
}

func (v *NullableMountToEnv) Get() *MountToEnv {
	return v.value
}

func (v *NullableMountToEnv) Set(val *MountToEnv) {
	v.value = val
	v.isSet = true
}

func (v *NullableMountToEnv) IsSet() bool {
	return v.isSet
}

func (v *NullableMountToEnv) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMountToEnv(val *MountToEnv) *NullableMountToEnv {
	return &NullableMountToEnv{value: val, isSet: true}
}

func (v NullableMountToEnv) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMountToEnv) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}