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

// PublicImportMetadata An object detailing a summary of the import record outputs
type PublicImportMetadata struct {
	// Summarized outcomes of each row a developer attempted to import into HubSpot.
	Counters map[string]int32 `json:"counters"`
	// The IDs of files uploaded in the File Manager API.
	FileIds []string `json:"fileIds"`
	// The lists containing the imported objects.
	ObjectLists []PublicObjectListRecord `json:"objectLists"`
}

// NewPublicImportMetadata instantiates a new PublicImportMetadata object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicImportMetadata(counters map[string]int32, fileIds []string, objectLists []PublicObjectListRecord) *PublicImportMetadata {
	this := PublicImportMetadata{}
	this.Counters = counters
	this.FileIds = fileIds
	this.ObjectLists = objectLists
	return &this
}

// NewPublicImportMetadataWithDefaults instantiates a new PublicImportMetadata object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicImportMetadataWithDefaults() *PublicImportMetadata {
	this := PublicImportMetadata{}
	return &this
}

// GetCounters returns the Counters field value
func (o *PublicImportMetadata) GetCounters() map[string]int32 {
	if o == nil {
		var ret map[string]int32
		return ret
	}

	return o.Counters
}

// GetCountersOk returns a tuple with the Counters field value
// and a boolean to check if the value has been set.
func (o *PublicImportMetadata) GetCountersOk() (*map[string]int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Counters, true
}

// SetCounters sets field value
func (o *PublicImportMetadata) SetCounters(v map[string]int32) {
	o.Counters = v
}

// GetFileIds returns the FileIds field value
func (o *PublicImportMetadata) GetFileIds() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.FileIds
}

// GetFileIdsOk returns a tuple with the FileIds field value
// and a boolean to check if the value has been set.
func (o *PublicImportMetadata) GetFileIdsOk() ([]string, bool) {
	if o == nil {
		return nil, false
	}
	return o.FileIds, true
}

// SetFileIds sets field value
func (o *PublicImportMetadata) SetFileIds(v []string) {
	o.FileIds = v
}

// GetObjectLists returns the ObjectLists field value
func (o *PublicImportMetadata) GetObjectLists() []PublicObjectListRecord {
	if o == nil {
		var ret []PublicObjectListRecord
		return ret
	}

	return o.ObjectLists
}

// GetObjectListsOk returns a tuple with the ObjectLists field value
// and a boolean to check if the value has been set.
func (o *PublicImportMetadata) GetObjectListsOk() ([]PublicObjectListRecord, bool) {
	if o == nil {
		return nil, false
	}
	return o.ObjectLists, true
}

// SetObjectLists sets field value
func (o *PublicImportMetadata) SetObjectLists(v []PublicObjectListRecord) {
	o.ObjectLists = v
}

func (o PublicImportMetadata) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["counters"] = o.Counters
	}
	if true {
		toSerialize["fileIds"] = o.FileIds
	}
	if true {
		toSerialize["objectLists"] = o.ObjectLists
	}
	return json.Marshal(toSerialize)
}

type NullablePublicImportMetadata struct {
	value *PublicImportMetadata
	isSet bool
}

func (v NullablePublicImportMetadata) Get() *PublicImportMetadata {
	return v.value
}

func (v *NullablePublicImportMetadata) Set(val *PublicImportMetadata) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicImportMetadata) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicImportMetadata) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicImportMetadata(val *PublicImportMetadata) *NullablePublicImportMetadata {
	return &NullablePublicImportMetadata{value: val, isSet: true}
}

func (v NullablePublicImportMetadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicImportMetadata) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
