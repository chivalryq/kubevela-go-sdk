/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package worker

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the WorkerSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &WorkerSpec{}

// WorkerSpec struct for WorkerSpec
type WorkerSpec struct {
	// Commands to run in the container
	Cmd []string `json:"cmd,omitempty"`
	// Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	Cpu *string `json:"cpu,omitempty"`
	// Define arguments by using environment variables
	Env []Env `json:"env,omitempty"`
	// Which image would you like to use for your service +short=i
	Image *string `json:"image"`
	// Specify image pull policy for your service
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Specify image pull secrets for your service
	ImagePullSecrets []string     `json:"imagePullSecrets,omitempty"`
	LivenessProbe    *HealthProbe `json:"livenessProbe,omitempty"`
	// Specifies the attributes of the memory resource required for the container.
	Memory         *string       `json:"memory,omitempty"`
	ReadinessProbe *HealthProbe  `json:"readinessProbe,omitempty"`
	VolumeMounts   *VolumeMounts `json:"volumeMounts,omitempty"`
	// Deprecated field, use volumeMounts instead.
	Volumes []Volumes `json:"volumes,omitempty"`
}

// NewWorkerSpecWith instantiates a new WorkerSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewWorkerSpecWith(image string) *WorkerSpec {
	this := WorkerSpec{}
	this.Image = &image
	return &this
}

// NewWorkerSpecWithDefault instantiates a new WorkerSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerSpecWithDefault() *WorkerSpec {
	this := WorkerSpec{}
	return &this
}

// NewWorkerSpec is short for NewWorkerSpecWithDefault which instantiates a new WorkerSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkerSpec() *WorkerSpec {
	return NewWorkerSpecWithDefault()
}

// NewWorkerSpecEmpty instantiates a new WorkerSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewWorkerSpecEmpty() *WorkerSpec {
	this := WorkerSpec{}
	return &this
}

// NewWorkerSpecs converts a list WorkerSpec pointers to objects.
// This is helpful when the SetWorkerSpec requires a list of objects
func NewWorkerSpecList(ps ...*WorkerSpec) []WorkerSpec {
	objs := []WorkerSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this WorkerSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *WorkerComponent) Validate() error {
	if o.Properties.Image == nil {
		return errors.New("Image in WorkerSpec must be set")
	}
	// validate all nested properties
	if o.Properties.LivenessProbe != nil {
		if err := o.Properties.LivenessProbe.Validate(); err != nil {
			return err
		}
	}
	if o.Properties.ReadinessProbe != nil {
		if err := o.Properties.ReadinessProbe.Validate(); err != nil {
			return err
		}
	}
	if o.Properties.VolumeMounts != nil {
		if err := o.Properties.VolumeMounts.Validate(); err != nil {
			return err
		}
	}

	for i, v := range o.Base.Traits {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("traits[%d] %s in %s component is invalid: %w", i, v.DefType(), WorkerType, err)
		}
	}
	return nil
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *WorkerComponent) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		var ret []string
		return ret
	}
	return o.Properties.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		return nil, false
	}
	return o.Properties.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *WorkerComponent) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// Cmd:  Commands to run in the container
func (o *WorkerComponent) SetCmd(v []string) *WorkerComponent {
	o.Properties.Cmd = v
	return o
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *WorkerComponent) GetCpu() string {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		var ret string
		return ret
	}
	return *o.Properties.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		return nil, false
	}
	return o.Properties.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *WorkerComponent) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the cpu field.
// Cpu:  Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (o *WorkerComponent) SetCpu(v string) *WorkerComponent {
	o.Properties.Cpu = &v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *WorkerComponent) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret []Env
		return ret
	}
	return o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *WorkerComponent) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// Env:  Define arguments by using environment variables
func (o *WorkerComponent) SetEnv(v []Env) *WorkerComponent {
	o.Properties.Env = v
	return o
}

// GetImage returns the Image field value
func (o *WorkerComponent) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Image, true
}

// SetImage sets field value
func (o *WorkerComponent) SetImage(v string) *WorkerComponent {
	o.Properties.Image = &v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *WorkerComponent) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		var ret string
		return ret
	}
	return *o.Properties.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		return nil, false
	}
	return o.Properties.ImagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *WorkerComponent) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullPolicy) {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// ImagePullPolicy:  Specify image pull policy for your service
func (o *WorkerComponent) SetImagePullPolicy(v string) *WorkerComponent {
	o.Properties.ImagePullPolicy = &v
	return o
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *WorkerComponent) GetImagePullSecrets() []string {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		var ret []string
		return ret
	}
	return o.Properties.ImagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetImagePullSecretsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		return nil, false
	}
	return o.Properties.ImagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *WorkerComponent) HasImagePullSecrets() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullSecrets) {
		return true
	}

	return false
}

// SetImagePullSecrets gets a reference to the given []string and assigns it to the imagePullSecrets field.
// ImagePullSecrets:  Specify image pull secrets for your service
func (o *WorkerComponent) SetImagePullSecrets(v []string) *WorkerComponent {
	o.Properties.ImagePullSecrets = v
	return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *WorkerComponent) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		return nil, false
	}
	return o.Properties.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *WorkerComponent) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// LivenessProbe:
func (o *WorkerComponent) SetLivenessProbe(v HealthProbe) *WorkerComponent {
	o.Properties.LivenessProbe = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *WorkerComponent) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		var ret string
		return ret
	}
	return *o.Properties.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		return nil, false
	}
	return o.Properties.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *WorkerComponent) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// Memory:  Specifies the attributes of the memory resource required for the container.
