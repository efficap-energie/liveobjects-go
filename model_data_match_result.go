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

// DataMatchResult struct for DataMatchResult
type DataMatchResult struct {
	DataPredicateResult *bool `json:"dataPredicateResult,omitempty"`
	DataPredicateValid *bool `json:"dataPredicateValid,omitempty"`
	DataValid *bool `json:"dataValid,omitempty"`
}

// NewDataMatchResult instantiates a new DataMatchResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDataMatchResult() *DataMatchResult {
	this := DataMatchResult{}
	return &this
}

// NewDataMatchResultWithDefaults instantiates a new DataMatchResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDataMatchResultWithDefaults() *DataMatchResult {
	this := DataMatchResult{}
	return &this
}

// GetDataPredicateResult returns the DataPredicateResult field value if set, zero value otherwise.
func (o *DataMatchResult) GetDataPredicateResult() bool {
	if o == nil || o.DataPredicateResult == nil {
		var ret bool
		return ret
	}
	return *o.DataPredicateResult
}

// GetDataPredicateResultOk returns a tuple with the DataPredicateResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMatchResult) GetDataPredicateResultOk() (*bool, bool) {
	if o == nil || o.DataPredicateResult == nil {
		return nil, false
	}
	return o.DataPredicateResult, true
}

// HasDataPredicateResult returns a boolean if a field has been set.
func (o *DataMatchResult) HasDataPredicateResult() bool {
	if o != nil && o.DataPredicateResult != nil {
		return true
	}

	return false
}

// SetDataPredicateResult gets a reference to the given bool and assigns it to the DataPredicateResult field.
func (o *DataMatchResult) SetDataPredicateResult(v bool) {
	o.DataPredicateResult = &v
}

// GetDataPredicateValid returns the DataPredicateValid field value if set, zero value otherwise.
func (o *DataMatchResult) GetDataPredicateValid() bool {
	if o == nil || o.DataPredicateValid == nil {
		var ret bool
		return ret
	}
	return *o.DataPredicateValid
}

// GetDataPredicateValidOk returns a tuple with the DataPredicateValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMatchResult) GetDataPredicateValidOk() (*bool, bool) {
	if o == nil || o.DataPredicateValid == nil {
		return nil, false
	}
	return o.DataPredicateValid, true
}

// HasDataPredicateValid returns a boolean if a field has been set.
func (o *DataMatchResult) HasDataPredicateValid() bool {
	if o != nil && o.DataPredicateValid != nil {
		return true
	}

	return false
}

// SetDataPredicateValid gets a reference to the given bool and assigns it to the DataPredicateValid field.
func (o *DataMatchResult) SetDataPredicateValid(v bool) {
	o.DataPredicateValid = &v
}

// GetDataValid returns the DataValid field value if set, zero value otherwise.
func (o *DataMatchResult) GetDataValid() bool {
	if o == nil || o.DataValid == nil {
		var ret bool
		return ret
	}
	return *o.DataValid
}

// GetDataValidOk returns a tuple with the DataValid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DataMatchResult) GetDataValidOk() (*bool, bool) {
	if o == nil || o.DataValid == nil {
		return nil, false
	}
	return o.DataValid, true
}

// HasDataValid returns a boolean if a field has been set.
func (o *DataMatchResult) HasDataValid() bool {
	if o != nil && o.DataValid != nil {
		return true
	}

	return false
}

// SetDataValid gets a reference to the given bool and assigns it to the DataValid field.
func (o *DataMatchResult) SetDataValid(v bool) {
	o.DataValid = &v
}

func (o DataMatchResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DataPredicateResult != nil {
		toSerialize["dataPredicateResult"] = o.DataPredicateResult
	}
	if o.DataPredicateValid != nil {
		toSerialize["dataPredicateValid"] = o.DataPredicateValid
	}
	if o.DataValid != nil {
		toSerialize["dataValid"] = o.DataValid
	}
	return json.Marshal(toSerialize)
}

type NullableDataMatchResult struct {
	value *DataMatchResult
	isSet bool
}

func (v NullableDataMatchResult) Get() *DataMatchResult {
	return v.value
}

func (v *NullableDataMatchResult) Set(val *DataMatchResult) {
	v.value = val
	v.isSet = true
}

func (v NullableDataMatchResult) IsSet() bool {
	return v.isSet
}

func (v *NullableDataMatchResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDataMatchResult(val *DataMatchResult) *NullableDataMatchResult {
	return &NullableDataMatchResult{value: val, isSet: true}
}

func (v NullableDataMatchResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDataMatchResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


