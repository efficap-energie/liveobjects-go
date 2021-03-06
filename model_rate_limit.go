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

// RateLimit struct for RateLimit
type RateLimit struct {
	// http rate limit: maximum api calls allowed per time window
	HttpMaxCalls *int64 `json:"httpMaxCalls,omitempty"`
	// http rate limit: window size in seconds
	HttpWindowSize *int64 `json:"httpWindowSize,omitempty"`
	// mqtt bridge rate limit: maximum messages allowed per time window
	MqttBridgeMaxMessages *int64 `json:"mqttBridgeMaxMessages,omitempty"`
	// mqtt bridge rate limit: window size in seconds
	MqttBridgeWindowSize *int64 `json:"mqttBridgeWindowSize,omitempty"`
	// mqtt device rate limit: maximum messages allowed per time window
	MqttDeviceMaxMessages *int64 `json:"mqttDeviceMaxMessages,omitempty"`
	// mqtt device rate limit: window size in seconds
	MqttDeviceWindowSize *int64 `json:"mqttDeviceWindowSize,omitempty"`
}

// NewRateLimit instantiates a new RateLimit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRateLimit() *RateLimit {
	this := RateLimit{}
	return &this
}

// NewRateLimitWithDefaults instantiates a new RateLimit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRateLimitWithDefaults() *RateLimit {
	this := RateLimit{}
	return &this
}

// GetHttpMaxCalls returns the HttpMaxCalls field value if set, zero value otherwise.
func (o *RateLimit) GetHttpMaxCalls() int64 {
	if o == nil || o.HttpMaxCalls == nil {
		var ret int64
		return ret
	}
	return *o.HttpMaxCalls
}

// GetHttpMaxCallsOk returns a tuple with the HttpMaxCalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetHttpMaxCallsOk() (*int64, bool) {
	if o == nil || o.HttpMaxCalls == nil {
		return nil, false
	}
	return o.HttpMaxCalls, true
}

// HasHttpMaxCalls returns a boolean if a field has been set.
func (o *RateLimit) HasHttpMaxCalls() bool {
	if o != nil && o.HttpMaxCalls != nil {
		return true
	}

	return false
}

// SetHttpMaxCalls gets a reference to the given int64 and assigns it to the HttpMaxCalls field.
func (o *RateLimit) SetHttpMaxCalls(v int64) {
	o.HttpMaxCalls = &v
}

// GetHttpWindowSize returns the HttpWindowSize field value if set, zero value otherwise.
func (o *RateLimit) GetHttpWindowSize() int64 {
	if o == nil || o.HttpWindowSize == nil {
		var ret int64
		return ret
	}
	return *o.HttpWindowSize
}

// GetHttpWindowSizeOk returns a tuple with the HttpWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetHttpWindowSizeOk() (*int64, bool) {
	if o == nil || o.HttpWindowSize == nil {
		return nil, false
	}
	return o.HttpWindowSize, true
}

// HasHttpWindowSize returns a boolean if a field has been set.
func (o *RateLimit) HasHttpWindowSize() bool {
	if o != nil && o.HttpWindowSize != nil {
		return true
	}

	return false
}

// SetHttpWindowSize gets a reference to the given int64 and assigns it to the HttpWindowSize field.
func (o *RateLimit) SetHttpWindowSize(v int64) {
	o.HttpWindowSize = &v
}

// GetMqttBridgeMaxMessages returns the MqttBridgeMaxMessages field value if set, zero value otherwise.
func (o *RateLimit) GetMqttBridgeMaxMessages() int64 {
	if o == nil || o.MqttBridgeMaxMessages == nil {
		var ret int64
		return ret
	}
	return *o.MqttBridgeMaxMessages
}

// GetMqttBridgeMaxMessagesOk returns a tuple with the MqttBridgeMaxMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetMqttBridgeMaxMessagesOk() (*int64, bool) {
	if o == nil || o.MqttBridgeMaxMessages == nil {
		return nil, false
	}
	return o.MqttBridgeMaxMessages, true
}

// HasMqttBridgeMaxMessages returns a boolean if a field has been set.
func (o *RateLimit) HasMqttBridgeMaxMessages() bool {
	if o != nil && o.MqttBridgeMaxMessages != nil {
		return true
	}

	return false
}

// SetMqttBridgeMaxMessages gets a reference to the given int64 and assigns it to the MqttBridgeMaxMessages field.
func (o *RateLimit) SetMqttBridgeMaxMessages(v int64) {
	o.MqttBridgeMaxMessages = &v
}

// GetMqttBridgeWindowSize returns the MqttBridgeWindowSize field value if set, zero value otherwise.
func (o *RateLimit) GetMqttBridgeWindowSize() int64 {
	if o == nil || o.MqttBridgeWindowSize == nil {
		var ret int64
		return ret
	}
	return *o.MqttBridgeWindowSize
}

