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

// InputFieldDefinition struct for InputFieldDefinition
type InputFieldDefinition struct {
	IsRequired          bool                `json:"isRequired"`
	AutomationFieldType *string             `json:"automationFieldType,omitempty"`
	TypeDefinition      FieldTypeDefinition `json:"typeDefinition"`
	SupportedValueTypes []string            `json:"supportedValueTypes,omitempty"`
}

// NewInputFieldDefinition instantiates a new InputFieldDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInputFieldDefinition(isRequired bool, typeDefinition FieldTypeDefinition) *InputFieldDefinition {
	this := InputFieldDefinition{}
	this.IsRequired = isRequired
	this.TypeDefinition = typeDefinition
	return &this
}

// NewInputFieldDefinitionWithDefaults instantiates a new InputFieldDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInputFieldDefinitionWithDefaults() *InputFieldDefinition {
	this := InputFieldDefinition{}
	return &this
}

// GetIsRequired returns the IsRequired field value
func (o *InputFieldDefinition) GetIsRequired() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.IsRequired
}

// GetIsRequiredOk returns a tuple with the IsRequired field value
// and a boolean to check if the value has been set.
func (o *InputFieldDefinition) GetIsRequiredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.IsRequired, true
}

// SetIsRequired sets field value
func (o *InputFieldDefinition) SetIsRequired(v bool) {
	o.IsRequired = v
}

// GetAutomationFieldType returns the AutomationFieldType field value if set, zero value otherwise.
func (o *InputFieldDefinition) GetAutomationFieldType() string {
	if o == nil || o.AutomationFieldType == nil {
		var ret string
		return ret
	}
	return *o.AutomationFieldType
}

// GetAutomationFieldTypeOk returns a tuple with the AutomationFieldType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputFieldDefinition) GetAutomationFieldTypeOk() (*string, bool) {
	if o == nil || o.AutomationFieldType == nil {
		return nil, false
	}
	return o.AutomationFieldType, true
}

// HasAutomationFieldType returns a boolean if a field has been set.
func (o *InputFieldDefinition) HasAutomationFieldType() bool {
	if o != nil && o.AutomationFieldType != nil {
		return true
	}

	return false
}

// SetAutomationFieldType gets a reference to the given string and assigns it to the AutomationFieldType field.
func (o *InputFieldDefinition) SetAutomationFieldType(v string) {
	o.AutomationFieldType = &v
}

// GetTypeDefinition returns the TypeDefinition field value
func (o *InputFieldDefinition) GetTypeDefinition() FieldTypeDefinition {
	if o == nil {
		var ret FieldTypeDefinition
		return ret
	}

	return o.TypeDefinition
}

// GetTypeDefinitionOk returns a tuple with the TypeDefinition field value
// and a boolean to check if the value has been set.
func (o *InputFieldDefinition) GetTypeDefinitionOk() (*FieldTypeDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TypeDefinition, true
}

// SetTypeDefinition sets field value
func (o *InputFieldDefinition) SetTypeDefinition(v FieldTypeDefinition) {
	o.TypeDefinition = v
}

// GetSupportedValueTypes returns the SupportedValueTypes field value if set, zero value otherwise.
func (o *InputFieldDefinition) GetSupportedValueTypes() []string {
	if o == nil || o.SupportedValueTypes == nil {
		var ret []string
		return ret
	}
	return o.SupportedValueTypes
}

// GetSupportedValueTypesOk returns a tuple with the SupportedValueTypes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InputFieldDefinition) GetSupportedValueTypesOk() ([]string, bool) {
	if o == nil || o.SupportedValueTypes == nil {
		return nil, false
	}
	return o.SupportedValueTypes, true
}

// HasSupportedValueTypes returns a boolean if a field has been set.
func (o *InputFieldDefinition) HasSupportedValueTypes() bool {
	if o != nil && o.SupportedValueTypes != nil {
		return true
	}

	return false
}

// SetSupportedValueTypes gets a reference to the given []string and assigns it to the SupportedValueTypes field.
func (o *InputFieldDefinition) SetSupportedValueTypes(v []string) {
	o.SupportedValueTypes = v
}

func (o InputFieldDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["isRequired"] = o.IsRequired
	}
	if o.AutomationFieldType != nil {
		toSerialize["automationFieldType"] = o.AutomationFieldType
	}
	if true {
		toSerialize["typeDefinition"] = o.TypeDefinition
	}
	if o.SupportedValueTypes != nil {
		toSerialize["supportedValueTypes"] = o.SupportedValueTypes
	}
	return json.Marshal(toSerialize)
}

type NullableInputFieldDefinition struct {
	value *InputFieldDefinition
	isSet bool
}

func (v NullableInputFieldDefinition) Get() *InputFieldDefinition {
	return v.value
}

func (v *NullableInputFieldDefinition) Set(val *InputFieldDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableInputFieldDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableInputFieldDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInputFieldDefinition(val *InputFieldDefinition) *NullableInputFieldDefinition {
	return &NullableInputFieldDefinition{value: val, isSet: true}
}

func (v NullableInputFieldDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInputFieldDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
