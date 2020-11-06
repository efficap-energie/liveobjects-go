# \DeprecatedRouterTopicManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PublishUsingPOST**](DeprecatedRouterTopicManagementApi.md#PublishUsingPOST) | **Post** /api/v0/topics/router/{routingKey} | Publishing a message into a ROUTER topic



## PublishUsingPOST

> bool PublishUsingPOST(ctx, routingKey).XAPIKEY(xAPIKEY).Payload(payload).Execute()

Publishing a message into a ROUTER topic



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
    routingKey := "routingKey_example" // string | Routing Key
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    payload := 987 // string | payload (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedRouterTopicManagementApi.PublishUsingPOST(context.Background(), routingKey).XAPIKEY(xAPIKEY).Payload(payload).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedRouterTopicManagementApi.PublishUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PublishUsingPOST`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedRouterTopicManagementApi.PublishUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**routingKey** | **string** | Routing Key | 

### Other Parameters

Other parameters are passed through a pointer to a apiPublishUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **payload** | **string** | payload | 

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

