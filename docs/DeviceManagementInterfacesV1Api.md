# \DeviceManagementInterfacesV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddInterfaceToDeviceUsingPOST**](DeviceManagementInterfacesV1Api.md#AddInterfaceToDeviceUsingPOST) | **Post** /api/v1/deviceMgt/devices/{deviceId}/interfaces | Add an interface to a registered device
[**DeleteInterfaceUsingDELETE**](DeviceManagementInterfacesV1Api.md#DeleteInterfaceUsingDELETE) | **Delete** /api/v1/deviceMgt/devices/{deviceId}/interfaces/{interfaceId} | Delete an interface
[**GetInterfaceForADeviceUsingGET**](DeviceManagementInterfacesV1Api.md#GetInterfaceForADeviceUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/interfaces/{interfaceId} | Get a specific interface for a registered device
[**ListInterfacesForADeviceUsingGET**](DeviceManagementInterfacesV1Api.md#ListInterfacesForADeviceUsingGET) | **Get** /api/v1/deviceMgt/devices/{deviceId}/interfaces | Get the interface list for a registered device
[**UpdateInterfaceUsingPATCH**](DeviceManagementInterfacesV1Api.md#UpdateInterfaceUsingPATCH) | **Patch** /api/v1/deviceMgt/devices/{deviceId}/interfaces/{interfaceId} | Update an interface



## AddInterfaceToDeviceUsingPOST

> DeviceInterface AddInterfaceToDeviceUsingPOST(ctx, deviceId, xAPIKEY, optional)

Add an interface to a registered device

<p>The <strong>definition</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora definition</span>:</p><ul><li>\"devEUI\": interface EUI</li><li>\"appEUI\": appEUI of the interface</li><li>\"appKey\": appKey of the interface</li><li>\"activationType\": supported value: \"OTAA\"</li><li>\"profile\": profile of the interface</li><li>\"encoding\": (Optional) encoding type of the binary payload sent by the interface</li><li>\"connectivityOptions\": connectivity options used for the interface </li><ul><li>\"tdoa\": true/false </li><li>\"ackUl\": true/false</li></ul><li>\"connectivityPlan\": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style=\"text-decoration: underline;\">SMS definition</span>:</p><ul><li>\"msisdn\": interface msisdn</li><li>\"serverPhoneNumber\": (Optional) server phone number</li><li>\"encoding\": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style=\"text-decoration: underline;\">MQTT definition</span>:</p><ul><li>\"clientId\": interface client ID</li></ul><p><span style=\"text-decoration: underline;\">X-CONNECTOR definition</span>:</p><ul><li>\"nodeId\": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li></ul><br><br>Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***AddInterfaceToDeviceUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddInterfaceToDeviceUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **newDeviceInterface** | [**optional.Interface of NewDeviceInterface**](NewDeviceInterface.md)| The device interface to add | 

### Return type

[**DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteInterfaceUsingDELETE

> DeleteInterfaceUsingDELETE(ctx, deviceId, interfaceId, xAPIKEY)

Delete an interface

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**interfaceId** | **string**| Must be {connector}:{nodeId}. | 
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


## GetInterfaceForADeviceUsingGET

> DeviceInterface GetInterfaceForADeviceUsingGET(ctx, deviceId, interfaceId, xAPIKEY)

Get a specific interface for a registered device

<p>The <strong>definition</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora definition</span>:</p><ul><li>\"devEUI\": interface EUI</li><li>\"appEUI\": appEUI of the interface</li><li>\"appKey\": appKey of the interface</li><li>\"activationType\": supported value: \"OTAA\"</li><li>\"profile\": profile of the interface</li><li>\"encoding\": (Optional) encoding type of the binary payload sent by the interface</li><li>\"connectivityOptions\": connectivity options used for the interface </li><ul><li>\"tdoa\": true/false </li><li>\"ackUl\": true/false</li></ul><li>\"connectivityPlan\": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style=\"text-decoration: underline;\">SMS definition</span>:</p><ul><li>\"msisdn\": interface msisdn</li><li>\"serverPhoneNumber\": (Optional) server phone number</li><li>\"encoding\": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style=\"text-decoration: underline;\">MQTT definition</span>:</p><ul><li>\"clientId\": interface client ID</li></ul><p><span style=\"text-decoration: underline;\">X-CONNECTOR definition</span>:</p><ul><li>\"nodeId\": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li></ul><br /><p>The <strong>activity</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora activity</span>:</p><ul><li>\"lastActivationTs\": (Optional) last activation date of the interface</li><li>\"lastDeactivationTs\": (Optional) last deactivation date of the interface</li><li>\"lastSignalLevel\": (Optional) last signal level</li><li>\"lastBatteryLevel\": (Optional) last battery level</li><li>\"lastDlFcnt\": (Optional) last downlink frame counter</li><li>\"lastUlFcnt\": (Optional) last uplink frame counter</li></ul><p><span style=\"text-decoration: underline;\">SMS activity</span>:</p><ul><li>\"lastUplink\": (Optional) last uplink date of the interface</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul><li>\"lastDownlink\": (Optional) last downlink date of the connector node</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul></ul><p><span style=\"text-decoration: underline;\">MQTT activity</span>:</p><ul><li>\"apiKeyId\": api key ID used</li><li>\"mqttVersion\": mqtt version used</li><li>\"mqttUsername\": mqtt username used</li><li>\"mqttTimeout\": mqtt timeout</li><li>\"remoteAddress\"</li><li>\"lastSessionStartTime\"</li><li>\"lastSessionEndTime\"</li></ul><br><br>Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**interfaceId** | **string**| Must be {connector}:{nodeId} | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListInterfacesForADeviceUsingGET

> []DeviceInterface ListInterfacesForADeviceUsingGET(ctx, deviceId, xAPIKEY)

Get the interface list for a registered device

<p>The <strong>definition</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora definition</span>:</p><ul><li>\"devEUI\": interface EUI</li><li>\"appEUI\": appEUI of the interface</li><li>\"appKey\": appKey of the interface</li><li>\"activationType\": supported value: \"OTAA\"</li><li>\"profile\": profile of the interface</li><li>\"encoding\": (Optional) encoding type of the binary payload sent by the interface</li><li>\"connectivityOptions\": connectivity options used for the interface </li><ul><li>\"tdoa\": true/false </li><li>\"ackUl\": true/false</li></ul><li>\"connectivityPlan\": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style=\"text-decoration: underline;\">SMS definition</span>:</p><ul><li>\"msisdn\": interface msisdn</li><li>\"serverPhoneNumber\": (Optional) server phone number</li><li>\"encoding\": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style=\"text-decoration: underline;\">MQTT definition</span>:</p><ul><li>\"clientId\": interface client ID</li></ul><p><span style=\"text-decoration: underline;\">X-CONNECTOR definition</span>:</p><ul><li>\"nodeId\": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li></ul><br /><p>The <strong>activity</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora activity</span>:</p><ul><li>\"lastActivationTs\": (Optional) last activation date of the interface</li><li>\"lastDeactivationTs\": (Optional) last deactivation date of the interface</li><li>\"lastSignalLevel\": (Optional) last signal level</li><li>\"lastBatteryLevel\": (Optional) last battery level</li><li>\"lastDlFcnt\": (Optional) last downlink frame counter</li><li>\"lastUlFcnt\": (Optional) last uplink frame counter</li></ul><p><span style=\"text-decoration: underline;\">SMS activity</span>:</p><ul><li>\"lastUplink\": (Optional) last uplink date of the interface</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul><li>\"lastDownlink\": (Optional) last downlink date of the connector node</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul></ul><p><span style=\"text-decoration: underline;\">MQTT activity</span>:</p><ul><li>\"apiKeyId\": api key ID used</li><li>\"mqttVersion\": mqtt version used</li><li>\"mqttUsername\": mqtt username used</li><li>\"mqttTimeout\": mqtt timeout</li><li>\"remoteAddress\"</li><li>\"lastSessionStartTime\"</li><li>\"lastSessionEndTime\"</li></ul><br><br>Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**[]DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateInterfaceUsingPATCH

> DeviceInterface UpdateInterfaceUsingPATCH(ctx, deviceId, interfaceId, xAPIKEY, optional)

Update an interface

<p>The <strong>definition</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora definition</span>:</p><ul><li>\"devEUI\": interface EUI</li><li>\"appEUI\": appEUI of the interface</li><li>\"appKey\": appKey of the interface</li><li>\"activationType\": supported value: \"OTAA\"</li><li>\"profile\": profile of the interface</li><li>\"encoding\": (Optional) encoding type of the binary payload sent by the interface</li><li>\"connectivityOptions\": connectivity options used for the interface </li><ul><li>\"tdoa\": true/false </li><li>\"ackUl\": true/false</li></ul><li>\"connectivityPlan\": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style=\"text-decoration: underline;\">SMS definition</span>:</p><ul><li>\"msisdn\": interface msisdn</li><li>\"serverPhoneNumber\": (Optional) server phone number</li><li>\"encoding\": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style=\"text-decoration: underline;\">MQTT definition</span>:</p><ul><li>\"clientId\": interface client ID</li></ul><p><span style=\"text-decoration: underline;\">X-CONNECTOR definition</span>:</p><ul><li>\"nodeId\": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li></ul><br><br>Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string**| device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**interfaceId** | **string**| Must be {connector}:{nodeId} | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateInterfaceUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateInterfaceUsingPATCHOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of UpdateInterfaceReqWeb**](UpdateInterfaceReqWeb.md)| The fields to update | 

### Return type

[**DeviceInterface**](DeviceInterface.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

