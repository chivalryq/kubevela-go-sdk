/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package kustomize

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Semver type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Semver{}

// Semver SemVer gives a semantic version range to check against the tags available.
type Semver struct {
	// Range gives a semver range for the image tag; the highest version within the range that's a tag yields the latest image.
	Range *string `json:"range"`
}

// NewSemverWith instantiates a new Semver object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewSemverWith(range_ string) *Semver {
	this := Semver{}
	this.Range = &range_
	return &this
}

// NewSemverWithDefault instantiates a new Semver object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSemverWithDefault() *Semver {
	this := Semver{}
	return &this
}

// NewSemver is short for NewSemverWithDefault which instantiates a new Semver object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSemver() *Semver {
	return NewSemverWithDefault()
}

// NewSemverEmpty instantiates a new Semver object with no properties set.
// This constructor will not assign any default values to properties.
func NewSemverEmpty() *Semver {
	this := Semver{}
	return &this
}

// NewSemvers converts a list Semver pointers to objects.
// This is helpful when the SetSemver requires a list of objects
func NewSemverList(ps ...*Semver) []Semver {
	objs := []Semver{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Semver
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Semver) Validate() error {
	if o.Range == nil {
		return errors.New("Range in Semver must be set")
	}
	// validate all nested properties
	return nil
}

// GetRange returns the Range field value
func (o *Semver) GetRange() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Range
}

// GetRangeOk returns a tuple with the Range field value
// and a boolean to check if the value has been set.
func (o *Semver) GetRangeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Range, true
}

// SetRange sets field value
func (o *Semver) SetRange(v string) *Semver {
	o.Range = &v
	return o
}

func (o Semver) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Semver) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["range"] = o.Range
	return toSerialize, nil
}

type NullableSemver struct {
	value *Semver
	isSet bool
}

func (v *NullableSemver) Get() *Semver {
	return v.value
}

func (v *NullableSemver) Set(val *Semver) {
	v.value = val
	v.isSet = true
}

func (v *NullableSemver) IsSet() bool {
	return v.isSet
}

func (v *NullableSemver) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSemver(val *Semver) *NullableSemver {
	return &NullableSemver{value: val, isSet: true}
}

func (v NullableSemver) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSemver) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
