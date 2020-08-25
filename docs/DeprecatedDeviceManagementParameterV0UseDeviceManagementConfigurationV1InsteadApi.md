# \DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetParamUsingGET**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#GetAssetParamUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId}/params/{paramKey} | Get a specific asset parameter
[**GetAssetParamsUsingGET**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#GetAssetParamsUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId}/params | Get a specific asset list of parameters
[**SetAssetParamsUsingPOST**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#SetAssetParamsUsingPOST) | **Post** /api/v0/assets/{assetNamespace}/{assetId}/params | Update a specific asset list of parameters
[**SetDeviceParamUpdateStatusUsingPUT**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#SetDeviceParamUpdateStatusUsingPUT) | **Put** /api/v0/assets/{assetNamespace}/{assetId}/params/{paramKey}/status | Update the status of a specific asset parameter update
[**SetDeviceParamsUpdateStatusUsingPUT**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#SetDeviceParamsUpdateStatusUsingPUT) | **Put** /api/v0/assets/{assetNamespace}/{assetId}/params/status | Update the status of a specific asset parameters update



## GetAssetParamUsingGET

> AssetParameter GetAssetParamUsingGET(ctx, assetNamespace, assetId, paramKey, xAPIKEY)

Get a specific asset parameter

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**paramKey** | **string**| key identifying the targeted asset parameter. Expected string (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**AssetParameter**](AssetParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetParamsUsingGET

> map[string]AssetParameter GetAssetParamsUsingGET(ctx, assetNamespace, assetId, xAPIKEY)

Get a specific asset list of parameters

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**map[string]AssetParameter**](AssetParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAssetParamsUsingPOST

> SetAssetParamsUsingPOST(ctx, assetNamespace, assetId, xAPIKEY, optional)

Update a specific asset list of parameters

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetAssetParamsUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetAssetParamsUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **notifyTo** | **optional.String**| fifo used to relay status notification ex: fifo/~notif. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;, &#39;router/&#39;] and max length is 255 | 
 **newParamValues** | [**optional.Interface of map[string]AssetParameterValue**](AssetParameterValue.md)| new param value ( type ex :  INT32, UINT32 , STRING, FLOAT, BINARY ). Max number of parameters is 100. Parameter name max length is 128. Parameter value must be valid according to the type and max length is 2000. | 

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


## SetDeviceParamUpdateStatusUsingPUT

> SetDeviceParamUpdateStatusUsingPUT(ctx, assetNamespace, assetId, paramKey, xAPIKEY, optional)

Update the status of a specific asset parameter update

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**paramKey** | **string**| key identifying the targeted asset parameter. Expected string (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetDeviceParamUpdateStatusUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetDeviceParamUpdateStatusUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **force** | **optional.Bool**| force the status update | [default to false]
 **newStatus** | **optional.String**| future state of the parameter --&gt; CANCELED | 

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


## SetDeviceParamsUpdateStatusUsingPUT

> SetDeviceParamsUpdateStatusUsingPUT(ctx, assetNamespace, assetId, xAPIKEY, optional)

Update the status of a specific asset parameters update

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string**| asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string**| asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetDeviceParamsUpdateStatusUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetDeviceParamsUpdateStatusUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **optional.Bool**| force the status update | [default to false]
 **reqWeb** | [**optional.Interface of AssetParamsStatusUpdateReqWeb**](AssetParamsStatusUpdateReqWeb.md)| parameters keys and their future status. Maximum 100 parameters, and parameter name max length is 128 | 

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

