# \DeprecatedTriggersAndActionsTestApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestHttpPushUsingPOST**](DeprecatedTriggersAndActionsTestApi.md#TestHttpPushUsingPOST) | **Post** /api/v0/event2action/test/http-push | Post an http request for testing a webhook



## TestHttpPushUsingPOST

> HttpPushTestResult TestHttpPushUsingPOST(ctx, xAPIKEY, optional)

Post an http request for testing a webhook

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***TestHttpPushUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestHttpPushUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **req** | [**optional.Interface of HttpPushWebhookTest**](HttpPushWebhookTest.md)| Http push request to be tested | 

### Return type

[**HttpPushTestResult**](HttpPushTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

