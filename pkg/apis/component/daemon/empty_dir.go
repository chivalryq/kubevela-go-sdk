/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package daemon

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the EmptyDir type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &EmptyDir{}

// EmptyDir struct for EmptyDir
type EmptyDir struct {
	Medium    *string `json:"medium"`
	MountPath *string `json:"mountPath"`
	Name      *string `json:"name"`
}

// NewEmptyDirWith instantiates a new EmptyDir object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewEmptyDirWith(medium string, mountPath string, name string) *EmptyDir {
	this := EmptyDir{}
	this.Medium = &medium
	this.MountPath = &mountPath
	this.Name = &name
	return &this
}

// NewEmptyDirWithDefault instantiates a new EmptyDir object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmptyDirWithDefault() *EmptyDir {
	this := EmptyDir{}
	var medium string = ""
	this.Medium = &medium
	return &this
}

// NewEmptyDir is short for NewEmptyDirWithDefault which instantiates a new EmptyDir object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmptyDir() *EmptyDir {
	return NewEmptyDirWithDefault()
}

// NewEmptyDirEmpty instantiates a new EmptyDir object with no properties set.
// This constructor will not assign any default values to properties.
func NewEmptyDirEmpty() *EmptyDir {
	this := EmptyDir{}
	return &this
}

// NewEmptyDirs converts a list EmptyDir pointers to objects.
// This is helpful when the SetEmptyDir requires a list of objects
func NewEmptyDirList(ps ...*EmptyDir) []EmptyDir {
	objs := []EmptyDir{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this EmptyDir
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *EmptyDir) Validate() error {
	if o.Medium == nil {
		return errors.New("Medium in EmptyDir must be set")
	}
	if o.MountPath == nil {
		return errors.New("MountPath in EmptyDir must be set")
	}
	if o.Name == nil {
		return errors.New("Name in EmptyDir must be set")
	}
	// validate all nested properties
	return nil
}

// GetMedium returns the Medium field value
func (o *EmptyDir) GetMedium() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Medium
}

// GetMediumOk returns a tuple with the Medium field value
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetMediumOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Medium, true
}

// SetMedium sets field value
func (o *EmptyDir) SetMedium(v string) *EmptyDir {
	o.Medium = &v
	return o
}

// GetMountPath returns the MountPath field value
func (o *EmptyDir) GetMountPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.MountPath
}

// GetMountPathOk returns a tuple with the MountPath field value
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetMountPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.MountPath, true
}

// SetMountPath sets field value
func (o *EmptyDir) SetMountPath(v string) *EmptyDir {
	o.MountPath = &v
	return o
}

// GetName returns the Name field value
func (o *EmptyDir) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EmptyDir) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *EmptyDir) SetName(v string) *EmptyDir {
	o.Name = &v
	return o
}

func (o EmptyDir) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EmptyDir) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["medium"] = o.Medium
	toSerialize["mountPath"] = o.MountPath
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableEmptyDir struct {
	value *EmptyDir
	isSet bool
}

func (v *NullableEmptyDir) Get() *EmptyDir {
	return v.value
}

func (v *NullableEmptyDir) Set(val *EmptyDir) {
	v.value = val
	v.isSet = true
}

func (v *NullableEmptyDir) IsSet() bool {
	return v.isSet
}

func (v *NullableEmptyDir) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmptyDir(val *EmptyDir) *NullableEmptyDir {
	return &NullableEmptyDir{value: val, isSet: true}
}

func (v NullableEmptyDir) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmptyDir) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
