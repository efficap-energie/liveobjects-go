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

// AssetAlias struct for AssetAlias
type AssetAlias struct {
	Id *string `json:"id,omitempty"`
	Namespace *string `json:"namespace,omitempty"`
}

// NewAssetAlias instantiates a new AssetAlias object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetAlias() *AssetAlias {
	this := AssetAlias{}
	return &this
}

// NewAssetAliasWithDefaults instantiates a new AssetAlias object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetAliasWithDefaults() *AssetAlias {
	this := AssetAlias{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AssetAlias) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlias) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AssetAlias) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AssetAlias) SetId(v string) {
	o.Id = &v
}

// GetNamespace returns the Namespace field value if set, zero value otherwise.
func (o *AssetAlias) GetNamespace() string {
	if o == nil || o.Namespace == nil {
		var ret string
		return ret
	}
	return *o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetAlias) GetNamespaceOk() (*string, bool) {
	if o == nil || o.Namespace == nil {
		return nil, false
	}
	return o.Namespace, true
}

// HasNamespace returns a boolean if a field has been set.
func (o *AssetAlias) HasNamespace() bool {
	if o != nil && o.Namespace != nil {
		return true
	}

	return false
}

// SetNamespace gets a reference to the given string and assigns it to the Namespace field.
func (o *AssetAlias) SetNamespace(v string) {
	o.Namespace = &v
}

func (o AssetAlias) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.Namespace != nil {
		toSerialize["namespace"] = o.Namespace
	}
	return json.Marshal(toSerialize)
}

type NullableAssetAlias struct {
	value *AssetAlias
	isSet bool
}

func (v NullableAssetAlias) Get() *AssetAlias {
	return v.value
}

func (v *NullableAssetAlias) Set(val *AssetAlias) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetAlias) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetAlias) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetAlias(val *AssetAlias) *NullableAssetAlias {
	return &NullableAssetAlias{value: val, isSet: true}
}

func (v NullableAssetAlias) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetAlias) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


