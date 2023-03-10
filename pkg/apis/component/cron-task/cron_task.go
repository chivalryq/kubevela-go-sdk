/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package cron_task

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

// checks if the CronTaskSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &CronTaskSpec{}

// CronTaskSpec struct for CronTaskSpec
type CronTaskSpec struct {
	// The duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it
	ActiveDeadlineSeconds *int32 `json:"activeDeadlineSeconds,omitempty"`
	// Specify the annotations in the workload
	Annotations map[string]string `json:"annotations,omitempty"`
	// The number of retries before marking this job failed
	BackoffLimit *int32 `json:"backoffLimit"`
	// Commands to run in the container
	Cmd []string `json:"cmd,omitempty"`
	// Specifies how to treat concurrent executions of a Job
	ConcurrencyPolicy *string `json:"concurrencyPolicy"`
	// Specify number of tasks to run in parallel +short=c
	Count *int32 `json:"count"`
	// Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
	Cpu *string `json:"cpu,omitempty"`
	// Define arguments by using environment variables
	Env []Env `json:"env,omitempty"`
	// The number of failed finished jobs to retain
	FailedJobsHistoryLimit *int32 `json:"failedJobsHistoryLimit"`
	// An optional list of hosts and IPs that will be injected into the pod's hosts file
	HostAliases []HostAliases `json:"hostAliases,omitempty"`
	// Which image would you like to use for your service +short=i
	Image *string `json:"image"`
	// Specify image pull policy for your service
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	// Specify image pull secrets for your service
	ImagePullSecrets []string `json:"imagePullSecrets,omitempty"`
	// Specify the labels in the workload
	Labels        map[string]string `json:"labels,omitempty"`
	LivenessProbe *HealthProbe      `json:"livenessProbe,omitempty"`
	// Specifies the attributes of the memory resource required for the container.
	Memory         *string      `json:"memory,omitempty"`
	ReadinessProbe *HealthProbe `json:"readinessProbe,omitempty"`
	// Define the job restart policy, the value can only be Never or OnFailure. By default, it's Never.
	Restart *string `json:"restart"`
	// Specify the schedule in Cron format, see https://en.wikipedia.org/wiki/Cron
	Schedule *string `json:"schedule"`
	// Specify deadline in seconds for starting the job if it misses scheduled
	StartingDeadlineSeconds *int32 `json:"startingDeadlineSeconds,omitempty"`
	// The number of successful finished jobs to retain
	SuccessfulJobsHistoryLimit *int32 `json:"successfulJobsHistoryLimit"`
	// suspend subsequent executions
	Suspend *bool `json:"suspend"`
	// Limits the lifetime of a Job that has finished
	TtlSecondsAfterFinished *int32 `json:"ttlSecondsAfterFinished,omitempty"`
	// Declare volumes and volumeMounts
	Volumes []Volumes `json:"volumes,omitempty"`
}

// NewCronTaskSpecWith instantiates a new CronTaskSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewCronTaskSpecWith(backoffLimit int32, concurrencyPolicy string, count int32, failedJobsHistoryLimit int32, image string, restart string, schedule string, successfulJobsHistoryLimit int32, suspend bool) *CronTaskSpec {
	this := CronTaskSpec{}
	this.BackoffLimit = &backoffLimit
	this.ConcurrencyPolicy = &concurrencyPolicy
	this.Count = &count
	this.FailedJobsHistoryLimit = &failedJobsHistoryLimit
	this.Image = &image
	this.Restart = &restart
	this.Schedule = &schedule
	this.SuccessfulJobsHistoryLimit = &successfulJobsHistoryLimit
	this.Suspend = &suspend
	return &this
}

// NewCronTaskSpecWithDefault instantiates a new CronTaskSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronTaskSpecWithDefault() *CronTaskSpec {
	this := CronTaskSpec{}
	var backoffLimit int32 = 6
	this.BackoffLimit = &backoffLimit
	var concurrencyPolicy string = "Allow"
	this.ConcurrencyPolicy = &concurrencyPolicy
	var count int32 = 1
	this.Count = &count
	var failedJobsHistoryLimit int32 = 1
	this.FailedJobsHistoryLimit = &failedJobsHistoryLimit
	var restart string = "Never"
	this.Restart = &restart
	var successfulJobsHistoryLimit int32 = 3
	this.SuccessfulJobsHistoryLimit = &successfulJobsHistoryLimit
	var suspend bool = false
	this.Suspend = &suspend
	return &this
}

