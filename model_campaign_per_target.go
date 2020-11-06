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

// CampaignPerTarget struct for CampaignPerTarget
type CampaignPerTarget struct {
	Created string `json:"created"`
	// Device identifier
	Device string `json:"device"`
	Operations []CampaignOperationState `json:"operations"`
	// Global status of the sequence of operations
	Status string `json:"status"`
	Updated string `json:"updated"`
}

// NewCampaignPerTarget instantiates a new CampaignPerTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCampaignPerTarget(created string, device string, operations []CampaignOperationState, status string, updated string, ) *CampaignPerTarget {
	this := CampaignPerTarget{}
	this.Created = created
	this.Device = device
	this.Operations = operations
	this.Status = status
	this.Updated = updated
	return &this
}

// NewCampaignPerTargetWithDefaults instantiates a new CampaignPerTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCampaignPerTargetWithDefaults() *CampaignPerTarget {
	this := CampaignPerTarget{}
	return &this
}

// GetCreated returns the Created field value
func (o *CampaignPerTarget) GetCreated() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Created
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
func (o *CampaignPerTarget) GetCreatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Created, true
}

// SetCreated sets field value
func (o *CampaignPerTarget) SetCreated(v string) {
	o.Created = v
}

// GetDevice returns the Device field value
func (o *CampaignPerTarget) GetDevice() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Device
}

// GetDeviceOk returns a tuple with the Device field value
// and a boolean to check if the value has been set.
func (o *CampaignPerTarget) GetDeviceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Device, true
}

// SetDevice sets field value
func (o *CampaignPerTarget) SetDevice(v string) {
	o.Device = v
}

// GetOperations returns the Operations field value
func (o *CampaignPerTarget) GetOperations() []CampaignOperationState {
	if o == nil  {
		var ret []CampaignOperationState
		return ret
	}

	return o.Operations
}

// GetOperationsOk returns a tuple with the Operations field value
// and a boolean to check if the value has been set.
func (o *CampaignPerTarget) GetOperationsOk() (*[]CampaignOperationState, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Operations, true
}

// SetOperations sets field value
func (o *CampaignPerTarget) SetOperations(v []CampaignOperationState) {
	o.Operations = v
}

// GetStatus returns the Status field value
func (o *CampaignPerTarget) GetStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *CampaignPerTarget) GetStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *CampaignPerTarget) SetStatus(v string) {
	o.Status = v
}

// GetUpdated returns the Updated field value
func (o *CampaignPerTarget) GetUpdated() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Updated
}

// GetUpdatedOk returns a tuple with the Updated field value
// and a boolean to check if the value has been set.
func (o *CampaignPerTarget) GetUpdatedOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Updated, true
}

// SetUpdated sets field value
func (o *CampaignPerTarget) SetUpdated(v string) {
	o.Updated = v
}

func (o CampaignPerTarget) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["created"] = o.Created
	}
	if true {
		toSerialize["device"] = o.Device
	}
	if true {
		toSerialize["operations"] = o.Operations
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["updated"] = o.Updated
	}
	return json.Marshal(toSerialize)
}

type NullableCampaignPerTarget struct {
	value *CampaignPerTarget
	isSet bool
}

func (v NullableCampaignPerTarget) Get() *CampaignPerTarget {
	return v.value
}

func (v *NullableCampaignPerTarget) Set(val *CampaignPerTarget) {
	v.value = val
	v.isSet = true
}

func (v NullableCampaignPerTarget) IsSet() bool {
	return v.isSet
}

func (v *NullableCampaignPerTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCampaignPerTarget(val *CampaignPerTarget) *NullableCampaignPerTarget {
	return &NullableCampaignPerTarget{value: val, isSet: true}
}

func (v NullableCampaignPerTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCampaignPerTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


