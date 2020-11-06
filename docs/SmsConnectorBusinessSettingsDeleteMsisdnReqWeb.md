# SMSConnectorBusinessSettingsDeleteMsisdnReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Msisdns** | **[]string** | device&#39;s SIM Card identifier ex: \&quot;33666666667\&quot; | 
**ServerPhoneNumber** | **string** | server phone number ex: \&quot;20258\&quot; | 

## Methods

### NewSMSConnectorBusinessSettingsDeleteMsisdnReqWeb

`func NewSMSConnectorBusinessSettingsDeleteMsisdnReqWeb(msisdns []string, serverPhoneNumber string, ) *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb`

NewSMSConnectorBusinessSettingsDeleteMsisdnReqWeb instantiates a new SMSConnectorBusinessSettingsDeleteMsisdnReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSConnectorBusinessSettingsDeleteMsisdnReqWebWithDefaults

`func NewSMSConnectorBusinessSettingsDeleteMsisdnReqWebWithDefaults() *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb`

NewSMSConnectorBusinessSettingsDeleteMsisdnReqWebWithDefaults instantiates a new SMSConnectorBusinessSettingsDeleteMsisdnReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMsisdns

`func (o *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb) GetMsisdns() []string`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb) GetMsisdnsOk() (*[]string, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb) SetMsisdns(v []string)`

SetMsisdns sets Msisdns field to given value.


### GetServerPhoneNumber

`func (o *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb) GetServerPhoneNumber() string`

GetServerPhoneNumber returns the ServerPhoneNumber field if non-nil, zero value otherwise.

### GetServerPhoneNumberOk

`func (o *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb) GetServerPhoneNumberOk() (*string, bool)`

GetServerPhoneNumberOk returns a tuple with the ServerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPhoneNumber

`func (o *SMSConnectorBusinessSettingsDeleteMsisdnReqWeb) SetServerPhoneNumber(v string)`

SetServerPhoneNumber sets ServerPhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


