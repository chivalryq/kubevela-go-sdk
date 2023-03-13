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

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ utils.MappedNullable = &Message{}

// Message Specify the message that you want to sent, refer to [dingtalk messaging](https://developers.dingtalk.com/document/robots/custom-robot-access/title-72m-8ag-pqw)
type Message struct {
	ActionCard *ActionCard `json:"actionCard,omitempty"`
	At         *At         `json:"at,omitempty"`
	FeedCard   *FeedCard   `json:"feedCard,omitempty"`
	Link       *Link       `json:"link,omitempty"`
	Markdown   *Markdown   `json:"markdown,omitempty"`
	// msgType can be text, link, mardown, actionCard, feedCard
	Msgtype *string `json:"msgtype"`
	Text    *Text   `json:"text,omitempty"`
}

// NewMessageWith instantiates a new Message object
// This constructor will make sure properties required by API are set.
// For optional properties, it will set default values if they have been defined.
// The set of arguments will change when the set of required properties is changed
func NewMessageWith(msgtype string) *Message {
	this := Message{}
	this.Msgtype = &msgtype
	return &this
}

// NewMessageWithDefault instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefault() *Message {
	this := Message{}
	var msgtype string = "text"
	this.Msgtype = &msgtype
	return &this
}

// NewMessage is short for NewMessageWithDefault which instantiates a new Message object.
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessage() *Message {
	return NewMessageWithDefault()
}

// NewMessageEmpty instantiates a new Message object with no properties set.
// This constructor will not assign any default values to properties.
func NewMessageEmpty() *Message {
	this := Message{}
	return &this
}

// NewMessages converts a list Message pointers to objects.
// This is helpful when the SetMessage requires a list of objects
func NewMessageList(ps ...*Message) []Message {
	objs := []Message{}
	for _, p := range ps {
		objs = append(objs, *p)
	}
	return objs
}

// Validate validates this Message
// 1. If the required properties are not set, this will return an error
// 2. If properties are set, will check if nested required properties are set
func (o *Message) Validate() error {
	if o.Msgtype == nil {
		return errors.New("Msgtype in Message must be set")
	}
	// validate all nested properties
	if o.ActionCard != nil {
		if err := o.ActionCard.Validate(); err != nil {
			return err
		}
	}
	if o.At != nil {
		if err := o.At.Validate(); err != nil {
			return err
		}
	}
	if o.FeedCard != nil {
		if err := o.FeedCard.Validate(); err != nil {
			return err
		}
	}
	if o.Link != nil {
		if err := o.Link.Validate(); err != nil {
			return err
		}
	}
	if o.Markdown != nil {
		if err := o.Markdown.Validate(); err != nil {
			return err
		}
	}
	if o.Text != nil {
		if err := o.Text.Validate(); err != nil {
			return err
		}
	}
	return nil
}

// GetActionCard returns the ActionCard field value if set, zero value otherwise.
func (o *Message) GetActionCard() ActionCard {
	if o == nil || utils.IsNil(o.ActionCard) {
		var ret ActionCard
		return ret
	}
	return *o.ActionCard
}

// GetActionCardOk returns a tuple with the ActionCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetActionCardOk() (*ActionCard, bool) {
	if o == nil || utils.IsNil(o.ActionCard) {
		return nil, false
	}
	return o.ActionCard, true
}

// HasActionCard returns a boolean if a field has been set.
func (o *Message) HasActionCard() bool {
	if o != nil && !utils.IsNil(o.ActionCard) {
		return true
	}

	return false
}

// SetActionCard gets a reference to the given ActionCard and assigns it to the actionCard field.
// ActionCard:
func (o *Message) SetActionCard(v ActionCard) *Message {
	o.ActionCard = &v
	return o
}

// GetAt returns the At field value if set, zero value otherwise.
func (o *Message) GetAt() At {
	if o == nil || utils.IsNil(o.At) {
		var ret At
		return ret
	}
	return *o.At
}

// GetAtOk returns a tuple with the At field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetAtOk() (*At, bool) {
	if o == nil || utils.IsNil(o.At) {
		return nil, false
	}
	return o.At, true
}

// HasAt returns a boolean if a field has been set.
func (o *Message) HasAt() bool {
	if o != nil && !utils.IsNil(o.At) {
		return true
	}

	return false
}

