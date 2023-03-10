/*
Generated by cue.

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: no version
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package notification

import (
	"encoding/json"
	"errors"

	"github.com/chivalryq/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Content type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Content{}

// Content Specify the content of the email
type Content struct {
	// Specify the context body of the email
	Body *string `json:"body"`
	// Specify the subject of the email
	Subject *string `json:"subject"`
}

// NewContentWith instantiates a new Content object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewContentWith(body string, subject string) *Content {
	this := Content{}
	this.Body = &body
	this.Subject = &subject
	return &this
}

// NewContentWithDefault instantiates a new Content object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContentWithDefault() *Content {
	this := Content{}
	return &this
}

// NewContent is short for NewContentWithDefault which instantiates a new Content object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContent() *Content {
	return NewContentWithDefault()
}

// NewContentEmpty instantiates a new Content object with no properties set.
// This constructor will not assign any default values to properties.
func NewContentEmpty() *Content {
	this := Content{}
	return &this
}

// NewContents converts a list Content pointers to objects.
// This is helpful when the SetContent requires a list of objects
func NewContentList(ps ...*Content) []Content {
	objs := []Content{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Content
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Content) Validate() error {
	if o.Body == nil {
		return errors.New("Body in Content must be set")
	}
	if o.Subject == nil {
		return errors.New("Subject in Content must be set")
	}
	// validate all nested properties
	return nil
}

// GetBody returns the Body field value
func (o *Content) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *Content) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Body, true
}

// SetBody sets field value
func (o *Content) SetBody(v string) *Content {
	o.Body = &v
	return o
}

// GetSubject returns the Subject field value
func (o *Content) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *Content) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Subject, true
}

// SetSubject sets field value
func (o *Content) SetSubject(v string) *Content {
	o.Subject = &v
	return o
}

func (o Content) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Content) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["body"] = o.Body
	toSerialize["subject"] = o.Subject
	return toSerialize, nil
}

type NullableContent struct {
	value *Content
	isSet bool
}

func (v *NullableContent) Get() *Content {
	return v.value
}

func (v *NullableContent) Set(val *Content) {
	v.value = val
	v.isSet = true
}

func (v *NullableContent) IsSet() bool {
	return v.isSet
}

func (v *NullableContent) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContent(val *Content) *NullableContent {
	return &NullableContent{value: val, isSet: true}
}

func (v NullableContent) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContent) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
