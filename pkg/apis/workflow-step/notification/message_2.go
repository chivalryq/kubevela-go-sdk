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

	"github.com/kubevela-contrib/kubevela-go-sdk/pkg/apis/utils"
)

// checks if the Message2 type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Message2{}

// Message2 Specify the message that you want to sent, refer to [slack messaging](https://api.slack.com/reference/messaging/payload)
type Message2 struct {
	Attachments *Attachments `json:"attachments,omitempty"`
	Blocks      []Block      `json:"blocks,omitempty"`
	// Specify the message text format in markdown for slack notification
	Mrkdwn *bool `json:"mrkdwn,omitempty"`
	// Specify the message text for slack notification
	Text     *string `json:"text"`
	ThreadTs *string `json:"thread_ts,omitempty"`
}

// NewMessage2With instantiates a new Message2 object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewMessage2With(text string) *Message2 {
	this := Message2{}
	var mrkdwn bool = true
	this.Mrkdwn = &mrkdwn
	this.Text = &text
	return &this
}

// NewMessage2WithDefault instantiates a new Message2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessage2WithDefault() *Message2 {
	this := Message2{}
	var mrkdwn bool = true
	this.Mrkdwn = &mrkdwn
	return &this
}

// NewMessage2 is short for NewMessage2WithDefault which instantiates a new Message2 object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessage2() *Message2 {
	return NewMessage2WithDefault()
}

// NewMessage2Empty instantiates a new Message2 object with no properties set.
// This constructor will not assign any default values to properties.
func NewMessage2Empty() *Message2 {
	this := Message2{}
	return &this
}

// NewMessage2s converts a list Message2 pointers to objects.
// This is helpful when the SetMessage2 requires a list of objects
func NewMessage2List(ps ...*Message2) []Message2 {
	objs := []Message2{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Message2
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Message2) Validate() error {
	if o.Text == nil {
		return errors.New("Text in Message2 must be set")
	}
	// validate all nested properties
	if o.Attachments != nil {
		if err := o.Attachments.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetAttachments returns the Attachments field value if set, zero value otherwise.
func (o *Message2) GetAttachments() Attachments {
	if o == nil || utils.IsNil(o.Attachments) {
		var ret Attachments
		return ret
	}
	return *o.Attachments
}

// GetAttachmentsOk returns a tuple with the Attachments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message2) GetAttachmentsOk() (*Attachments, bool) {
	if o == nil || utils.IsNil(o.Attachments) {
		return nil, false
	}
	return o.Attachments, true
}

// HasAttachments returns a boolean if a field has been set.
func (o *Message2) HasAttachments() bool {
	if o != nil && !utils.IsNil(o.Attachments) {
		return true
	}

	return false
}

// SetAttachments gets a reference to the given Attachments and assigns it to the attachments field.
// Attachments:
func (o *Message2) SetAttachments(v Attachments) *Message2 {
	o.Attachments = &v
	return o
}

// GetBlocks returns the Blocks field value if set, zero value otherwise.
func (o *Message2) GetBlocks() []Block {
	if o == nil || utils.IsNil(o.Blocks) {
		var ret []Block
		return ret
	}
	return o.Blocks
}

// GetBlocksOk returns a tuple with the Blocks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message2) GetBlocksOk() ([]Block, bool) {
	if o == nil || utils.IsNil(o.Blocks) {
		return nil, false
	}
	return o.Blocks, true
}

// HasBlocks returns a boolean if a field has been set.
func (o *Message2) HasBlocks() bool {
	if o != nil && !utils.IsNil(o.Blocks) {
		return true
	}

	return false
}

// SetBlocks gets a reference to the given []Block and assigns it to the blocks field.
// Blocks:
func (o *Message2) SetBlocks(v []Block) *Message2 {
	o.Blocks = v
	return o
}

// GetMrkdwn returns the Mrkdwn field value if set, zero value otherwise.
func (o *Message2) GetMrkdwn() bool {
	if o == nil || utils.IsNil(o.Mrkdwn) {
		var ret bool
		return ret
	}
	return *o.Mrkdwn
}

// GetMrkdwnOk returns a tuple with the Mrkdwn field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message2) GetMrkdwnOk() (*bool, bool) {
	if o == nil || utils.IsNil(o.Mrkdwn) {
		return nil, false
	}
	return o.Mrkdwn, true
}

// HasMrkdwn returns a boolean if a field has been set.
func (o *Message2) HasMrkdwn() bool {
	if o != nil && !utils.IsNil(o.Mrkdwn) {
		return true
	}

	return false
}

// SetMrkdwn gets a reference to the given bool and assigns it to the mrkdwn field.
// Mrkdwn:  Specify the message text format in markdown for slack notification
func (o *Message2) SetMrkdwn(v bool) *Message2 {
	o.Mrkdwn = &v
	return o
}

// GetText returns the Text field value
func (o *Message2) GetText() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Text
}

// GetTextOk returns a tuple with the Text field value
// and a boolean to check if the value has been set.
func (o *Message2) GetTextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Text, true
}

// SetText sets field value
func (o *Message2) SetText(v string) *Message2 {
	o.Text = &v
	return o
}

// GetThreadTs returns the ThreadTs field value if set, zero value otherwise.
func (o *Message2) GetThreadTs() string {
	if o == nil || utils.IsNil(o.ThreadTs) {
		var ret string
		return ret
	}
	return *o.ThreadTs
}

// GetThreadTsOk returns a tuple with the ThreadTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message2) GetThreadTsOk() (*string, bool) {
	if o == nil || utils.IsNil(o.ThreadTs) {
		return nil, false
	}
	return o.ThreadTs, true
}

// HasThreadTs returns a boolean if a field has been set.
func (o *Message2) HasThreadTs() bool {
	if o != nil && !utils.IsNil(o.ThreadTs) {
		return true
	}

	return false
}

// SetThreadTs gets a reference to the given string and assigns it to the threadTs field.
// ThreadTs:
func (o *Message2) SetThreadTs(v string) *Message2 {
	o.ThreadTs = &v
	return o
}

func (o Message2) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message2) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.Attachments) {
		toSerialize["attachments"] = o.Attachments
	}
	if !utils.IsNil(o.Blocks) {
		toSerialize["blocks"] = o.Blocks
	}
	if !utils.IsNil(o.Mrkdwn) {
		toSerialize["mrkdwn"] = o.Mrkdwn
	}
	toSerialize["text"] = o.Text
	if !utils.IsNil(o.ThreadTs) {
		toSerialize["thread_ts"] = o.ThreadTs
	}
	return toSerialize, nil
}

type NullableMessage2 struct {
	value *Message2
	isSet bool
}

func (v *NullableMessage2) Get() *Message2 {
	return v.value
}

func (v *NullableMessage2) Set(val *Message2) {
	v.value = val
	v.isSet = true
}

func (v *NullableMessage2) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage2(val *Message2) *NullableMessage2 {
	return &NullableMessage2{value: val, isSet: true}
}

func (v NullableMessage2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
