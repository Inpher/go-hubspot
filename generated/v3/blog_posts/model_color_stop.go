/*
Posts

Use these endpoints for interacting with Blog Posts, Blog Authors, and Blog Tags

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package blog_posts

import (
	"encoding/json"
)

// ColorStop struct for ColorStop
type ColorStop struct {
	Color RGBAColor `json:"color"`
}

// NewColorStop instantiates a new ColorStop object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColorStop(color RGBAColor) *ColorStop {
	this := ColorStop{}
	this.Color = color
	return &this
}

// NewColorStopWithDefaults instantiates a new ColorStop object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColorStopWithDefaults() *ColorStop {
	this := ColorStop{}
	return &this
}

// GetColor returns the Color field value
func (o *ColorStop) GetColor() RGBAColor {
	if o == nil {
		var ret RGBAColor
		return ret
	}

	return o.Color
}

// GetColorOk returns a tuple with the Color field value
// and a boolean to check if the value has been set.
func (o *ColorStop) GetColorOk() (*RGBAColor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Color, true
}

// SetColor sets field value
func (o *ColorStop) SetColor(v RGBAColor) {
	o.Color = v
}

func (o ColorStop) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["color"] = o.Color
	}
	return json.Marshal(toSerialize)
}

type NullableColorStop struct {
	value *ColorStop
	isSet bool
}

func (v NullableColorStop) Get() *ColorStop {
	return v.value
}

func (v *NullableColorStop) Set(val *ColorStop) {
	v.value = val
	v.isSet = true
}

func (v NullableColorStop) IsSet() bool {
	return v.isSet
}

func (v *NullableColorStop) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColorStop(val *ColorStop) *NullableColorStop {
	return &NullableColorStop{value: val, isSet: true}
}

func (v NullableColorStop) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColorStop) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
