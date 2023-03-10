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

// checks if the WriteConnectionSecretToRef type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WriteConnectionSecretToRef{}

// WriteConnectionSecretToRef this specifies the namespace and name of a secret to which any connection details for this managed resource should be written.
type WriteConnectionSecretToRef struct {
	Name      *string `json:"name"`
	Namespace *string `json:"namespace"`
}

// NewWriteConnectionSecretToRefWith instantiates a new WriteConnectionSecretToRef object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewWriteConnectionSecretToRefWith(name string, namespace string) *WriteConnectionSecretToRef {
	this := WriteConnectionSecretToRef{}
	this.Name = &name
	this.Namespace = &namespace
	return &this
}

// NewWriteConnectionSecretToRefWithDefault instantiates a new WriteConnectionSecretToRef object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteConnectionSecretToRefWithDefault() *WriteConnectionSecretToRef {
	this := WriteConnectionSecretToRef{}
	return &this
}

// NewWriteConnectionSecretToRef is short for NewWriteConnectionSecretToRefWithDefault which instantiates a new WriteConnectionSecretToRef object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWriteConnectionSecretToRef() *WriteConnectionSecretToRef {
	return NewWriteConnectionSecretToRefWithDefault()
}

// NewWriteConnectionSecretToRefEmpty instantiates a new WriteConnectionSecretToRef object with no properties set.
// This constructor will not assign any default values to properties.
func NewWriteConnectionSecretToRefEmpty() *WriteConnectionSecretToRef {
	this := WriteConnectionSecretToRef{}
	return &this
}

// NewWriteConnectionSecretToRefs converts a list WriteConnectionSecretToRef pointers to objects.
// This is helpful when the SetWriteConnectionSecretToRef requires a list of objects
func NewWriteConnectionSecretToRefList(ps ...*WriteConnectionSecretToRef) []WriteConnectionSecretToRef {
	objs := []WriteConnectionSecretToRef{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this WriteConnectionSecretToRef
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *WriteConnectionSecretToRef) Validate() error {
	if o.Name == nil {
		return errors.New("Name in WriteConnectionSecretToRef must be set")
	}
	if o.Namespace == nil {
		return errors.New("Namespace in WriteConnectionSecretToRef must be set")
	}
	// validate all nested properties
	return nil
}

// GetName returns the Name field value
func (o *WriteConnectionSecretToRef) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WriteConnectionSecretToRef) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *WriteConnectionSecretToRef) SetName(v string) *WriteConnectionSecretToRef {
	o.Name = &v
	return o
}

// GetNamespace returns the Namespace field value
func (o *WriteConnectionSecretToRef) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *WriteConnectionSecretToRef) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Namespace, true
}

// SetNamespace sets field value
func (o *WriteConnectionSecretToRef) SetNamespace(v string) *WriteConnectionSecretToRef {
	o.Namespace = &v
	return o
}

func (o WriteConnectionSecretToRef) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WriteConnectionSecretToRef) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableWriteConnectionSecretToRef struct {
	value *WriteConnectionSecretToRef
	isSet bool
}

func (v *NullableWriteConnectionSecretToRef) Get() *WriteConnectionSecretToRef {
	return v.value
}

func (v *NullableWriteConnectionSecretToRef) Set(val *WriteConnectionSecretToRef) {
	v.value = val
	v.isSet = true
}

func (v *NullableWriteConnectionSecretToRef) IsSet() bool {
	return v.isSet
}

func (v *NullableWriteConnectionSecretToRef) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWriteConnectionSecretToRef(val *WriteConnectionSecretToRef) *NullableWriteConnectionSecretToRef {
	return &NullableWriteConnectionSecretToRef{value: val, isSet: true}
}

func (v NullableWriteConnectionSecretToRef) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWriteConnectionSecretToRef) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}