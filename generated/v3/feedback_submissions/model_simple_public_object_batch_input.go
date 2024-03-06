/*
Feedback Submissions

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package feedback_submissions

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SimplePublicObjectBatchInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplePublicObjectBatchInput{}

// SimplePublicObjectBatchInput struct for SimplePublicObjectBatchInput
type SimplePublicObjectBatchInput struct {
	IdProperty *string           `json:"idProperty,omitempty"`
	Id         string            `json:"id"`
	Properties map[string]string `json:"properties"`
}

type _SimplePublicObjectBatchInput SimplePublicObjectBatchInput

// NewSimplePublicObjectBatchInput instantiates a new SimplePublicObjectBatchInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectBatchInput(id string, properties map[string]string) *SimplePublicObjectBatchInput {
	this := SimplePublicObjectBatchInput{}
	this.Id = id
	this.Properties = properties
	return &this
}

// NewSimplePublicObjectBatchInputWithDefaults instantiates a new SimplePublicObjectBatchInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectBatchInputWithDefaults() *SimplePublicObjectBatchInput {
	this := SimplePublicObjectBatchInput{}
	return &this
}

// GetIdProperty returns the IdProperty field value if set, zero value otherwise.
func (o *SimplePublicObjectBatchInput) GetIdProperty() string {
	if o == nil || IsNil(o.IdProperty) {
		var ret string
		return ret
	}
	return *o.IdProperty
}

// GetIdPropertyOk returns a tuple with the IdProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectBatchInput) GetIdPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.IdProperty) {
		return nil, false
	}
	return o.IdProperty, true
}

// HasIdProperty returns a boolean if a field has been set.
func (o *SimplePublicObjectBatchInput) HasIdProperty() bool {
	if o != nil && !IsNil(o.IdProperty) {
		return true
	}

	return false
}

// SetIdProperty gets a reference to the given string and assigns it to the IdProperty field.
func (o *SimplePublicObjectBatchInput) SetIdProperty(v string) {
	o.IdProperty = &v
}

// GetId returns the Id field value
func (o *SimplePublicObjectBatchInput) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectBatchInput) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObjectBatchInput) SetId(v string) {
	o.Id = v
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectBatchInput) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectBatchInput) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectBatchInput) SetProperties(v map[string]string) {
	o.Properties = v
}

func (o SimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplePublicObjectBatchInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdProperty) {
		toSerialize["idProperty"] = o.IdProperty
	}
	toSerialize["id"] = o.Id
	toSerialize["properties"] = o.Properties
	return toSerialize, nil
}

func (o *SimplePublicObjectBatchInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
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

	varSimplePublicObjectBatchInput := _SimplePublicObjectBatchInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimplePublicObjectBatchInput)

	if err != nil {
		return err
	}

	*o = SimplePublicObjectBatchInput(varSimplePublicObjectBatchInput)

	return err
}

type NullableSimplePublicObjectBatchInput struct {
	value *SimplePublicObjectBatchInput
	isSet bool
}

func (v NullableSimplePublicObjectBatchInput) Get() *SimplePublicObjectBatchInput {
	return v.value
}

func (v *NullableSimplePublicObjectBatchInput) Set(val *SimplePublicObjectBatchInput) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectBatchInput) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectBatchInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectBatchInput(val *SimplePublicObjectBatchInput) *NullableSimplePublicObjectBatchInput {
	return &NullableSimplePublicObjectBatchInput{value: val, isSet: true}
}

func (v NullableSimplePublicObjectBatchInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectBatchInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
