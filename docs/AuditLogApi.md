# \AuditLogApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchUsingGET**](AuditLogApi.md#SearchUsingGET) | **Get** /api/v0/auditlog/messages | Retrieve messages available in your AuditLog



## SearchUsingGET

> []AuditLogMessage SearchUsingGET(ctx, xAPIKEY, optional)

Retrieve messages available in your AuditLog

Restricted to API keys with at least one of the following roles : LOGS_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SearchUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SearchUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **optional.String**| Search for messages after this date. Use ISO-8601 normalization. | 
 **to** | **optional.String**| Search for messages before this date. Use ISO-8601 normalization. | 
 **offset** | **optional.Int32**| Offset from the first result you want to fetch. offset + limit should not exceed 10.000. | [default to 0]
 **limit** | **optional.Int32**| Maximum amount of messages to return. limit should not exceed 1.000. | [default to 100]
 **sort** | **optional.String**| Sort order, based on timestamp field of the AuditLog message. | [default to desc]
 **filters** | **optional.String**| Filter query based on parameter name &#x3D; field path and parameter value &#x3D; value to search for : {fiedlName}&#x3D;{value}. You can put several filters that way, they will all be treated with an AND operator. Common filters field names you can use : level, category, subcategory, type, source.deviceId, source.nodeId, description, detailedDescription. e.g. : level&#x3D;error&amp;source.deviceId&#x3D;urn:lora:0E5EAB0ABCD00000 | 
 **any** | [**optional.Interface of []string**](string.md)| Search for AuditLog Messages where any of the fields contains all these values. | 

### Return type

[**[]AuditLogMessage**](AuditLogMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

