# \DeprecatedTriggersAndActionsApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST1**](DeprecatedTriggersAndActionsApi.md#CreateUsingPOST1) | **Post** /api/v0/event2action/actionPolicies | Create an ActionPolicy
[**DeleteUsingDELETE12**](DeprecatedTriggersAndActionsApi.md#DeleteUsingDELETE12) | **Delete** /api/v0/event2action/actionPolicies/{policyId} | Delete an ActionPolicy
[**ListUsingGET14**](DeprecatedTriggersAndActionsApi.md#ListUsingGET14) | **Get** /api/v0/event2action/actionPolicies | List ActionPolicies
[**RetrieveUsingGET1**](DeprecatedTriggersAndActionsApi.md#RetrieveUsingGET1) | **Get** /api/v0/event2action/actionPolicies/{policyId} | Retrieve an ActionPolicy
[**UpsertUsingPUT**](DeprecatedTriggersAndActionsApi.md#UpsertUsingPUT) | **Put** /api/v0/event2action/actionPolicies/{policyId} | Create or update an ActionPolicy



## CreateUsingPOST1

> ActionPolicyV0 CreateUsingPOST1(ctx).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()

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
    actionPolicy := *openapiclient.NewActionPolicyV0() // ActionPolicyV0 | ActionPolicy to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedTriggersAndActionsApi.CreateUsingPOST1(context.Background()).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedTriggersAndActionsApi.CreateUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsingPOST1`: ActionPolicyV0
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedTriggersAndActionsApi.CreateUsingPOST1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **actionPolicy** | [**ActionPolicyV0**](ActionPolicyV0.md) | ActionPolicy to add | 

### Return type

[**ActionPolicyV0**](ActionPolicyV0.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE12

> DeleteUsingDELETE12(ctx, policyId).XAPIKEY(xAPIKEY).Execute()

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
    policyId := "policyId_example" // string | id of the action to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedTriggersAndActionsApi.DeleteUsingDELETE12(context.Background(), policyId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedTriggersAndActionsApi.DeleteUsingDELETE12``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | id of the action to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE12Request struct via the builder pattern


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


## ListUsingGET14

> PageableActionPolicyV0 ListUsingGET14(ctx).XAPIKEY(xAPIKEY).TriggerType(triggerType).Size(size).Page(page).Execute()

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
    triggerType := "triggerType_example" // string | Filter on ActionPolicy trigger type. Can be either \"rules\" or \"messages\". (optional)
    size := 987 // int32 | the maximum number of items per page (optional) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedTriggersAndActionsApi.ListUsingGET14(context.Background()).XAPIKEY(xAPIKEY).TriggerType(triggerType).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedTriggersAndActionsApi.ListUsingGET14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET14`: PageableActionPolicyV0
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedTriggersAndActionsApi.ListUsingGET14`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET14Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **triggerType** | **string** | Filter on ActionPolicy trigger type. Can be either \&quot;rules\&quot; or \&quot;messages\&quot;. | 
 **size** | **int32** | the maximum number of items per page (optional) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**PageableActionPolicyV0**](Pageable«ActionPolicyV0».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUsingGET1

> ActionPolicyV0 RetrieveUsingGET1(ctx, policyId).XAPIKEY(xAPIKEY).Execute()

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
    policyId := "policyId_example" // string | id of the action to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedTriggersAndActionsApi.RetrieveUsingGET1(context.Background(), policyId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedTriggersAndActionsApi.RetrieveUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUsingGET1`: ActionPolicyV0
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedTriggersAndActionsApi.RetrieveUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | id of the action to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**ActionPolicyV0**](ActionPolicyV0.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpsertUsingPUT

> UpsertUsingPUT(ctx, policyId).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()

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
    policyId := "policyId_example" // string | id of the action to save
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    actionPolicy := *openapiclient.NewActionPolicyV0() // ActionPolicyV0 | ActionPolicy to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedTriggersAndActionsApi.UpsertUsingPUT(context.Background(), policyId).XAPIKEY(xAPIKEY).ActionPolicy(actionPolicy).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedTriggersAndActionsApi.UpsertUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**policyId** | **string** | id of the action to save | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpsertUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **actionPolicy** | [**ActionPolicyV0**](ActionPolicyV0.md) | ActionPolicy to add | 

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

