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

// checks if the Items type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Items{}

// Items struct for Items
type Items struct {
	Key  *string `json:"key"`
	Mode *int32  `json:"mode"`
	Path *string `json:"path"`
}

// NewItemsWith instantiates a new Items object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewItemsWith(key string, mode int32, path string) *Items {
	this := Items{}
	this.Key = &key
	this.Mode = &mode
	this.Path = &path
	return &this
}

// NewItemsWithDefault instantiates a new Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItemsWithDefault() *Items {
	this := Items{}
	var mode int32 = 511
	this.Mode = &mode
	return &this
}

// NewItems is short for NewItemsWithDefault which instantiates a new Items object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewItems() *Items {
	return NewItemsWithDefault()
}

// NewItemsEmpty instantiates a new Items object with no properties set.
// This constructor will not assign any default values to properties.
func NewItemsEmpty() *Items {
	this := Items{}
	return &this
}

// NewItemss converts a list Items pointers to objects.
// This is helpful when the SetItems requires a list of objects
func NewItemsList(ps ...*Items) []Items {
	objs := []Items{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Items
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Items) Validate() error {
	if o.Key == nil {
		return errors.New("Key in Items must be set")
	}
	if o.Mode == nil {
		return errors.New("Mode in Items must be set")
	}
	if o.Path == nil {
		return errors.New("Path in Items must be set")
	}
	// validate all nested properties
	return nil
}

// GetKey returns the Key field value
func (o *Items) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *Items) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key, true
}

// SetKey sets field value
func (o *Items) SetKey(v string) *Items {
	o.Key = &v
	return o
}

// GetMode returns the Mode field value
func (o *Items) GetMode() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value
// and a boolean to check if the value has been set.
func (o *Items) GetModeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Mode, true
}

// SetMode sets field value
func (o *Items) SetMode(v int32) *Items {
	o.Mode = &v
	return o
}

// GetPath returns the Path field value
func (o *Items) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Items) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Path, true
}

// SetPath sets field value
func (o *Items) SetPath(v string) *Items {
	o.Path = &v
	return o
}

func (o Items) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Items) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["mode"] = o.Mode
	toSerialize["path"] = o.Path
	return toSerialize, nil
}

type NullableItems struct {
	value *Items
	isSet bool
}

func (v *NullableItems) Get() *Items {
	return v.value
}

func (v *NullableItems) Set(val *Items) {
	v.value = val
	v.isSet = true
}

func (v *NullableItems) IsSet() bool {
	return v.isSet
}

func (v *NullableItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableItems(val *Items) *NullableItems {
	return &NullableItems{value: val, isSet: true}
}

func (v NullableItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
