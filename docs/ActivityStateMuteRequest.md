# ActivityStateMuteRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityRuleId** | Pointer to **string** | id of the targeted activityRule. At least one of deviceId/activityRuleId must be set. | [optional] 
**DeviceId** | Pointer to **string** | id of the targeted device. At least one of deviceId/activityRuleId must be set. | [optional] 
**NextAlarmOrder** | **string** | set the order type : mute or reset (from now) the next alarm | 

## Methods

### NewActivityStateMuteRequest

`func NewActivityStateMuteRequest(nextAlarmOrder string, ) *ActivityStateMuteRequest`

NewActivityStateMuteRequest instantiates a new ActivityStateMuteRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityStateMuteRequestWithDefaults

`func NewActivityStateMuteRequestWithDefaults() *ActivityStateMuteRequest`

NewActivityStateMuteRequestWithDefaults instantiates a new ActivityStateMuteRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityRuleId

`func (o *ActivityStateMuteRequest) GetActivityRuleId() string`

GetActivityRuleId returns the ActivityRuleId field if non-nil, zero value otherwise.

### GetActivityRuleIdOk

`func (o *ActivityStateMuteRequest) GetActivityRuleIdOk() (*string, bool)`

GetActivityRuleIdOk returns a tuple with the ActivityRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityRuleId

`func (o *ActivityStateMuteRequest) SetActivityRuleId(v string)`

SetActivityRuleId sets ActivityRuleId field to given value.

### HasActivityRuleId

`func (o *ActivityStateMuteRequest) HasActivityRuleId() bool`

HasActivityRuleId returns a boolean if a field has been set.

### GetDeviceId

`func (o *ActivityStateMuteRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ActivityStateMuteRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ActivityStateMuteRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *ActivityStateMuteRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetNextAlarmOrder

`func (o *ActivityStateMuteRequest) GetNextAlarmOrder() string`

GetNextAlarmOrder returns the NextAlarmOrder field if non-nil, zero value otherwise.

### GetNextAlarmOrderOk

`func (o *ActivityStateMuteRequest) GetNextAlarmOrderOk() (*string, bool)`

GetNextAlarmOrderOk returns a tuple with the NextAlarmOrder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAlarmOrder

`func (o *ActivityStateMuteRequest) SetNextAlarmOrder(v string)`

SetNextAlarmOrder sets NextAlarmOrder field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


