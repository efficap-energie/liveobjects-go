# \DeviceManagementForLPWAApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CountDevicesUsingGET2**](DeviceManagementForLPWAApi.md#CountDevicesUsingGET2) | **Get** /api/v0/vendors/lora/devices/count | Device counts
[**GetDeviceProfilesUsingGET**](DeviceManagementForLPWAApi.md#GetDeviceProfilesUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/profiles | Get LoRa device profiles
[**GetDeviceProfilesUsingGET1**](DeviceManagementForLPWAApi.md#GetDeviceProfilesUsingGET1) | **Get** /api/v0/vendors/lora/profiles | Get LoRa device profiles (use /api/v1/deviceMgt/connectors/lora/profiles instead)
[**GetDeviceUsingGET3**](DeviceManagementForLPWAApi.md#GetDeviceUsingGET3) | **Get** /api/v0/vendors/lora/devices/{devEUI} | Get a device (use Device management - Connector nodes - V1 instead)
[**GetGatewayUsingGET**](DeviceManagementForLPWAApi.md#GetGatewayUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/gateways/{id} | Get a gateway
[**GetMessageCountUsingGET**](DeviceManagementForLPWAApi.md#GetMessageCountUsingGET) | **Get** /api/v0/vendors/lora/data/count | List of message counts group by period
[**ListCommandUsingGET**](DeviceManagementForLPWAApi.md#ListCommandUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks | List commands
[**ListCommandUsingGET1**](DeviceManagementForLPWAApi.md#ListCommandUsingGET1) | **Get** /api/v0/vendors/lora/devices/{devEUI}/commands | List commands (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)
[**ListConnectivityPlanUsingGET**](DeviceManagementForLPWAApi.md#ListConnectivityPlanUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/connectivity | List connectivity plan
[**ListConnectivityPlansUsingGET**](DeviceManagementForLPWAApi.md#ListConnectivityPlansUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/connectivities | List connectivity plans
[**ListDevicesUsingGET3**](DeviceManagementForLPWAApi.md#ListDevicesUsingGET3) | **Get** /api/v0/vendors/lora/devices | List LoRa devices (use Device management - Connector nodes - V1 instead)
[**ListGatewaysUsingGET**](DeviceManagementForLPWAApi.md#ListGatewaysUsingGET) | **Get** /api/v1/deviceMgt/connectors/lora/gateways | List gateways
[**RegisterCommandUsingPOST**](DeviceManagementForLPWAApi.md#RegisterCommandUsingPOST) | **Post** /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks | Register a command
[**RegisterCommandUsingPOST1**](DeviceManagementForLPWAApi.md#RegisterCommandUsingPOST1) | **Post** /api/v0/vendors/lora/devices/{devEUI}/commands | Register a command (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)
[**RegisterDeviceUsingPOST**](DeviceManagementForLPWAApi.md#RegisterDeviceUsingPOST) | **Post** /api/v0/vendors/lora/devices | Register a LoRa device (use Device management - Interfaces - V1 instead)
[**UnregisterDeviceUsingDELETE**](DeviceManagementForLPWAApi.md#UnregisterDeviceUsingDELETE) | **Delete** /api/v0/vendors/lora/devices/{devEUI} | Unregister a device (use Device management - Connector nodes - V1 instead)
[**UpdateDeviceUsingPATCH1**](DeviceManagementForLPWAApi.md#UpdateDeviceUsingPATCH1) | **Patch** /api/v0/vendors/lora/devices/{devEUI} | Update a device (use Device management - Connector nodes - V1 instead)
[**UpdateGatewayUsingPATCH**](DeviceManagementForLPWAApi.md#UpdateGatewayUsingPATCH) | **Patch** /api/v1/deviceMgt/connectors/lora/gateways/{id} | Update a gateway.



## CountDevicesUsingGET2

> LoraDeviceStatsWeb CountDevicesUsingGET2(ctx, xAPIKEY)

Device counts

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**LoraDeviceStatsWeb**](LoraDeviceStatsWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetDeviceProfilesUsingGET

> []string GetDeviceProfilesUsingGET(ctx, xAPIKEY)

Get LoRa device profiles

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


## GetDeviceProfilesUsingGET1

> []string GetDeviceProfilesUsingGET1(ctx, xAPIKEY)

Get LoRa device profiles (use /api/v1/deviceMgt/connectors/lora/profiles instead)

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


## GetDeviceUsingGET3

> LoraDevice GetDeviceUsingGET3(ctx, devEUI, xAPIKEY)

Get a device (use Device management - Connector nodes - V1 instead)

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string**| DevEUI of the device | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetGatewayUsingGET

> LoraGatewayInfo GetGatewayUsingGET(ctx, id, xAPIKEY)

Get a gateway

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| identifier of the gateway | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**LoraGatewayInfo**](LoraGatewayInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMessageCountUsingGET

> DeferredListenableFutureResultstringstring GetMessageCountUsingGET(ctx, xAPIKEY, optional)

List of message counts group by period

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetMessageCountUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMessageCountUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **period** | **optional.String**| the period to group by | 
 **timezone** | **optional.String**| the timezone of the client | [default to UTC]

### Return type

[**DeferredListenableFutureResultstringstring**](DeferredListenableFutureResult«string,string».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommandUsingGET

> PageableLoraCommand ListCommandUsingGET(ctx, devEUI, xAPIKEY, optional)

List commands

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string**| DevEUI of the device | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListCommandUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCommandUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| The requested page number (optional) | [default to 0]
 **status** | **optional.String**| The command status | 
 **timeRange** | [**optional.Interface of []string**](string.md)| Filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange&#x3D;[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusiv, upperbound is exclusiv. | 
 **sort** | **optional.String**| Field used for sorting. Prefix with &#39;-&#39; for descending order | 

### Return type

[**PageableLoraCommand**](Pageable«LoraCommand».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCommandUsingGET1

> PageableLoraCommand ListCommandUsingGET1(ctx, devEUI, xAPIKEY, optional)

List commands (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string**| DevEUI of the device | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListCommandUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCommandUsingGET1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| The requested page number (optional) | [default to 0]
 **status** | **optional.String**| The command status | 
 **timeRange** | [**optional.Interface of []string**](string.md)| Filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange&#x3D;[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusive, upperbound is exclusive. | 
 **sort** | **optional.String**| Field used for sorting. Prefix with &#39;-&#39; for descending order | 

### Return type

[**PageableLoraCommand**](Pageable«LoraCommand».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListConnectivityPlanUsingGET

> []string ListConnectivityPlanUsingGET(ctx, xAPIKEY)

List connectivity plan

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


## ListConnectivityPlansUsingGET

> []LoraNetworkSubscriptionDetail ListConnectivityPlansUsingGET(ctx, xAPIKEY)

List connectivity plans

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**[]LoraNetworkSubscriptionDetail**](LoraNetworkSubscriptionDetail.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListDevicesUsingGET3

> PageableLoraDevice ListDevicesUsingGET3(ctx, xAPIKEY, optional)

List LoRa devices (use Device management - Connector nodes - V1 instead)

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListDevicesUsingGET3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListDevicesUsingGET3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| Requested page number (optional) | [default to 0]
 **devEUI** | **optional.String**| DevEUI regexp filter | 
 **name** | **optional.String**| Name regexp filter | 
 **status** | **optional.String**| Status filter | 
 **tags** | [**optional.Interface of []string**](string.md)| Tags filter | 
 **sort** | **optional.String**| Field used for sorting (prefix with &#39;-&#39; for descending order) | 

### Return type

[**PageableLoraDevice**](Pageable«LoraDevice».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGatewaysUsingGET

> PageableLoraGatewayInfo ListGatewaysUsingGET(ctx, xAPIKEY, optional)

List gateways

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListGatewaysUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGatewaysUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]
 **id** | **optional.String**| filtering by identifier (regexp filter) of gateway | 
 **name** | **optional.String**| filtering by name (regexp filter) of gateway | 
 **type_** | **optional.String**| filtering by type of gateway | 
 **status** | **optional.String**| filtering by status of gateway | 
 **manufacturer** | **optional.String**| filtering by manufacturer of gateway | 

### Return type

[**PageableLoraGatewayInfo**](Pageable«LoraGatewayInfo».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCommandUsingPOST

> LoraCommand RegisterCommandUsingPOST(ctx, devEUI, xAPIKEY, optional)

Register a command

Restricted to API keys with at least one of the following roles : DEVICE_R, DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string**| DevEUI of the device | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***RegisterCommandUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterCommandUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **req** | [**optional.Interface of LoraCommandWeb**](LoraCommandWeb.md)| A LoRa command | 

### Return type

[**LoraCommand**](LoraCommand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterCommandUsingPOST1

> LoraCommand RegisterCommandUsingPOST1(ctx, devEUI, xAPIKEY, optional)

Register a command (use /api/v1/deviceMgt/connectors/lora/nodes/{devEUI}/downlinks instead)

Restricted to API keys with at least one of the following roles : DEVICE_R, DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string**| DevEUI of the device | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***RegisterCommandUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterCommandUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **req** | [**optional.Interface of LoraCommandWeb**](LoraCommandWeb.md)| A LoRa command | 

### Return type

[**LoraCommand**](LoraCommand.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RegisterDeviceUsingPOST

> LoraDevice RegisterDeviceUsingPOST(ctx, xAPIKEY, optional)

Register a LoRa device (use Device management - Interfaces - V1 instead)

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***RegisterDeviceUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RegisterDeviceUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **req** | [**optional.Interface of LoraDeviceCreateReqWeb**](LoraDeviceCreateReqWeb.md)| A LoRa device | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UnregisterDeviceUsingDELETE

> bool UnregisterDeviceUsingDELETE(ctx, devEUI, xAPIKEY)

Unregister a device (use Device management - Connector nodes - V1 instead)

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string**| DevEUI of the device | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

**bool**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDeviceUsingPATCH1

> LoraDevice UpdateDeviceUsingPATCH1(ctx, devEUI, xAPIKEY, optional)

Update a device (use Device management - Connector nodes - V1 instead)

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devEUI** | **string**| DevEUI of the device | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateDeviceUsingPATCH1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateDeviceUsingPATCH1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **req** | [**optional.Interface of LoraDeviceUpdateReqWeb**](LoraDeviceUpdateReqWeb.md)| A LoRa device (add only fields you want to update) | 

### Return type

[**LoraDevice**](LoraDevice.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGatewayUsingPATCH

> LoraGatewayInfo UpdateGatewayUsingPATCH(ctx, id, xAPIKEY, loraGatewayData)

Update a gateway.

Set the location property to an empty object to remove the gateway location.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| identifier of the gateway | 
**xAPIKEY** | **string**| a valid API key | 
**loraGatewayData** | [**LoraGatewayData**](LoraGatewayData.md)| loraGatewayData | 

### Return type

[**LoraGatewayInfo**](LoraGatewayInfo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

