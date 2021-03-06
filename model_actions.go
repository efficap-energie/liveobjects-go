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

// Actions struct for Actions
type Actions struct {
	// A collection of Azure Event Hubs actions
	AzureEventHubs *[]AzureEventHubsAction `json:"azureEventHubs,omitempty"`
	// A collection of email actions
	Emails *[]EmailAction `json:"emails,omitempty"`
	// A collection of fifo publish actions
	FifoPublish *[]FifoPublishAction `json:"fifoPublish,omitempty"`
	// A collection of webhook actions
	HttpPush *[]HttpPushAction `json:"httpPush,omitempty"`
	// A collection of sms actions
	Sms *[]SMSAction `json:"sms,omitempty"`
}

// NewActions instantiates a new Actions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActions() *Actions {
	this := Actions{}
	return &this
}

// NewActionsWithDefaults instantiates a new Actions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActionsWithDefaults() *Actions {
	this := Actions{}
	return &this
}

// GetAzureEventHubs returns the AzureEventHubs field value if set, zero value otherwise.
func (o *Actions) GetAzureEventHubs() []AzureEventHubsAction {
	if o == nil || o.AzureEventHubs == nil {
		var ret []AzureEventHubsAction
		return ret
	}
	return *o.AzureEventHubs
}

// GetAzureEventHubsOk returns a tuple with the AzureEventHubs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actions) GetAzureEventHubsOk() (*[]AzureEventHubsAction, bool) {
	if o == nil || o.AzureEventHubs == nil {
		return nil, false
	}
	return o.AzureEventHubs, true
}

// HasAzureEventHubs returns a boolean if a field has been set.
func (o *Actions) HasAzureEventHubs() bool {
	if o != nil && o.AzureEventHubs != nil {
		return true
	}

	return false
}

// SetAzureEventHubs gets a reference to the given []AzureEventHubsAction and assigns it to the AzureEventHubs field.
func (o *Actions) SetAzureEventHubs(v []AzureEventHubsAction) {
	o.AzureEventHubs = &v
}

// GetEmails returns the Emails field value if set, zero value otherwise.
func (o *Actions) GetEmails() []EmailAction {
	if o == nil || o.Emails == nil {
		var ret []EmailAction
		return ret
	}
	return *o.Emails
}

// GetEmailsOk returns a tuple with the Emails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actions) GetEmailsOk() (*[]EmailAction, bool) {
	if o == nil || o.Emails == nil {
		return nil, false
	}
	return o.Emails, true
}

// HasEmails returns a boolean if a field has been set.
func (o *Actions) HasEmails() bool {
	if o != nil && o.Emails != nil {
		return true
	}

	return false
}

// SetEmails gets a reference to the given []EmailAction and assigns it to the Emails field.
func (o *Actions) SetEmails(v []EmailAction) {
	o.Emails = &v
}

// GetFifoPublish returns the FifoPublish field value if set, zero value otherwise.
func (o *Actions) GetFifoPublish() []FifoPublishAction {
	if o == nil || o.FifoPublish == nil {
		var ret []FifoPublishAction
		return ret
	}
	return *o.FifoPublish
}

// GetFifoPublishOk returns a tuple with the FifoPublish field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actions) GetFifoPublishOk() (*[]FifoPublishAction, bool) {
	if o == nil || o.FifoPublish == nil {
		return nil, false
	}
	return o.FifoPublish, true
}

// HasFifoPublish returns a boolean if a field has been set.
func (o *Actions) HasFifoPublish() bool {
	if o != nil && o.FifoPublish != nil {
		return true
	}

	return false
}

// SetFifoPublish gets a reference to the given []FifoPublishAction and assigns it to the FifoPublish field.
func (o *Actions) SetFifoPublish(v []FifoPublishAction) {
	o.FifoPublish = &v
}

// GetHttpPush returns the HttpPush field value if set, zero value otherwise.
func (o *Actions) GetHttpPush() []HttpPushAction {
	if o == nil || o.HttpPush == nil {
		var ret []HttpPushAction
		return ret
	}
	return *o.HttpPush
}

// GetHttpPushOk returns a tuple with the HttpPush field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actions) GetHttpPushOk() (*[]HttpPushAction, bool) {
	if o == nil || o.HttpPush == nil {
		return nil, false
	}
	return o.HttpPush, true
}

// HasHttpPush returns a boolean if a field has been set.
func (o *Actions) HasHttpPush() bool {
	if o != nil && o.HttpPush != nil {
		return true
	}

	return false
}

// SetHttpPush gets a reference to the given []HttpPushAction and assigns it to the HttpPush field.
func (o *Actions) SetHttpPush(v []HttpPushAction) {
	o.HttpPush = &v
}

// GetSms returns the Sms field value if set, zero value otherwise.
func (o *Actions) GetSms() []SMSAction {
	if o == nil || o.Sms == nil {
		var ret []SMSAction
		return ret
	}
	return *o.Sms
}

// GetSmsOk returns a tuple with the Sms field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Actions) GetSmsOk() (*[]SMSAction, bool) {
	if o == nil || o.Sms == nil {
		return nil, false
	}
	return o.Sms, true
}

// HasSms returns a boolean if a field has been set.
func (o *Actions) HasSms() bool {
	if o != nil && o.Sms != nil {
		return true
	}

	return false
}

// SetSms gets a reference to the given []SMSAction and assigns it to the Sms field.
func (o *Actions) SetSms(v []SMSAction) {
	o.Sms = &v
}

func (o Actions) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AzureEventHubs != nil {
		toSerialize["azureEventHubs"] = o.AzureEventHubs
	}
	if o.Emails != nil {
		toSerialize["emails"] = o.Emails
	}
	if o.FifoPublish != nil {
		toSerialize["fifoPublish"] = o.FifoPublish
	}
	if o.HttpPush != nil {
		toSerialize["httpPush"] = o.HttpPush
	}
	if o.Sms != nil {
		toSerialize["sms"] = o.Sms
	}
	return json.Marshal(toSerialize)
}

type NullableActions struct {
	value *Actions
	isSet bool
}

func (v NullableActions) Get() *Actions {
	return v.value
}

func (v *NullableActions) Set(val *Actions) {
	v.value = val
	v.isSet = true
}

func (v NullableActions) IsSet() bool {
	return v.isSet
}

func (v *NullableActions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActions(val *Actions) *NullableActions {
	return &NullableActions{value: val, isSet: true}
}

func (v NullableActions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


