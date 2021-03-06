/*
 * Live Objects REST API Guide v2.13.3
 *
 * API description for Live Objects service
 *
 * API version: 2.13.3
 * Contact: liveobjects.support@orange.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package liveobjects

import (
	"encoding/json"
)

// LoraDeviceStatsWeb struct for LoraDeviceStatsWeb
type LoraDeviceStatsWeb struct {
	Activated *int64 `json:"activated,omitempty"`
	Deactivated *int64 `json:"deactivated,omitempty"`
}

// NewLoraDeviceStatsWeb instantiates a new LoraDeviceStatsWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoraDeviceStatsWeb() *LoraDeviceStatsWeb {
	this := LoraDeviceStatsWeb{}
	return &this
}

// NewLoraDeviceStatsWebWithDefaults instantiates a new LoraDeviceStatsWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoraDeviceStatsWebWithDefaults() *LoraDeviceStatsWeb {
	this := LoraDeviceStatsWeb{}
	return &this
}

// GetActivated returns the Activated field value if set, zero value otherwise.
func (o *LoraDeviceStatsWeb) GetActivated() int64 {
	if o == nil || o.Activated == nil {
		var ret int64
		return ret
	}
	return *o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceStatsWeb) GetActivatedOk() (*int64, bool) {
	if o == nil || o.Activated == nil {
		return nil, false
	}
	return o.Activated, true
}

// HasActivated returns a boolean if a field has been set.
func (o *LoraDeviceStatsWeb) HasActivated() bool {
	if o != nil && o.Activated != nil {
		return true
	}

	return false
}

// SetActivated gets a reference to the given int64 and assigns it to the Activated field.
func (o *LoraDeviceStatsWeb) SetActivated(v int64) {
	o.Activated = &v
}

// GetDeactivated returns the Deactivated field value if set, zero value otherwise.
func (o *LoraDeviceStatsWeb) GetDeactivated() int64 {
	if o == nil || o.Deactivated == nil {
		var ret int64
		return ret
	}
	return *o.Deactivated
}

// GetDeactivatedOk returns a tuple with the Deactivated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceStatsWeb) GetDeactivatedOk() (*int64, bool) {
	if o == nil || o.Deactivated == nil {
		return nil, false
	}
	return o.Deactivated, true
}

// HasDeactivated returns a boolean if a field has been set.
func (o *LoraDeviceStatsWeb) HasDeactivated() bool {
	if o != nil && o.Deactivated != nil {
		return true
	}

	return false
}

// SetDeactivated gets a reference to the given int64 and assigns it to the Deactivated field.
func (o *LoraDeviceStatsWeb) SetDeactivated(v int64) {
	o.Deactivated = &v
}

func (o LoraDeviceStatsWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Activated != nil {
		toSerialize["activated"] = o.Activated
	}
	if o.Deactivated != nil {
		toSerialize["deactivated"] = o.Deactivated
	}
	return json.Marshal(toSerialize)
}

type NullableLoraDeviceStatsWeb struct {
	value *LoraDeviceStatsWeb
	isSet bool
}

func (v NullableLoraDeviceStatsWeb) Get() *LoraDeviceStatsWeb {
	return v.value
}

func (v *NullableLoraDeviceStatsWeb) Set(val *LoraDeviceStatsWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableLoraDeviceStatsWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableLoraDeviceStatsWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoraDeviceStatsWeb(val *LoraDeviceStatsWeb) *NullableLoraDeviceStatsWeb {
	return &NullableLoraDeviceStatsWeb{value: val, isSet: true}
}

func (v NullableLoraDeviceStatsWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoraDeviceStatsWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


