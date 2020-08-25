# \EventProcessingGeozoneApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE17**](EventProcessingGeozoneApi.md#DeleteUsingDELETE17) | **Delete** /api/v0/eventprocessing/geozones/{zoneId} | delete a geographic zone
[**GetUsingGET16**](EventProcessingGeozoneApi.md#GetUsingGET16) | **Get** /api/v0/eventprocessing/geozones/{zoneId} | retrieve a geozone from repository
[**ListUsingGET19**](EventProcessingGeozoneApi.md#ListUsingGET19) | **Get** /api/v0/eventprocessing/geozones | retrieve paginated list of geozones
[**SaveUsingPUT1**](EventProcessingGeozoneApi.md#SaveUsingPUT1) | **Put** /api/v0/eventprocessing/geozones/{zoneId} | Save a geographic zone



## DeleteUsingDELETE17

> DeleteUsingDELETE17(ctx, zoneId, xAPIKEY)

delete a geographic zone

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| the id of the zone to delete | 
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


## GetUsingGET16

> GeozoneContainer GetUsingGET16(ctx, zoneId, xAPIKEY)

retrieve a geozone from repository

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| the user-defined id of the zone to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**GeozoneContainer**](GeozoneContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET19

> []GeozoneContainer ListUsingGET19(ctx, xAPIKEY, optional)

retrieve paginated list of geozones

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET19Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET19Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| The requested page number. Starts from 0. | [default to 0]
 **size** | **optional.Int32**| The maximum number of items per page. Must be between 1 and 1000. | [default to 20]

### Return type

[**[]GeozoneContainer**](GeozoneContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveUsingPUT1

> SaveUsingPUT1(ctx, zoneId, xAPIKEY, optional)

Save a geographic zone

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string**| the id of the zone to save | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SaveUsingPUT1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SaveUsingPUT1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **geozoneContainerBase** | [**optional.Interface of GeozoneContainerBase**](GeozoneContainerBase.md)| JSON geo zone object to add | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

