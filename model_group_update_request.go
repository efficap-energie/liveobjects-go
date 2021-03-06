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

// GroupUpdateRequest struct for GroupUpdateRequest
type GroupUpdateRequest struct {
	// group description. Expected string (max 255 characters)
	Description *string `json:"description,omitempty"`
	// group id. Expected string (max 6 characters)
	Id *string `json:"id,omitempty"`
	// reference to group's parent (id). Expected string (max 6 characters)
	ParentId *string `json:"parentId,omitempty"`
	// group's local id in path (unique for groups in same parent).Authorized: letter (lowercase and uppercase), accented characters, number, space, dash, underscore and simple quote. A valid path must respect the following regular expression <code>^[\\wÀ-ÖØ-öø-ÿ' -]{1,255}</code>.Expected string (max 255 characters)
	PathNode string `json:"pathNode"`
}

// NewGroupUpdateRequest instantiates a new GroupUpdateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewGroupUpdateRequest(pathNode string, ) *GroupUpdateRequest {
	this := GroupUpdateRequest{}
	this.PathNode = pathNode
	return &this
}

// NewGroupUpdateRequestWithDefaults instantiates a new GroupUpdateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewGroupUpdateRequestWithDefaults() *GroupUpdateRequest {
	this := GroupUpdateRequest{}
	return &this
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *GroupUpdateRequest) GetDescription() string {
	if o == nil || o.Description == nil {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdateRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || o.Description == nil {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *GroupUpdateRequest) HasDescription() bool {
	if o != nil && o.Description != nil {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *GroupUpdateRequest) SetDescription(v string) {
	o.Description = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *GroupUpdateRequest) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdateRequest) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *GroupUpdateRequest) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *GroupUpdateRequest) SetId(v string) {
	o.Id = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *GroupUpdateRequest) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *GroupUpdateRequest) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *GroupUpdateRequest) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *GroupUpdateRequest) SetParentId(v string) {
	o.ParentId = &v
}

// GetPathNode returns the PathNode field value
func (o *GroupUpdateRequest) GetPathNode() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.PathNode
}

// GetPathNodeOk returns a tuple with the PathNode field value
// and a boolean to check if the value has been set.
func (o *GroupUpdateRequest) GetPathNodeOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.PathNode, true
}

// SetPathNode sets field value
func (o *GroupUpdateRequest) SetPathNode(v string) {
	o.PathNode = v
}

func (o GroupUpdateRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Description != nil {
		toSerialize["description"] = o.Description
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if true {
		toSerialize["pathNode"] = o.PathNode
	}
	return json.Marshal(toSerialize)
}

type NullableGroupUpdateRequest struct {
	value *GroupUpdateRequest
	isSet bool
}

func (v NullableGroupUpdateRequest) Get() *GroupUpdateRequest {
	return v.value
}

func (v *NullableGroupUpdateRequest) Set(val *GroupUpdateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableGroupUpdateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableGroupUpdateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableGroupUpdateRequest(val *GroupUpdateRequest) *NullableGroupUpdateRequest {
	return &NullableGroupUpdateRequest{value: val, isSet: true}
}

func (v NullableGroupUpdateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableGroupUpdateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


