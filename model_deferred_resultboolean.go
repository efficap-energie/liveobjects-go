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

// DeferredResultboolean struct for DeferredResultboolean
type DeferredResultboolean struct {
	Result *map[string]interface{} `json:"result,omitempty"`
	SetOrExpired *bool `json:"setOrExpired,omitempty"`
}

// NewDeferredResultboolean instantiates a new DeferredResultboolean object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeferredResultboolean() *DeferredResultboolean {
	this := DeferredResultboolean{}
	return &this
}

// NewDeferredResultbooleanWithDefaults instantiates a new DeferredResultboolean object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeferredResultbooleanWithDefaults() *DeferredResultboolean {
	this := DeferredResultboolean{}
	return &this
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *DeferredResultboolean) GetResult() map[string]interface{} {
	if o == nil || o.Result == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeferredResultboolean) GetResultOk() (*map[string]interface{}, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *DeferredResultboolean) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given map[string]interface{} and assigns it to the Result field.
func (o *DeferredResultboolean) SetResult(v map[string]interface{}) {
	o.Result = &v
}

// GetSetOrExpired returns the SetOrExpired field value if set, zero value otherwise.
func (o *DeferredResultboolean) GetSetOrExpired() bool {
	if o == nil || o.SetOrExpired == nil {
		var ret bool
		return ret
	}
	return *o.SetOrExpired
}

// GetSetOrExpiredOk returns a tuple with the SetOrExpired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeferredResultboolean) GetSetOrExpiredOk() (*bool, bool) {
	if o == nil || o.SetOrExpired == nil {
		return nil, false
	}
	return o.SetOrExpired, true
}

// HasSetOrExpired returns a boolean if a field has been set.
func (o *DeferredResultboolean) HasSetOrExpired() bool {
	if o != nil && o.SetOrExpired != nil {
		return true
	}

	return false
}

// SetSetOrExpired gets a reference to the given bool and assigns it to the SetOrExpired field.
func (o *DeferredResultboolean) SetSetOrExpired(v bool) {
	o.SetOrExpired = &v
}

func (o DeferredResultboolean) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.SetOrExpired != nil {
		toSerialize["setOrExpired"] = o.SetOrExpired
	}
	return json.Marshal(toSerialize)
}

type NullableDeferredResultboolean struct {
	value *DeferredResultboolean
	isSet bool
}

func (v NullableDeferredResultboolean) Get() *DeferredResultboolean {
	return v.value
}

func (v *NullableDeferredResultboolean) Set(val *DeferredResultboolean) {
	v.value = val
	v.isSet = true
}

func (v NullableDeferredResultboolean) IsSet() bool {
	return v.isSet
}

func (v *NullableDeferredResultboolean) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeferredResultboolean(val *DeferredResultboolean) *NullableDeferredResultboolean {
	return &NullableDeferredResultboolean{value: val, isSet: true}
}

func (v NullableDeferredResultboolean) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeferredResultboolean) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


