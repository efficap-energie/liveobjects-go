# \EventProcessingMatchingApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE18**](EventProcessingMatchingApi.md#DeleteUsingDELETE18) | **Delete** /api/v0/eventprocessing/matching-rule/{matchingRuleId} | Delete a MatchingRule
[**GetUsingGET17**](EventProcessingMatchingApi.md#GetUsingGET17) | **Get** /api/v0/eventprocessing/matching-rule/{matchingRuleId} | Retrieve a MatchingRule
[**ListUsingGET20**](EventProcessingMatchingApi.md#ListUsingGET20) | **Get** /api/v0/eventprocessing/matching-rule | Retrieve the list of all the MatchingRules or get a MatchingRule by its name
[**PostUsingPOST15**](EventProcessingMatchingApi.md#PostUsingPOST15) | **Post** /api/v0/eventprocessing/matching-rule | Create a MatchingRule
[**TestUsingPOST4**](EventProcessingMatchingApi.md#TestUsingPOST4) | **Post** /api/v0/eventprocessing/matching-rule/test | Test a JsonLogic predicate with some data sample
[**UpdateUsingPUT9**](EventProcessingMatchingApi.md#UpdateUsingPUT9) | **Put** /api/v0/eventprocessing/matching-rule/{matchingRuleId} | Update a MatchingRule



## DeleteUsingDELETE18

> DeleteUsingDELETE18(ctx, matchingRuleId).XAPIKEY(xAPIKEY).Execute()

Delete a MatchingRule



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
    matchingRuleId := "matchingRuleId_example" // string | id of the MatchingRule to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingMatchingApi.DeleteUsingDELETE18(context.Background(), matchingRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingMatchingApi.DeleteUsingDELETE18``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleId** | **string** | id of the MatchingRule to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE18Request struct via the builder pattern


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


## GetUsingGET17

> MatchingRule GetUsingGET17(ctx, matchingRuleId).XAPIKEY(xAPIKEY).Execute()

Retrieve a MatchingRule



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
    matchingRuleId := "matchingRuleId_example" // string | id of the MatchingRule to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingMatchingApi.GetUsingGET17(context.Background(), matchingRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingMatchingApi.GetUsingGET17``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET17`: MatchingRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingMatchingApi.GetUsingGET17`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleId** | **string** | id of the MatchingRule to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET17Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**MatchingRule**](MatchingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET20

> []MatchingRule ListUsingGET20(ctx).XAPIKEY(xAPIKEY).Name(name).Execute()

Retrieve the list of all the MatchingRules or get a MatchingRule by its name



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
    name := "name_example" // string | name of the MatchingRule to retrieve (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingMatchingApi.ListUsingGET20(context.Background()).XAPIKEY(xAPIKEY).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingMatchingApi.ListUsingGET20``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET20`: []MatchingRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingMatchingApi.ListUsingGET20`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET20Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **name** | **string** | name of the MatchingRule to retrieve | 

### Return type

[**[]MatchingRule**](MatchingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST15

> MatchingRule PostUsingPOST15(ctx).XAPIKEY(xAPIKEY).MatchRule(matchRule).Execute()

Create a MatchingRule



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
    matchRule := *openapiclient.NewMatchingRule("Name_example") // MatchingRule | MatchingRule to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingMatchingApi.PostUsingPOST15(context.Background()).XAPIKEY(xAPIKEY).MatchRule(matchRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingMatchingApi.PostUsingPOST15``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsingPOST15`: MatchingRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingMatchingApi.PostUsingPOST15`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsingPOST15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **matchRule** | [**MatchingRule**](MatchingRule.md) | MatchingRule to add | 

### Return type

[**MatchingRule**](MatchingRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST4

> DataMatchResult TestUsingPOST4(ctx).XAPIKEY(xAPIKEY).DataMatchTest(dataMatchTest).Execute()

Test a JsonLogic predicate with some data sample



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
    dataMatchTest := *openapiclient.NewDataMatchTest(123, 123) // DataMatchTest | jsonLogic predicate and json data of the evaluation test you want to perform (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingMatchingApi.TestUsingPOST4(context.Background()).XAPIKEY(xAPIKEY).DataMatchTest(dataMatchTest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingMatchingApi.TestUsingPOST4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestUsingPOST4`: DataMatchResult
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingMatchingApi.TestUsingPOST4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestUsingPOST4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **dataMatchTest** | [**DataMatchTest**](DataMatchTest.md) | jsonLogic predicate and json data of the evaluation test you want to perform | 

### Return type

[**DataMatchResult**](DataMatchResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT9

> UpdateUsingPUT9(ctx, matchingRuleId).XAPIKEY(xAPIKEY).MatchingRule(matchingRule).Execute()

Update a MatchingRule



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
    matchingRuleId := "matchingRuleId_example" // string | id of the MatchingRule to update
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    matchingRule := *openapiclient.NewMatchingRule("Name_example") // MatchingRule | updated MatchingRule (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingMatchingApi.UpdateUsingPUT9(context.Background(), matchingRuleId).XAPIKEY(xAPIKEY).MatchingRule(matchingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingMatchingApi.UpdateUsingPUT9``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**matchingRuleId** | **string** | id of the MatchingRule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsingPUT9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **matchingRule** | [**MatchingRule**](MatchingRule.md) | updated MatchingRule | 

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

