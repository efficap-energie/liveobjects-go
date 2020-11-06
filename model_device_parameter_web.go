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

// DeviceParameterWeb struct for DeviceParameterWeb
type DeviceParameterWeb struct {
	Reported *DeviceParameterValueWeb `json:"reported,omitempty"`
	Requested *DeviceParameterValueWeb `json:"requested,omitempty"`
	// configuration parameter synchronization status: NONE, PENDING, OK or FAILED
	SyncStatus *string `json:"syncStatus,omitempty"`
}

// NewDeviceParameterWeb instantiates a new DeviceParameterWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceParameterWeb() *DeviceParameterWeb {
	this := DeviceParameterWeb{}
	return &this
}

// NewDeviceParameterWebWithDefaults instantiates a new DeviceParameterWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceParameterWebWithDefaults() *DeviceParameterWeb {
	this := DeviceParameterWeb{}
	return &this
}

// GetReported returns the Reported field value if set, zero value otherwise.
func (o *DeviceParameterWeb) GetReported() DeviceParameterValueWeb {
	if o == nil || o.Reported == nil {
		var ret DeviceParameterValueWeb
		return ret
	}
	return *o.Reported
}

// GetReportedOk returns a tuple with the Reported field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceParameterWeb) GetReportedOk() (*DeviceParameterValueWeb, bool) {
	if o == nil || o.Reported == nil {
		return nil, false
	}
	return o.Reported, true
}

// HasReported returns a boolean if a field has been set.
func (o *DeviceParameterWeb) HasReported() bool {
	if o != nil && o.Reported != nil {
		return true
	}

	return false
}

// SetReported gets a reference to the given DeviceParameterValueWeb and assigns it to the Reported field.
func (o *DeviceParameterWeb) SetReported(v DeviceParameterValueWeb) {
	o.Reported = &v
}

// GetRequested returns the Requested field value if set, zero value otherwise.
func (o *DeviceParameterWeb) GetRequested() DeviceParameterValueWeb {
	if o == nil || o.Requested == nil {
		var ret DeviceParameterValueWeb
		return ret
	}
	return *o.Requested
}

// GetRequestedOk returns a tuple with the Requested field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceParameterWeb) GetRequestedOk() (*DeviceParameterValueWeb, bool) {
	if o == nil || o.Requested == nil {
		return nil, false
	}
	return o.Requested, true
}

// HasRequested returns a boolean if a field has been set.
func (o *DeviceParameterWeb) HasRequested() bool {
	if o != nil && o.Requested != nil {
		return true
	}

	return false
}

// SetRequested gets a reference to the given DeviceParameterValueWeb and assigns it to the Requested field.
func (o *DeviceParameterWeb) SetRequested(v DeviceParameterValueWeb) {
	o.Requested = &v
}

// GetSyncStatus returns the SyncStatus field value if set, zero value otherwise.
func (o *DeviceParameterWeb) GetSyncStatus() string {
	if o == nil || o.SyncStatus == nil {
		var ret string
		return ret
	}
	return *o.SyncStatus
}

// GetSyncStatusOk returns a tuple with the SyncStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceParameterWeb) GetSyncStatusOk() (*string, bool) {
	if o == nil || o.SyncStatus == nil {
		return nil, false
	}
	return o.SyncStatus, true
}

// HasSyncStatus returns a boolean if a field has been set.
func (o *DeviceParameterWeb) HasSyncStatus() bool {
	if o != nil && o.SyncStatus != nil {
		return true
	}

	return false
}

// SetSyncStatus gets a reference to the given string and assigns it to the SyncStatus field.
func (o *DeviceParameterWeb) SetSyncStatus(v string) {
	o.SyncStatus = &v
}

func (o DeviceParameterWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Reported != nil {
		toSerialize["reported"] = o.Reported
	}
	if o.Requested != nil {
		toSerialize["requested"] = o.Requested
	}
	if o.SyncStatus != nil {
		toSerialize["syncStatus"] = o.SyncStatus
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceParameterWeb struct {
	value *DeviceParameterWeb
	isSet bool
}

func (v NullableDeviceParameterWeb) Get() *DeviceParameterWeb {
	return v.value
}

func (v *NullableDeviceParameterWeb) Set(val *DeviceParameterWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceParameterWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceParameterWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceParameterWeb(val *DeviceParameterWeb) *NullableDeviceParameterWeb {
	return &NullableDeviceParameterWeb{value: val, isSet: true}
}

func (v NullableDeviceParameterWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceParameterWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


