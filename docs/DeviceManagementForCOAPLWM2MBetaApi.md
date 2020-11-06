# \DeviceManagementForCOAPLWM2MBetaApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceUsingGET4**](DeviceManagementForCOAPLWM2MBetaApi.md#GetDeviceUsingGET4) | **Get** /api/v0/vendors/lwm2m/identity/{ep} | Find a device
[**ListDevicesUsingGET4**](DeviceManagementForCOAPLWM2MBetaApi.md#ListDevicesUsingGET4) | **Get** /api/v0/vendors/lwm2m/identities | List lwm2m devices
[**RegisterDeviceUsingPOST1**](DeviceManagementForCOAPLWM2MBetaApi.md#RegisterDeviceUsingPOST1) | **Post** /api/v0/vendors/lwm2m/identities | Register LWM2M/DTLS identity
[**UnregisterDeviceUsingDELETE1**](DeviceManagementForCOAPLWM2MBetaApi.md#UnregisterDeviceUsingDELETE1) | **Delete** /api/v0/vendors/lwm2m/identity/{ep} | Unregister a device
[**UpdateDeviceUsingPUT**](DeviceManagementForCOAPLWM2MBetaApi.md#UpdateDeviceUsingPUT) | **Put** /api/v0/vendors/lwm2m/identity/{ep} | Update a device



## GetDeviceUsingGET4

> LWM2MDevice GetDeviceUsingGET4(ctx, ep).XAPIKEY(xAPIKEY).Execute()

Find a device



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
    ep := "ep_example" // string | the device's endpoint
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForCOAPLWM2MBetaApi.GetDeviceUsingGET4(context.Background(), ep).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForCOAPLWM2MBetaApi.GetDeviceUsingGET4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceUsingGET4`: LWM2MDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForCOAPLWM2MBetaApi.GetDeviceUsingGET4`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ep** | **string** | the device&#39;s endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceUsingGET4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**LWM2MDevice**](LWM2MDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicesUsingGET4

> Lwm2mDevicePageWeb ListDevicesUsingGET4(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Ep(ep).Tags(tags).Execute()

List lwm2m devices



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
    page := 987 // int32 | the requested page number (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | sort (optional)
    ep := "ep_example" // string | filter, regexp on endpoint to filter list (optional)
    tags := []string{"Inner_example"} // []string | filter, required tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForCOAPLWM2MBetaApi.ListDevicesUsingGET4(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Ep(ep).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForCOAPLWM2MBetaApi.ListDevicesUsingGET4``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevicesUsingGET4`: Lwm2mDevicePageWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForCOAPLWM2MBetaApi.ListDevicesUsingGET4`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesUsingGET4Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number | [default to 0]
 **sort** | [**[]string**](string.md) | sort | 
 **ep** | **string** | filter, regexp on endpoint to filter list | 
 **tags** | [**[]string**](string.md) | filter, required tags | 

### Return type

[**Lwm2mDevicePageWeb**](Lwm2mDevicePageWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDeviceUsingPOST1

> LWM2MDevice RegisterDeviceUsingPOST1(ctx).XAPIKEY(xAPIKEY).Device(device).Execute()

Register LWM2M/DTLS identity



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
    device := *openapiclient.NewLWM2MDevice() // LWM2MDevice | device

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForCOAPLWM2MBetaApi.RegisterDeviceUsingPOST1(context.Background()).XAPIKEY(xAPIKEY).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForCOAPLWM2MBetaApi.RegisterDeviceUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDeviceUsingPOST1`: LWM2MDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForCOAPLWM2MBetaApi.RegisterDeviceUsingPOST1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDeviceUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **device** | [**LWM2MDevice**](LWM2MDevice.md) | device | 

### Return type

[**LWM2MDevice**](LWM2MDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterDeviceUsingDELETE1

> UnregisterDeviceUsingDELETE1(ctx, ep).XAPIKEY(xAPIKEY).Execute()

Unregister a device



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
    ep := "ep_example" // string | the device's endpoint
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForCOAPLWM2MBetaApi.UnregisterDeviceUsingDELETE1(context.Background(), ep).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForCOAPLWM2MBetaApi.UnregisterDeviceUsingDELETE1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ep** | **string** | the device&#39;s endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterDeviceUsingDELETE1Request struct via the builder pattern


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


## UpdateDeviceUsingPUT

> LWM2MDevice UpdateDeviceUsingPUT(ctx, ep).XAPIKEY(xAPIKEY).Device(device).Execute()

Update a device



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
    ep := "ep_example" // string | the device's endpoint
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    device := *openapiclient.NewLWM2MDevice() // LWM2MDevice | device

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForCOAPLWM2MBetaApi.UpdateDeviceUsingPUT(context.Background(), ep).XAPIKEY(xAPIKEY).Device(device).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForCOAPLWM2MBetaApi.UpdateDeviceUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceUsingPUT`: LWM2MDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForCOAPLWM2MBetaApi.UpdateDeviceUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ep** | **string** | the device&#39;s endpoint | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **device** | [**LWM2MDevice**](LWM2MDevice.md) | device | 

### Return type

[**LWM2MDevice**](LWM2MDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

