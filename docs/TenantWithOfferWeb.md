# TenantWithOfferWeb

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
**OfferAndOptions** | Pointer to [**[]OffersAndOptionsReqWeb**](OffersAndOptionsReqWeb.md) |  | [optional] 
**OfferNames** | Pointer to **[]string** |  | [optional] 
**OfferSettings** | Pointer to [**map[string]SettingsSection**](SettingsSection.md) |  | [optional] 
**ParentId** | Pointer to **string** |  | [optional] 
**Properties** | Pointer to **map[string]string** |  | [optional] 
**Siret** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**TechnicalSettings** | Pointer to [**map[string]SettingsSection**](SettingsSection.md) |  | [optional] 
**UserSettings** | Pointer to [**map[string]SettingsSection**](SettingsSection.md) |  | [optional] 

## Methods

### NewTenantWithOfferWeb

`func NewTenantWithOfferWeb() *TenantWithOfferWeb`

NewTenantWithOfferWeb instantiates a new TenantWithOfferWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantWithOfferWebWithDefaults

`func NewTenantWithOfferWebWithDefaults() *TenantWithOfferWeb`

NewTenantWithOfferWebWithDefaults instantiates a new TenantWithOfferWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccountType

`func (o *TenantWithOfferWeb) GetAccountType() string`

GetAccountType returns the AccountType field if non-nil, zero value otherwise.

### GetAccountTypeOk

`func (o *TenantWithOfferWeb) GetAccountTypeOk() (*string, bool)`

GetAccountTypeOk returns a tuple with the AccountType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccountType

`func (o *TenantWithOfferWeb) SetAccountType(v string)`

SetAccountType sets AccountType field to given value.

### HasAccountType

`func (o *TenantWithOfferWeb) HasAccountType() bool`

HasAccountType returns a boolean if a field has been set.

### GetAdminInfo

`func (o *TenantWithOfferWeb) GetAdminInfo() AdminInfo`

GetAdminInfo returns the AdminInfo field if non-nil, zero value otherwise.

### GetAdminInfoOk

`func (o *TenantWithOfferWeb) GetAdminInfoOk() (*AdminInfo, bool)`

GetAdminInfoOk returns a tuple with the AdminInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminInfo

`func (o *TenantWithOfferWeb) SetAdminInfo(v AdminInfo)`

SetAdminInfo sets AdminInfo field to given value.

### HasAdminInfo

`func (o *TenantWithOfferWeb) HasAdminInfo() bool`

HasAdminInfo returns a boolean if a field has been set.

### GetCountry

`func (o *TenantWithOfferWeb) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *TenantWithOfferWeb) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *TenantWithOfferWeb) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *TenantWithOfferWeb) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetCreationTs

`func (o *TenantWithOfferWeb) GetCreationTs() int64`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *TenantWithOfferWeb) GetCreationTsOk() (*int64, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *TenantWithOfferWeb) SetCreationTs(v int64)`

SetCreationTs sets CreationTs field to given value.

### HasCreationTs

`func (o *TenantWithOfferWeb) HasCreationTs() bool`

HasCreationTs returns a boolean if a field has been set.

### GetDeletionTs

`func (o *TenantWithOfferWeb) GetDeletionTs() int64`

GetDeletionTs returns the DeletionTs field if non-nil, zero value otherwise.

### GetDeletionTsOk

`func (o *TenantWithOfferWeb) GetDeletionTsOk() (*int64, bool)`

GetDeletionTsOk returns a tuple with the DeletionTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletionTs

`func (o *TenantWithOfferWeb) SetDeletionTs(v int64)`

SetDeletionTs sets DeletionTs field to given value.

### HasDeletionTs

`func (o *TenantWithOfferWeb) HasDeletionTs() bool`

HasDeletionTs returns a boolean if a field has been set.

### GetId

`func (o *TenantWithOfferWeb) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantWithOfferWeb) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantWithOfferWeb) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantWithOfferWeb) HasId() bool`

HasId returns a boolean if a field has been set.

### GetIdentityProvider

`func (o *TenantWithOfferWeb) GetIdentityProvider() string`

GetIdentityProvider returns the IdentityProvider field if non-nil, zero value otherwise.

### GetIdentityProviderOk

`func (o *TenantWithOfferWeb) GetIdentityProviderOk() (*string, bool)`

GetIdentityProviderOk returns a tuple with the IdentityProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdentityProvider

`func (o *TenantWithOfferWeb) SetIdentityProvider(v string)`

SetIdentityProvider sets IdentityProvider field to given value.

### HasIdentityProvider

`func (o *TenantWithOfferWeb) HasIdentityProvider() bool`

HasIdentityProvider returns a boolean if a field has been set.

### GetMainUserId

`func (o *TenantWithOfferWeb) GetMainUserId() string`

GetMainUserId returns the MainUserId field if non-nil, zero value otherwise.

### GetMainUserIdOk

`func (o *TenantWithOfferWeb) GetMainUserIdOk() (*string, bool)`

GetMainUserIdOk returns a tuple with the MainUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUserId

`func (o *TenantWithOfferWeb) SetMainUserId(v string)`

SetMainUserId sets MainUserId field to given value.

### HasMainUserId

`func (o *TenantWithOfferWeb) HasMainUserId() bool`

HasMainUserId returns a boolean if a field has been set.

### GetMasterApiKeyId

`func (o *TenantWithOfferWeb) GetMasterApiKeyId() string`

GetMasterApiKeyId returns the MasterApiKeyId field if non-nil, zero value otherwise.

