/*
Authors

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package authors

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SetNewLanguagePrimaryRequestVNext type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SetNewLanguagePrimaryRequestVNext{}

// SetNewLanguagePrimaryRequestVNext Request body object for setting a new primary language.
type SetNewLanguagePrimaryRequestVNext struct {
	// ID of object to set as primary in multi-language group.
	Id string `json:"id"`
}

type _SetNewLanguagePrimaryRequestVNext SetNewLanguagePrimaryRequestVNext

// NewSetNewLanguagePrimaryRequestVNext instantiates a new SetNewLanguagePrimaryRequestVNext object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSetNewLanguagePrimaryRequestVNext(id string) *SetNewLanguagePrimaryRequestVNext {
	this := SetNewLanguagePrimaryRequestVNext{}
	this.Id = id
	return &this
}

// NewSetNewLanguagePrimaryRequestVNextWithDefaults instantiates a new SetNewLanguagePrimaryRequestVNext object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSetNewLanguagePrimaryRequestVNextWithDefaults() *SetNewLanguagePrimaryRequestVNext {
	this := SetNewLanguagePrimaryRequestVNext{}
	return &this
}

// GetId returns the Id field value
func (o *SetNewLanguagePrimaryRequestVNext) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SetNewLanguagePrimaryRequestVNext) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SetNewLanguagePrimaryRequestVNext) SetId(v string) {
	o.Id = v
}

func (o SetNewLanguagePrimaryRequestVNext) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SetNewLanguagePrimaryRequestVNext) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *SetNewLanguagePrimaryRequestVNext) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varSetNewLanguagePrimaryRequestVNext := _SetNewLanguagePrimaryRequestVNext{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSetNewLanguagePrimaryRequestVNext)

	if err != nil {
		return err
	}

	*o = SetNewLanguagePrimaryRequestVNext(varSetNewLanguagePrimaryRequestVNext)

	return err
}

type NullableSetNewLanguagePrimaryRequestVNext struct {
	value *SetNewLanguagePrimaryRequestVNext
	isSet bool
}

func (v NullableSetNewLanguagePrimaryRequestVNext) Get() *SetNewLanguagePrimaryRequestVNext {
	return v.value
}

func (v *NullableSetNewLanguagePrimaryRequestVNext) Set(val *SetNewLanguagePrimaryRequestVNext) {
	v.value = val
	v.isSet = true
}

func (v NullableSetNewLanguagePrimaryRequestVNext) IsSet() bool {
	return v.isSet
}

func (v *NullableSetNewLanguagePrimaryRequestVNext) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSetNewLanguagePrimaryRequestVNext(val *SetNewLanguagePrimaryRequestVNext) *NullableSetNewLanguagePrimaryRequestVNext {
	return &NullableSetNewLanguagePrimaryRequestVNext{value: val, isSet: true}
}

func (v NullableSetNewLanguagePrimaryRequestVNext) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSetNewLanguagePrimaryRequestVNext) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
