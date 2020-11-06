# \DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelResourceUpdateUsingPOST**](DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.md#CancelResourceUpdateUsingPOST) | **Post** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/cancelUpdate | Cancel asset resource update
[**GetAllAssetResourcesUsingGET**](DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.md#GetAllAssetResourcesUsingGET) | **Get** /api/v0/rm/asset/{assetIdNamespace}/{assetId} | List the asset&#39;s resources (use /api/v1/deviceMgt/devices/{deviceId}/resources instead)
[**GetLastUpdateUsingGET**](DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.md#GetLastUpdateUsingGET) | **Get** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/update | Get the asset&#39;s resources update status (use /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId}/lastUpdate instead)
[**GetUpdateHistoryUsingGET**](DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.md#GetUpdateHistoryUsingGET) | **Get** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/update/history | Get the asset&#39;s resources update history (use /api/v1/deviceMgt/devices/{deviceId}/firmwareUpdates instead)
[**SetAssetTargetResourceVersionUsingPUT**](DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.md#SetAssetTargetResourceVersionUsingPUT) | **Put** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/resource/{resourceId}/targetversion | Set asset&#39;s target resource version (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} instead)



## CancelResourceUpdateUsingPOST

> CancelResourceUpdateUsingPOST(ctx, assetIdNamespace, assetId).XAPIKEY(xAPIKEY).Force(force).Execute()

Cancel asset resource update



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
    assetIdNamespace := "assetIdNamespace_example" // string | the asset id namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | the asset id. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | force update status change (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.CancelResourceUpdateUsingPOST(context.Background(), assetIdNamespace, assetId).XAPIKEY(xAPIKEY).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.CancelResourceUpdateUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string** | the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiCancelResourceUpdateUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | force update status change | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllAssetResourcesUsingGET

> PageableAssetResourceWeb GetAllAssetResourcesUsingGET(ctx, assetIdNamespace, assetId).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()

List the asset's resources (use /api/v1/deviceMgt/devices/{deviceId}/resources instead)



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
    assetIdNamespace := "assetIdNamespace_example" // string | the asset id namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | the asset id. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetAllAssetResourcesUsingGET(context.Background(), assetIdNamespace, assetId).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetAllAssetResourcesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllAssetResourcesUsingGET`: PageableAssetResourceWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetAllAssetResourcesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string** | the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAllAssetResourcesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**PageableAssetResourceWeb**](Pageable«AssetResourceWeb».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastUpdateUsingGET

> ResourceUpdateWeb GetLastUpdateUsingGET(ctx, assetIdNamespace, assetId).XAPIKEY(xAPIKEY).ResourceId(resourceId).Execute()

Get the asset's resources update status (use /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId}/lastUpdate instead)



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
    assetIdNamespace := "assetIdNamespace_example" // string | the asset id namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | the asset id. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetLastUpdateUsingGET(context.Background(), assetIdNamespace, assetId).XAPIKEY(xAPIKEY).ResourceId(resourceId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetLastUpdateUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastUpdateUsingGET`: ResourceUpdateWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetLastUpdateUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string** | the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastUpdateUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **resourceId** | **string** | the resource id. Expected string (max 255 characters) | 

### Return type

[**ResourceUpdateWeb**](ResourceUpdateWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateHistoryUsingGET

> PageableResourceUpdateWeb GetUpdateHistoryUsingGET(ctx, assetIdNamespace, assetId).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()

Get the asset's resources update history (use /api/v1/deviceMgt/devices/{deviceId}/firmwareUpdates instead)



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
    assetIdNamespace := "assetIdNamespace_example" // string | the asset id namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | the asset id. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetUpdateHistoryUsingGET(context.Background(), assetIdNamespace, assetId).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetUpdateHistoryUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUpdateHistoryUsingGET`: PageableResourceUpdateWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.GetUpdateHistoryUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string** | the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUpdateHistoryUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**PageableResourceUpdateWeb**](Pageable«ResourceUpdateWeb».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAssetTargetResourceVersionUsingPUT

> AssetResourceWeb SetAssetTargetResourceVersionUsingPUT(ctx, assetIdNamespace, assetId, resourceId).XAPIKEY(xAPIKEY).NotifyTo(notifyTo).TargetVersion(targetVersion).Execute()

Set asset's target resource version (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} instead)



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
    assetIdNamespace := "assetIdNamespace_example" // string | the asset id namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | the asset id. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    notifyTo := "notifyTo_example" // string | fifo used to relay status notification ex: fifo/~notif. A valid 'notify to' starts with one of ['fifo/', 'pubsub/','router/']and max length is 255 (optional)
    targetVersion := *openapiclient.NewAssetSetTargetResourceVersionReqWeb("TargetVersion_example") // AssetSetTargetResourceVersionReqWeb | the target resource version. Expected string (max 255 characters) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.SetAssetTargetResourceVersionUsingPUT(context.Background(), assetIdNamespace, assetId, resourceId).XAPIKEY(xAPIKEY).NotifyTo(notifyTo).TargetVersion(targetVersion).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.SetAssetTargetResourceVersionUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetAssetTargetResourceVersionUsingPUT`: AssetResourceWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementResourceV0UseDeviceManagementResourceV1InsteadApi.SetAssetTargetResourceVersionUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string** | the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAssetTargetResourceVersionUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAPIKEY** | **string** | a valid API key | 
 **notifyTo** | **string** | fifo used to relay status notification ex: fifo/~notif. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;,&#39;router/&#39;]and max length is 255 | 
 **targetVersion** | [**AssetSetTargetResourceVersionReqWeb**](AssetSetTargetResourceVersionReqWeb.md) | the target resource version. Expected string (max 255 characters) | 

### Return type

[**AssetResourceWeb**](AssetResourceWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

