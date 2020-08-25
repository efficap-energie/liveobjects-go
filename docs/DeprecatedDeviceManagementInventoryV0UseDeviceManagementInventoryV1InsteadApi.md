# \DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssetUsingPOST**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#CreateAssetUsingPOST) | **Post** /api/v0/assets | Create a device
[**DeleteDeviceStatusUsingDELETE**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#DeleteDeviceStatusUsingDELETE) | **Delete** /api/v0/assets/{assetNamespace}/{assetId} | Delete a device
[**GetAssetStatusUsingGET**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#GetAssetStatusUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId} | Get a device status
[**ListAssetNamespacesUsingGET**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#ListAssetNamespacesUsingGET) | **Get** /api/v0/inventory/namespaces | Enumerates the used asset namespaces
[**ListAssetsUsingGET**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#ListAssetsUsingGET) | **Get** /api/v0/assets | List registered assets status
[**UpdateAssetUsingPUT**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#UpdateAssetUsingPUT) | **Put** /api/v0/assets/{assetNamespace}/{assetId} | Update a device



## CreateAssetUsingPOST

> Asset CreateAssetUsingPOST(ctx, xAPIKEY, optional)

Create a device

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateAssetUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateAssetUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**optional.Interface of AssetCreateReqWeb**](AssetCreateReqWeb.md)| The device to register | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceStatusUsingDELETE

> DeleteDeviceStatusUsingDELETE(ctx, assetNamespace, assetId, xAPIKEY)

Delete a device

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| targeted for deletion asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| targeted for deletion asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
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


## GetAssetStatusUsingGET

> Asset GetAssetStatusUsingGET(ctx, assetNamespace, assetId, xAPIKEY)

Get a device status

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| requested asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| requested asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetNamespacesUsingGET

> []string ListAssetNamespacesUsingGET(ctx, xAPIKEY)

Enumerates the used asset namespaces

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


## ListAssetsUsingGET

> PageableAsset ListAssetsUsingGET(ctx, xAPIKEY, optional)

List registered assets status

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListAssetsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListAssetsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int64**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int64**| the requested page number (optional) | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| sorting list by attributes.  Supported columns: namespace, id, creationTs, lastUpdateTs, connected, groupPath. | 
 **namespace** | **optional.String**| filter list by asset namespace. Expected string (max 128 characters) | 
 **groupPath** | **optional.String**| filter list by asset groupPath.  (with optional use of wildcard &#39;/_*&#39; at the end of search term)Expected string (max 255 characters) | 
 **groupId** | **optional.String**| filter list by asset groupId. Expected string (max 6 characters) | 
 **id** | **optional.String**| filter list by asset id. (with optional use of wildcard &#39;*&#39; at the beginning or end of search term) Expected string (max 128 characters) | 
 **connected** | **optional.Bool**| filter list by asset connected state | 
 **name** | **optional.String**| filter list by asset name.  (with optional use of wildcard &#39;*&#39; at the beginning or end of search term)Expected string (max 255 characters) | 
 **tags** | [**optional.Interface of []string**](string.md)| filter list by asset tags. Max number of tags depends of your offer settings. Tag value max length is 32. | 
 **propertyFilterName** | **optional.String**| Multiple filters, Example : assets?property.temperature&#x3D;25&amp;property.humidity&#x3D;58...[cannot be tested in swagger-ui]. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | 

### Return type

[**PageableAsset**](Pageable«Asset».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetUsingPUT

> Asset UpdateAssetUsingPUT(ctx, assetNamespace, assetId, xAPIKEY, optional)

Update a device

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| namespace targeted to update assets. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| asset identifier to update. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateAssetUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateAssetUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of AssetUpdateReqWeb**](AssetUpdateReqWeb.md)| the updated asset definition | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

