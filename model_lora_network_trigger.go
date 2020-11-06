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

// LoraNetworkTrigger struct for LoraNetworkTrigger
type LoraNetworkTrigger struct {
	Filter *LoraNetworkFilter `json:"filter,omitempty"`
	// requested object version
	Version int32 `json:"version"`
}

// NewLoraNetworkTrigger instantiates a new LoraNetworkTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoraNetworkTrigger(version int32, ) *LoraNetworkTrigger {
	this := LoraNetworkTrigger{}
	this.Version = version
	return &this
}

// NewLoraNetworkTriggerWithDefaults instantiates a new LoraNetworkTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoraNetworkTriggerWithDefaults() *LoraNetworkTrigger {
	this := LoraNetworkTrigger{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *LoraNetworkTrigger) GetFilter() LoraNetworkFilter {
	if o == nil || o.Filter == nil {
		var ret LoraNetworkFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraNetworkTrigger) GetFilterOk() (*LoraNetworkFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *LoraNetworkTrigger) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given LoraNetworkFilter and assigns it to the Filter field.
func (o *LoraNetworkTrigger) SetFilter(v LoraNetworkFilter) {
	o.Filter = &v
}

// GetVersion returns the Version field value
func (o *LoraNetworkTrigger) GetVersion() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *LoraNetworkTrigger) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *LoraNetworkTrigger) SetVersion(v int32) {
	o.Version = v
}

func (o LoraNetworkTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableLoraNetworkTrigger struct {
	value *LoraNetworkTrigger
	isSet bool
}

func (v NullableLoraNetworkTrigger) Get() *LoraNetworkTrigger {
	return v.value
}

func (v *NullableLoraNetworkTrigger) Set(val *LoraNetworkTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableLoraNetworkTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableLoraNetworkTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoraNetworkTrigger(val *LoraNetworkTrigger) *NullableLoraNetworkTrigger {
	return &NullableLoraNetworkTrigger{value: val, isSet: true}
}

func (v NullableLoraNetworkTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoraNetworkTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


