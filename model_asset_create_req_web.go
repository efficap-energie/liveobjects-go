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

// AssetCreateReqWeb struct for AssetCreateReqWeb
type AssetCreateReqWeb struct {
	// user-readable description of the asset. Expected string (max 500 characters)
	Description *string `json:"description,omitempty"`
	// associate the asset to a group using the group id. Expected string (max 6 characters)
	GroupId *string `json:"groupId,omitempty"`
	// associate the asset to a group using the group path. Authorized: letter (lowercase and uppercase), accented characters, number, space, dash, underscore and simple quote. A valid path must respect the following regular expression <code>^[\\wÀ-ÖØ-öø-ÿ' -]{1,255}</code>.Expected string (max 255 characters)
	GroupPath *string `json:"groupPath,omitempty"`
	// asset identifier value. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
	Id string `json:"id"`
	// asset user-readable name (for display in web portal).Expected string (max 255 characters)
	Name *string `json:"name,omitempty"`
	// asset identifier namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
	Namespace string `json:"namespace"`
	// map of key/value string pairs detailed properties of the device. Max number of properties depends of your offer settings. A property name must not include following characters <code>$.</code> and max length is 128. Invalid property names are : 'class', '_class'. Property value max length is 256.
	Properties *map[string]string `json:"properties,omitempty"`
	// series of tags associated with the asset. Max number of tags depends of your offer settings. Tag value max length is 32.
	Tags *[]string `json:"tags,omitempty"`
}

// NewAssetCreateReqWeb instantiates a new AssetCreateReqWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAssetCreateReqWeb(id string, namespace string, ) *AssetCreateReqWeb {
	this := AssetCreateReqWeb{}
	this.Id = id
	this.Namespace = namespace
	return &this
}

// NewAssetCreateReqWebWithDefaults instantiates a new AssetCreateReqWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAssetCreateReqWebWithDefaults() *AssetCreateReqWeb {
	this := AssetCreateReqWeb{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AssetCreateReqWeb) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AssetCreateReqWeb) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AssetCreateReqWeb) SetDescription(v string) {
	o.Description = &v
}

// GetGroupId returns the GroupId field value if set, zero value otherwise.
func (o *AssetCreateReqWeb) GetGroupId() string {
	if o == nil || o.GroupId == nil {
		var ret string
		return ret
	}
	return *o.GroupId
}

// GetGroupIdOk returns a tuple with the GroupId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetGroupIdOk() (*string, bool) {
	if o == nil || o.GroupId == nil {
		return nil, false
	}
	return o.GroupId, true
}

// HasGroupId returns a boolean if a field has been set.
func (o *AssetCreateReqWeb) HasGroupId() bool {
	if o != nil && o.GroupId != nil {
		return true
	}

	return false
}

// SetGroupId gets a reference to the given string and assigns it to the GroupId field.
func (o *AssetCreateReqWeb) SetGroupId(v string) {
	o.GroupId = &v
}

// GetGroupPath returns the GroupPath field value if set, zero value otherwise.
func (o *AssetCreateReqWeb) GetGroupPath() string {
	if o == nil || o.GroupPath == nil {
		var ret string
		return ret
	}
	return *o.GroupPath
}

// GetGroupPathOk returns a tuple with the GroupPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetGroupPathOk() (*string, bool) {
	if o == nil || o.GroupPath == nil {
		return nil, false
	}
	return o.GroupPath, true
}

// HasGroupPath returns a boolean if a field has been set.
func (o *AssetCreateReqWeb) HasGroupPath() bool {
	if o != nil && o.GroupPath != nil {
		return true
	}

	return false
}

// SetGroupPath gets a reference to the given string and assigns it to the GroupPath field.
func (o *AssetCreateReqWeb) SetGroupPath(v string) {
	o.GroupPath = &v
}

// GetId returns the Id field value
func (o *AssetCreateReqWeb) GetId() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetIdOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *AssetCreateReqWeb) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AssetCreateReqWeb) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AssetCreateReqWeb) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AssetCreateReqWeb) SetName(v string) {
	o.Name = &v
}

// GetNamespace returns the Namespace field value
func (o *AssetCreateReqWeb) GetNamespace() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Namespace
}

// GetNamespaceOk returns a tuple with the Namespace field value
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetNamespaceOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Namespace, true
}

// SetNamespace sets field value
func (o *AssetCreateReqWeb) SetNamespace(v string) {
	o.Namespace = v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *AssetCreateReqWeb) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *AssetCreateReqWeb) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *AssetCreateReqWeb) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *AssetCreateReqWeb) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AssetCreateReqWeb) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *AssetCreateReqWeb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *AssetCreateReqWeb) SetTags(v []string) {
	o.Tags = &v
}

func (o AssetCreateReqWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.GroupId != nil {
		toSerialize["groupId"] = o.GroupId
	}
	if o.GroupPath != nil {
		toSerialize["groupPath"] = o.GroupPath
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["namespace"] = o.Namespace
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	return json.Marshal(toSerialize)
}

type NullableAssetCreateReqWeb struct {
	value *AssetCreateReqWeb
	isSet bool
}

func (v NullableAssetCreateReqWeb) Get() *AssetCreateReqWeb {
	return v.value
}

func (v *NullableAssetCreateReqWeb) Set(val *AssetCreateReqWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableAssetCreateReqWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableAssetCreateReqWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAssetCreateReqWeb(val *AssetCreateReqWeb) *NullableAssetCreateReqWeb {
	return &NullableAssetCreateReqWeb{value: val, isSet: true}
}

func (v NullableAssetCreateReqWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAssetCreateReqWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


