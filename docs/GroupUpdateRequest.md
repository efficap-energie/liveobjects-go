# GroupUpdateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | group description. Expected string (max 255 characters) | [optional] 
**Id** | Pointer to **string** | group id. Expected string (max 6 characters) | [optional] 
**ParentId** | Pointer to **string** | reference to group&#39;s parent (id). Expected string (max 6 characters) | [optional] 
**PathNode** | **string** | group&#39;s local id in path (unique for groups in same parent).Authorized: letter (lowercase and uppercase), accented characters, number, space, dash, underscore and simple quote. A valid path must respect the following regular expression &lt;code&gt;^[\\wÀ-ÖØ-öø-ÿ&#39; -]{1,255}&lt;/code&gt;.Expected string (max 255 characters) | 

## Methods

### NewGroupUpdateRequest

`func NewGroupUpdateRequest(pathNode string, ) *GroupUpdateRequest`

NewGroupUpdateRequest instantiates a new GroupUpdateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupUpdateRequestWithDefaults

`func NewGroupUpdateRequestWithDefaults() *GroupUpdateRequest`

NewGroupUpdateRequestWithDefaults instantiates a new GroupUpdateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GroupUpdateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GroupUpdateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GroupUpdateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GroupUpdateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GroupUpdateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GroupUpdateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GroupUpdateRequest) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GroupUpdateRequest) HasId() bool`

HasId returns a boolean if a field has been set.

### GetParentId

`func (o *GroupUpdateRequest) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *GroupUpdateRequest) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *GroupUpdateRequest) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *GroupUpdateRequest) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetPathNode

`func (o *GroupUpdateRequest) GetPathNode() string`

GetPathNode returns the PathNode field if non-nil, zero value otherwise.

### GetPathNodeOk

`func (o *GroupUpdateRequest) GetPathNodeOk() (*string, bool)`

GetPathNodeOk returns a tuple with the PathNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathNode

`func (o *GroupUpdateRequest) SetPathNode(v string)`

SetPathNode sets PathNode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


