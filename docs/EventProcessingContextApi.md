# \EventProcessingContextApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ClearUsingDELETE**](EventProcessingContextApi.md#ClearUsingDELETE) | **Delete** /api/v0/eventprocessing/context | Delete all context entries
[**DeleteUsingDELETE15**](EventProcessingContextApi.md#DeleteUsingDELETE15) | **Delete** /api/v0/eventprocessing/context/{contextKey} | Delete a context key
[**GetUsingGET14**](EventProcessingContextApi.md#GetUsingGET14) | **Get** /api/v0/eventprocessing/context/{contextKey} | Retrieve a context
[**ListUsingGET17**](EventProcessingContextApi.md#ListUsingGET17) | **Get** /api/v0/eventprocessing/context | Retrieve the list of contexts with optional tag filtering
[**SaveUsingPUT**](EventProcessingContextApi.md#SaveUsingPUT) | **Put** /api/v0/eventprocessing/context/{contextKey} | Save a context



## ClearUsingDELETE

> int64 ClearUsingDELETE(ctx).XAPIKEY(xAPIKEY).Execute()

Delete all context entries



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingContextApi.ClearUsingDELETE(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingContextApi.ClearUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ClearUsingDELETE`: int64
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingContextApi.ClearUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiClearUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE15

> DeleteUsingDELETE15(ctx, contextKey).XAPIKEY(xAPIKEY).Execute()

Delete a context key



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
    contextKey := "contextKey_example" // string | id of the context to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingContextApi.DeleteUsingDELETE15(context.Background(), contextKey).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingContextApi.DeleteUsingDELETE15``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextKey** | **string** | id of the context to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET14

> ContextContainer GetUsingGET14(ctx, contextKey).XAPIKEY(xAPIKEY).Execute()

Retrieve a context



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
    contextKey := "contextKey_example" // string | id of the context to get
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingContextApi.GetUsingGET14(context.Background(), contextKey).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingContextApi.GetUsingGET14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET14`: ContextContainer
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingContextApi.GetUsingGET14`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextKey** | **string** | id of the context to get | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**ContextContainer**](ContextContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET17

> []ContextContainer ListUsingGET17(ctx).XAPIKEY(xAPIKEY).Page(page).Size(size).Tags(tags).Execute()

Retrieve the list of contexts with optional tag filtering



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
    page := 987 // int32 | The requested page number (optional) (default to 0)
    size := 987 // int32 | The maximum number of items per page (optional) (default to 20)
    tags := []string{"Inner_example"} // []string | Filtering tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingContextApi.ListUsingGET17(context.Background()).XAPIKEY(xAPIKEY).Page(page).Size(size).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingContextApi.ListUsingGET17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET17`: []ContextContainer
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingContextApi.ListUsingGET17`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **page** | **int32** | The requested page number | [default to 0]
 **size** | **int32** | The maximum number of items per page | [default to 20]
 **tags** | [**[]string**](string.md) | Filtering tags | 

### Return type

[**[]ContextContainer**](ContextContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveUsingPUT

> ContextContainer SaveUsingPUT(ctx, contextKey).XAPIKEY(xAPIKEY).ContextContainer(contextContainer).Execute()

Save a context



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
    contextKey := "contextKey_example" // string | id of the context to save
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    contextContainer := *openapiclient.NewContextContainer(123) // ContextContainer | JSON context object to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingContextApi.SaveUsingPUT(context.Background(), contextKey).XAPIKEY(xAPIKEY).ContextContainer(contextContainer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingContextApi.SaveUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SaveUsingPUT`: ContextContainer
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingContextApi.SaveUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**contextKey** | **string** | id of the context to save | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **contextContainer** | [**ContextContainer**](ContextContainer.md) | JSON context object to add | 

### Return type

[**ContextContainer**](ContextContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

