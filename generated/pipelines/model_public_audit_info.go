/*
CRM Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"encoding/json"
	"time"
)

// PublicAuditInfo struct for PublicAuditInfo
type PublicAuditInfo struct {
	PortalId   int32                  `json:"portalId"`
	Identifier string                 `json:"identifier"`
	Action     string                 `json:"action"`
	Timestamp  *time.Time             `json:"timestamp,omitempty"`
	Message    *string                `json:"message,omitempty"`
	RawObject  map[string]interface{} `json:"rawObject,omitempty"`
	FromUserId *int32                 `json:"fromUserId,omitempty"`
}

// NewPublicAuditInfo instantiates a new PublicAuditInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAuditInfo(portalId int32, identifier string, action string) *PublicAuditInfo {
	this := PublicAuditInfo{}
	this.PortalId = portalId
	this.Identifier = identifier
	this.Action = action
	return &this
}

// NewPublicAuditInfoWithDefaults instantiates a new PublicAuditInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAuditInfoWithDefaults() *PublicAuditInfo {
	this := PublicAuditInfo{}
	return &this
}

// GetPortalId returns the PortalId field value
func (o *PublicAuditInfo) GetPortalId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.PortalId
}

// GetPortalIdOk returns a tuple with the PortalId field value
// and a boolean to check if the value has been set.
func (o *PublicAuditInfo) GetPortalIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PortalId, true
}

// SetPortalId sets field value
func (o *PublicAuditInfo) SetPortalId(v int32) {
	o.PortalId = v
}

// GetIdentifier returns the Identifier field value
func (o *PublicAuditInfo) GetIdentifier() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Identifier
}

// GetIdentifierOk returns a tuple with the Identifier field value
// and a boolean to check if the value has been set.
func (o *PublicAuditInfo) GetIdentifierOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identifier, true
}

// SetIdentifier sets field value
func (o *PublicAuditInfo) SetIdentifier(v string) {
	o.Identifier = v
}

// GetAction returns the Action field value
func (o *PublicAuditInfo) GetAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Action
}

// GetActionOk returns a tuple with the Action field value
// and a boolean to check if the value has been set.
func (o *PublicAuditInfo) GetActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Action, true
}

// SetAction sets field value
func (o *PublicAuditInfo) SetAction(v string) {
	o.Action = v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *PublicAuditInfo) GetTimestamp() time.Time {
	if o == nil || o.Timestamp == nil {
		var ret time.Time
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAuditInfo) GetTimestampOk() (*time.Time, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *PublicAuditInfo) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given time.Time and assigns it to the Timestamp field.
func (o *PublicAuditInfo) SetTimestamp(v time.Time) {
	o.Timestamp = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *PublicAuditInfo) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAuditInfo) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *PublicAuditInfo) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *PublicAuditInfo) SetMessage(v string) {
	o.Message = &v
}

// GetRawObject returns the RawObject field value if set, zero value otherwise.
func (o *PublicAuditInfo) GetRawObject() map[string]interface{} {
	if o == nil || o.RawObject == nil {
		var ret map[string]interface{}
		return ret
	}
	return o.RawObject
}

// GetRawObjectOk returns a tuple with the RawObject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAuditInfo) GetRawObjectOk() (map[string]interface{}, bool) {
	if o == nil || o.RawObject == nil {
		return nil, false
	}
	return o.RawObject, true
}

// HasRawObject returns a boolean if a field has been set.
func (o *PublicAuditInfo) HasRawObject() bool {
	if o != nil && o.RawObject != nil {
		return true
	}

	return false
}

// SetRawObject gets a reference to the given map[string]interface{} and assigns it to the RawObject field.
func (o *PublicAuditInfo) SetRawObject(v map[string]interface{}) {
	o.RawObject = v
}

// GetFromUserId returns the FromUserId field value if set, zero value otherwise.
func (o *PublicAuditInfo) GetFromUserId() int32 {
	if o == nil || o.FromUserId == nil {
		var ret int32
		return ret
	}
	return *o.FromUserId
}

// GetFromUserIdOk returns a tuple with the FromUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicAuditInfo) GetFromUserIdOk() (*int32, bool) {
	if o == nil || o.FromUserId == nil {
		return nil, false
	}
	return o.FromUserId, true
}

// HasFromUserId returns a boolean if a field has been set.
func (o *PublicAuditInfo) HasFromUserId() bool {
	if o != nil && o.FromUserId != nil {
		return true
	}

	return false
}

// SetFromUserId gets a reference to the given int32 and assigns it to the FromUserId field.
func (o *PublicAuditInfo) SetFromUserId(v int32) {
	o.FromUserId = &v
}

func (o PublicAuditInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["portalId"] = o.PortalId
	}
	if true {
		toSerialize["identifier"] = o.Identifier
	}
	if true {
		toSerialize["action"] = o.Action
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.RawObject != nil {
		toSerialize["rawObject"] = o.RawObject
	}
	if o.FromUserId != nil {
		toSerialize["fromUserId"] = o.FromUserId
	}
	return json.Marshal(toSerialize)
}

type NullablePublicAuditInfo struct {
	value *PublicAuditInfo
	isSet bool
}

func (v NullablePublicAuditInfo) Get() *PublicAuditInfo {
	return v.value
}

func (v *NullablePublicAuditInfo) Set(val *PublicAuditInfo) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAuditInfo) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAuditInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAuditInfo(val *PublicAuditInfo) *NullablePublicAuditInfo {
	return &NullablePublicAuditInfo{value: val, isSet: true}
}

func (v NullablePublicAuditInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAuditInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
