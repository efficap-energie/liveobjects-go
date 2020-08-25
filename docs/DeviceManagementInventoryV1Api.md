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

> Device CreateDeviceUsingPOST(ctx, xAPIKEY, optional)

Create a device

Please refer to the 'Device Management > Interfaces' API notes for more information about 'interfaces.[x].definition' content<br><br>Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateDeviceUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateDeviceUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of DeviceCreateRequest**](DeviceCreateRequest.md)| The device to register | 

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

> DeleteDeviceUsingDELETE(ctx, deviceId, xAPIKEY)

Delete a device

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| Target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> []DeviceStreamsResponseWeb GetDeviceStreamsUsingGET(ctx, deviceId, xAPIKEY, optional)

Get a device's streamIds

Restricted to API keys with at least one of the following roles : DATA_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| Targeted device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetDeviceStreamsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeviceStreamsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| maximum number of return items (optional, max 100 items) | [default to 10]

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

> Device GetDeviceUsingGET2(ctx, deviceId, xAPIKEY)

Get a device

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> []Device ListDevicesUsingGET2(ctx, xAPIKEY, optional)

List registered devices

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListDevicesUsingGET2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDevicesUsingGET2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **offset** | **optional.Int32**| number of items to skip (optional) | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| sorting list by attributes. Supported columns: id, name, created, updated, group, interfaces.status, interfaces.enabled, interfaces.lastContact). Example: [\&quot;urn\&quot;,\&quot;-created\&quot;].  | 
 **id** | **optional.String**| filter list by device identifier (regexp). Expected string (max 269 characters) | 
 **groupPath** | **optional.String**| filter list by device groupPath.  (with optional use of wildcard &#39;/_*&#39; at the end of search term)Expected string (max 255 characters) | 
 **groupId** | **optional.String**| filter list by device groupId. Expected string (max 6 characters) | 
 **name** | **optional.String**| filter list by device name.  (with optional use of wildcard &#39;*&#39; at the beginning or end of search term)Expected string (max 255 characters) | 
 **tags** | [**optional.Interface of []string**](string.md)| filter list by device tags. Max number of tags depends of your offer settings. Tag value max length is 32. | 
 **connectors** | [**optional.Interface of []string**](string.md)| list devices with interfaces of the specified connector(s). precede the connector with \&quot;-\&quot; to exclude it, use \&quot;none\&quot; to list devices with no interfaces. Example: \&quot;mqtt, -lora, none\&quot;. | 
 **interfacesNodeId** | **optional.String**| filter list by nodeId. Expected string (max 269 characters) | 
 **interfacesStatus** | [**optional.Interface of []string**](string.md)| filter list by interface status. Supported values: REGISTERED, INITIALIZING, INITIALIZED, ONLINE, OFFLINE, ACTIVATED, REACTIVATED, DEACTIVATED, CONNECTIVITY_ERROR, DELETED. | 
 **interfacesEnabled** | **optional.Bool**| filter list by interface enabled state. | 
 **propertyFilterName** | **optional.String**| Multiple filters, Example : devices?property.temperature&#x3D;25&amp;property.humidity&#x3D;58...[cannot be tested in swagger-ui]. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | 
 **activityStates** | [**optional.Interface of []string**](string.md)| filter list by activity state. Supported values: ACTIVE, SILENT, UNKNOWN, NOT_MONITORED | 
 **filterQuery** | **optional.String**| device filter expression using RSQL notation.  Supported device properties are &#39;groupPath&#39;, &#39;groupId&#39;, &#39;tags&#39;, &#39;connector&#39;, &#39;properties&#39;. Supported RSQL operators are &#39;&#x3D;&#x3D;&#39;,&#39;!&#x3D;&#39;,&#39;&#x3D;in&#x3D;&#39;,&#39;&#x3D;out&#x3D;&#39;,&#39;&#x3D;re&#x3D;&#39;,&#39;&#x3D;lt&#x3D;&#39;,&#39;&#x3D;le&#x3D;&#39;,&#39;&#x3D;gt&#x3D;&#39;,&#39;&#x3D;ge&#x3D;&#39;,&#39;and&#39;,&#39;or&#39;. Expected string (max 1000 characters) | 
 **fields** | [**optional.Interface of []string**](string.md)| list of fields to return.  Amongst: &#39;id&#39;, &#39;name&#39;, &#39;description&#39;, &#39;tags&#39;, &#39;properties&#39;, &#39;group&#39;, &#39;interfaces&#39;, &#39;config&#39;, &#39;firmwares&#39;, &#39;activityState&#39;, &#39;defaultDataStreamId&#39;, &#39;created&#39;, &#39;updated&#39;), default: id, name, tags &amp; group | 
 **xTotalCount** | **optional.Bool**| true if a total count must be returned in response | [default to false]

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

> ErrorResponseWeb UpdateDeviceUsingPATCH(ctx, deviceId, xAPIKEY, deviceUpdate)

Update a device

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| Targeted device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 
**deviceUpdate** | [**DeviceUpdate**](DeviceUpdate.md)| deviceUpdate | 

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

