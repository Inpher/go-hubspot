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

// checks if the PublicObjectRequestOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicObjectRequestOptions{}

// PublicObjectRequestOptions struct for PublicObjectRequestOptions
type PublicObjectRequestOptions struct {
	Properties []string `json:"properties"`
}

type _PublicObjectRequestOptions PublicObjectRequestOptions

// NewPublicObjectRequestOptions instantiates a new PublicObjectRequestOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicObjectRequestOptions(properties []string) *PublicObjectRequestOptions {
	this := PublicObjectRequestOptions{}
	this.Properties = properties
	return &this
}

// NewPublicObjectRequestOptionsWithDefaults instantiates a new PublicObjectRequestOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicObjectRequestOptionsWithDefaults() *PublicObjectRequestOptions {
	this := PublicObjectRequestOptions{}
	return &this
}

// GetProperties returns the Properties field value
func (o *PublicObjectRequestOptions) GetProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *PublicObjectRequestOptions) GetPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *PublicObjectRequestOptions) SetProperties(v []string) {
	o.Properties = v
}

func (o PublicObjectRequestOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicObjectRequestOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

func (o *PublicObjectRequestOptions) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"properties",
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

	varPublicObjectRequestOptions := _PublicObjectRequestOptions{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicObjectRequestOptions)

	if err != nil {
		return err
	}

	*o = PublicObjectRequestOptions(varPublicObjectRequestOptions)

	return err
}

type NullablePublicObjectRequestOptions struct {
	value *PublicObjectRequestOptions
	isSet bool
}

func (v NullablePublicObjectRequestOptions) Get() *PublicObjectRequestOptions {
	return v.value
}

func (v *NullablePublicObjectRequestOptions) Set(val *PublicObjectRequestOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicObjectRequestOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicObjectRequestOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicObjectRequestOptions(val *PublicObjectRequestOptions) *NullablePublicObjectRequestOptions {
	return &NullablePublicObjectRequestOptions{value: val, isSet: true}
}

func (v NullablePublicObjectRequestOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicObjectRequestOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
