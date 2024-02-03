/*
Imports

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package imports

import (
	"encoding/json"
)

// ImportRowCore struct for ImportRowCore
type ImportRowCore struct {
	RowData    []string `json:"rowData"`
	LineNumber int32    `json:"lineNumber"`
	PageName   *string  `json:"pageName,omitempty"`
	FileId     int32    `json:"fileId"`
}

// NewImportRowCore instantiates a new ImportRowCore object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewImportRowCore(rowData []string, lineNumber int32, fileId int32) *ImportRowCore {
	this := ImportRowCore{}
	this.RowData = rowData
	this.LineNumber = lineNumber
	this.FileId = fileId
	return &this
}

// NewImportRowCoreWithDefaults instantiates a new ImportRowCore object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewImportRowCoreWithDefaults() *ImportRowCore {
	this := ImportRowCore{}
	return &this
}

// GetRowData returns the RowData field value
func (o *ImportRowCore) GetRowData() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.RowData
}

// GetRowDataOk returns a tuple with the RowData field value
// and a boolean to check if the value has been set.
func (o *ImportRowCore) GetRowDataOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.RowData, true
}

// SetRowData sets field value
func (o *ImportRowCore) SetRowData(v []string) {
	o.RowData = v
}

// GetLineNumber returns the LineNumber field value
func (o *ImportRowCore) GetLineNumber() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.LineNumber
}

// GetLineNumberOk returns a tuple with the LineNumber field value
// and a boolean to check if the value has been set.
func (o *ImportRowCore) GetLineNumberOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LineNumber, true
}

// SetLineNumber sets field value
func (o *ImportRowCore) SetLineNumber(v int32) {
	o.LineNumber = v
}

// GetPageName returns the PageName field value if set, zero value otherwise.
func (o *ImportRowCore) GetPageName() string {
	if o == nil || o.PageName == nil {
		var ret string
		return ret
	}
	return *o.PageName
}

// GetPageNameOk returns a tuple with the PageName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ImportRowCore) GetPageNameOk() (*string, bool) {
	if o == nil || o.PageName == nil {
		return nil, false
	}
	return o.PageName, true
}

// HasPageName returns a boolean if a field has been set.
func (o *ImportRowCore) HasPageName() bool {
	if o != nil && o.PageName != nil {
		return true
	}

	return false
}

// SetPageName gets a reference to the given string and assigns it to the PageName field.
func (o *ImportRowCore) SetPageName(v string) {
	o.PageName = &v
}

// GetFileId returns the FileId field value
func (o *ImportRowCore) GetFileId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.FileId
}

// GetFileIdOk returns a tuple with the FileId field value
// and a boolean to check if the value has been set.
func (o *ImportRowCore) GetFileIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.FileId, true
}

// SetFileId sets field value
func (o *ImportRowCore) SetFileId(v int32) {
	o.FileId = v
}

func (o ImportRowCore) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["rowData"] = o.RowData
	}
	if true {
		toSerialize["lineNumber"] = o.LineNumber
	}
	if o.PageName != nil {
		toSerialize["pageName"] = o.PageName
	}
	if true {
		toSerialize["fileId"] = o.FileId
	}
	return json.Marshal(toSerialize)
}

type NullableImportRowCore struct {
	value *ImportRowCore
	isSet bool
}

func (v NullableImportRowCore) Get() *ImportRowCore {
	return v.value
}

func (v *NullableImportRowCore) Set(val *ImportRowCore) {
	v.value = val
	v.isSet = true
}

func (v NullableImportRowCore) IsSet() bool {
	return v.isSet
}

func (v *NullableImportRowCore) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableImportRowCore(val *ImportRowCore) *NullableImportRowCore {
	return &NullableImportRowCore{value: val, isSet: true}
}

func (v NullableImportRowCore) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableImportRowCore) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}