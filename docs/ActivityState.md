# ActivityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityRule** | [**ActivityRule**](ActivityRule.md) |  | 
**DeviceId** | **string** | id of the targeted device | 
**LastActivity** | Pointer to **string** | timestamp (ISO 8601) of the last registered contact of the device | [optional] 
**NextAlarm** | **string** | timestamp (ISO 8601) of the next alarm for this rule if the device remains silent | 
**NumberOfAlarmReminders** | **int32** | number of times a reminder has been sent for the current state | 
**State** | **string** | current state of the device | 

## Methods

### NewActivityState

`func NewActivityState(activityRule ActivityRule, deviceId string, nextAlarm string, numberOfAlarmReminders int32, state string, ) *ActivityState`

NewActivityState instantiates a new ActivityState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityStateWithDefaults

`func NewActivityStateWithDefaults() *ActivityState`

NewActivityStateWithDefaults instantiates a new ActivityState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivityRule

`func (o *ActivityState) GetActivityRule() ActivityRule`

GetActivityRule returns the ActivityRule field if non-nil, zero value otherwise.

### GetActivityRuleOk

`func (o *ActivityState) GetActivityRuleOk() (*ActivityRule, bool)`

GetActivityRuleOk returns a tuple with the ActivityRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityRule

`func (o *ActivityState) SetActivityRule(v ActivityRule)`

SetActivityRule sets ActivityRule field to given value.


### GetDeviceId

`func (o *ActivityState) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *ActivityState) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *ActivityState) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.


### GetLastActivity

`func (o *ActivityState) GetLastActivity() string`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *ActivityState) GetLastActivityOk() (*string, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *ActivityState) SetLastActivity(v string)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *ActivityState) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetNextAlarm

`func (o *ActivityState) GetNextAlarm() string`

GetNextAlarm returns the NextAlarm field if non-nil, zero value otherwise.

### GetNextAlarmOk

`func (o *ActivityState) GetNextAlarmOk() (*string, bool)`

GetNextAlarmOk returns a tuple with the NextAlarm field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNextAlarm

`func (o *ActivityState) SetNextAlarm(v string)`

SetNextAlarm sets NextAlarm field to given value.


### GetNumberOfAlarmReminders

`func (o *ActivityState) GetNumberOfAlarmReminders() int32`

GetNumberOfAlarmReminders returns the NumberOfAlarmReminders field if non-nil, zero value otherwise.

### GetNumberOfAlarmRemindersOk

`func (o *ActivityState) GetNumberOfAlarmRemindersOk() (*int32, bool)`

GetNumberOfAlarmRemindersOk returns a tuple with the NumberOfAlarmReminders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfAlarmReminders

`func (o *ActivityState) SetNumberOfAlarmReminders(v int32)`

SetNumberOfAlarmReminders sets NumberOfAlarmReminders field to given value.


### GetState

`func (o *ActivityState) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *ActivityState) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *ActivityState) SetState(v string)`

SetState sets State field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


