# ActivityState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivityRule** | [**ActivityRule**](ActivityRule.md) |  | 
**DeviceId** | **string** | id of the targeted device | 
**LastActivity** | **string** | timestamp (ISO 8601) of the last registered contact of the device | [optional] 
**NextAlarm** | **string** | timestamp (ISO 8601) of the next alarm for this rule if the device remains silent | 
**NumberOfAlarmReminders** | **int32** | number of times a reminder has been sent for the current state | 
**State** | **string** | current state of the device | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


