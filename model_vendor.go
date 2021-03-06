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

// Vendor struct for Vendor
type Vendor struct {
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
}

// NewVendor instantiates a new Vendor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVendor() *Vendor {
	this := Vendor{}
	return &this
}

// NewVendorWithDefaults instantiates a new Vendor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVendorWithDefaults() *Vendor {
	this := Vendor{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Vendor) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Vendor) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Vendor) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Vendor) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Vendor) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Vendor) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Vendor) SetName(v string) {
	o.Name = &v
}

func (o Vendor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableVendor struct {
	value *Vendor
	isSet bool
}

func (v NullableVendor) Get() *Vendor {
	return v.value
}

func (v *NullableVendor) Set(val *Vendor) {
	v.value = val
	v.isSet = true
}

func (v NullableVendor) IsSet() bool {
	return v.isSet
}

func (v *NullableVendor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVendor(val *Vendor) *NullableVendor {
	return &NullableVendor{value: val, isSet: true}
}

func (v NullableVendor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVendor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


