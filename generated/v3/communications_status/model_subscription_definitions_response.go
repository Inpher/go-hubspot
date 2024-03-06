/*
Subscriptions

Subscriptions allow contacts to control what forms of communications they receive. Contacts can decide whether they want to receive communication pertaining to a specific topic, brand, or an entire HubSpot account.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package communications_status

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SubscriptionDefinitionsResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionDefinitionsResponse{}

// SubscriptionDefinitionsResponse A collection of subscription definitions for the portal.
type SubscriptionDefinitionsResponse struct {
	// A list of all subscription definitions.
	SubscriptionDefinitions []SubscriptionDefinition `json:"subscriptionDefinitions"`
}

type _SubscriptionDefinitionsResponse SubscriptionDefinitionsResponse

// NewSubscriptionDefinitionsResponse instantiates a new SubscriptionDefinitionsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionDefinitionsResponse(subscriptionDefinitions []SubscriptionDefinition) *SubscriptionDefinitionsResponse {
	this := SubscriptionDefinitionsResponse{}
	this.SubscriptionDefinitions = subscriptionDefinitions
	return &this
}

// NewSubscriptionDefinitionsResponseWithDefaults instantiates a new SubscriptionDefinitionsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionDefinitionsResponseWithDefaults() *SubscriptionDefinitionsResponse {
	this := SubscriptionDefinitionsResponse{}
	return &this
}

// GetSubscriptionDefinitions returns the SubscriptionDefinitions field value
func (o *SubscriptionDefinitionsResponse) GetSubscriptionDefinitions() []SubscriptionDefinition {
	if o == nil {
		var ret []SubscriptionDefinition
		return ret
	}

	return o.SubscriptionDefinitions
}

// GetSubscriptionDefinitionsOk returns a tuple with the SubscriptionDefinitions field value
// and a boolean to check if the value has been set.
func (o *SubscriptionDefinitionsResponse) GetSubscriptionDefinitionsOk() ([]SubscriptionDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.SubscriptionDefinitions, true
}

// SetSubscriptionDefinitions sets field value
func (o *SubscriptionDefinitionsResponse) SetSubscriptionDefinitions(v []SubscriptionDefinition) {
	o.SubscriptionDefinitions = v
}

func (o SubscriptionDefinitionsResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionDefinitionsResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["subscriptionDefinitions"] = o.SubscriptionDefinitions
	return toSerialize, nil
}

func (o *SubscriptionDefinitionsResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"subscriptionDefinitions",
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

	varSubscriptionDefinitionsResponse := _SubscriptionDefinitionsResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriptionDefinitionsResponse)

	if err != nil {
		return err
	}

	*o = SubscriptionDefinitionsResponse(varSubscriptionDefinitionsResponse)

	return err
}

type NullableSubscriptionDefinitionsResponse struct {
	value *SubscriptionDefinitionsResponse
	isSet bool
}

func (v NullableSubscriptionDefinitionsResponse) Get() *SubscriptionDefinitionsResponse {
	return v.value
}

func (v *NullableSubscriptionDefinitionsResponse) Set(val *SubscriptionDefinitionsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionDefinitionsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionDefinitionsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionDefinitionsResponse(val *SubscriptionDefinitionsResponse) *NullableSubscriptionDefinitionsResponse {
	return &NullableSubscriptionDefinitionsResponse{value: val, isSet: true}
}

func (v NullableSubscriptionDefinitionsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionDefinitionsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
