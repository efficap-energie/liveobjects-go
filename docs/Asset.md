# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | **bool** | True if asset is considered connected | 
**CreationTs** | **int64** | Date/time when asset was first registered | 
**Description** | Pointer to **string** | device description | [optional] 
**GroupId** | Pointer to **string** | Group Id where this asset is localize | [optional] 
**GroupPath** | Pointer to **string** | Group path where this asset is localize | [optional] 
**Id** | **string** | Asset identifier | 
**LastUpdateTs** | Pointer to **int64** | Date/time when asset status has been lastly updated | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Additional asset info (from last device \&quot;connected\&quot; event) | [optional] 
**Name** | Pointer to **string** | human readable name | [optional] 
**Namespace** | **string** | Asset identifier namespace | 
**Path** | Pointer to [**[]AssetAlias**](AssetAlias.md) | Last path taken by a received message describing the asset status(i.e. connected/disconnected): the series of assets that forwarded the message. | [optional] 
**Properties** | Pointer to **map[string]string** | Static asset properties (from device provisioning | [optional] 
**Tags** | Pointer to **[]string** | Device tags | [optional] 
**TopicCommand** | Pointer to **string** | Topic where generic commands can be sent to asset (if asset is still connected) | [optional] 
**TopicParamUpdate** | Pointer to **string** | Topic where param update requests can be sent to asset (if asset is still connected) | [optional] 
**TopicResourceUpdate** | Pointer to **string** | Topic where resource update requests can be sent to asset (if asset is still connected) | [optional] 

## Methods

### NewAsset

`func NewAsset(connected bool, creationTs int64, id string, namespace string, ) *Asset`

NewAsset instantiates a new Asset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetWithDefaults

`func NewAssetWithDefaults() *Asset`

NewAssetWithDefaults instantiates a new Asset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnected

`func (o *Asset) GetConnected() bool`

GetConnected returns the Connected field if non-nil, zero value otherwise.

### GetConnectedOk

`func (o *Asset) GetConnectedOk() (*bool, bool)`

GetConnectedOk returns a tuple with the Connected field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnected

`func (o *Asset) SetConnected(v bool)`

SetConnected sets Connected field to given value.


### GetCreationTs

`func (o *Asset) GetCreationTs() int64`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *Asset) GetCreationTsOk() (*int64, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *Asset) SetCreationTs(v int64)`

SetCreationTs sets CreationTs field to given value.


### GetDescription

`func (o *Asset) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Asset) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Asset) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Asset) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupId

`func (o *Asset) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *Asset) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *Asset) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *Asset) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupPath

`func (o *Asset) GetGroupPath() string`

GetGroupPath returns the GroupPath field if non-nil, zero value otherwise.

### GetGroupPathOk

`func (o *Asset) GetGroupPathOk() (*string, bool)`

GetGroupPathOk returns a tuple with the GroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPath

`func (o *Asset) SetGroupPath(v string)`

SetGroupPath sets GroupPath field to given value.

### HasGroupPath

`func (o *Asset) HasGroupPath() bool`

HasGroupPath returns a boolean if a field has been set.

### GetId

`func (o *Asset) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Asset) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Asset) SetId(v string)`

SetId sets Id field to given value.


### GetLastUpdateTs

`func (o *Asset) GetLastUpdateTs() int64`

GetLastUpdateTs returns the LastUpdateTs field if non-nil, zero value otherwise.

### GetLastUpdateTsOk

`func (o *Asset) GetLastUpdateTsOk() (*int64, bool)`

GetLastUpdateTsOk returns a tuple with the LastUpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTs

`func (o *Asset) SetLastUpdateTs(v int64)`

SetLastUpdateTs sets LastUpdateTs field to given value.

### HasLastUpdateTs

`func (o *Asset) HasLastUpdateTs() bool`

HasLastUpdateTs returns a boolean if a field has been set.

### GetMetadata

`func (o *Asset) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Asset) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Asset) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Asset) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetName

`func (o *Asset) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Asset) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Asset) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Asset) HasName() bool`

HasName returns a boolean if a field has been set.

### GetNamespace

`func (o *Asset) GetNamespace() string`

GetNamespace returns the Namespace field if non-nil, zero value otherwise.

### GetNamespaceOk

`func (o *Asset) GetNamespaceOk() (*string, bool)`

GetNamespaceOk returns a tuple with the Namespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNamespace

`func (o *Asset) SetNamespace(v string)`

SetNamespace sets Namespace field to given value.


### GetPath

`func (o *Asset) GetPath() []AssetAlias`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Asset) GetPathOk() (*[]AssetAlias, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Asset) SetPath(v []AssetAlias)`

SetPath sets Path field to given value.

### HasPath

`func (o *Asset) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetProperties

`func (o *Asset) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Asset) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Asset) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Asset) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *Asset) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Asset) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Asset) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Asset) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTopicCommand

`func (o *Asset) GetTopicCommand() string`

GetTopicCommand returns the TopicCommand field if non-nil, zero value otherwise.

### GetTopicCommandOk

`func (o *Asset) GetTopicCommandOk() (*string, bool)`

GetTopicCommandOk returns a tuple with the TopicCommand field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicCommand

`func (o *Asset) SetTopicCommand(v string)`

SetTopicCommand sets TopicCommand field to given value.

### HasTopicCommand

`func (o *Asset) HasTopicCommand() bool`

HasTopicCommand returns a boolean if a field has been set.

### GetTopicParamUpdate

`func (o *Asset) GetTopicParamUpdate() string`

GetTopicParamUpdate returns the TopicParamUpdate field if non-nil, zero value otherwise.

### GetTopicParamUpdateOk

`func (o *Asset) GetTopicParamUpdateOk() (*string, bool)`

GetTopicParamUpdateOk returns a tuple with the TopicParamUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicParamUpdate

`func (o *Asset) SetTopicParamUpdate(v string)`

SetTopicParamUpdate sets TopicParamUpdate field to given value.

### HasTopicParamUpdate

`func (o *Asset) HasTopicParamUpdate() bool`

HasTopicParamUpdate returns a boolean if a field has been set.

### GetTopicResourceUpdate

`func (o *Asset) GetTopicResourceUpdate() string`

GetTopicResourceUpdate returns the TopicResourceUpdate field if non-nil, zero value otherwise.

### GetTopicResourceUpdateOk

`func (o *Asset) GetTopicResourceUpdateOk() (*string, bool)`

GetTopicResourceUpdateOk returns a tuple with the TopicResourceUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTopicResourceUpdate

`func (o *Asset) SetTopicResourceUpdate(v string)`

SetTopicResourceUpdate sets TopicResourceUpdate field to given value.

### HasTopicResourceUpdate

`func (o *Asset) HasTopicResourceUpdate() bool`

HasTopicResourceUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


