/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package export_data

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/common"
	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the ExportDataSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &ExportDataSpec{}

// ExportDataSpec struct for ExportDataSpec
type ExportDataSpec struct {
	// Specify the data to export
	Data map[string]interface{} `json:"data"`
	// Specify the kind of the export destination
	Kind *string `json:"kind"`
	// Specify the name of the export destination
	Name *string `json:"name,omitempty"`
	// Specify the namespace of the export destination
	Namespace *string `json:"namespace,omitempty"`
	// Specify the topology to export
	Topology *string `json:"topology,omitempty"`
}

// NewExportDataSpecWith instantiates a new ExportDataSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewExportDataSpecWith(data map[string]interface{}, kind string) *ExportDataSpec {
	this := ExportDataSpec{}
	this.Data = data
	this.Kind = &kind
	return &this
}

// NewExportDataSpecWithDefault instantiates a new ExportDataSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportDataSpecWithDefault() *ExportDataSpec {
	this := ExportDataSpec{}
	var kind string = "ConfigMap"
	this.Kind = &kind
	return &this
}

// NewExportDataSpec is short for NewExportDataSpecWithDefault which instantiates a new ExportDataSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportDataSpec() *ExportDataSpec {
	return NewExportDataSpecWithDefault()
}

// NewExportDataSpecEmpty instantiates a new ExportDataSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewExportDataSpecEmpty() *ExportDataSpec {
	this := ExportDataSpec{}
	return &this
}

// NewExportDataSpecs converts a list ExportDataSpec pointers to objects.
// This is helpful when the SetExportDataSpec requires a list of objects
func NewExportDataSpecList(ps ...*ExportDataSpec) []ExportDataSpec {
	objs := []ExportDataSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this ExportDataSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *ExportDataWorkflowStep) Validate() error {
	if o.Properties.Data == nil {
		return errors.New("Data in ExportDataSpec must be set")
	}
	if o.Properties.Kind == nil {
		return errors.New("Kind in ExportDataSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetData returns the Data field value
func (o *ExportDataWorkflowStep) GetData() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Properties.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *ExportDataWorkflowStep) GetDataOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Properties.Data, true
}

// SetData sets field value
func (o *ExportDataWorkflowStep) SetData(v map[string]interface{}) *ExportDataWorkflowStep {
	o.Properties.Data = v
	return o
}

// GetKind returns the Kind field value
func (o *ExportDataWorkflowStep) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *ExportDataWorkflowStep) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Kind, true
}

// SetKind sets field value
func (o *ExportDataWorkflowStep) SetKind(v string) *ExportDataWorkflowStep {
	o.Properties.Kind = &v
	return o
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *ExportDataWorkflowStep) GetName() string {
	if o == nil || utils.IsNil(o.Properties.Name) {
		var ret string
		return ret
	}
	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDataWorkflowStep) GetNameOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Name) {
		return nil, false
	}
	return o.Properties.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *ExportDataWorkflowStep) HasName() bool {
	if o != nil && !utils.IsNil(o.Properties.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the name field.
// Name:  Specify the name of the export destination
func (o *ExportDataWorkflowStep) SetName(v string) *ExportDataWorkflowStep {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *ExportDataWorkflowStep) GetNamespace() string {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		var ret string
		return ret
	}
	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDataWorkflowStep) GetNamespaceOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Namespace) {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *ExportDataWorkflowStep) HasNamespace() bool {
	if o != nil && !utils.IsNil(o.Properties.Namespace) {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the namespace field.
// Namespace:  Specify the namespace of the export destination
func (o *ExportDataWorkflowStep) SetNamespace(v string) *ExportDataWorkflowStep {
	o.Properties.Namespace = &v
	return o
}

// GetTopology returns the Topology field value if set, zero value otherwise.
func (o *ExportDataWorkflowStep) GetTopology() string {
	if o == nil || utils.IsNil(o.Properties.Topology) {
		var ret string
		return ret
	}
	return *o.Properties.Topology
}

// GetTopologyOk returns a tuple with the Topology field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportDataWorkflowStep) GetTopologyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Topology) {
		return nil, false
	}
	return o.Properties.Topology, true
}

// HasTopology returns a boolean if a field has been set.
func (o *ExportDataWorkflowStep) HasTopology() bool {
	if o != nil && !utils.IsNil(o.Properties.Topology) {
		return true
	}

	return false
}

