/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package resource

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ResourceSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ResourceSpec{}

// ResourceSpec struct for ResourceSpec
type ResourceSpec struct {
	// Specify the amount of cpu for requests and limits
	Cpu    *float32 `json:"cpu,omitempty"`
	Limits *Limits  `json:"limits,omitempty"`
	// Specify the amount of memory for requests and limits
	Memory   *string   `json:"memory,omitempty"`
	Requests *Requests `json:"requests,omitempty"`
}

// NewResourceSpecWith instantiates a new ResourceSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewResourceSpecWith() *ResourceSpec {
	this := ResourceSpec{}
	var cpu float32 = 1
	this.Cpu = &cpu
	var memory string = "2048Mi"
	this.Memory = &memory
	return &this
}

// NewResourceSpecWithDefault instantiates a new ResourceSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSpecWithDefault() *ResourceSpec {
	this := ResourceSpec{}
	var cpu float32 = 1
	this.Cpu = &cpu
	var memory string = "2048Mi"
	this.Memory = &memory
	return &this
}

// NewResourceSpec is short for NewResourceSpecWithDefault which instantiates a new ResourceSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewResourceSpec() *ResourceSpec {
	return NewResourceSpecWithDefault()
}

// NewResourceSpecEmpty instantiates a new ResourceSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewResourceSpecEmpty() *ResourceSpec {
	this := ResourceSpec{}
	return &this
}

// NewResourceSpecs converts a list ResourceSpec pointers to objects.
// This is helpful when the SetResourceSpec requires a list of objects
func NewResourceSpecList(ps ...*ResourceSpec) []ResourceSpec {
	objs := []ResourceSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ResourceSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ResourceTrait) Validate() error {
	// validate all nested properties
	if o.Properties.Limits != nil {
		if err := o.Properties.Limits.Validate(); err != nil {
			return err
		}
	}
	if o.Properties.Requests != nil {
		if err := o.Properties.Requests.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *ResourceTrait) GetCpu() float32 {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		var ret float32
		return ret
	}
	return *o.Properties.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetCpuOk() (*float32, bool) {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		return nil, false
	}
	return o.Properties.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *ResourceTrait) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given float32 and assigns it to the cpu field.
// Cpu:  Specify the amount of cpu for requests and limits
func (o *ResourceTrait) SetCpu(v float32) *ResourceTrait {
	o.Properties.Cpu = &v
	return o
}

// GetLimits returns the Limits field value if set, zero value otherwise.
func (o *ResourceTrait) GetLimits() Limits {
	if o == nil || utils.IsNil(o.Properties.Limits) {
		var ret Limits
		return ret
	}
	return *o.Properties.Limits
}

// GetLimitsOk returns a tuple with the Limits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetLimitsOk() (*Limits, bool) {
	if o == nil || utils.IsNil(o.Properties.Limits) {
		return nil, false
	}
	return o.Properties.Limits, true
}

// HasLimits returns a boolean if a field has been set.
func (o *ResourceTrait) HasLimits() bool {
	if o != nil && !utils.IsNil(o.Properties.Limits) {
		return true
	}

	return false
}

// SetLimits gets a reference to the given Limits and assigns it to the limits field.
// Limits:
func (o *ResourceTrait) SetLimits(v Limits) *ResourceTrait {
	o.Properties.Limits = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *ResourceTrait) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		var ret string
		return ret
	}
	return *o.Properties.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		return nil, false
	}
	return o.Properties.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *ResourceTrait) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// Memory:  Specify the amount of memory for requests and limits
func (o *ResourceTrait) SetMemory(v string) *ResourceTrait {
	o.Properties.Memory = &v
	return o
}

// GetRequests returns the Requests field value if set, zero value otherwise.
func (o *ResourceTrait) GetRequests() Requests {
	if o == nil || utils.IsNil(o.Properties.Requests) {
		var ret Requests
		return ret
	}
	return *o.Properties.Requests
}

// GetRequestsOk returns a tuple with the Requests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ResourceTrait) GetRequestsOk() (*Requests, bool) {
	if o == nil || utils.IsNil(o.Properties.Requests) {
		return nil, false
	}
	return o.Properties.Requests, true
}

// HasRequests returns a boolean if a field has been set.
func (o *ResourceTrait) HasRequests() bool {
	if o != nil && !utils.IsNil(o.Properties.Requests) {
		return true
	}

	return false
}

// SetRequests gets a reference to the given Requests and assigns it to the requests field.
// Requests:
func (o *ResourceTrait) SetRequests(v Requests) *ResourceTrait {
	o.Properties.Requests = &v
	return o
}

func (o ResourceSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ResourceSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !utils.IsNil(o.Limits) {
		toSerialize["limits"] = o.Limits
	}
	if !utils.IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !utils.IsNil(o.Requests) {
		toSerialize["requests"] = o.Requests
	}
	return toSerialize, nil
}

type NullableResourceSpec struct {
	value *ResourceSpec
	isSet bool
}

func (v *NullableResourceSpec) Get() *ResourceSpec {
	return v.value
}

func (v *NullableResourceSpec) Set(val *ResourceSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableResourceSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableResourceSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableResourceSpec(val *ResourceSpec) *NullableResourceSpec {
	return &NullableResourceSpec{value: val, isSet: true}
}

func (v NullableResourceSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableResourceSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ResourceType = "resource"

func init() {
	sdkcommon.RegisterTrait(ResourceType, FromTrait)
}

type ResourceTrait struct {
	Base       apis.TraitBase
	Properties ResourceSpec
}

func Resource() *ResourceTrait {
	r := &ResourceTrait{Base: apis.TraitBase{}}
	return r
}

func (r *ResourceTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(r.Properties),
		Type:       ResourceType,
	}
	return res
}

func (r *ResourceTrait) FromTrait(from common.ApplicationTrait) (*ResourceTrait, error) {
	var properties ResourceSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	r.Base.Type = ResourceType
	r.Properties = properties
	return r, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	r := &ResourceTrait{}
	return r.FromTrait(from)
}

func (r *ResourceTrait) DefType() string {
	return ResourceType
}
