# SMSConnectorBusinessSettingsCreationReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecoderName** | Pointer to **string** | decoder name ex: \&quot;decoderName\&quot; | [optional] 
**Msisdns** | [**[]Msisdns**](Msisdns.md) | device msisdn | 
**ServerPhoneNumber** | **string** | server phone number ex: \&quot;20258\&quot; | 

## Methods

### NewSMSConnectorBusinessSettingsCreationReqWeb

`func NewSMSConnectorBusinessSettingsCreationReqWeb(msisdns []Msisdns, serverPhoneNumber string, ) *SMSConnectorBusinessSettingsCreationReqWeb`

NewSMSConnectorBusinessSettingsCreationReqWeb instantiates a new SMSConnectorBusinessSettingsCreationReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSConnectorBusinessSettingsCreationReqWebWithDefaults

`func NewSMSConnectorBusinessSettingsCreationReqWebWithDefaults() *SMSConnectorBusinessSettingsCreationReqWeb`

NewSMSConnectorBusinessSettingsCreationReqWebWithDefaults instantiates a new SMSConnectorBusinessSettingsCreationReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecoderName

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) GetDecoderName() string`

GetDecoderName returns the DecoderName field if non-nil, zero value otherwise.

### GetDecoderNameOk

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) GetDecoderNameOk() (*string, bool)`

GetDecoderNameOk returns a tuple with the DecoderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoderName

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) SetDecoderName(v string)`

SetDecoderName sets DecoderName field to given value.

### HasDecoderName

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) HasDecoderName() bool`

HasDecoderName returns a boolean if a field has been set.

### GetMsisdns

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) GetMsisdns() []Msisdns`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) GetMsisdnsOk() (*[]Msisdns, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) SetMsisdns(v []Msisdns)`

SetMsisdns sets Msisdns field to given value.


### GetServerPhoneNumber

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) GetServerPhoneNumber() string`

GetServerPhoneNumber returns the ServerPhoneNumber field if non-nil, zero value otherwise.

### GetServerPhoneNumberOk

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) GetServerPhoneNumberOk() (*string, bool)`

GetServerPhoneNumberOk returns a tuple with the ServerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPhoneNumber

`func (o *SMSConnectorBusinessSettingsCreationReqWeb) SetServerPhoneNumber(v string)`

SetServerPhoneNumber sets ServerPhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


