/*
Products

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package products

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CollectionResponseAssociatedId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CollectionResponseAssociatedId{}

// CollectionResponseAssociatedId struct for CollectionResponseAssociatedId
type CollectionResponseAssociatedId struct {
	Paging  *Paging        `json:"paging,omitempty"`
	Results []AssociatedId `json:"results"`
}

type _CollectionResponseAssociatedId CollectionResponseAssociatedId

// NewCollectionResponseAssociatedId instantiates a new CollectionResponseAssociatedId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollectionResponseAssociatedId(results []AssociatedId) *CollectionResponseAssociatedId {
	this := CollectionResponseAssociatedId{}
	this.Results = results
	return &this
}

// NewCollectionResponseAssociatedIdWithDefaults instantiates a new CollectionResponseAssociatedId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollectionResponseAssociatedIdWithDefaults() *CollectionResponseAssociatedId {
	this := CollectionResponseAssociatedId{}
	return &this
}

// GetPaging returns the Paging field value if set, zero value otherwise.
func (o *CollectionResponseAssociatedId) GetPaging() Paging {
	if o == nil || IsNil(o.Paging) {
		var ret Paging
		return ret
	}
	return *o.Paging
}

// GetPagingOk returns a tuple with the Paging field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CollectionResponseAssociatedId) GetPagingOk() (*Paging, bool) {
	if o == nil || IsNil(o.Paging) {
		return nil, false
	}
	return o.Paging, true
}

// HasPaging returns a boolean if a field has been set.
func (o *CollectionResponseAssociatedId) HasPaging() bool {
	if o != nil && !IsNil(o.Paging) {
		return true
	}

	return false
}

// SetPaging gets a reference to the given Paging and assigns it to the Paging field.
func (o *CollectionResponseAssociatedId) SetPaging(v Paging) {
	o.Paging = &v
}

// GetResults returns the Results field value
func (o *CollectionResponseAssociatedId) GetResults() []AssociatedId {
	if o == nil {
		var ret []AssociatedId
		return ret
	}

	return o.Results
}

// GetResultsOk returns a tuple with the Results field value
// and a boolean to check if the value has been set.
func (o *CollectionResponseAssociatedId) GetResultsOk() ([]AssociatedId, bool) {
	if o == nil {
		return nil, false
	}
	return o.Results, true
}

// SetResults sets field value
func (o *CollectionResponseAssociatedId) SetResults(v []AssociatedId) {
	o.Results = v
}

func (o CollectionResponseAssociatedId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CollectionResponseAssociatedId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Paging) {
		toSerialize["paging"] = o.Paging
	}
	toSerialize["results"] = o.Results
	return toSerialize, nil
}

func (o *CollectionResponseAssociatedId) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"results",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varCollectionResponseAssociatedId := _CollectionResponseAssociatedId{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCollectionResponseAssociatedId)

	if err != nil {
		return err
	}

	*o = CollectionResponseAssociatedId(varCollectionResponseAssociatedId)

	return err
}

type NullableCollectionResponseAssociatedId struct {
	value *CollectionResponseAssociatedId
	isSet bool
}

func (v NullableCollectionResponseAssociatedId) Get() *CollectionResponseAssociatedId {
	return v.value
}

func (v *NullableCollectionResponseAssociatedId) Set(val *CollectionResponseAssociatedId) {
	v.value = val
	v.isSet = true
}

func (v NullableCollectionResponseAssociatedId) IsSet() bool {
	return v.isSet
}

func (v *NullableCollectionResponseAssociatedId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollectionResponseAssociatedId(val *CollectionResponseAssociatedId) *NullableCollectionResponseAssociatedId {
	return &NullableCollectionResponseAssociatedId{value: val, isSet: true}
}

func (v NullableCollectionResponseAssociatedId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollectionResponseAssociatedId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
