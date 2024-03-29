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

// PublicDefaultAssociation struct for PublicDefaultAssociation
type PublicDefaultAssociation struct {
	From            PublicObjectId  `json:"from"`
	To              PublicObjectId  `json:"to"`
	AssociationSpec AssociationSpec `json:"associationSpec"`
}

// NewPublicDefaultAssociation instantiates a new PublicDefaultAssociation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicDefaultAssociation(from PublicObjectId, to PublicObjectId, associationSpec AssociationSpec) *PublicDefaultAssociation {
	this := PublicDefaultAssociation{}
	this.From = from
	this.To = to
	this.AssociationSpec = associationSpec
	return &this
}

// NewPublicDefaultAssociationWithDefaults instantiates a new PublicDefaultAssociation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicDefaultAssociationWithDefaults() *PublicDefaultAssociation {
	this := PublicDefaultAssociation{}
	return &this
}

// GetFrom returns the From field value
func (o *PublicDefaultAssociation) GetFrom() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.From
}

// GetFromOk returns a tuple with the From field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociation) GetFromOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.From, true
}

// SetFrom sets field value
func (o *PublicDefaultAssociation) SetFrom(v PublicObjectId) {
	o.From = v
}

// GetTo returns the To field value
func (o *PublicDefaultAssociation) GetTo() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociation) GetToOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PublicDefaultAssociation) SetTo(v PublicObjectId) {
	o.To = v
}

// GetAssociationSpec returns the AssociationSpec field value
func (o *PublicDefaultAssociation) GetAssociationSpec() AssociationSpec {
	if o == nil {
		var ret AssociationSpec
		return ret
	}

	return o.AssociationSpec
}

// GetAssociationSpecOk returns a tuple with the AssociationSpec field value
// and a boolean to check if the value has been set.
func (o *PublicDefaultAssociation) GetAssociationSpecOk() (*AssociationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AssociationSpec, true
}

// SetAssociationSpec sets field value
func (o *PublicDefaultAssociation) SetAssociationSpec(v AssociationSpec) {
	o.AssociationSpec = v
}

func (o PublicDefaultAssociation) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["from"] = o.From
	}
	if true {
		toSerialize["to"] = o.To
	}
	if true {
		toSerialize["associationSpec"] = o.AssociationSpec
	}
	return json.Marshal(toSerialize)
}

type NullablePublicDefaultAssociation struct {
	value *PublicDefaultAssociation
	isSet bool
}

func (v NullablePublicDefaultAssociation) Get() *PublicDefaultAssociation {
	return v.value
}

func (v *NullablePublicDefaultAssociation) Set(val *PublicDefaultAssociation) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicDefaultAssociation) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicDefaultAssociation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicDefaultAssociation(val *PublicDefaultAssociation) *NullablePublicDefaultAssociation {
	return &NullablePublicDefaultAssociation{value: val, isSet: true}
}

func (v NullablePublicDefaultAssociation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicDefaultAssociation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
