/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package apply_component

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ApplyComponentSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ApplyComponentSpec{}

// ApplyComponentSpec struct for ApplyComponentSpec
type ApplyComponentSpec struct {
	// Specify the cluster
	Cluster *string `json:"cluster"`
	// Specify the component name to apply
	Component *string `json:"component"`
}

// NewApplyComponentSpecWith instantiates a new ApplyComponentSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewApplyComponentSpecWith(cluster string, component string) *ApplyComponentSpec {
	this := ApplyComponentSpec{}
	this.Cluster = &cluster
	this.Component = &component
	return &this
}

// NewApplyComponentSpecWithDefault instantiates a new ApplyComponentSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyComponentSpecWithDefault() *ApplyComponentSpec {
	this := ApplyComponentSpec{}
	var cluster string = ""
	this.Cluster = &cluster
	return &this
}

// NewApplyComponentSpec is short for NewApplyComponentSpecWithDefault which instantiates a new ApplyComponentSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplyComponentSpec() *ApplyComponentSpec {
	return NewApplyComponentSpecWithDefault()
}

// NewApplyComponentSpecEmpty instantiates a new ApplyComponentSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewApplyComponentSpecEmpty() *ApplyComponentSpec {
	this := ApplyComponentSpec{}
	return &this
}

// NewApplyComponentSpecs converts a list ApplyComponentSpec pointers to objects.
// This is helpful when the SetApplyComponentSpec requires a list of objects
func NewApplyComponentSpecList(ps ...*ApplyComponentSpec) []ApplyComponentSpec {
	objs := []ApplyComponentSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ApplyComponentSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ApplyComponentWorkflowStep) Validate() error {
	if o.Properties.Cluster == nil {
		return errors.New("Cluster in ApplyComponentSpec must be set")
	}
	if o.Properties.Component == nil {
		return errors.New("Component in ApplyComponentSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetCluster returns the Cluster field value
func (o *ApplyComponentWorkflowStep) GetCluster() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Cluster
}

// GetClusterOk returns a tuple with the Cluster field value
// and a boolean to check if the value has been set.
func (o *ApplyComponentWorkflowStep) GetClusterOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Cluster, true
}

// SetCluster sets field value
func (o *ApplyComponentWorkflowStep) SetCluster(v string) *ApplyComponentWorkflowStep {
	o.Properties.Cluster = &v
	return o
}

// GetComponent returns the Component field value
func (o *ApplyComponentWorkflowStep) GetComponent() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Component
}

// GetComponentOk returns a tuple with the Component field value
// and a boolean to check if the value has been set.
func (o *ApplyComponentWorkflowStep) GetComponentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Component, true
}

// SetComponent sets field value
func (o *ApplyComponentWorkflowStep) SetComponent(v string) *ApplyComponentWorkflowStep {
	o.Properties.Component = &v
	return o
}

func (o ApplyComponentSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplyComponentSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["cluster"] = o.Cluster
	toSerialize["component"] = o.Component
	return toSerialize, nil
}

type NullableApplyComponentSpec struct {
	value *ApplyComponentSpec
	isSet bool
}

func (v *NullableApplyComponentSpec) Get() *ApplyComponentSpec {
	return v.value
}

func (v *NullableApplyComponentSpec) Set(val *ApplyComponentSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableApplyComponentSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableApplyComponentSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplyComponentSpec(val *ApplyComponentSpec) *NullableApplyComponentSpec {
	return &NullableApplyComponentSpec{value: val, isSet: true}
}

func (v NullableApplyComponentSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplyComponentSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ApplyComponentType = "apply-component"

func init() {
	sdkcommon.RegisterWorkflowStep(ApplyComponentType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ApplyComponentType, FromWorkflowSubStep)
}

type ApplyComponentWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ApplyComponentSpec
}

func ApplyComponent(name string) *ApplyComponentWorkflowStep {
	a := &ApplyComponentWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ApplyComponentType,
	}}
	return a
}

func (a *ApplyComponentWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range a.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  a.Base.DependsOn,
		If:         a.Base.If,
		Inputs:     a.Base.Inputs,
		Meta:       a.Base.Meta,
		Name:       a.Base.Name,
		Outputs:    a.Base.Outputs,
		Properties: util.Object2RawExtension(a.Properties),
		SubSteps:   subSteps,
		Timeout:    a.Base.Timeout,
		Type:       ApplyComponentType,
	}
	return res
}

func (a *ApplyComponentWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ApplyComponentWorkflowStep, error) {
	var properties ApplyComponentSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := a.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	a.Base.Name = from.Name
	a.Base.DependsOn = from.DependsOn
	a.Base.Inputs = from.Inputs
	a.Base.Outputs = from.Outputs
	a.Base.If = from.If
	a.Base.Timeout = from.Timeout
	a.Base.Meta = from.Meta
	a.Base.Type = ApplyComponentType
	a.Properties = properties
	a.Base.SubSteps = subSteps
	return a, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	a := &ApplyComponentWorkflowStep{}
	return a.FromWorkflowStep(from)
}

func (a *ApplyComponentWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ApplyComponentWorkflowStep, error) {
	var properties ApplyComponentSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	a.Base.Name = from.Name
	a.Base.DependsOn = from.DependsOn
	a.Base.Inputs = from.Inputs
	a.Base.Outputs = from.Outputs
	a.Base.If = from.If
	a.Base.Timeout = from.Timeout
	a.Base.Meta = from.Meta
	a.Base.Type = ApplyComponentType
	a.Properties = properties
	return a, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	a := &ApplyComponentWorkflowStep{}
	return a.FromWorkflowSubStep(from)
}

func (a *ApplyComponentWorkflowStep) WorkflowStepName() string {
	return a.Base.Name
}

func (a *ApplyComponentWorkflowStep) DefType() string {
	return ApplyComponentType
}

func (a *ApplyComponentWorkflowStep) If(_if string) *ApplyComponentWorkflowStep {
	a.Base.If = _if
	return a
}

func (a *ApplyComponentWorkflowStep) Alias(alias string) *ApplyComponentWorkflowStep {
	a.Base.Meta.Alias = alias
	return a
}

func (a *ApplyComponentWorkflowStep) Timeout(timeout string) *ApplyComponentWorkflowStep {
	a.Base.Timeout = timeout
	return a
}

func (a *ApplyComponentWorkflowStep) DependsOn(dependsOn []string) *ApplyComponentWorkflowStep {
	a.Base.DependsOn = dependsOn
	return a
}

func (a *ApplyComponentWorkflowStep) Inputs(input common.StepInputs) *ApplyComponentWorkflowStep {
	a.Base.Inputs = input
	return a
}

func (a *ApplyComponentWorkflowStep) Outputs(output common.StepOutputs) *ApplyComponentWorkflowStep {
	a.Base.Outputs = output
	return a
}
