/*
Schemas

The CRM uses schemas to define how custom objects should store and represent information in the HubSpot CRM. Schemas define details about an object's type, properties, and associations. The schema can be uniquely identified by its **object type ID**.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package schemas

import (
	"encoding/json"
	"time"
)

// ObjectSchema Defines an object schema, including its properties and associations.
type ObjectSchema struct {
	Labels ObjectTypeDefinitionLabels `json:"labels"`
	// The names of properties that should be **required** when creating an object of this type.
	RequiredProperties []string `json:"requiredProperties"`
	// Names of properties that will be indexed for this object type in by HubSpot's product search.
	SearchableProperties []string `json:"searchableProperties"`
	// The name of the primary property for this object. This will be displayed as primary on the HubSpot record page for this object type.
	PrimaryDisplayProperty *string `json:"primaryDisplayProperty,omitempty"`
	// The names of secondary properties for this object. These will be displayed as secondary on the HubSpot record page for this object type.
	SecondaryDisplayProperties []string `json:"secondaryDisplayProperties"`
	Archived                   bool     `json:"archived"`
	// A unique ID for this schema's object type. Will be defined as {meta-type}-{unique ID}.
	Id string `json:"id"`
	// An assigned unique ID for the object, including portal ID and object name.
	FullyQualifiedName string `json:"fullyQualifiedName"`
	// When the object schema was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// When the object schema was last updated.
	UpdatedAt    *time.Time `json:"updatedAt,omitempty"`
	ObjectTypeId string     `json:"objectTypeId"`
	// Properties defined for this object type.
	Properties []Property `json:"properties"`
	// Associations defined for a given object type.
	Associations []AssociationDefinition `json:"associations"`
	// A unique name for the schema's object type.
	Name string `json:"name"`
}

// NewObjectSchema instantiates a new ObjectSchema object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewObjectSchema(labels ObjectTypeDefinitionLabels, requiredProperties []string, searchableProperties []string, secondaryDisplayProperties []string, archived bool, id string, fullyQualifiedName string, objectTypeId string, properties []Property, associations []AssociationDefinition, name string) *ObjectSchema {
	this := ObjectSchema{}
	this.Labels = labels
	this.RequiredProperties = requiredProperties
	this.SearchableProperties = searchableProperties
	this.SecondaryDisplayProperties = secondaryDisplayProperties
	this.Archived = archived
	this.Id = id
	this.FullyQualifiedName = fullyQualifiedName
	this.ObjectTypeId = objectTypeId
	this.Properties = properties
	this.Associations = associations
	this.Name = name
	return &this
}

// NewObjectSchemaWithDefaults instantiates a new ObjectSchema object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewObjectSchemaWithDefaults() *ObjectSchema {
	this := ObjectSchema{}
	return &this
}

// GetLabels returns the Labels field value
func (o *ObjectSchema) GetLabels() ObjectTypeDefinitionLabels {
	if o == nil {
		var ret ObjectTypeDefinitionLabels
		return ret
	}

	return o.Labels
}

// GetLabelsOk returns a tuple with the Labels field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetLabelsOk() (*ObjectTypeDefinitionLabels, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Labels, true
}

// SetLabels sets field value
func (o *ObjectSchema) SetLabels(v ObjectTypeDefinitionLabels) {
	o.Labels = v
}

// GetRequiredProperties returns the RequiredProperties field value
func (o *ObjectSchema) GetRequiredProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RequiredProperties
}

// GetRequiredPropertiesOk returns a tuple with the RequiredProperties field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetRequiredPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequiredProperties, true
}

// SetRequiredProperties sets field value
func (o *ObjectSchema) SetRequiredProperties(v []string) {
	o.RequiredProperties = v
}

// GetSearchableProperties returns the SearchableProperties field value
func (o *ObjectSchema) GetSearchableProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SearchableProperties
}

// GetSearchablePropertiesOk returns a tuple with the SearchableProperties field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetSearchablePropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SearchableProperties, true
}

// SetSearchableProperties sets field value
func (o *ObjectSchema) SetSearchableProperties(v []string) {
	o.SearchableProperties = v
}

// GetPrimaryDisplayProperty returns the PrimaryDisplayProperty field value if set, zero value otherwise.
func (o *ObjectSchema) GetPrimaryDisplayProperty() string {
	if o == nil || o.PrimaryDisplayProperty == nil {
		var ret string
		return ret
	}
	return *o.PrimaryDisplayProperty
}

// GetPrimaryDisplayPropertyOk returns a tuple with the PrimaryDisplayProperty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetPrimaryDisplayPropertyOk() (*string, bool) {
	if o == nil || o.PrimaryDisplayProperty == nil {
		return nil, false
	}
	return o.PrimaryDisplayProperty, true
}

// HasPrimaryDisplayProperty returns a boolean if a field has been set.
func (o *ObjectSchema) HasPrimaryDisplayProperty() bool {
	if o != nil && o.PrimaryDisplayProperty != nil {
		return true
	}

	return false
}

// SetPrimaryDisplayProperty gets a reference to the given string and assigns it to the PrimaryDisplayProperty field.
func (o *ObjectSchema) SetPrimaryDisplayProperty(v string) {
	o.PrimaryDisplayProperty = &v
}

// GetSecondaryDisplayProperties returns the SecondaryDisplayProperties field value
func (o *ObjectSchema) GetSecondaryDisplayProperties() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.SecondaryDisplayProperties
}

// GetSecondaryDisplayPropertiesOk returns a tuple with the SecondaryDisplayProperties field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetSecondaryDisplayPropertiesOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.SecondaryDisplayProperties, true
}

// SetSecondaryDisplayProperties sets field value
func (o *ObjectSchema) SetSecondaryDisplayProperties(v []string) {
	o.SecondaryDisplayProperties = v
}

// GetArchived returns the Archived field value
func (o *ObjectSchema) GetArchived() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Archived
}

// GetArchivedOk returns a tuple with the Archived field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetArchivedOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Archived, true
}

// SetArchived sets field value
func (o *ObjectSchema) SetArchived(v bool) {
	o.Archived = v
}

// GetId returns the Id field value
func (o *ObjectSchema) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ObjectSchema) SetId(v string) {
	o.Id = v
}

// GetFullyQualifiedName returns the FullyQualifiedName field value
func (o *ObjectSchema) GetFullyQualifiedName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FullyQualifiedName
}

// GetFullyQualifiedNameOk returns a tuple with the FullyQualifiedName field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetFullyQualifiedNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FullyQualifiedName, true
}

// SetFullyQualifiedName sets field value
func (o *ObjectSchema) SetFullyQualifiedName(v string) {
	o.FullyQualifiedName = v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *ObjectSchema) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *ObjectSchema) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *ObjectSchema) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *ObjectSchema) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *ObjectSchema) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *ObjectSchema) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetObjectTypeId returns the ObjectTypeId field value
func (o *ObjectSchema) GetObjectTypeId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ObjectTypeId
}

// GetObjectTypeIdOk returns a tuple with the ObjectTypeId field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetObjectTypeIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ObjectTypeId, true
}

// SetObjectTypeId sets field value
func (o *ObjectSchema) SetObjectTypeId(v string) {
	o.ObjectTypeId = v
}

// GetProperties returns the Properties field value
func (o *ObjectSchema) GetProperties() []Property {
	if o == nil {
		var ret []Property
		return ret
	}

	return o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetPropertiesOk() ([]Property, bool) {
	if o == nil {
		return nil, false
	}
	return o.Properties, true
}

// SetProperties sets field value
func (o *ObjectSchema) SetProperties(v []Property) {
	o.Properties = v
}

// GetAssociations returns the Associations field value
func (o *ObjectSchema) GetAssociations() []AssociationDefinition {
	if o == nil {
		var ret []AssociationDefinition
		return ret
	}

	return o.Associations
}

// GetAssociationsOk returns a tuple with the Associations field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetAssociationsOk() ([]AssociationDefinition, bool) {
	if o == nil {
		return nil, false
	}
	return o.Associations, true
}

// SetAssociations sets field value
func (o *ObjectSchema) SetAssociations(v []AssociationDefinition) {
	o.Associations = v
}

// GetName returns the Name field value
func (o *ObjectSchema) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ObjectSchema) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ObjectSchema) SetName(v string) {
	o.Name = v
}

func (o ObjectSchema) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["labels"] = o.Labels
	}
	if true {
		toSerialize["requiredProperties"] = o.RequiredProperties
	}
	if true {
		toSerialize["searchableProperties"] = o.SearchableProperties
	}
	if o.PrimaryDisplayProperty != nil {
		toSerialize["primaryDisplayProperty"] = o.PrimaryDisplayProperty
	}
	if true {
		toSerialize["secondaryDisplayProperties"] = o.SecondaryDisplayProperties
	}
	if true {
		toSerialize["archived"] = o.Archived
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["fullyQualifiedName"] = o.FullyQualifiedName
	}
	if o.CreatedAt != nil {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if o.UpdatedAt != nil {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if true {
		toSerialize["objectTypeId"] = o.ObjectTypeId
	}
	if true {
		toSerialize["properties"] = o.Properties
	}
	if true {
		toSerialize["associations"] = o.Associations
	}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableObjectSchema struct {
	value *ObjectSchema
	isSet bool
}

func (v NullableObjectSchema) Get() *ObjectSchema {
	return v.value
}

func (v *NullableObjectSchema) Set(val *ObjectSchema) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectSchema) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectSchema) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectSchema(val *ObjectSchema) *NullableObjectSchema {
	return &NullableObjectSchema{value: val, isSet: true}
}

func (v NullableObjectSchema) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectSchema) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
