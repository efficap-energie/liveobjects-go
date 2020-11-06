# \EventProcessingActivityApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteUsingDELETE14**](EventProcessingActivityApi.md#DeleteUsingDELETE14) | **Delete** /api/v0/eventprocessing/activity/rules/{activityRuleId} | Delete an ActivityRule
[**GetStatesUsingGET**](EventProcessingActivityApi.md#GetStatesUsingGET) | **Get** /api/v0/eventprocessing/activity/states | Retrieve the list of all the ActivityStates linked to a specific device and/or rule
[**GetUsingGET13**](EventProcessingActivityApi.md#GetUsingGET13) | **Get** /api/v0/eventprocessing/activity/rules/{activityRuleId} | Retrieve an ActivityRule
[**ListUsingGET16**](EventProcessingActivityApi.md#ListUsingGET16) | **Get** /api/v0/eventprocessing/activity/rules | Retrieve the list of all the ActivityRules or get an ActivityRule by its name
[**MuteUsingPUT**](EventProcessingActivityApi.md#MuteUsingPUT) | **Put** /api/v0/eventprocessing/activity/states/mute | Mute or reset nextAlarm of ActivityStates targeted by a specific deviceId/activityRuleId
[**PostUsingPOST13**](EventProcessingActivityApi.md#PostUsingPOST13) | **Post** /api/v0/eventprocessing/activity/rules | Create an ActivityRule
[**UpdateUsingPUT7**](EventProcessingActivityApi.md#UpdateUsingPUT7) | **Put** /api/v0/eventprocessing/activity/rules/{activityRuleId} | Update an ActivityRule



## DeleteUsingDELETE14

> DeleteUsingDELETE14(ctx, activityRuleId).XAPIKEY(xAPIKEY).Execute()

Delete an ActivityRule



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
    activityRuleId := "activityRuleId_example" // string | id of the ActivityRule to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingActivityApi.DeleteUsingDELETE14(context.Background(), activityRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingActivityApi.DeleteUsingDELETE14``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityRuleId** | **string** | id of the ActivityRule to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE14Request struct via the builder pattern


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


## GetStatesUsingGET

> []ActivityState GetStatesUsingGET(ctx).XAPIKEY(xAPIKEY).DeviceId(deviceId).ActivityRuleId(activityRuleId).Limit(limit).BookmarkDeviceId(bookmarkDeviceId).Execute()

Retrieve the list of all the ActivityStates linked to a specific device and/or rule



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
    deviceId := "deviceId_example" // string | id of the device targeted by the states to retrieve. At least one of deviceId/activityRuleId must be set. (optional)
    activityRuleId := "activityRuleId_example" // string | id of the rule targeted by the states to retrieve. At least one of deviceId/activityRuleId must be set. (optional)
    limit := 987 // int32 | when listing by activityRuleId, thousands of devices could be targeted. Indicates the number of AvtivityStates to return (one per targeted device). Default is 20, max is 1000. To get next results use bookmarkDeviceId field. (optional)
    bookmarkDeviceId := "bookmarkDeviceId_example" // string | 'bookmark' of previous listing by activityRuleId request : this is the last deviceId retrieved in previous request. If null, first ActivtyStates results will be returned. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingActivityApi.GetStatesUsingGET(context.Background()).XAPIKEY(xAPIKEY).DeviceId(deviceId).ActivityRuleId(activityRuleId).Limit(limit).BookmarkDeviceId(bookmarkDeviceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingActivityApi.GetStatesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetStatesUsingGET`: []ActivityState
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingActivityApi.GetStatesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetStatesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **deviceId** | **string** | id of the device targeted by the states to retrieve. At least one of deviceId/activityRuleId must be set. | 
 **activityRuleId** | **string** | id of the rule targeted by the states to retrieve. At least one of deviceId/activityRuleId must be set. | 
 **limit** | **int32** | when listing by activityRuleId, thousands of devices could be targeted. Indicates the number of AvtivityStates to return (one per targeted device). Default is 20, max is 1000. To get next results use bookmarkDeviceId field. | 
 **bookmarkDeviceId** | **string** | &#39;bookmark&#39; of previous listing by activityRuleId request : this is the last deviceId retrieved in previous request. If null, first ActivtyStates results will be returned. | 

### Return type

[**[]ActivityState**](ActivityState.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUsingGET13

> ActivityRule GetUsingGET13(ctx, activityRuleId).XAPIKEY(xAPIKEY).Execute()

Retrieve an ActivityRule



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
    activityRuleId := "activityRuleId_example" // string | id of the ActivityRule to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingActivityApi.GetUsingGET13(context.Background(), activityRuleId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingActivityApi.GetUsingGET13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET13`: ActivityRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingActivityApi.GetUsingGET13`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityRuleId** | **string** | id of the ActivityRule to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**ActivityRule**](ActivityRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET16

> []ActivityRule ListUsingGET16(ctx).XAPIKEY(xAPIKEY).Name(name).Execute()

Retrieve the list of all the ActivityRules or get an ActivityRule by its name



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
    name := "name_example" // string | name of the ActivityRule to retrieve (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingActivityApi.ListUsingGET16(context.Background()).XAPIKEY(xAPIKEY).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingActivityApi.ListUsingGET16``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET16`: []ActivityRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingActivityApi.ListUsingGET16`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET16Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **name** | **string** | name of the ActivityRule to retrieve | 

### Return type

[**[]ActivityRule**](ActivityRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## MuteUsingPUT

> int64 MuteUsingPUT(ctx).XAPIKEY(xAPIKEY).NextAlarmRequest(nextAlarmRequest).Execute()

Mute or reset nextAlarm of ActivityStates targeted by a specific deviceId/activityRuleId



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
    nextAlarmRequest := *openapiclient.NewActivityStateMuteRequest("NextAlarmOrder_example") // ActivityStateMuteRequest | nextAlarmRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingActivityApi.MuteUsingPUT(context.Background()).XAPIKEY(xAPIKEY).NextAlarmRequest(nextAlarmRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingActivityApi.MuteUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `MuteUsingPUT`: int64
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingActivityApi.MuteUsingPUT`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiMuteUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **nextAlarmRequest** | [**ActivityStateMuteRequest**](ActivityStateMuteRequest.md) | nextAlarmRequest | 

### Return type

**int64**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST13

> ActivityRule PostUsingPOST13(ctx).XAPIKEY(xAPIKEY).ActivityRule(activityRule).Execute()

Create an ActivityRule



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
    activityRule := *openapiclient.NewActivityRule("Name_example", *openapiclient.NewSilentPolicy("Duration_example"), *openapiclient.NewTargets()) // ActivityRule | ActivityRule to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingActivityApi.PostUsingPOST13(context.Background()).XAPIKEY(xAPIKEY).ActivityRule(activityRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingActivityApi.PostUsingPOST13``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsingPOST13`: ActivityRule
    fmt.Fprintf(os.Stdout, "Response from `EventProcessingActivityApi.PostUsingPOST13`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsingPOST13Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **activityRule** | [**ActivityRule**](ActivityRule.md) | ActivityRule to add | 

### Return type

[**ActivityRule**](ActivityRule.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUsingPUT7

> UpdateUsingPUT7(ctx, activityRuleId).XAPIKEY(xAPIKEY).ActivityRule(activityRule).Execute()

Update an ActivityRule



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
    activityRuleId := "activityRuleId_example" // string | id of the ActivityRule to update
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    activityRule := *openapiclient.NewActivityRule("Name_example", *openapiclient.NewSilentPolicy("Duration_example"), *openapiclient.NewTargets()) // ActivityRule | updated ActivityRule (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.EventProcessingActivityApi.UpdateUsingPUT7(context.Background(), activityRuleId).XAPIKEY(xAPIKEY).ActivityRule(activityRule).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `EventProcessingActivityApi.UpdateUsingPUT7``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**activityRuleId** | **string** | id of the ActivityRule to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUsingPUT7Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **activityRule** | [**ActivityRule**](ActivityRule.md) | updated ActivityRule | 

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

