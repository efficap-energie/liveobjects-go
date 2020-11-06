# \DeviceManagementConfigurationV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceConfigParameterUsingGET**](DeviceManagementConfigurationV1Api.md#GetDeviceConfigParameterUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/config/parameters/{paramKey} | Get state of a specific device configuration parameter
[**GetDeviceConfigParametersUsingGET**](DeviceManagementConfigurationV1Api.md#GetDeviceConfigParametersUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/config/parameters | Get a description of the device configuration parameters
[**GetDeviceConfigurationUsingGET**](DeviceManagementConfigurationV1Api.md#GetDeviceConfigurationUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/config | Get a description of the device configuration
[**SetDeviceParamUpdateStatusUsingPUT1**](DeviceManagementConfigurationV1Api.md#SetDeviceParamUpdateStatusUsingPUT1) | **Put** /api/v1/deviceMgt/devices/{deviceId}/config/parameters/{paramKey}/status | Update the status of a specific device parameter
[**SetMultipleDeviceConfigParamsUsingPOST**](DeviceManagementConfigurationV1Api.md#SetMultipleDeviceConfigParamsUsingPOST) | **Post** /api/v1/deviceMgt/devices/{deviceId}/config | Set requested values for a set of device configuration parameters



## GetDeviceConfigParameterUsingGET

> DeviceParameterWeb GetDeviceConfigParameterUsingGET(ctx, deviceId, paramKey).XAPIKEY(xAPIKEY).Execute()

Get state of a specific device configuration parameter



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
    paramKey := "paramKey_example" // string | config parameter identifier/key. Expected string (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementConfigurationV1Api.GetDeviceConfigParameterUsingGET(context.Background(), deviceId, paramKey).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConfigurationV1Api.GetDeviceConfigParameterUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceConfigParameterUsingGET`: DeviceParameterWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementConfigurationV1Api.GetDeviceConfigParameterUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**paramKey** | **string** | config parameter identifier/key. Expected string (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigParameterUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**DeviceParameterWeb**](DeviceParameterWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceConfigParametersUsingGET

> map[string]DeviceParameterWeb GetDeviceConfigParametersUsingGET(ctx, deviceId).XAPIKEY(xAPIKEY).Execute()

Get a description of the device configuration parameters



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
    resp, r, err := api_client.DeviceManagementConfigurationV1Api.GetDeviceConfigParametersUsingGET(context.Background(), deviceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConfigurationV1Api.GetDeviceConfigParametersUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceConfigParametersUsingGET`: map[string]DeviceParameterWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementConfigurationV1Api.GetDeviceConfigParametersUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigParametersUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**map[string]DeviceParameterWeb**](DeviceParameterWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceConfigurationUsingGET

> DeviceConfigWeb GetDeviceConfigurationUsingGET(ctx, deviceId).XAPIKEY(xAPIKEY).Execute()

Get a description of the device configuration



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
    resp, r, err := api_client.DeviceManagementConfigurationV1Api.GetDeviceConfigurationUsingGET(context.Background(), deviceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConfigurationV1Api.GetDeviceConfigurationUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceConfigurationUsingGET`: DeviceConfigWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementConfigurationV1Api.GetDeviceConfigurationUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceConfigurationUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**DeviceConfigWeb**](DeviceConfigWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceParamUpdateStatusUsingPUT1

> SetDeviceParamUpdateStatusUsingPUT1(ctx, deviceId, paramKey).XAPIKEY(xAPIKEY).Force(force).NewStatus(newStatus).Execute()

Update the status of a specific device parameter



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
    paramKey := "paramKey_example" // string | key identifying the targeted device parameter. Expected string (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | force the update of the parameter status (optional) (default to false)
    newStatus := "newStatus_example" // string | future state of the parameter --> CANCELED (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementConfigurationV1Api.SetDeviceParamUpdateStatusUsingPUT1(context.Background(), deviceId, paramKey).XAPIKEY(xAPIKEY).Force(force).NewStatus(newStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConfigurationV1Api.SetDeviceParamUpdateStatusUsingPUT1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**paramKey** | **string** | key identifying the targeted device parameter. Expected string (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceParamUpdateStatusUsingPUT1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | force the update of the parameter status | [default to false]
 **newStatus** | **string** | future state of the parameter --&gt; CANCELED | 

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


## SetMultipleDeviceConfigParamsUsingPOST

> SetMultipleDeviceConfigParamsUsingPOST(ctx, deviceId).XAPIKEY(xAPIKEY).Req(req).Execute()

Set requested values for a set of device configuration parameters



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
    req := *openapiclient.NewDeviceParametersSetRequest() // DeviceParametersSetRequest | a map of requested configuration parameter values (with type: 'INT32', 'UINT32', 'STRING', 'FLOAT', or 'BINARY') (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementConfigurationV1Api.SetMultipleDeviceConfigParamsUsingPOST(context.Background(), deviceId).XAPIKEY(xAPIKEY).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConfigurationV1Api.SetMultipleDeviceConfigParamsUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetMultipleDeviceConfigParamsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **req** | [**DeviceParametersSetRequest**](DeviceParametersSetRequest.md) | a map of requested configuration parameter values (with type: &#39;INT32&#39;, &#39;UINT32&#39;, &#39;STRING&#39;, &#39;FLOAT&#39;, or &#39;BINARY&#39;) | 

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

