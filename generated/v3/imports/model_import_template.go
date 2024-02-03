/*
Imports

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imports

import (
	"encoding/json"
)

// ImportTemplate struct for ImportTemplate
type ImportTemplate struct {
	TemplateType string `json:"templateType"`
	TemplateId   int32  `json:"templateId"`
}

// NewImportTemplate instantiates a new ImportTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportTemplate(templateType string, templateId int32) *ImportTemplate {
	this := ImportTemplate{}
	this.TemplateType = templateType
	this.TemplateId = templateId
	return &this
}

// NewImportTemplateWithDefaults instantiates a new ImportTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportTemplateWithDefaults() *ImportTemplate {
	this := ImportTemplate{}
	return &this
}

// GetTemplateType returns the TemplateType field value
func (o *ImportTemplate) GetTemplateType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TemplateType
}

// GetTemplateTypeOk returns a tuple with the TemplateType field value
// and a boolean to check if the value has been set.
func (o *ImportTemplate) GetTemplateTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateType, true
}

// SetTemplateType sets field value
func (o *ImportTemplate) SetTemplateType(v string) {
	o.TemplateType = v
}

// GetTemplateId returns the TemplateId field value
func (o *ImportTemplate) GetTemplateId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value
// and a boolean to check if the value has been set.
func (o *ImportTemplate) GetTemplateIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TemplateId, true
}

// SetTemplateId sets field value
func (o *ImportTemplate) SetTemplateId(v int32) {
	o.TemplateId = v
}

func (o ImportTemplate) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["templateType"] = o.TemplateType
	}
	if true {
		toSerialize["templateId"] = o.TemplateId
	}
	return json.Marshal(toSerialize)
}

type NullableImportTemplate struct {
	value *ImportTemplate
	isSet bool
}

func (v NullableImportTemplate) Get() *ImportTemplate {
	return v.value
}

func (v *NullableImportTemplate) Set(val *ImportTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableImportTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableImportTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportTemplate(val *ImportTemplate) *NullableImportTemplate {
	return &NullableImportTemplate{value: val, isSet: true}
}

func (v NullableImportTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}