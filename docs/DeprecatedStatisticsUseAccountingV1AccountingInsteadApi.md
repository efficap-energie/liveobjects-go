# \DeprecatedStatisticsUseAccountingV1AccountingInsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantStatisticsUsingGET3**](DeprecatedStatisticsUseAccountingV1AccountingInsteadApi.md#GetTenantStatisticsUsingGET3) | **Get** /api/v0/statistics/tenant/{tenantId} | Get tenant statistics for a specific tenant and a range of dates



## GetTenantStatisticsUsingGET3

> TenantStatsWeb GetTenantStatisticsUsingGET3(ctx, tenantId).StartDate(startDate).EndDate(endDate).XAPIKEY(xAPIKEY).Execute()

Get tenant statistics for a specific tenant and a range of dates



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
    tenantId := "tenantId_example" // string | the id of your tenant ex: 57xxxxxxxxxxxxxxxxxxxxxx
    startDate := "startDate_example" // string | the requested start date as yyyy-MM-dd
    endDate := "endDate_example" // string | the requested end date as yyyy-MM-dd
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedStatisticsUseAccountingV1AccountingInsteadApi.GetTenantStatisticsUsingGET3(context.Background(), tenantId).StartDate(startDate).EndDate(endDate).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedStatisticsUseAccountingV1AccountingInsteadApi.GetTenantStatisticsUsingGET3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantStatisticsUsingGET3`: TenantStatsWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedStatisticsUseAccountingV1AccountingInsteadApi.GetTenantStatisticsUsingGET3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string** | the id of your tenant ex: 57xxxxxxxxxxxxxxxxxxxxxx | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantStatisticsUsingGET3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDate** | **string** | the requested start date as yyyy-MM-dd | 
 **endDate** | **string** | the requested end date as yyyy-MM-dd | 
 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**TenantStatsWeb**](TenantStatsWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

