/*
Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
)

// PropertyValue struct for PropertyValue
type PropertyValue struct {
	//
	SourceId string `json:"sourceId"`
	//
	SelectedByUser bool `json:"selectedByUser"`
	//
	SourceLabel string `json:"sourceLabel"`
	//
	Source               string `json:"source"`
	UpdatedByUserId      *int32 `json:"updatedByUserId,omitempty"`
	PersistenceTimestamp *int64 `json:"persistenceTimestamp,omitempty"`
	// Source metadata encoded as a base64 string. For example: `ZXhhbXBsZSBzdHJpbmc=`
	SourceMetadata string `json:"sourceMetadata"`
	//
	SourceVid []int64 `json:"sourceVid"`
	//
	RequestId string `json:"requestId"`
	//
	Name                               string `json:"name"`
	UseTimestampAsPersistenceTimestamp *bool  `json:"useTimestampAsPersistenceTimestamp,omitempty"`
	//
	Value string `json:"value"`
	//
	SelectedByUserTimestamp int64 `json:"selectedByUserTimestamp"`
	//
	Timestamp    int64 `json:"timestamp"`
	IsLargeValue *bool `json:"isLargeValue,omitempty"`
}

// NewPropertyValue instantiates a new PropertyValue object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPropertyValue(sourceId string, selectedByUser bool, sourceLabel string, source string, sourceMetadata string, sourceVid []int64, requestId string, name string, value string, selectedByUserTimestamp int64, timestamp int64) *PropertyValue {
	this := PropertyValue{}
	this.SourceId = sourceId
	this.SelectedByUser = selectedByUser
	this.SourceLabel = sourceLabel
	this.Source = source
	this.SourceMetadata = sourceMetadata
	this.SourceVid = sourceVid
	this.RequestId = requestId
	this.Name = name
	this.Value = value
	this.SelectedByUserTimestamp = selectedByUserTimestamp
	this.Timestamp = timestamp
	return &this
}

// NewPropertyValueWithDefaults instantiates a new PropertyValue object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPropertyValueWithDefaults() *PropertyValue {
	this := PropertyValue{}
	return &this
}

// GetSourceId returns the SourceId field value
func (o *PropertyValue) GetSourceId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceId
}

// GetSourceIdOk returns a tuple with the SourceId field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetSourceIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceId, true
}

// SetSourceId sets field value
func (o *PropertyValue) SetSourceId(v string) {
	o.SourceId = v
}

// GetSelectedByUser returns the SelectedByUser field value
func (o *PropertyValue) GetSelectedByUser() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.SelectedByUser
}

// GetSelectedByUserOk returns a tuple with the SelectedByUser field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetSelectedByUserOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedByUser, true
}

// SetSelectedByUser sets field value
func (o *PropertyValue) SetSelectedByUser(v bool) {
	o.SelectedByUser = v
}

// GetSourceLabel returns the SourceLabel field value
func (o *PropertyValue) GetSourceLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceLabel
}

// GetSourceLabelOk returns a tuple with the SourceLabel field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetSourceLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceLabel, true
}

// SetSourceLabel sets field value
func (o *PropertyValue) SetSourceLabel(v string) {
	o.SourceLabel = v
}

// GetSource returns the Source field value
func (o *PropertyValue) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PropertyValue) SetSource(v string) {
	o.Source = v
}

// GetUpdatedByUserId returns the UpdatedByUserId field value if set, zero value otherwise.
func (o *PropertyValue) GetUpdatedByUserId() int32 {
	if o == nil || o.UpdatedByUserId == nil {
		var ret int32
		return ret
	}
	return *o.UpdatedByUserId
}

// GetUpdatedByUserIdOk returns a tuple with the UpdatedByUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetUpdatedByUserIdOk() (*int32, bool) {
	if o == nil || o.UpdatedByUserId == nil {
		return nil, false
	}
	return o.UpdatedByUserId, true
}

// HasUpdatedByUserId returns a boolean if a field has been set.
func (o *PropertyValue) HasUpdatedByUserId() bool {
	if o != nil && o.UpdatedByUserId != nil {
		return true
	}

	return false
}

// SetUpdatedByUserId gets a reference to the given int32 and assigns it to the UpdatedByUserId field.
func (o *PropertyValue) SetUpdatedByUserId(v int32) {
	o.UpdatedByUserId = &v
}

// GetPersistenceTimestamp returns the PersistenceTimestamp field value if set, zero value otherwise.
func (o *PropertyValue) GetPersistenceTimestamp() int64 {
	if o == nil || o.PersistenceTimestamp == nil {
		var ret int64
		return ret
	}
	return *o.PersistenceTimestamp
}

// GetPersistenceTimestampOk returns a tuple with the PersistenceTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetPersistenceTimestampOk() (*int64, bool) {
	if o == nil || o.PersistenceTimestamp == nil {
		return nil, false
	}
	return o.PersistenceTimestamp, true
}

// HasPersistenceTimestamp returns a boolean if a field has been set.
func (o *PropertyValue) HasPersistenceTimestamp() bool {
	if o != nil && o.PersistenceTimestamp != nil {
		return true
	}

	return false
}

// SetPersistenceTimestamp gets a reference to the given int64 and assigns it to the PersistenceTimestamp field.
func (o *PropertyValue) SetPersistenceTimestamp(v int64) {
	o.PersistenceTimestamp = &v
}

// GetSourceMetadata returns the SourceMetadata field value
func (o *PropertyValue) GetSourceMetadata() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.SourceMetadata
}

// GetSourceMetadataOk returns a tuple with the SourceMetadata field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetSourceMetadataOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SourceMetadata, true
}

// SetSourceMetadata sets field value
func (o *PropertyValue) SetSourceMetadata(v string) {
	o.SourceMetadata = v
}

// GetSourceVid returns the SourceVid field value
func (o *PropertyValue) GetSourceVid() []int64 {
	if o == nil {
		var ret []int64
		return ret
	}

	return o.SourceVid
}

// GetSourceVidOk returns a tuple with the SourceVid field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetSourceVidOk() ([]int64, bool) {
	if o == nil {
		return nil, false
	}
	return o.SourceVid, true
}

// SetSourceVid sets field value
func (o *PropertyValue) SetSourceVid(v []int64) {
	o.SourceVid = v
}

// GetRequestId returns the RequestId field value
func (o *PropertyValue) GetRequestId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.RequestId
}

// GetRequestIdOk returns a tuple with the RequestId field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.RequestId, true
}

// SetRequestId sets field value
func (o *PropertyValue) SetRequestId(v string) {
	o.RequestId = v
}

// GetName returns the Name field value
func (o *PropertyValue) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PropertyValue) SetName(v string) {
	o.Name = v
}

// GetUseTimestampAsPersistenceTimestamp returns the UseTimestampAsPersistenceTimestamp field value if set, zero value otherwise.
func (o *PropertyValue) GetUseTimestampAsPersistenceTimestamp() bool {
	if o == nil || o.UseTimestampAsPersistenceTimestamp == nil {
		var ret bool
		return ret
	}
	return *o.UseTimestampAsPersistenceTimestamp
}

// GetUseTimestampAsPersistenceTimestampOk returns a tuple with the UseTimestampAsPersistenceTimestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetUseTimestampAsPersistenceTimestampOk() (*bool, bool) {
	if o == nil || o.UseTimestampAsPersistenceTimestamp == nil {
		return nil, false
	}
	return o.UseTimestampAsPersistenceTimestamp, true
}

// HasUseTimestampAsPersistenceTimestamp returns a boolean if a field has been set.
func (o *PropertyValue) HasUseTimestampAsPersistenceTimestamp() bool {
	if o != nil && o.UseTimestampAsPersistenceTimestamp != nil {
		return true
	}

	return false
}

// SetUseTimestampAsPersistenceTimestamp gets a reference to the given bool and assigns it to the UseTimestampAsPersistenceTimestamp field.
func (o *PropertyValue) SetUseTimestampAsPersistenceTimestamp(v bool) {
	o.UseTimestampAsPersistenceTimestamp = &v
}

// GetValue returns the Value field value
func (o *PropertyValue) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *PropertyValue) SetValue(v string) {
	o.Value = v
}

// GetSelectedByUserTimestamp returns the SelectedByUserTimestamp field value
func (o *PropertyValue) GetSelectedByUserTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.SelectedByUserTimestamp
}

// GetSelectedByUserTimestampOk returns a tuple with the SelectedByUserTimestamp field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetSelectedByUserTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.SelectedByUserTimestamp, true
}

// SetSelectedByUserTimestamp sets field value
func (o *PropertyValue) SetSelectedByUserTimestamp(v int64) {
	o.SelectedByUserTimestamp = v
}

// GetTimestamp returns the Timestamp field value
func (o *PropertyValue) GetTimestamp() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetTimestampOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Timestamp, true
}

// SetTimestamp sets field value
func (o *PropertyValue) SetTimestamp(v int64) {
	o.Timestamp = v
}

// GetIsLargeValue returns the IsLargeValue field value if set, zero value otherwise.
func (o *PropertyValue) GetIsLargeValue() bool {
	if o == nil || o.IsLargeValue == nil {
		var ret bool
		return ret
	}
	return *o.IsLargeValue
}

// GetIsLargeValueOk returns a tuple with the IsLargeValue field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PropertyValue) GetIsLargeValueOk() (*bool, bool) {
	if o == nil || o.IsLargeValue == nil {
		return nil, false
	}
	return o.IsLargeValue, true
}

// HasIsLargeValue returns a boolean if a field has been set.
func (o *PropertyValue) HasIsLargeValue() bool {
	if o != nil && o.IsLargeValue != nil {
		return true
	}

	return false
}

// SetIsLargeValue gets a reference to the given bool and assigns it to the IsLargeValue field.
func (o *PropertyValue) SetIsLargeValue(v bool) {
	o.IsLargeValue = &v
}

func (o PropertyValue) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["sourceId"] = o.SourceId
	}
	if true {
		toSerialize["selectedByUser"] = o.SelectedByUser
	}
	if true {
		toSerialize["sourceLabel"] = o.SourceLabel
	}
	if true {
		toSerialize["source"] = o.Source
	}
	if o.UpdatedByUserId != nil {
		toSerialize["updatedByUserId"] = o.UpdatedByUserId
	}
	if o.PersistenceTimestamp != nil {
		toSerialize["persistenceTimestamp"] = o.PersistenceTimestamp
	}
	if true {
		toSerialize["sourceMetadata"] = o.SourceMetadata
	}
	if true {
		toSerialize["sourceVid"] = o.SourceVid
	}
	if true {
		toSerialize["requestId"] = o.RequestId
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.UseTimestampAsPersistenceTimestamp != nil {
		toSerialize["useTimestampAsPersistenceTimestamp"] = o.UseTimestampAsPersistenceTimestamp
	}
	if true {
		toSerialize["value"] = o.Value
	}
	if true {
		toSerialize["selectedByUserTimestamp"] = o.SelectedByUserTimestamp
	}
	if true {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.IsLargeValue != nil {
		toSerialize["isLargeValue"] = o.IsLargeValue
	}
	return json.Marshal(toSerialize)
}

type NullablePropertyValue struct {
	value *PropertyValue
	isSet bool
}

func (v NullablePropertyValue) Get() *PropertyValue {
	return v.value
}

func (v *NullablePropertyValue) Set(val *PropertyValue) {
	v.value = val
	v.isSet = true
}

func (v NullablePropertyValue) IsSet() bool {
	return v.isSet
}

func (v *NullablePropertyValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePropertyValue(val *PropertyValue) *NullablePropertyValue {
	return &NullablePropertyValue{value: val, isSet: true}
}

func (v NullablePropertyValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePropertyValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}