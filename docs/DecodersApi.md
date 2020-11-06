# \DecodersApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListUsingGET22**](DecodersApi.md#ListUsingGET22) | **Get** /api/v0/decoders | Retrieve the list of all decoders (binary, csv, js)



## ListUsingGET22

> []PayloadDescription ListUsingGET22(ctx).XAPIKEY(xAPIKEY).Tags(tags).Execute()

Retrieve the list of all decoders (binary, csv, js)



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
    resp, r, err := api_client.DecodersApi.ListUsingGET22(context.Background()).XAPIKEY(xAPIKEY).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersApi.ListUsingGET22``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET22`: []PayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersApi.ListUsingGET22`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET22Request struct via the builder pattern


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

