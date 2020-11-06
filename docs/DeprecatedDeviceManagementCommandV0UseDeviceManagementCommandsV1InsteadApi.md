# \DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommandUsingPOST**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#AddCommandUsingPOST) | **Post** /api/v0/assets/{assetNamespace}/{assetId}/commands | Register a new command
[**DeleteCommandUsingDELETE**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#DeleteCommandUsingDELETE) | **Delete** /api/v0/commands/{commandId} | Delete a specific command
[**GetAssetCommandsUsingGET**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#GetAssetCommandsUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId}/commands | Get a list of commands targeting a specific asset
[**GetCommandStatusUsingGET**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#GetCommandStatusUsingGET) | **Get** /api/v0/commands/{commandId}/status | Get the status of a specific command
[**GetCommandUsingGET**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#GetCommandUsingGET) | **Get** /api/v0/commands/{commandId} | Get a specific command
[**ListCommandsUsingGET**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#ListCommandsUsingGET) | **Get** /api/v0/commands | List registered commands
[**SetCommandStatusUsingPUT**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#SetCommandStatusUsingPUT) | **Put** /api/v0/commands/{commandId}/status | Update the status of specific command 



## AddCommandUsingPOST

> CommandV0 AddCommandUsingPOST(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).NotifyTo(notifyTo).Command(command).Execute()

Register a new command



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
    assetNamespace := "assetNamespace_example" // string | target asset namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | target asset identifier. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    notifyTo := "notifyTo_example" // string | fifo used to relay status notification ex: fifo/~notif. A valid 'notify to' starts with one of ['fifo/', 'pubsub/', 'router/'] and max length is 255 (optional)
    command := *openapiclient.NewAssetCommandWeb() // AssetCommandWeb | command to send to the device (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.AddCommandUsingPOST(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).NotifyTo(notifyTo).Command(command).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.AddCommandUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCommandUsingPOST`: CommandV0
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.AddCommandUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | target asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | target asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCommandUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **notifyTo** | **string** | fifo used to relay status notification ex: fifo/~notif. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;, &#39;router/&#39;] and max length is 255 | 
 **command** | [**AssetCommandWeb**](AssetCommandWeb.md) | command to send to the device | 

### Return type

[**CommandV0**](CommandV0.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommandUsingDELETE

> DeleteCommandUsingDELETE(ctx, commandId).XAPIKEY(xAPIKEY).Execute()

Delete a specific command



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
    commandId := "commandId_example" // string | identifier of specific command. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.DeleteCommandUsingDELETE(context.Background(), commandId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.DeleteCommandUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | identifier of specific command. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCommandUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

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


## GetAssetCommandsUsingGET

> PageableCommandV0 GetAssetCommandsUsingGET(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Execute()

Get a list of commands targeting a specific asset



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
    assetNamespace := "assetNamespace_example" // string | requested commands target asset namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | requested commands target asset identifier. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    size := 987 // int64 | page size (optional) (default to 20)
    page := 987 // int64 | page (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | sorting list by attributes (supported columns:  id, status, created). DefaultValue : -created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetAssetCommandsUsingGET(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetAssetCommandsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetCommandsUsingGET`: PageableCommandV0
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetAssetCommandsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | requested commands target asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | requested commands target asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetCommandsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int64** | page size | [default to 20]
 **page** | **int64** | page | [default to 0]
 **sort** | [**[]string**](string.md) | sorting list by attributes (supported columns:  id, status, created). DefaultValue : -created | 

### Return type

[**PageableCommandV0**](Pageable«CommandV0».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandStatusUsingGET

> SimpleStringWeb GetCommandStatusUsingGET(ctx, commandId).XAPIKEY(xAPIKEY).Execute()

Get the status of a specific command



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
    commandId := "commandId_example" // string | identifier of specific command. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetCommandStatusUsingGET(context.Background(), commandId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetCommandStatusUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandStatusUsingGET`: SimpleStringWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetCommandStatusUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | identifier of specific command. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandStatusUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**SimpleStringWeb**](SimpleStringWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandUsingGET

> CommandV0 GetCommandUsingGET(ctx, commandId).XAPIKEY(xAPIKEY).Execute()

Get a specific command



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
    commandId := "commandId_example" // string | identifier of specific command. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetCommandUsingGET(context.Background(), commandId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetCommandUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandUsingGET`: CommandV0
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.GetCommandUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | identifier of specific command. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**CommandV0**](CommandV0.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommandsUsingGET

> PageableCommandV0 ListCommandsUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Execute()

List registered commands



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
    size := 987 // int64 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int64 | the requested page number (optional) (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | sorting list by attributes (supported columns:  id, status, created). DefaultValue : -created (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.ListCommandsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.ListCommandsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommandsUsingGET`: PageableCommandV0
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.ListCommandsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListCommandsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int64** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int64** | the requested page number (optional) | [default to 0]
 **sort** | [**[]string**](string.md) | sorting list by attributes (supported columns:  id, status, created). DefaultValue : -created | 

### Return type

[**PageableCommandV0**](Pageable«CommandV0».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCommandStatusUsingPUT

> SetCommandStatusUsingPUT(ctx, commandId).XAPIKEY(xAPIKEY).Force(force).NewStatus(newStatus).Execute()

Update the status of specific command 



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
    commandId := "commandId_example" // string | identifier of specific command. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | force the status update (optional) (default to false)
    newStatus := "newStatus_example" // string | future state of the command --> CANCELED (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.SetCommandStatusUsingPUT(context.Background(), commandId).XAPIKEY(xAPIKEY).Force(force).NewStatus(newStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.SetCommandStatusUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | identifier of specific command. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetCommandStatusUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | force the status update | [default to false]
 **newStatus** | **string** | future state of the command --&gt; CANCELED | 

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

