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

// LoraDeviceUpdateReqWeb struct for LoraDeviceUpdateReqWeb
type LoraDeviceUpdateReqWeb struct {
	// Application identifier (AppEUI)
	AppEUI *string `json:"appEUI,omitempty"`
	// Application Key (AppKey)
	AppKey *string `json:"appKey,omitempty"`
	ConnectivityOptions *LoraConnectivityOptionsWeb `json:"connectivityOptions,omitempty"`
	// Device connectivity plan.
	ConnectivityPlan *string `json:"connectivityPlan,omitempty"`
	// LoraDeviceUpdateReqWeb status
	DeviceStatus *string `json:"deviceStatus,omitempty"`
	// LoRa device encoding
	Encoding *string `json:"encoding,omitempty"`
	// LoRa device name
	Name *string `json:"name,omitempty"`
	// LoRa device profile
	Profile *string `json:"profile,omitempty"`
	// List of tags, used to tag uplink from this device
	Tags *[]string `json:"tags,omitempty"`
}

// NewLoraDeviceUpdateReqWeb instantiates a new LoraDeviceUpdateReqWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLoraDeviceUpdateReqWeb() *LoraDeviceUpdateReqWeb {
	this := LoraDeviceUpdateReqWeb{}
	return &this
}

// NewLoraDeviceUpdateReqWebWithDefaults instantiates a new LoraDeviceUpdateReqWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLoraDeviceUpdateReqWebWithDefaults() *LoraDeviceUpdateReqWeb {
	this := LoraDeviceUpdateReqWeb{}
	return &this
}

// GetAppEUI returns the AppEUI field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetAppEUI() string {
	if o == nil || o.AppEUI == nil {
		var ret string
		return ret
	}
	return *o.AppEUI
}

// GetAppEUIOk returns a tuple with the AppEUI field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetAppEUIOk() (*string, bool) {
	if o == nil || o.AppEUI == nil {
		return nil, false
	}
	return o.AppEUI, true
}

// HasAppEUI returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasAppEUI() bool {
	if o != nil && o.AppEUI != nil {
		return true
	}

	return false
}

// SetAppEUI gets a reference to the given string and assigns it to the AppEUI field.
func (o *LoraDeviceUpdateReqWeb) SetAppEUI(v string) {
	o.AppEUI = &v
}

// GetAppKey returns the AppKey field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetAppKey() string {
	if o == nil || o.AppKey == nil {
		var ret string
		return ret
	}
	return *o.AppKey
}

// GetAppKeyOk returns a tuple with the AppKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetAppKeyOk() (*string, bool) {
	if o == nil || o.AppKey == nil {
		return nil, false
	}
	return o.AppKey, true
}

// HasAppKey returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasAppKey() bool {
	if o != nil && o.AppKey != nil {
		return true
	}

	return false
}

// SetAppKey gets a reference to the given string and assigns it to the AppKey field.
func (o *LoraDeviceUpdateReqWeb) SetAppKey(v string) {
	o.AppKey = &v
}

// GetConnectivityOptions returns the ConnectivityOptions field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetConnectivityOptions() LoraConnectivityOptionsWeb {
	if o == nil || o.ConnectivityOptions == nil {
		var ret LoraConnectivityOptionsWeb
		return ret
	}
	return *o.ConnectivityOptions
}

// GetConnectivityOptionsOk returns a tuple with the ConnectivityOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetConnectivityOptionsOk() (*LoraConnectivityOptionsWeb, bool) {
	if o == nil || o.ConnectivityOptions == nil {
		return nil, false
	}
	return o.ConnectivityOptions, true
}

// HasConnectivityOptions returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasConnectivityOptions() bool {
	if o != nil && o.ConnectivityOptions != nil {
		return true
	}

	return false
}

// SetConnectivityOptions gets a reference to the given LoraConnectivityOptionsWeb and assigns it to the ConnectivityOptions field.
func (o *LoraDeviceUpdateReqWeb) SetConnectivityOptions(v LoraConnectivityOptionsWeb) {
	o.ConnectivityOptions = &v
}

// GetConnectivityPlan returns the ConnectivityPlan field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetConnectivityPlan() string {
	if o == nil || o.ConnectivityPlan == nil {
		var ret string
		return ret
	}
	return *o.ConnectivityPlan
}

// GetConnectivityPlanOk returns a tuple with the ConnectivityPlan field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetConnectivityPlanOk() (*string, bool) {
	if o == nil || o.ConnectivityPlan == nil {
		return nil, false
	}
	return o.ConnectivityPlan, true
}

