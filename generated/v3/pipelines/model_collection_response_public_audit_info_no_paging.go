/*
Pipelines

Pipelines represent distinct stages in a workflow, like closing a deal or servicing a support ticket. These endpoints provide access to read and modify pipelines in HubSpot. Pipelines support `deals` and `tickets` object types.  ## Pipeline ID validation  When calling endpoints that take pipelineId as a parameter, that ID must correspond to an existing, un-archived pipeline. Otherwise the request will fail with a `404 Not Found` response.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package pipelines

import (
	"encoding/json"
)

// CollectionResponsePublicAuditInfoNoPaging struct for CollectionResponsePublicAuditInfoNoPaging
type CollectionResponsePublicAuditInfoNoPaging struct {
	Results []PublicAuditInfo `json:"results"`
}

// NewCollectionResponsePublicAuditInfoNoPaging instantiates a new CollectionResponsePublicAuditInfoNoPaging object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponsePublicAuditInfoNoPaging(results []PublicAuditInfo) *CollectionResponsePublicAuditInfoNoPaging {
	this := CollectionResponsePublicAuditInfoNoPaging{}
	this.Results = results
	return &this
}

// NewCollectionResponsePublicAuditInfoNoPagingWithDefaults instantiates a new CollectionResponsePublicAuditInfoNoPaging object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponsePublicAuditInfoNoPagingWithDefaults() *CollectionResponsePublicAuditInfoNoPaging {
	this := CollectionResponsePublicAuditInfoNoPaging{}
	return &this
}

// GetResults returns the Results field value
func (o *CollectionResponsePublicAuditInfoNoPaging) GetResults() []PublicAuditInfo {
	if o == nil {
		var ret []PublicAuditInfo
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponsePublicAuditInfoNoPaging) GetResultsOk() ([]PublicAuditInfo, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponsePublicAuditInfoNoPaging) SetResults(v []PublicAuditInfo) {
	o.Results = v
}

func (o CollectionResponsePublicAuditInfoNoPaging) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["results"] = o.Results
	}
	return json.Marshal(toSerialize)
}

type NullableCollectionResponsePublicAuditInfoNoPaging struct {
	value *CollectionResponsePublicAuditInfoNoPaging
	isSet bool
}

func (v NullableCollectionResponsePublicAuditInfoNoPaging) Get() *CollectionResponsePublicAuditInfoNoPaging {
	return v.value
}

func (v *NullableCollectionResponsePublicAuditInfoNoPaging) Set(val *CollectionResponsePublicAuditInfoNoPaging) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponsePublicAuditInfoNoPaging) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponsePublicAuditInfoNoPaging) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponsePublicAuditInfoNoPaging(val *CollectionResponsePublicAuditInfoNoPaging) *NullableCollectionResponsePublicAuditInfoNoPaging {
	return &NullableCollectionResponsePublicAuditInfoNoPaging{value: val, isSet: true}
}

func (v NullableCollectionResponsePublicAuditInfoNoPaging) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponsePublicAuditInfoNoPaging) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}