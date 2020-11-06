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

// PageableFifoTopic struct for PageableFifoTopic
type PageableFifoTopic struct {
	// number of the current page: starts at 0.
	Page int64 `json:"page"`
	// number of data per page (= maximum number of data in the associated list of data:the last page can have less data)
	Size int64 `json:"size"`
	// total count of data in the complete list.
	TotalCount int64 `json:"totalCount"`
	// list of data in this page.
	Data []FifoTopic `json:"data"`
}

// NewPageableFifoTopic instantiates a new PageableFifoTopic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageableFifoTopic(page int64, size int64, totalCount int64, data []FifoTopic, ) *PageableFifoTopic {
	this := PageableFifoTopic{}
	this.Page = page
	this.Size = size
	this.TotalCount = totalCount
	this.Data = data
	return &this
}

// NewPageableFifoTopicWithDefaults instantiates a new PageableFifoTopic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageableFifoTopicWithDefaults() *PageableFifoTopic {
	this := PageableFifoTopic{}
	return &this
}

// GetPage returns the Page field value
func (o *PageableFifoTopic) GetPage() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Page
}

// GetPageOk returns a tuple with the Page field value
// and a boolean to check if the value has been set.
func (o *PageableFifoTopic) GetPageOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Page, true
}

// SetPage sets field value
func (o *PageableFifoTopic) SetPage(v int64) {
	o.Page = v
}

// GetSize returns the Size field value
func (o *PageableFifoTopic) GetSize() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.Size
}

// GetSizeOk returns a tuple with the Size field value
// and a boolean to check if the value has been set.
func (o *PageableFifoTopic) GetSizeOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Size, true
}

// SetSize sets field value
func (o *PageableFifoTopic) SetSize(v int64) {
	o.Size = v
}

// GetTotalCount returns the TotalCount field value
func (o *PageableFifoTopic) GetTotalCount() int64 {
	if o == nil  {
		var ret int64
		return ret
	}

	return o.TotalCount
}

// GetTotalCountOk returns a tuple with the TotalCount field value
// and a boolean to check if the value has been set.
func (o *PageableFifoTopic) GetTotalCountOk() (*int64, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.TotalCount, true
}

// SetTotalCount sets field value
func (o *PageableFifoTopic) SetTotalCount(v int64) {
	o.TotalCount = v
}

// GetData returns the Data field value
func (o *PageableFifoTopic) GetData() []FifoTopic {
	if o == nil  {
		var ret []FifoTopic
		return ret
	}

	return o.Data
}

// GetDataOk returns a tuple with the Data field value
// and a boolean to check if the value has been set.
func (o *PageableFifoTopic) GetDataOk() (*[]FifoTopic, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Data, true
}

// SetData sets field value
func (o *PageableFifoTopic) SetData(v []FifoTopic) {
	o.Data = v
}

func (o PageableFifoTopic) MarshalJSON() ([]byte, error) {
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

type NullablePageableFifoTopic struct {
	value *PageableFifoTopic
	isSet bool
}

func (v NullablePageableFifoTopic) Get() *PageableFifoTopic {
	return v.value
}

func (v *NullablePageableFifoTopic) Set(val *PageableFifoTopic) {
	v.value = val
	v.isSet = true
}

func (v NullablePageableFifoTopic) IsSet() bool {
	return v.isSet
}

func (v *NullablePageableFifoTopic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageableFifoTopic(val *PageableFifoTopic) *NullablePageableFifoTopic {
	return &NullablePageableFifoTopic{value: val, isSet: true}
}

func (v NullablePageableFifoTopic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageableFifoTopic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


