/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package topologyspreadconstraints

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the TopologyspreadconstraintsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TopologyspreadconstraintsSpec{}

// TopologyspreadconstraintsSpec struct for TopologyspreadconstraintsSpec
type TopologyspreadconstraintsSpec struct {
	Constraints []Constraints `json:"constraints"`
}

// NewTopologyspreadconstraintsSpecWith instantiates a new TopologyspreadconstraintsSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewTopologyspreadconstraintsSpecWith(constraints []Constraints) *TopologyspreadconstraintsSpec {
	this := TopologyspreadconstraintsSpec{}
	this.Constraints = constraints
	return &this
}

// NewTopologyspreadconstraintsSpecWithDefault instantiates a new TopologyspreadconstraintsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyspreadconstraintsSpecWithDefault() *TopologyspreadconstraintsSpec {
	this := TopologyspreadconstraintsSpec{}
	return &this
}

// NewTopologyspreadconstraintsSpec is short for NewTopologyspreadconstraintsSpecWithDefault which instantiates a new TopologyspreadconstraintsSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTopologyspreadconstraintsSpec() *TopologyspreadconstraintsSpec {
	return NewTopologyspreadconstraintsSpecWithDefault()
}

// NewTopologyspreadconstraintsSpecEmpty instantiates a new TopologyspreadconstraintsSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewTopologyspreadconstraintsSpecEmpty() *TopologyspreadconstraintsSpec {
	this := TopologyspreadconstraintsSpec{}
	return &this
}

// NewTopologyspreadconstraintsSpecs converts a list TopologyspreadconstraintsSpec pointers to objects.
// This is helpful when the SetTopologyspreadconstraintsSpec requires a list of objects
func NewTopologyspreadconstraintsSpecList(ps ...*TopologyspreadconstraintsSpec) []TopologyspreadconstraintsSpec {
	objs := []TopologyspreadconstraintsSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this TopologyspreadconstraintsSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *TopologyspreadconstraintsTrait) Validate() error {
	if o.Properties.Constraints == nil {
		return errors.New("Constraints in TopologyspreadconstraintsSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetConstraints returns the Constraints field value
func (o *TopologyspreadconstraintsTrait) GetConstraints() []Constraints {
	if o == nil {
		var ret []Constraints
		return ret
	}

	return o.Properties.Constraints
}

// GetConstraintsOk returns a tuple with the Constraints field value
// and a boolean to check if the value has been set.
func (o *TopologyspreadconstraintsTrait) GetConstraintsOk() ([]Constraints, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Constraints, true
}

// SetConstraints sets field value
func (o *TopologyspreadconstraintsTrait) SetConstraints(v []Constraints) *TopologyspreadconstraintsTrait {
	o.Properties.Constraints = v
	return o
}

func (o TopologyspreadconstraintsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TopologyspreadconstraintsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["constraints"] = o.Constraints
	return toSerialize, nil
}

type NullableTopologyspreadconstraintsSpec struct {
	value *TopologyspreadconstraintsSpec
	isSet bool
}

func (v *NullableTopologyspreadconstraintsSpec) Get() *TopologyspreadconstraintsSpec {
	return v.value
}

func (v *NullableTopologyspreadconstraintsSpec) Set(val *TopologyspreadconstraintsSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableTopologyspreadconstraintsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableTopologyspreadconstraintsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTopologyspreadconstraintsSpec(val *TopologyspreadconstraintsSpec) *NullableTopologyspreadconstraintsSpec {
	return &NullableTopologyspreadconstraintsSpec{value: val, isSet: true}
}

func (v NullableTopologyspreadconstraintsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTopologyspreadconstraintsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const TopologyspreadconstraintsType = "topologyspreadconstraints"

func init() {
	sdkcommon.RegisterTrait(TopologyspreadconstraintsType, FromTrait)
}

type TopologyspreadconstraintsTrait struct {
	Base       apis.TraitBase
	Properties TopologyspreadconstraintsSpec
}

func Topologyspreadconstraints() *TopologyspreadconstraintsTrait {
	t := &TopologyspreadconstraintsTrait{Base: apis.TraitBase{}}
	return t
}

func (t *TopologyspreadconstraintsTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(t.Properties),
		Type:       TopologyspreadconstraintsType,
	}
	return res
}

func (t *TopologyspreadconstraintsTrait) FromTrait(from common.ApplicationTrait) (*TopologyspreadconstraintsTrait, error) {
	var properties TopologyspreadconstraintsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	t.Base.Type = TopologyspreadconstraintsType
	t.Properties = properties
	return t, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	t := &TopologyspreadconstraintsTrait{}
	return t.FromTrait(from)
}

func (t *TopologyspreadconstraintsTrait) DefType() string {
	return TopologyspreadconstraintsType
}
