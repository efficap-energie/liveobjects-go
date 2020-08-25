# SilentPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** | the duration (ISO 8601) after which the device will be considered as silent if it had no contact with the platform. If less than 10 minutes, this duration will be handled as a 10 minutes delay. A silent event/alarm will be sent after expiration. | 
**RepeatInterval** | **string** | if the device previously triggered a silent alarm and still had no contact with the platform; duration (ISO 8601) after which a new silent event/alarm will be sent. If less than 10 minutes, this duration will be handled as a 10 minutes delay. If null, no further alarm will be sent until the new active/silent cycle. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


