# \NotificationApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendMessageUsingPOST**](NotificationApi.md#SendMessageUsingPOST) | **Post** /api/v0/contact | API to send message



## SendMessageUsingPOST

> SendMessageUsingPOST(ctx).XAPIKEY(xAPIKEY).Request(request).Execute()

API to send message



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
    request := *openapiclient.NewSendMessageRequest(*openapiclient.NewContactMessage(), map[string]string{ "key": "Inner_example"}) // SendMessageRequest | contact and message are mandatory (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.NotificationApi.SendMessageUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `NotificationApi.SendMessageUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendMessageUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**SendMessageRequest**](SendMessageRequest.md) | contact and message are mandatory | 

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

