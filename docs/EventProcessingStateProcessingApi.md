# \EventProcessingStateProcessingApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE19**](EventProcessingStateProcessingApi.md#DeleteUsingDELETE19) | **Delete** /api/v0/eventprocessing/stateprocessing-rule/{stateProcessingRuleId} | Delete a StateProcessingRule
[**GetUsingGET18**](EventProcessingStateProcessingApi.md#GetUsingGET18) | **Get** /api/v0/eventprocessing/stateprocessing-rule/{stateProcessingRuleId} | Retrieve a StateProcessingRule
[**ListUsingGET21**](EventProcessingStateProcessingApi.md#ListUsingGET21) | **Get** /api/v0/eventprocessing/stateprocessing-rule | Retrieve the list of all the StateProcessingRules or get a StateProcessingRule by its name
[**PostUsingPOST16**](EventProcessingStateProcessingApi.md#PostUsingPOST16) | **Post** /api/v0/eventprocessing/stateprocessing-rule | Create a StateProcessingRule
[**TestUsingPOST5**](EventProcessingStateProcessingApi.md#TestUsingPOST5) | **Post** /api/v0/eventprocessing/stateprocessing-rule/test | test a  State Processing function
[**UpdateUsingPUT10**](EventProcessingStateProcessingApi.md#UpdateUsingPUT10) | **Put** /api/v0/eventprocessing/stateprocessing-rule/{stateProcessingRuleId} | Update a StateProcessingRule



## DeleteUsingDELETE19

> DeleteUsingDELETE19(ctx, stateProcessingRuleId).XAPIKEY(xAPIKEY).Execute()

Delete a StateProcessingRule



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
    stateProcessingRuleId := "stateProcessingRuleId_example" // string | id of the StateProcessingRule to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingStateProcessingApi.DeleteUsingDELETE19(context.Background(), stateProcessingRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingStateProcessingApi.DeleteUsingDELETE19``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateProcessingRuleId** | **string** | id of the StateProcessingRule to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE19Request struct via the builder pattern


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


## GetUsingGET18

> StateProcessingRule GetUsingGET18(ctx, stateProcessingRuleId).XAPIKEY(xAPIKEY).Execute()

Retrieve a StateProcessingRule



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
    stateProcessingRuleId := "stateProcessingRuleId_example" // string | id of the StateProcessingRule to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingStateProcessingApi.GetUsingGET18(context.Background(), stateProcessingRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingStateProcessingApi.GetUsingGET18``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET18`: StateProcessingRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingStateProcessingApi.GetUsingGET18`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateProcessingRuleId** | **string** | id of the StateProcessingRule to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET18Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**StateProcessingRule**](StateProcessingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET21

> []StateProcessingRule ListUsingGET21(ctx).XAPIKEY(xAPIKEY).Name(name).Execute()

Retrieve the list of all the StateProcessingRules or get a StateProcessingRule by its name



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
    name := "name_example" // string | name of the StateProcessingRules to retrieve (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingStateProcessingApi.ListUsingGET21(context.Background()).XAPIKEY(xAPIKEY).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingStateProcessingApi.ListUsingGET21``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET21`: []StateProcessingRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingStateProcessingApi.ListUsingGET21`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET21Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **name** | **string** | name of the StateProcessingRules to retrieve | 

### Return type

[**[]StateProcessingRule**](StateProcessingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST16

> StateProcessingRule PostUsingPOST16(ctx).XAPIKEY(xAPIKEY).StateProcessingRule(stateProcessingRule).Execute()

Create a StateProcessingRule



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
    stateProcessingRule := *openapiclient.NewStateProcessingRule("Name_example") // StateProcessingRule | StateProcessingRule to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingStateProcessingApi.PostUsingPOST16(context.Background()).XAPIKEY(xAPIKEY).StateProcessingRule(stateProcessingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingStateProcessingApi.PostUsingPOST16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsingPOST16`: StateProcessingRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingStateProcessingApi.PostUsingPOST16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsingPOST16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **stateProcessingRule** | [**StateProcessingRule**](StateProcessingRule.md) | StateProcessingRule to add | 

### Return type

[**StateProcessingRule**](StateProcessingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST5

> StateProcessingFunctionTestResult TestUsingPOST5(ctx).XAPIKEY(xAPIKEY).TestRequest(testRequest).Execute()

test a  State Processing function



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
    testRequest := *openapiclient.NewStateProcessingFunctionTest(*openapiclient.NewNewData(), 123) // StateProcessingFunctionTest | State Processing function test request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingStateProcessingApi.TestUsingPOST5(context.Background()).XAPIKEY(xAPIKEY).TestRequest(testRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingStateProcessingApi.TestUsingPOST5``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestUsingPOST5`: StateProcessingFunctionTestResult
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingStateProcessingApi.TestUsingPOST5`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestUsingPOST5Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **testRequest** | [**StateProcessingFunctionTest**](StateProcessingFunctionTest.md) | State Processing function test request | 

### Return type

[**StateProcessingFunctionTestResult**](StateProcessingFunctionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT10

> UpdateUsingPUT10(ctx, stateProcessingRuleId).XAPIKEY(xAPIKEY).StateProcessingRule(stateProcessingRule).Execute()

Update a StateProcessingRule



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
    stateProcessingRuleId := "stateProcessingRuleId_example" // string | id of the StateProcessingRule to update
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    stateProcessingRule := *openapiclient.NewStateProcessingRule("Name_example") // StateProcessingRule | updated StateProcessingRule (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingStateProcessingApi.UpdateUsingPUT10(context.Background(), stateProcessingRuleId).XAPIKEY(xAPIKEY).StateProcessingRule(stateProcessingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingStateProcessingApi.UpdateUsingPUT10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**stateProcessingRuleId** | **string** | id of the StateProcessingRule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsingPUT10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **stateProcessingRule** | [**StateProcessingRule**](StateProcessingRule.md) | updated StateProcessingRule | 

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

