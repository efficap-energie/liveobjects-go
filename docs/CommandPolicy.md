# CommandPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationInSeconds** | **int64** | Expiration in seconds since command creation date. Default is no expiry. Min value is 5 seconds | [optional] 
**AckTimeoutInSeconds** | **int64** | Ack timeout in seconds since command was sent. Default is no ack timeout. Min value is 10 seconds | [optional] 
**AckMode** | **string** | Ack mode for this command. NONE (or AUTO) ack means that the command is automatically acknowledged (set to &#39;PROCESSED&#39; status) as the command is sent to the device. NETWORK ack requires a reception acknowledge. APPLICATIVE (or DEVICE) ack requires a command response from the device to change its status. Default ack mode is connectivity dependant. | [optional] 
**Attempts** | **int32** | Number of attemps in case of ERROR. Default to 1 | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


