/*
Webhooks Webhooks

Provides a way for apps to subscribe to certain change events in HubSpot. Once configured, apps will receive event payloads containing details about the changes at a specified target URL. There can only be one target URL for receiving event notifications per app.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhooks

import (
	"encoding/json"
	"time"
)

// BatchResponseSubscriptionResponse struct for BatchResponseSubscriptionResponse
type BatchResponseSubscriptionResponse struct {
	CompletedAt time.Time              `json:"completedAt"`
	RequestedAt *time.Time             `json:"requestedAt,omitempty"`
	StartedAt   time.Time              `json:"startedAt"`
	Links       *map[string]string     `json:"links,omitempty"`
	Results     []SubscriptionResponse `json:"results"`
	Status      string                 `json:"status"`
}

// NewBatchResponseSubscriptionResponse instantiates a new BatchResponseSubscriptionResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBatchResponseSubscriptionResponse(completedAt time.Time, startedAt time.Time, results []SubscriptionResponse, status string) *BatchResponseSubscriptionResponse {
	this := BatchResponseSubscriptionResponse{}
	this.CompletedAt = completedAt
	this.StartedAt = startedAt
	this.Results = results
	this.Status = status
	return &this
}

// NewBatchResponseSubscriptionResponseWithDefaults instantiates a new BatchResponseSubscriptionResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBatchResponseSubscriptionResponseWithDefaults() *BatchResponseSubscriptionResponse {
	this := BatchResponseSubscriptionResponse{}
	return &this
}

// GetCompletedAt returns the CompletedAt field value
func (o *BatchResponseSubscriptionResponse) GetCompletedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CompletedAt
}

// GetCompletedAtOk returns a tuple with the CompletedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponse) GetCompletedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CompletedAt, true
}

// SetCompletedAt sets field value
func (o *BatchResponseSubscriptionResponse) SetCompletedAt(v time.Time) {
	o.CompletedAt = v
}

// GetRequestedAt returns the RequestedAt field value if set, zero value otherwise.
func (o *BatchResponseSubscriptionResponse) GetRequestedAt() time.Time {
	if o == nil || o.RequestedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.RequestedAt
}

// GetRequestedAtOk returns a tuple with the RequestedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponse) GetRequestedAtOk() (*time.Time, bool) {
	if o == nil || o.RequestedAt == nil {
		return nil, false
	}
	return o.RequestedAt, true
}

// HasRequestedAt returns a boolean if a field has been set.
func (o *BatchResponseSubscriptionResponse) HasRequestedAt() bool {
	if o != nil && o.RequestedAt != nil {
		return true
	}

	return false
}

// SetRequestedAt gets a reference to the given time.Time and assigns it to the RequestedAt field.
func (o *BatchResponseSubscriptionResponse) SetRequestedAt(v time.Time) {
	o.RequestedAt = &v
}

// GetStartedAt returns the StartedAt field value
func (o *BatchResponseSubscriptionResponse) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponse) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *BatchResponseSubscriptionResponse) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetLinks returns the Links field value if set, zero value otherwise.
func (o *BatchResponseSubscriptionResponse) GetLinks() map[string]string {
	if o == nil || o.Links == nil {
		var ret map[string]string
		return ret
	}
	return *o.Links
}

// GetLinksOk returns a tuple with the Links field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponse) GetLinksOk() (*map[string]string, bool) {
	if o == nil || o.Links == nil {
		return nil, false
	}
	return o.Links, true
}

// HasLinks returns a boolean if a field has been set.
func (o *BatchResponseSubscriptionResponse) HasLinks() bool {
	if o != nil && o.Links != nil {
		return true
	}

	return false
}

// SetLinks gets a reference to the given map[string]string and assigns it to the Links field.
func (o *BatchResponseSubscriptionResponse) SetLinks(v map[string]string) {
	o.Links = &v
}

// GetResults returns the Results field value
func (o *BatchResponseSubscriptionResponse) GetResults() []SubscriptionResponse {
	if o == nil {
		var ret []SubscriptionResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponse) GetResultsOk() ([]SubscriptionResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *BatchResponseSubscriptionResponse) SetResults(v []SubscriptionResponse) {
	o.Results = v
}

// GetStatus returns the Status field value
func (o *BatchResponseSubscriptionResponse) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *BatchResponseSubscriptionResponse) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *BatchResponseSubscriptionResponse) SetStatus(v string) {
	o.Status = v
}

func (o BatchResponseSubscriptionResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["completedAt"] = o.CompletedAt
	}
	if o.RequestedAt != nil {
		toSerialize["requestedAt"] = o.RequestedAt
	}
	if true {
		toSerialize["startedAt"] = o.StartedAt
	}
	if o.Links != nil {
		toSerialize["links"] = o.Links
	}
	if true {
		toSerialize["results"] = o.Results
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableBatchResponseSubscriptionResponse struct {
	value *BatchResponseSubscriptionResponse
	isSet bool
}

func (v NullableBatchResponseSubscriptionResponse) Get() *BatchResponseSubscriptionResponse {
	return v.value
}

func (v *NullableBatchResponseSubscriptionResponse) Set(val *BatchResponseSubscriptionResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableBatchResponseSubscriptionResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableBatchResponseSubscriptionResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBatchResponseSubscriptionResponse(val *BatchResponseSubscriptionResponse) *NullableBatchResponseSubscriptionResponse {
	return &NullableBatchResponseSubscriptionResponse{value: val, isSet: true}
}

func (v NullableBatchResponseSubscriptionResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBatchResponseSubscriptionResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
