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

// checks if the MatchExpressions type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &MatchExpressions{}

// MatchExpressions struct for MatchExpressions
type MatchExpressions struct {
	Key      *string  `json:"key"`
	Operator *string  `json:"operator"`
	Values   []string `json:"values,omitempty"`
}

// NewMatchExpressionsWith instantiates a new MatchExpressions object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewMatchExpressionsWith(key string, operator string) *MatchExpressions {
	this := MatchExpressions{}
	this.Key = &key
	this.Operator = &operator
	return &this
}

// NewMatchExpressionsWithDefault instantiates a new MatchExpressions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchExpressionsWithDefault() *MatchExpressions {
	this := MatchExpressions{}
	var operator string = "In"
	this.Operator = &operator
	return &this
}

// NewMatchExpressions is short for NewMatchExpressionsWithDefault which instantiates a new MatchExpressions object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMatchExpressions() *MatchExpressions {
	return NewMatchExpressionsWithDefault()
}

// NewMatchExpressionsEmpty instantiates a new MatchExpressions object with no properties set.
// This constructor will not assign any default values to properties.
func NewMatchExpressionsEmpty() *MatchExpressions {
	this := MatchExpressions{}
	return &this
}

// NewMatchExpressionss converts a list MatchExpressions pointers to objects.
// This is helpful when the SetMatchExpressions requires a list of objects
func NewMatchExpressionsList(ps ...*MatchExpressions) []MatchExpressions {
	objs := []MatchExpressions{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this MatchExpressions
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *MatchExpressions) Validate() error {
	if o.Key == nil {
		return errors.New("Key in MatchExpressions must be set")
	}
	if o.Operator == nil {
		return errors.New("Operator in MatchExpressions must be set")
	}
	// validate all nested properties
	return nil
}

// GetKey returns the Key field value
func (o *MatchExpressions) GetKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value
// and a boolean to check if the value has been set.
func (o *MatchExpressions) GetKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Key, true
}

// SetKey sets field value
func (o *MatchExpressions) SetKey(v string) *MatchExpressions {
	o.Key = &v
	return o
}

// GetOperator returns the Operator field value
func (o *MatchExpressions) GetOperator() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Operator
}

// GetOperatorOk returns a tuple with the Operator field value
// and a boolean to check if the value has been set.
func (o *MatchExpressions) GetOperatorOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Operator, true
}

// SetOperator sets field value
func (o *MatchExpressions) SetOperator(v string) *MatchExpressions {
	o.Operator = &v
	return o
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *MatchExpressions) GetValues() []string {
	if o == nil || utils.IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MatchExpressions) GetValuesOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *MatchExpressions) HasValues() bool {
	if o != nil && !utils.IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the values field.
// Values:
func (o *MatchExpressions) SetValues(v []string) *MatchExpressions {
	o.Values = v
	return o
}

func (o MatchExpressions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MatchExpressions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["key"] = o.Key
	toSerialize["operator"] = o.Operator
	if !utils.IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableMatchExpressions struct {
	value *MatchExpressions
	isSet bool
}

func (v *NullableMatchExpressions) Get() *MatchExpressions {
	return v.value
}

func (v *NullableMatchExpressions) Set(val *MatchExpressions) {
	v.value = val
	v.isSet = true
}

func (v *NullableMatchExpressions) IsSet() bool {
	return v.isSet
}

func (v *NullableMatchExpressions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMatchExpressions(val *MatchExpressions) *NullableMatchExpressions {
	return &NullableMatchExpressions{value: val, isSet: true}
}

func (v NullableMatchExpressions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMatchExpressions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
