# \PartnersManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddOptionToTenantUsingPOST**](PartnersManagementApi.md#AddOptionToTenantUsingPOST) | **Post** /api/v0/partners/tenants/{tenantId}/options | Subscribe an option for a partner&#39;s tenant
[**CreateTenantByPartnerUsingPOST**](PartnersManagementApi.md#CreateTenantByPartnerUsingPOST) | **Post** /api/v0/partners/tenants | Create a partner&#39;s tenant



## AddOptionToTenantUsingPOST

> AddOptionToTenantUsingPOST(ctx, tenantId).Authorization(authorization).Req(req).Execute()

Subscribe an option for a partner's tenant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    tenantId := "tenantId_example" // string | Tenant identifier
    authorization := "authorization_example" // string | Bearer {token}
    req := *openapiclient.NewOptionPartnerTenant("Id_example") // OptionPartnerTenant | req

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartnersManagementApi.AddOptionToTenantUsingPOST(context.Background(), tenantId).Authorization(authorization).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersManagementApi.AddOptionToTenantUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | Tenant identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddOptionToTenantUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **authorization** | **string** | Bearer {token} | 
 **req** | [**OptionPartnerTenant**](OptionPartnerTenant.md) | req | 

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

> PartnerTenant CreateTenantByPartnerUsingPOST(ctx).Authorization(authorization).Req(req).Execute()

Create a partner's tenant



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    authorization := "authorization_example" // string | Bearer {token}
    req := *openapiclient.NewAddPartnerTenantRequest("Email_example", "Language_example", "Login_example", "Name_example", "OfferMappingId_example") // AddPartnerTenantRequest | req

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartnersManagementApi.CreateTenantByPartnerUsingPOST(context.Background()).Authorization(authorization).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersManagementApi.CreateTenantByPartnerUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateTenantByPartnerUsingPOST`: PartnerTenant
    fmt.Fprintf(os.Stdout, "Response from `PartnersManagementApi.CreateTenantByPartnerUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantByPartnerUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer {token} | 
 **req** | [**AddPartnerTenantRequest**](AddPartnerTenantRequest.md) | req | 

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

