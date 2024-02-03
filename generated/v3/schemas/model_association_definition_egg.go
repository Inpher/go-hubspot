/*
Schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemas

import (
	"encoding/json"
)

// AssociationDefinitionEgg Defines an association between two object types.
type AssociationDefinitionEgg struct {
	// ID of the primary object type to link from.
	FromObjectTypeId string `json:"fromObjectTypeId"`
	// A unique name for this association.
	Name *string `json:"name,omitempty"`
	// ID of the target object type ID to link to.
	ToObjectTypeId string `json:"toObjectTypeId"`
}

// NewAssociationDefinitionEgg instantiates a new AssociationDefinitionEgg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssociationDefinitionEgg(fromObjectTypeId string, toObjectTypeId string) *AssociationDefinitionEgg {
	this := AssociationDefinitionEgg{}
	this.FromObjectTypeId = fromObjectTypeId
	this.ToObjectTypeId = toObjectTypeId
	return &this
}

// NewAssociationDefinitionEggWithDefaults instantiates a new AssociationDefinitionEgg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssociationDefinitionEggWithDefaults() *AssociationDefinitionEgg {
	this := AssociationDefinitionEgg{}
	return &this
}

// GetFromObjectTypeId returns the FromObjectTypeId field value
func (o *AssociationDefinitionEgg) GetFromObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FromObjectTypeId
}

// GetFromObjectTypeIdOk returns a tuple with the FromObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *AssociationDefinitionEgg) GetFromObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FromObjectTypeId, true
}

// SetFromObjectTypeId sets field value
func (o *AssociationDefinitionEgg) SetFromObjectTypeId(v string) {
	o.FromObjectTypeId = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssociationDefinitionEgg) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssociationDefinitionEgg) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssociationDefinitionEgg) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssociationDefinitionEgg) SetName(v string) {
	o.Name = &v
}

// GetToObjectTypeId returns the ToObjectTypeId field value
func (o *AssociationDefinitionEgg) GetToObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ToObjectTypeId
}

// GetToObjectTypeIdOk returns a tuple with the ToObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *AssociationDefinitionEgg) GetToObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ToObjectTypeId, true
}

// SetToObjectTypeId sets field value
func (o *AssociationDefinitionEgg) SetToObjectTypeId(v string) {
	o.ToObjectTypeId = v
}

func (o AssociationDefinitionEgg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["fromObjectTypeId"] = o.FromObjectTypeId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["toObjectTypeId"] = o.ToObjectTypeId
	}
	return json.Marshal(toSerialize)
}

type NullableAssociationDefinitionEgg struct {
	value *AssociationDefinitionEgg
	isSet bool
}

func (v NullableAssociationDefinitionEgg) Get() *AssociationDefinitionEgg {
	return v.value
}

func (v *NullableAssociationDefinitionEgg) Set(val *AssociationDefinitionEgg) {
	v.value = val
	v.isSet = true
}

func (v NullableAssociationDefinitionEgg) IsSet() bool {
	return v.isSet
}

func (v *NullableAssociationDefinitionEgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssociationDefinitionEgg(val *AssociationDefinitionEgg) *NullableAssociationDefinitionEgg {
	return &NullableAssociationDefinitionEgg{value: val, isSet: true}
}

func (v NullableAssociationDefinitionEgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssociationDefinitionEgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}