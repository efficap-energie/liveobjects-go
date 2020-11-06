# \TriggersAndActionsApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST2**](TriggersAndActionsApi.md#CreateUsingPOST2) | **Post** /api/v1/event2action/actionPolicies | Create an ActionPolicy
[**DeleteUsingDELETE13**](TriggersAndActionsApi.md#DeleteUsingDELETE13) | **Delete** /api/v1/event2action/actionPolicies/{policyId} | Delete an ActionPolicy
[**ListUsingGET15**](TriggersAndActionsApi.md#ListUsingGET15) | **Get** /api/v1/event2action/actionPolicies | List ActionPolicies
[**RetrieveUsingGET2**](TriggersAndActionsApi.md#RetrieveUsingGET2) | **Get** /api/v1/event2action/actionPolicies/{policyId} | Retrieve an ActionPolicy
[**UpsertUsingPUT1**](TriggersAndActionsApi.md#UpsertUsingPUT1) | **Put** /api/v1/event2action/actionPolicies/{policyId} | Create or update an ActionPolicy



## CreateUsingPOST2

> ActionPolicy CreateUsingPOST2(ctx).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()

Create an ActionPolicy



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
    actionPolicy := *openapiclient.NewActionPolicy(*openapiclient.NewActions(), "Name_example", *openapiclient.NewActionTriggers()) // ActionPolicy | ActionPolicy to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TriggersAndActionsApi.CreateUsingPOST2(context.Background()).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAndActionsApi.CreateUsingPOST2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsingPOST2`: ActionPolicy
    fmt.Fprintf(os.Stdout, "Response from `TriggersAndActionsApi.CreateUsingPOST2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsingPOST2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **actionPolicy** | [**ActionPolicy**](ActionPolicy.md) | ActionPolicy to add | 

### Return type

[**ActionPolicy**](ActionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE13

> DeleteUsingDELETE13(ctx, policyId).XAPIKEY(xAPIKEY).Execute()

Delete an ActionPolicy



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
    policyId := "policyId_example" // string | id of the ActionPolicy to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TriggersAndActionsApi.DeleteUsingDELETE13(context.Background(), policyId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAndActionsApi.DeleteUsingDELETE13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | id of the ActionPolicy to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE13Request struct via the builder pattern


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


## ListUsingGET15

> []ActionPolicy ListUsingGET15(ctx).XAPIKEY(xAPIKEY).TriggerType(triggerType).Execute()

List ActionPolicies



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
    triggerType := "triggerType_example" // string | Filter on action policy trigger type. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TriggersAndActionsApi.ListUsingGET15(context.Background()).XAPIKEY(xAPIKEY).TriggerType(triggerType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAndActionsApi.ListUsingGET15``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET15`: []ActionPolicy
    fmt.Fprintf(os.Stdout, "Response from `TriggersAndActionsApi.ListUsingGET15`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET15Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **triggerType** | **string** | Filter on action policy trigger type. | 

### Return type

[**[]ActionPolicy**](ActionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUsingGET2

> ActionPolicy RetrieveUsingGET2(ctx, policyId).XAPIKEY(xAPIKEY).Execute()

Retrieve an ActionPolicy



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
    policyId := "policyId_example" // string | id of the ActionPolicy to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TriggersAndActionsApi.RetrieveUsingGET2(context.Background(), policyId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAndActionsApi.RetrieveUsingGET2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUsingGET2`: ActionPolicy
    fmt.Fprintf(os.Stdout, "Response from `TriggersAndActionsApi.RetrieveUsingGET2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | id of the ActionPolicy to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUsingGET2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**ActionPolicy**](ActionPolicy.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUsingPUT1

> UpsertUsingPUT1(ctx, policyId).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()

Create or update an ActionPolicy



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
    policyId := "policyId_example" // string | id of the ActionPolicy to save
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    actionPolicy := *openapiclient.NewActionPolicy(*openapiclient.NewActions(), "Name_example", *openapiclient.NewActionTriggers()) // ActionPolicy | ActionPolicy to save (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.TriggersAndActionsApi.UpsertUsingPUT1(context.Background(), policyId).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `TriggersAndActionsApi.UpsertUsingPUT1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | id of the ActionPolicy to save | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUsingPUT1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **actionPolicy** | [**ActionPolicy**](ActionPolicy.md) | ActionPolicy to save | 

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

