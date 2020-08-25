# MatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPredicate** | **map[string]interface{}** | the JsonLogic (http://jsonlogic.com/) pattern matching that will trigger an event for each new data that match this predicate. This JsonLogic predicate needs to return a boolean. | [optional] 
**Enabled** | **bool** | activate or not the MatchingRule. Default is false. | [optional] 
**Id** | **string** | id of the MatchingRule. Should be null when used for POST. | [optional] 
**Name** | **string** | name of the MatchingRule. Must be unique. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


