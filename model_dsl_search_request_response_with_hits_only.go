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

// DslSearchRequestResponseWithHitsOnly struct for DslSearchRequestResponseWithHitsOnly
type DslSearchRequestResponseWithHitsOnly struct {
	Hits *[]SearchDataMessage `json:"hits,omitempty"`
	TookMillis *int64 `json:"tookMillis,omitempty"`
	TotalHits *int64 `json:"totalHits,omitempty"`
}

// NewDslSearchRequestResponseWithHitsOnly instantiates a new DslSearchRequestResponseWithHitsOnly object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDslSearchRequestResponseWithHitsOnly() *DslSearchRequestResponseWithHitsOnly {
	this := DslSearchRequestResponseWithHitsOnly{}
	return &this
}

// NewDslSearchRequestResponseWithHitsOnlyWithDefaults instantiates a new DslSearchRequestResponseWithHitsOnly object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDslSearchRequestResponseWithHitsOnlyWithDefaults() *DslSearchRequestResponseWithHitsOnly {
	this := DslSearchRequestResponseWithHitsOnly{}
	return &this
}

// GetHits returns the Hits field value if set, zero value otherwise.
func (o *DslSearchRequestResponseWithHitsOnly) GetHits() []SearchDataMessage {
	if o == nil || o.Hits == nil {
		var ret []SearchDataMessage
		return ret
	}
	return *o.Hits
}

// GetHitsOk returns a tuple with the Hits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DslSearchRequestResponseWithHitsOnly) GetHitsOk() (*[]SearchDataMessage, bool) {
	if o == nil || o.Hits == nil {
		return nil, false
	}
	return o.Hits, true
}

// HasHits returns a boolean if a field has been set.
func (o *DslSearchRequestResponseWithHitsOnly) HasHits() bool {
	if o != nil && o.Hits != nil {
		return true
	}

	return false
}

// SetHits gets a reference to the given []SearchDataMessage and assigns it to the Hits field.
func (o *DslSearchRequestResponseWithHitsOnly) SetHits(v []SearchDataMessage) {
	o.Hits = &v
}

// GetTookMillis returns the TookMillis field value if set, zero value otherwise.
func (o *DslSearchRequestResponseWithHitsOnly) GetTookMillis() int64 {
	if o == nil || o.TookMillis == nil {
		var ret int64
		return ret
	}
	return *o.TookMillis
}

// GetTookMillisOk returns a tuple with the TookMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DslSearchRequestResponseWithHitsOnly) GetTookMillisOk() (*int64, bool) {
	if o == nil || o.TookMillis == nil {
		return nil, false
	}
	return o.TookMillis, true
}

// HasTookMillis returns a boolean if a field has been set.
func (o *DslSearchRequestResponseWithHitsOnly) HasTookMillis() bool {
	if o != nil && o.TookMillis != nil {
		return true
	}

	return false
}

// SetTookMillis gets a reference to the given int64 and assigns it to the TookMillis field.
func (o *DslSearchRequestResponseWithHitsOnly) SetTookMillis(v int64) {
	o.TookMillis = &v
}

// GetTotalHits returns the TotalHits field value if set, zero value otherwise.
func (o *DslSearchRequestResponseWithHitsOnly) GetTotalHits() int64 {
	if o == nil || o.TotalHits == nil {
		var ret int64
		return ret
	}
	return *o.TotalHits
}

// GetTotalHitsOk returns a tuple with the TotalHits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DslSearchRequestResponseWithHitsOnly) GetTotalHitsOk() (*int64, bool) {
	if o == nil || o.TotalHits == nil {
		return nil, false
	}
	return o.TotalHits, true
}

// HasTotalHits returns a boolean if a field has been set.
func (o *DslSearchRequestResponseWithHitsOnly) HasTotalHits() bool {
	if o != nil && o.TotalHits != nil {
		return true
	}

	return false
}

// SetTotalHits gets a reference to the given int64 and assigns it to the TotalHits field.
func (o *DslSearchRequestResponseWithHitsOnly) SetTotalHits(v int64) {
	o.TotalHits = &v
}

func (o DslSearchRequestResponseWithHitsOnly) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Hits != nil {
		toSerialize["hits"] = o.Hits
	}
	if o.TookMillis != nil {
		toSerialize["tookMillis"] = o.TookMillis
	}
	if o.TotalHits != nil {
		toSerialize["totalHits"] = o.TotalHits
	}
	return json.Marshal(toSerialize)
}

type NullableDslSearchRequestResponseWithHitsOnly struct {
	value *DslSearchRequestResponseWithHitsOnly
	isSet bool
}

func (v NullableDslSearchRequestResponseWithHitsOnly) Get() *DslSearchRequestResponseWithHitsOnly {
	return v.value
}

func (v *NullableDslSearchRequestResponseWithHitsOnly) Set(val *DslSearchRequestResponseWithHitsOnly) {
	v.value = val
	v.isSet = true
}

func (v NullableDslSearchRequestResponseWithHitsOnly) IsSet() bool {
	return v.isSet
}

func (v *NullableDslSearchRequestResponseWithHitsOnly) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDslSearchRequestResponseWithHitsOnly(val *DslSearchRequestResponseWithHitsOnly) *NullableDslSearchRequestResponseWithHitsOnly {
	return &NullableDslSearchRequestResponseWithHitsOnly{value: val, isSet: true}
}

func (v NullableDslSearchRequestResponseWithHitsOnly) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDslSearchRequestResponseWithHitsOnly) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


