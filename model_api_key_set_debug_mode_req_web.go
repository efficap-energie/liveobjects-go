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

// ApiKeySetDebugModeReqWeb Set ApiKey debug mode request
type ApiKeySetDebugModeReqWeb struct {
	// The state of the debug mode
	Activated bool `json:"activated"`
	// The duration during which the debug mode will be activated (default = 15 min)
	DurationSeconds *int64 `json:"durationSeconds,omitempty"`
}

// NewApiKeySetDebugModeReqWeb instantiates a new ApiKeySetDebugModeReqWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApiKeySetDebugModeReqWeb(activated bool, ) *ApiKeySetDebugModeReqWeb {
	this := ApiKeySetDebugModeReqWeb{}
	this.Activated = activated
	return &this
}

// NewApiKeySetDebugModeReqWebWithDefaults instantiates a new ApiKeySetDebugModeReqWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApiKeySetDebugModeReqWebWithDefaults() *ApiKeySetDebugModeReqWeb {
	this := ApiKeySetDebugModeReqWeb{}
	return &this
}

// GetActivated returns the Activated field value
func (o *ApiKeySetDebugModeReqWeb) GetActivated() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Activated
}

// GetActivatedOk returns a tuple with the Activated field value
// and a boolean to check if the value has been set.
func (o *ApiKeySetDebugModeReqWeb) GetActivatedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Activated, true
}

// SetActivated sets field value
func (o *ApiKeySetDebugModeReqWeb) SetActivated(v bool) {
	o.Activated = v
}

// GetDurationSeconds returns the DurationSeconds field value if set, zero value otherwise.
func (o *ApiKeySetDebugModeReqWeb) GetDurationSeconds() int64 {
	if o == nil || o.DurationSeconds == nil {
		var ret int64
		return ret
	}
	return *o.DurationSeconds
}

// GetDurationSecondsOk returns a tuple with the DurationSeconds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApiKeySetDebugModeReqWeb) GetDurationSecondsOk() (*int64, bool) {
	if o == nil || o.DurationSeconds == nil {
		return nil, false
	}
	return o.DurationSeconds, true
}

// HasDurationSeconds returns a boolean if a field has been set.
func (o *ApiKeySetDebugModeReqWeb) HasDurationSeconds() bool {
	if o != nil && o.DurationSeconds != nil {
		return true
	}

	return false
}

// SetDurationSeconds gets a reference to the given int64 and assigns it to the DurationSeconds field.
func (o *ApiKeySetDebugModeReqWeb) SetDurationSeconds(v int64) {
	o.DurationSeconds = &v
}

func (o ApiKeySetDebugModeReqWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["activated"] = o.Activated
	}
	if o.DurationSeconds != nil {
		toSerialize["durationSeconds"] = o.DurationSeconds
	}
	return json.Marshal(toSerialize)
}

type NullableApiKeySetDebugModeReqWeb struct {
	value *ApiKeySetDebugModeReqWeb
	isSet bool
}

func (v NullableApiKeySetDebugModeReqWeb) Get() *ApiKeySetDebugModeReqWeb {
	return v.value
}

func (v *NullableApiKeySetDebugModeReqWeb) Set(val *ApiKeySetDebugModeReqWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableApiKeySetDebugModeReqWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableApiKeySetDebugModeReqWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApiKeySetDebugModeReqWeb(val *ApiKeySetDebugModeReqWeb) *NullableApiKeySetDebugModeReqWeb {
	return &NullableApiKeySetDebugModeReqWeb{value: val, isSet: true}
}

func (v NullableApiKeySetDebugModeReqWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApiKeySetDebugModeReqWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


