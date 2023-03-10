/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package helm_labels

import (
	"encoding/json"
	"errors"

	"github.com/oam-dev/kubevela-core-api/apis/core.oam.dev/common"
	"github.com/oam-dev/kubevela-core-api/pkg/oam/util"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis"
	sdkcommon "github.com/chivalryq/kubevela-go-sdk/pkg/apis/common"
	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the HelmLabelsSpec type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HelmLabelsSpec{}

// HelmLabelsSpec struct for HelmLabelsSpec
type HelmLabelsSpec struct {
	ApiVersion *string           `json:"apiVersion"`
	Kind       *string           `json:"kind"`
	Labels     map[string]string `json:"labels,omitempty"`
	Name       *string           `json:"name"`
	Namespace  *string           `json:"namespace"`
}

// NewHelmLabelsSpecWith instantiates a new HelmLabelsSpec object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewHelmLabelsSpecWith(apiVersion string, kind string, name string, namespace string) *HelmLabelsSpec {
	this := HelmLabelsSpec{}
	this.ApiVersion = &apiVersion
	this.Kind = &kind
	this.Name = &name
	this.Namespace = &namespace
	return &this
}

// NewHelmLabelsSpecWithDefault instantiates a new HelmLabelsSpec object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmLabelsSpecWithDefault() *HelmLabelsSpec {
	this := HelmLabelsSpec{}
	var apiVersion string = "apps/v1"
	this.ApiVersion = &apiVersion
	var kind string = "Deployment"
	this.Kind = &kind
	return &this
}

// NewHelmLabelsSpec is short for NewHelmLabelsSpecWithDefault which instantiates a new HelmLabelsSpec object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHelmLabelsSpec() *HelmLabelsSpec {
	return NewHelmLabelsSpecWithDefault()
}

// NewHelmLabelsSpecEmpty instantiates a new HelmLabelsSpec object with no properties set.
// This constructor will not assign any default values to properties.
func NewHelmLabelsSpecEmpty() *HelmLabelsSpec {
	this := HelmLabelsSpec{}
	return &this
}

// NewHelmLabelsSpecs converts a list HelmLabelsSpec pointers to objects.
// This is helpful when the SetHelmLabelsSpec requires a list of objects
func NewHelmLabelsSpecList(ps ...*HelmLabelsSpec) []HelmLabelsSpec {
	objs := []HelmLabelsSpec{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this HelmLabelsSpec
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *HelmLabelsTrait) Validate() error {
	if o.Properties.ApiVersion == nil {
		return errors.New("ApiVersion in HelmLabelsSpec must be set")
	}
	if o.Properties.Kind == nil {
		return errors.New("Kind in HelmLabelsSpec must be set")
	}
	if o.Properties.Name == nil {
		return errors.New("Name in HelmLabelsSpec must be set")
	}
	if o.Properties.Namespace == nil {
		return errors.New("Namespace in HelmLabelsSpec must be set")
	}
	// validate all nested properties
	return nil
}

// GetApiVersion returns the ApiVersion field value
func (o *HelmLabelsTrait) GetApiVersion() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.ApiVersion
}

// GetApiVersionOk returns a tuple with the ApiVersion field value
// and a boolean to check if the value has been set.
func (o *HelmLabelsTrait) GetApiVersionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.ApiVersion, true
}

// SetApiVersion sets field value
func (o *HelmLabelsTrait) SetApiVersion(v string) *HelmLabelsTrait {
	o.Properties.ApiVersion = &v
	return o
}

// GetKind returns the Kind field value
func (o *HelmLabelsTrait) GetKind() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *HelmLabelsTrait) GetKindOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Kind, true
}

// SetKind sets field value
func (o *HelmLabelsTrait) SetKind(v string) *HelmLabelsTrait {
	o.Properties.Kind = &v
	return o
}

