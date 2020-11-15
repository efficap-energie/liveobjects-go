# \DeprecatedTriggersAndActionsTestApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TestHttpPushUsingPOST**](DeprecatedTriggersAndActionsTestApi.md#TestHttpPushUsingPOST) | **Post** /api/v0/event2action/test/http-push | Post an http request for testing a webhook



## TestHttpPushUsingPOST

> HttpPushTestResult TestHttpPushUsingPOST(ctx).XAPIKEY(xAPIKEY).Req(req).Execute()

Post an http request for testing a webhook



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
    req := *openapiclient.NewHttpPushWebhookTest() // HttpPushWebhookTest | Http push request to be tested (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedTriggersAndActionsTestApi.TestHttpPushUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedTriggersAndActionsTestApi.TestHttpPushUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestHttpPushUsingPOST`: HttpPushTestResult
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedTriggersAndActionsTestApi.TestHttpPushUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestHttpPushUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **req** | [**HttpPushWebhookTest**](HttpPushWebhookTest.md) | Http push request to be tested | 

### Return type

[**HttpPushTestResult**](HttpPushTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

