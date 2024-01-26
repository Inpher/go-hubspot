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

// CallbackCompletionRequest struct for CallbackCompletionRequest
type CallbackCompletionRequest struct {
	OutputFields map[string]string `json:"outputFields"`
}

// NewCallbackCompletionRequest instantiates a new CallbackCompletionRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallbackCompletionRequest(outputFields map[string]string) *CallbackCompletionRequest {
	this := CallbackCompletionRequest{}
	this.OutputFields = outputFields
	return &this
}

// NewCallbackCompletionRequestWithDefaults instantiates a new CallbackCompletionRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallbackCompletionRequestWithDefaults() *CallbackCompletionRequest {
	this := CallbackCompletionRequest{}
	return &this
}

// GetOutputFields returns the OutputFields field value
func (o *CallbackCompletionRequest) GetOutputFields() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.OutputFields
}

// GetOutputFieldsOk returns a tuple with the OutputFields field value
// and a boolean to check if the value has been set.
func (o *CallbackCompletionRequest) GetOutputFieldsOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OutputFields, true
}

// SetOutputFields sets field value
func (o *CallbackCompletionRequest) SetOutputFields(v map[string]string) {
	o.OutputFields = v
}

func (o CallbackCompletionRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["outputFields"] = o.OutputFields
	}
	return json.Marshal(toSerialize)
}

type NullableCallbackCompletionRequest struct {
	value *CallbackCompletionRequest
	isSet bool
}

func (v NullableCallbackCompletionRequest) Get() *CallbackCompletionRequest {
	return v.value
}

func (v *NullableCallbackCompletionRequest) Set(val *CallbackCompletionRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCallbackCompletionRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCallbackCompletionRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallbackCompletionRequest(val *CallbackCompletionRequest) *NullableCallbackCompletionRequest {
	return &NullableCallbackCompletionRequest{value: val, isSet: true}
}

func (v NullableCallbackCompletionRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallbackCompletionRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
