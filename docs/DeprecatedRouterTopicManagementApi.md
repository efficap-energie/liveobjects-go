# \DeprecatedRouterTopicManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublishUsingPOST**](DeprecatedRouterTopicManagementApi.md#PublishUsingPOST) | **Post** /api/v0/topics/router/{routingKey} | Publishing a message into a ROUTER topic



## PublishUsingPOST

> bool PublishUsingPOST(ctx, routingKey, xAPIKEY, optional)

Publishing a message into a ROUTER topic

Restricted to API keys with at least one of the following roles : BUS_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingKey** | **string**| Routing Key | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PublishUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PublishUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **payload** | **optional.String**| payload | 

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

