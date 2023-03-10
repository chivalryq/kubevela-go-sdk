/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_once

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyOnceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyOnceSpec{}

// ApplyOnceSpec struct for ApplyOnceSpec
type ApplyOnceSpec struct {
	// Whether to enable apply-once for the whole application
	Enable *bool `json:"enable"`
	// Specify the rules for configuring apply-once policy in resource level
	Rules []ApplyOncePolicyRule `json:"rules,omitempty"`
}

// NewApplyOnceSpecWith instantiates a new ApplyOnceSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewApplyOnceSpecWith(enable bool) *ApplyOnceSpec {
	this := ApplyOnceSpec{}
	this.Enable = &enable
	return &this
}

// NewApplyOnceSpecWithDefault instantiates a new ApplyOnceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyOnceSpecWithDefault() *ApplyOnceSpec {
	this := ApplyOnceSpec{}
	var enable bool = false
	this.Enable = &enable
	return &this
}

// NewApplyOnceSpec is short for NewApplyOnceSpecWithDefault which instantiates a new ApplyOnceSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyOnceSpec() *ApplyOnceSpec {
	return NewApplyOnceSpecWithDefault()
}

// NewApplyOnceSpecEmpty instantiates a new ApplyOnceSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewApplyOnceSpecEmpty() *ApplyOnceSpec {
	this := ApplyOnceSpec{}
	return &this
}

// NewApplyOnceSpecs converts a list ApplyOnceSpec pointers to objects.
// This is helpful when the SetApplyOnceSpec requires a list of objects
func NewApplyOnceSpecList(ps ...*ApplyOnceSpec) []ApplyOnceSpec {
	objs := []ApplyOnceSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ApplyOnceSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ApplyOncePolicy) Validate() error {
	if o.Properties.Enable == nil {
		return errors.New("Enable in ApplyOnceSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetEnable returns the Enable field value
func (o *ApplyOncePolicy) GetEnable() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.Enable
}

// GetEnableOk returns a tuple with the Enable field value
// and a boolean to check if the value has been set.
func (o *ApplyOncePolicy) GetEnableOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Enable, true
}

// SetEnable sets field value
func (o *ApplyOncePolicy) SetEnable(v bool) *ApplyOncePolicy {
	o.Properties.Enable = &v
	return o
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *ApplyOncePolicy) GetRules() []ApplyOncePolicyRule {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		var ret []ApplyOncePolicyRule
		return ret
	}
	return o.Properties.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplyOncePolicy) GetRulesOk() ([]ApplyOncePolicyRule, bool) {
	if o == nil || utils.IsNil(o.Properties.Rules) {
		return nil, false
	}
	return o.Properties.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *ApplyOncePolicy) HasRules() bool {
	if o != nil && !utils.IsNil(o.Properties.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []ApplyOncePolicyRule and assigns it to the rules field.
// Rules:  Specify the rules for configuring apply-once policy in resource level
func (o *ApplyOncePolicy) SetRules(v []ApplyOncePolicyRule) *ApplyOncePolicy {
	o.Properties.Rules = v
	return o
}

func (o ApplyOnceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyOnceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["enable"] = o.Enable
	if !utils.IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	return toSerialize, nil
}

type NullableApplyOnceSpec struct {
	value *ApplyOnceSpec
	isSet bool
}

func (v *NullableApplyOnceSpec) Get() *ApplyOnceSpec {
	return v.value
}

func (v *NullableApplyOnceSpec) Set(val *ApplyOnceSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableApplyOnceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyOnceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyOnceSpec(val *ApplyOnceSpec) *NullableApplyOnceSpec {
	return &NullableApplyOnceSpec{value: val, isSet: true}
}

func (v NullableApplyOnceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyOnceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyOnceType = "apply-once"

func init() {
	sdkcommon.RegisterPolicy(ApplyOnceType, FromPolicy)
}

type ApplyOncePolicy struct {
	Base       apis.PolicyBase
	Properties ApplyOnceSpec
}

func ApplyOnce(name string) *ApplyOncePolicy {
	a := &ApplyOncePolicy{Base: apis.PolicyBase{
		Name: name,
		Type: ApplyOnceType,
	}}
	return a
}

func (a *ApplyOncePolicy) Build() v1beta1.AppPolicy {
	res := v1beta1.AppPolicy{
		Name:       a.Base.Name,
		Properties: util.Object2RawExtension(a.Properties),
		Type:       ApplyOnceType,
	}
	return res
}

func (a *ApplyOncePolicy) FromPolicy(from v1beta1.AppPolicy) (*ApplyOncePolicy, error) {
	var properties ApplyOnceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	a.Base.Name = from.Name
	a.Base.Type = ApplyOnceType
	a.Properties = properties
	return a, nil
}

func FromPolicy(from v1beta1.AppPolicy) (apis.Policy, error) {
	a := &ApplyOncePolicy{}
	return a.FromPolicy(from)
}

func (a *ApplyOncePolicy) PolicyName() string {
	return a.Base.Name
}

func (a *ApplyOncePolicy) DefType() string {
	return ApplyOnceType
}