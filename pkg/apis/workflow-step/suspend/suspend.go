/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package suspend

import (
	"encoding/json"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/v1beta1"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the SuspendSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &SuspendSpec{}

// SuspendSpec struct for SuspendSpec
type SuspendSpec struct {
	// Specify the wait duration time to resume workflow such as \"30s\", \"1min\" or \"2m15s\"
	Duration *string `json:"duration,omitempty"`
	// The suspend message to show
	Message *string `json:"message,omitempty"`
}

// NewSuspendSpecWith instantiates a new SuspendSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewSuspendSpecWith() *SuspendSpec {
	this := SuspendSpec{}
	return &this
}

// NewSuspendSpecWithDefault instantiates a new SuspendSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuspendSpecWithDefault() *SuspendSpec {
	this := SuspendSpec{}
	return &this
}

// NewSuspendSpec is short for NewSuspendSpecWithDefault which instantiates a new SuspendSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSuspendSpec() *SuspendSpec {
	return NewSuspendSpecWithDefault()
}

// NewSuspendSpecEmpty instantiates a new SuspendSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewSuspendSpecEmpty() *SuspendSpec {
	this := SuspendSpec{}
	return &this
}

// NewSuspendSpecs converts a list SuspendSpec pointers to objects.
// This is helpful when the SetSuspendSpec requires a list of objects
func NewSuspendSpecList(ps ...*SuspendSpec) []SuspendSpec {
	objs := []SuspendSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this SuspendSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *SuspendWorkflowStep) Validate() error {
	// validate all nested properties
	return nil
}

// GetDuration returns the Duration field value if set, zero value otherwise.
func (o *SuspendWorkflowStep) GetDuration() string {
	if o == nil || utils.IsNil(o.Properties.Duration) {
		var ret string
		return ret
	}
	return *o.Properties.Duration
}

// GetDurationOk returns a tuple with the Duration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuspendWorkflowStep) GetDurationOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Duration) {
		return nil, false
	}
	return o.Properties.Duration, true
}

// HasDuration returns a boolean if a field has been set.
func (o *SuspendWorkflowStep) HasDuration() bool {
	if o != nil && !utils.IsNil(o.Properties.Duration) {
		return true
	}

	return false
}

