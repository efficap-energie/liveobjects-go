/*
 * Live Objects REST API Guide v2.13.3
 *
 * API description for Live Objects service
 *
 * API version: 2.13.3
 * Contact: liveobjects.support@orange.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package liveobjects

import (
	"encoding/json"
)

// TenantWithOfferWeb struct for TenantWithOfferWeb
type TenantWithOfferWeb struct {
	AccountType *string `json:"accountType,omitempty"`
	AdminInfo *AdminInfo `json:"adminInfo,omitempty"`
	Country *string `json:"country,omitempty"`
	CreationTs *int64 `json:"creationTs,omitempty"`
	DeletionTs *int64 `json:"deletionTs,omitempty"`
	Id *string `json:"id,omitempty"`
	IdentityProvider *string `json:"identityProvider,omitempty"`
	MainUserId *string `json:"mainUserId,omitempty"`
	MasterApiKeyId *string `json:"masterApiKeyId,omitempty"`
	Name *string `json:"name,omitempty"`
	OfferAndOptions *[]OffersAndOptionsReqWeb `json:"offerAndOptions,omitempty"`
	OfferNames *[]string `json:"offerNames,omitempty"`
	OfferSettings *map[string]SettingsSection `json:"offerSettings,omitempty"`
	ParentId *string `json:"parentId,omitempty"`
	Properties *map[string]string `json:"properties,omitempty"`
	Siret *string `json:"siret,omitempty"`
	Tags *[]string `json:"tags,omitempty"`
	TechnicalSettings *map[string]SettingsSection `json:"technicalSettings,omitempty"`
	UserSettings *map[string]SettingsSection `json:"userSettings,omitempty"`
}

// NewTenantWithOfferWeb instantiates a new TenantWithOfferWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTenantWithOfferWeb() *TenantWithOfferWeb {
	this := TenantWithOfferWeb{}
	return &this
}

// NewTenantWithOfferWebWithDefaults instantiates a new TenantWithOfferWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTenantWithOfferWebWithDefaults() *TenantWithOfferWeb {
	this := TenantWithOfferWeb{}
	return &this
}

// GetAccountType returns the AccountType field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetAccountType() string {
	if o == nil || o.AccountType == nil {
		var ret string
		return ret
	}
	return *o.AccountType
}

// GetAccountTypeOk returns a tuple with the AccountType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetAccountTypeOk() (*string, bool) {
	if o == nil || o.AccountType == nil {
		return nil, false
	}
	return o.AccountType, true
}

// HasAccountType returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasAccountType() bool {
	if o != nil && o.AccountType != nil {
		return true
	}

	return false
}

// SetAccountType gets a reference to the given string and assigns it to the AccountType field.
func (o *TenantWithOfferWeb) SetAccountType(v string) {
	o.AccountType = &v
}

// GetAdminInfo returns the AdminInfo field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetAdminInfo() AdminInfo {
	if o == nil || o.AdminInfo == nil {
		var ret AdminInfo
		return ret
	}
	return *o.AdminInfo
}

// GetAdminInfoOk returns a tuple with the AdminInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetAdminInfoOk() (*AdminInfo, bool) {
	if o == nil || o.AdminInfo == nil {
		return nil, false
	}
	return o.AdminInfo, true
}

// HasAdminInfo returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasAdminInfo() bool {
	if o != nil && o.AdminInfo != nil {
		return true
	}

	return false
}

// SetAdminInfo gets a reference to the given AdminInfo and assigns it to the AdminInfo field.
func (o *TenantWithOfferWeb) SetAdminInfo(v AdminInfo) {
	o.AdminInfo = &v
}

// GetCountry returns the Country field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetCountry() string {
	if o == nil || o.Country == nil {
		var ret string
		return ret
	}
	return *o.Country
}

// GetCountryOk returns a tuple with the Country field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetCountryOk() (*string, bool) {
	if o == nil || o.Country == nil {
		return nil, false
	}
	return o.Country, true
}

// HasCountry returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasCountry() bool {
	if o != nil && o.Country != nil {
		return true
	}

	return false
}

// SetCountry gets a reference to the given string and assigns it to the Country field.
func (o *TenantWithOfferWeb) SetCountry(v string) {
	o.Country = &v
}

// GetCreationTs returns the CreationTs field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetCreationTs() int64 {
	if o == nil || o.CreationTs == nil {
		var ret int64
		return ret
	}
	return *o.CreationTs
}

// GetCreationTsOk returns a tuple with the CreationTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetCreationTsOk() (*int64, bool) {
	if o == nil || o.CreationTs == nil {
		return nil, false
	}
	return o.CreationTs, true
}

// HasCreationTs returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasCreationTs() bool {
	if o != nil && o.CreationTs != nil {
		return true
	}

	return false
}

// SetCreationTs gets a reference to the given int64 and assigns it to the CreationTs field.
func (o *TenantWithOfferWeb) SetCreationTs(v int64) {
	o.CreationTs = &v
}

// GetDeletionTs returns the DeletionTs field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetDeletionTs() int64 {
	if o == nil || o.DeletionTs == nil {
		var ret int64
		return ret
	}
	return *o.DeletionTs
}

// GetDeletionTsOk returns a tuple with the DeletionTs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetDeletionTsOk() (*int64, bool) {
	if o == nil || o.DeletionTs == nil {
		return nil, false
	}
	return o.DeletionTs, true
}

// HasDeletionTs returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasDeletionTs() bool {
	if o != nil && o.DeletionTs != nil {
		return true
	}

	return false
}

// SetDeletionTs gets a reference to the given int64 and assigns it to the DeletionTs field.
func (o *TenantWithOfferWeb) SetDeletionTs(v int64) {
	o.DeletionTs = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *TenantWithOfferWeb) SetId(v string) {
	o.Id = &v
}

// GetIdentityProvider returns the IdentityProvider field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetIdentityProvider() string {
	if o == nil || o.IdentityProvider == nil {
		var ret string
		return ret
	}
	return *o.IdentityProvider
}

// GetIdentityProviderOk returns a tuple with the IdentityProvider field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetIdentityProviderOk() (*string, bool) {
	if o == nil || o.IdentityProvider == nil {
		return nil, false
	}
	return o.IdentityProvider, true
}

// HasIdentityProvider returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasIdentityProvider() bool {
	if o != nil && o.IdentityProvider != nil {
		return true
	}

	return false
}

// SetIdentityProvider gets a reference to the given string and assigns it to the IdentityProvider field.
func (o *TenantWithOfferWeb) SetIdentityProvider(v string) {
	o.IdentityProvider = &v
}

// GetMainUserId returns the MainUserId field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetMainUserId() string {
	if o == nil || o.MainUserId == nil {
		var ret string
		return ret
	}
	return *o.MainUserId
}

// GetMainUserIdOk returns a tuple with the MainUserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetMainUserIdOk() (*string, bool) {
	if o == nil || o.MainUserId == nil {
		return nil, false
	}
	return o.MainUserId, true
}

// HasMainUserId returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasMainUserId() bool {
	if o != nil && o.MainUserId != nil {
		return true
	}

	return false
}

// SetMainUserId gets a reference to the given string and assigns it to the MainUserId field.
func (o *TenantWithOfferWeb) SetMainUserId(v string) {
	o.MainUserId = &v
}

// GetMasterApiKeyId returns the MasterApiKeyId field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetMasterApiKeyId() string {
	if o == nil || o.MasterApiKeyId == nil {
		var ret string
		return ret
	}
	return *o.MasterApiKeyId
}

// GetMasterApiKeyIdOk returns a tuple with the MasterApiKeyId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetMasterApiKeyIdOk() (*string, bool) {
	if o == nil || o.MasterApiKeyId == nil {
		return nil, false
	}
	return o.MasterApiKeyId, true
}

// HasMasterApiKeyId returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasMasterApiKeyId() bool {
	if o != nil && o.MasterApiKeyId != nil {
		return true
	}

	return false
}

// SetMasterApiKeyId gets a reference to the given string and assigns it to the MasterApiKeyId field.
func (o *TenantWithOfferWeb) SetMasterApiKeyId(v string) {
	o.MasterApiKeyId = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TenantWithOfferWeb) SetName(v string) {
	o.Name = &v
}

// GetOfferAndOptions returns the OfferAndOptions field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetOfferAndOptions() []OffersAndOptionsReqWeb {
	if o == nil || o.OfferAndOptions == nil {
		var ret []OffersAndOptionsReqWeb
		return ret
	}
	return *o.OfferAndOptions
}

// GetOfferAndOptionsOk returns a tuple with the OfferAndOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetOfferAndOptionsOk() (*[]OffersAndOptionsReqWeb, bool) {
	if o == nil || o.OfferAndOptions == nil {
		return nil, false
	}
	return o.OfferAndOptions, true
}

// HasOfferAndOptions returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasOfferAndOptions() bool {
	if o != nil && o.OfferAndOptions != nil {
		return true
	}

	return false
}

// SetOfferAndOptions gets a reference to the given []OffersAndOptionsReqWeb and assigns it to the OfferAndOptions field.
func (o *TenantWithOfferWeb) SetOfferAndOptions(v []OffersAndOptionsReqWeb) {
	o.OfferAndOptions = &v
}

// GetOfferNames returns the OfferNames field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetOfferNames() []string {
	if o == nil || o.OfferNames == nil {
		var ret []string
		return ret
	}
	return *o.OfferNames
}

// GetOfferNamesOk returns a tuple with the OfferNames field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetOfferNamesOk() (*[]string, bool) {
	if o == nil || o.OfferNames == nil {
		return nil, false
	}
	return o.OfferNames, true
}

// HasOfferNames returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasOfferNames() bool {
	if o != nil && o.OfferNames != nil {
		return true
	}

	return false
}

// SetOfferNames gets a reference to the given []string and assigns it to the OfferNames field.
func (o *TenantWithOfferWeb) SetOfferNames(v []string) {
	o.OfferNames = &v
}

// GetOfferSettings returns the OfferSettings field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetOfferSettings() map[string]SettingsSection {
	if o == nil || o.OfferSettings == nil {
		var ret map[string]SettingsSection
		return ret
	}
	return *o.OfferSettings
}

// GetOfferSettingsOk returns a tuple with the OfferSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetOfferSettingsOk() (*map[string]SettingsSection, bool) {
	if o == nil || o.OfferSettings == nil {
		return nil, false
	}
	return o.OfferSettings, true
}

// HasOfferSettings returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasOfferSettings() bool {
	if o != nil && o.OfferSettings != nil {
		return true
	}

	return false
}

// SetOfferSettings gets a reference to the given map[string]SettingsSection and assigns it to the OfferSettings field.
func (o *TenantWithOfferWeb) SetOfferSettings(v map[string]SettingsSection) {
	o.OfferSettings = &v
}

// GetParentId returns the ParentId field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetParentId() string {
	if o == nil || o.ParentId == nil {
		var ret string
		return ret
	}
	return *o.ParentId
}

// GetParentIdOk returns a tuple with the ParentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetParentIdOk() (*string, bool) {
	if o == nil || o.ParentId == nil {
		return nil, false
	}
	return o.ParentId, true
}

// HasParentId returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasParentId() bool {
	if o != nil && o.ParentId != nil {
		return true
	}

	return false
}

// SetParentId gets a reference to the given string and assigns it to the ParentId field.
func (o *TenantWithOfferWeb) SetParentId(v string) {
	o.ParentId = &v
}

// GetProperties returns the Properties field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetProperties() map[string]string {
	if o == nil || o.Properties == nil {
		var ret map[string]string
		return ret
	}
	return *o.Properties
}

// GetPropertiesOk returns a tuple with the Properties field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetPropertiesOk() (*map[string]string, bool) {
	if o == nil || o.Properties == nil {
		return nil, false
	}
	return o.Properties, true
}

// HasProperties returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasProperties() bool {
	if o != nil && o.Properties != nil {
		return true
	}

	return false
}

// SetProperties gets a reference to the given map[string]string and assigns it to the Properties field.
func (o *TenantWithOfferWeb) SetProperties(v map[string]string) {
	o.Properties = &v
}

// GetSiret returns the Siret field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetSiret() string {
	if o == nil || o.Siret == nil {
		var ret string
		return ret
	}
	return *o.Siret
}

// GetSiretOk returns a tuple with the Siret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetSiretOk() (*string, bool) {
	if o == nil || o.Siret == nil {
		return nil, false
	}
	return o.Siret, true
}

// HasSiret returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasSiret() bool {
	if o != nil && o.Siret != nil {
		return true
	}

	return false
}

// SetSiret gets a reference to the given string and assigns it to the Siret field.
func (o *TenantWithOfferWeb) SetSiret(v string) {
	o.Siret = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetTags() []string {
	if o == nil || o.Tags == nil {
		var ret []string
		return ret
	}
	return *o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetTagsOk() (*[]string, bool) {
	if o == nil || o.Tags == nil {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasTags() bool {
	if o != nil && o.Tags != nil {
		return true
	}

	return false
}

// SetTags gets a reference to the given []string and assigns it to the Tags field.
func (o *TenantWithOfferWeb) SetTags(v []string) {
	o.Tags = &v
}

// GetTechnicalSettings returns the TechnicalSettings field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetTechnicalSettings() map[string]SettingsSection {
	if o == nil || o.TechnicalSettings == nil {
		var ret map[string]SettingsSection
		return ret
	}
	return *o.TechnicalSettings
}

// GetTechnicalSettingsOk returns a tuple with the TechnicalSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetTechnicalSettingsOk() (*map[string]SettingsSection, bool) {
	if o == nil || o.TechnicalSettings == nil {
		return nil, false
	}
	return o.TechnicalSettings, true
}

// HasTechnicalSettings returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasTechnicalSettings() bool {
	if o != nil && o.TechnicalSettings != nil {
		return true
	}

	return false
}

// SetTechnicalSettings gets a reference to the given map[string]SettingsSection and assigns it to the TechnicalSettings field.
func (o *TenantWithOfferWeb) SetTechnicalSettings(v map[string]SettingsSection) {
	o.TechnicalSettings = &v
}

// GetUserSettings returns the UserSettings field value if set, zero value otherwise.
func (o *TenantWithOfferWeb) GetUserSettings() map[string]SettingsSection {
	if o == nil || o.UserSettings == nil {
		var ret map[string]SettingsSection
		return ret
	}
	return *o.UserSettings
}

// GetUserSettingsOk returns a tuple with the UserSettings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TenantWithOfferWeb) GetUserSettingsOk() (*map[string]SettingsSection, bool) {
	if o == nil || o.UserSettings == nil {
		return nil, false
	}
	return o.UserSettings, true
}

// HasUserSettings returns a boolean if a field has been set.
func (o *TenantWithOfferWeb) HasUserSettings() bool {
	if o != nil && o.UserSettings != nil {
		return true
	}

	return false
}

// SetUserSettings gets a reference to the given map[string]SettingsSection and assigns it to the UserSettings field.
func (o *TenantWithOfferWeb) SetUserSettings(v map[string]SettingsSection) {
	o.UserSettings = &v
}

func (o TenantWithOfferWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AccountType != nil {
		toSerialize["accountType"] = o.AccountType
	}
	if o.AdminInfo != nil {
		toSerialize["adminInfo"] = o.AdminInfo
	}
	if o.Country != nil {
		toSerialize["country"] = o.Country
	}
	if o.CreationTs != nil {
		toSerialize["creationTs"] = o.CreationTs
	}
	if o.DeletionTs != nil {
		toSerialize["deletionTs"] = o.DeletionTs
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if o.IdentityProvider != nil {
		toSerialize["identityProvider"] = o.IdentityProvider
	}
	if o.MainUserId != nil {
		toSerialize["mainUserId"] = o.MainUserId
	}
	if o.MasterApiKeyId != nil {
		toSerialize["masterApiKeyId"] = o.MasterApiKeyId
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OfferAndOptions != nil {
		toSerialize["offerAndOptions"] = o.OfferAndOptions
	}
	if o.OfferNames != nil {
		toSerialize["offerNames"] = o.OfferNames
	}
	if o.OfferSettings != nil {
		toSerialize["offerSettings"] = o.OfferSettings
	}
	if o.ParentId != nil {
		toSerialize["parentId"] = o.ParentId
	}
	if o.Properties != nil {
		toSerialize["properties"] = o.Properties
	}
	if o.Siret != nil {
		toSerialize["siret"] = o.Siret
	}
	if o.Tags != nil {
		toSerialize["tags"] = o.Tags
	}
	if o.TechnicalSettings != nil {
		toSerialize["technicalSettings"] = o.TechnicalSettings
	}
	if o.UserSettings != nil {
		toSerialize["userSettings"] = o.UserSettings
	}
	return json.Marshal(toSerialize)
}

type NullableTenantWithOfferWeb struct {
	value *TenantWithOfferWeb
	isSet bool
}

func (v NullableTenantWithOfferWeb) Get() *TenantWithOfferWeb {
	return v.value
}

func (v *NullableTenantWithOfferWeb) Set(val *TenantWithOfferWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableTenantWithOfferWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableTenantWithOfferWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTenantWithOfferWeb(val *TenantWithOfferWeb) *NullableTenantWithOfferWeb {
	return &NullableTenantWithOfferWeb{value: val, isSet: true}
}

func (v NullableTenantWithOfferWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTenantWithOfferWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


