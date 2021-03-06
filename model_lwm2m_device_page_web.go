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

// Lwm2mDevicePageWeb struct for Lwm2mDevicePageWeb
type Lwm2mDevicePageWeb struct {
	// number of the current page: starts at 0.
	Page int64 `json:"page"`
	// number of data per page (= maximum number of data in the associated list of data:the last page can have less data)
	Size int64 `json:"size"`
	// total count of data in the complete list.
	TotalCount int64 `json:"totalCount"`
	// list of data in this page.
	Data []LWM2MDevice `json:"data"`
}

// NewLwm2mDevicePageWeb instantiates a new Lwm2mDevicePageWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLwm2mDevicePageWeb(page int64, size int64, totalCount int64, data []LWM2MDevice, ) *Lwm2mDevicePageWeb {
	this := Lwm2mDevicePageWeb{}
	this.Page = page
	this.Size = size
	this.TotalCount = totalCount
	this.Data = data
	return &this
}

// NewLwm2mDevicePageWebWithDefaults instantiates a new Lwm2mDevicePageWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLwm2mDevicePageWebWithDefaults() *Lwm2mDevicePageWeb {
	this := Lwm2mDevicePageWeb{}
	return &this
}

// GetPage returns the Page field value
func (o *Lwm2mDevicePageWeb) GetPage() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *Lwm2mDevicePageWeb) GetPageOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *Lwm2mDevicePageWeb) SetPage(v int64) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *Lwm2mDevicePageWeb) GetSize() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *Lwm2mDevicePageWeb) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *Lwm2mDevicePageWeb) SetSize(v int64) {
	o.Size = v
}

// GetTotalCount returns the TotalCount field value
func (o *Lwm2mDevicePageWeb) GetTotalCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *Lwm2mDevicePageWeb) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *Lwm2mDevicePageWeb) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetData returns the Data field value
func (o *Lwm2mDevicePageWeb) GetData() []LWM2MDevice {
	if o == nil  {
		var ret []LWM2MDevice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *Lwm2mDevicePageWeb) GetDataOk() (*[]LWM2MDevice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *Lwm2mDevicePageWeb) SetData(v []LWM2MDevice) {
	o.Data = v
}

func (o Lwm2mDevicePageWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["page"] = o.Page
	}
	if true {
		toSerialize["size"] = o.Size
	}
	if true {
		toSerialize["totalCount"] = o.TotalCount
	}
	if true {
		toSerialize["data"] = o.Data
	}
	return json.Marshal(toSerialize)
}

type NullableLwm2mDevicePageWeb struct {
	value *Lwm2mDevicePageWeb
	isSet bool
}

func (v NullableLwm2mDevicePageWeb) Get() *Lwm2mDevicePageWeb {
	return v.value
}

func (v *NullableLwm2mDevicePageWeb) Set(val *Lwm2mDevicePageWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableLwm2mDevicePageWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableLwm2mDevicePageWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLwm2mDevicePageWeb(val *Lwm2mDevicePageWeb) *NullableLwm2mDevicePageWeb {
	return &NullableLwm2mDevicePageWeb{value: val, isSet: true}
}

func (v NullableLwm2mDevicePageWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLwm2mDevicePageWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


