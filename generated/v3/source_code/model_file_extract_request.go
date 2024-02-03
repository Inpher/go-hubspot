/*
CMS Source Code

Endpoints for interacting with files in the CMS Developer File System. These files include HTML templates, CSS, JS, modules, and other assets which are used to create CMS content.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package source_code

import (
	"encoding/json"
)

// FileExtractRequest struct for FileExtractRequest
type FileExtractRequest struct {
	Path string `json:"path"`
}

// NewFileExtractRequest instantiates a new FileExtractRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFileExtractRequest(path string) *FileExtractRequest {
	this := FileExtractRequest{}
	this.Path = path
	return &this
}

// NewFileExtractRequestWithDefaults instantiates a new FileExtractRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFileExtractRequestWithDefaults() *FileExtractRequest {
	this := FileExtractRequest{}
	return &this
}

// GetPath returns the Path field value
func (o *FileExtractRequest) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *FileExtractRequest) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *FileExtractRequest) SetPath(v string) {
	o.Path = v
}

func (o FileExtractRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	return json.Marshal(toSerialize)
}

type NullableFileExtractRequest struct {
	value *FileExtractRequest
	isSet bool
}

func (v NullableFileExtractRequest) Get() *FileExtractRequest {
	return v.value
}

func (v *NullableFileExtractRequest) Set(val *FileExtractRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFileExtractRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFileExtractRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFileExtractRequest(val *FileExtractRequest) *NullableFileExtractRequest {
	return &NullableFileExtractRequest{value: val, isSet: true}
}

func (v NullableFileExtractRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFileExtractRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}