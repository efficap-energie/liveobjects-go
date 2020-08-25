# \DeviceManagementCommandsV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommandUsingPOST1**](DeviceManagementCommandsV1Api.md#AddCommandUsingPOST1) | **Post** /api/v1/deviceMgt/devices/{deviceId}/commands | Register a new command
[**DeleteCommandUsingDELETE**](DeviceManagementCommandsV1Api.md#DeleteCommandUsingDELETE) | **Delete** /api/v1/deviceMgt/commands/{commandId} | Delete a specific command
[**GetAssetCommandsUsingGET1**](DeviceManagementCommandsV1Api.md#GetAssetCommandsUsingGET1) | **Get** /api/v1/deviceMgt/devices/{deviceId}/commands | List commands targeting a specific device
[**GetCommandStatusUsingGET**](DeviceManagementCommandsV1Api.md#GetCommandStatusUsingGET) | **Get** /api/v1/deviceMgt/commands/{commandId}/status | Get the status of a specific command
[**GetCommandUsingGET**](DeviceManagementCommandsV1Api.md#GetCommandUsingGET) | **Get** /api/v1/deviceMgt/commands/{commandId} | Get a specific command
[**SetCommandStatusUsingPUT**](DeviceManagementCommandsV1Api.md#SetCommandStatusUsingPUT) | **Put** /api/v1/deviceMgt/commands/{commandId}/status | Update the status of specific command 



## AddCommandUsingPOST1

> Command AddCommandUsingPOST1(ctx, deviceId, xAPIKEY, command, optional)

Register a new command

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| target device identifier (URN). A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 
**command** | [**CommandAddRequest**](CommandAddRequest.md)| new command request | 
 **optional** | ***AddCommandUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCommandUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **validate** | **optional.Bool**| Command will be validated by connector before registration. Default is \&quot;true\&quot; | [default to true]

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


## DeleteCommandUsingDELETE

> DeleteCommandUsingDELETE(ctx, commandId, xAPIKEY)

Delete a specific command

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string**| identifier of specific command. Expected identifier (max 24 characters) | 
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


## GetAssetCommandsUsingGET1

> []Command GetAssetCommandsUsingGET1(ctx, deviceId, xAPIKEY, optional)

List commands targeting a specific device

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| requested commands target device identifier (URN). A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetAssetCommandsUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAssetCommandsUsingGET1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **from** | **optional.String**| Search for commands created after this date. Use ISO-8601 normalization. | 
 **to** | **optional.String**| Search for commands created before this date. Use ISO-8601 normalization. | 
 **limit** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **offset** | **optional.Int32**| number of items to skip (optional) | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| sorting list by attributes. DefaultValue : -created.  Supported columns: id, status, created. Example: [\&quot;status\&quot;,\&quot;-created\&quot;].  | 
 **xTotalCount** | **optional.Bool**| true if a total count must be returned in response | [default to false]

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


## GetCommandStatusUsingGET

> SimpleStringWeb GetCommandStatusUsingGET(ctx, commandId, xAPIKEY)

Get the status of a specific command

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string**| identifier of specific command. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> Command GetCommandUsingGET(ctx, commandId, xAPIKEY)

Get a specific command

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string**| identifier of specific command. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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


## SetCommandStatusUsingPUT

> SetCommandStatusUsingPUT(ctx, commandId, xAPIKEY, newStatus, optional)

Update the status of specific command 

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string**| identifier of specific command. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
**newStatus** | **string**| future state of the command  --&gt; CANCELED | 
 **optional** | ***SetCommandStatusUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetCommandStatusUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **optional.Bool**| force the update of the command status | [default to false]

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

