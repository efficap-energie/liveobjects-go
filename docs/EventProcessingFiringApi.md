# \EventProcessingFiringApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE16**](EventProcessingFiringApi.md#DeleteUsingDELETE16) | **Delete** /api/v0/eventprocessing/firing-rule/{firingRuleId} | Delete a FiringRule
[**GetFiringGuardUsingGET**](EventProcessingFiringApi.md#GetFiringGuardUsingGET) | **Get** /api/v0/eventprocessing/firing-guard/{firingGuardId} | Get a FiringGuard
[**GetFiringGuardsUsingPOST**](EventProcessingFiringApi.md#GetFiringGuardsUsingPOST) | **Post** /api/v0/eventprocessing/firing-guard/search | Get FiringGuards linked to a FiringRule, and where FiringGuards selection criteria match.
[**GetUsingGET15**](EventProcessingFiringApi.md#GetUsingGET15) | **Get** /api/v0/eventprocessing/firing-rule/{firingRuleId} | Retrieve a FiringRule
[**ListUsingGET18**](EventProcessingFiringApi.md#ListUsingGET18) | **Get** /api/v0/eventprocessing/firing-rule | Retrieve the list of all the FiringRules or get a FiringRule by its name
[**PostUsingPOST14**](EventProcessingFiringApi.md#PostUsingPOST14) | **Post** /api/v0/eventprocessing/firing-rule | Create a FiringRule
[**RemoveFiringGuardUsingDELETE**](EventProcessingFiringApi.md#RemoveFiringGuardUsingDELETE) | **Delete** /api/v0/eventprocessing/firing-guard/{firingGuardId} | Remove a FiringGuard
[**RemoveFiringGuardsUsingDELETE**](EventProcessingFiringApi.md#RemoveFiringGuardsUsingDELETE) | **Delete** /api/v0/eventprocessing/firing-guard | Remove the FiringGuards linked to FiringRule, and where FiringGuards selection criteria match.
[**UpdateUsingPUT8**](EventProcessingFiringApi.md#UpdateUsingPUT8) | **Put** /api/v0/eventprocessing/firing-rule/{firingRuleId} | Update a FiringRule



## DeleteUsingDELETE16

> DeleteUsingDELETE16(ctx, firingRuleId, xAPIKEY)

Delete a FiringRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingRuleId** | **string**| id of the FiringRule to delete | 
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


## GetFiringGuardUsingGET

> FiringGuard GetFiringGuardUsingGET(ctx, firingGuardId, xAPIKEY)

Get a FiringGuard

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingGuardId** | **string**| id of the FiringGuard | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**FiringGuard**](FiringGuard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiringGuardsUsingPOST

> []FiringGuard GetFiringGuardsUsingPOST(ctx, xAPIKEY, optional)

Get FiringGuards linked to a FiringRule, and where FiringGuards selection criteria match.

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetFiringGuardsUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetFiringGuardsUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **searchParam** | [**optional.Interface of FiringGuardGetRequest**](FiringGuardGetRequest.md)| Search parameters to select FiringGuard | 

### Return type

[**[]FiringGuard**](FiringGuard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET15

> FiringRule GetUsingGET15(ctx, firingRuleId, xAPIKEY)

Retrieve a FiringRule

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingRuleId** | **string**| id of the FiringRule to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**FiringRule**](FiringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET18

> []FiringRule ListUsingGET18(ctx, xAPIKEY, optional)

Retrieve the list of all the FiringRules or get a FiringRule by its name

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET18Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET18Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the FiringRule to retrieve | 

### Return type

[**[]FiringRule**](FiringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST14

> FiringRule PostUsingPOST14(ctx, xAPIKEY, optional)

Create a FiringRule

Total number of FiringRules is limited. Contact the commercial team or see developer guide to get more information.<br><br>Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PostUsingPOST14Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUsingPOST14Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firingRule** | [**optional.Interface of FiringRule**](FiringRule.md)| FiringRule to add | 

### Return type

[**FiringRule**](FiringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFiringGuardUsingDELETE

> RemoveFiringGuardUsingDELETE(ctx, firingGuardId, xAPIKEY)

Remove a FiringGuard

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingGuardId** | **string**| id of the FiringGuard to remove | 
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


## RemoveFiringGuardsUsingDELETE

> int32 RemoveFiringGuardsUsingDELETE(ctx, xAPIKEY, optional)

Remove the FiringGuards linked to FiringRule, and where FiringGuards selection criteria match.

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***RemoveFiringGuardsUsingDELETEOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RemoveFiringGuardsUsingDELETEOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **resetParams** | [**optional.Interface of FiringGuardResetRequest**](FiringGuardResetRequest.md)| search parameters to select FiringGuard to remove | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT8

> UpdateUsingPUT8(ctx, firingRuleId, xAPIKEY, optional)

Update a FiringRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingRuleId** | **string**| id of the FiringRule to update | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUsingPUT8Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsingPUT8Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **firingRule** | [**optional.Interface of FiringRule**](FiringRule.md)| updated FiringRule | 

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

