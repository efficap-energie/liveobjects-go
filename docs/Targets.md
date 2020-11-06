# Targets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]string** | List of targeted device ids. This field should not be empty if groupPaths is empty. | [optional] 
**GroupPaths** | Pointer to [**[]GroupPath**](GroupPath.md) | list of targeted group paths. This field should not be empty if deviceIds is empty. | [optional] 

## Methods

### NewTargets

`func NewTargets() *Targets`

NewTargets instantiates a new Targets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTargetsWithDefaults

`func NewTargetsWithDefaults() *Targets`

NewTargetsWithDefaults instantiates a new Targets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *Targets) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *Targets) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *Targets) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *Targets) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetGroupPaths

`func (o *Targets) GetGroupPaths() []GroupPath`

GetGroupPaths returns the GroupPaths field if non-nil, zero value otherwise.

### GetGroupPathsOk

`func (o *Targets) GetGroupPathsOk() (*[]GroupPath, bool)`

GetGroupPathsOk returns a tuple with the GroupPaths field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPaths

`func (o *Targets) SetGroupPaths(v []GroupPath)`

SetGroupPaths sets GroupPaths field to given value.

### HasGroupPaths

`func (o *Targets) HasGroupPaths() bool`

HasGroupPaths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


