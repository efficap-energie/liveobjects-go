# \EventProcessingGeozoneApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE17**](EventProcessingGeozoneApi.md#DeleteUsingDELETE17) | **Delete** /api/v0/eventprocessing/geozones/{zoneId} | delete a geographic zone
[**GetUsingGET16**](EventProcessingGeozoneApi.md#GetUsingGET16) | **Get** /api/v0/eventprocessing/geozones/{zoneId} | retrieve a geozone from repository
[**ListUsingGET19**](EventProcessingGeozoneApi.md#ListUsingGET19) | **Get** /api/v0/eventprocessing/geozones | retrieve paginated list of geozones
[**SaveUsingPUT1**](EventProcessingGeozoneApi.md#SaveUsingPUT1) | **Put** /api/v0/eventprocessing/geozones/{zoneId} | Save a geographic zone



## DeleteUsingDELETE17

> DeleteUsingDELETE17(ctx, zoneId).XAPIKEY(xAPIKEY).Execute()

delete a geographic zone



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
    zoneId := "zoneId_example" // string | the id of the zone to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingGeozoneApi.DeleteUsingDELETE17(context.Background(), zoneId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingGeozoneApi.DeleteUsingDELETE17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | the id of the zone to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE17Request struct via the builder pattern


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


## GetUsingGET16

> GeozoneContainer GetUsingGET16(ctx, zoneId).XAPIKEY(xAPIKEY).Execute()

retrieve a geozone from repository



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
    zoneId := "zoneId_example" // string | the user-defined id of the zone to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingGeozoneApi.GetUsingGET16(context.Background(), zoneId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingGeozoneApi.GetUsingGET16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET16`: GeozoneContainer
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingGeozoneApi.GetUsingGET16`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | the user-defined id of the zone to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**GeozoneContainer**](GeozoneContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET19

> []GeozoneContainer ListUsingGET19(ctx).XAPIKEY(xAPIKEY).Page(page).Size(size).Execute()

retrieve paginated list of geozones



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
    page := 987 // int32 | The requested page number. Starts from 0. (optional) (default to 0)
    size := 987 // int32 | The maximum number of items per page. Must be between 1 and 1000. (optional) (default to 20)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingGeozoneApi.ListUsingGET19(context.Background()).XAPIKEY(xAPIKEY).Page(page).Size(size).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingGeozoneApi.ListUsingGET19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET19`: []GeozoneContainer
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingGeozoneApi.ListUsingGET19`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET19Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **page** | **int32** | The requested page number. Starts from 0. | [default to 0]
 **size** | **int32** | The maximum number of items per page. Must be between 1 and 1000. | [default to 20]

### Return type

[**[]GeozoneContainer**](GeozoneContainer.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SaveUsingPUT1

> SaveUsingPUT1(ctx, zoneId).XAPIKEY(xAPIKEY).GeozoneContainerBase(geozoneContainerBase).Execute()

Save a geographic zone



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
    zoneId := "zoneId_example" // string | the id of the zone to save
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    geozoneContainerBase := *openapiclient.NewGeozoneContainerBase() // GeozoneContainerBase | JSON geo zone object to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingGeozoneApi.SaveUsingPUT1(context.Background(), zoneId).XAPIKEY(xAPIKEY).GeozoneContainerBase(geozoneContainerBase).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingGeozoneApi.SaveUsingPUT1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**zoneId** | **string** | the id of the zone to save | 

### Other Parameters

Other parameters are passed through a pointer to a apiSaveUsingPUT1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **geozoneContainerBase** | [**GeozoneContainerBase**](GeozoneContainerBase.md) | JSON geo zone object to add | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

