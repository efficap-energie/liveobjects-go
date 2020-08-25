# \DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddCommandUsingPOST**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#AddCommandUsingPOST) | **Post** /api/v0/assets/{assetNamespace}/{assetId}/commands | Register a new command
[**DeleteCommandUsingDELETE1**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#DeleteCommandUsingDELETE1) | **Delete** /api/v0/commands/{commandId} | Delete a specific command
[**GetAssetCommandsUsingGET**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#GetAssetCommandsUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId}/commands | Get a list of commands targeting a specific asset
[**GetCommandStatusUsingGET1**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#GetCommandStatusUsingGET1) | **Get** /api/v0/commands/{commandId}/status | Get the status of a specific command
[**GetCommandUsingGET1**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#GetCommandUsingGET1) | **Get** /api/v0/commands/{commandId} | Get a specific command
[**ListCommandsUsingGET**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#ListCommandsUsingGET) | **Get** /api/v0/commands | List registered commands
[**SetCommandStatusUsingPUT1**](DeprecatedDeviceManagementCommandV0UseDeviceManagementCommandsV1InsteadApi.md#SetCommandStatusUsingPUT1) | **Put** /api/v0/commands/{commandId}/status | Update the status of specific command 



## AddCommandUsingPOST

> CommandV0 AddCommandUsingPOST(ctx, assetNamespace, assetId, xAPIKEY, optional)

Register a new command

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| target asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| target asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***AddCommandUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddCommandUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **notifyTo** | **optional.String**| fifo used to relay status notification ex: fifo/~notif. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;, &#39;router/&#39;] and max length is 255 | 
 **command** | [**optional.Interface of AssetCommandWeb**](AssetCommandWeb.md)| command to send to the device | 

### Return type

[**CommandV0**](CommandV0.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCommandUsingDELETE1

> DeleteCommandUsingDELETE1(ctx, commandId, xAPIKEY)

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


## GetAssetCommandsUsingGET

> PageableCommandV0 GetAssetCommandsUsingGET(ctx, assetNamespace, assetId, xAPIKEY, optional)

Get a list of commands targeting a specific asset

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| requested commands target asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| requested commands target asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetAssetCommandsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAssetCommandsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **size** | **optional.Int64**| page size | [default to 20]
 **page** | **optional.Int64**| page | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| sorting list by attributes (supported columns:  id, status, created). DefaultValue : -created | 

### Return type

[**PageableCommandV0**](Pageable«CommandV0».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCommandStatusUsingGET1

> SimpleStringWeb GetCommandStatusUsingGET1(ctx, commandId, xAPIKEY)

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


## GetCommandUsingGET1

> CommandV0 GetCommandUsingGET1(ctx, commandId, xAPIKEY)

Get a specific command

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string**| identifier of specific command. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**CommandV0**](CommandV0.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommandsUsingGET

> PageableCommandV0 ListCommandsUsingGET(ctx, xAPIKEY, optional)

List registered commands

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListCommandsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCommandsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int64**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int64**| the requested page number (optional) | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| sorting list by attributes (supported columns:  id, status, created). DefaultValue : -created | 

### Return type

[**PageableCommandV0**](Pageable«CommandV0».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetCommandStatusUsingPUT1

> SetCommandStatusUsingPUT1(ctx, commandId, xAPIKEY, optional)

Update the status of specific command 

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**commandId** | **string**| identifier of specific command. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetCommandStatusUsingPUT1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetCommandStatusUsingPUT1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| force the status update | [default to false]
 **newStatus** | **optional.String**| future state of the command --&gt; CANCELED | 

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

