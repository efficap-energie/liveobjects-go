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

// DayMetrics struct for DayMetrics
type DayMetrics struct {
	// connector (lora, mqtt, http) statistics
	Connectors *map[string]ConnectorAccounting `json:"connectors,omitempty"`
	// day in \"YYYY-MM-DD\" format.
	Day *string `json:"day,omitempty"`
	// service statistics
	Services *map[string]map[string]int64 `json:"services,omitempty"`
}

// NewDayMetrics instantiates a new DayMetrics object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDayMetrics() *DayMetrics {
	this := DayMetrics{}
	return &this
}

// NewDayMetricsWithDefaults instantiates a new DayMetrics object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDayMetricsWithDefaults() *DayMetrics {
	this := DayMetrics{}
	return &this
}

// GetConnectors returns the Connectors field value if set, zero value otherwise.
func (o *DayMetrics) GetConnectors() map[string]ConnectorAccounting {
	if o == nil || o.Connectors == nil {
		var ret map[string]ConnectorAccounting
		return ret
	}
	return *o.Connectors
}

// GetConnectorsOk returns a tuple with the Connectors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayMetrics) GetConnectorsOk() (*map[string]ConnectorAccounting, bool) {
	if o == nil || o.Connectors == nil {
		return nil, false
	}
	return o.Connectors, true
}

// HasConnectors returns a boolean if a field has been set.
func (o *DayMetrics) HasConnectors() bool {
	if o != nil && o.Connectors != nil {
		return true
	}

	return false
}

// SetConnectors gets a reference to the given map[string]ConnectorAccounting and assigns it to the Connectors field.
func (o *DayMetrics) SetConnectors(v map[string]ConnectorAccounting) {
	o.Connectors = &v
}

// GetDay returns the Day field value if set, zero value otherwise.
func (o *DayMetrics) GetDay() string {
	if o == nil || o.Day == nil {
		var ret string
		return ret
	}
	return *o.Day
}

// GetDayOk returns a tuple with the Day field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayMetrics) GetDayOk() (*string, bool) {
	if o == nil || o.Day == nil {
		return nil, false
	}
	return o.Day, true
}

// HasDay returns a boolean if a field has been set.
func (o *DayMetrics) HasDay() bool {
	if o != nil && o.Day != nil {
		return true
	}

	return false
}

// SetDay gets a reference to the given string and assigns it to the Day field.
func (o *DayMetrics) SetDay(v string) {
	o.Day = &v
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *DayMetrics) GetServices() map[string]map[string]int64 {
	if o == nil || o.Services == nil {
		var ret map[string]map[string]int64
		return ret
	}
	return *o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DayMetrics) GetServicesOk() (*map[string]map[string]int64, bool) {
	if o == nil || o.Services == nil {
		return nil, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *DayMetrics) HasServices() bool {
	if o != nil && o.Services != nil {
		return true
	}

	return false
}

// SetServices gets a reference to the given map[string]map[string]int64 and assigns it to the Services field.
func (o *DayMetrics) SetServices(v map[string]map[string]int64) {
	o.Services = &v
}

func (o DayMetrics) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Connectors != nil {
		toSerialize["connectors"] = o.Connectors
	}
	if o.Day != nil {
		toSerialize["day"] = o.Day
	}
	if o.Services != nil {
		toSerialize["services"] = o.Services
	}
	return json.Marshal(toSerialize)
}

type NullableDayMetrics struct {
	value *DayMetrics
	isSet bool
}

func (v NullableDayMetrics) Get() *DayMetrics {
	return v.value
}

func (v *NullableDayMetrics) Set(val *DayMetrics) {
	v.value = val
	v.isSet = true
}

func (v NullableDayMetrics) IsSet() bool {
	return v.isSet
}

func (v *NullableDayMetrics) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDayMetrics(val *DayMetrics) *NullableDayMetrics {
	return &NullableDayMetrics{value: val, isSet: true}
}

func (v NullableDayMetrics) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDayMetrics) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


