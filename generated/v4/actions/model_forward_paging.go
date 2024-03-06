/*
Automation Actions V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"encoding/json"
)

// checks if the ForwardPaging type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ForwardPaging{}

// ForwardPaging struct for ForwardPaging
type ForwardPaging struct {
	Next *NextPage `json:"next,omitempty"`
}

// NewForwardPaging instantiates a new ForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForwardPaging() *ForwardPaging {
	this := ForwardPaging{}
	return &this
}

// NewForwardPagingWithDefaults instantiates a new ForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForwardPagingWithDefaults() *ForwardPaging {
	this := ForwardPaging{}
	return &this
}

// GetNext returns the Next field value if set, zero value otherwise.
func (o *ForwardPaging) GetNext() NextPage {
	if o == nil || IsNil(o.Next) {
		var ret NextPage
		return ret
	}
	return *o.Next
}

// GetNextOk returns a tuple with the Next field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ForwardPaging) GetNextOk() (*NextPage, bool) {
	if o == nil || IsNil(o.Next) {
		return nil, false
	}
	return o.Next, true
}

// HasNext returns a boolean if a field has been set.
func (o *ForwardPaging) HasNext() bool {
	if o != nil && !IsNil(o.Next) {
		return true
	}

	return false
}

// SetNext gets a reference to the given NextPage and assigns it to the Next field.
func (o *ForwardPaging) SetNext(v NextPage) {
	o.Next = &v
}

func (o ForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ForwardPaging) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Next) {
		toSerialize["next"] = o.Next
	}
	return toSerialize, nil
}

type NullableForwardPaging struct {
	value *ForwardPaging
	isSet bool
}

func (v NullableForwardPaging) Get() *ForwardPaging {
	return v.value
}

func (v *NullableForwardPaging) Set(val *ForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForwardPaging(val *ForwardPaging) *NullableForwardPaging {
	return &NullableForwardPaging{value: val, isSet: true}
}

func (v NullableForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
