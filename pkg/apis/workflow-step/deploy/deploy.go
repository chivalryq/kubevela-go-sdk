/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package deploy

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

// checks if the DeploySpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &DeploySpec{}

// DeploySpec struct for DeploySpec
type DeploySpec struct {
	// If set to false, the workflow will suspend automatically before this step, default to be true.
	Auto *bool `json:"auto"`
	// If set false, this step will apply the components with the terraform workload.
	IgnoreTerraformComponent *bool `json:"ignoreTerraformComponent"`
	// Maximum number of concurrent delivered components.
	Parallelism *int32 `json:"parallelism"`
	// Declare the policies that used for this deployment. If not specified, the components will be deployed to the hub cluster.
	Policies []string `json:"policies"`
}

// NewDeploySpecWith instantiates a new DeploySpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewDeploySpecWith(auto bool, ignoreTerraformComponent bool, parallelism int32, policies []string) *DeploySpec {
	this := DeploySpec{}
	this.Auto = &auto
	this.IgnoreTerraformComponent = &ignoreTerraformComponent
	this.Parallelism = &parallelism
	this.Policies = policies
	return &this
}

// NewDeploySpecWithDefault instantiates a new DeploySpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySpecWithDefault() *DeploySpec {
	this := DeploySpec{}
	var auto bool = true
	this.Auto = &auto
	var ignoreTerraformComponent bool = true
	this.IgnoreTerraformComponent = &ignoreTerraformComponent
	var parallelism int32 = 5
	this.Parallelism = &parallelism
	return &this
}

// NewDeploySpec is short for NewDeploySpecWithDefault which instantiates a new DeploySpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeploySpec() *DeploySpec {
	return NewDeploySpecWithDefault()
}

// NewDeploySpecEmpty instantiates a new DeploySpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewDeploySpecEmpty() *DeploySpec {
	this := DeploySpec{}
	return &this
}

// NewDeploySpecs converts a list DeploySpec pointers to objects.
// This is helpful when the SetDeploySpec requires a list of objects
func NewDeploySpecList(ps ...*DeploySpec) []DeploySpec {
	objs := []DeploySpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this DeploySpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *DeployWorkflowStep) Validate() error {
	if o.Properties.Auto == nil {
		return errors.New("Auto in DeploySpec must be set")
	}
	if o.Properties.IgnoreTerraformComponent == nil {
		return errors.New("IgnoreTerraformComponent in DeploySpec must be set")
	}
	if o.Properties.Parallelism == nil {
		return errors.New("Parallelism in DeploySpec must be set")
	}
	if o.Properties.Policies == nil {
		return errors.New("Policies in DeploySpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetAuto returns the Auto field value
func (o *DeployWorkflowStep) GetAuto() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.Auto
}

// GetAutoOk returns a tuple with the Auto field value
// and a boolean to check if the value has been set.
func (o *DeployWorkflowStep) GetAutoOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Auto, true
}

// SetAuto sets field value
func (o *DeployWorkflowStep) SetAuto(v bool) *DeployWorkflowStep {
	o.Properties.Auto = &v
	return o
}

// GetIgnoreTerraformComponent returns the IgnoreTerraformComponent field value
func (o *DeployWorkflowStep) GetIgnoreTerraformComponent() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.IgnoreTerraformComponent
}

// GetIgnoreTerraformComponentOk returns a tuple with the IgnoreTerraformComponent field value
// and a boolean to check if the value has been set.
func (o *DeployWorkflowStep) GetIgnoreTerraformComponentOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.IgnoreTerraformComponent, true
}

// SetIgnoreTerraformComponent sets field value
func (o *DeployWorkflowStep) SetIgnoreTerraformComponent(v bool) *DeployWorkflowStep {
	o.Properties.IgnoreTerraformComponent = &v
	return o
}

// GetParallelism returns the Parallelism field value
func (o *DeployWorkflowStep) GetParallelism() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Parallelism
}

// GetParallelismOk returns a tuple with the Parallelism field value
// and a boolean to check if the value has been set.
func (o *DeployWorkflowStep) GetParallelismOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Parallelism, true
}

// SetParallelism sets field value
func (o *DeployWorkflowStep) SetParallelism(v int32) *DeployWorkflowStep {
	o.Properties.Parallelism = &v
	return o
}