// SetAt gets a reference to the given At and assigns it to the at field.
// At:
func (o *Message) SetAt(v At) *Message {
	o.At = &v
	return o
}

// GetFeedCard returns the FeedCard field value if set, zero value otherwise.
func (o *Message) GetFeedCard() FeedCard {
	if o == nil || utils.IsNil(o.FeedCard) {
		var ret FeedCard
		return ret
	}
	return *o.FeedCard
}

// GetFeedCardOk returns a tuple with the FeedCard field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetFeedCardOk() (*FeedCard, bool) {
	if o == nil || utils.IsNil(o.FeedCard) {
		return nil, false
	}
	return o.FeedCard, true
}

// HasFeedCard returns a boolean if a field has been set.
func (o *Message) HasFeedCard() bool {
	if o != nil && !utils.IsNil(o.FeedCard) {
		return true
	}

	return false
}

// SetFeedCard gets a reference to the given FeedCard and assigns it to the feedCard field.
// FeedCard:
func (o *Message) SetFeedCard(v FeedCard) *Message {
	o.FeedCard = &v
	return o
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *Message) GetLink() Link {
	if o == nil || utils.IsNil(o.Link) {
		var ret Link
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetLinkOk() (*Link, bool) {
	if o == nil || utils.IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *Message) HasLink() bool {
	if o != nil && !utils.IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given Link and assigns it to the link field.
// Link:
func (o *Message) SetLink(v Link) *Message {
	o.Link = &v
	return o
}

// GetMarkdown returns the Markdown field value if set, zero value otherwise.
func (o *Message) GetMarkdown() Markdown {
	if o == nil || utils.IsNil(o.Markdown) {
		var ret Markdown
		return ret
	}
	return *o.Markdown
}

// GetMarkdownOk returns a tuple with the Markdown field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetMarkdownOk() (*Markdown, bool) {
	if o == nil || utils.IsNil(o.Markdown) {
		return nil, false
	}
	return o.Markdown, true
}

// HasMarkdown returns a boolean if a field has been set.
func (o *Message) HasMarkdown() bool {
	if o != nil && !utils.IsNil(o.Markdown) {
		return true
	}

	return false
}

// SetMarkdown gets a reference to the given Markdown and assigns it to the markdown field.
// Markdown:
func (o *Message) SetMarkdown(v Markdown) *Message {
	o.Markdown = &v
	return o
}

// GetMsgtype returns the Msgtype field value
func (o *Message) GetMsgtype() string {
	if o == nil {
		var ret string
		return ret
	}

	return *o.Msgtype
}

// GetMsgtypeOk returns a tuple with the Msgtype field value
// and a boolean to check if the value has been set.
func (o *Message) GetMsgtypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Msgtype, true
}

// SetMsgtype sets field value
func (o *Message) SetMsgtype(v string) *Message {
	o.Msgtype = &v
	return o
}

// GetText returns the Text field value if set, zero value otherwise.
func (o *Message) GetText() Text {
	if o == nil || utils.IsNil(o.Text) {
		var ret Text
		return ret
	}
	return *o.Text
}

// GetTextOk returns a tuple with the Text field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetTextOk() (*Text, bool) {
	if o == nil || utils.IsNil(o.Text) {
		return nil, false
	}
	return o.Text, true
}

// HasText returns a boolean if a field has been set.
func (o *Message) HasText() bool {
	if o != nil && !utils.IsNil(o.Text) {
		return true
	}

	return false
}

// SetText gets a reference to the given Text and assigns it to the text field.
// Text:
func (o *Message) SetText(v Text) *Message {
	o.Text = &v
	return o
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !utils.IsNil(o.ActionCard) {
		toSerialize["actionCard"] = o.ActionCard
	}
	if !utils.IsNil(o.At) {
		toSerialize["at"] = o.At
	}
	if !utils.IsNil(o.FeedCard) {
		toSerialize["feedCard"] = o.FeedCard
	}
	if !utils.IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	if !utils.IsNil(o.Markdown) {
		toSerialize["markdown"] = o.Markdown
	}
	toSerialize["msgtype"] = o.Msgtype
	if !utils.IsNil(o.Text) {
		toSerialize["text"] = o.Text
	}
	return toSerialize, nil
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v *NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v *NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