// GetLabels returns the Labels field value if set, zero value otherwise.
func (o *HelmLabelsTrait) GetLabels() map[string]string {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		var ret map[string]string
		return ret
	}
	return o.Properties.Labels
}

// GetLabelsOk returns a tuple with the Labels field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HelmLabelsTrait) GetLabelsOk() (map[string]string, bool) {
	if o == nil || utils.IsNil(o.Properties.Labels) {
		return nil, false
	}
	return o.Properties.Labels, true
}

// HasLabels returns a boolean if a field has been set.
func (o *HelmLabelsTrait) HasLabels() bool {
	if o != nil && !utils.IsNil(o.Properties.Labels) {
		return true
	}

	return false
}

// SetLabels gets a reference to the given map[string]string and assigns it to the labels field.
// Labels:
func (o *HelmLabelsTrait) SetLabels(v map[string]string) *HelmLabelsTrait {
	o.Properties.Labels = v
	return o
}

// GetName returns the Name field value
func (o *HelmLabelsTrait) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *HelmLabelsTrait) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Name, true
}

// SetName sets field value
func (o *HelmLabelsTrait) SetName(v string) *HelmLabelsTrait {
	o.Properties.Name = &v
	return o
}

// GetNamespace returns the Namespace field value
func (o *HelmLabelsTrait) GetNamespace() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Properties.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *HelmLabelsTrait) GetNamespaceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties.Namespace, true
}

// SetNamespace sets field value
func (o *HelmLabelsTrait) SetNamespace(v string) *HelmLabelsTrait {
	o.Properties.Namespace = &v
	return o
}

func (o HelmLabelsSpec) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HelmLabelsSpec) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["apiVersion"] = o.ApiVersion
	toSerialize["kind"] = o.Kind
	if !utils.IsNil(o.Labels) {
		toSerialize["labels"] = o.Labels
	}
	toSerialize["name"] = o.Name
	toSerialize["namespace"] = o.Namespace
	return toSerialize, nil
}

type NullableHelmLabelsSpec struct {
	value *HelmLabelsSpec
	isSet bool
}

func (v *NullableHelmLabelsSpec) Get() *HelmLabelsSpec {
	return v.value
}

func (v *NullableHelmLabelsSpec) Set(val *HelmLabelsSpec) {
	v.value = val
	v.isSet = true
}

func (v *NullableHelmLabelsSpec) IsSet() bool {
	return v.isSet
}

func (v *NullableHelmLabelsSpec) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHelmLabelsSpec(val *HelmLabelsSpec) *NullableHelmLabelsSpec {
	return &NullableHelmLabelsSpec{value: val, isSet: true}
}

func (v NullableHelmLabelsSpec) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHelmLabelsSpec) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

const HelmLabelsType = "helm-labels"

func init() {
	sdkcommon.RegisterTrait(HelmLabelsType, FromTrait)
}

type HelmLabelsTrait struct {
	Base       apis.TraitBase
	Properties HelmLabelsSpec
}

func HelmLabels() *HelmLabelsTrait {
	h := &HelmLabelsTrait{Base: apis.TraitBase{}}
	return h
}

func (h *HelmLabelsTrait) Build() common.ApplicationTrait {
	res := common.ApplicationTrait{
		Properties: util.Object2RawExtension(h.Properties),
		Type:       HelmLabelsType,
	}
	return res
}

func (h *HelmLabelsTrait) FromTrait(from common.ApplicationTrait) (*HelmLabelsTrait, error) {
	var properties HelmLabelsSpec
	if from.Properties != nil {
		err := json.Unmarshal(from.Properties.Raw, &properties)
		if err != nil {
			return nil, err
		}
	}
	h.Base.Type = HelmLabelsType
	h.Properties = properties
	return h, nil
}

func FromTrait(from common.ApplicationTrait) (apis.Trait, error) {
	h := &HelmLabelsTrait{}
	return h.FromTrait(from)
}

func (h *HelmLabelsTrait) DefType() string {
	return HelmLabelsType
}