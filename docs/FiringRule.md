# FiringRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationKeys** | **[]string** | the list of jsonPath in the Data that will define on which group of data this FiringRule should be set. For instance &#39;streamId&#39;, &#39;metadata.source&#39;, &#39;value.type&#39;. | [optional] 
**Enabled** | **bool** | activate or not the FiringRule. Default is false. | [optional] 
**FiringType** | **string** | define the type of firing mechanism : ONCE, SLEEP, or ALWAYS | 
**Id** | **string** | id of the FiringRule. Should be null when used for POST. | [optional] 
**MatchingRuleIds** | **[]string** | the list of MatchingRule ids that will be handeld by this FiringRule. | [optional] 
**Name** | **string** | name of the FiringRule. Must be unique. | 
**SleepDuration** | **string** | sleep duration of the FiringRule. Is defined as a ISO-8601 Period string, restricted to days, hours, minutes and seconds. 1 day will always be equivalent to 24H, regardless of daylight saving time. eg. : &#39;P1D&#39;, &#39;PT30M&#39;. Must be set only for &#39;SLEEP&#39; FiringType. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


