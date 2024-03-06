/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the AssociatedId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AssociatedId{}

// AssociatedId struct for AssociatedId
type AssociatedId struct {
	Id   string `json:"id"`
	Type string `json:"type"`
}

type _AssociatedId AssociatedId

// NewAssociatedId instantiates a new AssociatedId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociatedId(id string, type_ string) *AssociatedId {
	this := AssociatedId{}
	this.Id = id
	this.Type = type_
	return &this
}

// NewAssociatedIdWithDefaults instantiates a new AssociatedId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociatedIdWithDefaults() *AssociatedId {
	this := AssociatedId{}
	return &this
}

// GetId returns the Id field value
func (o *AssociatedId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssociatedId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssociatedId) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *AssociatedId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *AssociatedId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *AssociatedId) SetType(v string) {
	o.Type = v
}

func (o AssociatedId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AssociatedId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *AssociatedId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"type",
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

	varAssociatedId := _AssociatedId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAssociatedId)

	if err != nil {
		return err
	}

	*o = AssociatedId(varAssociatedId)

	return err
}

type NullableAssociatedId struct {
	value *AssociatedId
	isSet bool
}

func (v NullableAssociatedId) Get() *AssociatedId {
	return v.value
}

func (v *NullableAssociatedId) Set(val *AssociatedId) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociatedId) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociatedId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociatedId(val *AssociatedId) *NullableAssociatedId {
	return &NullableAssociatedId{value: val, isSet: true}
}

func (v NullableAssociatedId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociatedId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
