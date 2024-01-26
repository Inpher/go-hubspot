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

// Option struct for Option
type Option struct {
	Hidden       bool    `json:"hidden"`
	DisplayOrder int32   `json:"displayOrder"`
	DoubleData   float32 `json:"doubleData"`
	Description  string  `json:"description"`
	ReadOnly     bool    `json:"readOnly"`
	Label        string  `json:"label"`
	Value        string  `json:"value"`
}

// NewOption instantiates a new Option object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOption(hidden bool, displayOrder int32, doubleData float32, description string, readOnly bool, label string, value string) *Option {
	this := Option{}
	this.Hidden = hidden
	this.DisplayOrder = displayOrder
	this.DoubleData = doubleData
	this.Description = description
	this.ReadOnly = readOnly
	this.Label = label
	this.Value = value
	return &this
}

// NewOptionWithDefaults instantiates a new Option object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionWithDefaults() *Option {
	this := Option{}
	return &this
}

// GetHidden returns the Hidden field value
func (o *Option) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *Option) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *Option) SetHidden(v bool) {
	o.Hidden = v
}

// GetDisplayOrder returns the DisplayOrder field value
func (o *Option) GetDisplayOrder() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value
// and a boolean to check if the value has been set.
func (o *Option) GetDisplayOrderOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DisplayOrder, true
}

// SetDisplayOrder sets field value
func (o *Option) SetDisplayOrder(v int32) {
	o.DisplayOrder = v
}

// GetDoubleData returns the DoubleData field value
func (o *Option) GetDoubleData() float32 {
	if o == nil {
		var ret float32
		return ret
	}

	return o.DoubleData
}

// GetDoubleDataOk returns a tuple with the DoubleData field value
// and a boolean to check if the value has been set.
func (o *Option) GetDoubleDataOk() (*float32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DoubleData, true
}

// SetDoubleData sets field value
func (o *Option) SetDoubleData(v float32) {
	o.DoubleData = v
}

// GetDescription returns the Description field value
func (o *Option) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *Option) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *Option) SetDescription(v string) {
	o.Description = v
}

// GetReadOnly returns the ReadOnly field value
func (o *Option) GetReadOnly() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.ReadOnly
}

// GetReadOnlyOk returns a tuple with the ReadOnly field value
// and a boolean to check if the value has been set.
func (o *Option) GetReadOnlyOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReadOnly, true
}

// SetReadOnly sets field value
func (o *Option) SetReadOnly(v bool) {
	o.ReadOnly = v
}

// GetLabel returns the Label field value
func (o *Option) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Option) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Option) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *Option) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *Option) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *Option) SetValue(v string) {
	o.Value = v
}

func (o Option) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["hidden"] = o.Hidden
	}
	if true {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if true {
		toSerialize["doubleData"] = o.DoubleData
	}
	if true {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["readOnly"] = o.ReadOnly
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableOption struct {
	value *Option
	isSet bool
}

func (v NullableOption) Get() *Option {
	return v.value
}

func (v *NullableOption) Set(val *Option) {
	v.value = val
	v.isSet = true
}

func (v NullableOption) IsSet() bool {
	return v.isSet
}

func (v *NullableOption) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOption(val *Option) *NullableOption {
	return &NullableOption{value: val, isSet: true}
}

func (v NullableOption) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOption) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
