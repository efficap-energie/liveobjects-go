# \PartnersManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOptionToTenantUsingPOST**](PartnersManagementApi.md#AddOptionToTenantUsingPOST) | **Post** /api/v0/partners/tenants/{tenantId}/options | Subscribe an option for a partner&#39;s tenant
[**CreateTenantByPartnerUsingPOST**](PartnersManagementApi.md#CreateTenantByPartnerUsingPOST) | **Post** /api/v0/partners/tenants | Create a partner&#39;s tenant



## AddOptionToTenantUsingPOST

> AddOptionToTenantUsingPOST(ctx, tenantId, authorization, req)

Subscribe an option for a partner's tenant

Usage of this API will be reported in your access log under 'admin_account' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string**| Tenant identifier | 
**authorization** | **string**| Bearer {token} | 
**req** | [**OptionPartnerTenant**](OptionPartnerTenant.md)| req | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateTenantByPartnerUsingPOST

> PartnerTenant CreateTenantByPartnerUsingPOST(ctx, authorization, req)

Create a partner's tenant

Usage of this API will be reported in your access log under 'admin_account' category.<br><br>Restricted to API keys with at least one of the following roles : IS_PARTNER.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Bearer {token} | 
**req** | [**AddPartnerTenantRequest**](AddPartnerTenantRequest.md)| req | 

### Return type

[**PartnerTenant**](PartnerTenant.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

