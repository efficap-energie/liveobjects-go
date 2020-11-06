# \DeviceManagementInterfacesV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInterfaceToDeviceUsingPOST**](DeviceManagementInterfacesV1Api.md#AddInterfaceToDeviceUsingPOST) | **Post** /api/v1/deviceMgt/devices/{deviceId}/interfaces | Add an interface to a registered device
[**DeleteInterfaceUsingDELETE**](DeviceManagementInterfacesV1Api.md#DeleteInterfaceUsingDELETE) | **Delete** /api/v1/deviceMgt/devices/{deviceId}/interfaces/{interfaceId} | Delete an interface
[**GetInterfaceForADeviceUsingGET**](DeviceManagementInterfacesV1Api.md#GetInterfaceForADeviceUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/interfaces/{interfaceId} | Get a specific interface for a registered device
[**ListInterfacesForADeviceUsingGET**](DeviceManagementInterfacesV1Api.md#ListInterfacesForADeviceUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/interfaces | Get the interface list for a registered device
[**UpdateInterfaceUsingPATCH**](DeviceManagementInterfacesV1Api.md#UpdateInterfaceUsingPATCH) | **Patch** /api/v1/deviceMgt/devices/{deviceId}/interfaces/{interfaceId} | Update an interface



## AddInterfaceToDeviceUsingPOST

> DeviceInterface AddInterfaceToDeviceUsingPOST(ctx, deviceId).XAPIKEY(xAPIKEY).NewDeviceInterface(newDeviceInterface).Execute()

Add an interface to a registered device



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
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    newDeviceInterface := *openapiclient.NewNewDeviceInterface() // NewDeviceInterface | The device interface to add (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInterfacesV1Api.AddInterfaceToDeviceUsingPOST(context.Background(), deviceId).XAPIKEY(xAPIKEY).NewDeviceInterface(newDeviceInterface).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInterfacesV1Api.AddInterfaceToDeviceUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddInterfaceToDeviceUsingPOST`: DeviceInterface
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInterfacesV1Api.AddInterfaceToDeviceUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddInterfaceToDeviceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **newDeviceInterface** | [**NewDeviceInterface**](NewDeviceInterface.md) | The device interface to add | 

### Return type

[**DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterfaceUsingDELETE

> DeleteInterfaceUsingDELETE(ctx, deviceId, interfaceId).XAPIKEY(xAPIKEY).Execute()

Delete an interface



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
    interfaceId := "interfaceId_example" // string | Must be {connector}:{nodeId}.
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInterfacesV1Api.DeleteInterfaceUsingDELETE(context.Background(), deviceId, interfaceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInterfacesV1Api.DeleteInterfaceUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**interfaceId** | **string** | Must be {connector}:{nodeId}. | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteInterfaceUsingDELETERequest struct via the builder pattern


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


## GetInterfaceForADeviceUsingGET

> DeviceInterface GetInterfaceForADeviceUsingGET(ctx, deviceId, interfaceId).XAPIKEY(xAPIKEY).Execute()

Get a specific interface for a registered device



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
    interfaceId := "interfaceId_example" // string | Must be {connector}:{nodeId}
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInterfacesV1Api.GetInterfaceForADeviceUsingGET(context.Background(), deviceId, interfaceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInterfacesV1Api.GetInterfaceForADeviceUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetInterfaceForADeviceUsingGET`: DeviceInterface
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInterfacesV1Api.GetInterfaceForADeviceUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**interfaceId** | **string** | Must be {connector}:{nodeId} | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInterfaceForADeviceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInterfacesForADeviceUsingGET

> []DeviceInterface ListInterfacesForADeviceUsingGET(ctx, deviceId).XAPIKEY(xAPIKEY).Execute()

Get the interface list for a registered device



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
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInterfacesV1Api.ListInterfacesForADeviceUsingGET(context.Background(), deviceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInterfacesV1Api.ListInterfacesForADeviceUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListInterfacesForADeviceUsingGET`: []DeviceInterface
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInterfacesV1Api.ListInterfacesForADeviceUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiListInterfacesForADeviceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**[]DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterfaceUsingPATCH

> DeviceInterface UpdateInterfaceUsingPATCH(ctx, deviceId, interfaceId).XAPIKEY(xAPIKEY).Body(body).Execute()

Update an interface



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
    interfaceId := "interfaceId_example" // string | Must be {connector}:{nodeId}
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    body := *openapiclient.NewUpdateInterfaceReqWeb() // UpdateInterfaceReqWeb | The fields to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInterfacesV1Api.UpdateInterfaceUsingPATCH(context.Background(), deviceId, interfaceId).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInterfacesV1Api.UpdateInterfaceUsingPATCH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateInterfaceUsingPATCH`: DeviceInterface
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInterfacesV1Api.UpdateInterfaceUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**interfaceId** | **string** | Must be {connector}:{nodeId} | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateInterfaceUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **body** | [**UpdateInterfaceReqWeb**](UpdateInterfaceReqWeb.md) | The fields to update | 

### Return type

[**DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

