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

// checks if the OptionInput type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OptionInput{}

// OptionInput struct for OptionInput
type OptionInput struct {
	// Hidden options won't be shown in HubSpot.
	Hidden bool `json:"hidden"`
	// Options are shown in order starting with the lowest positive integer value. Values of -1 will cause the option to be displayed after any positive values.
	DisplayOrder *int32 `json:"displayOrder,omitempty"`
	// A description of the option.
	Description *string `json:"description,omitempty"`
	// A human-readable option label that will be shown in HubSpot.
	Label string `json:"label"`
	// The internal value of the option, which must be used when setting the property value through the API.
	Value string `json:"value"`
}

type _OptionInput OptionInput

// NewOptionInput instantiates a new OptionInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOptionInput(hidden bool, label string, value string) *OptionInput {
	this := OptionInput{}
	this.Hidden = hidden
	this.Label = label
	this.Value = value
	return &this
}

// NewOptionInputWithDefaults instantiates a new OptionInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOptionInputWithDefaults() *OptionInput {
	this := OptionInput{}
	return &this
}

// GetHidden returns the Hidden field value
func (o *OptionInput) GetHidden() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Hidden
}

// GetHiddenOk returns a tuple with the Hidden field value
// and a boolean to check if the value has been set.
func (o *OptionInput) GetHiddenOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Hidden, true
}

// SetHidden sets field value
func (o *OptionInput) SetHidden(v bool) {
	o.Hidden = v
}

// GetDisplayOrder returns the DisplayOrder field value if set, zero value otherwise.
func (o *OptionInput) GetDisplayOrder() int32 {
	if o == nil || IsNil(o.DisplayOrder) {
		var ret int32
		return ret
	}
	return *o.DisplayOrder
}

// GetDisplayOrderOk returns a tuple with the DisplayOrder field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionInput) GetDisplayOrderOk() (*int32, bool) {
	if o == nil || IsNil(o.DisplayOrder) {
		return nil, false
	}
	return o.DisplayOrder, true
}

// HasDisplayOrder returns a boolean if a field has been set.
func (o *OptionInput) HasDisplayOrder() bool {
	if o != nil && !IsNil(o.DisplayOrder) {
		return true
	}

	return false
}

// SetDisplayOrder gets a reference to the given int32 and assigns it to the DisplayOrder field.
func (o *OptionInput) SetDisplayOrder(v int32) {
	o.DisplayOrder = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *OptionInput) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OptionInput) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *OptionInput) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *OptionInput) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value
func (o *OptionInput) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *OptionInput) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *OptionInput) SetLabel(v string) {
	o.Label = v
}

// GetValue returns the Value field value
func (o *OptionInput) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *OptionInput) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *OptionInput) SetValue(v string) {
	o.Value = v
}

func (o OptionInput) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OptionInput) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["hidden"] = o.Hidden
	if !IsNil(o.DisplayOrder) {
		toSerialize["displayOrder"] = o.DisplayOrder
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["label"] = o.Label
	toSerialize["value"] = o.Value
	return toSerialize, nil
}

func (o *OptionInput) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"hidden",
		"label",
		"value",
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

	varOptionInput := _OptionInput{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varOptionInput)

	if err != nil {
		return err
	}

	*o = OptionInput(varOptionInput)

	return err
}

type NullableOptionInput struct {
	value *OptionInput
	isSet bool
}

func (v NullableOptionInput) Get() *OptionInput {
	return v.value
}

func (v *NullableOptionInput) Set(val *OptionInput) {
	v.value = val
	v.isSet = true
}

func (v NullableOptionInput) IsSet() bool {
	return v.isSet
}

func (v *NullableOptionInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOptionInput(val *OptionInput) *NullableOptionInput {
	return &NullableOptionInput{value: val, isSet: true}
}

func (v NullableOptionInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOptionInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
