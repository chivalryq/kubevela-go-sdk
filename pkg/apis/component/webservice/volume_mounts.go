/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webservice

import (
	"encoding/json"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the VolumeMounts type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &VolumeMounts{}

// VolumeMounts struct for VolumeMounts
type VolumeMounts struct {
	// Mount ConfigMap type volume
	ConfigMap []ConfigMap `json:"configMap,omitempty"`
	// Mount EmptyDir type volume
	EmptyDir []EmptyDir `json:"emptyDir,omitempty"`
	// Mount HostPath type volume
	HostPath []HostPath `json:"hostPath,omitempty"`
	// Mount PVC type volume
	Pvc []Pvc `json:"pvc,omitempty"`
	// Mount Secret type volume
	Secret []Secret `json:"secret,omitempty"`
}

// NewVolumeMountsWith instantiates a new VolumeMounts object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewVolumeMountsWith() *VolumeMounts {
	this := VolumeMounts{}
	return &this
}

// NewVolumeMountsWithDefault instantiates a new VolumeMounts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeMountsWithDefault() *VolumeMounts {
	this := VolumeMounts{}
	return &this
}

// NewVolumeMounts is short for NewVolumeMountsWithDefault which instantiates a new VolumeMounts object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVolumeMounts() *VolumeMounts {
	return NewVolumeMountsWithDefault()
}

// NewVolumeMountsEmpty instantiates a new VolumeMounts object with no properties set.
// This constructor will not assign any default values to properties.
func NewVolumeMountsEmpty() *VolumeMounts {
	this := VolumeMounts{}
	return &this
}

// NewVolumeMountss converts a list VolumeMounts pointers to objects.
// This is helpful when the SetVolumeMounts requires a list of objects
func NewVolumeMountsList(ps ...*VolumeMounts) []VolumeMounts {
	objs := []VolumeMounts{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this VolumeMounts
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *VolumeMounts) Validate() error {
	// validate all nested properties
	return nil
}

// GetConfigMap returns the ConfigMap field value if set, zero value otherwise.
func (o *VolumeMounts) GetConfigMap() []ConfigMap {
	if o == nil || utils.IsNil(o.ConfigMap) {
		var ret []ConfigMap
		return ret
	}
	return o.ConfigMap
}

// GetConfigMapOk returns a tuple with the ConfigMap field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetConfigMapOk() ([]ConfigMap, bool) {
	if o == nil || utils.IsNil(o.ConfigMap) {
		return nil, false
	}
	return o.ConfigMap, true
}

// HasConfigMap returns a boolean if a field has been set.
func (o *VolumeMounts) HasConfigMap() bool {
	if o != nil && !utils.IsNil(o.ConfigMap) {
		return true
	}

	return false
}

// SetConfigMap gets a reference to the given []ConfigMap and assigns it to the configMap field.
// ConfigMap:  Mount ConfigMap type volume
func (o *VolumeMounts) SetConfigMap(v []ConfigMap) *VolumeMounts {
	o.ConfigMap = v
	return o
}

// GetEmptyDir returns the EmptyDir field value if set, zero value otherwise.
func (o *VolumeMounts) GetEmptyDir() []EmptyDir {
	if o == nil || utils.IsNil(o.EmptyDir) {
		var ret []EmptyDir
		return ret
	}
	return o.EmptyDir
}

// GetEmptyDirOk returns a tuple with the EmptyDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetEmptyDirOk() ([]EmptyDir, bool) {
	if o == nil || utils.IsNil(o.EmptyDir) {
		return nil, false
	}
	return o.EmptyDir, true
}

// HasEmptyDir returns a boolean if a field has been set.
func (o *VolumeMounts) HasEmptyDir() bool {
	if o != nil && !utils.IsNil(o.EmptyDir) {
		return true
	}

	return false
}

// SetEmptyDir gets a reference to the given []EmptyDir and assigns it to the emptyDir field.
// EmptyDir:  Mount EmptyDir type volume
func (o *VolumeMounts) SetEmptyDir(v []EmptyDir) *VolumeMounts {
	o.EmptyDir = v
	return o
}

// GetHostPath returns the HostPath field value if set, zero value otherwise.
func (o *VolumeMounts) GetHostPath() []HostPath {
	if o == nil || utils.IsNil(o.HostPath) {
		var ret []HostPath
		return ret
	}
	return o.HostPath
}

// GetHostPathOk returns a tuple with the HostPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetHostPathOk() ([]HostPath, bool) {
	if o == nil || utils.IsNil(o.HostPath) {
		return nil, false
	}
	return o.HostPath, true
}

// HasHostPath returns a boolean if a field has been set.
func (o *VolumeMounts) HasHostPath() bool {
	if o != nil && !utils.IsNil(o.HostPath) {
		return true
	}

	return false
}

// SetHostPath gets a reference to the given []HostPath and assigns it to the hostPath field.
// HostPath:  Mount HostPath type volume
func (o *VolumeMounts) SetHostPath(v []HostPath) *VolumeMounts {
	o.HostPath = v
	return o
}

// GetPvc returns the Pvc field value if set, zero value otherwise.
func (o *VolumeMounts) GetPvc() []Pvc {
	if o == nil || utils.IsNil(o.Pvc) {
		var ret []Pvc
		return ret
	}
	return o.Pvc
}

// GetPvcOk returns a tuple with the Pvc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetPvcOk() ([]Pvc, bool) {
	if o == nil || utils.IsNil(o.Pvc) {
		return nil, false
	}
	return o.Pvc, true
}

// HasPvc returns a boolean if a field has been set.
func (o *VolumeMounts) HasPvc() bool {
	if o != nil && !utils.IsNil(o.Pvc) {
		return true
	}

	return false
}

// SetPvc gets a reference to the given []Pvc and assigns it to the pvc field.
// Pvc:  Mount PVC type volume
func (o *VolumeMounts) SetPvc(v []Pvc) *VolumeMounts {
	o.Pvc = v
	return o
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *VolumeMounts) GetSecret() []Secret {
	if o == nil || utils.IsNil(o.Secret) {
		var ret []Secret
		return ret
	}
	return o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VolumeMounts) GetSecretOk() ([]Secret, bool) {
	if o == nil || utils.IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *VolumeMounts) HasSecret() bool {
	if o != nil && !utils.IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given []Secret and assigns it to the secret field.
// Secret:  Mount Secret type volume
func (o *VolumeMounts) SetSecret(v []Secret) *VolumeMounts {
	o.Secret = v
	return o
}

func (o VolumeMounts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VolumeMounts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ConfigMap) {
		toSerialize["configMap"] = o.ConfigMap
	}
	if !utils.IsNil(o.EmptyDir) {
		toSerialize["emptyDir"] = o.EmptyDir
	}
	if !utils.IsNil(o.HostPath) {
		toSerialize["hostPath"] = o.HostPath
	}
	if !utils.IsNil(o.Pvc) {
		toSerialize["pvc"] = o.Pvc
	}
	if !utils.IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	return toSerialize, nil
}

type NullableVolumeMounts struct {
	value *VolumeMounts
	isSet bool
}

func (v *NullableVolumeMounts) Get() *VolumeMounts {
	return v.value
}

func (v *NullableVolumeMounts) Set(val *VolumeMounts) {
	v.value = val
	v.isSet = true
}

func (v *NullableVolumeMounts) IsSet() bool {
	return v.isSet
}

func (v *NullableVolumeMounts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVolumeMounts(val *VolumeMounts) *NullableVolumeMounts {
	return &NullableVolumeMounts{value: val, isSet: true}
}

func (v NullableVolumeMounts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVolumeMounts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}