# SMSConnectorSendSMSReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Base64Payload** | Pointer to **string** | (deprecated) please use hexPayload | [optional] 
**HexPayload** | Pointer to **string** | payload as Hexadecimal PDU Message Entry (GSM 03.40) ex: \&quot;000100009100000E74747A0E4ACF416110BD3CA703\&quot;. If used, don&#39;t set textPayload and base64Payload. | [optional] 
**Msisdns** | **[]string** | device&#39;s SIM Card identifier(s) ex: [\&quot;33666666667\&quot;]. Must be registered in the business settings prior to send SMS only Orange France MSISDNs are supported for now | 
**ServerPhoneNumber** | **string** | server phone number ex: \&quot;20258\&quot;. Must be defined in OfferSettings. | 
**TextPayload** | Pointer to **string** | payload as text ex: \&quot;this is a test\&quot;. If used, don&#39;t set hexPayload and base64Payload. | [optional] 

## Methods

### NewSMSConnectorSendSMSReqWeb

`func NewSMSConnectorSendSMSReqWeb(msisdns []string, serverPhoneNumber string, ) *SMSConnectorSendSMSReqWeb`

NewSMSConnectorSendSMSReqWeb instantiates a new SMSConnectorSendSMSReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSConnectorSendSMSReqWebWithDefaults

`func NewSMSConnectorSendSMSReqWebWithDefaults() *SMSConnectorSendSMSReqWeb`

NewSMSConnectorSendSMSReqWebWithDefaults instantiates a new SMSConnectorSendSMSReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBase64Payload

`func (o *SMSConnectorSendSMSReqWeb) GetBase64Payload() string`

GetBase64Payload returns the Base64Payload field if non-nil, zero value otherwise.

### GetBase64PayloadOk

`func (o *SMSConnectorSendSMSReqWeb) GetBase64PayloadOk() (*string, bool)`

GetBase64PayloadOk returns a tuple with the Base64Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBase64Payload

`func (o *SMSConnectorSendSMSReqWeb) SetBase64Payload(v string)`

SetBase64Payload sets Base64Payload field to given value.

### HasBase64Payload

`func (o *SMSConnectorSendSMSReqWeb) HasBase64Payload() bool`

HasBase64Payload returns a boolean if a field has been set.

### GetHexPayload

`func (o *SMSConnectorSendSMSReqWeb) GetHexPayload() string`

GetHexPayload returns the HexPayload field if non-nil, zero value otherwise.

### GetHexPayloadOk

`func (o *SMSConnectorSendSMSReqWeb) GetHexPayloadOk() (*string, bool)`

GetHexPayloadOk returns a tuple with the HexPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHexPayload

`func (o *SMSConnectorSendSMSReqWeb) SetHexPayload(v string)`

SetHexPayload sets HexPayload field to given value.

### HasHexPayload

`func (o *SMSConnectorSendSMSReqWeb) HasHexPayload() bool`

HasHexPayload returns a boolean if a field has been set.

### GetMsisdns

`func (o *SMSConnectorSendSMSReqWeb) GetMsisdns() []string`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *SMSConnectorSendSMSReqWeb) GetMsisdnsOk() (*[]string, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *SMSConnectorSendSMSReqWeb) SetMsisdns(v []string)`

SetMsisdns sets Msisdns field to given value.


### GetServerPhoneNumber

`func (o *SMSConnectorSendSMSReqWeb) GetServerPhoneNumber() string`

GetServerPhoneNumber returns the ServerPhoneNumber field if non-nil, zero value otherwise.

### GetServerPhoneNumberOk

`func (o *SMSConnectorSendSMSReqWeb) GetServerPhoneNumberOk() (*string, bool)`

GetServerPhoneNumberOk returns a tuple with the ServerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPhoneNumber

`func (o *SMSConnectorSendSMSReqWeb) SetServerPhoneNumber(v string)`

SetServerPhoneNumber sets ServerPhoneNumber field to given value.


### GetTextPayload

`func (o *SMSConnectorSendSMSReqWeb) GetTextPayload() string`

GetTextPayload returns the TextPayload field if non-nil, zero value otherwise.

### GetTextPayloadOk

`func (o *SMSConnectorSendSMSReqWeb) GetTextPayloadOk() (*string, bool)`

GetTextPayloadOk returns a tuple with the TextPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTextPayload

`func (o *SMSConnectorSendSMSReqWeb) SetTextPayload(v string)`

SetTextPayload sets TextPayload field to given value.

### HasTextPayload

`func (o *SMSConnectorSendSMSReqWeb) HasTextPayload() bool`

HasTextPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


