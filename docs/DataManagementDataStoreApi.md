# \DataManagementDataStoreApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataMessageUsingPOST**](DataManagementDataStoreApi.md#AddDataMessageUsingPOST) | **Post** /api/v0/data/streams/{streamId} | Insert a new Data into the stream
[**RetrieveDataUsingGET**](DataManagementDataStoreApi.md#RetrieveDataUsingGET) | **Get** /api/v0/data/streams/{streamId} | Retrieve data from the streamId



## AddDataMessageUsingPOST

> string AddDataMessageUsingPOST(ctx, streamId).XAPIKEY(xAPIKEY).Data(data).Execute()

Insert a new Data into the stream



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
    streamId := "streamId_example" // string | StreamId in which the data will be added. Should not contains following character or space : ' \\ \" ; { } ( )
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    data := *openapiclient.NewDataWeb(123) // DataWeb | DataMessage to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementDataStoreApi.AddDataMessageUsingPOST(context.Background(), streamId).XAPIKEY(xAPIKEY).Data(data).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementDataStoreApi.AddDataMessageUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddDataMessageUsingPOST`: string
    fmt.Fprintf(os.Stdout, "Response from `DataManagementDataStoreApi.AddDataMessageUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | StreamId in which the data will be added. Should not contains following character or space : &#39; \\ \&quot; ; { } ( ) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddDataMessageUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **data** | [**DataWeb**](DataWeb.md) | DataMessage to add | 

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


## RetrieveDataUsingGET

> []DataStoredWeb RetrieveDataUsingGET(ctx, streamId).XAPIKEY(xAPIKEY).Limit(limit).TimeRange(timeRange).BookmarkId(bookmarkId).Execute()

Retrieve data from the streamId



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
    streamId := "streamId_example" // string | StreamId from which the data will be retrieved
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    limit := 987 // int32 | max number of data to return, value is limited to 1000 (optional) (default to 100)
    timeRange := []string{"Inner_example"} // []string | filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange=[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusiv, upperbound is exclusiv. (optional)
    bookmarkId := "bookmarkId_example" // string | id of the last document retrieved that can be used to paginate : first result will be the one following this document id (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DataManagementDataStoreApi.RetrieveDataUsingGET(context.Background(), streamId).XAPIKEY(xAPIKEY).Limit(limit).TimeRange(timeRange).BookmarkId(bookmarkId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DataManagementDataStoreApi.RetrieveDataUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveDataUsingGET`: []DataStoredWeb
    fmt.Fprintf(os.Stdout, "Response from `DataManagementDataStoreApi.RetrieveDataUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string** | StreamId from which the data will be retrieved | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveDataUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **limit** | **int32** | max number of data to return, value is limited to 1000 | [default to 100]
 **timeRange** | [**[]string**](string.md) | filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange&#x3D;[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusiv, upperbound is exclusiv. | 
 **bookmarkId** | **string** | id of the last document retrieved that can be used to paginate : first result will be the one following this document id | 

### Return type

[**[]DataStoredWeb**](DataStoredWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

