# \AccountingV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyStatisticsUsingGET2**](AccountingV1Api.md#GetDailyStatisticsUsingGET2) | **Get** /api/v1/accounting/daily | Get daily accounting metrics (Beta).
[**GetMonthlyStatisticsUsingGET2**](AccountingV1Api.md#GetMonthlyStatisticsUsingGET2) | **Get** /api/v1/accounting/monthly | Get monthly accounting metrics.



## GetDailyStatisticsUsingGET2

> TenantDayMetrics GetDailyStatisticsUsingGET2(ctx).XAPIKEY(xAPIKEY).StartDay(startDay).EndDay(endDay).Execute()

Get daily accounting metrics (Beta).



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
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    startDay := "startDay_example" // string | the requested start day as yyyy-MM-dd. If missing, the current day is used. (optional)
    endDay := "endDay_example" // string | the requested end day as yyyy-MM-dd. If missing, the current day is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountingV1Api.GetDailyStatisticsUsingGET2(context.Background()).XAPIKEY(xAPIKEY).StartDay(startDay).EndDay(endDay).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountingV1Api.GetDailyStatisticsUsingGET2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDailyStatisticsUsingGET2`: TenantDayMetrics
    fmt.Fprintf(os.Stdout, "Response from `AccountingV1Api.GetDailyStatisticsUsingGET2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDailyStatisticsUsingGET2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **startDay** | **string** | the requested start day as yyyy-MM-dd. If missing, the current day is used. | 
 **endDay** | **string** | the requested end day as yyyy-MM-dd. If missing, the current day is used. | 

### Return type

[**TenantDayMetrics**](TenantDayMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonthlyStatisticsUsingGET2

> TenantMonthMetrics GetMonthlyStatisticsUsingGET2(ctx).XAPIKEY(xAPIKEY).StartMonth(startMonth).EndMonth(endMonth).Execute()

Get monthly accounting metrics.



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
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    startMonth := "startMonth_example" // string | the requested start month as yyyy-MM. If missing, the current month is used. (optional)
    endMonth := "endMonth_example" // string | the requested end month as yyyy-MM. If missing, the current month is used. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AccountingV1Api.GetMonthlyStatisticsUsingGET2(context.Background()).XAPIKEY(xAPIKEY).StartMonth(startMonth).EndMonth(endMonth).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AccountingV1Api.GetMonthlyStatisticsUsingGET2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMonthlyStatisticsUsingGET2`: TenantMonthMetrics
    fmt.Fprintf(os.Stdout, "Response from `AccountingV1Api.GetMonthlyStatisticsUsingGET2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMonthlyStatisticsUsingGET2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **startMonth** | **string** | the requested start month as yyyy-MM. If missing, the current month is used. | 
 **endMonth** | **string** | the requested end month as yyyy-MM. If missing, the current month is used. | 

### Return type

[**TenantMonthMetrics**](TenantMonthMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