// SetDuration gets a reference to the given string and assigns it to the duration field.
// Duration:  Specify the wait duration time to resume workflow such as \"30s\", \"1min\" or \"2m15s\"
func (o *SuspendWorkflowStep) SetDuration(v string) *SuspendWorkflowStep {
	o.Properties.Duration = &v
	return o
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SuspendWorkflowStep) GetMessage() string {
	if o == nil || utils.IsNil(o.Properties.Message) {
		var ret string
		return ret
	}
	return *o.Properties.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SuspendWorkflowStep) GetMessageOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Message) {
		return nil, false
	}
	return o.Properties.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SuspendWorkflowStep) HasMessage() bool {
	if o != nil && !utils.IsNil(o.Properties.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the message field.
// Message:  The suspend message to show
func (o *SuspendWorkflowStep) SetMessage(v string) *SuspendWorkflowStep {
	o.Properties.Message = &v
	return o
}

func (o SuspendSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SuspendSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Duration) {
		toSerialize["duration"] = o.Duration
	}
	if !utils.IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	return toSerialize, nil
}

type NullableSuspendSpec struct {
	value *SuspendSpec
	isSet bool
}

func (v *NullableSuspendSpec) Get() *SuspendSpec {
	return v.value
}

func (v *NullableSuspendSpec) Set(val *SuspendSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableSuspendSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableSuspendSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSuspendSpec(val *SuspendSpec) *NullableSuspendSpec {
	return &NullableSuspendSpec{value: val, isSet: true}
}

func (v NullableSuspendSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSuspendSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const SuspendType = "suspend"

func init() {
	sdkcommon.RegisterWorkflowStep(SuspendType, FromWorkflowStep)
	sdkcommon.RegisterWorkflowSubStep(SuspendType, FromWorkflowSubStep)
}

type SuspendWorkflowStep struct {
	Base       apis.WorkflowStepBase
	Properties SuspendSpec
}

func Suspend(name string) *SuspendWorkflowStep {
	s := &SuspendWorkflowStep{Base: apis.WorkflowStepBase{
		Name: name,
		Type: SuspendType,
	}}
	return s
}

func (s *SuspendWorkflowStep) Build() v1beta1.WorkflowStep {
	_subSteps := make([]v1beta1.WorkflowStep, 0)
	for _, subStep := range s.Base.SubSteps {
		_subSteps = append(_subSteps, subStep.Build())
	}
	subSteps := make([]common.WorkflowSubStep, 0)
	for _, _s := range _subSteps {
		subSteps = append(subSteps, common.WorkflowSubStep{Name: _s.Name, DependsOn: _s.DependsOn, Inputs: _s.Inputs, Outputs: _s.Outputs, If: _s.If, Timeout: _s.Timeout, Meta: _s.Meta, Properties: _s.Properties, Type: _s.Type})
	}
	res := v1beta1.WorkflowStep{
		DependsOn:  s.Base.DependsOn,
		If:         s.Base.If,
		Inputs:     s.Base.Inputs,
		Meta:       s.Base.Meta,
		Name:       s.Base.Name,
		Outputs:    s.Base.Outputs,
		Properties: util.Object2RawExtension(s.Properties),
		SubSteps:   subSteps,
		Timeout:    s.Base.Timeout,
		Type:       SuspendType,
	}
	return res
}

func (s *SuspendWorkflowStep) FromWorkflowStep(from v1beta1.WorkflowStep) (*SuspendWorkflowStep, error) {
	var properties SuspendSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	subSteps := make([]apis.WorkflowStep, 0)
	for _, _s := range from.SubSteps {
		subStep, err := s.FromWorkflowSubStep(_s)
		if err != nil {
			return nil, err
		}
		subSteps = append(subSteps, subStep)
	}
	s.Base.Name = from.Name
	s.Base.DependsOn = from.DependsOn
	s.Base.Inputs = from.Inputs
	s.Base.Outputs = from.Outputs
	s.Base.If = from.If
	s.Base.Timeout = from.Timeout
	s.Base.Meta = from.Meta
	s.Base.Type = SuspendType
	s.Properties = properties
	s.Base.SubSteps = subSteps
	return s, nil
}

func FromWorkflowStep(from v1beta1.WorkflowStep) (apis.WorkflowStep, error) {
	s := &SuspendWorkflowStep{}
	return s.FromWorkflowStep(from)
}

func (s *SuspendWorkflowStep) FromWorkflowSubStep(from common.WorkflowSubStep) (*SuspendWorkflowStep, error) {
	var properties SuspendSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	s.Base.Name = from.Name
	s.Base.DependsOn = from.DependsOn
	s.Base.Inputs = from.Inputs
	s.Base.Outputs = from.Outputs
	s.Base.If = from.If
	s.Base.Timeout = from.Timeout
	s.Base.Meta = from.Meta
	s.Base.Type = SuspendType
	s.Properties = properties
	return s, nil
}

func FromWorkflowSubStep(from common.WorkflowSubStep) (apis.WorkflowStep, error) {
	s := &SuspendWorkflowStep{}
	return s.FromWorkflowSubStep(from)
}

func (s *SuspendWorkflowStep) WorkflowStepName() string {
	return s.Base.Name
}

func (s *SuspendWorkflowStep) DefType() string {
	return SuspendType
}

func (s *SuspendWorkflowStep) If(_if string) *SuspendWorkflowStep {
	s.Base.If = _if
	return s
}

func (s *SuspendWorkflowStep) Alias(alias string) *SuspendWorkflowStep {
	s.Base.Meta.Alias = alias
	return s
}

func (s *SuspendWorkflowStep) Timeout(timeout string) *SuspendWorkflowStep {
	s.Base.Timeout = timeout
	return s
}

func (s *SuspendWorkflowStep) DependsOn(dependsOn []string) *SuspendWorkflowStep {
	s.Base.DependsOn = dependsOn
	return s
}

func (s *SuspendWorkflowStep) Inputs(input common.StepInputs) *SuspendWorkflowStep {
	s.Base.Inputs = input
	return s
}

func (s *SuspendWorkflowStep) Outputs(output common.StepOutputs) *SuspendWorkflowStep {
	s.Base.Outputs = output
	return s
}