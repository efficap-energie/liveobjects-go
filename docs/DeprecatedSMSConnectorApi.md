# \DeprecatedSMSConnectorApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSMSUsingPOST**](DeprecatedSMSConnectorApi.md#SendSMSUsingPOST) | **Post** /api/v0/sms-connector/sms | send SMS by SMS Connector for a list of MSISDN



## SendSMSUsingPOST

> ConnectorStatusResponse SendSMSUsingPOST(ctx, xAPIKEY, optional)

send SMS by SMS Connector for a list of MSISDN

Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SendSMSUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SendSMSUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**optional.Interface of SmsConnectorSendSmsReqWeb**](SmsConnectorSendSmsReqWeb.md)| body to send SMS via SMS Connector (only one Payload field must be set) | 

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

