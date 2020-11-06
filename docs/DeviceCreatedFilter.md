# DeviceCreatedFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to **[]string** | list of filtered connectors. | [optional] 
**GroupPaths** | Pointer to [**[]GroupPath**](GroupPath.md) | list of filtered group paths. Null or empty to accept all group paths | [optional] 
**Tags** | Pointer to [**[][]string**](array.md) | list of groups of tags that should be contained in message tags. There is a match if at least one group of tags is a match. A group of tags is a match if the tags of the message contains all elements of this group.&lt;br&gt;e.g. [[\&quot;ALERT\&quot;]] will match all messages containing tag &#39;ALERT&#39;.&lt;br&gt;e.g. [[\&quot;HIGH\&quot;, \&quot;ALERT\&quot;],[\&quot;PROD\&quot;]] will match messages with either tag &#39;PROD&#39; or both tags &#39;ALERT&#39; and &#39;HIGH&#39;. | [optional] 

## Methods

### NewDeviceCreatedFilter

`func NewDeviceCreatedFilter() *DeviceCreatedFilter`

NewDeviceCreatedFilter instantiates a new DeviceCreatedFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreatedFilterWithDefaults

`func NewDeviceCreatedFilterWithDefaults() *DeviceCreatedFilter`

NewDeviceCreatedFilterWithDefaults instantiates a new DeviceCreatedFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *DeviceCreatedFilter) GetConnectors() []string`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DeviceCreatedFilter) GetConnectorsOk() (*[]string, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DeviceCreatedFilter) SetConnectors(v []string)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DeviceCreatedFilter) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetGroupPaths

`func (o *DeviceCreatedFilter) GetGroupPaths() []GroupPath`

GetGroupPaths returns the GroupPaths field if non-nil, zero value otherwise.

### GetGroupPathsOk

`func (o *DeviceCreatedFilter) GetGroupPathsOk() (*[]GroupPath, bool)`

GetGroupPathsOk returns a tuple with the GroupPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPaths

`func (o *DeviceCreatedFilter) SetGroupPaths(v []GroupPath)`

SetGroupPaths sets GroupPaths field to given value.

### HasGroupPaths

`func (o *DeviceCreatedFilter) HasGroupPaths() bool`

HasGroupPaths returns a boolean if a field has been set.

### GetTags

`func (o *DeviceCreatedFilter) GetTags() [][]string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceCreatedFilter) GetTagsOk() (*[][]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceCreatedFilter) SetTags(v [][]string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceCreatedFilter) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


