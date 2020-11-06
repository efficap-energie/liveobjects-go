# MessageSelectorFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to **[]string** | list of filtered connectors. | [optional] 
**DeviceIds** | Pointer to **[]string** | list of filtered devices. | [optional] 
**GroupPaths** | Pointer to [**[]GroupPath**](GroupPath.md) | list of filtered group paths | [optional] 

## Methods

### NewMessageSelectorFilter

`func NewMessageSelectorFilter() *MessageSelectorFilter`

NewMessageSelectorFilter instantiates a new MessageSelectorFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSelectorFilterWithDefaults

`func NewMessageSelectorFilterWithDefaults() *MessageSelectorFilter`

NewMessageSelectorFilterWithDefaults instantiates a new MessageSelectorFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *MessageSelectorFilter) GetConnectors() []string`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *MessageSelectorFilter) GetConnectorsOk() (*[]string, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *MessageSelectorFilter) SetConnectors(v []string)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *MessageSelectorFilter) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetDeviceIds

`func (o *MessageSelectorFilter) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *MessageSelectorFilter) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *MessageSelectorFilter) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *MessageSelectorFilter) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetGroupPaths

`func (o *MessageSelectorFilter) GetGroupPaths() []GroupPath`

GetGroupPaths returns the GroupPaths field if non-nil, zero value otherwise.

### GetGroupPathsOk

`func (o *MessageSelectorFilter) GetGroupPathsOk() (*[]GroupPath, bool)`

GetGroupPathsOk returns a tuple with the GroupPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPaths

`func (o *MessageSelectorFilter) SetGroupPaths(v []GroupPath)`

SetGroupPaths sets GroupPaths field to given value.

### HasGroupPaths

`func (o *MessageSelectorFilter) HasGroupPaths() bool`

HasGroupPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


