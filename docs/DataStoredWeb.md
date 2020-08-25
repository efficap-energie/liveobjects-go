# DataStoredWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **string** | ISO 8601 date time, when the data has been stored in the system | 
**Extra** | **map[string]string** | Extra info of the data: extra information enriched on the collected data (for example from the device inventory) that can be used for cross model or stream queries | [optional] 
**Id** | **string** | System id of the data. Can be use as bookmark to paginate results. | 
**Location** | [**Location**](Location.md) |  | [optional] 
**Metadata** | **map[string]interface{}** | Metadata of the stored value : may include user metadata (source) as well as additional information (protocol, server, user...) | [optional] 
**Model** | **string** | Model of the injected data. Data with the same model have coherent types in each value fields. | [optional] 
**StreamId** | **string** | streamId of the data | [optional] 
**Tags** | **[]string** | Tags of the data. Can be used to ease up advanced search through all streams and models | [optional] 
**Timestamp** | **string** | ISO 8601 date time, timestamp of the value | 
**Value** | **map[string]interface{}** | JsonObject | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


