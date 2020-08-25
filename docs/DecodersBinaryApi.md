# \DecodersBinaryApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUsingPUT8**](DecodersBinaryApi.md#ActivateUsingPUT8) | **Put** /api/v0/decoders/binary/{decoderId}/enabled | Activate or deactivate a decoder
[**DeleteUsingDELETE8**](DecodersBinaryApi.md#DeleteUsingDELETE8) | **Delete** /api/v0/decoders/binary/{decoderId} | Delete a binary decoder
[**GetUsingGET10**](DecodersBinaryApi.md#GetUsingGET10) | **Get** /api/v0/decoders/binary/{decoderId} | Retrieve a binary decoder
[**ListUsingGET10**](DecodersBinaryApi.md#ListUsingGET10) | **Get** /api/v0/decoders/binary | Retrieve the list of binary decoders
[**PostUsingPOST10**](DecodersBinaryApi.md#PostUsingPOST10) | **Post** /api/v0/decoders/binary | Create a binary decoder
[**PutUsingPUT2**](DecodersBinaryApi.md#PutUsingPUT2) | **Put** /api/v0/decoders/binary/{decoderId} | Update a binary decoder
[**TestUsingPOST2**](DecodersBinaryApi.md#TestUsingPOST2) | **Post** /api/v0/decoders/binary/test | Test a binary decoder format with an encoded payload



## ActivateUsingPUT8

> ActivateUsingPUT8(ctx, decoderId, xAPIKEY, optional)

Activate or deactivate a decoder

Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the binary decoder to activate or deactivate | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ActivateUsingPUT8Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ActivateUsingPUT8Opts struct


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


## DeleteUsingDELETE8

> DeleteUsingDELETE8(ctx, decoderId, xAPIKEY)

Delete a binary decoder

Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the binary decoder to delete | 
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


## GetUsingGET10

> BinaryPayloadDescription GetUsingGET10(ctx, decoderId, xAPIKEY)

Retrieve a binary decoder

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the binary decoder to retrieve | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET10

> []BinaryPayloadDescription ListUsingGET10(ctx, xAPIKEY, optional)

Retrieve the list of binary decoders

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET10Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET10Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | [**optional.Interface of []string**](string.md)| target filtering tags | 

### Return type

[**[]BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST10

> BinaryPayloadDescription PostUsingPOST10(ctx, xAPIKEY, optional)

Create a binary decoder

The number of binary decoders is limited to 100.<br><br>Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PostUsingPOST10Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PostUsingPOST10Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **binaryPayloadDescription** | [**optional.Interface of BinaryPayloadDescription**](BinaryPayloadDescription.md)| Binary decoder description | 

### Return type

[**BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUsingPUT2

> BinaryPayloadDescription PutUsingPUT2(ctx, decoderId, xAPIKEY, optional)

Update a binary decoder

Usage of this API will be reported in your access log under 'data_decoder' category.<br><br>Restricted to API keys with at least one of the following roles : DATA_PROCESSING_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string**| id of the binary decoder to update | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***PutUsingPUT2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a PutUsingPUT2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **binaryPayloadDescription** | [**optional.Interface of BinaryPayloadDescription**](BinaryPayloadDescription.md)| Binary decoder description | 

### Return type

[**BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST2

> PayloadDescriptionTestResult TestUsingPOST2(ctx, xAPIKEY, optional)

Test a binary decoder format with an encoded payload

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***TestUsingPOST2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a TestUsingPOST2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataDecodingTestRequest** | [**optional.Interface of BinaryPayloadDescriptionTestRequest**](BinaryPayloadDescriptionTestRequest.md)| Binary format and encoded payload on which decoding shall be peformed | 

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

