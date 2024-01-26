/*
Quotes

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package quotes

import (
	"encoding/json"
)

// PublicAssociationsForObject struct for PublicAssociationsForObject
type PublicAssociationsForObject struct {
	Types []AssociationSpec `json:"types"`
	To    PublicObjectId    `json:"to"`
}

// NewPublicAssociationsForObject instantiates a new PublicAssociationsForObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationsForObject(types []AssociationSpec, to PublicObjectId) *PublicAssociationsForObject {
	this := PublicAssociationsForObject{}
	this.Types = types
	this.To = to
	return &this
}

// NewPublicAssociationsForObjectWithDefaults instantiates a new PublicAssociationsForObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationsForObjectWithDefaults() *PublicAssociationsForObject {
	this := PublicAssociationsForObject{}
	return &this
}

// GetTypes returns the Types field value
func (o *PublicAssociationsForObject) GetTypes() []AssociationSpec {
	if o == nil {
		var ret []AssociationSpec
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationsForObject) GetTypesOk() ([]AssociationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *PublicAssociationsForObject) SetTypes(v []AssociationSpec) {
	o.Types = v
}

// GetTo returns the To field value
func (o *PublicAssociationsForObject) GetTo() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationsForObject) GetToOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PublicAssociationsForObject) SetTo(v PublicObjectId) {
	o.To = v
}

func (o PublicAssociationsForObject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["types"] = o.Types
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullablePublicAssociationsForObject struct {
	value *PublicAssociationsForObject
	isSet bool
}

func (v NullablePublicAssociationsForObject) Get() *PublicAssociationsForObject {
	return v.value
}

func (v *NullablePublicAssociationsForObject) Set(val *PublicAssociationsForObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationsForObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationsForObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationsForObject(val *PublicAssociationsForObject) *NullablePublicAssociationsForObject {
	return &NullablePublicAssociationsForObject{value: val, isSet: true}
}

func (v NullablePublicAssociationsForObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationsForObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