// NewCronTaskSpec is short for NewCronTaskSpecWithDefault which instantiates a new CronTaskSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCronTaskSpec() *CronTaskSpec {
	return NewCronTaskSpecWithDefault()
}

// NewCronTaskSpecEmpty instantiates a new CronTaskSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewCronTaskSpecEmpty() *CronTaskSpec {
	this := CronTaskSpec{}
	return &this
}

// NewCronTaskSpecs converts a list CronTaskSpec pointers to objects.
// This is helpful when the SetCronTaskSpec requires a list of objects
func NewCronTaskSpecList(ps ...*CronTaskSpec) []CronTaskSpec {
	objs := []CronTaskSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this CronTaskSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *CronTaskComponent) Validate() error {
	if o.Properties.BackoffLimit == nil {
		return errors.New("BackoffLimit in CronTaskSpec must be set")
	}
	if o.Properties.ConcurrencyPolicy == nil {
		return errors.New("ConcurrencyPolicy in CronTaskSpec must be set")
	}
	if o.Properties.Count == nil {
		return errors.New("Count in CronTaskSpec must be set")
	}
	if o.Properties.FailedJobsHistoryLimit == nil {
		return errors.New("FailedJobsHistoryLimit in CronTaskSpec must be set")
	}
	if o.Properties.Image == nil {
		return errors.New("Image in CronTaskSpec must be set")
	}
	if o.Properties.Restart == nil {
		return errors.New("Restart in CronTaskSpec must be set")
	}
	if o.Properties.Schedule == nil {
		return errors.New("Schedule in CronTaskSpec must be set")
	}
	if o.Properties.SuccessfulJobsHistoryLimit == nil {
		return errors.New("SuccessfulJobsHistoryLimit in CronTaskSpec must be set")
	}
	if o.Properties.Suspend == nil {
		return errors.New("Suspend in CronTaskSpec must be set")
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

	for i, v := range o.Base.Traits {
		if err := v.Validate(); err != nil {
			return fmt.Errorf("traits[%d] %s in %s component is invalid: %w", i, v.DefType(), CronTaskType, err)
		}
	}
	return nil
}

// GetActiveDeadlineSeconds returns the ActiveDeadlineSeconds field value if set, zero value otherwise.
func (o *CronTaskComponent) GetActiveDeadlineSeconds() int32 {
	if o == nil || utils.IsNil(o.Properties.ActiveDeadlineSeconds) {
		var ret int32
		return ret
	}
	return *o.Properties.ActiveDeadlineSeconds
}

// GetActiveDeadlineSecondsOk returns a tuple with the ActiveDeadlineSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetActiveDeadlineSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.ActiveDeadlineSeconds) {
		return nil, false
	}
	return o.Properties.ActiveDeadlineSeconds, true
}

// HasActiveDeadlineSeconds returns a boolean if a field has been set.
func (o *CronTaskComponent) HasActiveDeadlineSeconds() bool {
	if o != nil && !utils.IsNil(o.Properties.ActiveDeadlineSeconds) {
		return true
	}

	return false
}

// SetActiveDeadlineSeconds gets a reference to the given int32 and assigns it to the activeDeadlineSeconds field.
// ActiveDeadlineSeconds:  The duration in seconds relative to the startTime that the job may be continuously active before the system tries to terminate it
func (o *CronTaskComponent) SetActiveDeadlineSeconds(v int32) *CronTaskComponent {
	o.Properties.ActiveDeadlineSeconds = &v
	return o
}

// GetAnnotations returns the Annotations field value if set, zero value otherwise.
func (o *CronTaskComponent) GetAnnotations() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Annotations) {
		var ret map[string]string
		return ret
	}
	return o.Properties.Annotations
}

// GetAnnotationsOk returns a tuple with the Annotations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetAnnotationsOk() (map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Annotations) {
		return nil, false
	}
	return o.Properties.Annotations, true
}

// HasAnnotations returns a boolean if a field has been set.
func (o *CronTaskComponent) HasAnnotations() bool {
	if o != nil && !utils.IsNil(o.Properties.Annotations) {
		return true
	}

	return false
}

