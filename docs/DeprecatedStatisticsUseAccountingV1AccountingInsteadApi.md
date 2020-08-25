# \DeprecatedStatisticsUseAccountingV1AccountingInsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetTenantStatisticsUsingGET3**](DeprecatedStatisticsUseAccountingV1AccountingInsteadApi.md#GetTenantStatisticsUsingGET3) | **Get** /api/v0/statistics/tenant/{tenantId} | Get tenant statistics for a specific tenant and a range of dates



## GetTenantStatisticsUsingGET3

> TenantStatsWeb GetTenantStatisticsUsingGET3(ctx, tenantId, startDate, endDate, xAPIKEY)

Get tenant statistics for a specific tenant and a range of dates

There are 2 types of data in the statistics :  <br/>> counters : represent a flow (in/out) and are collected throughout the day. example : counters for device creation or device deletion.<br/>> inventories : represent a stock and are counted once a day. example : the total number of registered devices.<br/>The statistics, provided for a tenant and a range of dates, contain 2 main sections :<br/> > 'total': gives an aggregated view for each service. Counters are summed on the period. Inventories represent the max value on the period,<br/> > 'statisticsPerDay': daily statistics. displays the counter and inventory values for each day and each service.<br/><br/> The existing sections are:<br/> - Datazone : each service can store (data storage) or access (advanced search) data in an internal dedicated area: datazone regroups all statistics linked to these actions,<br/> - for HTTP requests: restDataZoneQuery (for dataZone/dataSearch), restOthers (for user/mail/sms), restTraffic (for FIFO/PubSub/Router),<br/> - for MQTT : mqttNoSecure, mqttWebSocketNoSecure (with WebSocket), mqttSSL (with SSL), mqttWebSocketSSL (with WebSocket/SSL)<br/> - for SMS : sms_'server phone number',<br/> - for Devices : the deviceInventory section gives the number of created/deleted devices during the day (counters) and the total number of devices (counted in the beginning of the day at 1.00 am).<br/> - for Msisdn : the msisdnInventory section gives the number of msisdn provisioned/unprovisioned during the day (counters) and the total number of registered msisdns (inventory).<br/> - overall : aggregated connector name; virtual connector containing the data of all the connectors aggregated,<br/><br/> For the current day, the traffic information is updated within 15 minutes. The number of sources/active devices is updated every minute.<br/><u>About the virtual messages:</u><br/>When the MsgIn / MsgOut message size is bigger than 5 kilobytes, a new indicator MsgVirtualIn / MsgVirtualOut  is used to count the number of packets of 5 kilobytes for that message.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantId** | **string**| the id of your tenant ex: 57xxxxxxxxxxxxxxxxxxxxxx | 
**startDate** | **string**| the requested start date as yyyy-MM-dd | 
**endDate** | **string**| the requested end date as yyyy-MM-dd | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**TenantStatsWeb**](TenantStatsWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

