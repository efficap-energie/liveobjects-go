# \DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBusinessSettingsUsingPOST**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#CreateBusinessSettingsUsingPOST) | **Post** /api/v0/sms-connector/settings/business | Create a new business settings of SMS Connector (use Device management - Interfaces - V1 instead)
[**DeleteBusinessSettingsByMsiSDNUsingDELETE**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#DeleteBusinessSettingsByMsiSDNUsingDELETE) | **Delete** /api/v0/sms-connector/settings/business/msisdn | Delete msisdn in business settings of SMS Connector
[**DeleteBusinessSettingsForOneMsiSDNUsingDELETE**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#DeleteBusinessSettingsForOneMsiSDNUsingDELETE) | **Delete** /api/v0/sms-connector/settings/business/{serverPhoneNumber}/msisdn/{msisdnNumber} | Delete one msisdn in business settings of SMS Connector
[**GetBusinessSettingsUsingGET**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#GetBusinessSettingsUsingGET) | **Get** /api/v0/sms-connector/settings/business | Get a business settings of SMS Connector
[**ListMsisdnUsingGET**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#ListMsisdnUsingGET) | **Get** /api/v0/sms-connector/settings/business/msisdn | List msisdn of business settings of SMS Connector
[**ListSettingsUsingGET**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#ListSettingsUsingGET) | **Get** /api/v0/sms-connector/settings | List all the business settings of the SMSConnector for a tenant
[**UpdateBusinessSettingsUsingPOST**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#UpdateBusinessSettingsUsingPOST) | **Post** /api/v0/sms-connector/settings/business/{serverPhoneNumber} | Update business settings of the SMSConnector



## CreateBusinessSettingsUsingPOST

> ConnectorStatusResponse CreateBusinessSettingsUsingPOST(ctx, xAPIKEY, optional)

Create a new business settings of SMS Connector (use Device management - Interfaces - V1 instead)

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateBusinessSettingsUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateBusinessSettingsUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**optional.Interface of SmsConnectorBusinessSettingsCreationReqWeb**](SmsConnectorBusinessSettingsCreationReqWeb.md)| body to create business settings for SMS Connector | 

### Return type

[**ConnectorStatusResponse**](ConnectorStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessSettingsByMsiSDNUsingDELETE

> ConnectorStatusResponse DeleteBusinessSettingsByMsiSDNUsingDELETE(ctx, xAPIKEY, optional)

Delete msisdn in business settings of SMS Connector

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***DeleteBusinessSettingsByMsiSDNUsingDELETEOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteBusinessSettingsByMsiSDNUsingDELETEOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**optional.Interface of SmsConnectorBusinessSettingsDeleteMsisdnReqWeb**](SmsConnectorBusinessSettingsDeleteMsisdnReqWeb.md)| body to delete business settings per msisdn for SMS Connector | 

### Return type

[**ConnectorStatusResponse**](ConnectorStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessSettingsForOneMsiSDNUsingDELETE

> DeleteBusinessSettingsForOneMsiSDNUsingDELETE(ctx, serverPhoneNumber, msisdnNumber, xAPIKEY)

Delete one msisdn in business settings of SMS Connector

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverPhoneNumber** | **string**| server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 
**msisdnNumber** | **string**| msisdn number ex: \&quot;0606060606\&quot; | 
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


## GetBusinessSettingsUsingGET

> BusinessSettings GetBusinessSettingsUsingGET(ctx, serverPhoneNumber, xAPIKEY, optional)

Get a business settings of SMS Connector

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverPhoneNumber** | **string**| the server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetBusinessSettingsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetBusinessSettingsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **decoderName** | **optional.String**| the decoder Name ex: \&quot;decoderName\&quot; | 

### Return type

[**BusinessSettings**](BusinessSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMsisdnUsingGET

> SmsConnectorMsisdnPageWeb ListMsisdnUsingGET(ctx, serverPhoneNumber, xAPIKEY, optional)

List msisdn of business settings of SMS Connector

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverPhoneNumber** | **string**| the server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListMsisdnUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListMsisdnUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **decoderName** | **optional.String**| the decoder Name ex: \&quot;decoderName\&quot; | 
 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**SmsConnectorMsisdnPageWeb**](SMSConnectorMsisdnPageWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSettingsUsingGET

> SmsConnectorBusinessSettingsPageWeb ListSettingsUsingGET(ctx, xAPIKEY, optional)

List all the business settings of the SMSConnector for a tenant

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListSettingsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListSettingsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**SmsConnectorBusinessSettingsPageWeb**](SMSConnectorBusinessSettingsPageWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessSettingsUsingPOST

> ConnectorStatusResponse UpdateBusinessSettingsUsingPOST(ctx, serverPhoneNumber, xAPIKEY, optional)

Update business settings of the SMSConnector

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverPhoneNumber** | **string**| server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateBusinessSettingsUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateBusinessSettingsUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **request** | [**optional.Interface of SmsConnectorBusinessSettingsUpdateReqWeb**](SmsConnectorBusinessSettingsUpdateReqWeb.md)| body to update business settings for SMS Connector | 

### Return type

[**ConnectorStatusResponse**](ConnectorStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