// SetAnnotations gets a reference to the given map[string]string and assigns it to the annotations field.
// Annotations:  Specify the annotations in the workload
func (o *CronTaskComponent) SetAnnotations(v map[string]string) *CronTaskComponent {
	o.Properties.Annotations = v
	return o
}

// GetBackoffLimit returns the BackoffLimit field value
func (o *CronTaskComponent) GetBackoffLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.BackoffLimit
}

// GetBackoffLimitOk returns a tuple with the BackoffLimit field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetBackoffLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.BackoffLimit, true
}

// SetBackoffLimit sets field value
func (o *CronTaskComponent) SetBackoffLimit(v int32) *CronTaskComponent {
	o.Properties.BackoffLimit = &v
	return o
}

// GetCmd returns the Cmd field value if set, zero value otherwise.
func (o *CronTaskComponent) GetCmd() []string {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		var ret []string
		return ret
	}
	return o.Properties.Cmd
}

// GetCmdOk returns a tuple with the Cmd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetCmdOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cmd) {
		return nil, false
	}
	return o.Properties.Cmd, true
}

// HasCmd returns a boolean if a field has been set.
func (o *CronTaskComponent) HasCmd() bool {
	if o != nil && !utils.IsNil(o.Properties.Cmd) {
		return true
	}

	return false
}

// SetCmd gets a reference to the given []string and assigns it to the cmd field.
// Cmd:  Commands to run in the container
func (o *CronTaskComponent) SetCmd(v []string) *CronTaskComponent {
	o.Properties.Cmd = v
	return o
}

// GetConcurrencyPolicy returns the ConcurrencyPolicy field value
func (o *CronTaskComponent) GetConcurrencyPolicy() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.ConcurrencyPolicy
}

// GetConcurrencyPolicyOk returns a tuple with the ConcurrencyPolicy field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetConcurrencyPolicyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ConcurrencyPolicy, true
}

// SetConcurrencyPolicy sets field value
func (o *CronTaskComponent) SetConcurrencyPolicy(v string) *CronTaskComponent {
	o.Properties.ConcurrencyPolicy = &v
	return o
}

// GetCount returns the Count field value
func (o *CronTaskComponent) GetCount() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.Count
}

// GetCountOk returns a tuple with the Count field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetCountOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Count, true
}

// SetCount sets field value
func (o *CronTaskComponent) SetCount(v int32) *CronTaskComponent {
	o.Properties.Count = &v
	return o
}

// GetCpu returns the Cpu field value if set, zero value otherwise.
func (o *CronTaskComponent) GetCpu() string {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		var ret string
		return ret
	}
	return *o.Properties.Cpu
}

// GetCpuOk returns a tuple with the Cpu field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetCpuOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Cpu) {
		return nil, false
	}
	return o.Properties.Cpu, true
}

// HasCpu returns a boolean if a field has been set.
func (o *CronTaskComponent) HasCpu() bool {
	if o != nil && !utils.IsNil(o.Properties.Cpu) {
		return true
	}

	return false
}

// SetCpu gets a reference to the given string and assigns it to the cpu field.
// Cpu:  Number of CPU units for the service, like `0.5` (0.5 CPU core), `1` (1 CPU core)
func (o *CronTaskComponent) SetCpu(v string) *CronTaskComponent {
	o.Properties.Cpu = &v
	return o
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *CronTaskComponent) GetEnv() []Env {
	if o == nil || utils.IsNil(o.Properties.Env) {
		var ret []Env
		return ret
	}
	return o.Properties.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetEnvOk() ([]Env, bool) {
	if o == nil || utils.IsNil(o.Properties.Env) {
		return nil, false
	}
	return o.Properties.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *CronTaskComponent) HasEnv() bool {
	if o != nil && !utils.IsNil(o.Properties.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []Env and assigns it to the env field.
// Env:  Define arguments by using environment variables
func (o *CronTaskComponent) SetEnv(v []Env) *CronTaskComponent {
	o.Properties.Env = v
	return o
}

// GetFailedJobsHistoryLimit returns the FailedJobsHistoryLimit field value
func (o *CronTaskComponent) GetFailedJobsHistoryLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.FailedJobsHistoryLimit
}

// GetFailedJobsHistoryLimitOk returns a tuple with the FailedJobsHistoryLimit field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetFailedJobsHistoryLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.FailedJobsHistoryLimit, true
}

// SetFailedJobsHistoryLimit sets field value
func (o *CronTaskComponent) SetFailedJobsHistoryLimit(v int32) *CronTaskComponent {
	o.Properties.FailedJobsHistoryLimit = &v
	return o
}

// GetHostAliases returns the HostAliases field value if set, zero value otherwise.
func (o *CronTaskComponent) GetHostAliases() []HostAliases {
	if o == nil || utils.IsNil(o.Properties.HostAliases) {
		var ret []HostAliases
		return ret
	}
	return o.Properties.HostAliases
}

// GetHostAliasesOk returns a tuple with the HostAliases field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetHostAliasesOk() ([]HostAliases, bool) {
	if o == nil || utils.IsNil(o.Properties.HostAliases) {
		return nil, false
	}
	return o.Properties.HostAliases, true
}

