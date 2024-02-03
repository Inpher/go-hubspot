/*
Webhooks Webhooks

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"encoding/json"
)

// SettingsChangeRequest New or updated webhook settings for an app.
type SettingsChangeRequest struct {
	Throttling ThrottlingSettings `json:"throttling"`
	// A publicly available URL for Hubspot to call where event payloads will be delivered. See [link-so-some-doc](#) for details about the format of these event payloads.
	TargetUrl string `json:"targetUrl"`
}

// NewSettingsChangeRequest instantiates a new SettingsChangeRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSettingsChangeRequest(throttling ThrottlingSettings, targetUrl string) *SettingsChangeRequest {
	this := SettingsChangeRequest{}
	this.Throttling = throttling
	this.TargetUrl = targetUrl
	return &this
}

// NewSettingsChangeRequestWithDefaults instantiates a new SettingsChangeRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSettingsChangeRequestWithDefaults() *SettingsChangeRequest {
	this := SettingsChangeRequest{}
	return &this
}

// GetThrottling returns the Throttling field value
func (o *SettingsChangeRequest) GetThrottling() ThrottlingSettings {
	if o == nil {
		var ret ThrottlingSettings
		return ret
	}

	return o.Throttling
}

// GetThrottlingOk returns a tuple with the Throttling field value
// and a boolean to check if the value has been set.
func (o *SettingsChangeRequest) GetThrottlingOk() (*ThrottlingSettings, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Throttling, true
}

// SetThrottling sets field value
func (o *SettingsChangeRequest) SetThrottling(v ThrottlingSettings) {
	o.Throttling = v
}

// GetTargetUrl returns the TargetUrl field value
func (o *SettingsChangeRequest) GetTargetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetUrl
}

// GetTargetUrlOk returns a tuple with the TargetUrl field value
// and a boolean to check if the value has been set.
func (o *SettingsChangeRequest) GetTargetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetUrl, true
}

// SetTargetUrl sets field value
func (o *SettingsChangeRequest) SetTargetUrl(v string) {
	o.TargetUrl = v
}

func (o SettingsChangeRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["throttling"] = o.Throttling
	}
	if true {
		toSerialize["targetUrl"] = o.TargetUrl
	}
	return json.Marshal(toSerialize)
}

type NullableSettingsChangeRequest struct {
	value *SettingsChangeRequest
	isSet bool
}

func (v NullableSettingsChangeRequest) Get() *SettingsChangeRequest {
	return v.value
}

func (v *NullableSettingsChangeRequest) Set(val *SettingsChangeRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableSettingsChangeRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableSettingsChangeRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSettingsChangeRequest(val *SettingsChangeRequest) *NullableSettingsChangeRequest {
	return &NullableSettingsChangeRequest{value: val, isSet: true}
}

func (v NullableSettingsChangeRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSettingsChangeRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}