# \PartnerDataPushApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataPushUsingPOST**](PartnerDataPushApi.md#DataPushUsingPOST) | **Post** /api/v0/partners/data/streams | Push data



## DataPushUsingPOST

> string DataPushUsingPOST(ctx, authorization, data)

Push data

Restricted to API keys with at least one of the following roles : IS_PARTNER.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Bearer {token} | 
**data** | [**PartnerDataItemSwagger**](PartnerDataItemSwagger.md)| data | 

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

