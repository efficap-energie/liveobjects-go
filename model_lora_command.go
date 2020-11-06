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

// LoraCommand struct for LoraCommand
type LoraCommand struct {
	// Status of the command. SENT: The command was injected into LPWA network core. ERROR: The command could injected into LPWA network core.
	CommandStatus string `json:"commandStatus"`
	// Network ack confirmation
	Confirmed bool `json:"confirmed"`
	// Date/time of the command creation
	CreationTs *string `json:"creationTs,omitempty"`
	// Hexadecimal raw data of the command
	Data string `json:"data"`
	// Unique id of the command
	Id string `json:"id"`
	// Port of the device on which the command was sent (cf. LoRaWan)
	Port int32 `json:"port"`
	// Date/time of the last command modification
	UpdateTs *string `json:"updateTs,omitempty"`
}

// NewLoraCommand instantiates a new LoraCommand object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoraCommand(commandStatus string, confirmed bool, data string, id string, port int32, ) *LoraCommand {
	this := LoraCommand{}
	this.CommandStatus = commandStatus
	this.Confirmed = confirmed
	this.Data = data
	this.Id = id
	this.Port = port
	return &this
}

// NewLoraCommandWithDefaults instantiates a new LoraCommand object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoraCommandWithDefaults() *LoraCommand {
	this := LoraCommand{}
	return &this
}

// GetCommandStatus returns the CommandStatus field value
func (o *LoraCommand) GetCommandStatus() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.CommandStatus
}

// GetCommandStatusOk returns a tuple with the CommandStatus field value
// and a boolean to check if the value has been set.
func (o *LoraCommand) GetCommandStatusOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.CommandStatus, true
}

// SetCommandStatus sets field value
func (o *LoraCommand) SetCommandStatus(v string) {
	o.CommandStatus = v
}

// GetConfirmed returns the Confirmed field value
func (o *LoraCommand) GetConfirmed() bool {
	if o == nil  {
		var ret bool
		return ret
	}

	return o.Confirmed
}

// GetConfirmedOk returns a tuple with the Confirmed field value
// and a boolean to check if the value has been set.
func (o *LoraCommand) GetConfirmedOk() (*bool, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Confirmed, true
}

// SetConfirmed sets field value
func (o *LoraCommand) SetConfirmed(v bool) {
	o.Confirmed = v
}

// GetCreationTs returns the CreationTs field value if set, zero value otherwise.
func (o *LoraCommand) GetCreationTs() string {
	if o == nil || o.CreationTs == nil {
		var ret string
		return ret
	}
	return *o.CreationTs
}

// GetCreationTsOk returns a tuple with the CreationTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraCommand) GetCreationTsOk() (*string, bool) {
	if o == nil || o.CreationTs == nil {
		return nil, false
	}
	return o.CreationTs, true
}

// HasCreationTs returns a boolean if a field has been set.
func (o *LoraCommand) HasCreationTs() bool {
	if o != nil && o.CreationTs != nil {
		return true
	}

	return false
}

// SetCreationTs gets a reference to the given string and assigns it to the CreationTs field.
func (o *LoraCommand) SetCreationTs(v string) {
	o.CreationTs = &v
}

// GetData returns the Data field value
func (o *LoraCommand) GetData() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *LoraCommand) GetDataOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *LoraCommand) SetData(v string) {
	o.Data = v
}

// GetId returns the Id field value
func (o *LoraCommand) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *LoraCommand) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *LoraCommand) SetId(v string) {
	o.Id = v
}

// GetPort returns the Port field value
func (o *LoraCommand) GetPort() int32 {
	if o == nil  {
		var ret int32
		return ret
	}

	return o.Port
}

// GetPortOk returns a tuple with the Port field value
// and a boolean to check if the value has been set.
func (o *LoraCommand) GetPortOk() (*int32, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Port, true
}

// SetPort sets field value
func (o *LoraCommand) SetPort(v int32) {
	o.Port = v
}

// GetUpdateTs returns the UpdateTs field value if set, zero value otherwise.
func (o *LoraCommand) GetUpdateTs() string {
	if o == nil || o.UpdateTs == nil {
		var ret string
		return ret
	}
	return *o.UpdateTs
}

// GetUpdateTsOk returns a tuple with the UpdateTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraCommand) GetUpdateTsOk() (*string, bool) {
	if o == nil || o.UpdateTs == nil {
		return nil, false
	}
	return o.UpdateTs, true
}

// HasUpdateTs returns a boolean if a field has been set.
func (o *LoraCommand) HasUpdateTs() bool {
	if o != nil && o.UpdateTs != nil {
		return true
	}

	return false
}

// SetUpdateTs gets a reference to the given string and assigns it to the UpdateTs field.
func (o *LoraCommand) SetUpdateTs(v string) {
	o.UpdateTs = &v
}

func (o LoraCommand) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["commandStatus"] = o.CommandStatus
	}
	if true {
		toSerialize["confirmed"] = o.Confirmed
	}
	if o.CreationTs != nil {
		toSerialize["creationTs"] = o.CreationTs
	}
	if true {
		toSerialize["data"] = o.Data
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["port"] = o.Port
	}
	if o.UpdateTs != nil {
		toSerialize["updateTs"] = o.UpdateTs
	}
	return json.Marshal(toSerialize)
}

type NullableLoraCommand struct {
	value *LoraCommand
	isSet bool
}

func (v NullableLoraCommand) Get() *LoraCommand {
	return v.value
}

func (v *NullableLoraCommand) Set(val *LoraCommand) {
	v.value = val
	v.isSet = true
}

func (v NullableLoraCommand) IsSet() bool {
	return v.isSet
}

func (v *NullableLoraCommand) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoraCommand(val *LoraCommand) *NullableLoraCommand {
	return &NullableLoraCommand{value: val, isSet: true}
}

func (v NullableLoraCommand) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoraCommand) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


