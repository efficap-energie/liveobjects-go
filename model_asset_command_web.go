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

// AssetCommandWeb Asset command data
type AssetCommandWeb struct {
	// Data fields / arguments of the command sent to the device. Expected map of string,string (max 100 elements, key max 255 characters, value is limited)
	Data *map[string]string `json:"data,omitempty"`
	// 'event' of the command message to send to the device (usually used to indicate called function). The length is limited
	Event *string `json:"event,omitempty"`
	// payload of the command message to send to the device. The length is limited
	Payload *string `json:"payload,omitempty"`
}

// NewAssetCommandWeb instantiates a new AssetCommandWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCommandWeb() *AssetCommandWeb {
	this := AssetCommandWeb{}
	return &this
}

// NewAssetCommandWebWithDefaults instantiates a new AssetCommandWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCommandWebWithDefaults() *AssetCommandWeb {
	this := AssetCommandWeb{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *AssetCommandWeb) GetData() map[string]string {
	if o == nil || o.Data == nil {
		var ret map[string]string
		return ret
	}
	return *o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCommandWeb) GetDataOk() (*map[string]string, bool) {
	if o == nil || o.Data == nil {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *AssetCommandWeb) HasData() bool {
	if o != nil && o.Data != nil {
		return true
	}

	return false
}

// SetData gets a reference to the given map[string]string and assigns it to the Data field.
func (o *AssetCommandWeb) SetData(v map[string]string) {
	o.Data = &v
}

// GetEvent returns the Event field value if set, zero value otherwise.
func (o *AssetCommandWeb) GetEvent() string {
	if o == nil || o.Event == nil {
		var ret string
		return ret
	}
	return *o.Event
}

// GetEventOk returns a tuple with the Event field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCommandWeb) GetEventOk() (*string, bool) {
	if o == nil || o.Event == nil {
		return nil, false
	}
	return o.Event, true
}

// HasEvent returns a boolean if a field has been set.
func (o *AssetCommandWeb) HasEvent() bool {
	if o != nil && o.Event != nil {
		return true
	}

	return false
}

// SetEvent gets a reference to the given string and assigns it to the Event field.
func (o *AssetCommandWeb) SetEvent(v string) {
	o.Event = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *AssetCommandWeb) GetPayload() string {
	if o == nil || o.Payload == nil {
		var ret string
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCommandWeb) GetPayloadOk() (*string, bool) {
	if o == nil || o.Payload == nil {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *AssetCommandWeb) HasPayload() bool {
	if o != nil && o.Payload != nil {
		return true
	}

	return false
}

// SetPayload gets a reference to the given string and assigns it to the Payload field.
func (o *AssetCommandWeb) SetPayload(v string) {
	o.Payload = &v
}

func (o AssetCommandWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Data != nil {
		toSerialize["data"] = o.Data
	}
	if o.Event != nil {
		toSerialize["event"] = o.Event
	}
	if o.Payload != nil {
		toSerialize["payload"] = o.Payload
	}
	return json.Marshal(toSerialize)
}

type NullableAssetCommandWeb struct {
	value *AssetCommandWeb
	isSet bool
}

func (v NullableAssetCommandWeb) Get() *AssetCommandWeb {
	return v.value
}

func (v *NullableAssetCommandWeb) Set(val *AssetCommandWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCommandWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCommandWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCommandWeb(val *AssetCommandWeb) *NullableAssetCommandWeb {
	return &NullableAssetCommandWeb{value: val, isSet: true}
}

func (v NullableAssetCommandWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCommandWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


