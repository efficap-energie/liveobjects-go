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

// DeviceCreatedFilter struct for DeviceCreatedFilter
type DeviceCreatedFilter struct {
	// list of filtered connectors.
	Connectors *[]string `json:"connectors,omitempty"`
	// list of filtered group paths. Null or empty to accept all group paths
	GroupPaths *[]GroupPath `json:"groupPaths,omitempty"`
	// list of groups of tags that should be contained in message tags. There is a match if at least one group of tags is a match. A group of tags is a match if the tags of the message contains all elements of this group.<br>e.g. [[\"ALERT\"]] will match all messages containing tag 'ALERT'.<br>e.g. [[\"HIGH\", \"ALERT\"],[\"PROD\"]] will match messages with either tag 'PROD' or both tags 'ALERT' and 'HIGH'.
	Tags *[][]string `json:"tags,omitempty"`
}

// NewDeviceCreatedFilter instantiates a new DeviceCreatedFilter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceCreatedFilter() *DeviceCreatedFilter {
	this := DeviceCreatedFilter{}
	return &this
}

// NewDeviceCreatedFilterWithDefaults instantiates a new DeviceCreatedFilter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceCreatedFilterWithDefaults() *DeviceCreatedFilter {
	this := DeviceCreatedFilter{}
	return &this
}

// GetConnectors returns the Connectors field value if set, zero value otherwise.
func (o *DeviceCreatedFilter) GetConnectors() []string {
	if o == nil || o.Connectors == nil {
		var ret []string
		return ret
	}
	return *o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedFilter) GetConnectorsOk() (*[]string, bool) {
	if o == nil || o.Connectors == nil {
		return nil, false
	}
	return o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *DeviceCreatedFilter) HasConnectors() bool {
	if o != nil && o.Connectors != nil {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given []string and assigns it to the Connectors field.
func (o *DeviceCreatedFilter) SetConnectors(v []string) {
	o.Connectors = &v
}

// GetGroupPaths returns the GroupPaths field value if set, zero value otherwise.
func (o *DeviceCreatedFilter) GetGroupPaths() []GroupPath {
	if o == nil || o.GroupPaths == nil {
		var ret []GroupPath
		return ret
	}
	return *o.GroupPaths
}

// GetGroupPathsOk returns a tuple with the GroupPaths field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedFilter) GetGroupPathsOk() (*[]GroupPath, bool) {
	if o == nil || o.GroupPaths == nil {
		return nil, false
	}
	return o.GroupPaths, true
}

// HasGroupPaths returns a boolean if a field has been set.
func (o *DeviceCreatedFilter) HasGroupPaths() bool {
	if o != nil && o.GroupPaths != nil {
		return true
	}

	return false
}

// SetGroupPaths gets a reference to the given []GroupPath and assigns it to the GroupPaths field.
func (o *DeviceCreatedFilter) SetGroupPaths(v []GroupPath) {
	o.GroupPaths = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *DeviceCreatedFilter) GetTags() [][]string {
	if o == nil || o.Tags == nil {
		var ret [][]string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceCreatedFilter) GetTagsOk() (*[][]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *DeviceCreatedFilter) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given [][]string and assigns it to the Tags field.
func (o *DeviceCreatedFilter) SetTags(v [][]string) {
	o.Tags = &v
}

func (o DeviceCreatedFilter) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connectors != nil {
		toSerialize["connectors"] = o.Connectors
	}
	if o.GroupPaths != nil {
		toSerialize["groupPaths"] = o.GroupPaths
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableDeviceCreatedFilter struct {
	value *DeviceCreatedFilter
	isSet bool
}

func (v NullableDeviceCreatedFilter) Get() *DeviceCreatedFilter {
	return v.value
}

func (v *NullableDeviceCreatedFilter) Set(val *DeviceCreatedFilter) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceCreatedFilter) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceCreatedFilter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceCreatedFilter(val *DeviceCreatedFilter) *NullableDeviceCreatedFilter {
	return &NullableDeviceCreatedFilter{value: val, isSet: true}
}

func (v NullableDeviceCreatedFilter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceCreatedFilter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


