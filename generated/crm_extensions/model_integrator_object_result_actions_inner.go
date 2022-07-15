/*
CRM cards

Allows an app to extend the CRM UI by surfacing custom cards in the sidebar of record pages. These cards are defined up-front as part of app configuration, then populated by external data fetch requests when the record page is accessed by a user.

API version: v3
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package crm_extensions

import (
	"encoding/json"
	"fmt"
)

// IntegratorObjectResultActionsInner - struct for IntegratorObjectResultActionsInner
type IntegratorObjectResultActionsInner struct {
	ActionHookActionBody *ActionHookActionBody
	IFrameActionBody     *IFrameActionBody
}

// ActionHookActionBodyAsIntegratorObjectResultActionsInner is a convenience function that returns ActionHookActionBody wrapped in IntegratorObjectResultActionsInner
func ActionHookActionBodyAsIntegratorObjectResultActionsInner(v *ActionHookActionBody) IntegratorObjectResultActionsInner {
	return IntegratorObjectResultActionsInner{
		ActionHookActionBody: v,
	}
}

// IFrameActionBodyAsIntegratorObjectResultActionsInner is a convenience function that returns IFrameActionBody wrapped in IntegratorObjectResultActionsInner
func IFrameActionBodyAsIntegratorObjectResultActionsInner(v *IFrameActionBody) IntegratorObjectResultActionsInner {
	return IntegratorObjectResultActionsInner{
		IFrameActionBody: v,
	}
}

// Unmarshal JSON data into one of the pointers in the struct
func (dst *IntegratorObjectResultActionsInner) UnmarshalJSON(data []byte) error {
	var err error
	match := 0
	// try to unmarshal data into ActionHookActionBody
	err = newStrictDecoder(data).Decode(&dst.ActionHookActionBody)
	if err == nil {
		jsonActionHookActionBody, _ := json.Marshal(dst.ActionHookActionBody)
		if string(jsonActionHookActionBody) == "{}" { // empty struct
			dst.ActionHookActionBody = nil
		} else {
			match++
		}
	} else {
		dst.ActionHookActionBody = nil
	}

	// try to unmarshal data into IFrameActionBody
	err = newStrictDecoder(data).Decode(&dst.IFrameActionBody)
	if err == nil {
		jsonIFrameActionBody, _ := json.Marshal(dst.IFrameActionBody)
		if string(jsonIFrameActionBody) == "{}" { // empty struct
			dst.IFrameActionBody = nil
		} else {
			match++
		}
	} else {
		dst.IFrameActionBody = nil
	}

	if match > 1 { // more than 1 match
		// reset to nil
		dst.ActionHookActionBody = nil
		dst.IFrameActionBody = nil

		return fmt.Errorf("Data matches more than one schema in oneOf(IntegratorObjectResultActionsInner)")
	} else if match == 1 {
		return nil // exactly one match
	} else { // no match
		return fmt.Errorf("Data failed to match schemas in oneOf(IntegratorObjectResultActionsInner)")
	}
}

// Marshal data from the first non-nil pointers in the struct to JSON
func (src IntegratorObjectResultActionsInner) MarshalJSON() ([]byte, error) {
	if src.ActionHookActionBody != nil {
		return json.Marshal(&src.ActionHookActionBody)
	}

	if src.IFrameActionBody != nil {
		return json.Marshal(&src.IFrameActionBody)
	}

	return nil, nil // no data in oneOf schemas
}

// Get the actual instance
func (obj *IntegratorObjectResultActionsInner) GetActualInstance() interface{} {
	if obj == nil {
		return nil
	}
	if obj.ActionHookActionBody != nil {
		return obj.ActionHookActionBody
	}

	if obj.IFrameActionBody != nil {
		return obj.IFrameActionBody
	}

	// all schemas are nil
	return nil
}

type NullableIntegratorObjectResultActionsInner struct {
	value *IntegratorObjectResultActionsInner
	isSet bool
}

func (v NullableIntegratorObjectResultActionsInner) Get() *IntegratorObjectResultActionsInner {
	return v.value
}

func (v *NullableIntegratorObjectResultActionsInner) Set(val *IntegratorObjectResultActionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableIntegratorObjectResultActionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableIntegratorObjectResultActionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIntegratorObjectResultActionsInner(val *IntegratorObjectResultActionsInner) *NullableIntegratorObjectResultActionsInner {
	return &NullableIntegratorObjectResultActionsInner{value: val, isSet: true}
}

func (v NullableIntegratorObjectResultActionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIntegratorObjectResultActionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}