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

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Resources type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Resources{}

// Resources struct for Resources
type Resources struct {
	Limits   *Limits   `json:"limits,omitempty"`
	Requests *Requests `json:"requests"`
}

// NewResourcesWith instantiates a new Resources object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewResourcesWith(requests Requests) *Resources {
	this := Resources{}
	this.Requests = &requests
	return &this
}

// NewResourcesWithDefault instantiates a new Resources object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourcesWithDefault() *Resources {
	this := Resources{}
	return &this
}

// NewResources is short for NewResourcesWithDefault which instantiates a new Resources object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResources() *Resources {
	return NewResourcesWithDefault()
}

// NewResourcesEmpty instantiates a new Resources object with no properties set.
// This constructor will not assign any default values to properties.
func NewResourcesEmpty() *Resources {
	this := Resources{}
	return &this
}

// NewResourcess converts a list Resources pointers to objects.
// This is helpful when the SetResources requires a list of objects
func NewResourcesList(ps ...*Resources) []Resources {
	objs := []Resources{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Resources
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Resources) Validate() error {
	if o.Requests == nil {
		return errors.New("Requests in Resources must be set")
	}
	// validate all nested properties
	if o.Limits != nil {
		if err := o.Limits.Validate(); err != nil {
			return err
		}
	}
	if o.Requests != nil {
		if err := o.Requests.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *Resources) GetLimits() Limits {
	if o == nil || utils.IsNil(o.Limits) {
		var ret Limits
		return ret
	}
	return *o.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Resources) GetLimitsOk() (*Limits, bool) {
	if o == nil || utils.IsNil(o.Limits) {
		return nil, false
	}
	return o.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *Resources) HasLimits() bool {
	if o != nil && !utils.IsNil(o.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given Limits and assigns it to the limits field.
// Limits:
func (o *Resources) SetLimits(v Limits) *Resources {
	o.Limits = &v
	return o
}

// GetRequests returns the Requests field value
func (o *Resources) GetRequests() Requests {
	if o == nil {
		var ret Requests
		return ret
	}

	return *o.Requests
}

// GetRequestsOk returns a tuple with the Requests field value
// and a boolean to check if the value has been set.
func (o *Resources) GetRequestsOk() (*Requests, bool) {
	if o == nil {
		return nil, false
	}
	return o.Requests, true
}

// SetRequests sets field value
func (o *Resources) SetRequests(v Requests) *Resources {
	o.Requests = &v
	return o
}

func (o Resources) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Resources) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	toSerialize["requests"] = o.Requests
	return toSerialize, nil
}

type NullableResources struct {
	value *Resources
	isSet bool
}

func (v *NullableResources) Get() *Resources {
	return v.value
}

func (v *NullableResources) Set(val *Resources) {
	v.value = val
	v.isSet = true
}

func (v *NullableResources) IsSet() bool {
	return v.isSet
}

func (v *NullableResources) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResources(val *Resources) *NullableResources {
	return &NullableResources{value: val, isSet: true}
}

func (v NullableResources) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResources) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
