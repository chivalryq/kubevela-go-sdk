/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package affinity

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Required type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Required{}

// Required Specify the required during scheduling ignored during execution
type Required struct {
	// Specify a list of node selector
	NodeSelectorTerms []NodeSelectorTerm `json:"nodeSelectorTerms"`
}

// NewRequiredWith instantiates a new Required object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewRequiredWith(nodeSelectorTerms []NodeSelectorTerm) *Required {
	this := Required{}
	this.NodeSelectorTerms = nodeSelectorTerms
	return &this
}

// NewRequiredWithDefault instantiates a new Required object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequiredWithDefault() *Required {
	this := Required{}
	return &this
}

// NewRequired is short for NewRequiredWithDefault which instantiates a new Required object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequired() *Required {
	return NewRequiredWithDefault()
}

// NewRequiredEmpty instantiates a new Required object with no properties set.
// This constructor will not assign any default values to properties.
func NewRequiredEmpty() *Required {
	this := Required{}
	return &this
}

// NewRequireds converts a list Required pointers to objects.
// This is helpful when the SetRequired requires a list of objects
func NewRequiredList(ps ...*Required) []Required {
	objs := []Required{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Required
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Required) Validate() error {
	if o.NodeSelectorTerms == nil {
		return errors.New("NodeSelectorTerms in Required must be set")
	}
	// validate all nested properties
	return nil
}

// GetNodeSelectorTerms returns the NodeSelectorTerms field value
func (o *Required) GetNodeSelectorTerms() []NodeSelectorTerm {
	if o == nil {
		var ret []NodeSelectorTerm
		return ret
	}

	return o.NodeSelectorTerms
}

// GetNodeSelectorTermsOk returns a tuple with the NodeSelectorTerms field value
// and a boolean to check if the value has been set.
func (o *Required) GetNodeSelectorTermsOk() ([]NodeSelectorTerm, bool) {
	if o == nil {
		return nil, false
	}
	return o.NodeSelectorTerms, true
}

// SetNodeSelectorTerms sets field value
func (o *Required) SetNodeSelectorTerms(v []NodeSelectorTerm) *Required {
	o.NodeSelectorTerms = v
	return o
}

func (o Required) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Required) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["nodeSelectorTerms"] = o.NodeSelectorTerms
	return toSerialize, nil
}

type NullableRequired struct {
	value *Required
	isSet bool
}

func (v *NullableRequired) Get() *Required {
	return v.value
}

func (v *NullableRequired) Set(val *Required) {
	v.value = val
	v.isSet = true
}

func (v *NullableRequired) IsSet() bool {
	return v.isSet
}

func (v *NullableRequired) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequired(val *Required) *NullableRequired {
	return &NullableRequired{value: val, isSet: true}
}

func (v NullableRequired) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequired) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}