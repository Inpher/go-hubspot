/*
Line Items

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package line_items

import (
	"encoding/json"
	"time"
)

// SimplePublicObjectWithAssociations struct for SimplePublicObjectWithAssociations
type SimplePublicObjectWithAssociations struct {
	Associations          *map[string]CollectionResponseAssociatedId `json:"associations,omitempty"`
	CreatedAt             time.Time                                  `json:"createdAt"`
	Archived              *bool                                      `json:"archived,omitempty"`
	ArchivedAt            *time.Time                                 `json:"archivedAt,omitempty"`
	PropertiesWithHistory *map[string][]ValueWithTimestamp           `json:"propertiesWithHistory,omitempty"`
	Id                    string                                     `json:"id"`
	Properties            map[string]string                          `json:"properties"`
	UpdatedAt             time.Time                                  `json:"updatedAt"`
}

// NewSimplePublicObjectWithAssociations instantiates a new SimplePublicObjectWithAssociations object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObjectWithAssociations(createdAt time.Time, id string, properties map[string]string, updatedAt time.Time) *SimplePublicObjectWithAssociations {
	this := SimplePublicObjectWithAssociations{}
	this.CreatedAt = createdAt
	this.Id = id
	this.Properties = properties
	this.UpdatedAt = updatedAt
	return &this
}

// NewSimplePublicObjectWithAssociationsWithDefaults instantiates a new SimplePublicObjectWithAssociations object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectWithAssociationsWithDefaults() *SimplePublicObjectWithAssociations {
	this := SimplePublicObjectWithAssociations{}
	return &this
}

// GetAssociations returns the Associations field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetAssociations() map[string]CollectionResponseAssociatedId {
	if o == nil || o.Associations == nil {
		var ret map[string]CollectionResponseAssociatedId
		return ret
	}
	return *o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetAssociationsOk() (*map[string]CollectionResponseAssociatedId, bool) {
	if o == nil || o.Associations == nil {
		return nil, false
	}
	return o.Associations, true
}

// HasAssociations returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasAssociations() bool {
	if o != nil && o.Associations != nil {
		return true
	}

	return false
}

// SetAssociations gets a reference to the given map[string]CollectionResponseAssociatedId and assigns it to the Associations field.
func (o *SimplePublicObjectWithAssociations) SetAssociations(v map[string]CollectionResponseAssociatedId) {
	o.Associations = &v
}

// GetCreatedAt returns the CreatedAt field value
func (o *SimplePublicObjectWithAssociations) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SimplePublicObjectWithAssociations) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetArchived() bool {
	if o == nil || o.Archived == nil {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetArchivedOk() (*bool, bool) {
	if o == nil || o.Archived == nil {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasArchived() bool {
	if o != nil && o.Archived != nil {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *SimplePublicObjectWithAssociations) SetArchived(v bool) {
	o.Archived = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetArchivedAt() time.Time {
	if o == nil || o.ArchivedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || o.ArchivedAt == nil {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasArchivedAt() bool {
	if o != nil && o.ArchivedAt != nil {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *SimplePublicObjectWithAssociations) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetPropertiesWithHistory returns the PropertiesWithHistory field value if set, zero value otherwise.
func (o *SimplePublicObjectWithAssociations) GetPropertiesWithHistory() map[string][]ValueWithTimestamp {
	if o == nil || o.PropertiesWithHistory == nil {
		var ret map[string][]ValueWithTimestamp
		return ret
	}
	return *o.PropertiesWithHistory
}

// GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetPropertiesWithHistoryOk() (*map[string][]ValueWithTimestamp, bool) {
	if o == nil || o.PropertiesWithHistory == nil {
		return nil, false
	}
	return o.PropertiesWithHistory, true
}

// HasPropertiesWithHistory returns a boolean if a field has been set.
func (o *SimplePublicObjectWithAssociations) HasPropertiesWithHistory() bool {
	if o != nil && o.PropertiesWithHistory != nil {
		return true
	}

	return false
}

// SetPropertiesWithHistory gets a reference to the given map[string][]ValueWithTimestamp and assigns it to the PropertiesWithHistory field.
func (o *SimplePublicObjectWithAssociations) SetPropertiesWithHistory(v map[string][]ValueWithTimestamp) {
	o.PropertiesWithHistory = &v
}

// GetId returns the Id field value
func (o *SimplePublicObjectWithAssociations) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObjectWithAssociations) SetId(v string) {
	o.Id = v
}

// GetProperties returns the Properties field value
func (o *SimplePublicObjectWithAssociations) GetProperties() map[string]string {
	if o == nil {
		var ret map[string]string
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Properties, true
}

// SetProperties sets field value
func (o *SimplePublicObjectWithAssociations) SetProperties(v map[string]string) {
	o.Properties = v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SimplePublicObjectWithAssociations) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObjectWithAssociations) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SimplePublicObjectWithAssociations) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SimplePublicObjectWithAssociations) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Associations != nil {
		toSerialize["associations"] = o.Associations
	}
	if true {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.Archived != nil {
		toSerialize["archived"] = o.Archived
	}
	if o.ArchivedAt != nil {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if o.PropertiesWithHistory != nil {
		toSerialize["propertiesWithHistory"] = o.PropertiesWithHistory
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableSimplePublicObjectWithAssociations struct {
	value *SimplePublicObjectWithAssociations
	isSet bool
}

func (v NullableSimplePublicObjectWithAssociations) Get() *SimplePublicObjectWithAssociations {
	return v.value
}

func (v *NullableSimplePublicObjectWithAssociations) Set(val *SimplePublicObjectWithAssociations) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObjectWithAssociations) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObjectWithAssociations) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObjectWithAssociations(val *SimplePublicObjectWithAssociations) *NullableSimplePublicObjectWithAssociations {
	return &NullableSimplePublicObjectWithAssociations{value: val, isSet: true}
}

func (v NullableSimplePublicObjectWithAssociations) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObjectWithAssociations) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
