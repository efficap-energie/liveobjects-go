# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignStatus** | **string** |  | [optional] [readonly] 
**Created** | **string** |  | [optional] [readonly] 
**Description** | **string** | Campaign description | [optional] 
**Id** | **string** | campaign id | 
**Name** | **string** | campaign name | 
**NumberOfTargets** | **int32** | number of targets after deduplication | [optional] [readonly] 
**Operations** | [**[]CampaignOperation**](CampaignOperation.md) | requested operations | 
**Options** | [**CampaignOptions**](CampaignOptions.md) |  | [optional] 
**Planning** | [**CampaignPlanning**](CampaignPlanning.md) |  | 
**Target** | [**DeviceSelector**](DeviceSelector.md) |  | 
**TotalTargetsPerStatus** | [**CampaignStatsPerStatus**](CampaignStatsPerStatus.md) |  | [optional] 
**Updated** | **string** |  | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


