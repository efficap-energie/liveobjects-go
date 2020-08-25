# \EventProcessingStateProcessingApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE19**](EventProcessingStateProcessingApi.md#DeleteUsingDELETE19) | **Delete** /api/v0/eventprocessing/stateprocessing-rule/{stateProcessingRuleId} | Delete a StateProcessingRule
[**GetUsingGET18**](EventProcessingStateProcessingApi.md#GetUsingGET18) | **Get** /api/v0/eventprocessing/stateprocessing-rule/{stateProcessingRuleId} | Retrieve a StateProcessingRule
[**ListUsingGET21**](EventProcessingStateProcessingApi.md#ListUsingGET21) | **Get** /api/v0/eventprocessing/stateprocessing-rule | Retrieve the list of all the StateProcessingRules or get a StateProcessingRule by its name
[**PostUsingPOST16**](EventProcessingStateProcessingApi.md#PostUsingPOST16) | **Post** /api/v0/eventprocessing/stateprocessing-rule | Create a StateProcessingRule
[**TestUsingPOST5**](EventProcessingStateProcessingApi.md#TestUsingPOST5) | **Post** /api/v0/eventprocessing/stateprocessing-rule/test | test a  State Processing function
[**UpdateUsingPUT10**](EventProcessingStateProcessingApi.md#UpdateUsingPUT10) | **Put** /api/v0/eventprocessing/stateprocessing-rule/{stateProcessingRuleId} | Update a StateProcessingRule



## DeleteUsingDELETE19

> DeleteUsingDELETE19(ctx, stateProcessingRuleId, xAPIKEY)

Delete a StateProcessingRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateProcessingRuleId** | **string**| id of the StateProcessingRule to delete | 
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


## GetUsingGET18

> StateProcessingRule GetUsingGET18(ctx, stateProcessingRuleId, xAPIKEY)

Retrieve a StateProcessingRule

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateProcessingRuleId** | **string**| id of the StateProcessingRule to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**StateProcessingRule**](StateProcessingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET21

> []StateProcessingRule ListUsingGET21(ctx, xAPIKEY, optional)

Retrieve the list of all the StateProcessingRules or get a StateProcessingRule by its name

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET21Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET21Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the StateProcessingRules to retrieve | 

### Return type

[**[]StateProcessingRule**](StateProcessingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST16

> StateProcessingRule PostUsingPOST16(ctx, xAPIKEY, optional)

Create a StateProcessingRule

Total number of StateProcessingRule is limited. Contact the commercial team or see developer guide to get more information.<br><br>Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PostUsingPOST16Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUsingPOST16Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **stateProcessingRule** | [**optional.Interface of StateProcessingRule**](StateProcessingRule.md)| StateProcessingRule to add | 

### Return type

[**StateProcessingRule**](StateProcessingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST5

> StateProcessingFunctionTestResult TestUsingPOST5(ctx, xAPIKEY, optional)

test a  State Processing function

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***TestUsingPOST5Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestUsingPOST5Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **testRequest** | [**optional.Interface of StateProcessingFunctionTest**](StateProcessingFunctionTest.md)| State Processing function test request | 

### Return type

[**StateProcessingFunctionTestResult**](StateProcessingFunctionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT10

> UpdateUsingPUT10(ctx, stateProcessingRuleId, xAPIKEY, optional)

Update a StateProcessingRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateProcessingRuleId** | **string**| id of the StateProcessingRule to update | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUsingPUT10Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsingPUT10Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **stateProcessingRule** | [**optional.Interface of StateProcessingRule**](StateProcessingRule.md)| updated StateProcessingRule | 

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

