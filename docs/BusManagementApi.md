# \BusManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST3**](BusManagementApi.md#CreateUsingPOST3) | **Post** /api/v0/topics/fifo | Create a FIFO
[**DeleteUsingDELETE20**](BusManagementApi.md#DeleteUsingDELETE20) | **Delete** /api/v0/topics/fifo/{fifoName} | Delete a FIFO
[**FifoPublishUsingPOST**](BusManagementApi.md#FifoPublishUsingPOST) | **Post** /api/v0/topics/fifo/{fifoName} | Publish a message into a FIFO
[**GetUsingGET19**](BusManagementApi.md#GetUsingGET19) | **Get** /api/v0/topics/fifo/{fifoName} | Get a FIFO
[**ListFifoTopicsUsingGET**](BusManagementApi.md#ListFifoTopicsUsingGET) | **Get** /api/v0/topics/fifo | List all FIFOs



## CreateUsingPOST3

> CreateUsingPOST3(ctx).XAPIKEY(xAPIKEY).RequestBody(requestBody).Execute()

Create a FIFO



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
    requestBody := *openapiclient.NewFifoCreateReqWeb(int64(123), "Name_example") // FifoCreateReqWeb | body for create a new fifo (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BusManagementApi.CreateUsingPOST3(context.Background()).XAPIKEY(xAPIKEY).RequestBody(requestBody).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusManagementApi.CreateUsingPOST3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsingPOST3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **requestBody** | [**FifoCreateReqWeb**](FifoCreateReqWeb.md) | body for create a new fifo | 

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


## DeleteUsingDELETE20

> DeleteUsingDELETE20(ctx, fifoName).XAPIKEY(xAPIKEY).Force(force).Execute()

Delete a FIFO



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
    fifoName := "fifoName_example" // string | name of the fifo. Fifo name must respect the following regular expression <code>[A-Za-z\\d-_~]+</code> (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | false: Live Objects prevents deleting fifo while there is at least one consumer, true: the fifo will be deleted in any cases. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BusManagementApi.DeleteUsingDELETE20(context.Background(), fifoName).XAPIKEY(xAPIKEY).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusManagementApi.DeleteUsingDELETE20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fifoName** | **string** | name of the fifo. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | false: Live Objects prevents deleting fifo while there is at least one consumer, true: the fifo will be deleted in any cases. | [default to false]

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


## FifoPublishUsingPOST

> bool FifoPublishUsingPOST(ctx, fifoName).XAPIKEY(xAPIKEY).Payload(payload).Execute()

Publish a message into a FIFO



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
    fifoName := "fifoName_example" // string | name of the fifo. Fifo name must respect the following regular expression <code>[A-Za-z\\d-_~]+</code> (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    payload := 987 // string | body for publish a message into your FIFO topic (your body is specific at your topic) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BusManagementApi.FifoPublishUsingPOST(context.Background(), fifoName).XAPIKEY(xAPIKEY).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusManagementApi.FifoPublishUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FifoPublishUsingPOST`: bool
    fmt.Fprintf(os.Stdout, "Response from `BusManagementApi.FifoPublishUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fifoName** | **string** | name of the fifo. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiFifoPublishUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **payload** | **string** | body for publish a message into your FIFO topic (your body is specific at your topic) | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json, 
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET19

> FifoTopic GetUsingGET19(ctx, fifoName).XAPIKEY(xAPIKEY).Execute()

Get a FIFO



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
    fifoName := "fifoName_example" // string | name of the fifo. Fifo name must respect the following regular expression <code>[A-Za-z\\d-_~]+</code> (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.BusManagementApi.GetUsingGET19(context.Background(), fifoName).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusManagementApi.GetUsingGET19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET19`: FifoTopic
    fmt.Fprintf(os.Stdout, "Response from `BusManagementApi.GetUsingGET19`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**fifoName** | **string** | name of the fifo. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET19Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**FifoTopic**](FifoTopic.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListFifoTopicsUsingGET

> PageableFifoTopic ListFifoTopicsUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()

List all FIFOs



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
    resp, r, err := api_client.BusManagementApi.ListFifoTopicsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BusManagementApi.ListFifoTopicsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListFifoTopicsUsingGET`: PageableFifoTopic
    fmt.Fprintf(os.Stdout, "Response from `BusManagementApi.ListFifoTopicsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListFifoTopicsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**PageableFifoTopic**](Pageable«FifoTopic».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

