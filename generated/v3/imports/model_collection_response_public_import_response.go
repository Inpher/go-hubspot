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

// CollectionResponsePublicImportResponse struct for CollectionResponsePublicImportResponse
type CollectionResponsePublicImportResponse struct {
	Paging  *Paging                `json:"paging,omitempty"`
	Results []PublicImportResponse `json:"results"`
}

// NewCollectionResponsePublicImportResponse instantiates a new CollectionResponsePublicImportResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponsePublicImportResponse(results []PublicImportResponse) *CollectionResponsePublicImportResponse {
	this := CollectionResponsePublicImportResponse{}
	this.Results = results
	return &this
}

// NewCollectionResponsePublicImportResponseWithDefaults instantiates a new CollectionResponsePublicImportResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponsePublicImportResponseWithDefaults() *CollectionResponsePublicImportResponse {
	this := CollectionResponsePublicImportResponse{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponsePublicImportResponse) GetPaging() Paging {
	if o == nil || o.Paging == nil {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponsePublicImportResponse) GetPagingOk() (*Paging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponsePublicImportResponse) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *CollectionResponsePublicImportResponse) SetPaging(v Paging) {
	o.Paging = &v
}

// GetResults returns the Results field value
func (o *CollectionResponsePublicImportResponse) GetResults() []PublicImportResponse {
	if o == nil {
		var ret []PublicImportResponse
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponsePublicImportResponse) GetResultsOk() ([]PublicImportResponse, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponsePublicImportResponse) SetResults(v []PublicImportResponse) {
	o.Results = v
}

func (o CollectionResponsePublicImportResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponsePublicImportResponse struct {
	value *CollectionResponsePublicImportResponse
	isSet bool
}

func (v NullableCollectionResponsePublicImportResponse) Get() *CollectionResponsePublicImportResponse {
	return v.value
}

func (v *NullableCollectionResponsePublicImportResponse) Set(val *CollectionResponsePublicImportResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponsePublicImportResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponsePublicImportResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponsePublicImportResponse(val *CollectionResponsePublicImportResponse) *NullableCollectionResponsePublicImportResponse {
	return &NullableCollectionResponsePublicImportResponse{value: val, isSet: true}
}

func (v NullableCollectionResponsePublicImportResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponsePublicImportResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
