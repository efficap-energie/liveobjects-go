# \DataManagementCustomPipelinesApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE11**](DataManagementCustomPipelinesApi.md#DeleteUsingDELETE11) | **Delete** /api/v0/pipelines/{pipelineId} | Delete a DataMessage pipeline
[**GetUsingGET12**](DataManagementCustomPipelinesApi.md#GetUsingGET12) | **Get** /api/v0/pipelines/{pipelineId} | Retrieve a DataMessage pipeline
[**ListUsingGET13**](DataManagementCustomPipelinesApi.md#ListUsingGET13) | **Get** /api/v0/pipelines | Retrieve the list of DataMessage pipelines, ordered by priorityLevel
[**PostUsingPOST12**](DataManagementCustomPipelinesApi.md#PostUsingPOST12) | **Post** /api/v0/pipelines | Create a DataMessage pipeline
[**UpdateUsingPUT6**](DataManagementCustomPipelinesApi.md#UpdateUsingPUT6) | **Put** /api/v0/pipelines/{pipelineId} | Update a DataMessage pipeline



## DeleteUsingDELETE11

> DeleteUsingDELETE11(ctx, pipelineId).XAPIKEY(xAPIKEY).Execute()

Delete a DataMessage pipeline



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
    pipelineId := "pipelineId_example" // string | id of the pipeline to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementCustomPipelinesApi.DeleteUsingDELETE11(context.Background(), pipelineId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementCustomPipelinesApi.DeleteUsingDELETE11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | id of the pipeline to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE11Request struct via the builder pattern


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


## GetUsingGET12

> Pipeline GetUsingGET12(ctx, pipelineId).XAPIKEY(xAPIKEY).Execute()

Retrieve a DataMessage pipeline



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
    pipelineId := "pipelineId_example" // string | id of the pipeline to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementCustomPipelinesApi.GetUsingGET12(context.Background(), pipelineId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementCustomPipelinesApi.GetUsingGET12``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET12`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `DataManagementCustomPipelinesApi.GetUsingGET12`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | id of the pipeline to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET13

> []Pipeline ListUsingGET13(ctx).XAPIKEY(xAPIKEY).Execute()

Retrieve the list of DataMessage pipelines, ordered by priorityLevel



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
    resp, r, err := api_client.DataManagementCustomPipelinesApi.ListUsingGET13(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementCustomPipelinesApi.ListUsingGET13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET13`: []Pipeline
    fmt.Fprintf(os.Stdout, "Response from `DataManagementCustomPipelinesApi.ListUsingGET13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**[]Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST12

> Pipeline PostUsingPOST12(ctx).XAPIKEY(xAPIKEY).PipelineDescription(pipelineDescription).Execute()

Create a DataMessage pipeline



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
    pipelineDescription := *openapiclient.NewPipeline(false, "Name_example", int64(123), []PipelineStep{*openapiclient.NewPipelineStep())) // Pipeline | pipeline to create (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementCustomPipelinesApi.PostUsingPOST12(context.Background()).XAPIKEY(xAPIKEY).PipelineDescription(pipelineDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementCustomPipelinesApi.PostUsingPOST12``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsingPOST12`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `DataManagementCustomPipelinesApi.PostUsingPOST12`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsingPOST12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **pipelineDescription** | [**Pipeline**](Pipeline.md) | pipeline to create | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT6

> Pipeline UpdateUsingPUT6(ctx, pipelineId).XAPIKEY(xAPIKEY).PipelineDescription(pipelineDescription).Execute()

Update a DataMessage pipeline



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
    pipelineId := "pipelineId_example" // string | id of the pipeline to update
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    pipelineDescription := *openapiclient.NewPipeline(false, "Name_example", int64(123), []PipelineStep{*openapiclient.NewPipelineStep())) // Pipeline | pipeline to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementCustomPipelinesApi.UpdateUsingPUT6(context.Background(), pipelineId).XAPIKEY(xAPIKEY).PipelineDescription(pipelineDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementCustomPipelinesApi.UpdateUsingPUT6``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUsingPUT6`: Pipeline
    fmt.Fprintf(os.Stdout, "Response from `DataManagementCustomPipelinesApi.UpdateUsingPUT6`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipelineId** | **string** | id of the pipeline to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsingPUT6Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **pipelineDescription** | [**Pipeline**](Pipeline.md) | pipeline to update | 

### Return type

[**Pipeline**](Pipeline.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

