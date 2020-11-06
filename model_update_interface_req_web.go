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

// UpdateInterfaceReqWeb struct for UpdateInterfaceReqWeb
type UpdateInterfaceReqWeb struct {
	// New device id this connector node must be associated with (or null)
	DeviceId *string `json:"deviceId,omitempty"`
	// New enabled state for this connector node (or null)
	Enabled *bool `json:"enabled,omitempty"`
	// New interface definition
	Definition *map[string]interface{} `json:"definition,omitempty"`
}

// NewUpdateInterfaceReqWeb instantiates a new UpdateInterfaceReqWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateInterfaceReqWeb() *UpdateInterfaceReqWeb {
	this := UpdateInterfaceReqWeb{}
	return &this
}

// NewUpdateInterfaceReqWebWithDefaults instantiates a new UpdateInterfaceReqWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateInterfaceReqWebWithDefaults() *UpdateInterfaceReqWeb {
	this := UpdateInterfaceReqWeb{}
	return &this
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *UpdateInterfaceReqWeb) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterfaceReqWeb) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *UpdateInterfaceReqWeb) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *UpdateInterfaceReqWeb) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UpdateInterfaceReqWeb) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterfaceReqWeb) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UpdateInterfaceReqWeb) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UpdateInterfaceReqWeb) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *UpdateInterfaceReqWeb) GetDefinition() map[string]interface{} {
	if o == nil || o.Definition == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateInterfaceReqWeb) GetDefinitionOk() (*map[string]interface{}, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *UpdateInterfaceReqWeb) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given map[string]interface{} and assigns it to the Definition field.
func (o *UpdateInterfaceReqWeb) SetDefinition(v map[string]interface{}) {
	o.Definition = &v
}

func (o UpdateInterfaceReqWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateInterfaceReqWeb struct {
	value *UpdateInterfaceReqWeb
	isSet bool
}

func (v NullableUpdateInterfaceReqWeb) Get() *UpdateInterfaceReqWeb {
	return v.value
}

func (v *NullableUpdateInterfaceReqWeb) Set(val *UpdateInterfaceReqWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateInterfaceReqWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateInterfaceReqWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateInterfaceReqWeb(val *UpdateInterfaceReqWeb) *NullableUpdateInterfaceReqWeb {
	return &NullableUpdateInterfaceReqWeb{value: val, isSet: true}
}

func (v NullableUpdateInterfaceReqWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateInterfaceReqWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


