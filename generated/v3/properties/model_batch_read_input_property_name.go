/*
Properties

All HubSpot objects store data in default and custom properties. These endpoints provide access to read and modify object properties in HubSpot.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package properties

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the BatchReadInputPropertyName type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BatchReadInputPropertyName{}

// BatchReadInputPropertyName struct for BatchReadInputPropertyName
type BatchReadInputPropertyName struct {
	Archived bool           `json:"archived"`
	Inputs   []PropertyName `json:"inputs"`
}

type _BatchReadInputPropertyName BatchReadInputPropertyName

// NewBatchReadInputPropertyName instantiates a new BatchReadInputPropertyName object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchReadInputPropertyName(archived bool, inputs []PropertyName) *BatchReadInputPropertyName {
	this := BatchReadInputPropertyName{}
	this.Archived = archived
	this.Inputs = inputs
	return &this
}

// NewBatchReadInputPropertyNameWithDefaults instantiates a new BatchReadInputPropertyName object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchReadInputPropertyNameWithDefaults() *BatchReadInputPropertyName {
	this := BatchReadInputPropertyName{}
	return &this
}

// GetArchived returns the Archived field value
func (o *BatchReadInputPropertyName) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputPropertyName) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *BatchReadInputPropertyName) SetArchived(v bool) {
	o.Archived = v
}

// GetInputs returns the Inputs field value
func (o *BatchReadInputPropertyName) GetInputs() []PropertyName {
	if o == nil {
		var ret []PropertyName
		return ret
	}

	return o.Inputs
}

// GetInputsOk returns a tuple with the Inputs field value
// and a boolean to check if the value has been set.
func (o *BatchReadInputPropertyName) GetInputsOk() ([]PropertyName, bool) {
	if o == nil {
		return nil, false
	}
	return o.Inputs, true
}

// SetInputs sets field value
func (o *BatchReadInputPropertyName) SetInputs(v []PropertyName) {
	o.Inputs = v
}

func (o BatchReadInputPropertyName) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BatchReadInputPropertyName) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["archived"] = o.Archived
	toSerialize["inputs"] = o.Inputs
	return toSerialize, nil
}

func (o *BatchReadInputPropertyName) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"archived",
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

	varBatchReadInputPropertyName := _BatchReadInputPropertyName{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varBatchReadInputPropertyName)

	if err != nil {
		return err
	}

	*o = BatchReadInputPropertyName(varBatchReadInputPropertyName)

	return err
}

type NullableBatchReadInputPropertyName struct {
	value *BatchReadInputPropertyName
	isSet bool
}

func (v NullableBatchReadInputPropertyName) Get() *BatchReadInputPropertyName {
	return v.value
}

func (v *NullableBatchReadInputPropertyName) Set(val *BatchReadInputPropertyName) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchReadInputPropertyName) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchReadInputPropertyName) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchReadInputPropertyName(val *BatchReadInputPropertyName) *NullableBatchReadInputPropertyName {
	return &NullableBatchReadInputPropertyName{value: val, isSet: true}
}

func (v NullableBatchReadInputPropertyName) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchReadInputPropertyName) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
