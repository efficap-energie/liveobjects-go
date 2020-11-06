# \DeviceManagementForLoRaApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountDevicesUsingGET2**](DeviceManagementForLoRaApi.md#CountDevicesUsingGET2) | **Get** /api/v0/vendors/lora/devices/count | Device counts
[**GetDeviceProfilesUsingGET**](DeviceManagementForLoRaApi.md#GetDeviceProfilesUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/profiles | Get LoRa device profiles
[**GetDeviceProfilesUsingGET1**](DeviceManagementForLoRaApi.md#GetDeviceProfilesUsingGET1) | **Get** /api/v0/vendors/lora/profiles | Get LoRa device profiles (use /api/v1/deviceMgt/connectors/lora/profiles instead)
[**GetDeviceUsingGET3**](DeviceManagementForLoRaApi.md#GetDeviceUsingGET3) | **Get** /api/v0/vendors/lora/devices/{devEUI} | Get a device (use Device management - Connector nodes - V1 instead)
[**GetGatewayUsingGET**](DeviceManagementForLoRaApi.md#GetGatewayUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/gateways/{id} | Get a gateway
[**GetMessageCountUsingGET**](DeviceManagementForLoRaApi.md#GetMessageCountUsingGET) | **Get** /api/v0/vendors/lora/data/count | List of message counts group by period
[**ListCommandUsingGET**](DeviceManagementForLoRaApi.md#ListCommandUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks | List commands
[**ListCommandUsingGET1**](DeviceManagementForLoRaApi.md#ListCommandUsingGET1) | **Get** /api/v0/vendors/lora/devices/{devEUI}/commands | List commands (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)
[**ListConnectivityPlanUsingGET**](DeviceManagementForLoRaApi.md#ListConnectivityPlanUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/connectivity | List connectivity plan
[**ListConnectivityPlansUsingGET**](DeviceManagementForLoRaApi.md#ListConnectivityPlansUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/connectivities | List connectivity plans
[**ListDevicesUsingGET3**](DeviceManagementForLoRaApi.md#ListDevicesUsingGET3) | **Get** /api/v0/vendors/lora/devices | List LoRa devices (use Device management - Connector nodes - V1 instead)
[**ListGatewaysUsingGET**](DeviceManagementForLoRaApi.md#ListGatewaysUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/gateways | List gateways
[**RegisterCommandUsingPOST**](DeviceManagementForLoRaApi.md#RegisterCommandUsingPOST) | **Post** /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks | Register a command
[**RegisterCommandUsingPOST1**](DeviceManagementForLoRaApi.md#RegisterCommandUsingPOST1) | **Post** /api/v0/vendors/lora/devices/{devEUI}/commands | Register a command (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)
[**RegisterDeviceUsingPOST**](DeviceManagementForLoRaApi.md#RegisterDeviceUsingPOST) | **Post** /api/v0/vendors/lora/devices | Register a LoRa device (use Device management - Interfaces - V1 instead)
[**UnregisterDeviceUsingDELETE**](DeviceManagementForLoRaApi.md#UnregisterDeviceUsingDELETE) | **Delete** /api/v0/vendors/lora/devices/{devEUI} | Unregister a device (use Device management - Connector nodes - V1 instead)
[**UpdateDeviceUsingPATCH1**](DeviceManagementForLoRaApi.md#UpdateDeviceUsingPATCH1) | **Patch** /api/v0/vendors/lora/devices/{devEUI} | Update a device (use Device management - Connector nodes - V1 instead)
[**UpdateGatewayUsingPATCH**](DeviceManagementForLoRaApi.md#UpdateGatewayUsingPATCH) | **Patch** /api/v1/deviceMgt/connectors/lora/gateways/{id} | Update a gateway.



## CountDevicesUsingGET2

> LoraDeviceStatsWeb CountDevicesUsingGET2(ctx).XAPIKEY(xAPIKEY).Execute()

Device counts



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.CountDevicesUsingGET2(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.CountDevicesUsingGET2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountDevicesUsingGET2`: LoraDeviceStatsWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.CountDevicesUsingGET2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCountDevicesUsingGET2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**LoraDeviceStatsWeb**](LoraDeviceStatsWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceProfilesUsingGET

> []string GetDeviceProfilesUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

Get LoRa device profiles



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.GetDeviceProfilesUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.GetDeviceProfilesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfilesUsingGET`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.GetDeviceProfilesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfilesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceProfilesUsingGET1

> []string GetDeviceProfilesUsingGET1(ctx).XAPIKEY(xAPIKEY).Execute()

Get LoRa device profiles (use /api/v1/deviceMgt/connectors/lora/profiles instead)



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.GetDeviceProfilesUsingGET1(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.GetDeviceProfilesUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceProfilesUsingGET1`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.GetDeviceProfilesUsingGET1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceProfilesUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceUsingGET3

> LoraDevice GetDeviceUsingGET3(ctx, devEUI).XAPIKEY(xAPIKEY).Execute()

Get a device (use Device management - Connector nodes - V1 instead)



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
    devEUI := "devEUI_example" // string | DevEUI of the device
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.GetDeviceUsingGET3(context.Background(), devEUI).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.GetDeviceUsingGET3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceUsingGET3`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.GetDeviceUsingGET3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string** | DevEUI of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceUsingGET3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayUsingGET

> LoraGatewayInfo GetGatewayUsingGET(ctx, id).XAPIKEY(xAPIKEY).Execute()

Get a gateway



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
    id := "id_example" // string | identifier of the gateway
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.GetGatewayUsingGET(context.Background(), id).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.GetGatewayUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGatewayUsingGET`: LoraGatewayInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.GetGatewayUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | identifier of the gateway | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGatewayUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**LoraGatewayInfo**](LoraGatewayInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageCountUsingGET

> DeferredListenableFutureResultstringstring GetMessageCountUsingGET(ctx).XAPIKEY(xAPIKEY).Period(period).Timezone(timezone).Execute()

List of message counts group by period



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
    period := "period_example" // string | the period to group by (optional)
    timezone := "timezone_example" // string | the timezone of the client (optional) (default to "UTC")

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.GetMessageCountUsingGET(context.Background()).XAPIKEY(xAPIKEY).Period(period).Timezone(timezone).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.GetMessageCountUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetMessageCountUsingGET`: DeferredListenableFutureResultstringstring
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.GetMessageCountUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetMessageCountUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **period** | **string** | the period to group by | 
 **timezone** | **string** | the timezone of the client | [default to &quot;UTC&quot;]

### Return type

[**DeferredListenableFutureResultstringstring**](DeferredListenableFutureResult«string,string».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommandUsingGET

> PageableLoraCommand ListCommandUsingGET(ctx, devEUI).XAPIKEY(xAPIKEY).Size(size).Page(page).Status(status).TimeRange(timeRange).Sort(sort).Execute()

List commands



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
    devEUI := "devEUI_example" // string | DevEUI of the device
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | The requested page number (optional) (optional) (default to 0)
    status := "status_example" // string | The command status (optional)
    timeRange := []string{"Inner_example"} // []string | Filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange=[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusiv, upperbound is exclusiv. (optional)
    sort := "sort_example" // string | Field used for sorting. Prefix with '-' for descending order (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.ListCommandUsingGET(context.Background(), devEUI).XAPIKEY(xAPIKEY).Size(size).Page(page).Status(status).TimeRange(timeRange).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.ListCommandUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommandUsingGET`: PageableLoraCommand
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.ListCommandUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string** | DevEUI of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommandUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | The requested page number (optional) | [default to 0]
 **status** | **string** | The command status | 
 **timeRange** | [**[]string**](string.md) | Filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange&#x3D;[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusiv, upperbound is exclusiv. | 
 **sort** | **string** | Field used for sorting. Prefix with &#39;-&#39; for descending order | 

### Return type

[**PageableLoraCommand**](Pageable«LoraCommand».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommandUsingGET1

> PageableLoraCommand ListCommandUsingGET1(ctx, devEUI).XAPIKEY(xAPIKEY).Size(size).Page(page).Status(status).TimeRange(timeRange).Sort(sort).Execute()

List commands (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)



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
    devEUI := "devEUI_example" // string | DevEUI of the device
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | The requested page number (optional) (optional) (default to 0)
    status := "status_example" // string | The command status (optional)
    timeRange := []string{"Inner_example"} // []string | Filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange=[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusive, upperbound is exclusive. (optional)
    sort := "sort_example" // string | Field used for sorting. Prefix with '-' for descending order (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.ListCommandUsingGET1(context.Background(), devEUI).XAPIKEY(xAPIKEY).Size(size).Page(page).Status(status).TimeRange(timeRange).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.ListCommandUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListCommandUsingGET1`: PageableLoraCommand
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.ListCommandUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string** | DevEUI of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiListCommandUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | The requested page number (optional) | [default to 0]
 **status** | **string** | The command status | 
 **timeRange** | [**[]string**](string.md) | Filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange&#x3D;[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusive, upperbound is exclusive. | 
 **sort** | **string** | Field used for sorting. Prefix with &#39;-&#39; for descending order | 

### Return type

[**PageableLoraCommand**](Pageable«LoraCommand».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectivityPlanUsingGET

> []string ListConnectivityPlanUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

List connectivity plan



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.ListConnectivityPlanUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.ListConnectivityPlanUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectivityPlanUsingGET`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.ListConnectivityPlanUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectivityPlanUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectivityPlansUsingGET

> []LoraNetworkSubscriptionDetail ListConnectivityPlansUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

List connectivity plans



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.ListConnectivityPlansUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.ListConnectivityPlansUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListConnectivityPlansUsingGET`: []LoraNetworkSubscriptionDetail
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.ListConnectivityPlansUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListConnectivityPlansUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**[]LoraNetworkSubscriptionDetail**](LoraNetworkSubscriptionDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicesUsingGET3

> PageableLoraDevice ListDevicesUsingGET3(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).DevEUI(devEUI).Name(name).Status(status).Tags(tags).Sort(sort).Execute()

List LoRa devices (use Device management - Connector nodes - V1 instead)



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
    page := 987 // int32 | Requested page number (optional) (optional) (default to 0)
    devEUI := "devEUI_example" // string | DevEUI regexp filter (optional)
    name := "name_example" // string | Name regexp filter (optional)
    status := "status_example" // string | Status filter (optional)
    tags := []string{"Inner_example"} // []string | Tags filter (optional)
    sort := "sort_example" // string | Field used for sorting (prefix with '-' for descending order) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.ListDevicesUsingGET3(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).DevEUI(devEUI).Name(name).Status(status).Tags(tags).Sort(sort).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.ListDevicesUsingGET3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevicesUsingGET3`: PageableLoraDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.ListDevicesUsingGET3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesUsingGET3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | Requested page number (optional) | [default to 0]
 **devEUI** | **string** | DevEUI regexp filter | 
 **name** | **string** | Name regexp filter | 
 **status** | **string** | Status filter | 
 **tags** | [**[]string**](string.md) | Tags filter | 
 **sort** | **string** | Field used for sorting (prefix with &#39;-&#39; for descending order) | 

### Return type

[**PageableLoraDevice**](Pageable«LoraDevice».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGatewaysUsingGET

> PageableLoraGatewayInfo ListGatewaysUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Id(id).Name(name).Type_(type_).Status(status).Manufacturer(manufacturer).Execute()

List gateways



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
    id := "id_example" // string | filtering by identifier (regexp filter) of gateway (optional)
    name := "name_example" // string | filtering by name (regexp filter) of gateway (optional)
    type_ := "type__example" // string | filtering by type of gateway (optional)
    status := "status_example" // string | filtering by status of gateway (optional)
    manufacturer := "manufacturer_example" // string | filtering by manufacturer of gateway (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.ListGatewaysUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Id(id).Name(name).Type_(type_).Status(status).Manufacturer(manufacturer).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.ListGatewaysUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGatewaysUsingGET`: PageableLoraGatewayInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.ListGatewaysUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGatewaysUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]
 **id** | **string** | filtering by identifier (regexp filter) of gateway | 
 **name** | **string** | filtering by name (regexp filter) of gateway | 
 **type_** | **string** | filtering by type of gateway | 
 **status** | **string** | filtering by status of gateway | 
 **manufacturer** | **string** | filtering by manufacturer of gateway | 

### Return type

[**PageableLoraGatewayInfo**](Pageable«LoraGatewayInfo».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCommandUsingPOST

> LoraCommand RegisterCommandUsingPOST(ctx, devEUI).XAPIKEY(xAPIKEY).Req(req).Execute()

Register a command



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
    devEUI := "devEUI_example" // string | DevEUI of the device
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    req := *openapiclient.NewLoraCommandWeb(true, "Data_example", 123) // LoraCommandWeb | A LoRa command (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.RegisterCommandUsingPOST(context.Background(), devEUI).XAPIKEY(xAPIKEY).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.RegisterCommandUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterCommandUsingPOST`: LoraCommand
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.RegisterCommandUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string** | DevEUI of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCommandUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **req** | [**LoraCommandWeb**](LoraCommandWeb.md) | A LoRa command | 

### Return type

[**LoraCommand**](LoraCommand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCommandUsingPOST1

> LoraCommand RegisterCommandUsingPOST1(ctx, devEUI).XAPIKEY(xAPIKEY).Req(req).Execute()

Register a command (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)



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
    devEUI := "devEUI_example" // string | DevEUI of the device
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    req := *openapiclient.NewLoraCommandWeb(true, "Data_example", 123) // LoraCommandWeb | A LoRa command (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.RegisterCommandUsingPOST1(context.Background(), devEUI).XAPIKEY(xAPIKEY).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.RegisterCommandUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterCommandUsingPOST1`: LoraCommand
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.RegisterCommandUsingPOST1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string** | DevEUI of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiRegisterCommandUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **req** | [**LoraCommandWeb**](LoraCommandWeb.md) | A LoRa command | 

### Return type

[**LoraCommand**](LoraCommand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDeviceUsingPOST

> LoraDevice RegisterDeviceUsingPOST(ctx).XAPIKEY(xAPIKEY).Req(req).Execute()

Register a LoRa device (use Device management - Interfaces - V1 instead)



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
    req := *openapiclient.NewLoraDeviceCreateReqWeb("ActivationType_example", "AppEUI_example", "AppKey_example", "DevEUI_example", "DeviceStatus_example", "Name_example", "Profile_example") // LoraDeviceCreateReqWeb | A LoRa device (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.RegisterDeviceUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.RegisterDeviceUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RegisterDeviceUsingPOST`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.RegisterDeviceUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiRegisterDeviceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **req** | [**LoraDeviceCreateReqWeb**](LoraDeviceCreateReqWeb.md) | A LoRa device | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterDeviceUsingDELETE

> bool UnregisterDeviceUsingDELETE(ctx, devEUI).XAPIKEY(xAPIKEY).Execute()

Unregister a device (use Device management - Connector nodes - V1 instead)



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
    devEUI := "devEUI_example" // string | DevEUI of the device
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.UnregisterDeviceUsingDELETE(context.Background(), devEUI).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.UnregisterDeviceUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UnregisterDeviceUsingDELETE`: bool
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.UnregisterDeviceUsingDELETE`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string** | DevEUI of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUnregisterDeviceUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceUsingPATCH1

> LoraDevice UpdateDeviceUsingPATCH1(ctx, devEUI).XAPIKEY(xAPIKEY).Req(req).Execute()

Update a device (use Device management - Connector nodes - V1 instead)



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
    devEUI := "devEUI_example" // string | DevEUI of the device
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    req := *openapiclient.NewLoraDeviceUpdateReqWeb() // LoraDeviceUpdateReqWeb | A LoRa device (add only fields you want to update) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.UpdateDeviceUsingPATCH1(context.Background(), devEUI).XAPIKEY(xAPIKEY).Req(req).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.UpdateDeviceUsingPATCH1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceUsingPATCH1`: LoraDevice
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.UpdateDeviceUsingPATCH1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string** | DevEUI of the device | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceUsingPATCH1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **req** | [**LoraDeviceUpdateReqWeb**](LoraDeviceUpdateReqWeb.md) | A LoRa device (add only fields you want to update) | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGatewayUsingPATCH

> LoraGatewayInfo UpdateGatewayUsingPATCH(ctx, id).XAPIKEY(xAPIKEY).LoraGatewayData(loraGatewayData).Execute()

Update a gateway.



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
    id := "id_example" // string | identifier of the gateway
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    loraGatewayData := *openapiclient.NewLoraGatewayData() // LoraGatewayData | loraGatewayData

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementForLoRaApi.UpdateGatewayUsingPATCH(context.Background(), id).XAPIKEY(xAPIKEY).LoraGatewayData(loraGatewayData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementForLoRaApi.UpdateGatewayUsingPATCH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGatewayUsingPATCH`: LoraGatewayInfo
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementForLoRaApi.UpdateGatewayUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | identifier of the gateway | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGatewayUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **loraGatewayData** | [**LoraGatewayData**](LoraGatewayData.md) | loraGatewayData | 

### Return type

[**LoraGatewayInfo**](LoraGatewayInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

