/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package storage

import (
	"encoding/json"
	"errors"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the DataSource type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DataSource{}

// DataSource struct for DataSource
type DataSource struct {
	ApiGroup *string `json:"apiGroup"`
	Kind     *string `json:"kind"`
	Name     *string `json:"name"`
}

// NewDataSourceWith instantiates a new DataSource object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewDataSourceWith(apiGroup string, kind string, name string) *DataSource {
	this := DataSource{}
	this.ApiGroup = &apiGroup
	this.Kind = &kind
	this.Name = &name
	return &this
}

// NewDataSourceWithDefault instantiates a new DataSource object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSourceWithDefault() *DataSource {
	this := DataSource{}
	return &this
}

// NewDataSource is short for NewDataSourceWithDefault which instantiates a new DataSource object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataSource() *DataSource {
	return NewDataSourceWithDefault()
}

// NewDataSourceEmpty instantiates a new DataSource object with no properties set.
// This constructor will not assign any default values to properties.
func NewDataSourceEmpty() *DataSource {
	this := DataSource{}
	return &this
}

// NewDataSources converts a list DataSource pointers to objects.
// This is helpful when the SetDataSource requires a list of objects
func NewDataSourceList(ps ...*DataSource) []DataSource {
	objs := []DataSource{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this DataSource
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *DataSource) Validate() error {
	if o.ApiGroup == nil {
		return errors.New("ApiGroup in DataSource must be set")
	}
	if o.Kind == nil {
		return errors.New("Kind in DataSource must be set")
	}
	if o.Name == nil {
		return errors.New("Name in DataSource must be set")
	}
	// validate all nested properties
	return nil
}

// GetApiGroup returns the ApiGroup field value
func (o *DataSource) GetApiGroup() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.ApiGroup
}

// GetApiGroupOk returns a tuple with the ApiGroup field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetApiGroupOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.ApiGroup, true
}

// SetApiGroup sets field value
func (o *DataSource) SetApiGroup(v string) *DataSource {
	o.ApiGroup = &v
	return o
}

// GetKind returns the Kind field value
func (o *DataSource) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Kind, true
}

// SetKind sets field value
func (o *DataSource) SetKind(v string) *DataSource {
	o.Kind = &v
	return o
}

// GetName returns the Name field value
func (o *DataSource) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *DataSource) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name, true
}

// SetName sets field value
func (o *DataSource) SetName(v string) *DataSource {
	o.Name = &v
	return o
}

func (o DataSource) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DataSource) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiGroup"] = o.ApiGroup
	toSerialize["kind"] = o.Kind
	toSerialize["name"] = o.Name
	return toSerialize, nil
}

type NullableDataSource struct {
	value *DataSource
	isSet bool
}

func (v *NullableDataSource) Get() *DataSource {
	return v.value
}

func (v *NullableDataSource) Set(val *DataSource) {
	v.value = val
	v.isSet = true
}

func (v *NullableDataSource) IsSet() bool {
	return v.isSet
}

func (v *NullableDataSource) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataSource(val *DataSource) *NullableDataSource {
	return &NullableDataSource{value: val, isSet: true}
}

func (v NullableDataSource) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataSource) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
