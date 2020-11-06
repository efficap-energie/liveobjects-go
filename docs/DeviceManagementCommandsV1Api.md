# \DeviceManagementCommandsV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommandUsingPOST1**](DeviceManagementCommandsV1Api.md#AddCommandUsingPOST1) | **Post** /api/v1/deviceMgt/devices/{deviceId}/commands | Register a new command
[**DeleteCommandUsingDELETE1**](DeviceManagementCommandsV1Api.md#DeleteCommandUsingDELETE1) | **Delete** /api/v1/deviceMgt/commands/{commandId} | Delete a specific command
[**GetAssetCommandsUsingGET1**](DeviceManagementCommandsV1Api.md#GetAssetCommandsUsingGET1) | **Get** /api/v1/deviceMgt/devices/{deviceId}/commands | List commands targeting a specific device
[**GetCommandStatusUsingGET1**](DeviceManagementCommandsV1Api.md#GetCommandStatusUsingGET1) | **Get** /api/v1/deviceMgt/commands/{commandId}/status | Get the status of a specific command
[**GetCommandUsingGET1**](DeviceManagementCommandsV1Api.md#GetCommandUsingGET1) | **Get** /api/v1/deviceMgt/commands/{commandId} | Get a specific command
[**SetCommandStatusUsingPUT1**](DeviceManagementCommandsV1Api.md#SetCommandStatusUsingPUT1) | **Put** /api/v1/deviceMgt/commands/{commandId}/status | Update the status of specific command 



## AddCommandUsingPOST1

> Command AddCommandUsingPOST1(ctx, deviceId).XAPIKEY(xAPIKEY).Command(command).Validate(validate).Execute()

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
    deviceId := "deviceId_example" // string | target device identifier (URN). A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    command := *openapiclient.NewCommandAddRequest(*openapiclient.NewCommandRequest("Connector_example", 123)) // CommandAddRequest | new command request
    validate := true // bool | Command will be validated by connector before registration. Default is \"true\" (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementCommandsV1Api.AddCommandUsingPOST1(context.Background(), deviceId).XAPIKEY(xAPIKEY).Command(command).Validate(validate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementCommandsV1Api.AddCommandUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddCommandUsingPOST1`: Command
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementCommandsV1Api.AddCommandUsingPOST1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | target device identifier (URN). A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddCommandUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **command** | [**CommandAddRequest**](CommandAddRequest.md) | new command request | 
 **validate** | **bool** | Command will be validated by connector before registration. Default is \&quot;true\&quot; | [default to true]

### Return type

[**Command**](Command.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommandUsingDELETE1

> DeleteCommandUsingDELETE1(ctx, commandId).XAPIKEY(xAPIKEY).Execute()

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
    resp, r, err := api_client.DeviceManagementCommandsV1Api.DeleteCommandUsingDELETE1(context.Background(), commandId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementCommandsV1Api.DeleteCommandUsingDELETE1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteCommandUsingDELETE1Request struct via the builder pattern


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


## GetAssetCommandsUsingGET1

> []Command GetAssetCommandsUsingGET1(ctx, deviceId).XAPIKEY(xAPIKEY).From(from).To(to).Limit(limit).Offset(offset).Sort(sort).XTotalCount(xTotalCount).Execute()

List commands targeting a specific device



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
    deviceId := "deviceId_example" // string | requested commands target device identifier (URN). A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    from := "from_example" // string | Search for commands created after this date. Use ISO-8601 normalization. (optional)
    to := "to_example" // string | Search for commands created before this date. Use ISO-8601 normalization. (optional)
    limit := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    offset := 987 // int32 | number of items to skip (optional) (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | sorting list by attributes. DefaultValue : -created.  Supported columns: id, status, created. Example: [\"status\",\"-created\"].  (optional)
    xTotalCount := true // bool | true if a total count must be returned in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementCommandsV1Api.GetAssetCommandsUsingGET1(context.Background(), deviceId).XAPIKEY(xAPIKEY).From(from).To(to).Limit(limit).Offset(offset).Sort(sort).XTotalCount(xTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementCommandsV1Api.GetAssetCommandsUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetCommandsUsingGET1`: []Command
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementCommandsV1Api.GetAssetCommandsUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | requested commands target device identifier (URN). A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetCommandsUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **from** | **string** | Search for commands created after this date. Use ISO-8601 normalization. | 
 **to** | **string** | Search for commands created before this date. Use ISO-8601 normalization. | 
 **limit** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **offset** | **int32** | number of items to skip (optional) | [default to 0]
 **sort** | [**[]string**](string.md) | sorting list by attributes. DefaultValue : -created.  Supported columns: id, status, created. Example: [\&quot;status\&quot;,\&quot;-created\&quot;].  | 
 **xTotalCount** | **bool** | true if a total count must be returned in response | [default to false]

### Return type

[**[]Command**](Command.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandStatusUsingGET1

> SimpleStringWeb GetCommandStatusUsingGET1(ctx, commandId).XAPIKEY(xAPIKEY).Execute()

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
    resp, r, err := api_client.DeviceManagementCommandsV1Api.GetCommandStatusUsingGET1(context.Background(), commandId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementCommandsV1Api.GetCommandStatusUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandStatusUsingGET1`: SimpleStringWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementCommandsV1Api.GetCommandStatusUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | identifier of specific command. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandStatusUsingGET1Request struct via the builder pattern


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


## GetCommandUsingGET1

> Command GetCommandUsingGET1(ctx, commandId).XAPIKEY(xAPIKEY).Execute()

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
    resp, r, err := api_client.DeviceManagementCommandsV1Api.GetCommandUsingGET1(context.Background(), commandId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementCommandsV1Api.GetCommandUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCommandUsingGET1`: Command
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementCommandsV1Api.GetCommandUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string** | identifier of specific command. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetCommandUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**Command**](Command.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCommandStatusUsingPUT1

> SetCommandStatusUsingPUT1(ctx, commandId).XAPIKEY(xAPIKEY).NewStatus(newStatus).Force(force).Execute()

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
    newStatus := "newStatus_example" // string | future state of the command  --> CANCELED
    force := true // bool | force the update of the command status (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementCommandsV1Api.SetCommandStatusUsingPUT1(context.Background(), commandId).XAPIKEY(xAPIKEY).NewStatus(newStatus).Force(force).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementCommandsV1Api.SetCommandStatusUsingPUT1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSetCommandStatusUsingPUT1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **newStatus** | **string** | future state of the command  --&gt; CANCELED | 
 **force** | **bool** | force the update of the command status | [default to false]

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

