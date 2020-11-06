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

// DeviceActivityTrigger struct for DeviceActivityTrigger
type DeviceActivityTrigger struct {
	Filter *DeviceActivityFilter `json:"filter,omitempty"`
	// requested object version
	Version int32 `json:"version"`
}

// NewDeviceActivityTrigger instantiates a new DeviceActivityTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceActivityTrigger(version int32, ) *DeviceActivityTrigger {
	this := DeviceActivityTrigger{}
	this.Version = version
	return &this
}

// NewDeviceActivityTriggerWithDefaults instantiates a new DeviceActivityTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceActivityTriggerWithDefaults() *DeviceActivityTrigger {
	this := DeviceActivityTrigger{}
	return &this
}

// GetFilter returns the Filter field value if set, zero value otherwise.
func (o *DeviceActivityTrigger) GetFilter() DeviceActivityFilter {
	if o == nil || o.Filter == nil {
		var ret DeviceActivityFilter
		return ret
	}
	return *o.Filter
}

// GetFilterOk returns a tuple with the Filter field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceActivityTrigger) GetFilterOk() (*DeviceActivityFilter, bool) {
	if o == nil || o.Filter == nil {
		return nil, false
	}
	return o.Filter, true
}

// HasFilter returns a boolean if a field has been set.
func (o *DeviceActivityTrigger) HasFilter() bool {
	if o != nil && o.Filter != nil {
		return true
	}

	return false
}

// SetFilter gets a reference to the given DeviceActivityFilter and assigns it to the Filter field.
func (o *DeviceActivityTrigger) SetFilter(v DeviceActivityFilter) {
	o.Filter = &v
}

// GetVersion returns the Version field value
func (o *DeviceActivityTrigger) GetVersion() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Version
}

// GetVersionOk returns a tuple with the Version field value
// and a boolean to check if the value has been set.
func (o *DeviceActivityTrigger) GetVersionOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Version, true
}

// SetVersion sets field value
func (o *DeviceActivityTrigger) SetVersion(v int32) {
	o.Version = v
}

func (o DeviceActivityTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Filter != nil {
		toSerialize["filter"] = o.Filter
	}
	if true {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceActivityTrigger struct {
	value *DeviceActivityTrigger
	isSet bool
}

func (v NullableDeviceActivityTrigger) Get() *DeviceActivityTrigger {
	return v.value
}

func (v *NullableDeviceActivityTrigger) Set(val *DeviceActivityTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceActivityTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceActivityTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceActivityTrigger(val *DeviceActivityTrigger) *NullableDeviceActivityTrigger {
	return &NullableDeviceActivityTrigger{value: val, isSet: true}
}

func (v NullableDeviceActivityTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceActivityTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


