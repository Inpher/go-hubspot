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

// checks if the PublicAssociationsForObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PublicAssociationsForObject{}

// PublicAssociationsForObject struct for PublicAssociationsForObject
type PublicAssociationsForObject struct {
	Types []AssociationSpec `json:"types"`
	To    PublicObjectId    `json:"to"`
}

type _PublicAssociationsForObject PublicAssociationsForObject

// NewPublicAssociationsForObject instantiates a new PublicAssociationsForObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicAssociationsForObject(types []AssociationSpec, to PublicObjectId) *PublicAssociationsForObject {
	this := PublicAssociationsForObject{}
	this.Types = types
	this.To = to
	return &this
}

// NewPublicAssociationsForObjectWithDefaults instantiates a new PublicAssociationsForObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicAssociationsForObjectWithDefaults() *PublicAssociationsForObject {
	this := PublicAssociationsForObject{}
	return &this
}

// GetTypes returns the Types field value
func (o *PublicAssociationsForObject) GetTypes() []AssociationSpec {
	if o == nil {
		var ret []AssociationSpec
		return ret
	}

	return o.Types
}

// GetTypesOk returns a tuple with the Types field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationsForObject) GetTypesOk() ([]AssociationSpec, bool) {
	if o == nil {
		return nil, false
	}
	return o.Types, true
}

// SetTypes sets field value
func (o *PublicAssociationsForObject) SetTypes(v []AssociationSpec) {
	o.Types = v
}

// GetTo returns the To field value
func (o *PublicAssociationsForObject) GetTo() PublicObjectId {
	if o == nil {
		var ret PublicObjectId
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *PublicAssociationsForObject) GetToOk() (*PublicObjectId, bool) {
	if o == nil {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *PublicAssociationsForObject) SetTo(v PublicObjectId) {
	o.To = v
}

func (o PublicAssociationsForObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PublicAssociationsForObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["types"] = o.Types
	toSerialize["to"] = o.To
	return toSerialize, nil
}

func (o *PublicAssociationsForObject) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"types",
		"to",
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

	varPublicAssociationsForObject := _PublicAssociationsForObject{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPublicAssociationsForObject)

	if err != nil {
		return err
	}

	*o = PublicAssociationsForObject(varPublicAssociationsForObject)

	return err
}

type NullablePublicAssociationsForObject struct {
	value *PublicAssociationsForObject
	isSet bool
}

func (v NullablePublicAssociationsForObject) Get() *PublicAssociationsForObject {
	return v.value
}

func (v *NullablePublicAssociationsForObject) Set(val *PublicAssociationsForObject) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicAssociationsForObject) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicAssociationsForObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicAssociationsForObject(val *PublicAssociationsForObject) *NullablePublicAssociationsForObject {
	return &NullablePublicAssociationsForObject{value: val, isSet: true}
}

func (v NullablePublicAssociationsForObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicAssociationsForObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
