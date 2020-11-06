# \CampaignManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCampaignUsingPUT**](CampaignManagementApi.md#CancelCampaignUsingPUT) | **Put** /api/v0/deviceMgt/campaigns/{campaignId}/cancel | Cancel a campaign
[**CreateCampaignUsingPOST**](CampaignManagementApi.md#CreateCampaignUsingPOST) | **Post** /api/v0/deviceMgt/campaigns | Create a campaign
[**DeleteCampaignUsingDELETE**](CampaignManagementApi.md#DeleteCampaignUsingDELETE) | **Delete** /api/v0/deviceMgt/campaigns/{campaignId} | Delete a campaign
[**GetCampaignDetailsUsingGET**](CampaignManagementApi.md#GetCampaignDetailsUsingGET) | **Get** /api/v0/deviceMgt/campaigns/{campaignId}/targets | Get the campaign status per device
[**GetCampaignUsingGET**](CampaignManagementApi.md#GetCampaignUsingGET) | **Get** /api/v0/deviceMgt/campaigns/{campaignId} | Get a campaign
[**ListCampaignsUsingGET**](CampaignManagementApi.md#ListCampaignsUsingGET) | **Get** /api/v0/deviceMgt/campaigns | List all campaigns
[**UpdateCampaignUsingPATCH**](CampaignManagementApi.md#UpdateCampaignUsingPATCH) | **Patch** /api/v0/deviceMgt/campaigns/{campaignId} | Update a campaign



## CancelCampaignUsingPUT

> Campaign CancelCampaignUsingPUT(ctx, campaignId).XAPIKEY(xAPIKEY).Force(force).Execute()

Cancel a campaign



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
    campaignId := "campaignId_example" // string | identifier of a campaign
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | When force is false, the campaign manager prevents devices with status not started from starting, cancel devices with status pending and let devices with status in progress finish properly. When force is true, the campaign manager aborts all pending and in progress sequences immediately. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignManagementApi.CancelCampaignUsingPUT(context.Background(), campaignId).XAPIKEY(xAPIKEY).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignManagementApi.CancelCampaignUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CancelCampaignUsingPUT`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignManagementApi.CancelCampaignUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | identifier of a campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelCampaignUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | When force is false, the campaign manager prevents devices with status not started from starting, cancel devices with status pending and let devices with status in progress finish properly. When force is true, the campaign manager aborts all pending and in progress sequences immediately. | [default to false]

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignUsingPOST

> Campaign CreateCampaignUsingPOST(ctx).XAPIKEY(xAPIKEY).Campaign(campaign).Execute()

Create a campaign



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
    campaign := *openapiclient.NewCampaignDefinition("Name_example", []CampaignOperation{*openapiclient.NewCampaignOperation("Action_example")), *openapiclient.NewCampaignPlanning("EndDate_example", "StartDate_example"), *openapiclient.NewDeviceSelector()) // CampaignDefinition | campaign

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignManagementApi.CreateCampaignUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Campaign(campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignManagementApi.CreateCampaignUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCampaignUsingPOST`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignManagementApi.CreateCampaignUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCampaignUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **campaign** | [**CampaignDefinition**](CampaignDefinition.md) | campaign | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignUsingDELETE

> ResponseEntity DeleteCampaignUsingDELETE(ctx, campaignId).XAPIKEY(xAPIKEY).Force(force).Execute()

Delete a campaign



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
    campaignId := "campaignId_example" // string | identifier of a campaign
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | cancelling mode. When force is false, the campaign manager prevents devices with status not started from starting, cancel devices with status pending and let devices with status in progress finish properly. When force is true, the campaign manager aborts all pending and in progress sequences immediately. (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignManagementApi.DeleteCampaignUsingDELETE(context.Background(), campaignId).XAPIKEY(xAPIKEY).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignManagementApi.DeleteCampaignUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteCampaignUsingDELETE`: ResponseEntity
    fmt.Fprintf(os.Stdout, "Response from `CampaignManagementApi.DeleteCampaignUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | identifier of a campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCampaignUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | cancelling mode. When force is false, the campaign manager prevents devices with status not started from starting, cancel devices with status pending and let devices with status in progress finish properly. When force is true, the campaign manager aborts all pending and in progress sequences immediately. | [default to false]

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignDetailsUsingGET

> CampaignPerTargetList GetCampaignDetailsUsingGET(ctx, campaignId).XAPIKEY(xAPIKEY).Status(status).Size(size).Page(page).Execute()

Get the campaign status per device



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
    campaignId := "campaignId_example" // string | identifier of a campaign
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    status := "status_example" // string | Status of the device (optional)
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignManagementApi.GetCampaignDetailsUsingGET(context.Background(), campaignId).XAPIKEY(xAPIKEY).Status(status).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignManagementApi.GetCampaignDetailsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignDetailsUsingGET`: CampaignPerTargetList
    fmt.Fprintf(os.Stdout, "Response from `CampaignManagementApi.GetCampaignDetailsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | identifier of a campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignDetailsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **status** | **string** | Status of the device | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**CampaignPerTargetList**](CampaignPerTargetList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignUsingGET

> Campaign GetCampaignUsingGET(ctx, campaignId).XAPIKEY(xAPIKEY).Execute()

Get a campaign



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
    campaignId := "campaignId_example" // string | identifier of a campaign
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignManagementApi.GetCampaignUsingGET(context.Background(), campaignId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignManagementApi.GetCampaignUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCampaignUsingGET`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignManagementApi.GetCampaignUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | identifier of a campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCampaignUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaignsUsingGET

> CampaignList ListCampaignsUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Status(status).Sort(sort).Execute()

List all campaigns



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
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)
    status := []string{"Inner_example"} // []string | status list filter (optional)
    sort := []string{"Inner_example"} // []string | sorting list by attributes (supported columns:  id, name, status, created, updated). DefaultValue : -created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignManagementApi.ListCampaignsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Status(status).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignManagementApi.ListCampaignsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCampaignsUsingGET`: CampaignList
    fmt.Fprintf(os.Stdout, "Response from `CampaignManagementApi.ListCampaignsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCampaignsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]
 **status** | [**[]string**](string.md) | status list filter | 
 **sort** | [**[]string**](string.md) | sorting list by attributes (supported columns:  id, name, status, created, updated). DefaultValue : -created | 

### Return type

[**CampaignList**](CampaignList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignUsingPATCH

> Campaign UpdateCampaignUsingPATCH(ctx, campaignId).XAPIKEY(xAPIKEY).Campaign(campaign).Execute()

Update a campaign



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
    campaignId := "campaignId_example" // string | identifier of a campaign
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    campaign := *openapiclient.NewCampaignUpdating() // CampaignUpdating | campaign

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CampaignManagementApi.UpdateCampaignUsingPATCH(context.Background(), campaignId).XAPIKEY(xAPIKEY).Campaign(campaign).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CampaignManagementApi.UpdateCampaignUsingPATCH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateCampaignUsingPATCH`: Campaign
    fmt.Fprintf(os.Stdout, "Response from `CampaignManagementApi.UpdateCampaignUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string** | identifier of a campaign | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCampaignUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **campaign** | [**CampaignUpdating**](CampaignUpdating.md) | campaign | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

