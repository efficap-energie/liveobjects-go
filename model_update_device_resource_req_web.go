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

// UpdateDeviceResourceReqWeb a resource update request
type UpdateDeviceResourceReqWeb struct {
	// the update metadata. Max number of metadata is 100. Metadata name max length is 255. Metadata value max length is 255.
	Metadata *map[string]string `json:"metadata,omitempty"`
	// requested resource version. Expected string (max 255 characters)
	Version *string `json:"version,omitempty"`
}

// NewUpdateDeviceResourceReqWeb instantiates a new UpdateDeviceResourceReqWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUpdateDeviceResourceReqWeb() *UpdateDeviceResourceReqWeb {
	this := UpdateDeviceResourceReqWeb{}
	return &this
}

// NewUpdateDeviceResourceReqWebWithDefaults instantiates a new UpdateDeviceResourceReqWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUpdateDeviceResourceReqWebWithDefaults() *UpdateDeviceResourceReqWeb {
	this := UpdateDeviceResourceReqWeb{}
	return &this
}

// GetMetadata returns the Metadata field value if set, zero value otherwise.
func (o *UpdateDeviceResourceReqWeb) GetMetadata() map[string]string {
	if o == nil || o.Metadata == nil {
		var ret map[string]string
		return ret
	}
	return *o.Metadata
}

// GetMetadataOk returns a tuple with the Metadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceResourceReqWeb) GetMetadataOk() (*map[string]string, bool) {
	if o == nil || o.Metadata == nil {
		return nil, false
	}
	return o.Metadata, true
}

// HasMetadata returns a boolean if a field has been set.
func (o *UpdateDeviceResourceReqWeb) HasMetadata() bool {
	if o != nil && o.Metadata != nil {
		return true
	}

	return false
}

// SetMetadata gets a reference to the given map[string]string and assigns it to the Metadata field.
func (o *UpdateDeviceResourceReqWeb) SetMetadata(v map[string]string) {
	o.Metadata = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *UpdateDeviceResourceReqWeb) GetVersion() string {
	if o == nil || o.Version == nil {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UpdateDeviceResourceReqWeb) GetVersionOk() (*string, bool) {
	if o == nil || o.Version == nil {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *UpdateDeviceResourceReqWeb) HasVersion() bool {
	if o != nil && o.Version != nil {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *UpdateDeviceResourceReqWeb) SetVersion(v string) {
	o.Version = &v
}

func (o UpdateDeviceResourceReqWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Metadata != nil {
		toSerialize["metadata"] = o.Metadata
	}
	if o.Version != nil {
		toSerialize["version"] = o.Version
	}
	return json.Marshal(toSerialize)
}

type NullableUpdateDeviceResourceReqWeb struct {
	value *UpdateDeviceResourceReqWeb
	isSet bool
}

func (v NullableUpdateDeviceResourceReqWeb) Get() *UpdateDeviceResourceReqWeb {
	return v.value
}

func (v *NullableUpdateDeviceResourceReqWeb) Set(val *UpdateDeviceResourceReqWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableUpdateDeviceResourceReqWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableUpdateDeviceResourceReqWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUpdateDeviceResourceReqWeb(val *UpdateDeviceResourceReqWeb) *NullableUpdateDeviceResourceReqWeb {
	return &NullableUpdateDeviceResourceReqWeb{value: val, isSet: true}
}

func (v NullableUpdateDeviceResourceReqWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUpdateDeviceResourceReqWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

