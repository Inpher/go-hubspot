/*
Posts

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AttachToLangPrimaryRequestVNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AttachToLangPrimaryRequestVNext{}

// AttachToLangPrimaryRequestVNext Request body object for attaching objects to multi-language groups.
type AttachToLangPrimaryRequestVNext struct {
	// Designated language of the object to add to a multi-language group.
	Language string `json:"language"`
	// ID of the object to add to a multi-language group.
	Id string `json:"id"`
	// ID of primary language object in multi-language group.
	PrimaryId string `json:"primaryId"`
	// Primary language of the multi-language group.
	PrimaryLanguage *string `json:"primaryLanguage,omitempty"`
}

type _AttachToLangPrimaryRequestVNext AttachToLangPrimaryRequestVNext

// NewAttachToLangPrimaryRequestVNext instantiates a new AttachToLangPrimaryRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttachToLangPrimaryRequestVNext(language string, id string, primaryId string) *AttachToLangPrimaryRequestVNext {
	this := AttachToLangPrimaryRequestVNext{}
	this.Language = language
	this.Id = id
	this.PrimaryId = primaryId
	return &this
}

// NewAttachToLangPrimaryRequestVNextWithDefaults instantiates a new AttachToLangPrimaryRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttachToLangPrimaryRequestVNextWithDefaults() *AttachToLangPrimaryRequestVNext {
	this := AttachToLangPrimaryRequestVNext{}
	return &this
}

// GetLanguage returns the Language field value
func (o *AttachToLangPrimaryRequestVNext) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *AttachToLangPrimaryRequestVNext) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *AttachToLangPrimaryRequestVNext) SetLanguage(v string) {
	o.Language = v
}

// GetId returns the Id field value
func (o *AttachToLangPrimaryRequestVNext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AttachToLangPrimaryRequestVNext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AttachToLangPrimaryRequestVNext) SetId(v string) {
	o.Id = v
}

// GetPrimaryId returns the PrimaryId field value
func (o *AttachToLangPrimaryRequestVNext) GetPrimaryId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PrimaryId
}

// GetPrimaryIdOk returns a tuple with the PrimaryId field value
// and a boolean to check if the value has been set.
func (o *AttachToLangPrimaryRequestVNext) GetPrimaryIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PrimaryId, true
}

// SetPrimaryId sets field value
func (o *AttachToLangPrimaryRequestVNext) SetPrimaryId(v string) {
	o.PrimaryId = v
}

// GetPrimaryLanguage returns the PrimaryLanguage field value if set, zero value otherwise.
func (o *AttachToLangPrimaryRequestVNext) GetPrimaryLanguage() string {
	if o == nil || IsNil(o.PrimaryLanguage) {
		var ret string
		return ret
	}
	return *o.PrimaryLanguage
}

// GetPrimaryLanguageOk returns a tuple with the PrimaryLanguage field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttachToLangPrimaryRequestVNext) GetPrimaryLanguageOk() (*string, bool) {
	if o == nil || IsNil(o.PrimaryLanguage) {
		return nil, false
	}
	return o.PrimaryLanguage, true
}

// HasPrimaryLanguage returns a boolean if a field has been set.
func (o *AttachToLangPrimaryRequestVNext) HasPrimaryLanguage() bool {
	if o != nil && !IsNil(o.PrimaryLanguage) {
		return true
	}

	return false
}

// SetPrimaryLanguage gets a reference to the given string and assigns it to the PrimaryLanguage field.
func (o *AttachToLangPrimaryRequestVNext) SetPrimaryLanguage(v string) {
	o.PrimaryLanguage = &v
}

func (o AttachToLangPrimaryRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AttachToLangPrimaryRequestVNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["language"] = o.Language
	toSerialize["id"] = o.Id
	toSerialize["primaryId"] = o.PrimaryId
	if !IsNil(o.PrimaryLanguage) {
		toSerialize["primaryLanguage"] = o.PrimaryLanguage
	}
	return toSerialize, nil
}

func (o *AttachToLangPrimaryRequestVNext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"language",
		"id",
		"primaryId",
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

	varAttachToLangPrimaryRequestVNext := _AttachToLangPrimaryRequestVNext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAttachToLangPrimaryRequestVNext)

	if err != nil {
		return err
	}

	*o = AttachToLangPrimaryRequestVNext(varAttachToLangPrimaryRequestVNext)

	return err
}

type NullableAttachToLangPrimaryRequestVNext struct {
	value *AttachToLangPrimaryRequestVNext
	isSet bool
}

func (v NullableAttachToLangPrimaryRequestVNext) Get() *AttachToLangPrimaryRequestVNext {
	return v.value
}

func (v *NullableAttachToLangPrimaryRequestVNext) Set(val *AttachToLangPrimaryRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableAttachToLangPrimaryRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableAttachToLangPrimaryRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttachToLangPrimaryRequestVNext(val *AttachToLangPrimaryRequestVNext) *NullableAttachToLangPrimaryRequestVNext {
	return &NullableAttachToLangPrimaryRequestVNext{value: val, isSet: true}
}

func (v NullableAttachToLangPrimaryRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttachToLangPrimaryRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
