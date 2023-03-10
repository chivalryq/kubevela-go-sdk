/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package lifecycle

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the HttpGet type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &HttpGet{}

// HttpGet struct for HttpGet
type HttpGet struct {
	Host        *string       `json:"host,omitempty"`
	HttpHeaders []HttpHeaders `json:"httpHeaders,omitempty"`
	Path        *string       `json:"path,omitempty"`
	Port        *int32        `json:"port"`
	Scheme      *string       `json:"scheme"`
}

// NewHttpGetWith instantiates a new HttpGet object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewHttpGetWith(port int32, scheme string) *HttpGet {
	this := HttpGet{}
	this.Port = &port
	this.Scheme = &scheme
	return &this
}

// NewHttpGetWithDefault instantiates a new HttpGet object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpGetWithDefault() *HttpGet {
	this := HttpGet{}
	var scheme string = "HTTP"
	this.Scheme = &scheme
	return &this
}

// NewHttpGet is short for NewHttpGetWithDefault which instantiates a new HttpGet object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpGet() *HttpGet {
	return NewHttpGetWithDefault()
}

// NewHttpGetEmpty instantiates a new HttpGet object with no properties set.
// This constructor will not assign any default values to properties.
func NewHttpGetEmpty() *HttpGet {
	this := HttpGet{}
	return &this
}

// NewHttpGets converts a list HttpGet pointers to objects.
// This is helpful when the SetHttpGet requires a list of objects
func NewHttpGetList(ps ...*HttpGet) []HttpGet {
	objs := []HttpGet{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this HttpGet
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *HttpGet) Validate() error {
	if o.Port == nil {
		return errors.New("Port in HttpGet must be set")
	}
	if o.Scheme == nil {
		return errors.New("Scheme in HttpGet must be set")
	}
	// validate all nested properties
	return nil
}

// GetHost returns the Host field value if set, zero value otherwise.
func (o *HttpGet) GetHost() string {
	if o == nil || utils.IsNil(o.Host) {
		var ret string
		return ret
	}
	return *o.Host
}

// GetHostOk returns a tuple with the Host field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetHostOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Host) {
		return nil, false
	}
	return o.Host, true
}

// HasHost returns a boolean if a field has been set.
func (o *HttpGet) HasHost() bool {
	if o != nil && !utils.IsNil(o.Host) {
		return true
	}

	return false
}

// SetHost gets a reference to the given string and assigns it to the host field.
// Host:
func (o *HttpGet) SetHost(v string) *HttpGet {
	o.Host = &v
	return o
}

// GetHttpHeaders returns the HttpHeaders field value if set, zero value otherwise.
func (o *HttpGet) GetHttpHeaders() []HttpHeaders {
	if o == nil || utils.IsNil(o.HttpHeaders) {
		var ret []HttpHeaders
		return ret
	}
	return o.HttpHeaders
}

// GetHttpHeadersOk returns a tuple with the HttpHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetHttpHeadersOk() ([]HttpHeaders, bool) {
	if o == nil || utils.IsNil(o.HttpHeaders) {
		return nil, false
	}
	return o.HttpHeaders, true
}

// HasHttpHeaders returns a boolean if a field has been set.
func (o *HttpGet) HasHttpHeaders() bool {
	if o != nil && !utils.IsNil(o.HttpHeaders) {
		return true
	}

	return false
}

// SetHttpHeaders gets a reference to the given []HttpHeaders and assigns it to the httpHeaders field.
// HttpHeaders:
func (o *HttpGet) SetHttpHeaders(v []HttpHeaders) *HttpGet {
	o.HttpHeaders = v
	return o
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *HttpGet) GetPath() string {
	if o == nil || utils.IsNil(o.Path) {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpGet) GetPathOk() (*string, bool) {
	if o == nil || utils.IsNil(o.Path) {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *HttpGet) HasPath() bool {
	if o != nil && !utils.IsNil(o.Path) {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the path field.
// Path:
func (o *HttpGet) SetPath(v string) *HttpGet {
	o.Path = &v
	return o
}

// GetPort returns the Port field value
func (o *HttpGet) GetPort() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return *o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *HttpGet) GetPortOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Port, true
}

// SetPort sets field value
func (o *HttpGet) SetPort(v int32) *HttpGet {
	o.Port = &v
	return o
}

// GetScheme returns the Scheme field value
func (o *HttpGet) GetScheme() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Scheme
}

// GetSchemeOk returns a tuple with the Scheme field value
// and a boolean to check if the value has been set.
func (o *HttpGet) GetSchemeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Scheme, true
}

// SetScheme sets field value
func (o *HttpGet) SetScheme(v string) *HttpGet {
	o.Scheme = &v
	return o
}

func (o HttpGet) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpGet) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Host) {
		toSerialize["host"] = o.Host
	}
	if !utils.IsNil(o.HttpHeaders) {
		toSerialize["httpHeaders"] = o.HttpHeaders
	}
	if !utils.IsNil(o.Path) {
		toSerialize["path"] = o.Path
	}
	toSerialize["port"] = o.Port
	toSerialize["scheme"] = o.Scheme
	return toSerialize, nil
}

type NullableHttpGet struct {
	value *HttpGet
	isSet bool
}

func (v *NullableHttpGet) Get() *HttpGet {
	return v.value
}

func (v *NullableHttpGet) Set(val *HttpGet) {
	v.value = val
	v.isSet = true
}

func (v *NullableHttpGet) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpGet) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpGet(val *HttpGet) *NullableHttpGet {
	return &NullableHttpGet{value: val, isSet: true}
}

func (v NullableHttpGet) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpGet) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
