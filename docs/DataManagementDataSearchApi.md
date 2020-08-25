# \DataManagementDataSearchApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DslQueryHitsOnlyUsingPOST**](DataManagementDataSearchApi.md#DslQueryHitsOnlyUsingPOST) | **Post** /api/v0/data/search/hits | Query an Elasticsearch Domain Specific Language request and get only hits result
[**DslQueryUsingPOST**](DataManagementDataSearchApi.md#DslQueryUsingPOST) | **Post** /api/v0/data/search | Query an Elasticsearch Domain Specific Language request



## DslQueryHitsOnlyUsingPOST

> []DataStoredWeb DslQueryHitsOnlyUsingPOST(ctx, xAPIKEY, optional)

Query an Elasticsearch Domain Specific Language request and get only hits result

return a json array of stored data that match the query.<br><br>Restricted to API keys with at least one of the following roles : DATA_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***DslQueryHitsOnlyUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DslQueryHitsOnlyUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dslRequest** | **optional.Map[string]interface{}**| elasticsearch DSL request | 

### Return type

[**[]DataStoredWeb**](DataStoredWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DslQueryUsingPOST

> map[string]interface{} DslQueryUsingPOST(ctx, xAPIKEY, optional)

Query an Elasticsearch Domain Specific Language request

return the json serialization of an Elasticsearch SearchResponse for the query<br><br>Restricted to API keys with at least one of the following roles : DATA_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***DslQueryUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DslQueryUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dslRequest** | **optional.Map[string]interface{}**| elasticsearch DSL request | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

