# SmsConnectorSendSmsReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base64Payload** | **string** | (deprecated) please use hexPayload | [optional] 
**HexPayload** | **string** | payload as Hexadecimal PDU Message Entry (GSM 03.40) ex: \&quot;000100009100000E74747A0E4ACF416110BD3CA703\&quot;. If used, don&#39;t set textPayload and base64Payload. | [optional] 
**Msisdns** | **[]string** | device&#39;s SIM Card identifier(s) ex: [\&quot;33666666667\&quot;]. Must be registered in the business settings prior to send SMS only Orange France MSISDNs are supported for now | 
**ServerPhoneNumber** | **string** | server phone number ex: \&quot;20258\&quot;. Must be defined in OfferSettings. | 
**TextPayload** | **string** | payload as text ex: \&quot;this is a test\&quot;. If used, don&#39;t set hexPayload and base64Payload. | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


