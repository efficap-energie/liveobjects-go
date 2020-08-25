# \DeviceManagementDevicesFirmwaresResourcesApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourceUsingPOST**](DeviceManagementDevicesFirmwaresResourcesApi.md#AddResourceUsingPOST) | **Post** /api/v0/rm | Add a resource in Resource Manager
[**AddResourceVersionUsingPOST**](DeviceManagementDevicesFirmwaresResourcesApi.md#AddResourceVersionUsingPOST) | **Post** /api/v0/rm/{resourceId}/version | Add a new version to a resource
[**CancelResourceUpdateUsingPOST**](DeviceManagementDevicesFirmwaresResourcesApi.md#CancelResourceUpdateUsingPOST) | **Post** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/cancelUpdate | Cancel asset resource update
[**DeleteResourceUsingDELETE**](DeviceManagementDevicesFirmwaresResourcesApi.md#DeleteResourceUsingDELETE) | **Delete** /api/v0/rm/{resourceId} | Delete a resource
[**DeleteResourceVersionUsingDELETE**](DeviceManagementDevicesFirmwaresResourcesApi.md#DeleteResourceVersionUsingDELETE) | **Delete** /api/v0/rm/{resourceId}/version/{resourceVersionId} | Delete a resource&#39;s version
[**GetAllAssetResourcesUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetAllAssetResourcesUsingGET) | **Get** /api/v0/rm/asset/{assetIdNamespace}/{assetId} | List the asset&#39;s resources (use /api/v1/deviceMgt/devices/{deviceId}/firmwares instead)
[**GetAllConnectorsUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetAllConnectorsUsingGET) | **Get** /api/v0/rm/connectors | Get all available connectors
[**GetConnectorMandatoryAndOptionalMetadataUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetConnectorMandatoryAndOptionalMetadataUsingGET) | **Get** /api/v0/rm/connectors/{connector}/metadata | Get mandatory and optional metadata of a connector
[**GetDeviceFirmwareUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetDeviceFirmwareUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId} | Get a specific device firmware
[**GetDeviceResourceUpdatesUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetDeviceResourceUpdatesUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwareUpdates | Get a list of device firmware updates
[**GetLastResourceUpdateUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetLastResourceUpdateUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId}/lastUpdate | Get info about last update of this device firmware
[**GetLastUpdateUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetLastUpdateUsingGET) | **Get** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/update | Get the asset&#39;s resources update status (use /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId}/lastUpdate instead)
[**GetResourceCompatibleVersionsUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetResourceCompatibleVersionsUsingGET) | **Get** /api/v0/rm/{resourceId}/compatibleVersion/{currentVersion} | List the versions from which a resource update to the version can be done
[**GetResourceUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetResourceUsingGET) | **Get** /api/v0/rm/{resourceId} | Get a resource
[**GetResourceVersionUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetResourceVersionUsingGET) | **Get** /api/v0/rm/{resourceId}/version/{resourceVersionId} | Get a resource&#39;s version
[**GetResourceVersionsUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetResourceVersionsUsingGET) | **Get** /api/v0/rm/{resourceId}/version | List resource&#39;s versions
[**GetResourcesUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetResourcesUsingGET) | **Get** /api/v0/rm | List all resources
[**GetUpdateHistoryUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#GetUpdateHistoryUsingGET) | **Get** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/update/history | Get the asset&#39;s resources update history (use /api/v1/deviceMgt/devices/{deviceId}/firmwareUpdates instead)
[**ListDeviceResourcesUsingGET**](DeviceManagementDevicesFirmwaresResourcesApi.md#ListDeviceResourcesUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/firmwares | Get a map of all device firmwares
[**SetAssetTargetResourceVersionUsingPUT**](DeviceManagementDevicesFirmwaresResourcesApi.md#SetAssetTargetResourceVersionUsingPUT) | **Put** /api/v0/rm/asset/{assetIdNamespace}/{assetId}/resource/{resourceId}/targetversion | Set asset&#39;s target resource version (use /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId} instead)
[**SetDeviceResourceVersionUsingPOST**](DeviceManagementDevicesFirmwaresResourcesApi.md#SetDeviceResourceVersionUsingPOST) | **Post** /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId} | Set device firmware versions
[**UpdateResourceUsingPUT**](DeviceManagementDevicesFirmwaresResourcesApi.md#UpdateResourceUsingPUT) | **Put** /api/v0/rm/{resourceId} | Update a resource
[**UpdateResourceVersionUsingPUT**](DeviceManagementDevicesFirmwaresResourcesApi.md#UpdateResourceVersionUsingPUT) | **Put** /api/v0/rm/{resourceId}/version/{resourceVersionId} | Update a resource&#39;s version



## AddResourceUsingPOST

> Resource AddResourceUsingPOST(ctx, xAPIKEY, optional)

Add a resource in Resource Manager

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***AddResourceUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddResourceUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**optional.Interface of ResourceAddReqWeb**](ResourceAddReqWeb.md)| body of add resource | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddResourceVersionUsingPOST

> ResourceVersion AddResourceVersionUsingPOST(ctx, resourceId, xAPIKEY, optional)

Add a new version to a resource

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***AddResourceVersionUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddResourceVersionUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**optional.Interface of ResourceVersionAddReqWeb**](ResourceVersionAddReqWeb.md)| body of add resource version | 

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CancelResourceUpdateUsingPOST

> CancelResourceUpdateUsingPOST(ctx, assetIdNamespace, assetId, xAPIKEY, optional)

Cancel asset resource update

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string**| the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CancelResourceUpdateUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelResourceUpdateUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **optional.Bool**| force update status change | [default to false]

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


## DeleteResourceUsingDELETE

> DeleteResourceUsingDELETE(ctx, resourceId, xAPIKEY)

Delete a resource

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
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


## DeleteResourceVersionUsingDELETE

> DeleteResourceVersionUsingDELETE(ctx, resourceId, resourceVersionId, xAPIKEY)

Delete a resource's version

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**resourceVersionId** | **string**| the resource version id. Expected string (max 255 characters) | 
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


## GetAllAssetResourcesUsingGET

> PageableAssetResourceWeb GetAllAssetResourcesUsingGET(ctx, assetIdNamespace, assetId, xAPIKEY, optional)

List the asset's resources (use /api/v1/deviceMgt/devices/{deviceId}/firmwares instead)

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string**| the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetAllAssetResourcesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetAllAssetResourcesUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**PageableAssetResourceWeb**](Pageable«AssetResourceWeb».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAllConnectorsUsingGET

> []string GetAllConnectorsUsingGET(ctx, xAPIKEY)

Get all available connectors

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

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


## GetConnectorMandatoryAndOptionalMetadataUsingGET

> MetadataResourceConnectors GetConnectorMandatoryAndOptionalMetadataUsingGET(ctx, connector, xAPIKEY)

Get mandatory and optional metadata of a connector

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string**| the resource&#39;s connector (ex. \&quot;http-updater\&quot;). A connector must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**MetadataResourceConnectors**](MetadataResourceConnectors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceFirmwareUsingGET

> DeviceFirmwareWeb GetDeviceFirmwareUsingGET(ctx, deviceId, firmwareId, xAPIKEY)

Get a specific device firmware

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**firmwareId** | **string**| device firmware identifier. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**DeviceFirmwareWeb**](DeviceFirmwareWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceResourceUpdatesUsingGET

> []FirmwareUpdateWeb GetDeviceResourceUpdatesUsingGET(ctx, deviceId, xAPIKEY, optional)

Get a list of device firmware updates

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| requested target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetDeviceResourceUpdatesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDeviceResourceUpdatesUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]
 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **xTotalCount** | **optional.Bool**| true if a total count must be returned in response | [default to false]

### Return type

[**[]FirmwareUpdateWeb**](FirmwareUpdateWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastResourceUpdateUsingGET

> FirmwareUpdateWeb GetLastResourceUpdateUsingGET(ctx, deviceId, firmwareId, xAPIKEY)

Get info about last update of this device firmware

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| requested device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**firmwareId** | **string**| device firmware identifier. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**FirmwareUpdateWeb**](FirmwareUpdateWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetLastUpdateUsingGET

> ResourceUpdateWeb GetLastUpdateUsingGET(ctx, assetIdNamespace, assetId, xAPIKEY, optional)

Get the asset's resources update status (use /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId}/lastUpdate instead)

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string**| the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetLastUpdateUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetLastUpdateUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceId** | **optional.String**| the resource id. Expected string (max 255 characters) | 

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


## GetResourceCompatibleVersionsUsingGET

> []string GetResourceCompatibleVersionsUsingGET(ctx, resourceId, currentVersion, xAPIKEY)

List the versions from which a resource update to the version can be done

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**currentVersion** | **string**| current version. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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


## GetResourceUsingGET

> Resource GetResourceUsingGET(ctx, resourceId, xAPIKEY)

Get a resource

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceVersionUsingGET

> ResourceVersion GetResourceVersionUsingGET(ctx, resourceId, resourceVersionId, xAPIKEY)

Get a resource's version

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**resourceVersionId** | **string**| the resource version id | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceVersionsUsingGET

> PageableResourceVersion GetResourceVersionsUsingGET(ctx, resourceId, xAPIKEY, optional)

List resource's versions

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetResourceVersionsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetResourceVersionsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**PageableResourceVersion**](Pageable«ResourceVersion».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcesUsingGET

> PageableResource GetResourcesUsingGET(ctx, xAPIKEY, optional)

List all resources

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetResourcesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetResourcesUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**PageableResource**](Pageable«Resource».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUpdateHistoryUsingGET

> PageableResourceUpdateWeb GetUpdateHistoryUsingGET(ctx, assetIdNamespace, assetId, xAPIKEY, optional)

Get the asset's resources update history (use /api/v1/deviceMgt/devices/{deviceId}/firmwareUpdates instead)

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string**| the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetUpdateHistoryUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUpdateHistoryUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**PageableResourceUpdateWeb**](Pageable«ResourceUpdateWeb».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDeviceResourcesUsingGET

> map[string]DeviceFirmwareWeb ListDeviceResourcesUsingGET(ctx, deviceId, xAPIKEY)

Get a map of all device firmwares

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| requested target device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**map[string]DeviceFirmwareWeb**](DeviceFirmwareWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAssetTargetResourceVersionUsingPUT

> AssetResourceWeb SetAssetTargetResourceVersionUsingPUT(ctx, assetIdNamespace, assetId, resourceId, xAPIKEY, optional)

Set asset's target resource version (use /api/v1/deviceMgt/devices/{deviceId}/firmwares/{firmwareId} instead)

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetIdNamespace** | **string**| the asset id namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| the asset id. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetAssetTargetResourceVersionUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetAssetTargetResourceVersionUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **notifyTo** | **optional.String**| fifo used to relay status notification ex: fifo/~notif. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;,&#39;router/&#39;]and max length is 255 | 
 **targetVersion** | [**optional.Interface of AssetSetTargetResourceVersionReqWeb**](AssetSetTargetResourceVersionReqWeb.md)| the target resource version. Expected string (max 255 characters) | 

### Return type

[**AssetResourceWeb**](AssetResourceWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceResourceVersionUsingPOST

> DeviceFirmwareWeb SetDeviceResourceVersionUsingPOST(ctx, deviceId, firmwareId, xAPIKEY, optional)

Set device firmware versions

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**firmwareId** | **string**| device firmware identifier. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetDeviceResourceVersionUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetDeviceResourceVersionUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **update** | [**optional.Interface of UpdateDeviceFirmwareReqWeb**](UpdateDeviceFirmwareReqWeb.md)| The device firmware version to register | 

### Return type

[**DeviceFirmwareWeb**](DeviceFirmwareWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceUsingPUT

> Resource UpdateResourceUsingPUT(ctx, resourceId, xAPIKEY, optional)

Update a resource

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource Id. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateResourceUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateResourceUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **resourceUpdate** | [**optional.Interface of ResourceUpdateReqWeb**](ResourceUpdateReqWeb.md)| body of resource update | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceVersionUsingPUT

> ResourceVersion UpdateResourceVersionUsingPUT(ctx, resourceId, resourceVersionId, xAPIKEY, optional)

Update a resource's version

Usage of this API will be reported in your access log under 'device_resource' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string**| the resource id. Expected string (max 255 characters) | 
**resourceVersionId** | **string**| the resource version id. Expected string (max 255 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateResourceVersionUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateResourceVersionUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **resourceVersionUpdate** | [**optional.Interface of ResourceVersionUpdateReqWeb**](ResourceVersionUpdateReqWeb.md)| body of resource version update | 

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

