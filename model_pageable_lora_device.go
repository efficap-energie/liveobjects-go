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

// PageableLoraDevice struct for PageableLoraDevice
type PageableLoraDevice struct {
	// number of the current page: starts at 0.
	Page int64 `json:"page"`
	// number of data per page (= maximum number of data in the associated list of data:the last page can have less data)
	Size int64 `json:"size"`
	// total count of data in the complete list.
	TotalCount int64 `json:"totalCount"`
	// list of data in this page.
	Data []LoraDevice `json:"data"`
}

// NewPageableLoraDevice instantiates a new PageableLoraDevice object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableLoraDevice(page int64, size int64, totalCount int64, data []LoraDevice, ) *PageableLoraDevice {
	this := PageableLoraDevice{}
	this.Page = page
	this.Size = size
	this.TotalCount = totalCount
	this.Data = data
	return &this
}

// NewPageableLoraDeviceWithDefaults instantiates a new PageableLoraDevice object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableLoraDeviceWithDefaults() *PageableLoraDevice {
	this := PageableLoraDevice{}
	return &this
}

// GetPage returns the Page field value
func (o *PageableLoraDevice) GetPage() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PageableLoraDevice) GetPageOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PageableLoraDevice) SetPage(v int64) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *PageableLoraDevice) GetSize() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PageableLoraDevice) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PageableLoraDevice) SetSize(v int64) {
	o.Size = v
}

// GetTotalCount returns the TotalCount field value
func (o *PageableLoraDevice) GetTotalCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PageableLoraDevice) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PageableLoraDevice) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetData returns the Data field value
func (o *PageableLoraDevice) GetData() []LoraDevice {
	if o == nil  {
		var ret []LoraDevice
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PageableLoraDevice) GetDataOk() (*[]LoraDevice, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PageableLoraDevice) SetData(v []LoraDevice) {
	o.Data = v
}

func (o PageableLoraDevice) MarshalJSON() ([]byte, error) {
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

type NullablePageableLoraDevice struct {
	value *PageableLoraDevice
	isSet bool
}

func (v NullablePageableLoraDevice) Get() *PageableLoraDevice {
	return v.value
}

func (v *NullablePageableLoraDevice) Set(val *PageableLoraDevice) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableLoraDevice) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableLoraDevice) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableLoraDevice(val *PageableLoraDevice) *NullablePageableLoraDevice {
	return &NullablePageableLoraDevice{value: val, isSet: true}
}

func (v NullablePageableLoraDevice) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableLoraDevice) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