// HasHostAliases returns a boolean if a field has been set.
func (o *CronTaskComponent) HasHostAliases() bool {
	if o != nil && !utils.IsNil(o.Properties.HostAliases) {
		return true
	}

	return false
}

// SetHostAliases gets a reference to the given []HostAliases and assigns it to the hostAliases field.
// HostAliases:  An optional list of hosts and IPs that will be injected into the pod's hosts file
func (o *CronTaskComponent) SetHostAliases(v []HostAliases) *CronTaskComponent {
	o.Properties.HostAliases = v
	return o
}

// GetImage returns the Image field value
func (o *CronTaskComponent) GetImage() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Image
}

// GetImageOk returns a tuple with the Image field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetImageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Image, true
}

// SetImage sets field value
func (o *CronTaskComponent) SetImage(v string) *CronTaskComponent {
	o.Properties.Image = &v
	return o
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *CronTaskComponent) GetImagePullPolicy() string {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		var ret string
		return ret
	}
	return *o.Properties.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullPolicy) {
		return nil, false
	}
	return o.Properties.ImagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *CronTaskComponent) HasImagePullPolicy() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullPolicy) {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the imagePullPolicy field.
// ImagePullPolicy:  Specify image pull policy for your service
func (o *CronTaskComponent) SetImagePullPolicy(v string) *CronTaskComponent {
	o.Properties.ImagePullPolicy = &v
	return o
}

// GetImagePullSecrets returns the ImagePullSecrets field value if set, zero value otherwise.
func (o *CronTaskComponent) GetImagePullSecrets() []string {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		var ret []string
		return ret
	}
	return o.Properties.ImagePullSecrets
}

// GetImagePullSecretsOk returns a tuple with the ImagePullSecrets field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetImagePullSecretsOk() ([]string, bool) {
	if o == nil || utils.IsNil(o.Properties.ImagePullSecrets) {
		return nil, false
	}
	return o.Properties.ImagePullSecrets, true
}

// HasImagePullSecrets returns a boolean if a field has been set.
func (o *CronTaskComponent) HasImagePullSecrets() bool {
	if o != nil && !utils.IsNil(o.Properties.ImagePullSecrets) {
		return true
	}

	return false
}

// SetImagePullSecrets gets a reference to the given []string and assigns it to the imagePullSecrets field.
// ImagePullSecrets:  Specify image pull secrets for your service
func (o *CronTaskComponent) SetImagePullSecrets(v []string) *CronTaskComponent {
	o.Properties.ImagePullSecrets = v
	return o
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *CronTaskComponent) GetLabels() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		var ret map[string]string
		return ret
	}
	return o.Properties.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetLabelsOk() (map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		return nil, false
	}
	return o.Properties.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *CronTaskComponent) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Properties.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the labels field.
// Labels:  Specify the labels in the workload
func (o *CronTaskComponent) SetLabels(v map[string]string) *CronTaskComponent {
	o.Properties.Labels = v
	return o
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *CronTaskComponent) GetLivenessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetLivenessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.LivenessProbe) {
		return nil, false
	}
	return o.Properties.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *CronTaskComponent) HasLivenessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given HealthProbe and assigns it to the livenessProbe field.
// LivenessProbe:
func (o *CronTaskComponent) SetLivenessProbe(v HealthProbe) *CronTaskComponent {
	o.Properties.LivenessProbe = &v
	return o
}

