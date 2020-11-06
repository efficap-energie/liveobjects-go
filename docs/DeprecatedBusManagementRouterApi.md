# \DeprecatedBusManagementRouterApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddBindingUsingPOST**](DeprecatedBusManagementRouterApi.md#AddBindingUsingPOST) | **Post** /api/v0/bindings | Add a binding for a FIFO queue
[**DeleteBindingUsingDELETE**](DeprecatedBusManagementRouterApi.md#DeleteBindingUsingDELETE) | **Delete** /api/v0/bindings/{routingKeyFilter}/{fifoName} | Delete a binding
[**ListBindingsUsingGET**](DeprecatedBusManagementRouterApi.md#ListBindingsUsingGET) | **Get** /api/v0/bindings | List registered bindings



## AddBindingUsingPOST

> AddBindingUsingPOST(ctx).XAPIKEY(xAPIKEY).Req(req).Execute()

Add a binding for a FIFO queue



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
    req := *openapiclient.NewCreateFifoBindingRequest("FifoId_example", "RoutingKeyFilter_example") // CreateFifoBindingRequest | body of create FIFO bindings (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedBusManagementRouterApi.AddBindingUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedBusManagementRouterApi.AddBindingUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddBindingUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **req** | [**CreateFifoBindingRequest**](CreateFifoBindingRequest.md) | body of create FIFO bindings | 

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


## DeleteBindingUsingDELETE

> DeleteBindingUsingDELETE(ctx, routingKeyFilter, fifoName).XAPIKEY(xAPIKEY).Execute()

Delete a binding



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
    routingKeyFilter := "routingKeyFilter_example" // string | routing key filter. Routing key must respect the following regular expression <code>[A-Za-z\\d-_:~#\\.\\*]+</code> (max 1000 characters)
    fifoName := "fifoName_example" // string | name of the fifo. Fifo name must respect the following regular expression <code>[A-Za-z\\d-_~]+</code> (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedBusManagementRouterApi.DeleteBindingUsingDELETE(context.Background(), routingKeyFilter, fifoName).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedBusManagementRouterApi.DeleteBindingUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingKeyFilter** | **string** | routing key filter. Routing key must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_:~#\\.\\*]+&lt;/code&gt; (max 1000 characters) | 
**fifoName** | **string** | name of the fifo. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBindingUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListBindingsUsingGET

> PageableFifoBinding ListBindingsUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()

List registered bindings



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
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedBusManagementRouterApi.ListBindingsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedBusManagementRouterApi.ListBindingsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListBindingsUsingGET`: PageableFifoBinding
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedBusManagementRouterApi.ListBindingsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListBindingsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**PageableFifoBinding**](Pageable«FifoBinding».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

