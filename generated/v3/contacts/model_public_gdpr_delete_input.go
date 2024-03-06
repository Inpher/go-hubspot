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

// checks if the PublicGdprDeleteInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicGdprDeleteInput{}

// PublicGdprDeleteInput struct for PublicGdprDeleteInput
type PublicGdprDeleteInput struct {
	IdProperty *string `json:"idProperty,omitempty"`
	ObjectId   string  `json:"objectId"`
}

type _PublicGdprDeleteInput PublicGdprDeleteInput

// NewPublicGdprDeleteInput instantiates a new PublicGdprDeleteInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicGdprDeleteInput(objectId string) *PublicGdprDeleteInput {
	this := PublicGdprDeleteInput{}
	this.ObjectId = objectId
	return &this
}

// NewPublicGdprDeleteInputWithDefaults instantiates a new PublicGdprDeleteInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicGdprDeleteInputWithDefaults() *PublicGdprDeleteInput {
	this := PublicGdprDeleteInput{}
	return &this
}

// GetIdProperty returns the IdProperty field value if set, zero value otherwise.
func (o *PublicGdprDeleteInput) GetIdProperty() string {
	if o == nil || IsNil(o.IdProperty) {
		var ret string
		return ret
	}
	return *o.IdProperty
}

// GetIdPropertyOk returns a tuple with the IdProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicGdprDeleteInput) GetIdPropertyOk() (*string, bool) {
	if o == nil || IsNil(o.IdProperty) {
		return nil, false
	}
	return o.IdProperty, true
}

// HasIdProperty returns a boolean if a field has been set.
func (o *PublicGdprDeleteInput) HasIdProperty() bool {
	if o != nil && !IsNil(o.IdProperty) {
		return true
	}

	return false
}

// SetIdProperty gets a reference to the given string and assigns it to the IdProperty field.
func (o *PublicGdprDeleteInput) SetIdProperty(v string) {
	o.IdProperty = &v
}

// GetObjectId returns the ObjectId field value
func (o *PublicGdprDeleteInput) GetObjectId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectId
}

// GetObjectIdOk returns a tuple with the ObjectId field value
// and a boolean to check if the value has been set.
func (o *PublicGdprDeleteInput) GetObjectIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectId, true
}

// SetObjectId sets field value
func (o *PublicGdprDeleteInput) SetObjectId(v string) {
	o.ObjectId = v
}

func (o PublicGdprDeleteInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicGdprDeleteInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.IdProperty) {
		toSerialize["idProperty"] = o.IdProperty
	}
	toSerialize["objectId"] = o.ObjectId
	return toSerialize, nil
}

func (o *PublicGdprDeleteInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"objectId",
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

	varPublicGdprDeleteInput := _PublicGdprDeleteInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicGdprDeleteInput)

	if err != nil {
		return err
	}

	*o = PublicGdprDeleteInput(varPublicGdprDeleteInput)

	return err
}

type NullablePublicGdprDeleteInput struct {
	value *PublicGdprDeleteInput
	isSet bool
}

func (v NullablePublicGdprDeleteInput) Get() *PublicGdprDeleteInput {
	return v.value
}

func (v *NullablePublicGdprDeleteInput) Set(val *PublicGdprDeleteInput) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicGdprDeleteInput) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicGdprDeleteInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicGdprDeleteInput(val *PublicGdprDeleteInput) *NullablePublicGdprDeleteInput {
	return &NullablePublicGdprDeleteInput{value: val, isSet: true}
}

func (v NullablePublicGdprDeleteInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicGdprDeleteInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