// HasConnectivityPlan returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasConnectivityPlan() bool {
	if o != nil && o.ConnectivityPlan != nil {
		return true
	}

	return false
}

// SetConnectivityPlan gets a reference to the given string and assigns it to the ConnectivityPlan field.
func (o *LoraDeviceUpdateReqWeb) SetConnectivityPlan(v string) {
	o.ConnectivityPlan = &v
}

// GetDeviceStatus returns the DeviceStatus field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetDeviceStatus() string {
	if o == nil || o.DeviceStatus == nil {
		var ret string
		return ret
	}
	return *o.DeviceStatus
}

// GetDeviceStatusOk returns a tuple with the DeviceStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetDeviceStatusOk() (*string, bool) {
	if o == nil || o.DeviceStatus == nil {
		return nil, false
	}
	return o.DeviceStatus, true
}

// HasDeviceStatus returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasDeviceStatus() bool {
	if o != nil && o.DeviceStatus != nil {
		return true
	}

	return false
}

// SetDeviceStatus gets a reference to the given string and assigns it to the DeviceStatus field.
func (o *LoraDeviceUpdateReqWeb) SetDeviceStatus(v string) {
	o.DeviceStatus = &v
}

// GetEncoding returns the Encoding field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetEncoding() string {
	if o == nil || o.Encoding == nil {
		var ret string
		return ret
	}
	return *o.Encoding
}

// GetEncodingOk returns a tuple with the Encoding field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetEncodingOk() (*string, bool) {
	if o == nil || o.Encoding == nil {
		return nil, false
	}
	return o.Encoding, true
}

// HasEncoding returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasEncoding() bool {
	if o != nil && o.Encoding != nil {
		return true
	}

	return false
}

// SetEncoding gets a reference to the given string and assigns it to the Encoding field.
func (o *LoraDeviceUpdateReqWeb) SetEncoding(v string) {
	o.Encoding = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *LoraDeviceUpdateReqWeb) SetName(v string) {
	o.Name = &v
}

// GetProfile returns the Profile field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetProfile() string {
	if o == nil || o.Profile == nil {
		var ret string
		return ret
	}
	return *o.Profile
}

// GetProfileOk returns a tuple with the Profile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetProfileOk() (*string, bool) {
	if o == nil || o.Profile == nil {
		return nil, false
	}
	return o.Profile, true
}

// HasProfile returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasProfile() bool {
	if o != nil && o.Profile != nil {
		return true
	}

	return false
}

// SetProfile gets a reference to the given string and assigns it to the Profile field.
func (o *LoraDeviceUpdateReqWeb) SetProfile(v string) {
	o.Profile = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *LoraDeviceUpdateReqWeb) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *LoraDeviceUpdateReqWeb) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *LoraDeviceUpdateReqWeb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *LoraDeviceUpdateReqWeb) SetTags(v []string) {
	o.Tags = &v
}

func (o LoraDeviceUpdateReqWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AppEUI != nil {
		toSerialize["appEUI"] = o.AppEUI
	}
	if o.AppKey != nil {
		toSerialize["appKey"] = o.AppKey
	}
	if o.ConnectivityOptions != nil {
		toSerialize["connectivityOptions"] = o.ConnectivityOptions
	}
	if o.ConnectivityPlan != nil {
		toSerialize["connectivityPlan"] = o.ConnectivityPlan
	}
	if o.DeviceStatus != nil {
		toSerialize["deviceStatus"] = o.DeviceStatus
	}
	if o.Encoding != nil {
		toSerialize["encoding"] = o.Encoding
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Profile != nil {
		toSerialize["profile"] = o.Profile
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableLoraDeviceUpdateReqWeb struct {
	value *LoraDeviceUpdateReqWeb
	isSet bool
}

func (v NullableLoraDeviceUpdateReqWeb) Get() *LoraDeviceUpdateReqWeb {
	return v.value
}

func (v *NullableLoraDeviceUpdateReqWeb) Set(val *LoraDeviceUpdateReqWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableLoraDeviceUpdateReqWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableLoraDeviceUpdateReqWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLoraDeviceUpdateReqWeb(val *LoraDeviceUpdateReqWeb) *NullableLoraDeviceUpdateReqWeb {
	return &NullableLoraDeviceUpdateReqWeb{value: val, isSet: true}
}

func (v NullableLoraDeviceUpdateReqWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLoraDeviceUpdateReqWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


