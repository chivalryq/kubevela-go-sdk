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

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the SourceOneOf1 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SourceOneOf1{}

// SourceOneOf1 struct for SourceOneOf1
type SourceOneOf1 struct {
	// specify the path of the terraform configuration
	Path *string `json:"path,omitempty"`
	// specify the remote url of the terraform configuration
	Remote *string `json:"remote"`
}

// NewSourceOneOf1With instantiates a new SourceOneOf1 object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewSourceOneOf1With(remote string) *SourceOneOf1 {
	this := SourceOneOf1{}
	this.Remote = &remote
	return &this
}

// NewSourceOneOf1WithDefault instantiates a new SourceOneOf1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceOneOf1WithDefault() *SourceOneOf1 {
	this := SourceOneOf1{}
	var remote string = "https://github.com/chivalryq/terraform-modules.git"
	this.Remote = &remote
	return &this
}

// NewSourceOneOf1 is short for NewSourceOneOf1WithDefault which instantiates a new SourceOneOf1 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSourceOneOf1() *SourceOneOf1 {
	return NewSourceOneOf1WithDefault()
}

// NewSourceOneOf1Empty instantiates a new SourceOneOf1 object with no properties set.
// This constructor will not assign any default values to properties.
func NewSourceOneOf1Empty() *SourceOneOf1 {
	this := SourceOneOf1{}
	return &this
}

// NewSourceOneOf1s converts a list SourceOneOf1 pointers to objects.
// This is helpful when the SetSourceOneOf1 requires a list of objects
func NewSourceOneOf1List(ps ...*SourceOneOf1) []SourceOneOf1 {
	objs := []SourceOneOf1{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this SourceOneOf1
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *SourceOneOf1) Validate() error {
	if o.Remote == nil {
		return errors.New("Remote in SourceOneOf1 must be set")
	}
	// validate all nested properties
	return nil
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *SourceOneOf1) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SourceOneOf1) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *SourceOneOf1) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the path field.
// Path:  specify the path of the terraform configuration
func (o *SourceOneOf1) SetPath(v string) *SourceOneOf1 {
	o.Path = &v
	return o
}

// GetRemote returns the Remote field value
func (o *SourceOneOf1) GetRemote() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Remote
}

// GetRemoteOk returns a tuple with the Remote field value
// and a boolean to check if the value has been set.
func (o *SourceOneOf1) GetRemoteOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Remote, true
}

// SetRemote sets field value
func (o *SourceOneOf1) SetRemote(v string) *SourceOneOf1 {
	o.Remote = &v
	return o
}

func (o SourceOneOf1) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SourceOneOf1) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	toSerialize["remote"] = o.Remote
	return toSerialize, nil
}

type NullableSourceOneOf1 struct {
	value *SourceOneOf1
	isSet bool
}

func (v *NullableSourceOneOf1) Get() *SourceOneOf1 {
	return v.value
}

func (v *NullableSourceOneOf1) Set(val *SourceOneOf1) {
	v.value = val
	v.isSet = true
}

func (v *NullableSourceOneOf1) IsSet() bool {
	return v.isSet
}

func (v *NullableSourceOneOf1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSourceOneOf1(val *SourceOneOf1) *NullableSourceOneOf1 {
	return &NullableSourceOneOf1{value: val, isSet: true}
}

func (v NullableSourceOneOf1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSourceOneOf1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
