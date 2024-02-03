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

// Column struct for Column
type Column struct {
	// Foreign table id referenced
	ForeignTableId *int64  `json:"foreignTableId,omitempty"`
	Description    *string `json:"description,omitempty"`
	// Label of the column
	Label string `json:"label"`
	// Type of the column
	Type string `json:"type"`
	// Number of options available
	OptionCount *int32 `json:"optionCount,omitempty"`
	// Foreign Ids
	ForeignIds []ForeignId `json:"foreignIds,omitempty"`
	Deleted    *bool       `json:"deleted,omitempty"`
	// Name of the column
	Name string `json:"name"`
	// Options to choose for select and multi-select columns
	Options []Option `json:"options,omitempty"`
	// Column width for HubDB UI
	Width *int32 `json:"width,omitempty"`
	// Column Id
	Id *string `json:"id,omitempty"`
	// Foreign ids
	ForeignIdsById *map[string]ForeignId `json:"foreignIdsById,omitempty"`
	// Foreign Column id
	ForeignColumnId *int32 `json:"foreignColumnId,omitempty"`
	// Foreign ids by name
	ForeignIdsByName *map[string]ForeignId `json:"foreignIdsByName,omitempty"`
}

// NewColumn instantiates a new Column object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewColumn(label string, type_ string, name string) *Column {
	this := Column{}
	this.Label = label
	this.Type = type_
	this.Name = name
	return &this
}

// NewColumnWithDefaults instantiates a new Column object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewColumnWithDefaults() *Column {
	this := Column{}
	return &this
}

// GetForeignTableId returns the ForeignTableId field value if set, zero value otherwise.
func (o *Column) GetForeignTableId() int64 {
	if o == nil || o.ForeignTableId == nil {
		var ret int64
		return ret
	}
	return *o.ForeignTableId
}

// GetForeignTableIdOk returns a tuple with the ForeignTableId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetForeignTableIdOk() (*int64, bool) {
	if o == nil || o.ForeignTableId == nil {
		return nil, false
	}
	return o.ForeignTableId, true
}

// HasForeignTableId returns a boolean if a field has been set.
func (o *Column) HasForeignTableId() bool {
	if o != nil && o.ForeignTableId != nil {
		return true
	}

	return false
}

// SetForeignTableId gets a reference to the given int64 and assigns it to the ForeignTableId field.
func (o *Column) SetForeignTableId(v int64) {
	o.ForeignTableId = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Column) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Column) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Column) SetDescription(v string) {
	o.Description = &v
}

// GetLabel returns the Label field value
func (o *Column) GetLabel() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Label
}

// GetLabelOk returns a tuple with the Label field value
// and a boolean to check if the value has been set.
func (o *Column) GetLabelOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Label, true
}

// SetLabel sets field value
func (o *Column) SetLabel(v string) {
	o.Label = v
}

// GetType returns the Type field value
func (o *Column) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Column) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Column) SetType(v string) {
	o.Type = v
}

// GetOptionCount returns the OptionCount field value if set, zero value otherwise.
func (o *Column) GetOptionCount() int32 {
	if o == nil || o.OptionCount == nil {
		var ret int32
		return ret
	}
	return *o.OptionCount
}

// GetOptionCountOk returns a tuple with the OptionCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetOptionCountOk() (*int32, bool) {
	if o == nil || o.OptionCount == nil {
		return nil, false
	}
	return o.OptionCount, true
}

// HasOptionCount returns a boolean if a field has been set.
func (o *Column) HasOptionCount() bool {
	if o != nil && o.OptionCount != nil {
		return true
	}

	return false
}

// SetOptionCount gets a reference to the given int32 and assigns it to the OptionCount field.
func (o *Column) SetOptionCount(v int32) {
	o.OptionCount = &v
}

// GetForeignIds returns the ForeignIds field value if set, zero value otherwise.
func (o *Column) GetForeignIds() []ForeignId {
	if o == nil || o.ForeignIds == nil {
		var ret []ForeignId
		return ret
	}
	return o.ForeignIds
}

// GetForeignIdsOk returns a tuple with the ForeignIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetForeignIdsOk() ([]ForeignId, bool) {
	if o == nil || o.ForeignIds == nil {
		return nil, false
	}
	return o.ForeignIds, true
}

// HasForeignIds returns a boolean if a field has been set.
func (o *Column) HasForeignIds() bool {
	if o != nil && o.ForeignIds != nil {
		return true
	}

	return false
}

// SetForeignIds gets a reference to the given []ForeignId and assigns it to the ForeignIds field.
func (o *Column) SetForeignIds(v []ForeignId) {
	o.ForeignIds = v
}

// GetDeleted returns the Deleted field value if set, zero value otherwise.
func (o *Column) GetDeleted() bool {
	if o == nil || o.Deleted == nil {
		var ret bool
		return ret
	}
	return *o.Deleted
}

// GetDeletedOk returns a tuple with the Deleted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetDeletedOk() (*bool, bool) {
	if o == nil || o.Deleted == nil {
		return nil, false
	}
	return o.Deleted, true
}

// HasDeleted returns a boolean if a field has been set.
func (o *Column) HasDeleted() bool {
	if o != nil && o.Deleted != nil {
		return true
	}

	return false
}

// SetDeleted gets a reference to the given bool and assigns it to the Deleted field.
func (o *Column) SetDeleted(v bool) {
	o.Deleted = &v
}

