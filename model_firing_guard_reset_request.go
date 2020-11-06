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

// FiringGuardResetRequest struct for FiringGuardResetRequest
type FiringGuardResetRequest struct {
	// firingRuleId linked to the FiringGuard to reset
	FiringRuleId string `json:"firingRuleId"`
	// criteria that should match the FiringGuard to reset
	SelectionCriteria *[]SelectionCriteria `json:"selectionCriteria,omitempty"`
}

// NewFiringGuardResetRequest instantiates a new FiringGuardResetRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFiringGuardResetRequest(firingRuleId string, ) *FiringGuardResetRequest {
	this := FiringGuardResetRequest{}
	this.FiringRuleId = firingRuleId
	return &this
}

// NewFiringGuardResetRequestWithDefaults instantiates a new FiringGuardResetRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFiringGuardResetRequestWithDefaults() *FiringGuardResetRequest {
	this := FiringGuardResetRequest{}
	return &this
}

// GetFiringRuleId returns the FiringRuleId field value
func (o *FiringGuardResetRequest) GetFiringRuleId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.FiringRuleId
}

// GetFiringRuleIdOk returns a tuple with the FiringRuleId field value
// and a boolean to check if the value has been set.
func (o *FiringGuardResetRequest) GetFiringRuleIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FiringRuleId, true
}

// SetFiringRuleId sets field value
func (o *FiringGuardResetRequest) SetFiringRuleId(v string) {
	o.FiringRuleId = v
}

// GetSelectionCriteria returns the SelectionCriteria field value if set, zero value otherwise.
func (o *FiringGuardResetRequest) GetSelectionCriteria() []SelectionCriteria {
	if o == nil || o.SelectionCriteria == nil {
		var ret []SelectionCriteria
		return ret
	}
	return *o.SelectionCriteria
}

// GetSelectionCriteriaOk returns a tuple with the SelectionCriteria field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FiringGuardResetRequest) GetSelectionCriteriaOk() (*[]SelectionCriteria, bool) {
	if o == nil || o.SelectionCriteria == nil {
		return nil, false
	}
	return o.SelectionCriteria, true
}

// HasSelectionCriteria returns a boolean if a field has been set.
func (o *FiringGuardResetRequest) HasSelectionCriteria() bool {
	if o != nil && o.SelectionCriteria != nil {
		return true
	}

	return false
}

// SetSelectionCriteria gets a reference to the given []SelectionCriteria and assigns it to the SelectionCriteria field.
func (o *FiringGuardResetRequest) SetSelectionCriteria(v []SelectionCriteria) {
	o.SelectionCriteria = &v
}

func (o FiringGuardResetRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["firingRuleId"] = o.FiringRuleId
	}
	if o.SelectionCriteria != nil {
		toSerialize["selectionCriteria"] = o.SelectionCriteria
	}
	return json.Marshal(toSerialize)
}

type NullableFiringGuardResetRequest struct {
	value *FiringGuardResetRequest
	isSet bool
}

func (v NullableFiringGuardResetRequest) Get() *FiringGuardResetRequest {
	return v.value
}

func (v *NullableFiringGuardResetRequest) Set(val *FiringGuardResetRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFiringGuardResetRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFiringGuardResetRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFiringGuardResetRequest(val *FiringGuardResetRequest) *NullableFiringGuardResetRequest {
	return &NullableFiringGuardResetRequest{value: val, isSet: true}
}

func (v NullableFiringGuardResetRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFiringGuardResetRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


