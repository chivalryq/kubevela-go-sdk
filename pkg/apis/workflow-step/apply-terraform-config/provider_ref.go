/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_terraform_config

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ProviderRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ProviderRef{}

// ProviderRef providerRef specifies the reference to Provider
type ProviderRef struct {
	Name      *string `json:"name"`
	Namespace *string `json:"namespace"`
}

// NewProviderRefWith instantiates a new ProviderRef object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewProviderRefWith(name string, namespace string) *ProviderRef {
	this := ProviderRef{}
	this.Name = &name
	this.Namespace = &namespace
	return &this
}

// NewProviderRefWithDefault instantiates a new ProviderRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderRefWithDefault() *ProviderRef {
	this := ProviderRef{}
	return &this
}

// NewProviderRef is short for NewProviderRefWithDefault which instantiates a new ProviderRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProviderRef() *ProviderRef {
	return NewProviderRefWithDefault()
}

// NewProviderRefEmpty instantiates a new ProviderRef object with no properties set.
// This constructor will not assign any default values to properties.
func NewProviderRefEmpty() *ProviderRef {
	this := ProviderRef{}
	return &this
}

// NewProviderRefs converts a list ProviderRef pointers to objects.
// This is helpful when the SetProviderRef requires a list of objects
func NewProviderRefList(ps ...*ProviderRef) []ProviderRef {
	objs := []ProviderRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ProviderRef
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ProviderRef) Validate() error {
	if o.Name == nil {
		return errors.New("Name in ProviderRef must be set")
	}
	if o.Namespace == nil {
		return errors.New("Namespace in ProviderRef must be set")
	}
	// validate all nested properties
	return nil
}

// GetName returns the Name field value
func (o *ProviderRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ProviderRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *ProviderRef) SetName(v string) *ProviderRef {
	o.Name = &v
	return o
}

// GetNamespace returns the Namespace field value
func (o *ProviderRef) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *ProviderRef) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace, true
}

// SetNamespace sets field value
func (o *ProviderRef) SetNamespace(v string) *ProviderRef {
	o.Namespace = &v
	return o
}

func (o ProviderRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProviderRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableProviderRef struct {
	value *ProviderRef
	isSet bool
}

func (v *NullableProviderRef) Get() *ProviderRef {
	return v.value
}

func (v *NullableProviderRef) Set(val *ProviderRef) {
	v.value = val
	v.isSet = true
}

func (v *NullableProviderRef) IsSet() bool {
	return v.isSet
}

func (v *NullableProviderRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProviderRef(val *ProviderRef) *NullableProviderRef {
	return &NullableProviderRef{value: val, isSet: true}
}

func (v NullableProviderRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProviderRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}