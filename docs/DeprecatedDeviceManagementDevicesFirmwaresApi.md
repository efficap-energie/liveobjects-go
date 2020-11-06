# \DeprecatedDeviceManagementDevicesFirmwaresApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceFirmwareUsingGET**](DeprecatedDeviceManagementDevicesFirmwaresApi.md#GetDeviceFirmwareUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId} | Get a specific device firmware (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} instead)
[**GetDeviceResourceUpdatesUsingGET**](DeprecatedDeviceManagementDevicesFirmwaresApi.md#GetDeviceResourceUpdatesUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwareUpdates | Get a list of device firmware updates (use /api/v1/deviceMgt/devices/{deviceId}/resources/updates
[**GetLastResourceUpdateUsingGET**](DeprecatedDeviceManagementDevicesFirmwaresApi.md#GetLastResourceUpdateUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId}/lastUpdate | Get info about last update of this device firmware (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId}/updates/latest instead)
[**ListDeviceResourcesUsingGET**](DeprecatedDeviceManagementDevicesFirmwaresApi.md#ListDeviceResourcesUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwares | Get a map of all device firmwares (use /api/v1/deviceMgt/devices/{deviceId}/resources instead)
[**SetDeviceResourceVersionUsingPOST**](DeprecatedDeviceManagementDevicesFirmwaresApi.md#SetDeviceResourceVersionUsingPOST) | **Post** /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId} | Set device firmware versions (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} instead)



## GetDeviceFirmwareUsingGET

> DeviceFirmwareWeb GetDeviceFirmwareUsingGET(ctx, deviceId, firmwareId).XAPIKEY(xAPIKEY).Execute()

Get a specific device firmware (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} instead)



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
    deviceId := "deviceId_example" // string | device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    firmwareId := "firmwareId_example" // string | device firmware identifier. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementDevicesFirmwaresApi.GetDeviceFirmwareUsingGET(context.Background(), deviceId, firmwareId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementDevicesFirmwaresApi.GetDeviceFirmwareUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceFirmwareUsingGET`: DeviceFirmwareWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementDevicesFirmwaresApi.GetDeviceFirmwareUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**firmwareId** | **string** | device firmware identifier. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceFirmwareUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**DeviceFirmwareWeb**](DeviceFirmwareWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceResourceUpdatesUsingGET

> []FirmwareUpdateWeb GetDeviceResourceUpdatesUsingGET(ctx, deviceId).XAPIKEY(xAPIKEY).Page(page).Size(size).XTotalCount(xTotalCount).Execute()

Get a list of device firmware updates (use /api/v1/deviceMgt/devices/{deviceId}/resources/updates



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
    deviceId := "deviceId_example" // string | requested target device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    xTotalCount := true // bool | true if a total count must be returned in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementDevicesFirmwaresApi.GetDeviceResourceUpdatesUsingGET(context.Background(), deviceId).XAPIKEY(xAPIKEY).Page(page).Size(size).XTotalCount(xTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementDevicesFirmwaresApi.GetDeviceResourceUpdatesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceResourceUpdatesUsingGET`: []FirmwareUpdateWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementDevicesFirmwaresApi.GetDeviceResourceUpdatesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | requested target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceResourceUpdatesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **page** | **int32** | the requested page number (optional) | [default to 0]
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **xTotalCount** | **bool** | true if a total count must be returned in response | [default to false]

### Return type

[**[]FirmwareUpdateWeb**](FirmwareUpdateWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastResourceUpdateUsingGET

> FirmwareUpdateWeb GetLastResourceUpdateUsingGET(ctx, deviceId, firmwareId).XAPIKEY(xAPIKEY).Execute()

Get info about last update of this device firmware (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId}/updates/latest instead)



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
    deviceId := "deviceId_example" // string | requested device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    firmwareId := "firmwareId_example" // string | device firmware identifier. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementDevicesFirmwaresApi.GetLastResourceUpdateUsingGET(context.Background(), deviceId, firmwareId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementDevicesFirmwaresApi.GetLastResourceUpdateUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastResourceUpdateUsingGET`: FirmwareUpdateWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementDevicesFirmwaresApi.GetLastResourceUpdateUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | requested device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**firmwareId** | **string** | device firmware identifier. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastResourceUpdateUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**FirmwareUpdateWeb**](FirmwareUpdateWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceResourcesUsingGET

> map[string]DeviceFirmwareWeb ListDeviceResourcesUsingGET(ctx, deviceId).XAPIKEY(xAPIKEY).Execute()

Get a map of all device firmwares (use /api/v1/deviceMgt/devices/{deviceId}/resources instead)



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
    deviceId := "deviceId_example" // string | requested target device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementDevicesFirmwaresApi.ListDeviceResourcesUsingGET(context.Background(), deviceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementDevicesFirmwaresApi.ListDeviceResourcesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceResourcesUsingGET`: map[string]DeviceFirmwareWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementDevicesFirmwaresApi.ListDeviceResourcesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | requested target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceResourcesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**map[string]DeviceFirmwareWeb**](DeviceFirmwareWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceResourceVersionUsingPOST

> DeviceFirmwareWeb SetDeviceResourceVersionUsingPOST(ctx, deviceId, firmwareId).XAPIKEY(xAPIKEY).Update(update).Execute()

Set device firmware versions (use /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} instead)



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
    deviceId := "deviceId_example" // string | device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    firmwareId := "firmwareId_example" // string | device firmware identifier. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    update := *openapiclient.NewUpdateDeviceFirmwareReqWeb() // UpdateDeviceFirmwareReqWeb | The device firmware version to register (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementDevicesFirmwaresApi.SetDeviceResourceVersionUsingPOST(context.Background(), deviceId, firmwareId).XAPIKEY(xAPIKEY).Update(update).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementDevicesFirmwaresApi.SetDeviceResourceVersionUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDeviceResourceVersionUsingPOST`: DeviceFirmwareWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementDevicesFirmwaresApi.SetDeviceResourceVersionUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**firmwareId** | **string** | device firmware identifier. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceResourceVersionUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **update** | [**UpdateDeviceFirmwareReqWeb**](UpdateDeviceFirmwareReqWeb.md) | The device firmware version to register | 

### Return type

[**DeviceFirmwareWeb**](DeviceFirmwareWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

