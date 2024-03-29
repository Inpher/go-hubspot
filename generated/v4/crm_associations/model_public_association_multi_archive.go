/*
CrmPublicAssociationsServiceV4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_associations

import (
	"encoding/json"
)

// PublicAssociationMultiArchive struct for PublicAssociationMultiArchive
type PublicAssociationMultiArchive struct {
	From PublicObjectId   `json:"from"`
	To   []PublicObjectId `json:"to"`
}

// NewPublicAssociationMultiArchive instantiates a new PublicAssociationMultiArchive object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationMultiArchive(from PublicObjectId, to []PublicObjectId) *PublicAssociationMultiArchive {
	this := PublicAssociationMultiArchive{}
	this.From = from
	this.To = to
	return &this
}

// NewPublicAssociationMultiArchiveWithDefaults instantiates a new PublicAssociationMultiArchive object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationMultiArchiveWithDefaults() *PublicAssociationMultiArchive {
	this := PublicAssociationMultiArchive{}
	return &this
}

// GetFrom returns the From field value
func (o *PublicAssociationMultiArchive) GetFrom() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationMultiArchive) GetFromOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *PublicAssociationMultiArchive) SetFrom(v PublicObjectId) {
	o.From = v
}

// GetTo returns the To field value
func (o *PublicAssociationMultiArchive) GetTo() []PublicObjectId {
	if o == nil {
		var ret []PublicObjectId
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationMultiArchive) GetToOk() ([]PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return o.To, true
}

// SetTo sets field value
func (o *PublicAssociationMultiArchive) SetTo(v []PublicObjectId) {
	o.To = v
}

func (o PublicAssociationMultiArchive) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullablePublicAssociationMultiArchive struct {
	value *PublicAssociationMultiArchive
	isSet bool
}

func (v NullablePublicAssociationMultiArchive) Get() *PublicAssociationMultiArchive {
	return v.value
}

func (v *NullablePublicAssociationMultiArchive) Set(val *PublicAssociationMultiArchive) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationMultiArchive) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationMultiArchive) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationMultiArchive(val *PublicAssociationMultiArchive) *NullablePublicAssociationMultiArchive {
	return &NullablePublicAssociationMultiArchive{value: val, isSet: true}
}

func (v NullablePublicAssociationMultiArchive) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationMultiArchive) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
