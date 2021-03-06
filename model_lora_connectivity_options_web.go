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

// LoraConnectivityOptionsWeb struct for LoraConnectivityOptionsWeb
type LoraConnectivityOptionsWeb struct {
	// Options for Ack uplink
	AckUl *bool `json:"ackUl,omitempty"`
	// Options for TDOA geolocation
	Tdoa *bool `json:"tdoa,omitempty"`
}

// NewLoraConnectivityOptionsWeb instantiates a new LoraConnectivityOptionsWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoraConnectivityOptionsWeb() *LoraConnectivityOptionsWeb {
	this := LoraConnectivityOptionsWeb{}
	return &this
}

// NewLoraConnectivityOptionsWebWithDefaults instantiates a new LoraConnectivityOptionsWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoraConnectivityOptionsWebWithDefaults() *LoraConnectivityOptionsWeb {
	this := LoraConnectivityOptionsWeb{}
	return &this
}

// GetAckUl returns the AckUl field value if set, zero value otherwise.
func (o *LoraConnectivityOptionsWeb) GetAckUl() bool {
	if o == nil || o.AckUl == nil {
		var ret bool
		return ret
	}
	return *o.AckUl
}

// GetAckUlOk returns a tuple with the AckUl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraConnectivityOptionsWeb) GetAckUlOk() (*bool, bool) {
	if o == nil || o.AckUl == nil {
		return nil, false
	}
	return o.AckUl, true
}

// HasAckUl returns a boolean if a field has been set.
func (o *LoraConnectivityOptionsWeb) HasAckUl() bool {
	if o != nil && o.AckUl != nil {
		return true
	}

	return false
}

// SetAckUl gets a reference to the given bool and assigns it to the AckUl field.
func (o *LoraConnectivityOptionsWeb) SetAckUl(v bool) {
	o.AckUl = &v
}

// GetTdoa returns the Tdoa field value if set, zero value otherwise.
func (o *LoraConnectivityOptionsWeb) GetTdoa() bool {
	if o == nil || o.Tdoa == nil {
		var ret bool
		return ret
	}
	return *o.Tdoa
}

// GetTdoaOk returns a tuple with the Tdoa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraConnectivityOptionsWeb) GetTdoaOk() (*bool, bool) {
	if o == nil || o.Tdoa == nil {
		return nil, false
	}
	return o.Tdoa, true
}

// HasTdoa returns a boolean if a field has been set.
func (o *LoraConnectivityOptionsWeb) HasTdoa() bool {
	if o != nil && o.Tdoa != nil {
		return true
	}

	return false
}

// SetTdoa gets a reference to the given bool and assigns it to the Tdoa field.
func (o *LoraConnectivityOptionsWeb) SetTdoa(v bool) {
	o.Tdoa = &v
}

func (o LoraConnectivityOptionsWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AckUl != nil {
		toSerialize["ackUl"] = o.AckUl
	}
	if o.Tdoa != nil {
		toSerialize["tdoa"] = o.Tdoa
	}
	return json.Marshal(toSerialize)
}

type NullableLoraConnectivityOptionsWeb struct {
	value *LoraConnectivityOptionsWeb
	isSet bool
}

func (v NullableLoraConnectivityOptionsWeb) Get() *LoraConnectivityOptionsWeb {
	return v.value
}

func (v *NullableLoraConnectivityOptionsWeb) Set(val *LoraConnectivityOptionsWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableLoraConnectivityOptionsWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableLoraConnectivityOptionsWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoraConnectivityOptionsWeb(val *LoraConnectivityOptionsWeb) *NullableLoraConnectivityOptionsWeb {
	return &NullableLoraConnectivityOptionsWeb{value: val, isSet: true}
}

func (v NullableLoraConnectivityOptionsWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoraConnectivityOptionsWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


