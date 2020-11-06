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

// CommandAddRequest struct for CommandAddRequest
type CommandAddRequest struct {
	Request CommandRequest `json:"request"`
	Policy *CommandPolicy `json:"policy,omitempty"`
}

// NewCommandAddRequest instantiates a new CommandAddRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandAddRequest(request CommandRequest, ) *CommandAddRequest {
	this := CommandAddRequest{}
	this.Request = request
	return &this
}

// NewCommandAddRequestWithDefaults instantiates a new CommandAddRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandAddRequestWithDefaults() *CommandAddRequest {
	this := CommandAddRequest{}
	return &this
}

// GetRequest returns the Request field value
func (o *CommandAddRequest) GetRequest() CommandRequest {
	if o == nil  {
		var ret CommandRequest
		return ret
	}

	return o.Request
}

// GetRequestOk returns a tuple with the Request field value
// and a boolean to check if the value has been set.
func (o *CommandAddRequest) GetRequestOk() (*CommandRequest, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Request, true
}

// SetRequest sets field value
func (o *CommandAddRequest) SetRequest(v CommandRequest) {
	o.Request = v
}

// GetPolicy returns the Policy field value if set, zero value otherwise.
func (o *CommandAddRequest) GetPolicy() CommandPolicy {
	if o == nil || o.Policy == nil {
		var ret CommandPolicy
		return ret
	}
	return *o.Policy
}

// GetPolicyOk returns a tuple with the Policy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandAddRequest) GetPolicyOk() (*CommandPolicy, bool) {
	if o == nil || o.Policy == nil {
		return nil, false
	}
	return o.Policy, true
}

// HasPolicy returns a boolean if a field has been set.
func (o *CommandAddRequest) HasPolicy() bool {
	if o != nil && o.Policy != nil {
		return true
	}

	return false
}

// SetPolicy gets a reference to the given CommandPolicy and assigns it to the Policy field.
func (o *CommandAddRequest) SetPolicy(v CommandPolicy) {
	o.Policy = &v
}

func (o CommandAddRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["request"] = o.Request
	}
	if o.Policy != nil {
		toSerialize["policy"] = o.Policy
	}
	return json.Marshal(toSerialize)
}

type NullableCommandAddRequest struct {
	value *CommandAddRequest
	isSet bool
}

func (v NullableCommandAddRequest) Get() *CommandAddRequest {
	return v.value
}

func (v *NullableCommandAddRequest) Set(val *CommandAddRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandAddRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandAddRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandAddRequest(val *CommandAddRequest) *NullableCommandAddRequest {
	return &NullableCommandAddRequest{value: val, isSet: true}
}

func (v NullableCommandAddRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandAddRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


