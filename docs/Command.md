# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | command unique identifier | [optional] 
**TargetDeviceId** | **string** | targeted device identifier (URN) | [optional] 
**Request** | [**CommandRequest**](CommandRequest.md) |  | [optional] 
**Response** | [**CommandResponse**](CommandResponse.md) |  | [optional] 
**Status** | **string** | command current status | [optional] 
**DeliveryStatus** | **string** | command current delivery status | [optional] 
**ErrorCode** | **string** | error code in case of ERROR status | [optional] 
**Policy** | [**CommandPolicy**](CommandPolicy.md) |  | [optional] 
**History** | [**[]CommandHistory**](CommandHistory.md) | command history | [optional] 
**NotifyTo** | **string** | topic where command status change notification are published | [optional] 
**Created** | **string** | command creation date/time | [optional] 
**Updated** | **string** | command last status update date/time | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


