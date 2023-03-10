/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package service_binding

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ServiceBindingSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ServiceBindingSpec{}

// ServiceBindingSpec struct for ServiceBindingSpec
type ServiceBindingSpec struct {
	// The mapping of environment variables to secret
	EnvMappings map[string]KeySecret `json:"envMappings"`
}

// NewServiceBindingSpecWith instantiates a new ServiceBindingSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewServiceBindingSpecWith(envMappings map[string]KeySecret) *ServiceBindingSpec {
	this := ServiceBindingSpec{}
	this.EnvMappings = envMappings
	return &this
}

// NewServiceBindingSpecWithDefault instantiates a new ServiceBindingSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingSpecWithDefault() *ServiceBindingSpec {
	this := ServiceBindingSpec{}
	return &this
}

// NewServiceBindingSpec is short for NewServiceBindingSpecWithDefault which instantiates a new ServiceBindingSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceBindingSpec() *ServiceBindingSpec {
	return NewServiceBindingSpecWithDefault()
}

// NewServiceBindingSpecEmpty instantiates a new ServiceBindingSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewServiceBindingSpecEmpty() *ServiceBindingSpec {
	this := ServiceBindingSpec{}
	return &this
}

// NewServiceBindingSpecs converts a list ServiceBindingSpec pointers to objects.
// This is helpful when the SetServiceBindingSpec requires a list of objects
func NewServiceBindingSpecList(ps ...*ServiceBindingSpec) []ServiceBindingSpec {
	objs := []ServiceBindingSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ServiceBindingSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ServiceBindingTrait) Validate() error {
	if o.Properties.EnvMappings == nil {
		return errors.New("EnvMappings in ServiceBindingSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetEnvMappings returns the EnvMappings field value
func (o *ServiceBindingTrait) GetEnvMappings() map[string]KeySecret {
	if o == nil {
		var ret map[string]KeySecret
		return ret
	}

	return o.Properties.EnvMappings
}

// GetEnvMappingsOk returns a tuple with the EnvMappings field value
// and a boolean to check if the value has been set.
func (o *ServiceBindingTrait) GetEnvMappingsOk() (map[string]KeySecret, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.EnvMappings, true
}

// SetEnvMappings sets field value
func (o *ServiceBindingTrait) SetEnvMappings(v map[string]KeySecret) *ServiceBindingTrait {
	o.Properties.EnvMappings = v
	return o
}

func (o ServiceBindingSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ServiceBindingSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["envMappings"] = o.EnvMappings
	return toSerialize, nil
}

type NullableServiceBindingSpec struct {
	value *ServiceBindingSpec
	isSet bool
}

func (v *NullableServiceBindingSpec) Get() *ServiceBindingSpec {
	return v.value
}

func (v *NullableServiceBindingSpec) Set(val *ServiceBindingSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableServiceBindingSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceBindingSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceBindingSpec(val *ServiceBindingSpec) *NullableServiceBindingSpec {
	return &NullableServiceBindingSpec{value: val, isSet: true}
}

func (v NullableServiceBindingSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceBindingSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ServiceBindingType = "service-binding"

func init() {
	sdkcommon.RegisterTrait(ServiceBindingType, FromTrait)
}

type ServiceBindingTrait struct {
	Base       apis.TraitBase
	Properties ServiceBindingSpec
}

func ServiceBinding() *ServiceBindingTrait {
	s := &ServiceBindingTrait{Base: apis.TraitBase{}}
	return s
}

func (s *ServiceBindingTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(s.Properties),
		Type:       ServiceBindingType,
	}
	return res
}

func (s *ServiceBindingTrait) FromTrait(from common.ApplicationTrait) (*ServiceBindingTrait, error) {
	var properties ServiceBindingSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	s.Base.Type = ServiceBindingType
	s.Properties = properties
	return s, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	s := &ServiceBindingTrait{}
	return s.FromTrait(from)
}

func (s *ServiceBindingTrait) DefType() string {
	return ServiceBindingType
}