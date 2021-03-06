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

// GeozoneContainer struct for GeozoneContainer
type GeozoneContainer struct {
	Description *string `json:"description,omitempty"`
	Geometry *Polygon `json:"geometry,omitempty"`
	Properties *map[string]interface{} `json:"properties,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	ZoneId *string `json:"zoneId,omitempty"`
}

// NewGeozoneContainer instantiates a new GeozoneContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGeozoneContainer() *GeozoneContainer {
	this := GeozoneContainer{}
	return &this
}

// NewGeozoneContainerWithDefaults instantiates a new GeozoneContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGeozoneContainerWithDefaults() *GeozoneContainer {
	this := GeozoneContainer{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GeozoneContainer) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeozoneContainer) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GeozoneContainer) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GeozoneContainer) SetDescription(v string) {
	o.Description = &v
}

// GetGeometry returns the Geometry field value if set, zero value otherwise.
func (o *GeozoneContainer) GetGeometry() Polygon {
	if o == nil || o.Geometry == nil {
		var ret Polygon
		return ret
	}
	return *o.Geometry
}

// GetGeometryOk returns a tuple with the Geometry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeozoneContainer) GetGeometryOk() (*Polygon, bool) {
	if o == nil || o.Geometry == nil {
		return nil, false
	}
	return o.Geometry, true
}

// HasGeometry returns a boolean if a field has been set.
func (o *GeozoneContainer) HasGeometry() bool {
	if o != nil && o.Geometry != nil {
		return true
	}

	return false
}

// SetGeometry gets a reference to the given Polygon and assigns it to the Geometry field.
func (o *GeozoneContainer) SetGeometry(v Polygon) {
	o.Geometry = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *GeozoneContainer) GetProperties() map[string]interface{} {
	if o == nil || o.Properties == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeozoneContainer) GetPropertiesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *GeozoneContainer) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]interface{} and assigns it to the Properties field.
func (o *GeozoneContainer) SetProperties(v map[string]interface{}) {
	o.Properties = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *GeozoneContainer) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeozoneContainer) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *GeozoneContainer) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *GeozoneContainer) SetTags(v []string) {
	o.Tags = &v
}

// GetZoneId returns the ZoneId field value if set, zero value otherwise.
func (o *GeozoneContainer) GetZoneId() string {
	if o == nil || o.ZoneId == nil {
		var ret string
		return ret
	}
	return *o.ZoneId
}

// GetZoneIdOk returns a tuple with the ZoneId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GeozoneContainer) GetZoneIdOk() (*string, bool) {
	if o == nil || o.ZoneId == nil {
		return nil, false
	}
	return o.ZoneId, true
}

// HasZoneId returns a boolean if a field has been set.
func (o *GeozoneContainer) HasZoneId() bool {
	if o != nil && o.ZoneId != nil {
		return true
	}

	return false
}

// SetZoneId gets a reference to the given string and assigns it to the ZoneId field.
func (o *GeozoneContainer) SetZoneId(v string) {
	o.ZoneId = &v
}

func (o GeozoneContainer) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Geometry != nil {
		toSerialize["geometry"] = o.Geometry
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.ZoneId != nil {
		toSerialize["zoneId"] = o.ZoneId
	}
	return json.Marshal(toSerialize)
}

type NullableGeozoneContainer struct {
	value *GeozoneContainer
	isSet bool
}

func (v NullableGeozoneContainer) Get() *GeozoneContainer {
	return v.value
}

func (v *NullableGeozoneContainer) Set(val *GeozoneContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableGeozoneContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableGeozoneContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGeozoneContainer(val *GeozoneContainer) *NullableGeozoneContainer {
	return &NullableGeozoneContainer{value: val, isSet: true}
}

func (v NullableGeozoneContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGeozoneContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


