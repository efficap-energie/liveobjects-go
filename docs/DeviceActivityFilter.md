# DeviceActivityFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceIds** | Pointer to **[]string** | list of filtered devices. Null or empty to accept all devices | [optional] 
**RuleIds** | Pointer to **[]string** | list of filtered rule Ids. Null or empty to accept all rule Ids | [optional] 

## Methods

### NewDeviceActivityFilter

`func NewDeviceActivityFilter() *DeviceActivityFilter`

NewDeviceActivityFilter instantiates a new DeviceActivityFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceActivityFilterWithDefaults

`func NewDeviceActivityFilterWithDefaults() *DeviceActivityFilter`

NewDeviceActivityFilterWithDefaults instantiates a new DeviceActivityFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceIds

`func (o *DeviceActivityFilter) GetDeviceIds() []string`

GetDeviceIds returns the DeviceIds field if non-nil, zero value otherwise.

### GetDeviceIdsOk

`func (o *DeviceActivityFilter) GetDeviceIdsOk() (*[]string, bool)`

GetDeviceIdsOk returns a tuple with the DeviceIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceIds

`func (o *DeviceActivityFilter) SetDeviceIds(v []string)`

SetDeviceIds sets DeviceIds field to given value.

### HasDeviceIds

`func (o *DeviceActivityFilter) HasDeviceIds() bool`

HasDeviceIds returns a boolean if a field has been set.

### GetRuleIds

`func (o *DeviceActivityFilter) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *DeviceActivityFilter) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *DeviceActivityFilter) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *DeviceActivityFilter) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


