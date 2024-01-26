/*
Marketing Events

These APIs allow you to interact with HubSpot's Marketing Events Extension. It allows you to: * Create, Read or update Marketing Event information in HubSpot * Specify whether a HubSpot contact has registered, attended or cancelled a registration to a Marketing Event. * Specify a URL that can be called to get the details of a Marketing Event.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package marketing_events_beta

import (
	"encoding/json"
	"time"
)

// MarketingEventCreateRequestParams struct for MarketingEventCreateRequestParams
type MarketingEventCreateRequestParams struct {
	// The start date and time of the marketing event.
	StartDateTime *time.Time `json:"startDateTime,omitempty"`
	// A list of PropertyValues. These can be whatever kind of property names and values you want. However, they must already exist on the HubSpot account's definition of the MarketingEvent Object. If they don't they will be filtered out and not set. In order to do this you'll need to create a new PropertyGroup on the HubSpot account's MarketingEvent object for your specific app and create the Custom Property you want to track on that HubSpot account. Do not create any new default properties on the MarketingEvent object as that will apply to all HubSpot accounts.
	CustomProperties []PropertyValue `json:"customProperties,omitempty"`
	// The accountId that is associated with this marketing event in the external event application.
	ExternalAccountId string `json:"externalAccountId"`
	// Indicates if the marketing event has been cancelled.  Defaults to `false`
	EventCancelled *bool `json:"eventCancelled,omitempty"`
	// The name of the organizer of the marketing event.
	EventOrganizer string `json:"eventOrganizer"`
	// A URL in the external event application where the marketing event can be managed.
	EventUrl *string `json:"eventUrl,omitempty"`
	// The id of the marketing event in the external event application.
	ExternalEventId string `json:"externalEventId"`
	// The description of the marketing event.
	EventDescription *string `json:"eventDescription,omitempty"`
	// The name of the marketing event.
	EventName string `json:"eventName"`
	// Describes what type of event this is.  For example: `WEBINAR`, `CONFERENCE`, `WORKSHOP`
	EventType *string `json:"eventType,omitempty"`
	// The end date and time of the marketing event.
	EndDateTime *time.Time `json:"endDateTime,omitempty"`
}

// NewMarketingEventCreateRequestParams instantiates a new MarketingEventCreateRequestParams object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMarketingEventCreateRequestParams(externalAccountId string, eventOrganizer string, externalEventId string, eventName string) *MarketingEventCreateRequestParams {
	this := MarketingEventCreateRequestParams{}
	this.ExternalAccountId = externalAccountId
	this.EventOrganizer = eventOrganizer
	this.ExternalEventId = externalEventId
	this.EventName = eventName
	return &this
}

// NewMarketingEventCreateRequestParamsWithDefaults instantiates a new MarketingEventCreateRequestParams object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMarketingEventCreateRequestParamsWithDefaults() *MarketingEventCreateRequestParams {
	this := MarketingEventCreateRequestParams{}
	return &this
}

// GetStartDateTime returns the StartDateTime field value if set, zero value otherwise.
func (o *MarketingEventCreateRequestParams) GetStartDateTime() time.Time {
	if o == nil || o.StartDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartDateTime
}

// GetStartDateTimeOk returns a tuple with the StartDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetStartDateTimeOk() (*time.Time, bool) {
	if o == nil || o.StartDateTime == nil {
		return nil, false
	}
	return o.StartDateTime, true
}

// HasStartDateTime returns a boolean if a field has been set.
func (o *MarketingEventCreateRequestParams) HasStartDateTime() bool {
	if o != nil && o.StartDateTime != nil {
		return true
	}

	return false
}

// SetStartDateTime gets a reference to the given time.Time and assigns it to the StartDateTime field.
func (o *MarketingEventCreateRequestParams) SetStartDateTime(v time.Time) {
	o.StartDateTime = &v
}

// GetCustomProperties returns the CustomProperties field value if set, zero value otherwise.
func (o *MarketingEventCreateRequestParams) GetCustomProperties() []PropertyValue {
	if o == nil || o.CustomProperties == nil {
		var ret []PropertyValue
		return ret
	}
	return o.CustomProperties
}

// GetCustomPropertiesOk returns a tuple with the CustomProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetCustomPropertiesOk() ([]PropertyValue, bool) {
	if o == nil || o.CustomProperties == nil {
		return nil, false
	}
	return o.CustomProperties, true
}

// HasCustomProperties returns a boolean if a field has been set.
func (o *MarketingEventCreateRequestParams) HasCustomProperties() bool {
	if o != nil && o.CustomProperties != nil {
		return true
	}

	return false
}

// SetCustomProperties gets a reference to the given []PropertyValue and assigns it to the CustomProperties field.
func (o *MarketingEventCreateRequestParams) SetCustomProperties(v []PropertyValue) {
	o.CustomProperties = v
}

// GetExternalAccountId returns the ExternalAccountId field value
func (o *MarketingEventCreateRequestParams) GetExternalAccountId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalAccountId
}

// GetExternalAccountIdOk returns a tuple with the ExternalAccountId field value
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetExternalAccountIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalAccountId, true
}

// SetExternalAccountId sets field value
func (o *MarketingEventCreateRequestParams) SetExternalAccountId(v string) {
	o.ExternalAccountId = v
}

// GetEventCancelled returns the EventCancelled field value if set, zero value otherwise.
func (o *MarketingEventCreateRequestParams) GetEventCancelled() bool {
	if o == nil || o.EventCancelled == nil {
		var ret bool
		return ret
	}
	return *o.EventCancelled
}

// GetEventCancelledOk returns a tuple with the EventCancelled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetEventCancelledOk() (*bool, bool) {
	if o == nil || o.EventCancelled == nil {
		return nil, false
	}
	return o.EventCancelled, true
}

// HasEventCancelled returns a boolean if a field has been set.
func (o *MarketingEventCreateRequestParams) HasEventCancelled() bool {
	if o != nil && o.EventCancelled != nil {
		return true
	}

	return false
}

// SetEventCancelled gets a reference to the given bool and assigns it to the EventCancelled field.
func (o *MarketingEventCreateRequestParams) SetEventCancelled(v bool) {
	o.EventCancelled = &v
}

// GetEventOrganizer returns the EventOrganizer field value
func (o *MarketingEventCreateRequestParams) GetEventOrganizer() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventOrganizer
}

// GetEventOrganizerOk returns a tuple with the EventOrganizer field value
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetEventOrganizerOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventOrganizer, true
}

// SetEventOrganizer sets field value
func (o *MarketingEventCreateRequestParams) SetEventOrganizer(v string) {
	o.EventOrganizer = v
}

// GetEventUrl returns the EventUrl field value if set, zero value otherwise.
func (o *MarketingEventCreateRequestParams) GetEventUrl() string {
	if o == nil || o.EventUrl == nil {
		var ret string
		return ret
	}
	return *o.EventUrl
}

// GetEventUrlOk returns a tuple with the EventUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetEventUrlOk() (*string, bool) {
	if o == nil || o.EventUrl == nil {
		return nil, false
	}
	return o.EventUrl, true
}

// HasEventUrl returns a boolean if a field has been set.
func (o *MarketingEventCreateRequestParams) HasEventUrl() bool {
	if o != nil && o.EventUrl != nil {
		return true
	}

	return false
}

// SetEventUrl gets a reference to the given string and assigns it to the EventUrl field.
func (o *MarketingEventCreateRequestParams) SetEventUrl(v string) {
	o.EventUrl = &v
}

// GetExternalEventId returns the ExternalEventId field value
func (o *MarketingEventCreateRequestParams) GetExternalEventId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalEventId
}

// GetExternalEventIdOk returns a tuple with the ExternalEventId field value
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetExternalEventIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalEventId, true
}

// SetExternalEventId sets field value
func (o *MarketingEventCreateRequestParams) SetExternalEventId(v string) {
	o.ExternalEventId = v
}

// GetEventDescription returns the EventDescription field value if set, zero value otherwise.
func (o *MarketingEventCreateRequestParams) GetEventDescription() string {
	if o == nil || o.EventDescription == nil {
		var ret string
		return ret
	}
	return *o.EventDescription
}

// GetEventDescriptionOk returns a tuple with the EventDescription field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetEventDescriptionOk() (*string, bool) {
	if o == nil || o.EventDescription == nil {
		return nil, false
	}
	return o.EventDescription, true
}

// HasEventDescription returns a boolean if a field has been set.
func (o *MarketingEventCreateRequestParams) HasEventDescription() bool {
	if o != nil && o.EventDescription != nil {
		return true
	}

	return false
}

// SetEventDescription gets a reference to the given string and assigns it to the EventDescription field.
func (o *MarketingEventCreateRequestParams) SetEventDescription(v string) {
	o.EventDescription = &v
}

// GetEventName returns the EventName field value
func (o *MarketingEventCreateRequestParams) GetEventName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventName
}

// GetEventNameOk returns a tuple with the EventName field value
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetEventNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventName, true
}

// SetEventName sets field value
func (o *MarketingEventCreateRequestParams) SetEventName(v string) {
	o.EventName = v
}

// GetEventType returns the EventType field value if set, zero value otherwise.
func (o *MarketingEventCreateRequestParams) GetEventType() string {
	if o == nil || o.EventType == nil {
		var ret string
		return ret
	}
	return *o.EventType
}

// GetEventTypeOk returns a tuple with the EventType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetEventTypeOk() (*string, bool) {
	if o == nil || o.EventType == nil {
		return nil, false
	}
	return o.EventType, true
}

// HasEventType returns a boolean if a field has been set.
func (o *MarketingEventCreateRequestParams) HasEventType() bool {
	if o != nil && o.EventType != nil {
		return true
	}

	return false
}

// SetEventType gets a reference to the given string and assigns it to the EventType field.
func (o *MarketingEventCreateRequestParams) SetEventType(v string) {
	o.EventType = &v
}

// GetEndDateTime returns the EndDateTime field value if set, zero value otherwise.
func (o *MarketingEventCreateRequestParams) GetEndDateTime() time.Time {
	if o == nil || o.EndDateTime == nil {
		var ret time.Time
		return ret
	}
	return *o.EndDateTime
}

// GetEndDateTimeOk returns a tuple with the EndDateTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MarketingEventCreateRequestParams) GetEndDateTimeOk() (*time.Time, bool) {
	if o == nil || o.EndDateTime == nil {
		return nil, false
	}
	return o.EndDateTime, true
}

// HasEndDateTime returns a boolean if a field has been set.
func (o *MarketingEventCreateRequestParams) HasEndDateTime() bool {
	if o != nil && o.EndDateTime != nil {
		return true
	}

	return false
}

// SetEndDateTime gets a reference to the given time.Time and assigns it to the EndDateTime field.
func (o *MarketingEventCreateRequestParams) SetEndDateTime(v time.Time) {
	o.EndDateTime = &v
}

func (o MarketingEventCreateRequestParams) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.StartDateTime != nil {
		toSerialize["startDateTime"] = o.StartDateTime
	}
	if o.CustomProperties != nil {
		toSerialize["customProperties"] = o.CustomProperties
	}
	if true {
		toSerialize["externalAccountId"] = o.ExternalAccountId
	}
	if o.EventCancelled != nil {
		toSerialize["eventCancelled"] = o.EventCancelled
	}
	if true {
		toSerialize["eventOrganizer"] = o.EventOrganizer
	}
	if o.EventUrl != nil {
		toSerialize["eventUrl"] = o.EventUrl
	}
	if true {
		toSerialize["externalEventId"] = o.ExternalEventId
	}
	if o.EventDescription != nil {
		toSerialize["eventDescription"] = o.EventDescription
	}
	if true {
		toSerialize["eventName"] = o.EventName
	}
	if o.EventType != nil {
		toSerialize["eventType"] = o.EventType
	}
	if o.EndDateTime != nil {
		toSerialize["endDateTime"] = o.EndDateTime
	}
	return json.Marshal(toSerialize)
}

type NullableMarketingEventCreateRequestParams struct {
	value *MarketingEventCreateRequestParams
	isSet bool
}

func (v NullableMarketingEventCreateRequestParams) Get() *MarketingEventCreateRequestParams {
	return v.value
}

func (v *NullableMarketingEventCreateRequestParams) Set(val *MarketingEventCreateRequestParams) {
	v.value = val
	v.isSet = true
}

func (v NullableMarketingEventCreateRequestParams) IsSet() bool {
	return v.isSet
}

func (v *NullableMarketingEventCreateRequestParams) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMarketingEventCreateRequestParams(val *MarketingEventCreateRequestParams) *NullableMarketingEventCreateRequestParams {
	return &NullableMarketingEventCreateRequestParams{value: val, isSet: true}
}

func (v NullableMarketingEventCreateRequestParams) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMarketingEventCreateRequestParams) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
