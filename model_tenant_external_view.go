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

// TenantExternalView struct for TenantExternalView
type TenantExternalView struct {
	// tenant Id
	Id *string `json:"id,omitempty"`
	// tenant name
	Name *string `json:"name,omitempty"`
	OfferAndOptions *OfferAndOptions `json:"offerAndOptions,omitempty"`
	// tenant properties
	Properties *map[string]string `json:"properties,omitempty"`
}

// NewTenantExternalView instantiates a new TenantExternalView object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantExternalView() *TenantExternalView {
	this := TenantExternalView{}
	return &this
}

// NewTenantExternalViewWithDefaults instantiates a new TenantExternalView object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantExternalViewWithDefaults() *TenantExternalView {
	this := TenantExternalView{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TenantExternalView) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantExternalView) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TenantExternalView) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TenantExternalView) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TenantExternalView) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantExternalView) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TenantExternalView) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TenantExternalView) SetName(v string) {
	o.Name = &v
}

// GetOfferAndOptions returns the OfferAndOptions field value if set, zero value otherwise.
func (o *TenantExternalView) GetOfferAndOptions() OfferAndOptions {
	if o == nil || o.OfferAndOptions == nil {
		var ret OfferAndOptions
		return ret
	}
	return *o.OfferAndOptions
}

// GetOfferAndOptionsOk returns a tuple with the OfferAndOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantExternalView) GetOfferAndOptionsOk() (*OfferAndOptions, bool) {
	if o == nil || o.OfferAndOptions == nil {
		return nil, false
	}
	return o.OfferAndOptions, true
}

// HasOfferAndOptions returns a boolean if a field has been set.
func (o *TenantExternalView) HasOfferAndOptions() bool {
	if o != nil && o.OfferAndOptions != nil {
		return true
	}

	return false
}

// SetOfferAndOptions gets a reference to the given OfferAndOptions and assigns it to the OfferAndOptions field.
func (o *TenantExternalView) SetOfferAndOptions(v OfferAndOptions) {
	o.OfferAndOptions = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TenantExternalView) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantExternalView) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TenantExternalView) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *TenantExternalView) SetProperties(v map[string]string) {
	o.Properties = &v
}

func (o TenantExternalView) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OfferAndOptions != nil {
		toSerialize["offerAndOptions"] = o.OfferAndOptions
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	return json.Marshal(toSerialize)
}

type NullableTenantExternalView struct {
	value *TenantExternalView
	isSet bool
}

func (v NullableTenantExternalView) Get() *TenantExternalView {
	return v.value
}

func (v *NullableTenantExternalView) Set(val *TenantExternalView) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantExternalView) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantExternalView) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantExternalView(val *TenantExternalView) *NullableTenantExternalView {
	return &NullableTenantExternalView{value: val, isSet: true}
}

func (v NullableTenantExternalView) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantExternalView) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


