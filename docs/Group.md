# Group

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTs** | **int64** | Date/time when group was first registered | 
**Description** | **string** | group description | 
**GroupId** | **string** | Group identifier | 
**LastUpdateTs** | **int64** | Date/time when group status has been lastly updated | 
**Parent** | **string** | parent | 
**Path** | **string** | path | 
**PathNode** | **string** | path node | 

## Methods

### NewGroup

`func NewGroup(creationTs int64, description string, groupId string, lastUpdateTs int64, parent string, path string, pathNode string, ) *Group`

NewGroup instantiates a new Group object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupWithDefaults

`func NewGroupWithDefaults() *Group`

NewGroupWithDefaults instantiates a new Group object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTs

`func (o *Group) GetCreationTs() int64`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *Group) GetCreationTsOk() (*int64, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *Group) SetCreationTs(v int64)`

SetCreationTs sets CreationTs field to given value.


### GetDescription

`func (o *Group) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Group) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Group) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetGroupId

`func (o *Group) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Group) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Group) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.


### GetLastUpdateTs

`func (o *Group) GetLastUpdateTs() int64`

GetLastUpdateTs returns the LastUpdateTs field if non-nil, zero value otherwise.

### GetLastUpdateTsOk

`func (o *Group) GetLastUpdateTsOk() (*int64, bool)`

GetLastUpdateTsOk returns a tuple with the LastUpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTs

`func (o *Group) SetLastUpdateTs(v int64)`

SetLastUpdateTs sets LastUpdateTs field to given value.


### GetParent

`func (o *Group) GetParent() string`

GetParent returns the Parent field if non-nil, zero value otherwise.

### GetParentOk

`func (o *Group) GetParentOk() (*string, bool)`

GetParentOk returns a tuple with the Parent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParent

`func (o *Group) SetParent(v string)`

SetParent sets Parent field to given value.


### GetPath

`func (o *Group) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Group) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Group) SetPath(v string)`

SetPath sets Path field to given value.


### GetPathNode

`func (o *Group) GetPathNode() string`

GetPathNode returns the PathNode field if non-nil, zero value otherwise.

### GetPathNodeOk

`func (o *Group) GetPathNodeOk() (*string, bool)`

GetPathNodeOk returns a tuple with the PathNode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPathNode

`func (o *Group) SetPathNode(v string)`

SetPathNode sets PathNode field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


