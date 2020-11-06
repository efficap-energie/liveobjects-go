# BusinessSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DecoderName** | Pointer to **string** |  | [optional] 
**Msisdns** | Pointer to [**[]Msisdns**](Msisdns.md) |  | [optional] 
**ServerPhoneNumber** | **string** | platform phone number use as source of message | 
**TenantId** | **string** | Tenant id | 

## Methods

### NewBusinessSettings

`func NewBusinessSettings(serverPhoneNumber string, tenantId string, ) *BusinessSettings`

NewBusinessSettings instantiates a new BusinessSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBusinessSettingsWithDefaults

`func NewBusinessSettingsWithDefaults() *BusinessSettings`

NewBusinessSettingsWithDefaults instantiates a new BusinessSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDecoderName

`func (o *BusinessSettings) GetDecoderName() string`

GetDecoderName returns the DecoderName field if non-nil, zero value otherwise.

### GetDecoderNameOk

`func (o *BusinessSettings) GetDecoderNameOk() (*string, bool)`

GetDecoderNameOk returns a tuple with the DecoderName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDecoderName

`func (o *BusinessSettings) SetDecoderName(v string)`

SetDecoderName sets DecoderName field to given value.

### HasDecoderName

`func (o *BusinessSettings) HasDecoderName() bool`

HasDecoderName returns a boolean if a field has been set.

### GetMsisdns

`func (o *BusinessSettings) GetMsisdns() []Msisdns`

GetMsisdns returns the Msisdns field if non-nil, zero value otherwise.

### GetMsisdnsOk

`func (o *BusinessSettings) GetMsisdnsOk() (*[]Msisdns, bool)`

GetMsisdnsOk returns a tuple with the Msisdns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMsisdns

`func (o *BusinessSettings) SetMsisdns(v []Msisdns)`

SetMsisdns sets Msisdns field to given value.

### HasMsisdns

`func (o *BusinessSettings) HasMsisdns() bool`

HasMsisdns returns a boolean if a field has been set.

### GetServerPhoneNumber

`func (o *BusinessSettings) GetServerPhoneNumber() string`

GetServerPhoneNumber returns the ServerPhoneNumber field if non-nil, zero value otherwise.

### GetServerPhoneNumberOk

`func (o *BusinessSettings) GetServerPhoneNumberOk() (*string, bool)`

GetServerPhoneNumberOk returns a tuple with the ServerPhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServerPhoneNumber

`func (o *BusinessSettings) SetServerPhoneNumber(v string)`

SetServerPhoneNumber sets ServerPhoneNumber field to given value.


### GetTenantId

`func (o *BusinessSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *BusinessSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *BusinessSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


