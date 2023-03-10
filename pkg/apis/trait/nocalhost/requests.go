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

// checks if the Requests type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Requests{}

// Requests struct for Requests
type Requests struct {
	Cpu    *string `json:"cpu"`
	Memory *string `json:"memory"`
}

// NewRequestsWith instantiates a new Requests object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewRequestsWith(cpu string, memory string) *Requests {
	this := Requests{}
	this.Cpu = &cpu
	this.Memory = &memory
	return &this
}

// NewRequestsWithDefault instantiates a new Requests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestsWithDefault() *Requests {
	this := Requests{}
	var cpu string = "0.5"
	this.Cpu = &cpu
	var memory string = "512Mi"
	this.Memory = &memory
	return &this
}

// NewRequests is short for NewRequestsWithDefault which instantiates a new Requests object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequests() *Requests {
	return NewRequestsWithDefault()
}

// NewRequestsEmpty instantiates a new Requests object with no properties set.
// This constructor will not assign any default values to properties.
func NewRequestsEmpty() *Requests {
	this := Requests{}
	return &this
}

// NewRequestss converts a list Requests pointers to objects.
// This is helpful when the SetRequests requires a list of objects
func NewRequestsList(ps ...*Requests) []Requests {
	objs := []Requests{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Requests
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Requests) Validate() error {
	if o.Cpu == nil {
		return errors.New("Cpu in Requests must be set")
	}
	if o.Memory == nil {
		return errors.New("Memory in Requests must be set")
	}
	// validate all nested properties
	return nil
}

// GetCpu returns the Cpu field value
func (o *Requests) GetCpu() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value
// and a boolean to check if the value has been set.
func (o *Requests) GetCpuOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Cpu, true
}

// SetCpu sets field value
func (o *Requests) SetCpu(v string) *Requests {
	o.Cpu = &v
	return o
}

// GetMemory returns the Memory field value
func (o *Requests) GetMemory() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Memory
}

// GetMemoryOk returns a tuple with the Memory field value
// and a boolean to check if the value has been set.
func (o *Requests) GetMemoryOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Memory, true
}

// SetMemory sets field value
func (o *Requests) SetMemory(v string) *Requests {
	o.Memory = &v
	return o
}

func (o Requests) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Requests) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpu"] = o.Cpu
	toSerialize["memory"] = o.Memory
	return toSerialize, nil
}

type NullableRequests struct {
	value *Requests
	isSet bool
}

func (v *NullableRequests) Get() *Requests {
	return v.value
}

func (v *NullableRequests) Set(val *Requests) {
	v.value = val
	v.isSet = true
}

func (v *NullableRequests) IsSet() bool {
	return v.isSet
}

func (v *NullableRequests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequests(val *Requests) *NullableRequests {
	return &NullableRequests{value: val, isSet: true}
}

func (v NullableRequests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
