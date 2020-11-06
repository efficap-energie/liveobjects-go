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

// ActivityStateMuteRequest request to mute or reset the nextAlarm of AvtivityStates targeting specific deviceId and/or ruleId
type ActivityStateMuteRequest struct {
	// id of the targeted activityRule. At least one of deviceId/activityRuleId must be set.
	ActivityRuleId *string `json:"activityRuleId,omitempty"`
	// id of the targeted device. At least one of deviceId/activityRuleId must be set.
	DeviceId *string `json:"deviceId,omitempty"`
	// set the order type : mute or reset (from now) the next alarm
	NextAlarmOrder string `json:"nextAlarmOrder"`
}

// NewActivityStateMuteRequest instantiates a new ActivityStateMuteRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivityStateMuteRequest(nextAlarmOrder string, ) *ActivityStateMuteRequest {
	this := ActivityStateMuteRequest{}
	this.NextAlarmOrder = nextAlarmOrder
	return &this
}

// NewActivityStateMuteRequestWithDefaults instantiates a new ActivityStateMuteRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivityStateMuteRequestWithDefaults() *ActivityStateMuteRequest {
	this := ActivityStateMuteRequest{}
	return &this
}

// GetActivityRuleId returns the ActivityRuleId field value if set, zero value otherwise.
func (o *ActivityStateMuteRequest) GetActivityRuleId() string {
	if o == nil || o.ActivityRuleId == nil {
		var ret string
		return ret
	}
	return *o.ActivityRuleId
}

// GetActivityRuleIdOk returns a tuple with the ActivityRuleId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityStateMuteRequest) GetActivityRuleIdOk() (*string, bool) {
	if o == nil || o.ActivityRuleId == nil {
		return nil, false
	}
	return o.ActivityRuleId, true
}

// HasActivityRuleId returns a boolean if a field has been set.
func (o *ActivityStateMuteRequest) HasActivityRuleId() bool {
	if o != nil && o.ActivityRuleId != nil {
		return true
	}

	return false
}

// SetActivityRuleId gets a reference to the given string and assigns it to the ActivityRuleId field.
func (o *ActivityStateMuteRequest) SetActivityRuleId(v string) {
	o.ActivityRuleId = &v
}

// GetDeviceId returns the DeviceId field value if set, zero value otherwise.
func (o *ActivityStateMuteRequest) GetDeviceId() string {
	if o == nil || o.DeviceId == nil {
		var ret string
		return ret
	}
	return *o.DeviceId
}

// GetDeviceIdOk returns a tuple with the DeviceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ActivityStateMuteRequest) GetDeviceIdOk() (*string, bool) {
	if o == nil || o.DeviceId == nil {
		return nil, false
	}
	return o.DeviceId, true
}

// HasDeviceId returns a boolean if a field has been set.
func (o *ActivityStateMuteRequest) HasDeviceId() bool {
	if o != nil && o.DeviceId != nil {
		return true
	}

	return false
}

// SetDeviceId gets a reference to the given string and assigns it to the DeviceId field.
func (o *ActivityStateMuteRequest) SetDeviceId(v string) {
	o.DeviceId = &v
}

// GetNextAlarmOrder returns the NextAlarmOrder field value
func (o *ActivityStateMuteRequest) GetNextAlarmOrder() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.NextAlarmOrder
}

// GetNextAlarmOrderOk returns a tuple with the NextAlarmOrder field value
// and a boolean to check if the value has been set.
func (o *ActivityStateMuteRequest) GetNextAlarmOrderOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.NextAlarmOrder, true
}

// SetNextAlarmOrder sets field value
func (o *ActivityStateMuteRequest) SetNextAlarmOrder(v string) {
	o.NextAlarmOrder = v
}

func (o ActivityStateMuteRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.ActivityRuleId != nil {
		toSerialize["activityRuleId"] = o.ActivityRuleId
	}
	if o.DeviceId != nil {
		toSerialize["deviceId"] = o.DeviceId
	}
	if true {
		toSerialize["nextAlarmOrder"] = o.NextAlarmOrder
	}
	return json.Marshal(toSerialize)
}

type NullableActivityStateMuteRequest struct {
	value *ActivityStateMuteRequest
	isSet bool
}

func (v NullableActivityStateMuteRequest) Get() *ActivityStateMuteRequest {
	return v.value
}

func (v *NullableActivityStateMuteRequest) Set(val *ActivityStateMuteRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActivityStateMuteRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActivityStateMuteRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivityStateMuteRequest(val *ActivityStateMuteRequest) *NullableActivityStateMuteRequest {
	return &NullableActivityStateMuteRequest{value: val, isSet: true}
}

func (v NullableActivityStateMuteRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivityStateMuteRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


