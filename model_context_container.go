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

// ContextContainer Hold tenant context data that can be used inside matching rules and state processing functions 
type ContextContainer struct {
	// context data.
	ContextData map[string]interface{} `json:"contextData"`
	// key for accessing context data. Should be null when used for PUT.
	ContextKey *string `json:"contextKey,omitempty"`
	// optional tags list for free use
	Tags *[]string `json:"tags,omitempty"`
}

// NewContextContainer instantiates a new ContextContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextContainer(contextData map[string]interface{}, ) *ContextContainer {
	this := ContextContainer{}
	this.ContextData = contextData
	return &this
}

// NewContextContainerWithDefaults instantiates a new ContextContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextContainerWithDefaults() *ContextContainer {
	this := ContextContainer{}
	return &this
}

// GetContextData returns the ContextData field value
func (o *ContextContainer) GetContextData() map[string]interface{} {
	if o == nil  {
		var ret map[string]interface{}
		return ret
	}

	return o.ContextData
}

// GetContextDataOk returns a tuple with the ContextData field value
// and a boolean to check if the value has been set.
func (o *ContextContainer) GetContextDataOk() (*map[string]interface{}, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContextData, true
}

// SetContextData sets field value
func (o *ContextContainer) SetContextData(v map[string]interface{}) {
	o.ContextData = v
}

// GetContextKey returns the ContextKey field value if set, zero value otherwise.
func (o *ContextContainer) GetContextKey() string {
	if o == nil || o.ContextKey == nil {
		var ret string
		return ret
	}
	return *o.ContextKey
}

// GetContextKeyOk returns a tuple with the ContextKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextContainer) GetContextKeyOk() (*string, bool) {
	if o == nil || o.ContextKey == nil {
		return nil, false
	}
	return o.ContextKey, true
}

// HasContextKey returns a boolean if a field has been set.
func (o *ContextContainer) HasContextKey() bool {
	if o != nil && o.ContextKey != nil {
		return true
	}

	return false
}

// SetContextKey gets a reference to the given string and assigns it to the ContextKey field.
func (o *ContextContainer) SetContextKey(v string) {
	o.ContextKey = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *ContextContainer) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextContainer) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *ContextContainer) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *ContextContainer) SetTags(v []string) {
	o.Tags = &v
}

func (o ContextContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contextData"] = o.ContextData
	}
	if o.ContextKey != nil {
		toSerialize["contextKey"] = o.ContextKey
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableContextContainer struct {
	value *ContextContainer
	isSet bool
}

func (v NullableContextContainer) Get() *ContextContainer {
	return v.value
}

func (v *NullableContextContainer) Set(val *ContextContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableContextContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableContextContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextContainer(val *ContextContainer) *NullableContextContainer {
	return &NullableContextContainer{value: val, isSet: true}
}

func (v NullableContextContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


