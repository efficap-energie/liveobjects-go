# \DecodersPublicApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsingGET23**](DecodersPublicApi.md#ListUsingGET23) | **Get** /api/v0/decoders/public | Retrieve the list of all public decoders (binary, csv, js)



## ListUsingGET23

> []PayloadDescription ListUsingGET23(ctx).XAPIKEY(xAPIKEY).Tags(tags).Execute()

Retrieve the list of all public decoders (binary, csv, js)



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
    tags := []string{"Inner_example"} // []string | target filtering tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersPublicApi.ListUsingGET23(context.Background()).XAPIKEY(xAPIKEY).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersPublicApi.ListUsingGET23``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET23`: []PayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersPublicApi.ListUsingGET23`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET23Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **tags** | [**[]string**](string.md) | target filtering tags | 

### Return type

[**[]PayloadDescription**](PayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

