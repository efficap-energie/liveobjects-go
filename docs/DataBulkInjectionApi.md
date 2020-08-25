# \DataBulkInjectionApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataBulkUsingPOST**](DataBulkInjectionApi.md#AddDataBulkUsingPOST) | **Post** /api/v0/data/bulk | Insert a bulk of new Data



## AddDataBulkUsingPOST

> string AddDataBulkUsingPOST(ctx, xAPIKEY, optional)

Insert a bulk of new Data

Restricted to API keys with at least one of the following roles : DATA_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***AddDataBulkUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddDataBulkUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dataBulk** | [**optional.Interface of []DataBulkItemWeb**](DataBulkItemWeb.md)| Bulk of Data to add | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

