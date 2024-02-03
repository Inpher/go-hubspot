/*
Hubdb

HubDB is a relational data store that presents data as rows, columns, and cells in a table, much like a spreadsheet. HubDB tables can be added or modified [in the HubSpot CMS](https://knowledge.hubspot.com/cos-general/how-to-edit-hubdb-tables), but you can also use the API endpoints documented here. For more information on HubDB tables and using their data on a HubSpot site, see the [CMS developers site](https://designers.hubspot.com/docs/tools/hubdb). You can also see the [documentation for dynamic pages](https://designers.hubspot.com/docs/tutorials/how-to-build-dynamic-pages-with-hubdb) for more details about the `useForPages` field.  HubDB tables support `draft` and `published` versions. This allows you to update data in the table, either for testing or to allow for a manual approval process, without affecting any live pages using the existing data. Draft data can be reviewed, and published by a user working in HubSpot or published via the API. Draft data can also be discarded, allowing users to go back to the published version of the data without disrupting it. If a table is set to be `allowed for public access`, you can access the published version of the table and rows without any authentication by specifying the portal id via the query parameter `portalId`.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package hubdb

import (
	"encoding/json"
)

// ForeignId struct for ForeignId
type ForeignId struct {
	//
	Name string `json:"name"`
	//
	Id string `json:"id"`
	//
	Type string `json:"type"`
}

// NewForeignId instantiates a new ForeignId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewForeignId(name string, id string, type_ string) *ForeignId {
	this := ForeignId{}
	this.Name = name
	this.Id = id
	this.Type = type_
	return &this
}

// NewForeignIdWithDefaults instantiates a new ForeignId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewForeignIdWithDefaults() *ForeignId {
	this := ForeignId{}
	return &this
}

// GetName returns the Name field value
func (o *ForeignId) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ForeignId) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ForeignId) SetName(v string) {
	o.Name = v
}

// GetId returns the Id field value
func (o *ForeignId) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ForeignId) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ForeignId) SetId(v string) {
	o.Id = v
}

// GetType returns the Type field value
func (o *ForeignId) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ForeignId) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ForeignId) SetType(v string) {
	o.Type = v
}

func (o ForeignId) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableForeignId struct {
	value *ForeignId
	isSet bool
}

func (v NullableForeignId) Get() *ForeignId {
	return v.value
}

func (v *NullableForeignId) Set(val *ForeignId) {
	v.value = val
	v.isSet = true
}

func (v NullableForeignId) IsSet() bool {
	return v.isSet
}

func (v *NullableForeignId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableForeignId(val *ForeignId) *NullableForeignId {
	return &NullableForeignId{value: val, isSet: true}
}

func (v NullableForeignId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableForeignId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}