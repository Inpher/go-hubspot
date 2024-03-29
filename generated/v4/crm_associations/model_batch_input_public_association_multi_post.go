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

// BatchInputPublicAssociationMultiPost struct for BatchInputPublicAssociationMultiPost
type BatchInputPublicAssociationMultiPost struct {
	Inputs []PublicAssociationMultiPost `json:"inputs"`
}

// NewBatchInputPublicAssociationMultiPost instantiates a new BatchInputPublicAssociationMultiPost object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputPublicAssociationMultiPost(inputs []PublicAssociationMultiPost) *BatchInputPublicAssociationMultiPost {
	this := BatchInputPublicAssociationMultiPost{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputPublicAssociationMultiPostWithDefaults instantiates a new BatchInputPublicAssociationMultiPost object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputPublicAssociationMultiPostWithDefaults() *BatchInputPublicAssociationMultiPost {
	this := BatchInputPublicAssociationMultiPost{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputPublicAssociationMultiPost) GetInputs() []PublicAssociationMultiPost {
	if o == nil {
		var ret []PublicAssociationMultiPost
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputPublicAssociationMultiPost) GetInputsOk() ([]PublicAssociationMultiPost, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputPublicAssociationMultiPost) SetInputs(v []PublicAssociationMultiPost) {
	o.Inputs = v
}

func (o BatchInputPublicAssociationMultiPost) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["inputs"] = o.Inputs
	}
	return json.Marshal(toSerialize)
}

type NullableBatchInputPublicAssociationMultiPost struct {
	value *BatchInputPublicAssociationMultiPost
	isSet bool
}

func (v NullableBatchInputPublicAssociationMultiPost) Get() *BatchInputPublicAssociationMultiPost {
	return v.value
}

func (v *NullableBatchInputPublicAssociationMultiPost) Set(val *BatchInputPublicAssociationMultiPost) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputPublicAssociationMultiPost) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputPublicAssociationMultiPost) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputPublicAssociationMultiPost(val *BatchInputPublicAssociationMultiPost) *NullableBatchInputPublicAssociationMultiPost {
	return &NullableBatchInputPublicAssociationMultiPost{value: val, isSet: true}
}

func (v NullableBatchInputPublicAssociationMultiPost) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputPublicAssociationMultiPost) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
