/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package nocalhost

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Limits type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Limits{}

// Limits struct for Limits
type Limits struct {
	Cpu    *string `json:"cpu"`
	Memory *string `json:"memory"`
}

// NewLimitsWith instantiates a new Limits object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewLimitsWith(cpu string, memory string) *Limits {
	this := Limits{}
	this.Cpu = &cpu
	this.Memory = &memory
	return &this
}

// NewLimitsWithDefault instantiates a new Limits object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimitsWithDefault() *Limits {
	this := Limits{}
	var cpu string = "2"
	this.Cpu = &cpu
	var memory string = "2Gi"
	this.Memory = &memory
	return &this
}

// NewLimits is short for NewLimitsWithDefault which instantiates a new Limits object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLimits() *Limits {
	return NewLimitsWithDefault()
}

// NewLimitsEmpty instantiates a new Limits object with no properties set.
// This constructor will not assign any default values to properties.
func NewLimitsEmpty() *Limits {
	this := Limits{}
	return &this
}

// NewLimitss converts a list Limits pointers to objects.
// This is helpful when the SetLimits requires a list of objects
func NewLimitsList(ps ...*Limits) []Limits {
	objs := []Limits{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Limits
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Limits) Validate() error {
	if o.Cpu == nil {
		return errors.New("Cpu in Limits must be set")
	}
	if o.Memory == nil {
		return errors.New("Memory in Limits must be set")
	}
	// validate all nested properties
	return nil
}

// GetCpu returns the Cpu field value
func (o *Limits) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Limits) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu, true
}

// SetCpu sets field value
func (o *Limits) SetCpu(v string) *Limits {
	o.Cpu = &v
	return o
}

// GetMemory returns the Memory field value
func (o *Limits) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Limits) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory, true
}

// SetMemory sets field value
func (o *Limits) SetMemory(v string) *Limits {
	o.Memory = &v
	return o
}

func (o Limits) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Limits) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	return toSerialize, nil
}

type NullableLimits struct {
	value *Limits
	isSet bool
}

func (v *NullableLimits) Get() *Limits {
	return v.value
}

func (v *NullableLimits) Set(val *Limits) {
	v.value = val
	v.isSet = true
}

func (v *NullableLimits) IsSet() bool {
	return v.isSet
}

func (v *NullableLimits) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLimits(val *Limits) *NullableLimits {
	return &NullableLimits{value: val, isSet: true}
}

func (v NullableLimits) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLimits) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
