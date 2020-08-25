# \TenantAccountsApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMyTenantUsingGET**](TenantAccountsApi.md#GetMyTenantUsingGET) | **Get** /api/v0/tenants/me | Get details of your account



## GetMyTenantUsingGET

> Tenant GetMyTenantUsingGET(ctx, xAPIKEY)

Get details of your account

<p><strong>Beware</strong>: The information contained in the following 3 fields is restricted to Live Objects Portal application only. It may change without notice, no backward compatibility is guaranteed to customer applications<ul><li>\"offerSettings\"<li>\"technicalSettings\"<li>\"userSettings\"<ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**Tenant**](Tenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

