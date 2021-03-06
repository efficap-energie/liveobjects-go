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

// FifoPublishAction struct for FifoPublishAction
type FifoPublishAction struct {
	// Target fifo to publish into. Refer to Topics API for managing these fifos
	FifoName *string `json:"fifoName,omitempty"`
}

// NewFifoPublishAction instantiates a new FifoPublishAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFifoPublishAction() *FifoPublishAction {
	this := FifoPublishAction{}
	return &this
}

// NewFifoPublishActionWithDefaults instantiates a new FifoPublishAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFifoPublishActionWithDefaults() *FifoPublishAction {
	this := FifoPublishAction{}
	return &this
}

// GetFifoName returns the FifoName field value if set, zero value otherwise.
func (o *FifoPublishAction) GetFifoName() string {
	if o == nil || o.FifoName == nil {
		var ret string
		return ret
	}
	return *o.FifoName
}

// GetFifoNameOk returns a tuple with the FifoName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FifoPublishAction) GetFifoNameOk() (*string, bool) {
	if o == nil || o.FifoName == nil {
		return nil, false
	}
	return o.FifoName, true
}

// HasFifoName returns a boolean if a field has been set.
func (o *FifoPublishAction) HasFifoName() bool {
	if o != nil && o.FifoName != nil {
		return true
	}

	return false
}

// SetFifoName gets a reference to the given string and assigns it to the FifoName field.
func (o *FifoPublishAction) SetFifoName(v string) {
	o.FifoName = &v
}

func (o FifoPublishAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.FifoName != nil {
		toSerialize["fifoName"] = o.FifoName
	}
	return json.Marshal(toSerialize)
}

type NullableFifoPublishAction struct {
	value *FifoPublishAction
	isSet bool
}

func (v NullableFifoPublishAction) Get() *FifoPublishAction {
	return v.value
}

func (v *NullableFifoPublishAction) Set(val *FifoPublishAction) {
	v.value = val
	v.isSet = true
}

func (v NullableFifoPublishAction) IsSet() bool {
	return v.isSet
}

func (v *NullableFifoPublishAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFifoPublishAction(val *FifoPublishAction) *NullableFifoPublishAction {
	return &NullableFifoPublishAction{value: val, isSet: true}
}

func (v NullableFifoPublishAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFifoPublishAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


