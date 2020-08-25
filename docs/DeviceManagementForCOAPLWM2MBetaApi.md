# \DeviceManagementForCOAPLWM2MBetaApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDeviceUsingGET4**](DeviceManagementForCOAPLWM2MBetaApi.md#GetDeviceUsingGET4) | **Get** /api/v0/vendors/lwm2m/identity/{ep} | Find a device
[**ListDevicesUsingGET4**](DeviceManagementForCOAPLWM2MBetaApi.md#ListDevicesUsingGET4) | **Get** /api/v0/vendors/lwm2m/identities | List lwm2m devices
[**RegisterDeviceUsingPOST1**](DeviceManagementForCOAPLWM2MBetaApi.md#RegisterDeviceUsingPOST1) | **Post** /api/v0/vendors/lwm2m/identities | Register LWM2M/DTLS identity
[**UnregisterDeviceUsingDELETE1**](DeviceManagementForCOAPLWM2MBetaApi.md#UnregisterDeviceUsingDELETE1) | **Delete** /api/v0/vendors/lwm2m/identity/{ep} | Unregister a device
[**UpdateDeviceUsingPUT**](DeviceManagementForCOAPLWM2MBetaApi.md#UpdateDeviceUsingPUT) | **Put** /api/v0/vendors/lwm2m/identity/{ep} | Update a device



## GetDeviceUsingGET4

> Lwm2MDevice GetDeviceUsingGET4(ctx, ep, xAPIKEY)

Find a device

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ep** | **string**| the device&#39;s endpoint | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**Lwm2MDevice**](LWM2MDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicesUsingGET4

> Lwm2mDevicePageWeb ListDevicesUsingGET4(ctx, xAPIKEY, optional)

List lwm2m devices

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListDevicesUsingGET4Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDevicesUsingGET4Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number | [default to 0]
 **sort** | [**optional.Interface of []string**](string.md)| sort | 
 **ep** | **optional.String**| filter, regexp on endpoint to filter list | 
 **tags** | [**optional.Interface of []string**](string.md)| filter, required tags | 

### Return type

[**Lwm2mDevicePageWeb**](Lwm2mDevicePageWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDeviceUsingPOST1

> Lwm2MDevice RegisterDeviceUsingPOST1(ctx, xAPIKEY, device)

Register LWM2M/DTLS identity

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
**device** | [**Lwm2MDevice**](Lwm2MDevice.md)| device | 

### Return type

[**Lwm2MDevice**](LWM2MDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterDeviceUsingDELETE1

> UnregisterDeviceUsingDELETE1(ctx, ep, xAPIKEY)

Unregister a device

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ep** | **string**| the device&#39;s endpoint | 
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


## UpdateDeviceUsingPUT

> Lwm2MDevice UpdateDeviceUsingPUT(ctx, ep, xAPIKEY, device)

Update a device

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**ep** | **string**| the device&#39;s endpoint | 
**xAPIKEY** | **string**| a valid API key | 
**device** | [**Lwm2MDevice**](Lwm2MDevice.md)| device | 

### Return type

[**Lwm2MDevice**](LWM2MDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

