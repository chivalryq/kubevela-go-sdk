/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cpuscaler

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the CpuscalerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CpuscalerSpec{}

// CpuscalerSpec struct for CpuscalerSpec
type CpuscalerSpec struct {
	// Specify the average CPU utilization, for example, 50 means the CPU usage is 50%
	CpuUtil *int32 `json:"cpuUtil"`
	// Specify the maximum number of of replicas to which the autoscaler can scale up
	Max *int32 `json:"max"`
	// Specify the minimal number of replicas to which the autoscaler can scale down
	Min *int32 `json:"min"`
	// Specify the apiVersion of scale target
	TargetAPIVersion *string `json:"targetAPIVersion"`
	// Specify the kind of scale target
	TargetKind *string `json:"targetKind"`
}

// NewCpuscalerSpecWith instantiates a new CpuscalerSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewCpuscalerSpecWith(cpuUtil int32, max int32, min int32, targetAPIVersion string, targetKind string) *CpuscalerSpec {
	this := CpuscalerSpec{}
	this.CpuUtil = &cpuUtil
	this.Max = &max
	this.Min = &min
	this.TargetAPIVersion = &targetAPIVersion
	this.TargetKind = &targetKind
	return &this
}

// NewCpuscalerSpecWithDefault instantiates a new CpuscalerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuscalerSpecWithDefault() *CpuscalerSpec {
	this := CpuscalerSpec{}
	var cpuUtil int32 = 50
	this.CpuUtil = &cpuUtil
	var max int32 = 10
	this.Max = &max
	var min int32 = 1
	this.Min = &min
	var targetAPIVersion string = "apps/v1"
	this.TargetAPIVersion = &targetAPIVersion
	var targetKind string = "Deployment"
	this.TargetKind = &targetKind
	return &this
}

// NewCpuscalerSpec is short for NewCpuscalerSpecWithDefault which instantiates a new CpuscalerSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCpuscalerSpec() *CpuscalerSpec {
	return NewCpuscalerSpecWithDefault()
}

// NewCpuscalerSpecEmpty instantiates a new CpuscalerSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewCpuscalerSpecEmpty() *CpuscalerSpec {
	this := CpuscalerSpec{}
	return &this
}

// NewCpuscalerSpecs converts a list CpuscalerSpec pointers to objects.
// This is helpful when the SetCpuscalerSpec requires a list of objects
func NewCpuscalerSpecList(ps ...*CpuscalerSpec) []CpuscalerSpec {
	objs := []CpuscalerSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this CpuscalerSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *CpuscalerTrait) Validate() error {
	if o.Properties.CpuUtil == nil {
		return errors.New("CpuUtil in CpuscalerSpec must be set")
	}
	if o.Properties.Max == nil {
		return errors.New("Max in CpuscalerSpec must be set")
	}
	if o.Properties.Min == nil {
		return errors.New("Min in CpuscalerSpec must be set")
	}
	if o.Properties.TargetAPIVersion == nil {
		return errors.New("TargetAPIVersion in CpuscalerSpec must be set")
	}
	if o.Properties.TargetKind == nil {
		return errors.New("TargetKind in CpuscalerSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetCpuUtil returns the CpuUtil field value
func (o *CpuscalerTrait) GetCpuUtil() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.CpuUtil
}

// GetCpuUtilOk returns a tuple with the CpuUtil field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetCpuUtilOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.CpuUtil, true
}

// SetCpuUtil sets field value
func (o *CpuscalerTrait) SetCpuUtil(v int32) *CpuscalerTrait {
	o.Properties.CpuUtil = &v
	return o
}

// GetMax returns the Max field value
func (o *CpuscalerTrait) GetMax() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Max
}

// GetMaxOk returns a tuple with the Max field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetMaxOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Max, true
}

// SetMax sets field value
func (o *CpuscalerTrait) SetMax(v int32) *CpuscalerTrait {
	o.Properties.Max = &v
	return o
}

// GetMin returns the Min field value
func (o *CpuscalerTrait) GetMin() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Min
}

// GetMinOk returns a tuple with the Min field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetMinOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Min, true
}

// SetMin sets field value
func (o *CpuscalerTrait) SetMin(v int32) *CpuscalerTrait {
	o.Properties.Min = &v
	return o
}

// GetTargetAPIVersion returns the TargetAPIVersion field value
func (o *CpuscalerTrait) GetTargetAPIVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.TargetAPIVersion
}

// GetTargetAPIVersionOk returns a tuple with the TargetAPIVersion field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetTargetAPIVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.TargetAPIVersion, true
}

// SetTargetAPIVersion sets field value
func (o *CpuscalerTrait) SetTargetAPIVersion(v string) *CpuscalerTrait {
	o.Properties.TargetAPIVersion = &v
	return o
}

// GetTargetKind returns the TargetKind field value
func (o *CpuscalerTrait) GetTargetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.TargetKind
}

// GetTargetKindOk returns a tuple with the TargetKind field value
// and a boolean to check if the value has been set.
func (o *CpuscalerTrait) GetTargetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.TargetKind, true
}

// SetTargetKind sets field value
func (o *CpuscalerTrait) SetTargetKind(v string) *CpuscalerTrait {
	o.Properties.TargetKind = &v
	return o
}

func (o CpuscalerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CpuscalerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cpuUtil"] = o.CpuUtil
	toSerialize["max"] = o.Max
	toSerialize["min"] = o.Min
	toSerialize["targetAPIVersion"] = o.TargetAPIVersion
	toSerialize["targetKind"] = o.TargetKind
	return toSerialize, nil
}

type NullableCpuscalerSpec struct {
	value *CpuscalerSpec
	isSet bool
}

func (v *NullableCpuscalerSpec) Get() *CpuscalerSpec {
	return v.value
}

func (v *NullableCpuscalerSpec) Set(val *CpuscalerSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableCpuscalerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCpuscalerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCpuscalerSpec(val *CpuscalerSpec) *NullableCpuscalerSpec {
	return &NullableCpuscalerSpec{value: val, isSet: true}
}

func (v NullableCpuscalerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCpuscalerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const CpuscalerType = "cpuscaler"

func init() {
	sdkcommon.RegisterTrait(CpuscalerType, FromTrait)
}

type CpuscalerTrait struct {
	Base       apis.TraitBase
	Properties CpuscalerSpec
}

func Cpuscaler() *CpuscalerTrait {
	c := &CpuscalerTrait{Base: apis.TraitBase{}}
	return c
}

func (c *CpuscalerTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(c.Properties),
		Type:       CpuscalerType,
	}
	return res
}

func (c *CpuscalerTrait) FromTrait(from common.ApplicationTrait) (*CpuscalerTrait, error) {
	var properties CpuscalerSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	c.Base.Type = CpuscalerType
	c.Properties = properties
	return c, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	c := &CpuscalerTrait{}
	return c.FromTrait(from)
}

func (c *CpuscalerTrait) DefType() string {
	return CpuscalerType
}