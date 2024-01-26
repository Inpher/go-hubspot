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

// CollectionResponseWithTotalHubDbTableRowV3ForwardPaging struct for CollectionResponseWithTotalHubDbTableRowV3ForwardPaging
type CollectionResponseWithTotalHubDbTableRowV3ForwardPaging struct {
	//
	Total  int32          `json:"total"`
	Paging *ForwardPaging `json:"paging,omitempty"`
	//
	Results []HubDbTableRowV3 `json:"results"`
}

// NewCollectionResponseWithTotalHubDbTableRowV3ForwardPaging instantiates a new CollectionResponseWithTotalHubDbTableRowV3ForwardPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseWithTotalHubDbTableRowV3ForwardPaging(total int32, results []HubDbTableRowV3) *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging {
	this := CollectionResponseWithTotalHubDbTableRowV3ForwardPaging{}
	this.Total = total
	this.Results = results
	return &this
}

// NewCollectionResponseWithTotalHubDbTableRowV3ForwardPagingWithDefaults instantiates a new CollectionResponseWithTotalHubDbTableRowV3ForwardPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseWithTotalHubDbTableRowV3ForwardPagingWithDefaults() *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging {
	this := CollectionResponseWithTotalHubDbTableRowV3ForwardPaging{}
	return &this
}

// GetTotal returns the Total field value
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) SetTotal(v int32) {
	o.Total = v
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) GetPaging() ForwardPaging {
	if o == nil || o.Paging == nil {
		var ret ForwardPaging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) GetPagingOk() (*ForwardPaging, bool) {
	if o == nil || o.Paging == nil {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) HasPaging() bool {
	if o != nil && o.Paging != nil {
		return true
	}

	return false
}

// SetPaging gets a reference to the given ForwardPaging and assigns it to the Paging field.
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) SetPaging(v ForwardPaging) {
	o.Paging = &v
}

// GetResults returns the Results field value
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) GetResults() []HubDbTableRowV3 {
	if o == nil {
		var ret []HubDbTableRowV3
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) GetResultsOk() ([]HubDbTableRowV3, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) SetResults(v []HubDbTableRowV3) {
	o.Results = v
}

func (o CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["total"] = o.Total
	}
	if o.Paging != nil {
		toSerialize["paging"] = o.Paging
	}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging struct {
	value *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging
	isSet bool
}

func (v NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging) Get() *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging {
	return v.value
}

func (v *NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging) Set(val *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging(val *CollectionResponseWithTotalHubDbTableRowV3ForwardPaging) *NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging {
	return &NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging{value: val, isSet: true}
}

func (v NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseWithTotalHubDbTableRowV3ForwardPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
