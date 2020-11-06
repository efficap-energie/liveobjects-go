# \DeviceManagementResourcesV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceResourceUpdatesUsingGET1**](DeviceManagementResourcesV1Api.md#GetDeviceResourceUpdatesUsingGET1) | **Get** /api/v1/deviceMgt/devices/{deviceId}/resources/updates | Get a list of device resource updates
[**GetDeviceResourceUsingGET**](DeviceManagementResourcesV1Api.md#GetDeviceResourceUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} | Get a specific device resource
[**GetLastResourceUpdateUsingGET1**](DeviceManagementResourcesV1Api.md#GetLastResourceUpdateUsingGET1) | **Get** /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId}/updates/latest | Get info about last update of this device resource
[**ListDeviceResourcesUsingGET1**](DeviceManagementResourcesV1Api.md#ListDeviceResourcesUsingGET1) | **Get** /api/v1/deviceMgt/devices/{deviceId}/resources | Get a map of all device resources
[**SetDeviceResourceVersionUsingPOST1**](DeviceManagementResourcesV1Api.md#SetDeviceResourceVersionUsingPOST1) | **Post** /api/v1/deviceMgt/devices/{deviceId}/resources/{resourceId} | Set device resource versions



## GetDeviceResourceUpdatesUsingGET1

> []ResourceUpdateWeb GetDeviceResourceUpdatesUsingGET1(ctx, deviceId).XAPIKEY(xAPIKEY).Offset(offset).Limit(limit).XTotalCount(xTotalCount).Execute()

Get a list of device resource updates



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
    offset := 987 // int32 | the requested page number (optional) (optional) (default to 0)
    limit := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    xTotalCount := true // bool | true if a total count must be returned in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesV1Api.GetDeviceResourceUpdatesUsingGET1(context.Background(), deviceId).XAPIKEY(xAPIKEY).Offset(offset).Limit(limit).XTotalCount(xTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesV1Api.GetDeviceResourceUpdatesUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceResourceUpdatesUsingGET1`: []ResourceUpdateWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesV1Api.GetDeviceResourceUpdatesUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | requested target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceResourceUpdatesUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **offset** | **int32** | the requested page number (optional) | [default to 0]
 **limit** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **xTotalCount** | **bool** | true if a total count must be returned in response | [default to false]

### Return type

[**[]ResourceUpdateWeb**](ResourceUpdateWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceResourceUsingGET

> DeviceResourceWeb GetDeviceResourceUsingGET(ctx, deviceId, resourceId).XAPIKEY(xAPIKEY).Execute()

Get a specific device resource



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
    resourceId := "resourceId_example" // string | device resource identifier. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesV1Api.GetDeviceResourceUsingGET(context.Background(), deviceId, resourceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesV1Api.GetDeviceResourceUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceResourceUsingGET`: DeviceResourceWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesV1Api.GetDeviceResourceUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**resourceId** | **string** | device resource identifier. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceResourceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**DeviceResourceWeb**](DeviceResourceWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastResourceUpdateUsingGET1

> ResourceUpdateWeb GetLastResourceUpdateUsingGET1(ctx, deviceId, resourceId).XAPIKEY(xAPIKEY).Execute()

Get info about last update of this device resource



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
    resourceId := "resourceId_example" // string | device resource identifier. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesV1Api.GetLastResourceUpdateUsingGET1(context.Background(), deviceId, resourceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesV1Api.GetLastResourceUpdateUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetLastResourceUpdateUsingGET1`: ResourceUpdateWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesV1Api.GetLastResourceUpdateUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | requested device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**resourceId** | **string** | device resource identifier. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetLastResourceUpdateUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

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


## ListDeviceResourcesUsingGET1

> map[string]DeviceResourceWeb ListDeviceResourcesUsingGET1(ctx, deviceId).XAPIKEY(xAPIKEY).Execute()

Get a map of all device resources



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
    resp, r, err := api_client.DeviceManagementResourcesV1Api.ListDeviceResourcesUsingGET1(context.Background(), deviceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesV1Api.ListDeviceResourcesUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDeviceResourcesUsingGET1`: map[string]DeviceResourceWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesV1Api.ListDeviceResourcesUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | requested target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListDeviceResourcesUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**map[string]DeviceResourceWeb**](DeviceResourceWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceResourceVersionUsingPOST1

> DeviceResourceWeb SetDeviceResourceVersionUsingPOST1(ctx, deviceId, resourceId).XAPIKEY(xAPIKEY).Update(update).Execute()

Set device resource versions



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
    resourceId := "resourceId_example" // string | device resource identifier. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    update := *openapiclient.NewUpdateDeviceResourceReqWeb() // UpdateDeviceResourceReqWeb | The device resource version to register (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesV1Api.SetDeviceResourceVersionUsingPOST1(context.Background(), deviceId, resourceId).XAPIKEY(xAPIKEY).Update(update).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesV1Api.SetDeviceResourceVersionUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetDeviceResourceVersionUsingPOST1`: DeviceResourceWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesV1Api.SetDeviceResourceVersionUsingPOST1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**resourceId** | **string** | device resource identifier. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceResourceVersionUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **update** | [**UpdateDeviceResourceReqWeb**](UpdateDeviceResourceReqWeb.md) | The device resource version to register | 

### Return type

[**DeviceResourceWeb**](DeviceResourceWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