// GetMemory returns the Memory field value if set, zero value otherwise.
func (o *CronTaskComponent) GetMemory() string {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		var ret string
		return ret
	}
	return *o.Properties.Memory
}

// GetMemoryOk returns a tuple with the Memory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetMemoryOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Properties.Memory) {
		return nil, false
	}
	return o.Properties.Memory, true
}

// HasMemory returns a boolean if a field has been set.
func (o *CronTaskComponent) HasMemory() bool {
	if o != nil && !utils.IsNil(o.Properties.Memory) {
		return true
	}

	return false
}

// SetMemory gets a reference to the given string and assigns it to the memory field.
// Memory:  Specifies the attributes of the memory resource required for the container.
func (o *CronTaskComponent) SetMemory(v string) *CronTaskComponent {
	o.Properties.Memory = &v
	return o
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *CronTaskComponent) GetReadinessProbe() HealthProbe {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		var ret HealthProbe
		return ret
	}
	return *o.Properties.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetReadinessProbeOk() (*HealthProbe, bool) {
	if o == nil || utils.IsNil(o.Properties.ReadinessProbe) {
		return nil, false
	}
	return o.Properties.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *CronTaskComponent) HasReadinessProbe() bool {
	if o != nil && !utils.IsNil(o.Properties.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given HealthProbe and assigns it to the readinessProbe field.
// ReadinessProbe:
func (o *CronTaskComponent) SetReadinessProbe(v HealthProbe) *CronTaskComponent {
	o.Properties.ReadinessProbe = &v
	return o
}

// GetRestart returns the Restart field value
func (o *CronTaskComponent) GetRestart() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Restart
}

// GetRestartOk returns a tuple with the Restart field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetRestartOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Restart, true
}

// SetRestart sets field value
func (o *CronTaskComponent) SetRestart(v string) *CronTaskComponent {
	o.Properties.Restart = &v
	return o
}

// GetSchedule returns the Schedule field value
func (o *CronTaskComponent) GetSchedule() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Schedule
}

// GetScheduleOk returns a tuple with the Schedule field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetScheduleOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Schedule, true
}

// SetSchedule sets field value
func (o *CronTaskComponent) SetSchedule(v string) *CronTaskComponent {
	o.Properties.Schedule = &v
	return o
}

// GetStartingDeadlineSeconds returns the StartingDeadlineSeconds field value if set, zero value otherwise.
func (o *CronTaskComponent) GetStartingDeadlineSeconds() int32 {
	if o == nil || utils.IsNil(o.Properties.StartingDeadlineSeconds) {
		var ret int32
		return ret
	}
	return *o.Properties.StartingDeadlineSeconds
}

// GetStartingDeadlineSecondsOk returns a tuple with the StartingDeadlineSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetStartingDeadlineSecondsOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.StartingDeadlineSeconds) {
		return nil, false
	}
	return o.Properties.StartingDeadlineSeconds, true
}

// HasStartingDeadlineSeconds returns a boolean if a field has been set.
func (o *CronTaskComponent) HasStartingDeadlineSeconds() bool {
	if o != nil && !utils.IsNil(o.Properties.StartingDeadlineSeconds) {
		return true
	}

	return false
}

// SetStartingDeadlineSeconds gets a reference to the given int32 and assigns it to the startingDeadlineSeconds field.
// StartingDeadlineSeconds:  Specify deadline in seconds for starting the job if it misses scheduled
func (o *CronTaskComponent) SetStartingDeadlineSeconds(v int32) *CronTaskComponent {
	o.Properties.StartingDeadlineSeconds = &v
	return o
}

// GetSuccessfulJobsHistoryLimit returns the SuccessfulJobsHistoryLimit field value
func (o *CronTaskComponent) GetSuccessfulJobsHistoryLimit() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Properties.SuccessfulJobsHistoryLimit
}

// GetSuccessfulJobsHistoryLimitOk returns a tuple with the SuccessfulJobsHistoryLimit field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetSuccessfulJobsHistoryLimitOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.SuccessfulJobsHistoryLimit, true
}

// SetSuccessfulJobsHistoryLimit sets field value
func (o *CronTaskComponent) SetSuccessfulJobsHistoryLimit(v int32) *CronTaskComponent {
	o.Properties.SuccessfulJobsHistoryLimit = &v
	return o
}

