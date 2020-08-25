# \DecodersApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsingGET22**](DecodersApi.md#ListUsingGET22) | **Get** /api/v0/decoders | Retrieve the list of all decoders (binary, csv, js)



## ListUsingGET22

> []PayloadDescription ListUsingGET22(ctx, xAPIKEY, optional)

Retrieve the list of all decoders (binary, csv, js)

Restricted to API keys with at least one of the following roles : DATA_PROCESSING_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsingGET22Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsingGET22Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **tags** | [**optional.Interface of []string**](string.md)| target filtering tags | 

### Return type

[**[]PayloadDescription**](PayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

