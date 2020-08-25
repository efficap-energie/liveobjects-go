# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **string** | pipeline&#39;s creation date (ISO-8601). Should be null for post and update action. | [optional] 
**Description** | **string** | pipeline description (length limit : 2000). | [optional] 
**Enabled** | **bool** | pipeline activation. Default is false. | 
**Filter** | [**PipelineFilter**](PipelineFilter.md) |  | [optional] 
**Id** | **string** | id of the pipeline. Should be null when used for POST. Required for update. | [optional] 
**Name** | **string** | pipeline name (length limit : 1000). | 
**PriorityLevel** | **int64** | pipeline&#39;s priority level. When several pipelines match the input filter, then the priority is given to the pipeline with lowest priorityLevel. If several pipelines match filter and the share the same priorityLevel, then the oldest pipeline is chosen (based on &#39;created&#39; field). | 
**Steps** | [**[]PipelineStep**](PipelineStep.md) | list of steps of the pipelines. Should contain at least 1 step. Available steps are defined in developer&#39;s guide. | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