// GetSuspend returns the Suspend field value
func (o *CronTaskComponent) GetSuspend() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return *o.Properties.Suspend
}

// GetSuspendOk returns a tuple with the Suspend field value
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetSuspendOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Suspend, true
}

// SetSuspend sets field value
func (o *CronTaskComponent) SetSuspend(v bool) *CronTaskComponent {
	o.Properties.Suspend = &v
	return o
}

// GetTtlSecondsAfterFinished returns the TtlSecondsAfterFinished field value if set, zero value otherwise.
func (o *CronTaskComponent) GetTtlSecondsAfterFinished() int32 {
	if o == nil || utils.IsNil(o.Properties.TtlSecondsAfterFinished) {
		var ret int32
		return ret
	}
	return *o.Properties.TtlSecondsAfterFinished
}

// GetTtlSecondsAfterFinishedOk returns a tuple with the TtlSecondsAfterFinished field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetTtlSecondsAfterFinishedOk() (*int32, bool) {
	if o == nil || utils.IsNil(o.Properties.TtlSecondsAfterFinished) {
		return nil, false
	}
	return o.Properties.TtlSecondsAfterFinished, true
}

// HasTtlSecondsAfterFinished returns a boolean if a field has been set.
func (o *CronTaskComponent) HasTtlSecondsAfterFinished() bool {
	if o != nil && !utils.IsNil(o.Properties.TtlSecondsAfterFinished) {
		return true
	}

	return false
}

// SetTtlSecondsAfterFinished gets a reference to the given int32 and assigns it to the ttlSecondsAfterFinished field.
// TtlSecondsAfterFinished:  Limits the lifetime of a Job that has finished
func (o *CronTaskComponent) SetTtlSecondsAfterFinished(v int32) *CronTaskComponent {
	o.Properties.TtlSecondsAfterFinished = &v
	return o
}

// GetVolumes returns the Volumes field value if set, zero value otherwise.
func (o *CronTaskComponent) GetVolumes() []Volumes {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		var ret []Volumes
		return ret
	}
	return o.Properties.Volumes
}

// GetVolumesOk returns a tuple with the Volumes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CronTaskComponent) GetVolumesOk() ([]Volumes, bool) {
	if o == nil || utils.IsNil(o.Properties.Volumes) {
		return nil, false
	}
	return o.Properties.Volumes, true
}

// HasVolumes returns a boolean if a field has been set.
func (o *CronTaskComponent) HasVolumes() bool {
	if o != nil && !utils.IsNil(o.Properties.Volumes) {
		return true
	}

	return false
}

// SetVolumes gets a reference to the given []Volumes and assigns it to the volumes field.
// Volumes:  Declare volumes and volumeMounts
func (o *CronTaskComponent) SetVolumes(v []Volumes) *CronTaskComponent {
	o.Properties.Volumes = v
	return o
}

func (o CronTaskSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CronTaskSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ActiveDeadlineSeconds) {
		toSerialize["activeDeadlineSeconds"] = o.ActiveDeadlineSeconds
	}
	if !utils.IsNil(o.Annotations) {
		toSerialize["annotations"] = o.Annotations
	}
	toSerialize["backoffLimit"] = o.BackoffLimit
	if !utils.IsNil(o.Cmd) {
		toSerialize["cmd"] = o.Cmd
	}
	toSerialize["concurrencyPolicy"] = o.ConcurrencyPolicy
	toSerialize["count"] = o.Count
	if !utils.IsNil(o.Cpu) {
		toSerialize["cpu"] = o.Cpu
	}
	if !utils.IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	toSerialize["failedJobsHistoryLimit"] = o.FailedJobsHistoryLimit
	if !utils.IsNil(o.HostAliases) {
		toSerialize["hostAliases"] = o.HostAliases
	}
	toSerialize["image"] = o.Image
	if !utils.IsNil(o.ImagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	}
	if !utils.IsNil(o.ImagePullSecrets) {
		toSerialize["imagePullSecrets"] = o.ImagePullSecrets
	}
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
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
	toSerialize["restart"] = o.Restart
	toSerialize["schedule"] = o.Schedule
	if !utils.IsNil(o.StartingDeadlineSeconds) {
		toSerialize["startingDeadlineSeconds"] = o.StartingDeadlineSeconds
	}
	toSerialize["successfulJobsHistoryLimit"] = o.SuccessfulJobsHistoryLimit
	toSerialize["suspend"] = o.Suspend
	if !utils.IsNil(o.TtlSecondsAfterFinished) {
		toSerialize["ttlSecondsAfterFinished"] = o.TtlSecondsAfterFinished
	}
	if !utils.IsNil(o.Volumes) {
		toSerialize["volumes"] = o.Volumes
	}
	return toSerialize, nil
}

