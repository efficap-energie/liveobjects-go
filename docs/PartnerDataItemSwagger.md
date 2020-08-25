# PartnerDataItemSwagger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Model** | **string** | Model of the injected data. Data with the same model must have coherent types in each value fields. This field is required to perform an advanced search on the data. Can not contains space or dot characters. Should not be empty or blank. | [optional] 
**StreamId** | **string** | StreamId in which the data will be added. Should not contains following character or space : &#39; \\ \&quot; ; { } ( ) | [optional] 
**Tags** | **[]string** | Easy way to &#39;tag&#39; the data in order to ease up advanced search through all streams and models | [optional] 
**Timestamp** | **string** | ISO 8601 date time | [optional] 
**Value** | **map[string]interface{}** | Must be a JsonObject (primitive objects like string or int are not allowed) | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


