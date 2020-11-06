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

// EmailAction struct for EmailAction
type EmailAction struct {
	Cc *[]string `json:"cc,omitempty"`
	Cci *[]string `json:"cci,omitempty"`
	// A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the email content
	ContentTemplate string `json:"contentTemplate"`
	// A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the email subject
	SubjectTemplate string `json:"subjectTemplate"`
	To []string `json:"to"`
}

// NewEmailAction instantiates a new EmailAction object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEmailAction(contentTemplate string, subjectTemplate string, to []string, ) *EmailAction {
	this := EmailAction{}
	this.ContentTemplate = contentTemplate
	this.SubjectTemplate = subjectTemplate
	this.To = to
	return &this
}

// NewEmailActionWithDefaults instantiates a new EmailAction object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEmailActionWithDefaults() *EmailAction {
	this := EmailAction{}
	return &this
}

// GetCc returns the Cc field value if set, zero value otherwise.
func (o *EmailAction) GetCc() []string {
	if o == nil || o.Cc == nil {
		var ret []string
		return ret
	}
	return *o.Cc
}

// GetCcOk returns a tuple with the Cc field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAction) GetCcOk() (*[]string, bool) {
	if o == nil || o.Cc == nil {
		return nil, false
	}
	return o.Cc, true
}

// HasCc returns a boolean if a field has been set.
func (o *EmailAction) HasCc() bool {
	if o != nil && o.Cc != nil {
		return true
	}

	return false
}

// SetCc gets a reference to the given []string and assigns it to the Cc field.
func (o *EmailAction) SetCc(v []string) {
	o.Cc = &v
}

// GetCci returns the Cci field value if set, zero value otherwise.
func (o *EmailAction) GetCci() []string {
	if o == nil || o.Cci == nil {
		var ret []string
		return ret
	}
	return *o.Cci
}

// GetCciOk returns a tuple with the Cci field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EmailAction) GetCciOk() (*[]string, bool) {
	if o == nil || o.Cci == nil {
		return nil, false
	}
	return o.Cci, true
}

// HasCci returns a boolean if a field has been set.
func (o *EmailAction) HasCci() bool {
	if o != nil && o.Cci != nil {
		return true
	}

	return false
}

// SetCci gets a reference to the given []string and assigns it to the Cci field.
func (o *EmailAction) SetCci(v []string) {
	o.Cci = &v
}

// GetContentTemplate returns the ContentTemplate field value
func (o *EmailAction) GetContentTemplate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.ContentTemplate
}

// GetContentTemplateOk returns a tuple with the ContentTemplate field value
// and a boolean to check if the value has been set.
func (o *EmailAction) GetContentTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.ContentTemplate, true
}

// SetContentTemplate sets field value
func (o *EmailAction) SetContentTemplate(v string) {
	o.ContentTemplate = v
}

// GetSubjectTemplate returns the SubjectTemplate field value
func (o *EmailAction) GetSubjectTemplate() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.SubjectTemplate
}

// GetSubjectTemplateOk returns a tuple with the SubjectTemplate field value
// and a boolean to check if the value has been set.
func (o *EmailAction) GetSubjectTemplateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.SubjectTemplate, true
}

// SetSubjectTemplate sets field value
func (o *EmailAction) SetSubjectTemplate(v string) {
	o.SubjectTemplate = v
}

// GetTo returns the To field value
func (o *EmailAction) GetTo() []string {
	if o == nil  {
		var ret []string
		return ret
	}

	return o.To
}

// GetToOk returns a tuple with the To field value
// and a boolean to check if the value has been set.
func (o *EmailAction) GetToOk() (*[]string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.To, true
}

// SetTo sets field value
func (o *EmailAction) SetTo(v []string) {
	o.To = v
}

func (o EmailAction) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Cc != nil {
		toSerialize["cc"] = o.Cc
	}
	if o.Cci != nil {
		toSerialize["cci"] = o.Cci
	}
	if true {
		toSerialize["contentTemplate"] = o.ContentTemplate
	}
	if true {
		toSerialize["subjectTemplate"] = o.SubjectTemplate
	}
	if true {
		toSerialize["to"] = o.To
	}
	return json.Marshal(toSerialize)
}

type NullableEmailAction struct {
	value *EmailAction
	isSet bool
}

func (v NullableEmailAction) Get() *EmailAction {
	return v.value
}

func (v *NullableEmailAction) Set(val *EmailAction) {
	v.value = val
	v.isSet = true
}

func (v NullableEmailAction) IsSet() bool {
	return v.isSet
}

func (v *NullableEmailAction) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEmailAction(val *EmailAction) *NullableEmailAction {
	return &NullableEmailAction{value: val, isSet: true}
}

func (v NullableEmailAction) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEmailAction) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


