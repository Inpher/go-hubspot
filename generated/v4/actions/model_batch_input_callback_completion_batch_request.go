/*
Automation Actions V4

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v4
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package actions

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchInputCallbackCompletionBatchRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputCallbackCompletionBatchRequest{}

// BatchInputCallbackCompletionBatchRequest struct for BatchInputCallbackCompletionBatchRequest
type BatchInputCallbackCompletionBatchRequest struct {
	Inputs []CallbackCompletionBatchRequest `json:"inputs"`
}

type _BatchInputCallbackCompletionBatchRequest BatchInputCallbackCompletionBatchRequest

// NewBatchInputCallbackCompletionBatchRequest instantiates a new BatchInputCallbackCompletionBatchRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputCallbackCompletionBatchRequest(inputs []CallbackCompletionBatchRequest) *BatchInputCallbackCompletionBatchRequest {
	this := BatchInputCallbackCompletionBatchRequest{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputCallbackCompletionBatchRequestWithDefaults instantiates a new BatchInputCallbackCompletionBatchRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputCallbackCompletionBatchRequestWithDefaults() *BatchInputCallbackCompletionBatchRequest {
	this := BatchInputCallbackCompletionBatchRequest{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputCallbackCompletionBatchRequest) GetInputs() []CallbackCompletionBatchRequest {
	if o == nil {
		var ret []CallbackCompletionBatchRequest
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputCallbackCompletionBatchRequest) GetInputsOk() ([]CallbackCompletionBatchRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputCallbackCompletionBatchRequest) SetInputs(v []CallbackCompletionBatchRequest) {
	o.Inputs = v
}

func (o BatchInputCallbackCompletionBatchRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputCallbackCompletionBatchRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputCallbackCompletionBatchRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"inputs",
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

	varBatchInputCallbackCompletionBatchRequest := _BatchInputCallbackCompletionBatchRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputCallbackCompletionBatchRequest)

	if err != nil {
		return err
	}

	*o = BatchInputCallbackCompletionBatchRequest(varBatchInputCallbackCompletionBatchRequest)

	return err
}

type NullableBatchInputCallbackCompletionBatchRequest struct {
	value *BatchInputCallbackCompletionBatchRequest
	isSet bool
}

func (v NullableBatchInputCallbackCompletionBatchRequest) Get() *BatchInputCallbackCompletionBatchRequest {
	return v.value
}

func (v *NullableBatchInputCallbackCompletionBatchRequest) Set(val *BatchInputCallbackCompletionBatchRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputCallbackCompletionBatchRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputCallbackCompletionBatchRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputCallbackCompletionBatchRequest(val *BatchInputCallbackCompletionBatchRequest) *NullableBatchInputCallbackCompletionBatchRequest {
	return &NullableBatchInputCallbackCompletionBatchRequest{value: val, isSet: true}
}

func (v NullableBatchInputCallbackCompletionBatchRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputCallbackCompletionBatchRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