func (o *WorkerComponent) SetMemory(v string) *WorkerComponent {
	o.Properties.Memory = &v
	return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *WorkerComponent) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		return nil, false
	}
	return o.Properties.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *WorkerComponent) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// ReadinessProbe:
func (o *WorkerComponent) SetReadinessProbe(v HealthProbe) *WorkerComponent {
	o.Properties.ReadinessProbe = &v
	return o
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *WorkerComponent) GetVolumeMounts() VolumeMounts {
	if o == nil || utils.IsNil(o.Properties.VolumeMounts) {
		var ret VolumeMounts
		return ret
	}
	return *o.Properties.VolumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetVolumeMountsOk() (*VolumeMounts, bool) {
	if o == nil || utils.IsNil(o.Properties.VolumeMounts) {
		return nil, false
	}
	return o.Properties.VolumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *WorkerComponent) HasVolumeMounts() bool {
	if o != nil && !utils.IsNil(o.Properties.VolumeMounts) {
		return true
	}

	return false
}

// SetVolumeMounts gets a reference to the given VolumeMounts and assigns it to the volumeMounts field.
// VolumeMounts:
func (o *WorkerComponent) SetVolumeMounts(v VolumeMounts) *WorkerComponent {
	o.Properties.VolumeMounts = &v
	return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *WorkerComponent) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		var ret []Volumes
		return ret
	}
	return o.Properties.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkerComponent) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		return nil, false
	}
	return o.Properties.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *WorkerComponent) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.Properties.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volumes and assigns it to the volumes field.
// Volumes:  Deprecated field, use volumeMounts instead.
func (o *WorkerComponent) SetVolumes(v []Volumes) *WorkerComponent {
	o.Properties.Volumes = v
	return o
}

func (o WorkerSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkerSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	if !utils.IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	toSerialize["image"] = o.Image
	if !utils.IsNil(o.ImagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	}
	if !utils.IsNil(o.ImagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.ImagePullSecrets
	}
	if !utils.IsNil(o.LivenessProbe) {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	if !utils.IsNil(o.Memory) {
		toSerialize["memory"] = o.Memory
	}
	if !utils.IsNil(o.ReadinessProbe) {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	if !utils.IsNil(o.VolumeMounts) {
		toSerialize["volumeMounts"] = o.VolumeMounts
	}
	if !utils.IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableWorkerSpec struct {
	value *WorkerSpec
	isSet bool
}

func (v *NullableWorkerSpec) Get() *WorkerSpec {
	return v.value
}

func (v *NullableWorkerSpec) Set(val *WorkerSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableWorkerSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkerSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkerSpec(val *WorkerSpec) *NullableWorkerSpec {
	return &NullableWorkerSpec{value: val, isSet: true}
}

func (v NullableWorkerSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkerSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const WorkerType = "worker"

func init() {
	sdkcommon.RegisterComponent(WorkerType, FromComponent)
}

type WorkerComponent struct {
	Base       apis.ComponentBase
	Properties WorkerSpec
}

func Worker(name string) *WorkerComponent {
	w := &WorkerComponent{Base: apis.ComponentBase{
		Name: name,
		Type: WorkerType,
	}}
	return w
}

func (w *WorkerComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range w.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  w.Base.DependsOn,
		Inputs:     w.Base.Inputs,
		Name:       w.Base.Name,
		Outputs:    w.Base.Outputs,
		Properties: util.Object2RawExtension(w.Properties),
		Traits:     traits,
		Type:       WorkerType,
	}
	return res
}

func (w *WorkerComponent) FromComponent(from common.ApplicationComponent) (*WorkerComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		w.Base.Traits = append(w.Base.Traits, _t)
	}
	var properties WorkerSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	w.Base.Name = from.Name
	w.Base.DependsOn = from.DependsOn
	w.Base.Inputs = from.Inputs
	w.Base.Outputs = from.Outputs
	w.Base.Type = WorkerType
	w.Properties = properties
	return w, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	w := &WorkerComponent{}
	return w.FromComponent(from)
}

func (w *WorkerComponent) SetTraits(traits ...apis.Trait) *WorkerComponent {
	for _, addTrait := range traits {
		found := false
		for i, _t := range w.Base.Traits {
			if _t.DefType() == addTrait.DefType() {
				w.Base.Traits[i] = addTrait
				found = true
				break
			}
			if !found {
				w.Base.Traits = append(w.Base.Traits, addTrait)
			}
		}
	}
	return w
}

func (w *WorkerComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range w.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (w *WorkerComponent) GetAllTraits() []apis.Trait {
	return w.Base.Traits
}

func (w *WorkerComponent) ComponentName() string {
	return w.Base.Name
}

func (w *WorkerComponent) DefType() string {
	return WorkerType
}

func (w *WorkerComponent) DependsOn(dependsOn []string) *WorkerComponent {
	w.Base.DependsOn = dependsOn
	return w
}

func (w *WorkerComponent) Inputs(input common.StepInputs) *WorkerComponent {
	w.Base.Inputs = input
	return w
}

func (w *WorkerComponent) Outputs(output common.StepOutputs) *WorkerComponent {
	w.Base.Outputs = output
	return w
}

func (w *WorkerComponent) AddDependsOn(dependsOn string) *WorkerComponent {
	w.Base.DependsOn = append(w.Base.DependsOn, dependsOn)
	return w
}