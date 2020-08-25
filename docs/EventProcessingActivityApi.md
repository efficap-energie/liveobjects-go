# \EventProcessingActivityApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE14**](EventProcessingActivityApi.md#DeleteUsingDELETE14) | **Delete** /api/v0/eventprocessing/activity/rules/{activityRuleId} | Delete an ActivityRule
[**GetStatesUsingGET**](EventProcessingActivityApi.md#GetStatesUsingGET) | **Get** /api/v0/eventprocessing/activity/states | Retrieve the list of all the ActivityStates linked to a specific device and/or rule
[**GetUsingGET13**](EventProcessingActivityApi.md#GetUsingGET13) | **Get** /api/v0/eventprocessing/activity/rules/{activityRuleId} | Retrieve an ActivityRule
[**ListUsingGET16**](EventProcessingActivityApi.md#ListUsingGET16) | **Get** /api/v0/eventprocessing/activity/rules | Retrieve the list of all the ActivityRules or get an ActivityRule by its name
[**MuteUsingPUT**](EventProcessingActivityApi.md#MuteUsingPUT) | **Put** /api/v0/eventprocessing/activity/states/mute | Mute or reset nextAlarm of ActivityStates targeted by a specific deviceId/activityRuleId
[**PostUsingPOST13**](EventProcessingActivityApi.md#PostUsingPOST13) | **Post** /api/v0/eventprocessing/activity/rules | Create an ActivityRule
[**UpdateUsingPUT7**](EventProcessingActivityApi.md#UpdateUsingPUT7) | **Put** /api/v0/eventprocessing/activity/rules/{activityRuleId} | Update an ActivityRule



## DeleteUsingDELETE14

> DeleteUsingDELETE14(ctx, activityRuleId, xAPIKEY)

Delete an ActivityRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityRuleId** | **string**| id of the ActivityRule to delete | 
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


## GetStatesUsingGET

> []ActivityState GetStatesUsingGET(ctx, xAPIKEY, optional)

Retrieve the list of all the ActivityStates linked to a specific device and/or rule

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetStatesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetStatesUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **deviceId** | **optional.String**| id of the device targeted by the states to retrieve. At least one of deviceId/activityRuleId must be set. | 
 **activityRuleId** | **optional.String**| id of the rule targeted by the states to retrieve. At least one of deviceId/activityRuleId must be set. | 
 **limit** | **optional.Int32**| when listing by activityRuleId, thousands of devices could be targeted. Indicates the number of AvtivityStates to return (one per targeted device). Default is 20, max is 1000. To get next results use bookmarkDeviceId field. | 
 **bookmarkDeviceId** | **optional.String**| &#39;bookmark&#39; of previous listing by activityRuleId request : this is the last deviceId retrieved in previous request. If null, first ActivtyStates results will be returned. | 

### Return type

[**[]ActivityState**](ActivityState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET13

> ActivityRule GetUsingGET13(ctx, activityRuleId, xAPIKEY)

Retrieve an ActivityRule

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityRuleId** | **string**| id of the ActivityRule to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**ActivityRule**](ActivityRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET16

> []ActivityRule ListUsingGET16(ctx, xAPIKEY, optional)

Retrieve the list of all the ActivityRules or get an ActivityRule by its name

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET16Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET16Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **optional.String**| name of the ActivityRule to retrieve | 

### Return type

[**[]ActivityRule**](ActivityRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteUsingPUT

> int64 MuteUsingPUT(ctx, xAPIKEY, nextAlarmRequest)

Mute or reset nextAlarm of ActivityStates targeted by a specific deviceId/activityRuleId

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
**nextAlarmRequest** | [**ActivityStateMuteRequest**](ActivityStateMuteRequest.md)| nextAlarmRequest | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST13

> ActivityRule PostUsingPOST13(ctx, xAPIKEY, optional)

Create an ActivityRule

Total number of ActivityRules is limited. Contact the commercial team or see developer guide to get more information.<br><br>Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PostUsingPOST13Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUsingPOST13Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activityRule** | [**optional.Interface of ActivityRule**](ActivityRule.md)| ActivityRule to add | 

### Return type

[**ActivityRule**](ActivityRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT7

> UpdateUsingPUT7(ctx, activityRuleId, xAPIKEY, optional)

Update an ActivityRule

Usage of this API will be reported in your access log under 'alarming' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityRuleId** | **string**| id of the ActivityRule to update | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUsingPUT7Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUsingPUT7Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **activityRule** | [**optional.Interface of ActivityRule**](ActivityRule.md)| updated ActivityRule | 

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

