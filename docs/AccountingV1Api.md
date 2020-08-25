# \AccountingV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetDailyStatisticsUsingGET2**](AccountingV1Api.md#GetDailyStatisticsUsingGET2) | **Get** /api/v1/accounting/daily | Get daily accounting metrics (Beta).
[**GetMonthlyStatisticsUsingGET2**](AccountingV1Api.md#GetMonthlyStatisticsUsingGET2) | **Get** /api/v1/accounting/monthly | Get monthly accounting metrics.



## GetDailyStatisticsUsingGET2

> TenantDayMetrics GetDailyStatisticsUsingGET2(ctx, xAPIKEY, optional)

Get daily accounting metrics (Beta).

<p>Get accounting metrics for a given period (from one day to 2 months) over the past 18 months. The data is aggregated by days.</p><p>The response contains 2 main sections : <strong>tenant</strong> and <strong>days</strong>.</p><p>The \"<strong>tenant\"</strong> section gives useful information related to the customer. It returns the current values, whatever has been specified in startDay and endDay parameters :</p><ul><li>id : Live Objects identifier</li><li>name : tenant name in Live Objects</li><li>detailed offer and options, including the display name</li><li>properties : custom fields</li></ul><p><br />The \"<strong>days\"</strong> section provides connectors and service metrics for each day.</p><ul><li><strong>connectors</strong> section (LoRa, mqtt (device), mqtt application, http, sms<strong>_[server phone number]</strong>, external) :<ul><li><strong>traffic</strong> metrics:<ul><li><strong>msg</strong>: number of messages (no unit)</li><li><strong>bytes</strong>: sum of the volume of all messages (unit: bytes)</li><li><strong>virtualMsg</strong>: number of messages split by a size of 5kB (no unit).<br />Example: 1 msg with a size of 7kB = 1 msg and 2 virtualMsg</li></ul></li><li><strong>inventory</strong> metrics:<ul><li><strong>maxNumberOfEnabledNodes</strong>: LoRa specific: max number of simultaneously registered LoRa interfaces for the day. Corresponding LoRa status : all status except DEACTIVATED.</li></ul></li><li><strong>service</strong> section:<ul><li><strong>deviceManagement</strong></li><ul><li><strong>numberOfSuccessfulDownloads</strong>: count of all firmware downloads in success</li></ul><li><strong>storage</strong></li><ul><li><strong>storedDataMessagesInMB</strong>: the total size of messages stored in database.This size does not include technical data (for ex. Index).Values in MB are rounded up.The value differs from the sum of “IN” bytes sent to the platform because it includes data coming from enrichment/decoding/event processing features.This information is available from June 2019.</li></ul></ul></li></ul></ul><p><span style=\"text-decoration: underline;\">Additional info</span> :</p><p>Traffic is split by connector type (mqtt, mqtt application, sms_[server phone number], LoRa, http, external) and by directions (in/out).</p><ul><li><strong>In</strong>: means incoming traffic (i.e. from \"outside\" into Live Objects) (mostly from objects) also called <strong>Uplink</strong>.</li><li><strong>Out</strong>: any outgoing message (i.e. from Live Objects to \"outside\") except protocol messages (MQTT CONNACK, SUBACK … ) and messages consumed by applications on top of Live Objects (subscribe in MQTT application mode).<br /><br /></li></ul><p><strong>http specific:</strong></p><ul><li>IN traffic: sum (POST /api/v0/data/bulk &amp; POST /api/v0/data/streams/{streamId})</li></ul><p><strong>mqtt specific:</strong></p><ul><li>IN & OUT traffic: after may 2020 : the mqtt section contains only traffic in device mode. Mqtt application mode is computed only for incoming traffic (IN).</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetDailyStatisticsUsingGET2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetDailyStatisticsUsingGET2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startDay** | **optional.String**| the requested start day as yyyy-MM-dd. If missing, the current day is used. | 
 **endDay** | **optional.String**| the requested end day as yyyy-MM-dd. If missing, the current day is used. | 

### Return type

[**TenantDayMetrics**](TenantDayMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetMonthlyStatisticsUsingGET2

> TenantMonthMetrics GetMonthlyStatisticsUsingGET2(ctx, xAPIKEY, optional)

Get monthly accounting metrics.

<p>Get accounting metrics for a given period (up to 18 months). The data is aggregated by months.</p><p>The response contains 2 main sections : <strong>tenant</strong> and <strong>months</strong>.</p><p>The \"<strong>tenant\"</strong> section gives useful information related to the customer. It returns the current values, whatever has been specified in startMonth and endMonth parameters :</p><ul><li>id : Live Objects identifier</li><li>name : tenant name in Live Objects</li><li>detailed offer and options, including the display name</li><li>properties : custom fields</li></ul><p><br />The \"<strong>months\"</strong> section provides connectors and service metrics for each month.</p><ul><li><strong>connectors</strong> section (LoRa, mqtt (device), mqtt application, http, sms<strong>_[server phone number]</strong>, external) :<ul><li><strong>traffic</strong> metrics:<ul><li><strong>msg</strong>: number of messages (no unit)</li><li><strong>bytes</strong>: sum of the volume of all messages (unit: bytes)</li><li><strong>virtualMsg</strong>: number of messages split by a size of 5kB (no unit).<br />Example: 1 msg with a size of 7kB = 1 msg and 2 virtualMsg</li></ul></li><li><strong>inventory</strong> metrics:<ul><li><strong>numberOfCommunicatingNodes</strong>: this counter is available by connector: LoRa, mqtt, sms_[server phone number] or external. It is a unique count of devices that send at least one uplink message to the platform (except for mqtt, see below). More precisely, by connector:<ul><li>LoRa: distinct devEUIs sending an uplink</li><li>mqtt: distinct clientIds that have been connected at least one time</li><li>external: distinct nodeIds behind an external connector that have been declared online or have published a dataMessage</li><li>sms: distinct MSISDNs sending a message</li></ul></li><li><strong>numberOfActivatedNodes</strong>: LoRa specific: count of unique devices allowed to communicate and paired to the network. Corresponding LoRa status : ACTIVATED or INITIALIZED.</li><li><strong>numberOfActivatedTdoaNodes</strong>: LoRa specific: count of unique Activated devices that had TDOA lora option for the month.</li><li><strong>numberOfActivatedAckUlNodes</strong>: LoRa specific: count of unique Activated devices that had AckUL lora option for the month. </ul></ul></li></ul><ul><li><strong>service</strong> section:<ul><li><strong>deviceManagement</strong></li><ul><li><strong>numberOfSuccessfulDownloads</strong>: count of all firmware downloads in success</li></ul><li><strong>storage</strong></li><ul><li><strong>storedDataMessagesInMB</strong>: the total size of messages stored in database.This size does not include technical data (for ex. Index).Values in MB are rounded up.The value differs from the sum of “IN” bytes sent to the platform because it includes data coming from enrichment/decoding/event processing features.This information is available from June 2019.</li></ul></ul></li></ul><p><span style=\"text-decoration: underline;\">Additional info</span> :</p><p>Traffic is split by connector type (mqtt, mqtt application, sms_[server phone number], LoRa, http, external) and by directions (in/out).</p><ul><li><strong>In</strong>: means incoming traffic (i.e. from \"outside\" into Live Objects) (mostly from objects) also called <strong>Uplink</strong>.</li><li><strong>Out</strong>: any outgoing message (i.e. from Live Objects to \"outside\") except protocol messages (MQTT CONNACK, SUBACK … ) and messages consumed by applications on top of Live Objects (subscribe in MQTT application mode).<br /><br /></li></ul><p><strong>http specific:</strong></p><ul><li>IN traffic: sum (POST /api/v0/data/bulk &amp; POST /api/v0/data/streams/{streamId})</li></ul><p><strong>mqtt specific:</strong></p><ul><li>IN & OUT traffic: after may 2020 : the mqtt section contains only traffic in device mode. Mqtt application mode is computed only for incoming traffic (IN).</li></ul>

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetMonthlyStatisticsUsingGET2Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetMonthlyStatisticsUsingGET2Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **startMonth** | **optional.String**| the requested start month as yyyy-MM. If missing, the current month is used. | 
 **endMonth** | **optional.String**| the requested end month as yyyy-MM. If missing, the current month is used. | 

### Return type

[**TenantMonthMetrics**](TenantMonthMetrics.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

