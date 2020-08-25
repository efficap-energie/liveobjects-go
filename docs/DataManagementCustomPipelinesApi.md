# \DataManagementCustomPipelinesApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE11**](DataManagementCustomPipelinesApi.md#DeleteUsingDELETE11) | **Delete** /api/v0/pipelines/{pipelineId} | Delete a DataMessage pipeline
[**GetUsingGET12**](DataManagementCustomPipelinesApi.md#GetUsingGET12) | **Get** /api/v0/pipelines/{pipelineId} | Retrieve a DataMessage pipeline
[**ListUsingGET13**](DataManagementCustomPipelinesApi.md#ListUsingGET13) | **Get** /api/v0/pipelines | Retrieve the list of DataMessage pipelines, ordered by priorityLevel
[**PostUsingPOST12**](DataManagementCustomPipelinesApi.md#PostUsingPOST12) | **Post** /api/v0/pipelines | Create a DataMessage pipeline
[**UpdateUsingPUT6**](DataManagementCustomPipelinesApi.md#UpdateUsingPUT6) | **Put** /api/v0/pipelines/{pipelineId} | Update a DataMessage pipeline



## DeleteUsingDELETE11

> DeleteUsingDELETE11(ctx, pipelineId, xAPIKEY)

Delete a DataMessage pipeline

Usage of this API will be reported in your access log under 'data_pipeline' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string**| id of the pipeline to delete | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET12

> Pipeline GetUsingGET12(ctx, pipelineId, xAPIKEY)

Retrieve a DataMessage pipeline

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string**| id of the pipeline to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET13

> []Pipeline ListUsingGET13(ctx, xAPIKEY)

Retrieve the list of DataMessage pipelines, ordered by priorityLevel

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST12

> Pipeline PostUsingPOST12(ctx, xAPIKEY, optional)

Create a DataMessage pipeline

Usage of this API will be reported in your access log under 'data_pipeline' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PostUsingPOST12Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUsingPOST12Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **pipelineDescription** | [**optional.Interface of Pipeline**](Pipeline.md)| pipeline to create | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT6

> Pipeline UpdateUsingPUT6(ctx, pipelineId, xAPIKEY, optional)

Update a DataMessage pipeline

Usage of this API will be reported in your access log under 'data_pipeline' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string**| id of the pipeline to update | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUsingPUT6Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsingPUT6Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **pipelineDescription** | [**optional.Interface of Pipeline**](Pipeline.md)| pipeline to update | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

