# \BusManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST3**](BusManagementApi.md#CreateUsingPOST3) | **Post** /api/v0/topics/fifo | Create a FIFO
[**DeleteUsingDELETE20**](BusManagementApi.md#DeleteUsingDELETE20) | **Delete** /api/v0/topics/fifo/{fifoName} | Delete a FIFO
[**FifoPublishUsingPOST**](BusManagementApi.md#FifoPublishUsingPOST) | **Post** /api/v0/topics/fifo/{fifoName} | Publish a message into a FIFO
[**GetUsingGET19**](BusManagementApi.md#GetUsingGET19) | **Get** /api/v0/topics/fifo/{fifoName} | Get a FIFO
[**ListFifoTopicsUsingGET**](BusManagementApi.md#ListFifoTopicsUsingGET) | **Get** /api/v0/topics/fifo | List all FIFOs



## CreateUsingPOST3

> CreateUsingPOST3(ctx, xAPIKEY, optional)

Create a FIFO

Usage of this API will be reported in your access log under 'routing' category.<br><br>Restricted to API keys with at least one of the following roles : BUS_CONFIG_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateUsingPOST3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUsingPOST3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **requestBody** | [**optional.Interface of FifoCreateReqWeb**](FifoCreateReqWeb.md)| body for create a new fifo | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE20

> DeleteUsingDELETE20(ctx, fifoName, xAPIKEY)

Delete a FIFO

Usage of this API will be reported in your access log under 'routing' category.<br><br>Restricted to API keys with at least one of the following roles : BUS_CONFIG_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fifoName** | **string**| name of the fifo. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FifoPublishUsingPOST

> bool FifoPublishUsingPOST(ctx, fifoName, xAPIKEY, optional)

Publish a message into a FIFO

Restricted to API keys with at least one of the following roles : BUS_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fifoName** | **string**| name of the fifo. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***FifoPublishUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a FifoPublishUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **payload** | **optional.String**| body for publish a message into your FIFO topic (your body is specific at your topic) | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, 
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET19

> FifoTopic GetUsingGET19(ctx, fifoName, xAPIKEY)

Get a FIFO

Restricted to API keys with at least one of the following roles : BUS_CONFIG_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fifoName** | **string**| name of the fifo. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**FifoTopic**](FifoTopic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFifoTopicsUsingGET

> PageableFifoTopic ListFifoTopicsUsingGET(ctx, xAPIKEY, optional)

List all FIFOs

Restricted to API keys with at least one of the following roles : BUS_CONFIG_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListFifoTopicsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListFifoTopicsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**PageableFifoTopic**](Pageable«FifoTopic».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