// GetMqttBridgeWindowSizeOk returns a tuple with the MqttBridgeWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetMqttBridgeWindowSizeOk() (*int64, bool) {
	if o == nil || o.MqttBridgeWindowSize == nil {
		return nil, false
	}
	return o.MqttBridgeWindowSize, true
}

// HasMqttBridgeWindowSize returns a boolean if a field has been set.
func (o *RateLimit) HasMqttBridgeWindowSize() bool {
	if o != nil && o.MqttBridgeWindowSize != nil {
		return true
	}

	return false
}

// SetMqttBridgeWindowSize gets a reference to the given int64 and assigns it to the MqttBridgeWindowSize field.
func (o *RateLimit) SetMqttBridgeWindowSize(v int64) {
	o.MqttBridgeWindowSize = &v
}

// GetMqttDeviceMaxMessages returns the MqttDeviceMaxMessages field value if set, zero value otherwise.
func (o *RateLimit) GetMqttDeviceMaxMessages() int64 {
	if o == nil || o.MqttDeviceMaxMessages == nil {
		var ret int64
		return ret
	}
	return *o.MqttDeviceMaxMessages
}

// GetMqttDeviceMaxMessagesOk returns a tuple with the MqttDeviceMaxMessages field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetMqttDeviceMaxMessagesOk() (*int64, bool) {
	if o == nil || o.MqttDeviceMaxMessages == nil {
		return nil, false
	}
	return o.MqttDeviceMaxMessages, true
}

// HasMqttDeviceMaxMessages returns a boolean if a field has been set.
func (o *RateLimit) HasMqttDeviceMaxMessages() bool {
	if o != nil && o.MqttDeviceMaxMessages != nil {
		return true
	}

	return false
}

// SetMqttDeviceMaxMessages gets a reference to the given int64 and assigns it to the MqttDeviceMaxMessages field.
func (o *RateLimit) SetMqttDeviceMaxMessages(v int64) {
	o.MqttDeviceMaxMessages = &v
}

// GetMqttDeviceWindowSize returns the MqttDeviceWindowSize field value if set, zero value otherwise.
func (o *RateLimit) GetMqttDeviceWindowSize() int64 {
	if o == nil || o.MqttDeviceWindowSize == nil {
		var ret int64
		return ret
	}
	return *o.MqttDeviceWindowSize
}

// GetMqttDeviceWindowSizeOk returns a tuple with the MqttDeviceWindowSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RateLimit) GetMqttDeviceWindowSizeOk() (*int64, bool) {
	if o == nil || o.MqttDeviceWindowSize == nil {
		return nil, false
	}
	return o.MqttDeviceWindowSize, true
}

// HasMqttDeviceWindowSize returns a boolean if a field has been set.
func (o *RateLimit) HasMqttDeviceWindowSize() bool {
	if o != nil && o.MqttDeviceWindowSize != nil {
		return true
	}

	return false
}

// SetMqttDeviceWindowSize gets a reference to the given int64 and assigns it to the MqttDeviceWindowSize field.
func (o *RateLimit) SetMqttDeviceWindowSize(v int64) {
	o.MqttDeviceWindowSize = &v
}

func (o RateLimit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpMaxCalls != nil {
		toSerialize["httpMaxCalls"] = o.HttpMaxCalls
	}
	if o.HttpWindowSize != nil {
		toSerialize["httpWindowSize"] = o.HttpWindowSize
	}
	if o.MqttBridgeMaxMessages != nil {
		toSerialize["mqttBridgeMaxMessages"] = o.MqttBridgeMaxMessages
	}
	if o.MqttBridgeWindowSize != nil {
		toSerialize["mqttBridgeWindowSize"] = o.MqttBridgeWindowSize
	}
	if o.MqttDeviceMaxMessages != nil {
		toSerialize["mqttDeviceMaxMessages"] = o.MqttDeviceMaxMessages
	}
	if o.MqttDeviceWindowSize != nil {
		toSerialize["mqttDeviceWindowSize"] = o.MqttDeviceWindowSize
	}
	return json.Marshal(toSerialize)
}

type NullableRateLimit struct {
	value *RateLimit
	isSet bool
}

func (v NullableRateLimit) Get() *RateLimit {
	return v.value
}

func (v *NullableRateLimit) Set(val *RateLimit) {
	v.value = val
	v.isSet = true
}

func (v NullableRateLimit) IsSet() bool {
	return v.isSet
}

func (v *NullableRateLimit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRateLimit(val *RateLimit) *NullableRateLimit {
	return &NullableRateLimit{value: val, isSet: true}
}

func (v NullableRateLimit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRateLimit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


