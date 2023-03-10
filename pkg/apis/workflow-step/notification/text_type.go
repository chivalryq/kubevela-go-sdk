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

// checks if the TextType type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &TextType{}

// TextType struct for TextType
type TextType struct {
	Emoji    *bool   `json:"emoji,omitempty"`
	Text     *string `json:"text"`
	Type     *string `json:"type"`
	Verbatim *bool   `json:"verbatim,omitempty"`
}

// NewTextTypeWith instantiates a new TextType object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewTextTypeWith(text string, type_ string) *TextType {
	this := TextType{}
	this.Text = &text
	this.Type = &type_
	return &this
}

// NewTextTypeWithDefault instantiates a new TextType object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextTypeWithDefault() *TextType {
	this := TextType{}
	return &this
}

// NewTextType is short for NewTextTypeWithDefault which instantiates a new TextType object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTextType() *TextType {
	return NewTextTypeWithDefault()
}

// NewTextTypeEmpty instantiates a new TextType object with no properties set.
// This constructor will not assign any default values to properties.
func NewTextTypeEmpty() *TextType {
	this := TextType{}
	return &this
}

// NewTextTypes converts a list TextType pointers to objects.
// This is helpful when the SetTextType requires a list of objects
func NewTextTypeList(ps ...*TextType) []TextType {
	objs := []TextType{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this TextType
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *TextType) Validate() error {
	if o.Text == nil {
		return errors.New("Text in TextType must be set")
	}
	if o.Type == nil {
		return errors.New("Type in TextType must be set")
	}
	// validate all nested properties
	return nil
}

// GetEmoji returns the Emoji field value if set, zero value otherwise.
func (o *TextType) GetEmoji() bool {
	if o == nil || utils.IsNil(o.Emoji) {
		var ret bool
		return ret
	}
	return *o.Emoji
}

// GetEmojiOk returns a tuple with the Emoji field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextType) GetEmojiOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Emoji) {
		return nil, false
	}
	return o.Emoji, true
}

// HasEmoji returns a boolean if a field has been set.
func (o *TextType) HasEmoji() bool {
	if o != nil && !utils.IsNil(o.Emoji) {
		return true
	}

	return false
}

// SetEmoji gets a reference to the given bool and assigns it to the emoji field.
// Emoji:
func (o *TextType) SetEmoji(v bool) *TextType {
	o.Emoji = &v
	return o
}

// GetText returns the Text field value
func (o *TextType) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *TextType) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text, true
}

// SetText sets field value
func (o *TextType) SetText(v string) *TextType {
	o.Text = &v
	return o
}

// GetType returns the Type field value
func (o *TextType) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *TextType) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Type, true
}

// SetType sets field value
func (o *TextType) SetType(v string) *TextType {
	o.Type = &v
	return o
}

// GetVerbatim returns the Verbatim field value if set, zero value otherwise.
func (o *TextType) GetVerbatim() bool {
	if o == nil || utils.IsNil(o.Verbatim) {
		var ret bool
		return ret
	}
	return *o.Verbatim
}

// GetVerbatimOk returns a tuple with the Verbatim field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TextType) GetVerbatimOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Verbatim) {
		return nil, false
	}
	return o.Verbatim, true
}

// HasVerbatim returns a boolean if a field has been set.
func (o *TextType) HasVerbatim() bool {
	if o != nil && !utils.IsNil(o.Verbatim) {
		return true
	}

	return false
}

// SetVerbatim gets a reference to the given bool and assigns it to the verbatim field.
// Verbatim:
func (o *TextType) SetVerbatim(v bool) *TextType {
	o.Verbatim = &v
	return o
}

func (o TextType) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TextType) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Emoji) {
		toSerialize["emoji"] = o.Emoji
	}
	toSerialize["text"] = o.Text
	toSerialize["type"] = o.Type
	if !utils.IsNil(o.Verbatim) {
		toSerialize["verbatim"] = o.Verbatim
	}
	return toSerialize, nil
}

type NullableTextType struct {
	value *TextType
	isSet bool
}

func (v *NullableTextType) Get() *TextType {
	return v.value
}

func (v *NullableTextType) Set(val *TextType) {
	v.value = val
	v.isSet = true
}

func (v *NullableTextType) IsSet() bool {
	return v.isSet
}

func (v *NullableTextType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTextType(val *TextType) *NullableTextType {
	return &NullableTextType{value: val, isSet: true}
}

func (v NullableTextType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTextType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
