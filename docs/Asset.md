# Asset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connected** | **bool** | True if asset is considered connected | 
**CreationTs** | **int64** | Date/time when asset was first registered | 
**Description** | **string** | device description | [optional] 
**GroupId** | **string** | Group Id where this asset is localize | [optional] 
**GroupPath** | **string** | Group path where this asset is localize | [optional] 
**Id** | **string** | Asset identifier | 
**LastUpdateTs** | **int64** | Date/time when asset status has been lastly updated | [optional] 
**Metadata** | **map[string]interface{}** | Additional asset info (from last device \&quot;connected\&quot; event) | [optional] 
**Name** | **string** | human readable name | [optional] 
**Namespace** | **string** | Asset identifier namespace | 
**Path** | [**[]AssetAlias**](AssetAlias.md) | Last path taken by a received message describing the asset status(i.e. connected/disconnected): the series of assets that forwarded the message. | [optional] 
**Properties** | **map[string]string** | Static asset properties (from device provisioning | [optional] 
**Tags** | **[]string** | Device tags | [optional] 
**TopicCommand** | **string** | Topic where generic commands can be sent to asset (if asset is still connected) | [optional] 
**TopicParamUpdate** | **string** | Topic where param update requests can be sent to asset (if asset is still connected) | [optional] 
**TopicResourceUpdate** | **string** | Topic where resource update requests can be sent to asset (if asset is still connected) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


