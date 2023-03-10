/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hpa

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the PodCustomMetrics type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &PodCustomMetrics{}

// PodCustomMetrics struct for PodCustomMetrics
type PodCustomMetrics struct {
	// Specify name of custom metrics
	Name *string `json:"name"`
	// Specify target value of custom metrics
	Value *string `json:"value"`
}

// NewPodCustomMetricsWith instantiates a new PodCustomMetrics object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewPodCustomMetricsWith(name string, value string) *PodCustomMetrics {
	this := PodCustomMetrics{}
	this.Name = &name
	this.Value = &value
	return &this
}

// NewPodCustomMetricsWithDefault instantiates a new PodCustomMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodCustomMetricsWithDefault() *PodCustomMetrics {
	this := PodCustomMetrics{}
	return &this
}

// NewPodCustomMetrics is short for NewPodCustomMetricsWithDefault which instantiates a new PodCustomMetrics object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPodCustomMetrics() *PodCustomMetrics {
	return NewPodCustomMetricsWithDefault()
}

// NewPodCustomMetricsEmpty instantiates a new PodCustomMetrics object with no properties set.
// This constructor will not assign any default values to properties.
func NewPodCustomMetricsEmpty() *PodCustomMetrics {
	this := PodCustomMetrics{}
	return &this
}

// NewPodCustomMetricss converts a list PodCustomMetrics pointers to objects.
// This is helpful when the SetPodCustomMetrics requires a list of objects
func NewPodCustomMetricsList(ps ...*PodCustomMetrics) []PodCustomMetrics {
	objs := []PodCustomMetrics{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this PodCustomMetrics
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *PodCustomMetrics) Validate() error {
	if o.Name == nil {
		return errors.New("Name in PodCustomMetrics must be set")
	}
	if o.Value == nil {
		return errors.New("Value in PodCustomMetrics must be set")
	}
	// validate all nested properties
	return nil
}

// GetName returns the Name field value
func (o *PodCustomMetrics) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PodCustomMetrics) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *PodCustomMetrics) SetName(v string) *PodCustomMetrics {
	o.Name = &v
	return o
}

// GetValue returns the Value field value
func (o *PodCustomMetrics) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PodCustomMetrics) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Value, true
}

// SetValue sets field value
func (o *PodCustomMetrics) SetValue(v string) *PodCustomMetrics {
	o.Value = &v
	return o
}

func (o PodCustomMetrics) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PodCustomMetrics) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

type NullablePodCustomMetrics struct {
	value *PodCustomMetrics
	isSet bool
}

func (v *NullablePodCustomMetrics) Get() *PodCustomMetrics {
	return v.value
}

func (v *NullablePodCustomMetrics) Set(val *PodCustomMetrics) {
	v.value = val
	v.isSet = true
}

func (v *NullablePodCustomMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullablePodCustomMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePodCustomMetrics(val *PodCustomMetrics) *NullablePodCustomMetrics {
	return &NullablePodCustomMetrics{value: val, isSet: true}
}

func (v NullablePodCustomMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePodCustomMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}