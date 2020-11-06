# Tenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AccountType** | Pointer to **string** |  | [optional] 
**AdminInfo** | Pointer to [**AdminInfo**](AdminInfo.md) |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**CreationTs** | Pointer to **int64** |  | [optional] 
**DeletionTs** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**IdentityProvider** | Pointer to **string** |  | [optional] 
**MainUserId** | Pointer to **string** |  | [optional] 
**MasterApiKeyId** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** |  | [optional] 
**OfferSettings** | Pointer to [**map[string]SettingsSection**](SettingsSection.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Siret** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TechnicalSettings** | Pointer to [**map[string]SettingsSection**](SettingsSection.md) |  | [optional] 
**UserSettings** | Pointer to [**map[string]SettingsSection**](SettingsSection.md) |  | [optional] 

## Methods

### NewTenant

`func NewTenant() *Tenant`

NewTenant instantiates a new Tenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithDefaults

`func NewTenantWithDefaults() *Tenant`

NewTenantWithDefaults instantiates a new Tenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *Tenant) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *Tenant) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *Tenant) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *Tenant) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAdminInfo

`func (o *Tenant) GetAdminInfo() AdminInfo`

GetAdminInfo returns the AdminInfo field if non-nil, zero value otherwise.

### GetAdminInfoOk

`func (o *Tenant) GetAdminInfoOk() (*AdminInfo, bool)`

GetAdminInfoOk returns a tuple with the AdminInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminInfo

`func (o *Tenant) SetAdminInfo(v AdminInfo)`

SetAdminInfo sets AdminInfo field to given value.

### HasAdminInfo

`func (o *Tenant) HasAdminInfo() bool`

HasAdminInfo returns a boolean if a field has been set.

### GetCountry

`func (o *Tenant) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Tenant) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Tenant) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *Tenant) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreationTs

`func (o *Tenant) GetCreationTs() int64`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *Tenant) GetCreationTsOk() (*int64, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *Tenant) SetCreationTs(v int64)`

SetCreationTs sets CreationTs field to given value.

### HasCreationTs

`func (o *Tenant) HasCreationTs() bool`

HasCreationTs returns a boolean if a field has been set.

### GetDeletionTs

`func (o *Tenant) GetDeletionTs() int64`

GetDeletionTs returns the DeletionTs field if non-nil, zero value otherwise.

### GetDeletionTsOk

`func (o *Tenant) GetDeletionTsOk() (*int64, bool)`

GetDeletionTsOk returns a tuple with the DeletionTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTs

`func (o *Tenant) SetDeletionTs(v int64)`

SetDeletionTs sets DeletionTs field to given value.

### HasDeletionTs

`func (o *Tenant) HasDeletionTs() bool`

HasDeletionTs returns a boolean if a field has been set.

### GetId

`func (o *Tenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Tenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Tenant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Tenant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *Tenant) GetIdentityProvider() string`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *Tenant) GetIdentityProviderOk() (*string, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *Tenant) SetIdentityProvider(v string)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *Tenant) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetMainUserId

`func (o *Tenant) GetMainUserId() string`

GetMainUserId returns the MainUserId field if non-nil, zero value otherwise.

### GetMainUserIdOk

`func (o *Tenant) GetMainUserIdOk() (*string, bool)`

GetMainUserIdOk returns a tuple with the MainUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserId

`func (o *Tenant) SetMainUserId(v string)`

SetMainUserId sets MainUserId field to given value.

### HasMainUserId

`func (o *Tenant) HasMainUserId() bool`

HasMainUserId returns a boolean if a field has been set.

### GetMasterApiKeyId

`func (o *Tenant) GetMasterApiKeyId() string`

GetMasterApiKeyId returns the MasterApiKeyId field if non-nil, zero value otherwise.

### GetMasterApiKeyIdOk

`func (o *Tenant) GetMasterApiKeyIdOk() (*string, bool)`

GetMasterApiKeyIdOk returns a tuple with the MasterApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterApiKeyId

`func (o *Tenant) SetMasterApiKeyId(v string)`

SetMasterApiKeyId sets MasterApiKeyId field to given value.

### HasMasterApiKeyId

`func (o *Tenant) HasMasterApiKeyId() bool`

HasMasterApiKeyId returns a boolean if a field has been set.

### GetName

`func (o *Tenant) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Tenant) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Tenant) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Tenant) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferSettings

`func (o *Tenant) GetOfferSettings() map[string]SettingsSection`

GetOfferSettings returns the OfferSettings field if non-nil, zero value otherwise.

### GetOfferSettingsOk

`func (o *Tenant) GetOfferSettingsOk() (*map[string]SettingsSection, bool)`

GetOfferSettingsOk returns a tuple with the OfferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferSettings

`func (o *Tenant) SetOfferSettings(v map[string]SettingsSection)`

SetOfferSettings sets OfferSettings field to given value.

### HasOfferSettings

`func (o *Tenant) HasOfferSettings() bool`

HasOfferSettings returns a boolean if a field has been set.

### GetParentId

`func (o *Tenant) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *Tenant) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *Tenant) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *Tenant) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProperties

`func (o *Tenant) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Tenant) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Tenant) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Tenant) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSiret

`func (o *Tenant) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *Tenant) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *Tenant) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *Tenant) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetTags

`func (o *Tenant) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Tenant) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Tenant) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Tenant) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTechnicalSettings

`func (o *Tenant) GetTechnicalSettings() map[string]SettingsSection`

GetTechnicalSettings returns the TechnicalSettings field if non-nil, zero value otherwise.

### GetTechnicalSettingsOk

`func (o *Tenant) GetTechnicalSettingsOk() (*map[string]SettingsSection, bool)`

GetTechnicalSettingsOk returns a tuple with the TechnicalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalSettings

`func (o *Tenant) SetTechnicalSettings(v map[string]SettingsSection)`

SetTechnicalSettings sets TechnicalSettings field to given value.

### HasTechnicalSettings

`func (o *Tenant) HasTechnicalSettings() bool`

HasTechnicalSettings returns a boolean if a field has been set.

### GetUserSettings

`func (o *Tenant) GetUserSettings() map[string]SettingsSection`

GetUserSettings returns the UserSettings field if non-nil, zero value otherwise.

### GetUserSettingsOk

`func (o *Tenant) GetUserSettingsOk() (*map[string]SettingsSection, bool)`

GetUserSettingsOk returns a tuple with the UserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSettings

`func (o *Tenant) SetUserSettings(v map[string]SettingsSection)`

SetUserSettings sets UserSettings field to given value.

### HasUserSettings

`func (o *Tenant) HasUserSettings() bool`

HasUserSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