### GetMasterApiKeyIdOk

`func (o *TenantWithOfferWeb) GetMasterApiKeyIdOk() (*string, bool)`

GetMasterApiKeyIdOk returns a tuple with the MasterApiKeyId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterApiKeyId

`func (o *TenantWithOfferWeb) SetMasterApiKeyId(v string)`

SetMasterApiKeyId sets MasterApiKeyId field to given value.

### HasMasterApiKeyId

`func (o *TenantWithOfferWeb) HasMasterApiKeyId() bool`

HasMasterApiKeyId returns a boolean if a field has been set.

### GetName

`func (o *TenantWithOfferWeb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantWithOfferWeb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantWithOfferWeb) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantWithOfferWeb) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferAndOptions

`func (o *TenantWithOfferWeb) GetOfferAndOptions() []OffersAndOptionsReqWeb`

GetOfferAndOptions returns the OfferAndOptions field if non-nil, zero value otherwise.

### GetOfferAndOptionsOk

`func (o *TenantWithOfferWeb) GetOfferAndOptionsOk() (*[]OffersAndOptionsReqWeb, bool)`

GetOfferAndOptionsOk returns a tuple with the OfferAndOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferAndOptions

`func (o *TenantWithOfferWeb) SetOfferAndOptions(v []OffersAndOptionsReqWeb)`

SetOfferAndOptions sets OfferAndOptions field to given value.

### HasOfferAndOptions

`func (o *TenantWithOfferWeb) HasOfferAndOptions() bool`

HasOfferAndOptions returns a boolean if a field has been set.

### GetOfferNames

`func (o *TenantWithOfferWeb) GetOfferNames() []string`

GetOfferNames returns the OfferNames field if non-nil, zero value otherwise.

### GetOfferNamesOk

`func (o *TenantWithOfferWeb) GetOfferNamesOk() (*[]string, bool)`

GetOfferNamesOk returns a tuple with the OfferNames field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferNames

`func (o *TenantWithOfferWeb) SetOfferNames(v []string)`

SetOfferNames sets OfferNames field to given value.

### HasOfferNames

`func (o *TenantWithOfferWeb) HasOfferNames() bool`

HasOfferNames returns a boolean if a field has been set.

### GetOfferSettings

`func (o *TenantWithOfferWeb) GetOfferSettings() map[string]SettingsSection`

GetOfferSettings returns the OfferSettings field if non-nil, zero value otherwise.

### GetOfferSettingsOk

`func (o *TenantWithOfferWeb) GetOfferSettingsOk() (*map[string]SettingsSection, bool)`

GetOfferSettingsOk returns a tuple with the OfferSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferSettings

`func (o *TenantWithOfferWeb) SetOfferSettings(v map[string]SettingsSection)`

SetOfferSettings sets OfferSettings field to given value.

### HasOfferSettings

`func (o *TenantWithOfferWeb) HasOfferSettings() bool`

HasOfferSettings returns a boolean if a field has been set.

### GetParentId

`func (o *TenantWithOfferWeb) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *TenantWithOfferWeb) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *TenantWithOfferWeb) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *TenantWithOfferWeb) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetProperties

`func (o *TenantWithOfferWeb) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TenantWithOfferWeb) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TenantWithOfferWeb) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TenantWithOfferWeb) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetSiret

`func (o *TenantWithOfferWeb) GetSiret() string`

GetSiret returns the Siret field if non-nil, zero value otherwise.

### GetSiretOk

`func (o *TenantWithOfferWeb) GetSiretOk() (*string, bool)`

GetSiretOk returns a tuple with the Siret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSiret

`func (o *TenantWithOfferWeb) SetSiret(v string)`

SetSiret sets Siret field to given value.

### HasSiret

`func (o *TenantWithOfferWeb) HasSiret() bool`

HasSiret returns a boolean if a field has been set.

### GetTags

`func (o *TenantWithOfferWeb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *TenantWithOfferWeb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *TenantWithOfferWeb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *TenantWithOfferWeb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTechnicalSettings

`func (o *TenantWithOfferWeb) GetTechnicalSettings() map[string]SettingsSection`

GetTechnicalSettings returns the TechnicalSettings field if non-nil, zero value otherwise.

### GetTechnicalSettingsOk

`func (o *TenantWithOfferWeb) GetTechnicalSettingsOk() (*map[string]SettingsSection, bool)`

GetTechnicalSettingsOk returns a tuple with the TechnicalSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnicalSettings

`func (o *TenantWithOfferWeb) SetTechnicalSettings(v map[string]SettingsSection)`

SetTechnicalSettings sets TechnicalSettings field to given value.

### HasTechnicalSettings

`func (o *TenantWithOfferWeb) HasTechnicalSettings() bool`

HasTechnicalSettings returns a boolean if a field has been set.

### GetUserSettings

`func (o *TenantWithOfferWeb) GetUserSettings() map[string]SettingsSection`

GetUserSettings returns the UserSettings field if non-nil, zero value otherwise.

### GetUserSettingsOk

`func (o *TenantWithOfferWeb) GetUserSettingsOk() (*map[string]SettingsSection, bool)`

GetUserSettingsOk returns a tuple with the UserSettings field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserSettings

`func (o *TenantWithOfferWeb) SetUserSettings(v map[string]SettingsSection)`

SetUserSettings sets UserSettings field to given value.

### HasUserSettings

`func (o *TenantWithOfferWeb) HasUserSettings() bool`

HasUserSettings returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


