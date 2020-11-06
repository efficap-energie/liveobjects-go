# TenantSettingsUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Settings** | **map[string]interface{}** | Content of the section to set | 
**SettingsRevision** | **int32** | Current revision of the section, for history of value changes. | 
**SettingsVersion** | Pointer to **string** | Version of the settings model | [optional] 

## Methods

### NewTenantSettingsUpdateReqWeb

`func NewTenantSettingsUpdateReqWeb(settings map[string]interface{}, settingsRevision int32, ) *TenantSettingsUpdateReqWeb`

NewTenantSettingsUpdateReqWeb instantiates a new TenantSettingsUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantSettingsUpdateReqWebWithDefaults

`func NewTenantSettingsUpdateReqWebWithDefaults() *TenantSettingsUpdateReqWeb`

NewTenantSettingsUpdateReqWebWithDefaults instantiates a new TenantSettingsUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSettings

`func (o *TenantSettingsUpdateReqWeb) GetSettings() map[string]interface{}`

GetSettings returns the Settings field if non-nil, zero value otherwise.

### GetSettingsOk

`func (o *TenantSettingsUpdateReqWeb) GetSettingsOk() (*map[string]interface{}, bool)`

GetSettingsOk returns a tuple with the Settings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettings

`func (o *TenantSettingsUpdateReqWeb) SetSettings(v map[string]interface{})`

SetSettings sets Settings field to given value.


### GetSettingsRevision

`func (o *TenantSettingsUpdateReqWeb) GetSettingsRevision() int32`

GetSettingsRevision returns the SettingsRevision field if non-nil, zero value otherwise.

### GetSettingsRevisionOk

`func (o *TenantSettingsUpdateReqWeb) GetSettingsRevisionOk() (*int32, bool)`

GetSettingsRevisionOk returns a tuple with the SettingsRevision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsRevision

`func (o *TenantSettingsUpdateReqWeb) SetSettingsRevision(v int32)`

SetSettingsRevision sets SettingsRevision field to given value.


### GetSettingsVersion

`func (o *TenantSettingsUpdateReqWeb) GetSettingsVersion() string`

GetSettingsVersion returns the SettingsVersion field if non-nil, zero value otherwise.

### GetSettingsVersionOk

`func (o *TenantSettingsUpdateReqWeb) GetSettingsVersionOk() (*string, bool)`

GetSettingsVersionOk returns a tuple with the SettingsVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSettingsVersion

`func (o *TenantSettingsUpdateReqWeb) SetSettingsVersion(v string)`

SetSettingsVersion sets SettingsVersion field to given value.

### HasSettingsVersion

`func (o *TenantSettingsUpdateReqWeb) HasSettingsVersion() bool`

HasSettingsVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