// GetPolicies returns the Policies field value
func (o *DeployWorkflowStep) GetPolicies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties.Policies
}

// GetPoliciesOk returns a tuple with the Policies field value
// and a boolean to check if the value has been set.
func (o *DeployWorkflowStep) GetPoliciesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Policies, true
}

// SetPolicies sets field value
func (o *DeployWorkflowStep) SetPolicies(v []string) *DeployWorkflowStep {
	o.Properties.Policies = v
	return o
}

func (o DeploySpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeploySpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["auto"] = o.Auto
	toSerialize["ignoreTerraformComponent"] = o.IgnoreTerraformComponent
	toSerialize["parallelism"] = o.Parallelism
	toSerialize["policies"] = o.Policies
	return toSerialize, nil
}

type NullableDeploySpec struct {
	value *DeploySpec
	isSet bool
}

func (v *NullableDeploySpec) Get() *DeploySpec {
	return v.value
}

func (v *NullableDeploySpec) Set(val *DeploySpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableDeploySpec) IsSet() bool {
	return v.isSet
}

func (v *NullableDeploySpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeploySpec(val *DeploySpec) *NullableDeploySpec {
	return &NullableDeploySpec{value: val, isSet: true}
}

func (v NullableDeploySpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeploySpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const DeployType = "deploy"

func init() {
	sdkcommon.RegisterWorkflowStep(DeployType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(DeployType, FromWorkflowSubStep)
}

type DeployWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties DeploySpec
}

func Deploy(name string) *DeployWorkflowStep {
	d := &DeployWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: DeployType,
	}}
	return d
}

func (d *DeployWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range d.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  d.Base.DependsOn,
		If:         d.Base.If,
		Inputs:     d.Base.Inputs,
		Meta:       d.Base.Meta,
		Name:       d.Base.Name,
		Outputs:    d.Base.Outputs,
		Properties: util.Object2RawExtension(d.Properties),
		SubSteps:   subSteps,
		Timeout:    d.Base.Timeout,
		Type:       DeployType,
	}
	return res
}

func (d *DeployWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*DeployWorkflowStep, error) {
	var properties DeploySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := d.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	d.Base.Name = from.Name
	d.Base.DependsOn = from.DependsOn
	d.Base.Inputs = from.Inputs
	d.Base.Outputs = from.Outputs
	d.Base.If = from.If
	d.Base.Timeout = from.Timeout
	d.Base.Meta = from.Meta
	d.Base.Type = DeployType
	d.Properties = properties
	d.Base.SubSteps = subSteps
	return d, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	d := &DeployWorkflowStep{}
	return d.FromWorkflowStep(from)
}

func (d *DeployWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*DeployWorkflowStep, error) {
	var properties DeploySpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	d.Base.Name = from.Name
	d.Base.DependsOn = from.DependsOn
	d.Base.Inputs = from.Inputs
	d.Base.Outputs = from.Outputs
	d.Base.If = from.If
	d.Base.Timeout = from.Timeout
	d.Base.Meta = from.Meta
	d.Base.Type = DeployType
	d.Properties = properties
	return d, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	d := &DeployWorkflowStep{}
	return d.FromWorkflowSubStep(from)
}

func (d *DeployWorkflowStep) WorkflowStepName() string {
	return d.Base.Name
}

func (d *DeployWorkflowStep) DefType() string {
	return DeployType
}

func (d *DeployWorkflowStep) If(_if string) *DeployWorkflowStep {
	d.Base.If = _if
	return d
}

func (d *DeployWorkflowStep) Alias(alias string) *DeployWorkflowStep {
	d.Base.Meta.Alias = alias
	return d
}

func (d *DeployWorkflowStep) Timeout(timeout string) *DeployWorkflowStep {
	d.Base.Timeout = timeout
	return d
}

func (d *DeployWorkflowStep) DependsOn(dependsOn []string) *DeployWorkflowStep {
	d.Base.DependsOn = dependsOn
	return d
}

func (d *DeployWorkflowStep) Inputs(input common.StepInputs) *DeployWorkflowStep {
	d.Base.Inputs = input
	return d
}

func (d *DeployWorkflowStep) Outputs(output common.StepOutputs) *DeployWorkflowStep {
	d.Base.Outputs = output
	return d
}
