# \EventProcessingMatchingApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE18**](EventProcessingMatchingApi.md#DeleteUsingDELETE18) | **Delete** /api/v0/eventprocessing/matching-rule/{matchingRuleId} | Delete a MatchingRule
[**GetUsingGET17**](EventProcessingMatchingApi.md#GetUsingGET17) | **Get** /api/v0/eventprocessing/matching-rule/{matchingRuleId} | Retrieve a MatchingRule
[**ListUsingGET20**](EventProcessingMatchingApi.md#ListUsingGET20) | **Get** /api/v0/eventprocessing/matching-rule | Retrieve the list of all the MatchingRules or get a MatchingRule by its name
[**PostUsingPOST15**](EventProcessingMatchingApi.md#PostUsingPOST15) | **Post** /api/v0/eventprocessing/matching-rule | Create a MatchingRule
[**TestUsingPOST4**](EventProcessingMatchingApi.md#TestUsingPOST4) | **Post** /api/v0/eventprocessing/matching-rule/test | Test a JsonLogic predicate with some data sample
[**UpdateUsingPUT9**](EventProcessingMatchingApi.md#UpdateUsingPUT9) | **Put** /api/v0/eventprocessing/matching-rule/{matchingRuleId} | Update a MatchingRule



## DeleteUsingDELETE18

> DeleteUsingDELETE18(ctx, matchingRuleId, xAPIKEY)

Delete a MatchingRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleId** | **string**| id of the MatchingRule to delete | 
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


## GetUsingGET17

> MatchingRule GetUsingGET17(ctx, matchingRuleId, xAPIKEY)

Retrieve a MatchingRule

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleId** | **string**| id of the MatchingRule to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**MatchingRule**](MatchingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET20

> []MatchingRule ListUsingGET20(ctx, xAPIKEY, optional)

Retrieve the list of all the MatchingRules or get a MatchingRule by its name

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET20Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET20Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the MatchingRule to retrieve | 

### Return type

[**[]MatchingRule**](MatchingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST15

> MatchingRule PostUsingPOST15(ctx, xAPIKEY, optional)

Create a MatchingRule

Total number of MatchingRules is limited. Contact the commercial team or see developer guide to get more information.<br><br>Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PostUsingPOST15Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUsingPOST15Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **matchRule** | [**optional.Interface of MatchingRule**](MatchingRule.md)| MatchingRule to add | 

### Return type

[**MatchingRule**](MatchingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST4

> DataMatchResult TestUsingPOST4(ctx, xAPIKEY, optional)

Test a JsonLogic predicate with some data sample

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***TestUsingPOST4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestUsingPOST4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataMatchTest** | [**optional.Interface of DataMatchTest**](DataMatchTest.md)| jsonLogic predicate and json data of the evaluation test you want to perform | 

### Return type

[**DataMatchResult**](DataMatchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT9

> UpdateUsingPUT9(ctx, matchingRuleId, xAPIKEY, optional)

Update a MatchingRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleId** | **string**| id of the MatchingRule to update | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUsingPUT9Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsingPUT9Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **matchingRule** | [**optional.Interface of MatchingRule**](MatchingRule.md)| updated MatchingRule | 

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

