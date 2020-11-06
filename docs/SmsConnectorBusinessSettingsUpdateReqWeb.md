# SMSConnectorBusinessSettingsUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecoderName** | Pointer to **string** | decoder name ex: \&quot;decoderName\&quot; | [optional] 
**Msisdns** | Pointer to [**[]Msisdns**](Msisdns.md) | device&#39;s SIM Card identifier ex: \&quot;33666666667\&quot; | [optional] 
**NewDecoderName** | Pointer to **string** | new decoder name ex: \&quot;newDecoderName\&quot; | [optional] 

## Methods

### NewSMSConnectorBusinessSettingsUpdateReqWeb

`func NewSMSConnectorBusinessSettingsUpdateReqWeb() *SMSConnectorBusinessSettingsUpdateReqWeb`

NewSMSConnectorBusinessSettingsUpdateReqWeb instantiates a new SMSConnectorBusinessSettingsUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSConnectorBusinessSettingsUpdateReqWebWithDefaults

`func NewSMSConnectorBusinessSettingsUpdateReqWebWithDefaults() *SMSConnectorBusinessSettingsUpdateReqWeb`

NewSMSConnectorBusinessSettingsUpdateReqWebWithDefaults instantiates a new SMSConnectorBusinessSettingsUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecoderName

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) GetDecoderName() string`

GetDecoderName returns the DecoderName field if non-nil, zero value otherwise.

### GetDecoderNameOk

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) GetDecoderNameOk() (*string, bool)`

GetDecoderNameOk returns a tuple with the DecoderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoderName

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) SetDecoderName(v string)`

SetDecoderName sets DecoderName field to given value.

### HasDecoderName

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) HasDecoderName() bool`

HasDecoderName returns a boolean if a field has been set.

### GetMsisdns

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) GetMsisdns() []Msisdns`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) GetMsisdnsOk() (*[]Msisdns, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) SetMsisdns(v []Msisdns)`

SetMsisdns sets Msisdns field to given value.

### HasMsisdns

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) HasMsisdns() bool`

HasMsisdns returns a boolean if a field has been set.

### GetNewDecoderName

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) GetNewDecoderName() string`

GetNewDecoderName returns the NewDecoderName field if non-nil, zero value otherwise.

### GetNewDecoderNameOk

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) GetNewDecoderNameOk() (*string, bool)`

GetNewDecoderNameOk returns a tuple with the NewDecoderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewDecoderName

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) SetNewDecoderName(v string)`

SetNewDecoderName sets NewDecoderName field to given value.

### HasNewDecoderName

`func (o *SMSConnectorBusinessSettingsUpdateReqWeb) HasNewDecoderName() bool`

HasNewDecoderName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


