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

// PageableAssetResourceWeb struct for PageableAssetResourceWeb
type PageableAssetResourceWeb struct {
	// number of the current page: starts at 0.
	Page int64 `json:"page"`
	// number of data per page (= maximum number of data in the associated list of data:the last page can have less data)
	Size int64 `json:"size"`
	// total count of data in the complete list.
	TotalCount int64 `json:"totalCount"`
	// list of data in this page.
	Data []AssetResourceWeb `json:"data"`
}

// NewPageableAssetResourceWeb instantiates a new PageableAssetResourceWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableAssetResourceWeb(page int64, size int64, totalCount int64, data []AssetResourceWeb, ) *PageableAssetResourceWeb {
	this := PageableAssetResourceWeb{}
	this.Page = page
	this.Size = size
	this.TotalCount = totalCount
	this.Data = data
	return &this
}

// NewPageableAssetResourceWebWithDefaults instantiates a new PageableAssetResourceWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableAssetResourceWebWithDefaults() *PageableAssetResourceWeb {
	this := PageableAssetResourceWeb{}
	return &this
}

// GetPage returns the Page field value
func (o *PageableAssetResourceWeb) GetPage() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PageableAssetResourceWeb) GetPageOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PageableAssetResourceWeb) SetPage(v int64) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *PageableAssetResourceWeb) GetSize() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PageableAssetResourceWeb) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PageableAssetResourceWeb) SetSize(v int64) {
	o.Size = v
}

// GetTotalCount returns the TotalCount field value
func (o *PageableAssetResourceWeb) GetTotalCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PageableAssetResourceWeb) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PageableAssetResourceWeb) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetData returns the Data field value
func (o *PageableAssetResourceWeb) GetData() []AssetResourceWeb {
	if o == nil  {
		var ret []AssetResourceWeb
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PageableAssetResourceWeb) GetDataOk() (*[]AssetResourceWeb, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PageableAssetResourceWeb) SetData(v []AssetResourceWeb) {
	o.Data = v
}

func (o PageableAssetResourceWeb) MarshalJSON() ([]byte, error) {
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

type NullablePageableAssetResourceWeb struct {
	value *PageableAssetResourceWeb
	isSet bool
}

func (v NullablePageableAssetResourceWeb) Get() *PageableAssetResourceWeb {
	return v.value
}

func (v *NullablePageableAssetResourceWeb) Set(val *PageableAssetResourceWeb) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableAssetResourceWeb) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableAssetResourceWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableAssetResourceWeb(val *PageableAssetResourceWeb) *NullablePageableAssetResourceWeb {
	return &NullablePageableAssetResourceWeb{value: val, isSet: true}
}

func (v NullablePageableAssetResourceWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableAssetResourceWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


