# \DecodersCSVApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUsingPUT9**](DecodersCSVApi.md#ActivateUsingPUT9) | **Put** /api/v0/decoders/csv/{decoderId}/enabled | Activate or deactivate a decoder
[**DeleteUsingDELETE10**](DecodersCSVApi.md#DeleteUsingDELETE10) | **Delete** /api/v0/decoders/csv/{decoderId} | Delete a csv decoder
[**GetUsingGET11**](DecodersCSVApi.md#GetUsingGET11) | **Get** /api/v0/decoders/csv/{decoderId} | Retrieve a csv decoder
[**ListUsingGET12**](DecodersCSVApi.md#ListUsingGET12) | **Get** /api/v0/decoders/csv | Retrieve the list of csv decoders filtered by tags
[**PostUsingPOST11**](DecodersCSVApi.md#PostUsingPOST11) | **Post** /api/v0/decoders/csv | Create a csv decoder
[**PutUsingPUT3**](DecodersCSVApi.md#PutUsingPUT3) | **Put** /api/v0/decoders/csv/{decoderId} | Update a csv decoder
[**TestUsingPOST3**](DecodersCSVApi.md#TestUsingPOST3) | **Post** /api/v0/decoders/csv/test | Test a csv decoder format with an encoded payload



## ActivateUsingPUT9

> ActivateUsingPUT9(ctx, decoderId, xAPIKEY, optional)

Activate or deactivate a decoder

Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the csv decoder to activate or deactivate | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ActivateUsingPUT9Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ActivateUsingPUT9Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **enabled** | **optional.Bool**| true to activate, false otherwise | 

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


## DeleteUsingDELETE10

> DeleteUsingDELETE10(ctx, decoderId, xAPIKEY)

Delete a csv decoder

Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the csv decoder to delete | 
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


## GetUsingGET11

> CsvPayloadDescription GetUsingGET11(ctx, decoderId, xAPIKEY)

Retrieve a csv decoder

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the csv decoder to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET12

> []CsvPayloadDescription ListUsingGET12(ctx, xAPIKEY, optional)

Retrieve the list of csv decoders filtered by tags

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET12Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET12Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | [**optional.Interface of []string**](string.md)| target filtering tags | 

### Return type

[**[]CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST11

> CsvPayloadDescription PostUsingPOST11(ctx, xAPIKEY, optional)

Create a csv decoder

The number of csv decoders is limited to 100.<br><br>Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PostUsingPOST11Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUsingPOST11Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **csvPayloadDescription** | [**optional.Interface of CsvPayloadDescription**](CsvPayloadDescription.md)| Csv decoder description | 

### Return type

[**CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUsingPUT3

> CsvPayloadDescription PutUsingPUT3(ctx, decoderId, xAPIKEY, optional)

Update a csv decoder

Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the csv decoder to update | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PutUsingPUT3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutUsingPUT3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **csvPayloadDescription** | [**optional.Interface of CsvPayloadDescription**](CsvPayloadDescription.md)| Csv decoder description | 

### Return type

[**CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST3

> PayloadDescriptionTestResult TestUsingPOST3(ctx, xAPIKEY, optional)

Test a csv decoder format with an encoded payload

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***TestUsingPOST3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestUsingPOST3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataDecodingTestRequest** | [**optional.Interface of CsvPayloadDescriptionTestRequest**](CsvPayloadDescriptionTestRequest.md)| Csv format and encoded payload on which decoding shall be peformed | 

### Return type

[**PayloadDescriptionTestResult**](PayloadDescriptionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

