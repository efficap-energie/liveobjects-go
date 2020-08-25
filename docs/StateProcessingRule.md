# StateProcessingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | activate or not the StateProcessingRule. Default is false. | [optional] 
**FilterPredicate** | **map[string]interface{}** | the JsonLogic (http://jsonlogic.com/) filter predicate applied before state function | [optional] 
**Id** | **string** | id of the StateProcessingRule. Should be null when used for POST. | [optional] 
**Name** | **string** | name of the StateProcessingRule. Must be unique. | 
**StateFunction** | **map[string]interface{}** | the JsonLogic (http://jsonlogic.com/) state function | [optional] 
**StateKeyPath** | **string** | the json path applied on dataMessage used to compute the stateKey | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


