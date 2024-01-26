/*
Tags

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package tags

import (
	"encoding/json"
	"time"
)

// Tag Model definition for a Tag.
type Tag struct {
	// The timestamp (ISO8601 format) when this Blog Tag was deleted.
	DeletedAt time.Time `json:"deletedAt"`
	Created   time.Time `json:"created"`
	// The name of the tag.
	Name string `json:"name"`
	// The explicitly defined ISO 639 language code of the tag.
	Language string `json:"language"`
	// The unique ID of the Blog Tag.
	Id string `json:"id"`
	// ID of the primary tag this object was translated from.
	TranslatedFromId int64     `json:"translatedFromId"`
	Updated          time.Time `json:"updated"`
}

// NewTag instantiates a new Tag object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTag(deletedAt time.Time, created time.Time, name string, language string, id string, translatedFromId int64, updated time.Time) *Tag {
	this := Tag{}
	this.DeletedAt = deletedAt
	this.Created = created
	this.Name = name
	this.Language = language
	this.Id = id
	this.TranslatedFromId = translatedFromId
	this.Updated = updated
	return &this
}

// NewTagWithDefaults instantiates a new Tag object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTagWithDefaults() *Tag {
	this := Tag{}
	return &this
}

// GetDeletedAt returns the DeletedAt field value
func (o *Tag) GetDeletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.DeletedAt
}

// GetDeletedAtOk returns a tuple with the DeletedAt field value
// and a boolean to check if the value has been set.
func (o *Tag) GetDeletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DeletedAt, true
}

// SetDeletedAt sets field value
func (o *Tag) SetDeletedAt(v time.Time) {
	o.DeletedAt = v
}

// GetCreated returns the Created field value
func (o *Tag) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *Tag) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *Tag) SetCreated(v time.Time) {
	o.Created = v
}

// GetName returns the Name field value
func (o *Tag) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Tag) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Tag) SetName(v string) {
	o.Name = v
}

// GetLanguage returns the Language field value
func (o *Tag) GetLanguage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Language
}

// GetLanguageOk returns a tuple with the Language field value
// and a boolean to check if the value has been set.
func (o *Tag) GetLanguageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Language, true
}

// SetLanguage sets field value
func (o *Tag) SetLanguage(v string) {
	o.Language = v
}

// GetId returns the Id field value
func (o *Tag) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Tag) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Tag) SetId(v string) {
	o.Id = v
}

// GetTranslatedFromId returns the TranslatedFromId field value
func (o *Tag) GetTranslatedFromId() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.TranslatedFromId
}

// GetTranslatedFromIdOk returns a tuple with the TranslatedFromId field value
// and a boolean to check if the value has been set.
func (o *Tag) GetTranslatedFromIdOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TranslatedFromId, true
}

// SetTranslatedFromId sets field value
func (o *Tag) SetTranslatedFromId(v int64) {
	o.TranslatedFromId = v
}

// GetUpdated returns the Updated field value
func (o *Tag) GetUpdated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *Tag) GetUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *Tag) SetUpdated(v time.Time) {
	o.Updated = v
}

func (o Tag) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["deletedAt"] = o.DeletedAt
	}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["language"] = o.Language
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["translatedFromId"] = o.TranslatedFromId
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableTag struct {
	value *Tag
	isSet bool
}

func (v NullableTag) Get() *Tag {
	return v.value
}

func (v *NullableTag) Set(val *Tag) {
	v.value = val
	v.isSet = true
}

func (v NullableTag) IsSet() bool {
	return v.isSet
}

func (v *NullableTag) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTag(val *Tag) *NullableTag {
	return &NullableTag{value: val, isSet: true}
}

func (v NullableTag) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTag) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