type NullableCronTaskSpec struct {
	value *CronTaskSpec
	isSet bool
}

func (v *NullableCronTaskSpec) Get() *CronTaskSpec {
	return v.value
}

func (v *NullableCronTaskSpec) Set(val *CronTaskSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableCronTaskSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableCronTaskSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCronTaskSpec(val *CronTaskSpec) *NullableCronTaskSpec {
	return &NullableCronTaskSpec{value: val, isSet: true}
}

func (v NullableCronTaskSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCronTaskSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const CronTaskType = "cron-task"

func init() {
	sdkcommon.RegisterComponent(CronTaskType, FromComponent)
}

type CronTaskComponent struct {
	Base       apis.ComponentBase
	Properties CronTaskSpec
}

func CronTask(name string) *CronTaskComponent {
	c := &CronTaskComponent{Base: apis.ComponentBase{
		Name: name,
		Type: CronTaskType,
	}}
	return c
}

func (c *CronTaskComponent) Build() common.ApplicationComponent {
	traits := make([]common.ApplicationTrait, 0)
	for _, trait := range c.Base.Traits {
		traits = append(traits, trait.Build())
	}
	res := common.ApplicationComponent{
		DependsOn:  c.Base.DependsOn,
		Inputs:     c.Base.Inputs,
		Name:       c.Base.Name,
		Outputs:    c.Base.Outputs,
		Properties: util.Object2RawExtension(c.Properties),
		Traits:     traits,
		Type:       CronTaskType,
	}
	return res
}

func (c *CronTaskComponent) FromComponent(from common.ApplicationComponent) (*CronTaskComponent, error) {
	for _, trait := range from.Traits {
		_t, err := sdkcommon.FromTrait(&trait)
		if err != nil {
			return nil, err
		}
		c.Base.Traits = append(c.Base.Traits, _t)
	}
	var properties CronTaskSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	c.Base.Name = from.Name
	c.Base.DependsOn = from.DependsOn
	c.Base.Inputs = from.Inputs
	c.Base.Outputs = from.Outputs
	c.Base.Type = CronTaskType
	c.Properties = properties
	return c, nil
}

func FromComponent(from common.ApplicationComponent) (apis.Component, error) {
	c := &CronTaskComponent{}
	return c.FromComponent(from)
}

func (c *CronTaskComponent) SetTraits(traits ...apis.Trait) *CronTaskComponent {
	for _, addTrait := range traits {
		found := false
		for i, _t := range c.Base.Traits {
			if _t.DefType() == addTrait.DefType() {
				c.Base.Traits[i] = addTrait
				found = true
				break
			}
			if !found {
				c.Base.Traits = append(c.Base.Traits, addTrait)
			}
		}
	}
	return c
}

func (c *CronTaskComponent) GetTrait(typ string) apis.Trait {
	for _, _t := range c.Base.Traits {
		if _t.DefType() == typ {
			return _t
		}
	}
	return nil
}

func (c *CronTaskComponent) GetAllTraits() []apis.Trait {
	return c.Base.Traits
}

func (c *CronTaskComponent) ComponentName() string {
	return c.Base.Name
}

func (c *CronTaskComponent) DefType() string {
	return CronTaskType
}

func (c *CronTaskComponent) DependsOn(dependsOn []string) *CronTaskComponent {
	c.Base.DependsOn = dependsOn
	return c
}

func (c *CronTaskComponent) Inputs(input common.StepInputs) *CronTaskComponent {
	c.Base.Inputs = input
	return c
}

func (c *CronTaskComponent) Outputs(output common.StepOutputs) *CronTaskComponent {
	c.Base.Outputs = output
	return c
}

func (c *CronTaskComponent) AddDependsOn(dependsOn string) *CronTaskComponent {
	c.Base.DependsOn = append(c.Base.DependsOn, dependsOn)
	return c
}