// GetName returns the Name field value
func (o *Column) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Column) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Column) SetName(v string) {
	o.Name = v
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *Column) GetOptions() []Option {
	if o == nil || o.Options == nil {
		var ret []Option
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetOptionsOk() ([]Option, bool) {
	if o == nil || o.Options == nil {
		return nil, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *Column) HasOptions() bool {
	if o != nil && o.Options != nil {
		return true
	}

	return false
}

// SetOptions gets a reference to the given []Option and assigns it to the Options field.
func (o *Column) SetOptions(v []Option) {
	o.Options = v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *Column) GetWidth() int32 {
	if o == nil || o.Width == nil {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetWidthOk() (*int32, bool) {
	if o == nil || o.Width == nil {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *Column) HasWidth() bool {
	if o != nil && o.Width != nil {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *Column) SetWidth(v int32) {
	o.Width = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Column) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Column) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Column) SetId(v string) {
	o.Id = &v
}

// GetForeignIdsById returns the ForeignIdsById field value if set, zero value otherwise.
func (o *Column) GetForeignIdsById() map[string]ForeignId {
	if o == nil || o.ForeignIdsById == nil {
		var ret map[string]ForeignId
		return ret
	}
	return *o.ForeignIdsById
}

// GetForeignIdsByIdOk returns a tuple with the ForeignIdsById field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetForeignIdsByIdOk() (*map[string]ForeignId, bool) {
	if o == nil || o.ForeignIdsById == nil {
		return nil, false
	}
	return o.ForeignIdsById, true
}

// HasForeignIdsById returns a boolean if a field has been set.
func (o *Column) HasForeignIdsById() bool {
	if o != nil && o.ForeignIdsById != nil {
		return true
	}

	return false
}

// SetForeignIdsById gets a reference to the given map[string]ForeignId and assigns it to the ForeignIdsById field.
func (o *Column) SetForeignIdsById(v map[string]ForeignId) {
	o.ForeignIdsById = &v
}

// GetForeignColumnId returns the ForeignColumnId field value if set, zero value otherwise.
func (o *Column) GetForeignColumnId() int32 {
	if o == nil || o.ForeignColumnId == nil {
		var ret int32
		return ret
	}
	return *o.ForeignColumnId
}

// GetForeignColumnIdOk returns a tuple with the ForeignColumnId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetForeignColumnIdOk() (*int32, bool) {
	if o == nil || o.ForeignColumnId == nil {
		return nil, false
	}
	return o.ForeignColumnId, true
}

// HasForeignColumnId returns a boolean if a field has been set.
func (o *Column) HasForeignColumnId() bool {
	if o != nil && o.ForeignColumnId != nil {
		return true
	}

	return false
}

// SetForeignColumnId gets a reference to the given int32 and assigns it to the ForeignColumnId field.
func (o *Column) SetForeignColumnId(v int32) {
	o.ForeignColumnId = &v
}

// GetForeignIdsByName returns the ForeignIdsByName field value if set, zero value otherwise.
func (o *Column) GetForeignIdsByName() map[string]ForeignId {
	if o == nil || o.ForeignIdsByName == nil {
		var ret map[string]ForeignId
		return ret
	}
	return *o.ForeignIdsByName
}

// GetForeignIdsByNameOk returns a tuple with the ForeignIdsByName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Column) GetForeignIdsByNameOk() (*map[string]ForeignId, bool) {
	if o == nil || o.ForeignIdsByName == nil {
		return nil, false
	}
	return o.ForeignIdsByName, true
}

// HasForeignIdsByName returns a boolean if a field has been set.
func (o *Column) HasForeignIdsByName() bool {
	if o != nil && o.ForeignIdsByName != nil {
		return true
	}

	return false
}

// SetForeignIdsByName gets a reference to the given map[string]ForeignId and assigns it to the ForeignIdsByName field.
func (o *Column) SetForeignIdsByName(v map[string]ForeignId) {
	o.ForeignIdsByName = &v
}

func (o Column) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ForeignTableId != nil {
		toSerialize["foreignTableId"] = o.ForeignTableId
	}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if true {
		toSerialize["label"] = o.Label
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.OptionCount != nil {
		toSerialize["optionCount"] = o.OptionCount
	}
	if o.ForeignIds != nil {
		toSerialize["foreignIds"] = o.ForeignIds
	}
	if o.Deleted != nil {
		toSerialize["deleted"] = o.Deleted
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.Options != nil {
		toSerialize["options"] = o.Options
	}
	if o.Width != nil {
		toSerialize["width"] = o.Width
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ForeignIdsById != nil {
		toSerialize["foreignIdsById"] = o.ForeignIdsById
	}
	if o.ForeignColumnId != nil {
		toSerialize["foreignColumnId"] = o.ForeignColumnId
	}
	if o.ForeignIdsByName != nil {
		toSerialize["foreignIdsByName"] = o.ForeignIdsByName
	}
	return json.Marshal(toSerialize)
}

type NullableColumn struct {
	value *Column
	isSet bool
}

func (v NullableColumn) Get() *Column {
	return v.value
}

func (v *NullableColumn) Set(val *Column) {
	v.value = val
	v.isSet = true
}

func (v NullableColumn) IsSet() bool {
	return v.isSet
}

func (v *NullableColumn) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableColumn(val *Column) *NullableColumn {
	return &NullableColumn{value: val, isSet: true}
}

func (v NullableColumn) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableColumn) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}