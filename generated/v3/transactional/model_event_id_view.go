/*
Transactional Single Send

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package transactional

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the EventIdView type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EventIdView{}

// EventIdView The ID of a send event.
type EventIdView struct {
	// Time of event creation.
	Created time.Time `json:"created"`
	// Identifier of event.
	Id string `json:"id"`
}

type _EventIdView EventIdView

// NewEventIdView instantiates a new EventIdView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEventIdView(created time.Time, id string) *EventIdView {
	this := EventIdView{}
	this.Created = created
	this.Id = id
	return &this
}

// NewEventIdViewWithDefaults instantiates a new EventIdView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEventIdViewWithDefaults() *EventIdView {
	this := EventIdView{}
	return &this
}

// GetCreated returns the Created field value
func (o *EventIdView) GetCreated() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *EventIdView) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *EventIdView) SetCreated(v time.Time) {
	o.Created = v
}

// GetId returns the Id field value
func (o *EventIdView) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *EventIdView) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *EventIdView) SetId(v string) {
	o.Id = v
}

func (o EventIdView) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EventIdView) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["created"] = o.Created
	toSerialize["id"] = o.Id
	return toSerialize, nil
}

func (o *EventIdView) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"created",
		"id",
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

	varEventIdView := _EventIdView{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varEventIdView)

	if err != nil {
		return err
	}

	*o = EventIdView(varEventIdView)

	return err
}

type NullableEventIdView struct {
	value *EventIdView
	isSet bool
}

func (v NullableEventIdView) Get() *EventIdView {
	return v.value
}

func (v *NullableEventIdView) Set(val *EventIdView) {
	v.value = val
	v.isSet = true
}

func (v NullableEventIdView) IsSet() bool {
	return v.isSet
}

func (v *NullableEventIdView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEventIdView(val *EventIdView) *NullableEventIdView {
	return &NullableEventIdView{value: val, isSet: true}
}

func (v NullableEventIdView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEventIdView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
