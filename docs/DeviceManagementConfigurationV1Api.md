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

> DeviceParameterWeb GetDeviceConfigParameterUsingGET(ctx, deviceId, paramKey, xAPIKEY)

Get state of a specific device configuration parameter

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**paramKey** | **string**| config parameter identifier/key. Expected string (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> map[string]DeviceParameterWeb GetDeviceConfigParametersUsingGET(ctx, deviceId, xAPIKEY)

Get a description of the device configuration parameters

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> DeviceConfigWeb GetDeviceConfigurationUsingGET(ctx, deviceId, xAPIKEY)

Get a description of the device configuration

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> SetDeviceParamUpdateStatusUsingPUT1(ctx, deviceId, paramKey, xAPIKEY, optional)

Update the status of a specific device parameter

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**paramKey** | **string**| key identifying the targeted device parameter. Expected string (max 128 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetDeviceParamUpdateStatusUsingPUT1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetDeviceParamUpdateStatusUsingPUT1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **force** | **optional.Bool**| force the update of the parameter status | [default to false]
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


## SetMultipleDeviceConfigParamsUsingPOST

> SetMultipleDeviceConfigParamsUsingPOST(ctx, deviceId, xAPIKEY, optional)

Set requested values for a set of device configuration parameters

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetMultipleDeviceConfigParamsUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetMultipleDeviceConfigParamsUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **req** | [**optional.Interface of DeviceParametersSetRequest**](DeviceParametersSetRequest.md)| a map of requested configuration parameter values (with type: &#39;INT32&#39;, &#39;UINT32&#39;, &#39;STRING&#39;, &#39;FLOAT&#39;, or &#39;BINARY&#39;) | 

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

