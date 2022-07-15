/*
CRM Objects

CRM objects such as companies, contacts, deals, line items, products, tickets, and quotes are standard objects in HubSpot’s CRM. These core building blocks support custom properties, store critical information, and play a central role in the HubSpot application.  ## Supported Object Types  This API provides access to collections of CRM objects, which return a map of property names to values. Each object type has its own set of default properties, which can be found by exploring the [CRM Object Properties API](https://developers.hubspot.com/docs/methods/crm-properties/crm-properties-overview).  |Object Type |Properties returned by default | |--|--| | `companies` | `name`, `domain` | | `contacts` | `firstname`, `lastname`, `email` | | `deals` | `dealname`, `amount`, `closedate`, `pipeline`, `dealstage` | | `products` | `name`, `description`, `price` | | `tickets` | `content`, `hs_pipeline`, `hs_pipeline_stage`, `hs_ticket_category`, `hs_ticket_priority`, `subject` |  Find a list of all properties for an object type using the [CRM Object Properties](https://developers.hubspot.com/docs/methods/crm-properties/get-properties) API. e.g. `GET https://api.hubapi.com/properties/v2/companies/properties`. Change the properties returned in the response using the `properties` array in the request body.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package objects

import (
	"encoding/json"
)

// ErrorCategory struct for ErrorCategory
type ErrorCategory struct {
	Name       string `json:"name"`
	HttpStatus string `json:"httpStatus"`
}

// NewErrorCategory instantiates a new ErrorCategory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewErrorCategory(name string, httpStatus string) *ErrorCategory {
	this := ErrorCategory{}
	this.Name = name
	this.HttpStatus = httpStatus
	return &this
}

// NewErrorCategoryWithDefaults instantiates a new ErrorCategory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewErrorCategoryWithDefaults() *ErrorCategory {
	this := ErrorCategory{}
	return &this
}

// GetName returns the Name field value
func (o *ErrorCategory) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *ErrorCategory) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *ErrorCategory) SetName(v string) {
	o.Name = v
}

// GetHttpStatus returns the HttpStatus field value
func (o *ErrorCategory) GetHttpStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.HttpStatus
}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
func (o *ErrorCategory) GetHttpStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ErrorCategory) SetHttpStatus(v string) {
	o.HttpStatus = v
}

func (o ErrorCategory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["httpStatus"] = o.HttpStatus
	}
	return json.Marshal(toSerialize)
}

type NullableErrorCategory struct {
	value *ErrorCategory
	isSet bool
}

func (v NullableErrorCategory) Get() *ErrorCategory {
	return v.value
}

func (v *NullableErrorCategory) Set(val *ErrorCategory) {
	v.value = val
	v.isSet = true
}

func (v NullableErrorCategory) IsSet() bool {
	return v.isSet
}

func (v *NullableErrorCategory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableErrorCategory(val *ErrorCategory) *NullableErrorCategory {
	return &NullableErrorCategory{value: val, isSet: true}
}

func (v NullableErrorCategory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableErrorCategory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
