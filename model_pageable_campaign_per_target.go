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

// PageableCampaignPerTarget struct for PageableCampaignPerTarget
type PageableCampaignPerTarget struct {
	// number of the current page: starts at 0.
	Page int64 `json:"page"`
	// number of data per page (= maximum number of data in the associated list of data:the last page can have less data)
	Size int64 `json:"size"`
	// total count of data in the complete list.
	TotalCount int64 `json:"totalCount"`
	// list of data in this page.
	Data []CampaignPerTarget `json:"data"`
}

// NewPageableCampaignPerTarget instantiates a new PageableCampaignPerTarget object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableCampaignPerTarget(page int64, size int64, totalCount int64, data []CampaignPerTarget, ) *PageableCampaignPerTarget {
	this := PageableCampaignPerTarget{}
	this.Page = page
	this.Size = size
	this.TotalCount = totalCount
	this.Data = data
	return &this
}

// NewPageableCampaignPerTargetWithDefaults instantiates a new PageableCampaignPerTarget object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableCampaignPerTargetWithDefaults() *PageableCampaignPerTarget {
	this := PageableCampaignPerTarget{}
	return &this
}

// GetPage returns the Page field value
func (o *PageableCampaignPerTarget) GetPage() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PageableCampaignPerTarget) GetPageOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PageableCampaignPerTarget) SetPage(v int64) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *PageableCampaignPerTarget) GetSize() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PageableCampaignPerTarget) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PageableCampaignPerTarget) SetSize(v int64) {
	o.Size = v
}

// GetTotalCount returns the TotalCount field value
func (o *PageableCampaignPerTarget) GetTotalCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PageableCampaignPerTarget) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PageableCampaignPerTarget) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetData returns the Data field value
func (o *PageableCampaignPerTarget) GetData() []CampaignPerTarget {
	if o == nil  {
		var ret []CampaignPerTarget
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PageableCampaignPerTarget) GetDataOk() (*[]CampaignPerTarget, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PageableCampaignPerTarget) SetData(v []CampaignPerTarget) {
	o.Data = v
}

func (o PageableCampaignPerTarget) MarshalJSON() ([]byte, error) {
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

type NullablePageableCampaignPerTarget struct {
	value *PageableCampaignPerTarget
	isSet bool
}

func (v NullablePageableCampaignPerTarget) Get() *PageableCampaignPerTarget {
	return v.value
}

func (v *NullablePageableCampaignPerTarget) Set(val *PageableCampaignPerTarget) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableCampaignPerTarget) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableCampaignPerTarget) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableCampaignPerTarget(val *PageableCampaignPerTarget) *NullablePageableCampaignPerTarget {
	return &NullablePageableCampaignPerTarget{value: val, isSet: true}
}

func (v NullablePageableCampaignPerTarget) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableCampaignPerTarget) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


