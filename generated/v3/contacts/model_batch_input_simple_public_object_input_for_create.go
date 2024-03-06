/*
Contacts

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package contacts

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchInputSimplePublicObjectInputForCreate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchInputSimplePublicObjectInputForCreate{}

// BatchInputSimplePublicObjectInputForCreate struct for BatchInputSimplePublicObjectInputForCreate
type BatchInputSimplePublicObjectInputForCreate struct {
	Inputs []SimplePublicObjectInputForCreate `json:"inputs"`
}

type _BatchInputSimplePublicObjectInputForCreate BatchInputSimplePublicObjectInputForCreate

// NewBatchInputSimplePublicObjectInputForCreate instantiates a new BatchInputSimplePublicObjectInputForCreate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchInputSimplePublicObjectInputForCreate(inputs []SimplePublicObjectInputForCreate) *BatchInputSimplePublicObjectInputForCreate {
	this := BatchInputSimplePublicObjectInputForCreate{}
	this.Inputs = inputs
	return &this
}

// NewBatchInputSimplePublicObjectInputForCreateWithDefaults instantiates a new BatchInputSimplePublicObjectInputForCreate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchInputSimplePublicObjectInputForCreateWithDefaults() *BatchInputSimplePublicObjectInputForCreate {
	this := BatchInputSimplePublicObjectInputForCreate{}
	return &this
}

// GetInputs returns the Inputs field value
func (o *BatchInputSimplePublicObjectInputForCreate) GetInputs() []SimplePublicObjectInputForCreate {
	if o == nil {
		var ret []SimplePublicObjectInputForCreate
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchInputSimplePublicObjectInputForCreate) GetInputsOk() ([]SimplePublicObjectInputForCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchInputSimplePublicObjectInputForCreate) SetInputs(v []SimplePublicObjectInputForCreate) {
	o.Inputs = v
}

func (o BatchInputSimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchInputSimplePublicObjectInputForCreate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchInputSimplePublicObjectInputForCreate) UnmarshalJSON(data []byte) (err error) {
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

	varBatchInputSimplePublicObjectInputForCreate := _BatchInputSimplePublicObjectInputForCreate{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchInputSimplePublicObjectInputForCreate)

	if err != nil {
		return err
	}

	*o = BatchInputSimplePublicObjectInputForCreate(varBatchInputSimplePublicObjectInputForCreate)

	return err
}

type NullableBatchInputSimplePublicObjectInputForCreate struct {
	value *BatchInputSimplePublicObjectInputForCreate
	isSet bool
}

func (v NullableBatchInputSimplePublicObjectInputForCreate) Get() *BatchInputSimplePublicObjectInputForCreate {
	return v.value
}

func (v *NullableBatchInputSimplePublicObjectInputForCreate) Set(val *BatchInputSimplePublicObjectInputForCreate) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchInputSimplePublicObjectInputForCreate) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchInputSimplePublicObjectInputForCreate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchInputSimplePublicObjectInputForCreate(val *BatchInputSimplePublicObjectInputForCreate) *NullableBatchInputSimplePublicObjectInputForCreate {
	return &NullableBatchInputSimplePublicObjectInputForCreate{value: val, isSet: true}
}

func (v NullableBatchInputSimplePublicObjectInputForCreate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchInputSimplePublicObjectInputForCreate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
