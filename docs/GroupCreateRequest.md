# GroupCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | group description. Expected string (max 255 characters) | [optional] 
**Parent** | Pointer to **string** | reference to group&#39;s parent (id). Expected string (max 6 characters) | [optional] 
**PathNode** | **string** | group&#39;s local id in path (unique for groups in same parent).Authorized: letter (lowercase and uppercase), accented characters, number, space, dash, underscore and simple quote. A valid path must respect the following regular expression &lt;code&gt;^[\\wÀ-ÖØ-öø-ÿ&#39; -]{1,255}&lt;/code&gt;.Expected string (max 255 characters) | 

## Methods

### NewGroupCreateRequest

`func NewGroupCreateRequest(pathNode string, ) *GroupCreateRequest`

NewGroupCreateRequest instantiates a new GroupCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupCreateRequestWithDefaults

`func NewGroupCreateRequestWithDefaults() *GroupCreateRequest`

NewGroupCreateRequestWithDefaults instantiates a new GroupCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetParent

`func (o *GroupCreateRequest) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *GroupCreateRequest) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *GroupCreateRequest) SetParent(v string)`

SetParent sets Parent field to given value.

### HasParent

`func (o *GroupCreateRequest) HasParent() bool`

HasParent returns a boolean if a field has been set.

### GetPathNode

`func (o *GroupCreateRequest) GetPathNode() string`

GetPathNode returns the PathNode field if non-nil, zero value otherwise.

### GetPathNodeOk

`func (o *GroupCreateRequest) GetPathNodeOk() (*string, bool)`

GetPathNodeOk returns a tuple with the PathNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathNode

`func (o *GroupCreateRequest) SetPathNode(v string)`

SetPathNode sets PathNode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


