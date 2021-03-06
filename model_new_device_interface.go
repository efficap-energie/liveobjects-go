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

// NewDeviceInterface Definition of a new device interface
type NewDeviceInterface struct {
	// Connector ID. A connector must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
	Connector *string `json:"connector,omitempty"`
	// Indicates whether the interface is enabled
	Enabled *bool `json:"enabled,omitempty"`
	// Base definition. Expected string (max 10000 characters)
	Definition *map[string]interface{} `json:"definition,omitempty"`
}

// NewNewDeviceInterface instantiates a new NewDeviceInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNewDeviceInterface() *NewDeviceInterface {
	this := NewDeviceInterface{}
	return &this
}

// NewNewDeviceInterfaceWithDefaults instantiates a new NewDeviceInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNewDeviceInterfaceWithDefaults() *NewDeviceInterface {
	this := NewDeviceInterface{}
	return &this
}

// GetConnector returns the Connector field value if set, zero value otherwise.
func (o *NewDeviceInterface) GetConnector() string {
	if o == nil || o.Connector == nil {
		var ret string
		return ret
	}
	return *o.Connector
}

// GetConnectorOk returns a tuple with the Connector field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDeviceInterface) GetConnectorOk() (*string, bool) {
	if o == nil || o.Connector == nil {
		return nil, false
	}
	return o.Connector, true
}

// HasConnector returns a boolean if a field has been set.
func (o *NewDeviceInterface) HasConnector() bool {
	if o != nil && o.Connector != nil {
		return true
	}

	return false
}

// SetConnector gets a reference to the given string and assigns it to the Connector field.
func (o *NewDeviceInterface) SetConnector(v string) {
	o.Connector = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *NewDeviceInterface) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDeviceInterface) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *NewDeviceInterface) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *NewDeviceInterface) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetDefinition returns the Definition field value if set, zero value otherwise.
func (o *NewDeviceInterface) GetDefinition() map[string]interface{} {
	if o == nil || o.Definition == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Definition
}

// GetDefinitionOk returns a tuple with the Definition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *NewDeviceInterface) GetDefinitionOk() (*map[string]interface{}, bool) {
	if o == nil || o.Definition == nil {
		return nil, false
	}
	return o.Definition, true
}

// HasDefinition returns a boolean if a field has been set.
func (o *NewDeviceInterface) HasDefinition() bool {
	if o != nil && o.Definition != nil {
		return true
	}

	return false
}

// SetDefinition gets a reference to the given map[string]interface{} and assigns it to the Definition field.
func (o *NewDeviceInterface) SetDefinition(v map[string]interface{}) {
	o.Definition = &v
}

func (o NewDeviceInterface) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connector != nil {
		toSerialize["connector"] = o.Connector
	}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.Definition != nil {
		toSerialize["definition"] = o.Definition
	}
	return json.Marshal(toSerialize)
}

type NullableNewDeviceInterface struct {
	value *NewDeviceInterface
	isSet bool
}

func (v NullableNewDeviceInterface) Get() *NewDeviceInterface {
	return v.value
}

func (v *NullableNewDeviceInterface) Set(val *NewDeviceInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableNewDeviceInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableNewDeviceInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNewDeviceInterface(val *NewDeviceInterface) *NullableNewDeviceInterface {
	return &NullableNewDeviceInterface{value: val, isSet: true}
}

func (v NullableNewDeviceInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNewDeviceInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


