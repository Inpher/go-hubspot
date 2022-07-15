/*
Accounting Extension

These APIs allow you to interact with HubSpot's Accounting Extension. It allows you to: * Specify the URLs that HubSpot will use when making webhook requests to your external accounting system. * Respond to webhook calls made to your external accounting system by HubSpot

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package accounting

import (
	"encoding/json"
)

// CustomerSearchResponseExternal A response to a search for customers.
type CustomerSearchResponseExternal struct {
	// Designates if the response is a success ('OK') or failure ('ERR').
	Result string `json:"@result"`
	// The list of customers that matched the search criteria.
	Customers []AccountingExtensionCustomer `json:"customers"`
}

// NewCustomerSearchResponseExternal instantiates a new CustomerSearchResponseExternal object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCustomerSearchResponseExternal(result string, customers []AccountingExtensionCustomer) *CustomerSearchResponseExternal {
	this := CustomerSearchResponseExternal{}
	this.Result = result
	this.Customers = customers
	return &this
}

// NewCustomerSearchResponseExternalWithDefaults instantiates a new CustomerSearchResponseExternal object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCustomerSearchResponseExternalWithDefaults() *CustomerSearchResponseExternal {
	this := CustomerSearchResponseExternal{}
	return &this
}

// GetResult returns the Result field value
func (o *CustomerSearchResponseExternal) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *CustomerSearchResponseExternal) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *CustomerSearchResponseExternal) SetResult(v string) {
	o.Result = v
}

// GetCustomers returns the Customers field value
func (o *CustomerSearchResponseExternal) GetCustomers() []AccountingExtensionCustomer {
	if o == nil {
		var ret []AccountingExtensionCustomer
		return ret
	}

	return o.Customers
}

// GetCustomersOk returns a tuple with the Customers field value
// and a boolean to check if the value has been set.
func (o *CustomerSearchResponseExternal) GetCustomersOk() ([]AccountingExtensionCustomer, bool) {
	if o == nil {
		return nil, false
	}
	return o.Customers, true
}

// SetCustomers sets field value
func (o *CustomerSearchResponseExternal) SetCustomers(v []AccountingExtensionCustomer) {
	o.Customers = v
}

func (o CustomerSearchResponseExternal) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["@result"] = o.Result
	}
	if true {
		toSerialize["customers"] = o.Customers
	}
	return json.Marshal(toSerialize)
}

type NullableCustomerSearchResponseExternal struct {
	value *CustomerSearchResponseExternal
	isSet bool
}

func (v NullableCustomerSearchResponseExternal) Get() *CustomerSearchResponseExternal {
	return v.value
}

func (v *NullableCustomerSearchResponseExternal) Set(val *CustomerSearchResponseExternal) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomerSearchResponseExternal) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomerSearchResponseExternal) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomerSearchResponseExternal(val *CustomerSearchResponseExternal) *NullableCustomerSearchResponseExternal {
	return &NullableCustomerSearchResponseExternal{value: val, isSet: true}
}

func (v NullableCustomerSearchResponseExternal) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomerSearchResponseExternal) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
