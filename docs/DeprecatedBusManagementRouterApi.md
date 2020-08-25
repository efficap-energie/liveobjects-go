# \DeprecatedBusManagementRouterApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBindingUsingPOST**](DeprecatedBusManagementRouterApi.md#AddBindingUsingPOST) | **Post** /api/v0/bindings | Add a binding for a FIFO queue
[**DeleteBindingUsingDELETE**](DeprecatedBusManagementRouterApi.md#DeleteBindingUsingDELETE) | **Delete** /api/v0/bindings/{routingKeyFilter}/{fifoName} | Delete a binding
[**ListBindingsUsingGET**](DeprecatedBusManagementRouterApi.md#ListBindingsUsingGET) | **Get** /api/v0/bindings | List registered bindings



## AddBindingUsingPOST

> AddBindingUsingPOST(ctx, xAPIKEY, optional)

Add a binding for a FIFO queue

Usage of this API will be reported in your access log under 'routing' category.<br><br>Restricted to API keys with at least one of the following roles : BUS_CONFIG_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***AddBindingUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddBindingUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **req** | [**optional.Interface of CreateFifoBindingRequest**](CreateFifoBindingRequest.md)| body of create FIFO bindings | 

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


## DeleteBindingUsingDELETE

> DeleteBindingUsingDELETE(ctx, routingKeyFilter, fifoName, xAPIKEY)

Delete a binding

Usage of this API will be reported in your access log under 'routing' category.<br><br>Restricted to API keys with at least one of the following roles : BUS_CONFIG_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingKeyFilter** | **string**| routing key filter. Routing key must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_:~#\\.\\*]+&lt;/code&gt; (max 1000 characters) | 
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


## ListBindingsUsingGET

> PageableFifoBinding ListBindingsUsingGET(ctx, xAPIKEY, optional)

List registered bindings

Restricted to API keys with at least one of the following roles : BUS_CONFIG_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListBindingsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListBindingsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**PageableFifoBinding**](Pageable«FifoBinding».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

