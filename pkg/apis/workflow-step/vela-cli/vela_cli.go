/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package vela_cli

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

// checks if the VelaCliSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VelaCliSpec{}

// VelaCliSpec struct for VelaCliSpec
type VelaCliSpec struct {
	// Specify the name of the addon.
	AddonName *string `json:"addonName"`
	// Specify the vela command
	Command []string `json:"command"`
	// Specify the image
	Image *string `json:"image"`
	// specify serviceAccountName want to use
	ServiceAccountName *string  `json:"serviceAccountName"`
	Storage            *Storage `json:"storage,omitempty"`
}

// NewVelaCliSpecWith instantiates a new VelaCliSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewVelaCliSpecWith(addonName string, command []string, image string, serviceAccountName string) *VelaCliSpec {
	this := VelaCliSpec{}
	this.AddonName = &addonName
	this.Command = command
	this.Image = &image
	this.ServiceAccountName = &serviceAccountName
	return &this
}

// NewVelaCliSpecWithDefault instantiates a new VelaCliSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelaCliSpecWithDefault() *VelaCliSpec {
	this := VelaCliSpec{}
	var image string = "oamdev/vela-cli:v1.6.4"
	this.Image = &image
	var serviceAccountName string = "kubevela-vela-core"
	this.ServiceAccountName = &serviceAccountName
	return &this
}

// NewVelaCliSpec is short for NewVelaCliSpecWithDefault which instantiates a new VelaCliSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVelaCliSpec() *VelaCliSpec {
	return NewVelaCliSpecWithDefault()
}

// NewVelaCliSpecEmpty instantiates a new VelaCliSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewVelaCliSpecEmpty() *VelaCliSpec {
	this := VelaCliSpec{}
	return &this
}

// NewVelaCliSpecs converts a list VelaCliSpec pointers to objects.
// This is helpful when the SetVelaCliSpec requires a list of objects
func NewVelaCliSpecList(ps ...*VelaCliSpec) []VelaCliSpec {
	objs := []VelaCliSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this VelaCliSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *VelaCliWorkflowStep) Validate() error {
	if o.Properties.AddonName == nil {
		return errors.New("AddonName in VelaCliSpec must be set")
	}
	if o.Properties.Command == nil {
		return errors.New("Command in VelaCliSpec must be set")
	}
	if o.Properties.Image == nil {
		return errors.New("Image in VelaCliSpec must be set")
	}
	if o.Properties.ServiceAccountName == nil {
		return errors.New("ServiceAccountName in VelaCliSpec must be set")
	}
	// validate all nested properties
	if o.Properties.Storage != nil {
		if err := o.Properties.Storage.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetAddonName returns the AddonName field value
func (o *VelaCliWorkflowStep) GetAddonName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.AddonName
}

// GetAddonNameOk returns a tuple with the AddonName field value
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetAddonNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.AddonName, true
}

// SetAddonName sets field value
func (o *VelaCliWorkflowStep) SetAddonName(v string) *VelaCliWorkflowStep {
	o.Properties.AddonName = &v
	return o
}

// GetCommand returns the Command field value
func (o *VelaCliWorkflowStep) GetCommand() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties.Command
}

// GetCommandOk returns a tuple with the Command field value
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetCommandOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Command, true
}

// SetCommand sets field value
func (o *VelaCliWorkflowStep) SetCommand(v []string) *VelaCliWorkflowStep {
	o.Properties.Command = v
	return o
}

// GetImage returns the Image field value
func (o *VelaCliWorkflowStep) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Image, true
}

// SetImage sets field value
func (o *VelaCliWorkflowStep) SetImage(v string) *VelaCliWorkflowStep {
	o.Properties.Image = &v
	return o
}

// GetServiceAccountName returns the ServiceAccountName field value
func (o *VelaCliWorkflowStep) GetServiceAccountName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.ServiceAccountName
}

// GetServiceAccountNameOk returns a tuple with the ServiceAccountName field value
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetServiceAccountNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ServiceAccountName, true
}

// SetServiceAccountName sets field value
func (o *VelaCliWorkflowStep) SetServiceAccountName(v string) *VelaCliWorkflowStep {
	o.Properties.ServiceAccountName = &v
	return o
}

// GetStorage returns the Storage field value if set, zero value otherwise.
func (o *VelaCliWorkflowStep) GetStorage() Storage {
	if o == nil || utils.IsNil(o.Properties.Storage) {
		var ret Storage
		return ret
	}
	return *o.Properties.Storage
}

// GetStorageOk returns a tuple with the Storage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VelaCliWorkflowStep) GetStorageOk() (*Storage, bool) {
	if o == nil || utils.IsNil(o.Properties.Storage) {
		return nil, false
	}
	return o.Properties.Storage, true
}

