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

// DeviceParameterValueWeb struct for DeviceParameterValueWeb
type DeviceParameterValueWeb struct {
	// configuration parameter value type (INT32, UINT32, FLOAT, STRING or BINARY)
	Type *string `json:"type,omitempty"`
	// configuration parameter value (number for INT32/UINT32 type, string for STRING type,float for FLOAT type, base64-encoded string for BINARY type) 
	Value *map[string]interface{} `json:"value,omitempty"`
	// configuration parameter value associated date/time (ISO 8601)
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewDeviceParameterValueWeb instantiates a new DeviceParameterValueWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceParameterValueWeb() *DeviceParameterValueWeb {
	this := DeviceParameterValueWeb{}
	return &this
}

// NewDeviceParameterValueWebWithDefaults instantiates a new DeviceParameterValueWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceParameterValueWebWithDefaults() *DeviceParameterValueWeb {
	this := DeviceParameterValueWeb{}
	return &this
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *DeviceParameterValueWeb) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceParameterValueWeb) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *DeviceParameterValueWeb) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *DeviceParameterValueWeb) SetType(v string) {
	o.Type = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *DeviceParameterValueWeb) GetValue() map[string]interface{} {
	if o == nil || o.Value == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceParameterValueWeb) GetValueOk() (*map[string]interface{}, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *DeviceParameterValueWeb) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given map[string]interface{} and assigns it to the Value field.
func (o *DeviceParameterValueWeb) SetValue(v map[string]interface{}) {
	o.Value = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *DeviceParameterValueWeb) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceParameterValueWeb) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *DeviceParameterValueWeb) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *DeviceParameterValueWeb) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o DeviceParameterValueWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceParameterValueWeb struct {
	value *DeviceParameterValueWeb
	isSet bool
}

func (v NullableDeviceParameterValueWeb) Get() *DeviceParameterValueWeb {
	return v.value
}

func (v *NullableDeviceParameterValueWeb) Set(val *DeviceParameterValueWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceParameterValueWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceParameterValueWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceParameterValueWeb(val *DeviceParameterValueWeb) *NullableDeviceParameterValueWeb {
	return &NullableDeviceParameterValueWeb{value: val, isSet: true}
}

func (v NullableDeviceParameterValueWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceParameterValueWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


