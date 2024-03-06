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

// checks if the PreviousPage type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PreviousPage{}

// PreviousPage struct for PreviousPage
type PreviousPage struct {
	Before string  `json:"before"`
	Link   *string `json:"link,omitempty"`
}

type _PreviousPage PreviousPage

// NewPreviousPage instantiates a new PreviousPage object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviousPage(before string) *PreviousPage {
	this := PreviousPage{}
	this.Before = before
	return &this
}

// NewPreviousPageWithDefaults instantiates a new PreviousPage object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviousPageWithDefaults() *PreviousPage {
	this := PreviousPage{}
	return &this
}

// GetBefore returns the Before field value
func (o *PreviousPage) GetBefore() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Before
}

// GetBeforeOk returns a tuple with the Before field value
// and a boolean to check if the value has been set.
func (o *PreviousPage) GetBeforeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Before, true
}

// SetBefore sets field value
func (o *PreviousPage) SetBefore(v string) {
	o.Before = v
}

// GetLink returns the Link field value if set, zero value otherwise.
func (o *PreviousPage) GetLink() string {
	if o == nil || IsNil(o.Link) {
		var ret string
		return ret
	}
	return *o.Link
}

// GetLinkOk returns a tuple with the Link field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviousPage) GetLinkOk() (*string, bool) {
	if o == nil || IsNil(o.Link) {
		return nil, false
	}
	return o.Link, true
}

// HasLink returns a boolean if a field has been set.
func (o *PreviousPage) HasLink() bool {
	if o != nil && !IsNil(o.Link) {
		return true
	}

	return false
}

// SetLink gets a reference to the given string and assigns it to the Link field.
func (o *PreviousPage) SetLink(v string) {
	o.Link = &v
}

func (o PreviousPage) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PreviousPage) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["before"] = o.Before
	if !IsNil(o.Link) {
		toSerialize["link"] = o.Link
	}
	return toSerialize, nil
}

func (o *PreviousPage) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"before",
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

	varPreviousPage := _PreviousPage{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPreviousPage)

	if err != nil {
		return err
	}

	*o = PreviousPage(varPreviousPage)

	return err
}

type NullablePreviousPage struct {
	value *PreviousPage
	isSet bool
}

func (v NullablePreviousPage) Get() *PreviousPage {
	return v.value
}

func (v *NullablePreviousPage) Set(val *PreviousPage) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviousPage) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviousPage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviousPage(val *PreviousPage) *NullablePreviousPage {
	return &NullablePreviousPage{value: val, isSet: true}
}

func (v NullablePreviousPage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviousPage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
