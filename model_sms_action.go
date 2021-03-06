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

// SMSAction struct for SMSAction
type SMSAction struct {
	// A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the sms content
	ContentTemplate string `json:"contentTemplate"`
	// A list of valid msisdn
	DestinationPhoneNumbers []string `json:"destinationPhoneNumbers"`
}

// NewSMSAction instantiates a new SMSAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSMSAction(contentTemplate string, destinationPhoneNumbers []string, ) *SMSAction {
	this := SMSAction{}
	this.ContentTemplate = contentTemplate
	this.DestinationPhoneNumbers = destinationPhoneNumbers
	return &this
}

// NewSMSActionWithDefaults instantiates a new SMSAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSMSActionWithDefaults() *SMSAction {
	this := SMSAction{}
	return &this
}

// GetContentTemplate returns the ContentTemplate field value
func (o *SMSAction) GetContentTemplate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContentTemplate
}

// GetContentTemplateOk returns a tuple with the ContentTemplate field value
// and a boolean to check if the value has been set.
func (o *SMSAction) GetContentTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentTemplate, true
}

// SetContentTemplate sets field value
func (o *SMSAction) SetContentTemplate(v string) {
	o.ContentTemplate = v
}

// GetDestinationPhoneNumbers returns the DestinationPhoneNumbers field value
func (o *SMSAction) GetDestinationPhoneNumbers() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.DestinationPhoneNumbers
}

// GetDestinationPhoneNumbersOk returns a tuple with the DestinationPhoneNumbers field value
// and a boolean to check if the value has been set.
func (o *SMSAction) GetDestinationPhoneNumbersOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.DestinationPhoneNumbers, true
}

// SetDestinationPhoneNumbers sets field value
func (o *SMSAction) SetDestinationPhoneNumbers(v []string) {
	o.DestinationPhoneNumbers = v
}

func (o SMSAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["contentTemplate"] = o.ContentTemplate
	}
	if true {
		toSerialize["destinationPhoneNumbers"] = o.DestinationPhoneNumbers
	}
	return json.Marshal(toSerialize)
}

type NullableSMSAction struct {
	value *SMSAction
	isSet bool
}

func (v NullableSMSAction) Get() *SMSAction {
	return v.value
}

func (v *NullableSMSAction) Set(val *SMSAction) {
	v.value = val
	v.isSet = true
}

func (v NullableSMSAction) IsSet() bool {
	return v.isSet
}

func (v *NullableSMSAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSMSAction(val *SMSAction) *NullableSMSAction {
	return &NullableSMSAction{value: val, isSet: true}
}

func (v NullableSMSAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSMSAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