// HasStorage returns a boolean if a field has been set.
func (o *VelaCliWorkflowStep) HasStorage() bool {
	if o != nil && !utils.IsNil(o.Properties.Storage) {
		return true
	}

	return false
}

// SetStorage gets a reference to the given Storage and assigns it to the storage field.
// Storage:
func (o *VelaCliWorkflowStep) SetStorage(v Storage) *VelaCliWorkflowStep {
	o.Properties.Storage = &v
	return o
}

func (o VelaCliSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VelaCliSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["addonName"] = o.AddonName
	toSerialize["command"] = o.Command
	toSerialize["image"] = o.Image
	toSerialize["serviceAccountName"] = o.ServiceAccountName
	if !utils.IsNil(o.Storage) {
		toSerialize["storage"] = o.Storage
	}
	return toSerialize, nil
}

type NullableVelaCliSpec struct {
	value *VelaCliSpec
	isSet bool
}

func (v *NullableVelaCliSpec) Get() *VelaCliSpec {
	return v.value
}

func (v *NullableVelaCliSpec) Set(val *VelaCliSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableVelaCliSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableVelaCliSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVelaCliSpec(val *VelaCliSpec) *NullableVelaCliSpec {
	return &NullableVelaCliSpec{value: val, isSet: true}
}

func (v NullableVelaCliSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVelaCliSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const VelaCliType = "vela-cli"

func init() {
	sdkcommon.RegisterWorkflowStep(VelaCliType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(VelaCliType, FromWorkflowSubStep)
}

type VelaCliWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties VelaCliSpec
}

func VelaCli(name string) *VelaCliWorkflowStep {
	v := &VelaCliWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: VelaCliType,
	}}
	return v
}

func (v *VelaCliWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range v.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  v.Base.DependsOn,
		If:         v.Base.If,
		Inputs:     v.Base.Inputs,
		Meta:       v.Base.Meta,
		Name:       v.Base.Name,
		Outputs:    v.Base.Outputs,
		Properties: util.Object2RawExtension(v.Properties),
		SubSteps:   subSteps,
		Timeout:    v.Base.Timeout,
		Type:       VelaCliType,
	}
	return res
}

func (v *VelaCliWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*VelaCliWorkflowStep, error) {
	var properties VelaCliSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := v.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	v.Base.Name = from.Name
	v.Base.DependsOn = from.DependsOn
	v.Base.Inputs = from.Inputs
	v.Base.Outputs = from.Outputs
	v.Base.If = from.If
	v.Base.Timeout = from.Timeout
	v.Base.Meta = from.Meta
	v.Base.Type = VelaCliType
	v.Properties = properties
	v.Base.SubSteps = subSteps
	return v, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	v := &VelaCliWorkflowStep{}
	return v.FromWorkflowStep(from)
}

func (v *VelaCliWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*VelaCliWorkflowStep, error) {
	var properties VelaCliSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	v.Base.Name = from.Name
	v.Base.DependsOn = from.DependsOn
	v.Base.Inputs = from.Inputs
	v.Base.Outputs = from.Outputs
	v.Base.If = from.If
	v.Base.Timeout = from.Timeout
	v.Base.Meta = from.Meta
	v.Base.Type = VelaCliType
	v.Properties = properties
	return v, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	v := &VelaCliWorkflowStep{}
	return v.FromWorkflowSubStep(from)
}

func (v *VelaCliWorkflowStep) WorkflowStepName() string {
	return v.Base.Name
}

func (v *VelaCliWorkflowStep) DefType() string {
	return VelaCliType
}

func (v *VelaCliWorkflowStep) If(_if string) *VelaCliWorkflowStep {
	v.Base.If = _if
	return v
}

func (v *VelaCliWorkflowStep) Alias(alias string) *VelaCliWorkflowStep {
	v.Base.Meta.Alias = alias
	return v
}

func (v *VelaCliWorkflowStep) Timeout(timeout string) *VelaCliWorkflowStep {
	v.Base.Timeout = timeout
	return v
}

func (v *VelaCliWorkflowStep) DependsOn(dependsOn []string) *VelaCliWorkflowStep {
	v.Base.DependsOn = dependsOn
	return v
}

func (v *VelaCliWorkflowStep) Inputs(input common.StepInputs) *VelaCliWorkflowStep {
	v.Base.Inputs = input
	return v
}

func (v *VelaCliWorkflowStep) Outputs(output common.StepOutputs) *VelaCliWorkflowStep {
	v.Base.Outputs = output
	return v
}