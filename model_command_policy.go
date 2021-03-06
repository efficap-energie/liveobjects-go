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

// CommandPolicy struct for CommandPolicy
type CommandPolicy struct {
	// Expiration in seconds since command creation date. Default is no expiry. Min value is 5 seconds
	ExpirationInSeconds *int64 `json:"expirationInSeconds,omitempty"`
	// Ack timeout in seconds since command was sent. Default is no ack timeout. Min value is 10 seconds
	AckTimeoutInSeconds *int64 `json:"ackTimeoutInSeconds,omitempty"`
	// Ack mode for this command. NONE (or AUTO) ack means that the command is automatically acknowledged (set to 'PROCESSED' status) as the command is sent to the device. NETWORK ack requires a reception acknowledge. APPLICATIVE (or DEVICE) ack requires a command response from the device to change its status. Default ack mode is connectivity dependant.
	AckMode *string `json:"ackMode,omitempty"`
	// Number of attemps in case of ERROR. Default to 1
	Attempts *int32 `json:"attempts,omitempty"`
}

// NewCommandPolicy instantiates a new CommandPolicy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandPolicy() *CommandPolicy {
	this := CommandPolicy{}
	return &this
}

// NewCommandPolicyWithDefaults instantiates a new CommandPolicy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandPolicyWithDefaults() *CommandPolicy {
	this := CommandPolicy{}
	return &this
}

// GetExpirationInSeconds returns the ExpirationInSeconds field value if set, zero value otherwise.
func (o *CommandPolicy) GetExpirationInSeconds() int64 {
	if o == nil || o.ExpirationInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.ExpirationInSeconds
}

// GetExpirationInSecondsOk returns a tuple with the ExpirationInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandPolicy) GetExpirationInSecondsOk() (*int64, bool) {
	if o == nil || o.ExpirationInSeconds == nil {
		return nil, false
	}
	return o.ExpirationInSeconds, true
}

// HasExpirationInSeconds returns a boolean if a field has been set.
func (o *CommandPolicy) HasExpirationInSeconds() bool {
	if o != nil && o.ExpirationInSeconds != nil {
		return true
	}

	return false
}

// SetExpirationInSeconds gets a reference to the given int64 and assigns it to the ExpirationInSeconds field.
func (o *CommandPolicy) SetExpirationInSeconds(v int64) {
	o.ExpirationInSeconds = &v
}

// GetAckTimeoutInSeconds returns the AckTimeoutInSeconds field value if set, zero value otherwise.
func (o *CommandPolicy) GetAckTimeoutInSeconds() int64 {
	if o == nil || o.AckTimeoutInSeconds == nil {
		var ret int64
		return ret
	}
	return *o.AckTimeoutInSeconds
}

// GetAckTimeoutInSecondsOk returns a tuple with the AckTimeoutInSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandPolicy) GetAckTimeoutInSecondsOk() (*int64, bool) {
	if o == nil || o.AckTimeoutInSeconds == nil {
		return nil, false
	}
	return o.AckTimeoutInSeconds, true
}

// HasAckTimeoutInSeconds returns a boolean if a field has been set.
func (o *CommandPolicy) HasAckTimeoutInSeconds() bool {
	if o != nil && o.AckTimeoutInSeconds != nil {
		return true
	}

	return false
}

// SetAckTimeoutInSeconds gets a reference to the given int64 and assigns it to the AckTimeoutInSeconds field.
func (o *CommandPolicy) SetAckTimeoutInSeconds(v int64) {
	o.AckTimeoutInSeconds = &v
}

// GetAckMode returns the AckMode field value if set, zero value otherwise.
func (o *CommandPolicy) GetAckMode() string {
	if o == nil || o.AckMode == nil {
		var ret string
		return ret
	}
	return *o.AckMode
}

// GetAckModeOk returns a tuple with the AckMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandPolicy) GetAckModeOk() (*string, bool) {
	if o == nil || o.AckMode == nil {
		return nil, false
	}
	return o.AckMode, true
}

// HasAckMode returns a boolean if a field has been set.
func (o *CommandPolicy) HasAckMode() bool {
	if o != nil && o.AckMode != nil {
		return true
	}

	return false
}

// SetAckMode gets a reference to the given string and assigns it to the AckMode field.
func (o *CommandPolicy) SetAckMode(v string) {
	o.AckMode = &v
}

// GetAttempts returns the Attempts field value if set, zero value otherwise.
func (o *CommandPolicy) GetAttempts() int32 {
	if o == nil || o.Attempts == nil {
		var ret int32
		return ret
	}
	return *o.Attempts
}

// GetAttemptsOk returns a tuple with the Attempts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandPolicy) GetAttemptsOk() (*int32, bool) {
	if o == nil || o.Attempts == nil {
		return nil, false
	}
	return o.Attempts, true
}

// HasAttempts returns a boolean if a field has been set.
func (o *CommandPolicy) HasAttempts() bool {
	if o != nil && o.Attempts != nil {
		return true
	}

	return false
}

// SetAttempts gets a reference to the given int32 and assigns it to the Attempts field.
func (o *CommandPolicy) SetAttempts(v int32) {
	o.Attempts = &v
}

func (o CommandPolicy) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ExpirationInSeconds != nil {
		toSerialize["expirationInSeconds"] = o.ExpirationInSeconds
	}
	if o.AckTimeoutInSeconds != nil {
		toSerialize["ackTimeoutInSeconds"] = o.AckTimeoutInSeconds
	}
	if o.AckMode != nil {
		toSerialize["ackMode"] = o.AckMode
	}
	if o.Attempts != nil {
		toSerialize["attempts"] = o.Attempts
	}
	return json.Marshal(toSerialize)
}

type NullableCommandPolicy struct {
	value *CommandPolicy
	isSet bool
}

func (v NullableCommandPolicy) Get() *CommandPolicy {
	return v.value
}

func (v *NullableCommandPolicy) Set(val *CommandPolicy) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandPolicy) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandPolicy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandPolicy(val *CommandPolicy) *NullableCommandPolicy {
	return &NullableCommandPolicy{value: val, isSet: true}
}

func (v NullableCommandPolicy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandPolicy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


