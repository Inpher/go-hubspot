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

// ObjectSchemaEgg Defines a new object type, its properties, and associations.
type ObjectSchemaEgg struct {
	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties,omitempty"`
	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties []string `json:"requiredProperties"`
	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties []string `json:"searchableProperties,omitempty"`
	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty *string `json:"primaryDisplayProperty,omitempty"`
	// A unique name for this object. For internal use only.
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// Associations defined for this object type.
	AssociatedObjects []string `json:"associatedObjects"`
	// Properties defined for this object type.
	Properties []ObjectTypePropertyCreate `json:"properties"`
	Labels     ObjectTypeDefinitionLabels `json:"labels"`
}

// NewObjectSchemaEgg instantiates a new ObjectSchemaEgg object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectSchemaEgg(requiredProperties []string, name string, associatedObjects []string, properties []ObjectTypePropertyCreate, labels ObjectTypeDefinitionLabels) *ObjectSchemaEgg {
	this := ObjectSchemaEgg{}
	this.RequiredProperties = requiredProperties
	this.Name = name
	this.AssociatedObjects = associatedObjects
	this.Properties = properties
	this.Labels = labels
	return &this
}

// NewObjectSchemaEggWithDefaults instantiates a new ObjectSchemaEgg object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectSchemaEggWithDefaults() *ObjectSchemaEgg {
	this := ObjectSchemaEgg{}
	return &this
}

// GetSecondaryDisplayProperties returns the SecondaryDisplayProperties field value if set, zero value otherwise.
func (o *ObjectSchemaEgg) GetSecondaryDisplayProperties() []string {
	if o == nil || o.SecondaryDisplayProperties == nil {
		var ret []string
		return ret
	}
	return o.SecondaryDisplayProperties
}

// GetSecondaryDisplayPropertiesOk returns a tuple with the SecondaryDisplayProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetSecondaryDisplayPropertiesOk() ([]string, bool) {
	if o == nil || o.SecondaryDisplayProperties == nil {
		return nil, false
	}
	return o.SecondaryDisplayProperties, true
}

// HasSecondaryDisplayProperties returns a boolean if a field has been set.
func (o *ObjectSchemaEgg) HasSecondaryDisplayProperties() bool {
	if o != nil && o.SecondaryDisplayProperties != nil {
		return true
	}

	return false
}

// SetSecondaryDisplayProperties gets a reference to the given []string and assigns it to the SecondaryDisplayProperties field.
func (o *ObjectSchemaEgg) SetSecondaryDisplayProperties(v []string) {
	o.SecondaryDisplayProperties = v
}

// GetRequiredProperties returns the RequiredProperties field value
func (o *ObjectSchemaEgg) GetRequiredProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequiredProperties
}

// GetRequiredPropertiesOk returns a tuple with the RequiredProperties field value
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetRequiredPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredProperties, true
}

// SetRequiredProperties sets field value
func (o *ObjectSchemaEgg) SetRequiredProperties(v []string) {
	o.RequiredProperties = v
}

// GetSearchableProperties returns the SearchableProperties field value if set, zero value otherwise.
func (o *ObjectSchemaEgg) GetSearchableProperties() []string {
	if o == nil || o.SearchableProperties == nil {
		var ret []string
		return ret
	}
	return o.SearchableProperties
}

// GetSearchablePropertiesOk returns a tuple with the SearchableProperties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetSearchablePropertiesOk() ([]string, bool) {
	if o == nil || o.SearchableProperties == nil {
		return nil, false
	}
	return o.SearchableProperties, true
}

// HasSearchableProperties returns a boolean if a field has been set.
func (o *ObjectSchemaEgg) HasSearchableProperties() bool {
	if o != nil && o.SearchableProperties != nil {
		return true
	}

	return false
}

// SetSearchableProperties gets a reference to the given []string and assigns it to the SearchableProperties field.
func (o *ObjectSchemaEgg) SetSearchableProperties(v []string) {
	o.SearchableProperties = v
}

