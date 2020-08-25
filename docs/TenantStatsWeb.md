# TenantStatsWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | **map[string]string** |  | [optional] 
**StatisticsPerDay** | [**map[string]map[string]ConnectorStatistics**](map.md) | all statistics per day | [optional] 
**StatisticsPerMonth** | [**map[string]map[string]ConnectorStatistics**](map.md) | Statistics per month : number of distinct sources (lora, mqtt, datazone) | [optional] 
**TenantId** | **string** | The tenant Id | [optional] 
**TenantName** | **string** | The tenant name | [optional] 
**Total** | [**map[string]ConnectorStatistics**](ConnectorStatistics.md) | Aggregation of all statistics | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


