# \EventProcessingFiringApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE16**](EventProcessingFiringApi.md#DeleteUsingDELETE16) | **Delete** /api/v0/eventprocessing/firing-rule/{firingRuleId} | Delete a FiringRule
[**GetFiringGuardUsingGET**](EventProcessingFiringApi.md#GetFiringGuardUsingGET) | **Get** /api/v0/eventprocessing/firing-guard/{firingGuardId} | Get a FiringGuard
[**GetFiringGuardsUsingPOST**](EventProcessingFiringApi.md#GetFiringGuardsUsingPOST) | **Post** /api/v0/eventprocessing/firing-guard/search | Get FiringGuards linked to a FiringRule, and where FiringGuards selection criteria match.
[**GetUsingGET15**](EventProcessingFiringApi.md#GetUsingGET15) | **Get** /api/v0/eventprocessing/firing-rule/{firingRuleId} | Retrieve a FiringRule
[**ListUsingGET18**](EventProcessingFiringApi.md#ListUsingGET18) | **Get** /api/v0/eventprocessing/firing-rule | Retrieve the list of all the FiringRules or get a FiringRule by its name
[**PostUsingPOST14**](EventProcessingFiringApi.md#PostUsingPOST14) | **Post** /api/v0/eventprocessing/firing-rule | Create a FiringRule
[**RemoveFiringGuardUsingDELETE**](EventProcessingFiringApi.md#RemoveFiringGuardUsingDELETE) | **Delete** /api/v0/eventprocessing/firing-guard/{firingGuardId} | Remove a FiringGuard
[**RemoveFiringGuardsUsingDELETE**](EventProcessingFiringApi.md#RemoveFiringGuardsUsingDELETE) | **Delete** /api/v0/eventprocessing/firing-guard | Remove the FiringGuards linked to FiringRule, and where FiringGuards selection criteria match.
[**UpdateUsingPUT8**](EventProcessingFiringApi.md#UpdateUsingPUT8) | **Put** /api/v0/eventprocessing/firing-rule/{firingRuleId} | Update a FiringRule



## DeleteUsingDELETE16

> DeleteUsingDELETE16(ctx, firingRuleId).XAPIKEY(xAPIKEY).Execute()

Delete a FiringRule



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
    firingRuleId := "firingRuleId_example" // string | id of the FiringRule to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.DeleteUsingDELETE16(context.Background(), firingRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.DeleteUsingDELETE16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingRuleId** | **string** | id of the FiringRule to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE16Request struct via the builder pattern


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


## GetFiringGuardUsingGET

> FiringGuard GetFiringGuardUsingGET(ctx, firingGuardId).XAPIKEY(xAPIKEY).Execute()

Get a FiringGuard



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
    firingGuardId := "firingGuardId_example" // string | id of the FiringGuard
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.GetFiringGuardUsingGET(context.Background(), firingGuardId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.GetFiringGuardUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFiringGuardUsingGET`: FiringGuard
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingFiringApi.GetFiringGuardUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingGuardId** | **string** | id of the FiringGuard | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFiringGuardUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**FiringGuard**](FiringGuard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetFiringGuardsUsingPOST

> []FiringGuard GetFiringGuardsUsingPOST(ctx).XAPIKEY(xAPIKEY).SearchParam(searchParam).Execute()

Get FiringGuards linked to a FiringRule, and where FiringGuards selection criteria match.



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
    searchParam := *openapiclient.NewFiringGuardGetRequest("FiringRuleId_example") // FiringGuardGetRequest | Search parameters to select FiringGuard (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.GetFiringGuardsUsingPOST(context.Background()).XAPIKEY(xAPIKEY).SearchParam(searchParam).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.GetFiringGuardsUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetFiringGuardsUsingPOST`: []FiringGuard
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingFiringApi.GetFiringGuardsUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetFiringGuardsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **searchParam** | [**FiringGuardGetRequest**](FiringGuardGetRequest.md) | Search parameters to select FiringGuard | 

### Return type

[**[]FiringGuard**](FiringGuard.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET15

> FiringRule GetUsingGET15(ctx, firingRuleId).XAPIKEY(xAPIKEY).Execute()

Retrieve a FiringRule



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
    firingRuleId := "firingRuleId_example" // string | id of the FiringRule to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.GetUsingGET15(context.Background(), firingRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.GetUsingGET15``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET15`: FiringRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingFiringApi.GetUsingGET15`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingRuleId** | **string** | id of the FiringRule to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**FiringRule**](FiringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET18

> []FiringRule ListUsingGET18(ctx).XAPIKEY(xAPIKEY).Name(name).Execute()

Retrieve the list of all the FiringRules or get a FiringRule by its name



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
    name := "name_example" // string | name of the FiringRule to retrieve (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.ListUsingGET18(context.Background()).XAPIKEY(xAPIKEY).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.ListUsingGET18``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET18`: []FiringRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingFiringApi.ListUsingGET18`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET18Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **name** | **string** | name of the FiringRule to retrieve | 

### Return type

[**[]FiringRule**](FiringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST14

> FiringRule PostUsingPOST14(ctx).XAPIKEY(xAPIKEY).FiringRule(firingRule).Execute()

Create a FiringRule



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
    firingRule := *openapiclient.NewFiringRule("FiringType_example", "Name_example") // FiringRule | FiringRule to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.PostUsingPOST14(context.Background()).XAPIKEY(xAPIKEY).FiringRule(firingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.PostUsingPOST14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsingPOST14`: FiringRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingFiringApi.PostUsingPOST14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsingPOST14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **firingRule** | [**FiringRule**](FiringRule.md) | FiringRule to add | 

### Return type

[**FiringRule**](FiringRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RemoveFiringGuardUsingDELETE

> RemoveFiringGuardUsingDELETE(ctx, firingGuardId).XAPIKEY(xAPIKEY).Execute()

Remove a FiringGuard



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
    firingGuardId := "firingGuardId_example" // string | id of the FiringGuard to remove
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.RemoveFiringGuardUsingDELETE(context.Background(), firingGuardId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.RemoveFiringGuardUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingGuardId** | **string** | id of the FiringGuard to remove | 

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFiringGuardUsingDELETERequest struct via the builder pattern


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


## RemoveFiringGuardsUsingDELETE

> int32 RemoveFiringGuardsUsingDELETE(ctx).XAPIKEY(xAPIKEY).ResetParams(resetParams).Execute()

Remove the FiringGuards linked to FiringRule, and where FiringGuards selection criteria match.



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
    resetParams := *openapiclient.NewFiringGuardResetRequest("FiringRuleId_example") // FiringGuardResetRequest | search parameters to select FiringGuard to remove (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.RemoveFiringGuardsUsingDELETE(context.Background()).XAPIKEY(xAPIKEY).ResetParams(resetParams).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.RemoveFiringGuardsUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RemoveFiringGuardsUsingDELETE`: int32
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingFiringApi.RemoveFiringGuardsUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRemoveFiringGuardsUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **resetParams** | [**FiringGuardResetRequest**](FiringGuardResetRequest.md) | search parameters to select FiringGuard to remove | 

### Return type

**int32**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT8

> UpdateUsingPUT8(ctx, firingRuleId).XAPIKEY(xAPIKEY).FiringRule(firingRule).Execute()

Update a FiringRule



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
    firingRuleId := "firingRuleId_example" // string | id of the FiringRule to update
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    firingRule := *openapiclient.NewFiringRule("FiringType_example", "Name_example") // FiringRule | updated FiringRule (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingFiringApi.UpdateUsingPUT8(context.Background(), firingRuleId).XAPIKEY(xAPIKEY).FiringRule(firingRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingFiringApi.UpdateUsingPUT8``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**firingRuleId** | **string** | id of the FiringRule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsingPUT8Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **firingRule** | [**FiringRule**](FiringRule.md) | updated FiringRule | 

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

