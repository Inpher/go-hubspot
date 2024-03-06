/*
Products

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package products

import (
	"bytes"
	"encoding/json"
	"fmt"
	"time"
)

// checks if the SimplePublicObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SimplePublicObject{}

// SimplePublicObject struct for SimplePublicObject
type SimplePublicObject struct {
	CreatedAt             time.Time                        `json:"createdAt"`
	Archived              *bool                            `json:"archived,omitempty"`
	ArchivedAt            *time.Time                       `json:"archivedAt,omitempty"`
	PropertiesWithHistory *map[string][]ValueWithTimestamp `json:"propertiesWithHistory,omitempty"`
	Id                    string                           `json:"id"`
	Properties            *map[string]string               `json:"properties,omitempty"`
	UpdatedAt             time.Time                        `json:"updatedAt"`
}

type _SimplePublicObject SimplePublicObject

// NewSimplePublicObject instantiates a new SimplePublicObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimplePublicObject(createdAt time.Time, id string, updatedAt time.Time) *SimplePublicObject {
	this := SimplePublicObject{}
	this.CreatedAt = createdAt
	this.Id = id
	this.UpdatedAt = updatedAt
	return &this
}

// NewSimplePublicObjectWithDefaults instantiates a new SimplePublicObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimplePublicObjectWithDefaults() *SimplePublicObject {
	this := SimplePublicObject{}
	return &this
}

// GetCreatedAt returns the CreatedAt field value
func (o *SimplePublicObject) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *SimplePublicObject) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetArchived returns the Archived field value if set, zero value otherwise.
func (o *SimplePublicObject) GetArchived() bool {
	if o == nil || IsNil(o.Archived) {
		var ret bool
		return ret
	}
	return *o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetArchivedOk() (*bool, bool) {
	if o == nil || IsNil(o.Archived) {
		return nil, false
	}
	return o.Archived, true
}

// HasArchived returns a boolean if a field has been set.
func (o *SimplePublicObject) HasArchived() bool {
	if o != nil && !IsNil(o.Archived) {
		return true
	}

	return false
}

// SetArchived gets a reference to the given bool and assigns it to the Archived field.
func (o *SimplePublicObject) SetArchived(v bool) {
	o.Archived = &v
}

// GetArchivedAt returns the ArchivedAt field value if set, zero value otherwise.
func (o *SimplePublicObject) GetArchivedAt() time.Time {
	if o == nil || IsNil(o.ArchivedAt) {
		var ret time.Time
		return ret
	}
	return *o.ArchivedAt
}

// GetArchivedAtOk returns a tuple with the ArchivedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetArchivedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ArchivedAt) {
		return nil, false
	}
	return o.ArchivedAt, true
}

// HasArchivedAt returns a boolean if a field has been set.
func (o *SimplePublicObject) HasArchivedAt() bool {
	if o != nil && !IsNil(o.ArchivedAt) {
		return true
	}

	return false
}

// SetArchivedAt gets a reference to the given time.Time and assigns it to the ArchivedAt field.
func (o *SimplePublicObject) SetArchivedAt(v time.Time) {
	o.ArchivedAt = &v
}

// GetPropertiesWithHistory returns the PropertiesWithHistory field value if set, zero value otherwise.
func (o *SimplePublicObject) GetPropertiesWithHistory() map[string][]ValueWithTimestamp {
	if o == nil || IsNil(o.PropertiesWithHistory) {
		var ret map[string][]ValueWithTimestamp
		return ret
	}
	return *o.PropertiesWithHistory
}

// GetPropertiesWithHistoryOk returns a tuple with the PropertiesWithHistory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetPropertiesWithHistoryOk() (*map[string][]ValueWithTimestamp, bool) {
	if o == nil || IsNil(o.PropertiesWithHistory) {
		return nil, false
	}
	return o.PropertiesWithHistory, true
}

// HasPropertiesWithHistory returns a boolean if a field has been set.
func (o *SimplePublicObject) HasPropertiesWithHistory() bool {
	if o != nil && !IsNil(o.PropertiesWithHistory) {
		return true
	}

	return false
}

// SetPropertiesWithHistory gets a reference to the given map[string][]ValueWithTimestamp and assigns it to the PropertiesWithHistory field.
func (o *SimplePublicObject) SetPropertiesWithHistory(v map[string][]ValueWithTimestamp) {
	o.PropertiesWithHistory = &v
}

// GetId returns the Id field value
func (o *SimplePublicObject) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *SimplePublicObject) SetId(v string) {
	o.Id = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *SimplePublicObject) GetProperties() map[string]string {
	if o == nil || IsNil(o.Properties) {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || IsNil(o.Properties) {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *SimplePublicObject) HasProperties() bool {
	if o != nil && !IsNil(o.Properties) {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *SimplePublicObject) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetUpdatedAt returns the UpdatedAt field value
func (o *SimplePublicObject) GetUpdatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value
// and a boolean to check if the value has been set.
func (o *SimplePublicObject) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UpdatedAt, true
}

// SetUpdatedAt sets field value
func (o *SimplePublicObject) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = v
}

func (o SimplePublicObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SimplePublicObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["createdAt"] = o.CreatedAt
	if !IsNil(o.Archived) {
		toSerialize["archived"] = o.Archived
	}
	if !IsNil(o.ArchivedAt) {
		toSerialize["archivedAt"] = o.ArchivedAt
	}
	if !IsNil(o.PropertiesWithHistory) {
		toSerialize["propertiesWithHistory"] = o.PropertiesWithHistory
	}
	toSerialize["id"] = o.Id
	if !IsNil(o.Properties) {
		toSerialize["properties"] = o.Properties
	}
	toSerialize["updatedAt"] = o.UpdatedAt
	return toSerialize, nil
}

func (o *SimplePublicObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"createdAt",
		"id",
		"updatedAt",
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

	varSimplePublicObject := _SimplePublicObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSimplePublicObject)

	if err != nil {
		return err
	}

	*o = SimplePublicObject(varSimplePublicObject)

	return err
}

type NullableSimplePublicObject struct {
	value *SimplePublicObject
	isSet bool
}

func (v NullableSimplePublicObject) Get() *SimplePublicObject {
	return v.value
}

func (v *NullableSimplePublicObject) Set(val *SimplePublicObject) {
	v.value = val
	v.isSet = true
}

func (v NullableSimplePublicObject) IsSet() bool {
	return v.isSet
}

func (v *NullableSimplePublicObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimplePublicObject(val *SimplePublicObject) *NullableSimplePublicObject {
	return &NullableSimplePublicObject{value: val, isSet: true}
}

func (v NullableSimplePublicObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimplePublicObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
