# \EventProcessingContextApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearUsingDELETE**](EventProcessingContextApi.md#ClearUsingDELETE) | **Delete** /api/v0/eventprocessing/context | Delete all context entries
[**DeleteUsingDELETE15**](EventProcessingContextApi.md#DeleteUsingDELETE15) | **Delete** /api/v0/eventprocessing/context/{contextKey} | Delete a context key
[**GetUsingGET14**](EventProcessingContextApi.md#GetUsingGET14) | **Get** /api/v0/eventprocessing/context/{contextKey} | Retrieve a context
[**ListUsingGET17**](EventProcessingContextApi.md#ListUsingGET17) | **Get** /api/v0/eventprocessing/context | Retrieve the list of contexts with optional tag filtering
[**SaveUsingPUT**](EventProcessingContextApi.md#SaveUsingPUT) | **Put** /api/v0/eventprocessing/context/{contextKey} | Save a context



## ClearUsingDELETE

> int64 ClearUsingDELETE(ctx, xAPIKEY)

Delete all context entries

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE15

> DeleteUsingDELETE15(ctx, contextKey, xAPIKEY)

Delete a context key

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextKey** | **string**| id of the context to delete | 
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


## GetUsingGET14

> ContextContainer GetUsingGET14(ctx, contextKey, xAPIKEY)

Retrieve a context

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextKey** | **string**| id of the context to get | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**ContextContainer**](ContextContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET17

> []ContextContainer ListUsingGET17(ctx, xAPIKEY, optional)

Retrieve the list of contexts with optional tag filtering

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET17Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET17Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **optional.Int32**| The requested page number | [default to 0]
 **size** | **optional.Int32**| The maximum number of items per page | [default to 20]
 **tags** | [**optional.Interface of []string**](string.md)| Filtering tags | 

### Return type

[**[]ContextContainer**](ContextContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveUsingPUT

> ContextContainer SaveUsingPUT(ctx, contextKey, xAPIKEY, optional)

Save a context

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextKey** | **string**| id of the context to save | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SaveUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SaveUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **contextContainer** | [**optional.Interface of ContextContainer**](ContextContainer.md)| JSON context object to add | 

### Return type

[**ContextContainer**](ContextContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

