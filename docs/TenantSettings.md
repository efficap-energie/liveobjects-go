# TenantSettings

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | Pointer to [**map[string]SettingsSection**](SettingsSection.md) |  | [optional] 
**TenantId** | Pointer to **string** |  | [optional] 

## Methods

### NewTenantSettings

`func NewTenantSettings() *TenantSettings`

NewTenantSettings instantiates a new TenantSettings object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSettingsWithDefaults

`func NewTenantSettingsWithDefaults() *TenantSettings`

NewTenantSettingsWithDefaults instantiates a new TenantSettings object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *TenantSettings) GetSettings() map[string]SettingsSection`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TenantSettings) GetSettingsOk() (*map[string]SettingsSection, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TenantSettings) SetSettings(v map[string]SettingsSection)`

SetSettings sets Settings field to given value.

### HasSettings

`func (o *TenantSettings) HasSettings() bool`

HasSettings returns a boolean if a field has been set.

### GetTenantId

`func (o *TenantSettings) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantSettings) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantSettings) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantSettings) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


