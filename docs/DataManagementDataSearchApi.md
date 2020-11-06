# \DataManagementDataSearchApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DslQueryHitsOnlyUsingPOST**](DataManagementDataSearchApi.md#DslQueryHitsOnlyUsingPOST) | **Post** /api/v0/data/search/hits | Query an Elasticsearch Domain Specific Language request and get only hits result
[**DslQueryUsingPOST**](DataManagementDataSearchApi.md#DslQueryUsingPOST) | **Post** /api/v0/data/search | Query an Elasticsearch Domain Specific Language request



## DslQueryHitsOnlyUsingPOST

> []DataStoredWeb DslQueryHitsOnlyUsingPOST(ctx).XAPIKEY(xAPIKEY).DslRequest(dslRequest).Execute()

Query an Elasticsearch Domain Specific Language request and get only hits result



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
    dslRequest := 987 // map[string]interface{} | elasticsearch DSL request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementDataSearchApi.DslQueryHitsOnlyUsingPOST(context.Background()).XAPIKEY(xAPIKEY).DslRequest(dslRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementDataSearchApi.DslQueryHitsOnlyUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DslQueryHitsOnlyUsingPOST`: []DataStoredWeb
    fmt.Fprintf(os.Stdout, "Response from `DataManagementDataSearchApi.DslQueryHitsOnlyUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDslQueryHitsOnlyUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **dslRequest** | **map[string]interface{}** | elasticsearch DSL request | 

### Return type

[**[]DataStoredWeb**](DataStoredWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DslQueryUsingPOST

> map[string]interface{} DslQueryUsingPOST(ctx).XAPIKEY(xAPIKEY).DslRequest(dslRequest).Execute()

Query an Elasticsearch Domain Specific Language request



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
    dslRequest := 987 // map[string]interface{} | elasticsearch DSL request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementDataSearchApi.DslQueryUsingPOST(context.Background()).XAPIKEY(xAPIKEY).DslRequest(dslRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementDataSearchApi.DslQueryUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DslQueryUsingPOST`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `DataManagementDataSearchApi.DslQueryUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDslQueryUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **dslRequest** | **map[string]interface{}** | elasticsearch DSL request | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

