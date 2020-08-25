# UpdateDeviceFirmwareReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | **map[string]string** | the update metadata. Max number of metadata is 100. Metadata name max length is 255. Metadata value max length is 255. | [optional] 
**NotifyTo** | **string** | (optional) topic where firmware update status change events must be published. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;,&#39;router/&#39;]and max length is 255 | [optional] 
**Version** | **string** | requested firmware version. Expected string (max 255 characters) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


