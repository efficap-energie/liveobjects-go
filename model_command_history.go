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

// CommandHistory struct for CommandHistory
type CommandHistory struct {
	DeliveryStatus *string `json:"deliveryStatus,omitempty"`
	ErrorCode *string `json:"errorCode,omitempty"`
	NodeId *string `json:"nodeId,omitempty"`
	Status *string `json:"status,omitempty"`
	Timestamp *string `json:"timestamp,omitempty"`
}

// NewCommandHistory instantiates a new CommandHistory object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandHistory() *CommandHistory {
	this := CommandHistory{}
	return &this
}

// NewCommandHistoryWithDefaults instantiates a new CommandHistory object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandHistoryWithDefaults() *CommandHistory {
	this := CommandHistory{}
	return &this
}

// GetDeliveryStatus returns the DeliveryStatus field value if set, zero value otherwise.
func (o *CommandHistory) GetDeliveryStatus() string {
	if o == nil || o.DeliveryStatus == nil {
		var ret string
		return ret
	}
	return *o.DeliveryStatus
}

// GetDeliveryStatusOk returns a tuple with the DeliveryStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandHistory) GetDeliveryStatusOk() (*string, bool) {
	if o == nil || o.DeliveryStatus == nil {
		return nil, false
	}
	return o.DeliveryStatus, true
}

// HasDeliveryStatus returns a boolean if a field has been set.
func (o *CommandHistory) HasDeliveryStatus() bool {
	if o != nil && o.DeliveryStatus != nil {
		return true
	}

	return false
}

// SetDeliveryStatus gets a reference to the given string and assigns it to the DeliveryStatus field.
func (o *CommandHistory) SetDeliveryStatus(v string) {
	o.DeliveryStatus = &v
}

// GetErrorCode returns the ErrorCode field value if set, zero value otherwise.
func (o *CommandHistory) GetErrorCode() string {
	if o == nil || o.ErrorCode == nil {
		var ret string
		return ret
	}
	return *o.ErrorCode
}

// GetErrorCodeOk returns a tuple with the ErrorCode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandHistory) GetErrorCodeOk() (*string, bool) {
	if o == nil || o.ErrorCode == nil {
		return nil, false
	}
	return o.ErrorCode, true
}

// HasErrorCode returns a boolean if a field has been set.
func (o *CommandHistory) HasErrorCode() bool {
	if o != nil && o.ErrorCode != nil {
		return true
	}

	return false
}

// SetErrorCode gets a reference to the given string and assigns it to the ErrorCode field.
func (o *CommandHistory) SetErrorCode(v string) {
	o.ErrorCode = &v
}

// GetNodeId returns the NodeId field value if set, zero value otherwise.
func (o *CommandHistory) GetNodeId() string {
	if o == nil || o.NodeId == nil {
		var ret string
		return ret
	}
	return *o.NodeId
}

// GetNodeIdOk returns a tuple with the NodeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandHistory) GetNodeIdOk() (*string, bool) {
	if o == nil || o.NodeId == nil {
		return nil, false
	}
	return o.NodeId, true
}

// HasNodeId returns a boolean if a field has been set.
func (o *CommandHistory) HasNodeId() bool {
	if o != nil && o.NodeId != nil {
		return true
	}

	return false
}

// SetNodeId gets a reference to the given string and assigns it to the NodeId field.
func (o *CommandHistory) SetNodeId(v string) {
	o.NodeId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *CommandHistory) GetStatus() string {
	if o == nil || o.Status == nil {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandHistory) GetStatusOk() (*string, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *CommandHistory) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *CommandHistory) SetStatus(v string) {
	o.Status = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *CommandHistory) GetTimestamp() string {
	if o == nil || o.Timestamp == nil {
		var ret string
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandHistory) GetTimestampOk() (*string, bool) {
	if o == nil || o.Timestamp == nil {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *CommandHistory) HasTimestamp() bool {
	if o != nil && o.Timestamp != nil {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given string and assigns it to the Timestamp field.
func (o *CommandHistory) SetTimestamp(v string) {
	o.Timestamp = &v
}

func (o CommandHistory) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DeliveryStatus != nil {
		toSerialize["deliveryStatus"] = o.DeliveryStatus
	}
	if o.ErrorCode != nil {
		toSerialize["errorCode"] = o.ErrorCode
	}
	if o.NodeId != nil {
		toSerialize["nodeId"] = o.NodeId
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.Timestamp != nil {
		toSerialize["timestamp"] = o.Timestamp
	}
	return json.Marshal(toSerialize)
}

type NullableCommandHistory struct {
	value *CommandHistory
	isSet bool
}

func (v NullableCommandHistory) Get() *CommandHistory {
	return v.value
}

func (v *NullableCommandHistory) Set(val *CommandHistory) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandHistory) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandHistory) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandHistory(val *CommandHistory) *NullableCommandHistory {
	return &NullableCommandHistory{value: val, isSet: true}
}

func (v NullableCommandHistory) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandHistory) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


