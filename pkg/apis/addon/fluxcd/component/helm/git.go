/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package helm

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Git type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Git{}

// Git struct for Git
type Git struct {
	// The Git reference to checkout and monitor for changes, defaults to main branch
	Branch *string `json:"branch"`
}

// NewGitWith instantiates a new Git object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewGitWith(branch string) *Git {
	this := Git{}
	this.Branch = &branch
	return &this
}

// NewGitWithDefault instantiates a new Git object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGitWithDefault() *Git {
	this := Git{}
	var branch string = "main"
	this.Branch = &branch
	return &this
}

// NewGit is short for NewGitWithDefault which instantiates a new Git object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGit() *Git {
	return NewGitWithDefault()
}

// NewGitEmpty instantiates a new Git object with no properties set.
// This constructor will not assign any default values to properties.
func NewGitEmpty() *Git {
	this := Git{}
	return &this
}

// NewGits converts a list Git pointers to objects.
// This is helpful when the SetGit requires a list of objects
func NewGitList(ps ...*Git) []Git {
	objs := []Git{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Git
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Git) Validate() error {
	if o.Branch == nil {
		return errors.New("Branch in Git must be set")
	}
	// validate all nested properties
	return nil
}

// GetBranch returns the Branch field value
func (o *Git) GetBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value
// and a boolean to check if the value has been set.
func (o *Git) GetBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Branch, true
}

// SetBranch sets field value
func (o *Git) SetBranch(v string) *Git {
	o.Branch = &v
	return o
}

func (o Git) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Git) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["branch"] = o.Branch
	return toSerialize, nil
}

type NullableGit struct {
	value *Git
	isSet bool
}

func (v *NullableGit) Get() *Git {
	return v.value
}

func (v *NullableGit) Set(val *Git) {
	v.value = val
	v.isSet = true
}

func (v *NullableGit) IsSet() bool {
	return v.isSet
}

func (v *NullableGit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGit(val *Git) *NullableGit {
	return &NullableGit{value: val, isSet: true}
}

func (v NullableGit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
