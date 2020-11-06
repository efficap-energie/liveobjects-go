# SMSConnectorSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Decoders** | **[]string** | decoders use to translate messages received | 
**Enabled** | **bool** | indicate if setting is enabled or not | 
**Limits** | Pointer to [**Thresholds**](Thresholds.md) |  | [optional] 
**ServerPhoneNumber** | **string** | platform phone number use as source of message | 

## Methods

### NewSMSConnectorSettings

`func NewSMSConnectorSettings(decoders []string, enabled bool, serverPhoneNumber string, ) *SMSConnectorSettings`

NewSMSConnectorSettings instantiates a new SMSConnectorSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSConnectorSettingsWithDefaults

`func NewSMSConnectorSettingsWithDefaults() *SMSConnectorSettings`

NewSMSConnectorSettingsWithDefaults instantiates a new SMSConnectorSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecoders

`func (o *SMSConnectorSettings) GetDecoders() []string`

GetDecoders returns the Decoders field if non-nil, zero value otherwise.

### GetDecodersOk

`func (o *SMSConnectorSettings) GetDecodersOk() (*[]string, bool)`

GetDecodersOk returns a tuple with the Decoders field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoders

`func (o *SMSConnectorSettings) SetDecoders(v []string)`

SetDecoders sets Decoders field to given value.


### GetEnabled

`func (o *SMSConnectorSettings) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *SMSConnectorSettings) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *SMSConnectorSettings) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetLimits

`func (o *SMSConnectorSettings) GetLimits() Thresholds`

GetLimits returns the Limits field if non-nil, zero value otherwise.

### GetLimitsOk

`func (o *SMSConnectorSettings) GetLimitsOk() (*Thresholds, bool)`

GetLimitsOk returns a tuple with the Limits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimits

`func (o *SMSConnectorSettings) SetLimits(v Thresholds)`

SetLimits sets Limits field to given value.

### HasLimits

`func (o *SMSConnectorSettings) HasLimits() bool`

HasLimits returns a boolean if a field has been set.

### GetServerPhoneNumber

`func (o *SMSConnectorSettings) GetServerPhoneNumber() string`

GetServerPhoneNumber returns the ServerPhoneNumber field if non-nil, zero value otherwise.

### GetServerPhoneNumberOk

`func (o *SMSConnectorSettings) GetServerPhoneNumberOk() (*string, bool)`

GetServerPhoneNumberOk returns a tuple with the ServerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPhoneNumber

`func (o *SMSConnectorSettings) SetServerPhoneNumber(v string)`

SetServerPhoneNumber sets ServerPhoneNumber field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


