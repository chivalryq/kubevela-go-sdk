/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package k8s_update_strategy

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the RollingStrategy type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &RollingStrategy{}

// RollingStrategy Specify the parameters of rollong update strategy
type RollingStrategy struct {
	MaxSurge       *string `json:"maxSurge"`
	MaxUnavailable *string `json:"maxUnavailable"`
	Partition      *int32  `json:"partition"`
}

// NewRollingStrategyWith instantiates a new RollingStrategy object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewRollingStrategyWith(maxSurge string, maxUnavailable string, partition int32) *RollingStrategy {
	this := RollingStrategy{}
	this.MaxSurge = &maxSurge
	this.MaxUnavailable = &maxUnavailable
	this.Partition = &partition
	return &this
}

// NewRollingStrategyWithDefault instantiates a new RollingStrategy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollingStrategyWithDefault() *RollingStrategy {
	this := RollingStrategy{}
	var maxSurge string = "25%"
	this.MaxSurge = &maxSurge
	var maxUnavailable string = "25%"
	this.MaxUnavailable = &maxUnavailable
	var partition int32 = 0
	this.Partition = &partition
	return &this
}

// NewRollingStrategy is short for NewRollingStrategyWithDefault which instantiates a new RollingStrategy object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRollingStrategy() *RollingStrategy {
	return NewRollingStrategyWithDefault()
}

// NewRollingStrategyEmpty instantiates a new RollingStrategy object with no properties set.
// This constructor will not assign any default values to properties.
func NewRollingStrategyEmpty() *RollingStrategy {
	this := RollingStrategy{}
	return &this
}

// NewRollingStrategys converts a list RollingStrategy pointers to objects.
// This is helpful when the SetRollingStrategy requires a list of objects
func NewRollingStrategyList(ps ...*RollingStrategy) []RollingStrategy {
	objs := []RollingStrategy{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this RollingStrategy
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *RollingStrategy) Validate() error {
	if o.MaxSurge == nil {
		return errors.New("MaxSurge in RollingStrategy must be set")
	}
	if o.MaxUnavailable == nil {
		return errors.New("MaxUnavailable in RollingStrategy must be set")
	}
	if o.Partition == nil {
		return errors.New("Partition in RollingStrategy must be set")
	}
	// validate all nested properties
	return nil
}

// GetMaxSurge returns the MaxSurge field value
func (o *RollingStrategy) GetMaxSurge() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.MaxSurge
}

// GetMaxSurgeOk returns a tuple with the MaxSurge field value
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetMaxSurgeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxSurge, true
}

// SetMaxSurge sets field value
func (o *RollingStrategy) SetMaxSurge(v string) *RollingStrategy {
	o.MaxSurge = &v
	return o
}

// GetMaxUnavailable returns the MaxUnavailable field value
func (o *RollingStrategy) GetMaxUnavailable() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.MaxUnavailable
}

// GetMaxUnavailableOk returns a tuple with the MaxUnavailable field value
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetMaxUnavailableOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MaxUnavailable, true
}

// SetMaxUnavailable sets field value
func (o *RollingStrategy) SetMaxUnavailable(v string) *RollingStrategy {
	o.MaxUnavailable = &v
	return o
}

// GetPartition returns the Partition field value
func (o *RollingStrategy) GetPartition() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Partition
}

// GetPartitionOk returns a tuple with the Partition field value
// and a boolean to check if the value has been set.
func (o *RollingStrategy) GetPartitionOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Partition, true
}

// SetPartition sets field value
func (o *RollingStrategy) SetPartition(v int32) *RollingStrategy {
	o.Partition = &v
	return o
}

func (o RollingStrategy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RollingStrategy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["maxSurge"] = o.MaxSurge
	toSerialize["maxUnavailable"] = o.MaxUnavailable
	toSerialize["partition"] = o.Partition
	return toSerialize, nil
}

type NullableRollingStrategy struct {
	value *RollingStrategy
	isSet bool
}

func (v *NullableRollingStrategy) Get() *RollingStrategy {
	return v.value
}

func (v *NullableRollingStrategy) Set(val *RollingStrategy) {
	v.value = val
	v.isSet = true
}

func (v *NullableRollingStrategy) IsSet() bool {
	return v.isSet
}

func (v *NullableRollingStrategy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRollingStrategy(val *RollingStrategy) *NullableRollingStrategy {
	return &NullableRollingStrategy{value: val, isSet: true}
}

func (v NullableRollingStrategy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRollingStrategy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
