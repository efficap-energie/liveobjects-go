# \DeviceManagementConnectorNodesV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNodeUsingDELETE**](DeviceManagementConnectorNodesV1Api.md#DeleteNodeUsingDELETE) | **Delete** /api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId} | Delete a connector node
[**GetNodeUsingGET**](DeviceManagementConnectorNodesV1Api.md#GetNodeUsingGET) | **Get** /api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId} | Get details of connector node
[**ListNodesUsingGET**](DeviceManagementConnectorNodesV1Api.md#ListNodesUsingGET) | **Get** /api/v1/deviceMgt/connectors/{connector}/nodes | List all connector nodes
[**UpdateNodeUsingPATCH**](DeviceManagementConnectorNodesV1Api.md#UpdateNodeUsingPATCH) | **Patch** /api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId} | Update a connector node



## DeleteNodeUsingDELETE

> DeleteNodeUsingDELETE(ctx, connector, nodeId, xAPIKEY)

Delete a connector node

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string**| the connector id | 
**nodeId** | **string**| the node id | 
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


## GetNodeUsingGET

> ConnectorNode GetNodeUsingGET(ctx, connector, nodeId, xAPIKEY)

Get details of connector node

<p>The <strong>definition</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora definition</span>:</p><ul><li>\"devEUI\": interface EUI</li><li>\"appEUI\": appEUI of the interface</li><li>\"appKey\": appKey of the interface</li><li>\"activationType\": supported value: \"OTAA\"</li><li>\"profile\": profile of the interface</li><li>\"encoding\": (Optional) encoding type of the binary payload sent by the interface</li><li>\"connectivityOptions\": connectivity options used for the interface </li><ul><li>\"tdoa\": true/false </li><li>\"ackUl\": true/false</li></ul><li>\"connectivityPlan\": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style=\"text-decoration: underline;\">SMS definition</span>:</p><ul><li>\"msisdn\": interface msisdn</li><li>\"serverPhoneNumber\": (Optional) server phone number</li><li>\"encoding\": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style=\"text-decoration: underline;\">MQTT definition</span>:</p><ul><li>\"clientId\": interface client ID</li></ul><p><span style=\"text-decoration: underline;\">X-CONNECTOR definition</span>:</p><ul><li>\"nodeId\": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li></ul><br /><p>The <strong>activity</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora activity</span>:</p><ul><li>\"lastActivationTs\": (Optional) last activation date of the interface</li><li>\"lastDeactivationTs\": (Optional) last deactivation date of the interface</li><li>\"lastSignalLevel\": (Optional) last signal level</li><li>\"lastBatteryLevel\": (Optional) last battery level</li><li>\"lastDlFcnt\": (Optional) last downlink frame counter</li><li>\"lastUlFcnt\": (Optional) last uplink frame counter</li></ul><p><span style=\"text-decoration: underline;\">SMS activity</span>:</p><ul><li>\"lastUplink\": (Optional) last uplink date of the interface</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul><li>\"lastDownlink\": (Optional) last downlink date of the connector node</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul></ul><p><span style=\"text-decoration: underline;\">MQTT activity</span>:</p><ul><li>\"apiKeyId\": api key ID used</li><li>\"mqttVersion\": mqtt version used</li><li>\"mqttUsername\": mqtt username used</li><li>\"mqttTimeout\": mqtt timeout</li><li>\"remoteAddress\"</li><li>\"lastSessionStartTime\"</li><li>\"lastSessionEndTime\"</li></ul><br><br>Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string**| the connector | 
**nodeId** | **string**| the node id | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**ConnectorNode**](ConnectorNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodesUsingGET

> []ConnectorNode ListNodesUsingGET(ctx, connector, xAPIKEY, optional)

List all connector nodes

<p>The <strong>definition</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora definition</span>:</p><ul><li>\"devEUI\": interface EUI</li><li>\"appEUI\": appEUI of the interface</li><li>\"appKey\": appKey of the interface</li><li>\"activationType\": supported value: \"OTAA\"</li><li>\"profile\": profile of the interface</li><li>\"encoding\": (Optional) encoding type of the binary payload sent by the interface</li><li>\"connectivityOptions\": connectivity options used for the interface </li><ul><li>\"tdoa\": true/false </li><li>\"ackUl\": true/false</li></ul><li>\"connectivityPlan\": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style=\"text-decoration: underline;\">SMS definition</span>:</p><ul><li>\"msisdn\": interface msisdn</li><li>\"serverPhoneNumber\": (Optional) server phone number</li><li>\"encoding\": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style=\"text-decoration: underline;\">MQTT definition</span>:</p><ul><li>\"clientId\": interface client ID</li></ul><p><span style=\"text-decoration: underline;\">X-CONNECTOR definition</span>:</p><ul><li>\"nodeId\": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li></ul><br /><p>The <strong>activity</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora activity</span>:</p><ul><li>\"lastActivationTs\": (Optional) last activation date of the interface</li><li>\"lastDeactivationTs\": (Optional) last deactivation date of the interface</li><li>\"lastSignalLevel\": (Optional) last signal level</li><li>\"lastBatteryLevel\": (Optional) last battery level</li><li>\"lastDlFcnt\": (Optional) last downlink frame counter</li><li>\"lastUlFcnt\": (Optional) last uplink frame counter</li></ul><p><span style=\"text-decoration: underline;\">SMS activity</span>:</p><ul><li>\"lastUplink\": (Optional) last uplink date of the interface</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul><li>\"lastDownlink\": (Optional) last downlink date of the connector node</li><ul><li>\"timestamp\": date of the activity</li><li>\"serverPhoneNumber\": server phone number used</li></ul></ul><p><span style=\"text-decoration: underline;\">MQTT activity</span>:</p><ul><li>\"apiKeyId\": api key ID used</li><li>\"mqttVersion\": mqtt version used</li><li>\"mqttUsername\": mqtt username used</li><li>\"mqttTimeout\": mqtt timeout</li><li>\"remoteAddress\"</li><li>\"lastSessionStartTime\"</li><li>\"lastSessionEndTime\"</li></ul><br><br>Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string**| the connector | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListNodesUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListNodesUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **offset** | **optional.Int32**| number of items to skip (optional) | [default to 0]
 **xTotalCount** | **optional.Bool**| true if a total count must be returned in response | [default to false]

### Return type

[**[]ConnectorNode**](ConnectorNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeUsingPATCH

> ConnectorNode UpdateNodeUsingPATCH(ctx, connector, nodeId, xAPIKEY, optional)

Update a connector node

<p>The <strong>definition</strong> depends on connector.</p><p><span style=\"text-decoration: underline;\">Lora definition</span>:</p><ul><li>\"devEUI\": interface EUI</li><li>\"appEUI\": appEUI of the interface</li><li>\"appKey\": appKey of the interface</li><li>\"activationType\": supported value: \"OTAA\"</li><li>\"profile\": profile of the interface</li><li>\"encoding\": (Optional) encoding type of the binary payload sent by the interface</li><li>\"connectivityOptions\": connectivity options used for the interface </li><ul><li>\"tdoa\": true/false </li><li>\"ackUl\": true/false</li></ul><li>\"connectivityPlan\": connectivity plan to use for the interface. Should be used preferably at connectivityOptions but at least one of two shall be defined.</li></ul><p><span style=\"text-decoration: underline;\">SMS definition</span>:</p><ul><li>\"msisdn\": interface msisdn</li><li>\"serverPhoneNumber\": (Optional) server phone number</li><li>\"encoding\": (Optional) name of the decoder that will be used to decode received SMS</li></ul><p><span style=\"text-decoration: underline;\">MQTT definition</span>:</p><ul><li>\"clientId\": interface client ID</li></ul><p><span style=\"text-decoration: underline;\">X-CONNECTOR definition</span>:</p><ul><li>\"nodeId\": id used by the external connector to identify the communicating device (e.g. PUB connector/v1/nodes/{nodeId}/status)</li></ul><br><br>Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string**| the connector id | 
**nodeId** | **string**| the node id | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateNodeUsingPATCHOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateNodeUsingPATCHOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**optional.Interface of UpdateConnectorNodeRequest**](UpdateConnectorNodeRequest.md)| The fields to update | 

### Return type

[**ConnectorNode**](ConnectorNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

