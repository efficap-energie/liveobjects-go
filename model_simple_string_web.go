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

// SimpleStringWeb struct for SimpleStringWeb
type SimpleStringWeb struct {
	Content *string `json:"content,omitempty"`
}

// NewSimpleStringWeb instantiates a new SimpleStringWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSimpleStringWeb() *SimpleStringWeb {
	this := SimpleStringWeb{}
	return &this
}

// NewSimpleStringWebWithDefaults instantiates a new SimpleStringWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSimpleStringWebWithDefaults() *SimpleStringWeb {
	this := SimpleStringWeb{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *SimpleStringWeb) GetContent() string {
	if o == nil || o.Content == nil {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SimpleStringWeb) GetContentOk() (*string, bool) {
	if o == nil || o.Content == nil {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *SimpleStringWeb) HasContent() bool {
	if o != nil && o.Content != nil {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *SimpleStringWeb) SetContent(v string) {
	o.Content = &v
}

func (o SimpleStringWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Content != nil {
		toSerialize["content"] = o.Content
	}
	return json.Marshal(toSerialize)
}

type NullableSimpleStringWeb struct {
	value *SimpleStringWeb
	isSet bool
}

func (v NullableSimpleStringWeb) Get() *SimpleStringWeb {
	return v.value
}

func (v *NullableSimpleStringWeb) Set(val *SimpleStringWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableSimpleStringWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableSimpleStringWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSimpleStringWeb(val *SimpleStringWeb) *NullableSimpleStringWeb {
	return &NullableSimpleStringWeb{value: val, isSet: true}
}

func (v NullableSimpleStringWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSimpleStringWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


