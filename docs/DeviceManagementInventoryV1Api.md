# \DeviceManagementInventoryV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDeviceUsingPOST**](DeviceManagementInventoryV1Api.md#CreateDeviceUsingPOST) | **Post** /api/v1/deviceMgt/devices | Create a device
[**DeleteDeviceUsingDELETE**](DeviceManagementInventoryV1Api.md#DeleteDeviceUsingDELETE) | **Delete** /api/v1/deviceMgt/devices/{deviceId} | Delete a device
[**GetDeviceStreamsUsingGET**](DeviceManagementInventoryV1Api.md#GetDeviceStreamsUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/data/streams | Get a device&#39;s streamIds
[**GetDeviceUsingGET2**](DeviceManagementInventoryV1Api.md#GetDeviceUsingGET2) | **Get** /api/v1/deviceMgt/devices/{deviceId} | Get a device
[**ListDevicesUsingGET2**](DeviceManagementInventoryV1Api.md#ListDevicesUsingGET2) | **Get** /api/v1/deviceMgt/devices | List registered devices
[**UpdateDeviceUsingPATCH**](DeviceManagementInventoryV1Api.md#UpdateDeviceUsingPATCH) | **Patch** /api/v1/deviceMgt/devices/{deviceId} | Update a device



## CreateDeviceUsingPOST

> Device CreateDeviceUsingPOST(ctx).XAPIKEY(xAPIKEY).Body(body).Execute()

Create a device



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
    body := *openapiclient.NewDeviceCreateRequest("Id_example") // DeviceCreateRequest | The device to register (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInventoryV1Api.CreateDeviceUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInventoryV1Api.CreateDeviceUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateDeviceUsingPOST`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInventoryV1Api.CreateDeviceUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateDeviceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **body** | [**DeviceCreateRequest**](DeviceCreateRequest.md) | The device to register | 

### Return type

[**Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceUsingDELETE

> DeleteDeviceUsingDELETE(ctx, deviceId).XAPIKEY(xAPIKEY).Execute()

Delete a device



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
    deviceId := "deviceId_example" // string | Target device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInventoryV1Api.DeleteDeviceUsingDELETE(context.Background(), deviceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInventoryV1Api.DeleteDeviceUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceUsingDELETERequest struct via the builder pattern


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


## GetDeviceStreamsUsingGET

> []DeviceStreamsResponseWeb GetDeviceStreamsUsingGET(ctx, deviceId).XAPIKEY(xAPIKEY).Limit(limit).Execute()

Get a device's streamIds



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
    deviceId := "deviceId_example" // string | Targeted device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    limit := 987 // int32 | maximum number of return items (optional, max 100 items) (optional) (default to 10)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInventoryV1Api.GetDeviceStreamsUsingGET(context.Background(), deviceId).XAPIKEY(xAPIKEY).Limit(limit).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInventoryV1Api.GetDeviceStreamsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceStreamsUsingGET`: []DeviceStreamsResponseWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInventoryV1Api.GetDeviceStreamsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Targeted device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceStreamsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **limit** | **int32** | maximum number of return items (optional, max 100 items) | [default to 10]

### Return type

[**[]DeviceStreamsResponseWeb**](DeviceStreamsResponseWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceUsingGET2

> Device GetDeviceUsingGET2(ctx, deviceId).XAPIKEY(xAPIKEY).Execute()

Get a device



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
    resp, r, err := api_client.DeviceManagementInventoryV1Api.GetDeviceUsingGET2(context.Background(), deviceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInventoryV1Api.GetDeviceUsingGET2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetDeviceUsingGET2`: Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInventoryV1Api.GetDeviceUsingGET2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceUsingGET2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicesUsingGET2

> []Device ListDevicesUsingGET2(ctx).XAPIKEY(xAPIKEY).Limit(limit).Offset(offset).Sort(sort).Id(id).GroupPath(groupPath).GroupId(groupId).Name(name).Tags(tags).Connectors(connectors).InterfacesNodeId(interfacesNodeId).InterfacesStatus(interfacesStatus).InterfacesEnabled(interfacesEnabled).PropertyFilterName(propertyFilterName).ActivityStates(activityStates).FilterQuery(filterQuery).Fields(fields).XTotalCount(xTotalCount).Execute()

List registered devices



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
    limit := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    offset := 987 // int32 | number of items to skip (optional) (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | sorting list by attributes. Supported columns: id, name, created, updated, group, interfaces.status, interfaces.enabled, interfaces.lastContact). Example: [\"urn\",\"-created\"].  (optional)
    id := "id_example" // string | filter list by device identifier (regexp). Expected string (max 269 characters) (optional)
    groupPath := "groupPath_example" // string | filter list by device groupPath.  (with optional use of wildcard '/_*' at the end of search term)Expected string (max 255 characters) (optional)
    groupId := "groupId_example" // string | filter list by device groupId. Expected string (max 6 characters) (optional)
    name := "name_example" // string | filter list by device name.  (with optional use of wildcard '*' at the beginning or end of search term)Expected string (max 255 characters) (optional)
    tags := []string{"Inner_example"} // []string | filter list by device tags. Max number of tags depends of your offer settings. Tag value max length is 32. (optional)
    connectors := []string{"Inner_example"} // []string | list devices with interfaces of the specified connector(s). precede the connector with \"-\" to exclude it, use \"none\" to list devices with no interfaces. Example: \"mqtt, -lora, none\". (optional)
    interfacesNodeId := "interfacesNodeId_example" // string | filter list by nodeId. Expected string (max 269 characters) (optional)
    interfacesStatus := []string{"Inner_example"} // []string | filter list by interface status. Supported values: REGISTERED, INITIALIZING, INITIALIZED, ONLINE, OFFLINE, ACTIVATED, REACTIVATED, DEACTIVATED, CONNECTIVITY_ERROR, DELETED. (optional)
    interfacesEnabled := true // bool | filter list by interface enabled state. (optional)
    propertyFilterName := "propertyFilterName_example" // string | Multiple filters, Example : devices?property.temperature=25&property.humidity=58...[cannot be tested in swagger-ui]. Max number of properties depends of your offer settings. A property name must not include following characters <code>$.</code> and max length is 128. Invalid property names are : 'class', '_class'. Property value max length is 256. (optional)
    activityStates := []string{"Inner_example"} // []string | filter list by activity state. Supported values: ACTIVE, SILENT, UNKNOWN, NOT_MONITORED (optional)
    filterQuery := "filterQuery_example" // string | device filter expression using RSQL notation.  Supported device properties are 'groupPath', 'groupId', 'tags', 'connector', 'properties'. Supported RSQL operators are '==','!=','=in=','=out=','=re=','=lt=','=le=','=gt=','=ge=','and','or'. Expected string (max 1000 characters) (optional)
    fields := []string{"Inner_example"} // []string | list of fields to return.  Amongst: 'id', 'name', 'description', 'tags', 'properties', 'group', 'interfaces', 'config', 'resources', 'activityState', 'defaultDataStreamId', 'created', 'updated'), default: id, name, tags & group (optional)
    xTotalCount := true // bool | true if a total count must be returned in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInventoryV1Api.ListDevicesUsingGET2(context.Background()).XAPIKEY(xAPIKEY).Limit(limit).Offset(offset).Sort(sort).Id(id).GroupPath(groupPath).GroupId(groupId).Name(name).Tags(tags).Connectors(connectors).InterfacesNodeId(interfacesNodeId).InterfacesStatus(interfacesStatus).InterfacesEnabled(interfacesEnabled).PropertyFilterName(propertyFilterName).ActivityStates(activityStates).FilterQuery(filterQuery).Fields(fields).XTotalCount(xTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInventoryV1Api.ListDevicesUsingGET2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListDevicesUsingGET2`: []Device
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInventoryV1Api.ListDevicesUsingGET2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListDevicesUsingGET2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **limit** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **offset** | **int32** | number of items to skip (optional) | [default to 0]
 **sort** | [**[]string**](string.md) | sorting list by attributes. Supported columns: id, name, created, updated, group, interfaces.status, interfaces.enabled, interfaces.lastContact). Example: [\&quot;urn\&quot;,\&quot;-created\&quot;].  | 
 **id** | **string** | filter list by device identifier (regexp). Expected string (max 269 characters) | 
 **groupPath** | **string** | filter list by device groupPath.  (with optional use of wildcard &#39;/_*&#39; at the end of search term)Expected string (max 255 characters) | 
 **groupId** | **string** | filter list by device groupId. Expected string (max 6 characters) | 
 **name** | **string** | filter list by device name.  (with optional use of wildcard &#39;*&#39; at the beginning or end of search term)Expected string (max 255 characters) | 
 **tags** | [**[]string**](string.md) | filter list by device tags. Max number of tags depends of your offer settings. Tag value max length is 32. | 
 **connectors** | [**[]string**](string.md) | list devices with interfaces of the specified connector(s). precede the connector with \&quot;-\&quot; to exclude it, use \&quot;none\&quot; to list devices with no interfaces. Example: \&quot;mqtt, -lora, none\&quot;. | 
 **interfacesNodeId** | **string** | filter list by nodeId. Expected string (max 269 characters) | 
 **interfacesStatus** | [**[]string**](string.md) | filter list by interface status. Supported values: REGISTERED, INITIALIZING, INITIALIZED, ONLINE, OFFLINE, ACTIVATED, REACTIVATED, DEACTIVATED, CONNECTIVITY_ERROR, DELETED. | 
 **interfacesEnabled** | **bool** | filter list by interface enabled state. | 
 **propertyFilterName** | **string** | Multiple filters, Example : devices?property.temperature&#x3D;25&amp;property.humidity&#x3D;58...[cannot be tested in swagger-ui]. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | 
 **activityStates** | [**[]string**](string.md) | filter list by activity state. Supported values: ACTIVE, SILENT, UNKNOWN, NOT_MONITORED | 
 **filterQuery** | **string** | device filter expression using RSQL notation.  Supported device properties are &#39;groupPath&#39;, &#39;groupId&#39;, &#39;tags&#39;, &#39;connector&#39;, &#39;properties&#39;. Supported RSQL operators are &#39;&#x3D;&#x3D;&#39;,&#39;!&#x3D;&#39;,&#39;&#x3D;in&#x3D;&#39;,&#39;&#x3D;out&#x3D;&#39;,&#39;&#x3D;re&#x3D;&#39;,&#39;&#x3D;lt&#x3D;&#39;,&#39;&#x3D;le&#x3D;&#39;,&#39;&#x3D;gt&#x3D;&#39;,&#39;&#x3D;ge&#x3D;&#39;,&#39;and&#39;,&#39;or&#39;. Expected string (max 1000 characters) | 
 **fields** | [**[]string**](string.md) | list of fields to return.  Amongst: &#39;id&#39;, &#39;name&#39;, &#39;description&#39;, &#39;tags&#39;, &#39;properties&#39;, &#39;group&#39;, &#39;interfaces&#39;, &#39;config&#39;, &#39;resources&#39;, &#39;activityState&#39;, &#39;defaultDataStreamId&#39;, &#39;created&#39;, &#39;updated&#39;), default: id, name, tags &amp; group | 
 **xTotalCount** | **bool** | true if a total count must be returned in response | [default to false]

### Return type

[**[]Device**](Device.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceUsingPATCH

> ErrorResponseWeb UpdateDeviceUsingPATCH(ctx, deviceId).XAPIKEY(xAPIKEY).DeviceUpdate(deviceUpdate).Execute()

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
    deviceId := "deviceId_example" // string | Targeted device identifier. A Live Objects URN 'urn:lo:nsid:{namespace}:{id}' must respect the following regular expression <code>^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$</code> (max 269 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    deviceUpdate := *openapiclient.NewDeviceUpdate("Id_example") // DeviceUpdate | deviceUpdate

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementInventoryV1Api.UpdateDeviceUsingPATCH(context.Background(), deviceId).XAPIKEY(xAPIKEY).DeviceUpdate(deviceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementInventoryV1Api.UpdateDeviceUsingPATCH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateDeviceUsingPATCH`: ErrorResponseWeb
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementInventoryV1Api.UpdateDeviceUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Targeted device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDeviceUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **deviceUpdate** | [**DeviceUpdate**](DeviceUpdate.md) | deviceUpdate | 

### Return type

[**ErrorResponseWeb**](ErrorResponseWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

