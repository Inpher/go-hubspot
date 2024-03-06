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

// checks if the PublicExecutionTranslationRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicExecutionTranslationRule{}

// PublicExecutionTranslationRule struct for PublicExecutionTranslationRule
type PublicExecutionTranslationRule struct {
	LabelName  string                            `json:"labelName"`
	Conditions map[string]map[string]interface{} `json:"conditions"`
}

type _PublicExecutionTranslationRule PublicExecutionTranslationRule

// NewPublicExecutionTranslationRule instantiates a new PublicExecutionTranslationRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicExecutionTranslationRule(labelName string, conditions map[string]map[string]interface{}) *PublicExecutionTranslationRule {
	this := PublicExecutionTranslationRule{}
	this.LabelName = labelName
	this.Conditions = conditions
	return &this
}

// NewPublicExecutionTranslationRuleWithDefaults instantiates a new PublicExecutionTranslationRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicExecutionTranslationRuleWithDefaults() *PublicExecutionTranslationRule {
	this := PublicExecutionTranslationRule{}
	return &this
}

// GetLabelName returns the LabelName field value
func (o *PublicExecutionTranslationRule) GetLabelName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LabelName
}

// GetLabelNameOk returns a tuple with the LabelName field value
// and a boolean to check if the value has been set.
func (o *PublicExecutionTranslationRule) GetLabelNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LabelName, true
}

// SetLabelName sets field value
func (o *PublicExecutionTranslationRule) SetLabelName(v string) {
	o.LabelName = v
}

// GetConditions returns the Conditions field value
func (o *PublicExecutionTranslationRule) GetConditions() map[string]map[string]interface{} {
	if o == nil {
		var ret map[string]map[string]interface{}
		return ret
	}

	return o.Conditions
}

// GetConditionsOk returns a tuple with the Conditions field value
// and a boolean to check if the value has been set.
func (o *PublicExecutionTranslationRule) GetConditionsOk() (map[string]map[string]interface{}, bool) {
	if o == nil {
		return map[string]map[string]interface{}{}, false
	}
	return o.Conditions, true
}

// SetConditions sets field value
func (o *PublicExecutionTranslationRule) SetConditions(v map[string]map[string]interface{}) {
	o.Conditions = v
}

func (o PublicExecutionTranslationRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicExecutionTranslationRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["labelName"] = o.LabelName
	toSerialize["conditions"] = o.Conditions
	return toSerialize, nil
}

func (o *PublicExecutionTranslationRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"labelName",
		"conditions",
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

	varPublicExecutionTranslationRule := _PublicExecutionTranslationRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicExecutionTranslationRule)

	if err != nil {
		return err
	}

	*o = PublicExecutionTranslationRule(varPublicExecutionTranslationRule)

	return err
}

type NullablePublicExecutionTranslationRule struct {
	value *PublicExecutionTranslationRule
	isSet bool
}

func (v NullablePublicExecutionTranslationRule) Get() *PublicExecutionTranslationRule {
	return v.value
}

func (v *NullablePublicExecutionTranslationRule) Set(val *PublicExecutionTranslationRule) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicExecutionTranslationRule) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicExecutionTranslationRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicExecutionTranslationRule(val *PublicExecutionTranslationRule) *NullablePublicExecutionTranslationRule {
	return &NullablePublicExecutionTranslationRule{value: val, isSet: true}
}

func (v NullablePublicExecutionTranslationRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicExecutionTranslationRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
