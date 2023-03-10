/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lifecycle

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the LifecycleSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &LifecycleSpec{}

// LifecycleSpec struct for LifecycleSpec
type LifecycleSpec struct {
	PostStart *LifeCycleHandler `json:"postStart,omitempty"`
	PreStop   *LifeCycleHandler `json:"preStop,omitempty"`
}

// NewLifecycleSpecWith instantiates a new LifecycleSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewLifecycleSpecWith() *LifecycleSpec {
	this := LifecycleSpec{}
	return &this
}

// NewLifecycleSpecWithDefault instantiates a new LifecycleSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleSpecWithDefault() *LifecycleSpec {
	this := LifecycleSpec{}
	return &this
}

// NewLifecycleSpec is short for NewLifecycleSpecWithDefault which instantiates a new LifecycleSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLifecycleSpec() *LifecycleSpec {
	return NewLifecycleSpecWithDefault()
}

// NewLifecycleSpecEmpty instantiates a new LifecycleSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewLifecycleSpecEmpty() *LifecycleSpec {
	this := LifecycleSpec{}
	return &this
}

// NewLifecycleSpecs converts a list LifecycleSpec pointers to objects.
// This is helpful when the SetLifecycleSpec requires a list of objects
func NewLifecycleSpecList(ps ...*LifecycleSpec) []LifecycleSpec {
	objs := []LifecycleSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this LifecycleSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *LifecycleTrait) Validate() error {
	// validate all nested properties
	if o.Properties.PostStart != nil {
		if err := o.Properties.PostStart.Validate(); err != nil {
			return err
		}
	}
	if o.Properties.PreStop != nil {
		if err := o.Properties.PreStop.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetPostStart returns the PostStart field value if set, zero value otherwise.
func (o *LifecycleTrait) GetPostStart() LifeCycleHandler {
	if o == nil || utils.IsNil(o.Properties.PostStart) {
		var ret LifeCycleHandler
		return ret
	}
	return *o.Properties.PostStart
}

// GetPostStartOk returns a tuple with the PostStart field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleTrait) GetPostStartOk() (*LifeCycleHandler, bool) {
	if o == nil || utils.IsNil(o.Properties.PostStart) {
		return nil, false
	}
	return o.Properties.PostStart, true
}

// HasPostStart returns a boolean if a field has been set.
func (o *LifecycleTrait) HasPostStart() bool {
	if o != nil && !utils.IsNil(o.Properties.PostStart) {
		return true
	}

	return false
}

// SetPostStart gets a reference to the given LifeCycleHandler and assigns it to the postStart field.
// PostStart:
func (o *LifecycleTrait) SetPostStart(v LifeCycleHandler) *LifecycleTrait {
	o.Properties.PostStart = &v
	return o
}

// GetPreStop returns the PreStop field value if set, zero value otherwise.
func (o *LifecycleTrait) GetPreStop() LifeCycleHandler {
	if o == nil || utils.IsNil(o.Properties.PreStop) {
		var ret LifeCycleHandler
		return ret
	}
	return *o.Properties.PreStop
}

// GetPreStopOk returns a tuple with the PreStop field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LifecycleTrait) GetPreStopOk() (*LifeCycleHandler, bool) {
	if o == nil || utils.IsNil(o.Properties.PreStop) {
		return nil, false
	}
	return o.Properties.PreStop, true
}

// HasPreStop returns a boolean if a field has been set.
func (o *LifecycleTrait) HasPreStop() bool {
	if o != nil && !utils.IsNil(o.Properties.PreStop) {
		return true
	}

	return false
}

// SetPreStop gets a reference to the given LifeCycleHandler and assigns it to the preStop field.
// PreStop:
func (o *LifecycleTrait) SetPreStop(v LifeCycleHandler) *LifecycleTrait {
	o.Properties.PreStop = &v
	return o
}

func (o LifecycleSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o LifecycleSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.PostStart) {
		toSerialize["postStart"] = o.PostStart
	}
	if !utils.IsNil(o.PreStop) {
		toSerialize["preStop"] = o.PreStop
	}
	return toSerialize, nil
}

type NullableLifecycleSpec struct {
	value *LifecycleSpec
	isSet bool
}

func (v *NullableLifecycleSpec) Get() *LifecycleSpec {
	return v.value
}

func (v *NullableLifecycleSpec) Set(val *LifecycleSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableLifecycleSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycleSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycleSpec(val *LifecycleSpec) *NullableLifecycleSpec {
	return &NullableLifecycleSpec{value: val, isSet: true}
}

func (v NullableLifecycleSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycleSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const LifecycleType = "lifecycle"

func init() {
	sdkcommon.RegisterTrait(LifecycleType, FromTrait)
}

type LifecycleTrait struct {
	Base       apis.TraitBase
	Properties LifecycleSpec
}

func Lifecycle() *LifecycleTrait {
	l := &LifecycleTrait{Base: apis.TraitBase{}}
	return l
}

func (l *LifecycleTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(l.Properties),
		Type:       LifecycleType,
	}
	return res
}

func (l *LifecycleTrait) FromTrait(from common.ApplicationTrait) (*LifecycleTrait, error) {
	var properties LifecycleSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	l.Base.Type = LifecycleType
	l.Properties = properties
	return l, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	l := &LifecycleTrait{}
	return l.FromTrait(from)
}

func (l *LifecycleTrait) DefType() string {
	return LifecycleType
}
