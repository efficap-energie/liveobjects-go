# \PartnerDataPushApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DataPushUsingPOST**](PartnerDataPushApi.md#DataPushUsingPOST) | **Post** /api/v0/partners/data/streams | Push data



## DataPushUsingPOST

> string DataPushUsingPOST(ctx).Authorization(authorization).Data(data).Execute()

Push data



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
    data := *openapiclient.NewPartnerDataItemSwagger(*openapiclient.NewMetadata(), 123) // PartnerDataItemSwagger | data

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartnerDataPushApi.DataPushUsingPOST(context.Background()).Authorization(authorization).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnerDataPushApi.DataPushUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DataPushUsingPOST`: string
    fmt.Fprintf(os.Stdout, "Response from `PartnerDataPushApi.DataPushUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDataPushUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Bearer {token} | 
 **data** | [**PartnerDataItemSwagger**](PartnerDataItemSwagger.md) | data | 

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