// SetTopology gets a reference to the given string and assigns it to the topology field.
// Topology:  Specify the topology to export
func (o *ExportDataWorkflowStep) SetTopology(v string) *ExportDataWorkflowStep {
	o.Properties.Topology = &v
	return o
}

func (o ExportDataSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExportDataSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["data"] = o.Data
	toSerialize["kind"] = o.Kind
	if !utils.IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !utils.IsNil(o.Namespace) {
		toSerialize["namespace"] = o.Namespace
	}
	if !utils.IsNil(o.Topology) {
		toSerialize["topology"] = o.Topology
	}
	return toSerialize, nil
}

type NullableExportDataSpec struct {
	value *ExportDataSpec
	isSet bool
}

func (v *NullableExportDataSpec) Get() *ExportDataSpec {
	return v.value
}

func (v *NullableExportDataSpec) Set(val *ExportDataSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableExportDataSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableExportDataSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportDataSpec(val *ExportDataSpec) *NullableExportDataSpec {
	return &NullableExportDataSpec{value: val, isSet: true}
}

func (v NullableExportDataSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportDataSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const ExportDataType = "export-data"

func init() {
	sdkcommon.RegisterWorkflowStep(ExportDataType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(ExportDataType, FromWorkflowSubStep)
}

type ExportDataWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties ExportDataSpec
}

func ExportData(name string) *ExportDataWorkflowStep {
	e := &ExportDataWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: ExportDataType,
	}}
	return e
}

func (e *ExportDataWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range e.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  e.Base.DependsOn,
		If:         e.Base.If,
		Inputs:     e.Base.Inputs,
		Meta:       e.Base.Meta,
		Name:       e.Base.Name,
		Outputs:    e.Base.Outputs,
		Properties: util.Object2RawExtension(e.Properties),
		SubSteps:   subSteps,
		Timeout:    e.Base.Timeout,
		Type:       ExportDataType,
	}
	return res
}

func (e *ExportDataWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*ExportDataWorkflowStep, error) {
	var properties ExportDataSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := e.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	e.Base.Name = from.Name
	e.Base.DependsOn = from.DependsOn
	e.Base.Inputs = from.Inputs
	e.Base.Outputs = from.Outputs
	e.Base.If = from.If
	e.Base.Timeout = from.Timeout
	e.Base.Meta = from.Meta
	e.Base.Type = ExportDataType
	e.Properties = properties
	e.Base.SubSteps = subSteps
	return e, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	e := &ExportDataWorkflowStep{}
	return e.FromWorkflowStep(from)
}

func (e *ExportDataWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*ExportDataWorkflowStep, error) {
	var properties ExportDataSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	e.Base.Name = from.Name
	e.Base.DependsOn = from.DependsOn
	e.Base.Inputs = from.Inputs
	e.Base.Outputs = from.Outputs
	e.Base.If = from.If
	e.Base.Timeout = from.Timeout
	e.Base.Meta = from.Meta
	e.Base.Type = ExportDataType
	e.Properties = properties
	return e, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	e := &ExportDataWorkflowStep{}
	return e.FromWorkflowSubStep(from)
}

func (e *ExportDataWorkflowStep) WorkflowStepName() string {
	return e.Base.Name
}

func (e *ExportDataWorkflowStep) DefType() string {
	return ExportDataType
}

func (e *ExportDataWorkflowStep) If(_if string) *ExportDataWorkflowStep {
	e.Base.If = _if
	return e
}

func (e *ExportDataWorkflowStep) Alias(alias string) *ExportDataWorkflowStep {
	e.Base.Meta.Alias = alias
	return e
}

func (e *ExportDataWorkflowStep) Timeout(timeout string) *ExportDataWorkflowStep {
	e.Base.Timeout = timeout
	return e
}

func (e *ExportDataWorkflowStep) DependsOn(dependsOn []string) *ExportDataWorkflowStep {
	e.Base.DependsOn = dependsOn
	return e
}

func (e *ExportDataWorkflowStep) Inputs(input common.StepInputs) *ExportDataWorkflowStep {
	e.Base.Inputs = input
	return e
}

func (e *ExportDataWorkflowStep) Outputs(output common.StepOutputs) *ExportDataWorkflowStep {
	e.Base.Outputs = output
	return e
}
