/*
Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SubscriberVidResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriberVidResponse{}

// SubscriberVidResponse struct for SubscriberVidResponse
type SubscriberVidResponse struct {
	Vid int32 `json:"vid"`
}

type _SubscriberVidResponse SubscriberVidResponse

// NewSubscriberVidResponse instantiates a new SubscriberVidResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriberVidResponse(vid int32) *SubscriberVidResponse {
	this := SubscriberVidResponse{}
	this.Vid = vid
	return &this
}

// NewSubscriberVidResponseWithDefaults instantiates a new SubscriberVidResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriberVidResponseWithDefaults() *SubscriberVidResponse {
	this := SubscriberVidResponse{}
	return &this
}

// GetVid returns the Vid field value
func (o *SubscriberVidResponse) GetVid() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Vid
}

// GetVidOk returns a tuple with the Vid field value
// and a boolean to check if the value has been set.
func (o *SubscriberVidResponse) GetVidOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Vid, true
}

// SetVid sets field value
func (o *SubscriberVidResponse) SetVid(v int32) {
	o.Vid = v
}

func (o SubscriberVidResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriberVidResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["vid"] = o.Vid
	return toSerialize, nil
}

func (o *SubscriberVidResponse) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"vid",
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

	varSubscriberVidResponse := _SubscriberVidResponse{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSubscriberVidResponse)

	if err != nil {
		return err
	}

	*o = SubscriberVidResponse(varSubscriberVidResponse)

	return err
}

type NullableSubscriberVidResponse struct {
	value *SubscriberVidResponse
	isSet bool
}

func (v NullableSubscriberVidResponse) Get() *SubscriberVidResponse {
	return v.value
}

func (v *NullableSubscriberVidResponse) Set(val *SubscriberVidResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriberVidResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriberVidResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriberVidResponse(val *SubscriberVidResponse) *NullableSubscriberVidResponse {
	return &NullableSubscriberVidResponse{value: val, isSet: true}
}

func (v NullableSubscriberVidResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriberVidResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