// GetPrimaryDisplayProperty returns the PrimaryDisplayProperty field value if set, zero value otherwise.
func (o *ObjectSchemaEgg) GetPrimaryDisplayProperty() string {
	if o == nil || o.PrimaryDisplayProperty == nil {
		var ret string
		return ret
	}
	return *o.PrimaryDisplayProperty
}

// GetPrimaryDisplayPropertyOk returns a tuple with the PrimaryDisplayProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetPrimaryDisplayPropertyOk() (*string, bool) {
	if o == nil || o.PrimaryDisplayProperty == nil {
		return nil, false
	}
	return o.PrimaryDisplayProperty, true
}

// HasPrimaryDisplayProperty returns a boolean if a field has been set.
func (o *ObjectSchemaEgg) HasPrimaryDisplayProperty() bool {
	if o != nil && o.PrimaryDisplayProperty != nil {
		return true
	}

	return false
}

// SetPrimaryDisplayProperty gets a reference to the given string and assigns it to the PrimaryDisplayProperty field.
func (o *ObjectSchemaEgg) SetPrimaryDisplayProperty(v string) {
	o.PrimaryDisplayProperty = &v
}

// GetName returns the Name field value
func (o *ObjectSchemaEgg) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjectSchemaEgg) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *ObjectSchemaEgg) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *ObjectSchemaEgg) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *ObjectSchemaEgg) SetDescription(v string) {
	o.Description = &v
}

// GetAssociatedObjects returns the AssociatedObjects field value
func (o *ObjectSchemaEgg) GetAssociatedObjects() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.AssociatedObjects
}

// GetAssociatedObjectsOk returns a tuple with the AssociatedObjects field value
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetAssociatedObjectsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.AssociatedObjects, true
}

// SetAssociatedObjects sets field value
func (o *ObjectSchemaEgg) SetAssociatedObjects(v []string) {
	o.AssociatedObjects = v
}

// GetProperties returns the Properties field value
func (o *ObjectSchemaEgg) GetProperties() []ObjectTypePropertyCreate {
	if o == nil {
		var ret []ObjectTypePropertyCreate
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetPropertiesOk() ([]ObjectTypePropertyCreate, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *ObjectSchemaEgg) SetProperties(v []ObjectTypePropertyCreate) {
	o.Properties = v
}

// GetLabels returns the Labels field value
func (o *ObjectSchemaEgg) GetLabels() ObjectTypeDefinitionLabels {
	if o == nil {
		var ret ObjectTypeDefinitionLabels
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ObjectSchemaEgg) GetLabelsOk() (*ObjectTypeDefinitionLabels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *ObjectSchemaEgg) SetLabels(v ObjectTypeDefinitionLabels) {
	o.Labels = v
}

func (o ObjectSchemaEgg) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.SecondaryDisplayProperties != nil {
		toSerialize["secondaryDisplayProperties"] = o.SecondaryDisplayProperties
	}
	if true {
		toSerialize["requiredProperties"] = o.RequiredProperties
	}
	if o.SearchableProperties != nil {
		toSerialize["searchableProperties"] = o.SearchableProperties
	}
	if o.PrimaryDisplayProperty != nil {
		toSerialize["primaryDisplayProperty"] = o.PrimaryDisplayProperty
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["associatedObjects"] = o.AssociatedObjects
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["labels"] = o.Labels
	}
	return json.Marshal(toSerialize)
}

type NullableObjectSchemaEgg struct {
	value *ObjectSchemaEgg
	isSet bool
}

func (v NullableObjectSchemaEgg) Get() *ObjectSchemaEgg {
	return v.value
}

func (v *NullableObjectSchemaEgg) Set(val *ObjectSchemaEgg) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectSchemaEgg) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectSchemaEgg) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectSchemaEgg(val *ObjectSchemaEgg) *NullableObjectSchemaEgg {
	return &NullableObjectSchemaEgg{value: val, isSet: true}
}

func (v NullableObjectSchemaEgg) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectSchemaEgg) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}