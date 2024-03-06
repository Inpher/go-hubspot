/*
Transactional Single Send

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactional

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the PublicSingleSendEmail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicSingleSendEmail{}

// PublicSingleSendEmail A JSON object containing anything you want to override.
type PublicSingleSendEmail struct {
	// List of email addresses to send as Cc.
	Cc []string `json:"cc,omitempty"`
	// ID for a particular send. No more than one email will be sent per sendId.
	SendId *string `json:"sendId,omitempty"`
	// List of email addresses to send as Bcc.
	Bcc []string `json:"bcc,omitempty"`
	// List of Reply-To header values for the email.
	ReplyTo []string `json:"replyTo,omitempty"`
	// The From header for the email.
	From *string `json:"from,omitempty"`
	// The recipient of the email.
	To string `json:"to"`
}

type _PublicSingleSendEmail PublicSingleSendEmail

// NewPublicSingleSendEmail instantiates a new PublicSingleSendEmail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicSingleSendEmail(to string) *PublicSingleSendEmail {
	this := PublicSingleSendEmail{}
	this.To = to
	return &this
}

// NewPublicSingleSendEmailWithDefaults instantiates a new PublicSingleSendEmail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicSingleSendEmailWithDefaults() *PublicSingleSendEmail {
	this := PublicSingleSendEmail{}
	return &this
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetCc() []string {
	if o == nil || IsNil(o.Cc) {
		var ret []string
		return ret
	}
	return o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetCcOk() ([]string, bool) {
	if o == nil || IsNil(o.Cc) {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasCc() bool {
	if o != nil && !IsNil(o.Cc) {
		return true
	}

	return false
}

// SetCc gets a reference to the given []string and assigns it to the Cc field.
func (o *PublicSingleSendEmail) SetCc(v []string) {
	o.Cc = v
}

// GetSendId returns the SendId field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetSendId() string {
	if o == nil || IsNil(o.SendId) {
		var ret string
		return ret
	}
	return *o.SendId
}

// GetSendIdOk returns a tuple with the SendId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetSendIdOk() (*string, bool) {
	if o == nil || IsNil(o.SendId) {
		return nil, false
	}
	return o.SendId, true
}

// HasSendId returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasSendId() bool {
	if o != nil && !IsNil(o.SendId) {
		return true
	}

	return false
}

// SetSendId gets a reference to the given string and assigns it to the SendId field.
func (o *PublicSingleSendEmail) SetSendId(v string) {
	o.SendId = &v
}

// GetBcc returns the Bcc field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetBcc() []string {
	if o == nil || IsNil(o.Bcc) {
		var ret []string
		return ret
	}
	return o.Bcc
}

// GetBccOk returns a tuple with the Bcc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetBccOk() ([]string, bool) {
	if o == nil || IsNil(o.Bcc) {
		return nil, false
	}
	return o.Bcc, true
}

// HasBcc returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasBcc() bool {
	if o != nil && !IsNil(o.Bcc) {
		return true
	}

	return false
}

// SetBcc gets a reference to the given []string and assigns it to the Bcc field.
func (o *PublicSingleSendEmail) SetBcc(v []string) {
	o.Bcc = v
}

// GetReplyTo returns the ReplyTo field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetReplyTo() []string {
	if o == nil || IsNil(o.ReplyTo) {
		var ret []string
		return ret
	}
	return o.ReplyTo
}

// GetReplyToOk returns a tuple with the ReplyTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetReplyToOk() ([]string, bool) {
	if o == nil || IsNil(o.ReplyTo) {
		return nil, false
	}
	return o.ReplyTo, true
}

// HasReplyTo returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasReplyTo() bool {
	if o != nil && !IsNil(o.ReplyTo) {
		return true
	}

	return false
}

// SetReplyTo gets a reference to the given []string and assigns it to the ReplyTo field.
func (o *PublicSingleSendEmail) SetReplyTo(v []string) {
	o.ReplyTo = v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *PublicSingleSendEmail) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *PublicSingleSendEmail) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *PublicSingleSendEmail) SetFrom(v string) {
	o.From = &v
}

// GetTo returns the To field value
func (o *PublicSingleSendEmail) GetTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicSingleSendEmail) GetToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PublicSingleSendEmail) SetTo(v string) {
	o.To = v
}

func (o PublicSingleSendEmail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicSingleSendEmail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Cc) {
		toSerialize["cc"] = o.Cc
	}
	if !IsNil(o.SendId) {
		toSerialize["sendId"] = o.SendId
	}
	if !IsNil(o.Bcc) {
		toSerialize["bcc"] = o.Bcc
	}
	if !IsNil(o.ReplyTo) {
		toSerialize["replyTo"] = o.ReplyTo
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	toSerialize["to"] = o.To
	return toSerialize, nil
}

func (o *PublicSingleSendEmail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"to",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varPublicSingleSendEmail := _PublicSingleSendEmail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicSingleSendEmail)

	if err != nil {
		return err
	}

	*o = PublicSingleSendEmail(varPublicSingleSendEmail)

	return err
}

type NullablePublicSingleSendEmail struct {
	value *PublicSingleSendEmail
	isSet bool
}

func (v NullablePublicSingleSendEmail) Get() *PublicSingleSendEmail {
	return v.value
}

func (v *NullablePublicSingleSendEmail) Set(val *PublicSingleSendEmail) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicSingleSendEmail) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicSingleSendEmail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicSingleSendEmail(val *PublicSingleSendEmail) *NullablePublicSingleSendEmail {
	return &NullablePublicSingleSendEmail{value: val, isSet: true}
}

func (v NullablePublicSingleSendEmail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicSingleSendEmail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
