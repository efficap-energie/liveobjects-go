# \DataBulkInjectionApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataBulkUsingPOST**](DataBulkInjectionApi.md#AddDataBulkUsingPOST) | **Post** /api/v0/data/bulk | Insert a bulk of new Data



## AddDataBulkUsingPOST

> string AddDataBulkUsingPOST(ctx).XAPIKEY(xAPIKEY).DataBulk(dataBulk).Execute()

Insert a bulk of new Data



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
    dataBulk := []DataBulkItemWeb{*openapiclient.NewDataBulkItemWeb(123)} // []DataBulkItemWeb | Bulk of Data to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataBulkInjectionApi.AddDataBulkUsingPOST(context.Background()).XAPIKEY(xAPIKEY).DataBulk(dataBulk).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataBulkInjectionApi.AddDataBulkUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDataBulkUsingPOST`: string
    fmt.Fprintf(os.Stdout, "Response from `DataBulkInjectionApi.AddDataBulkUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddDataBulkUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **dataBulk** | [**[]DataBulkItemWeb**](DataBulkItemWeb.md) | Bulk of Data to add | 

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

