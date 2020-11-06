# AddPartnerTenantRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | email of the main user account | 
**Language** | **string** | Initial language of the user account | 
**Login** | **string** | login of the main user account (unique) | 
**Name** | **string** | name (unique) of the main user account | 
**OfferMappingId** | **string** | identifier of the offer mapping | 
**Options** | Pointer to [**[]OptionPartnerTenant**](OptionPartnerTenant.md) | options definition | [optional] 
**Properties** | Pointer to **map[string]string** | tenant properties | [optional] 

## Methods

### NewAddPartnerTenantRequest

`func NewAddPartnerTenantRequest(email string, language string, login string, name string, offerMappingId string, ) *AddPartnerTenantRequest`

NewAddPartnerTenantRequest instantiates a new AddPartnerTenantRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAddPartnerTenantRequestWithDefaults

`func NewAddPartnerTenantRequestWithDefaults() *AddPartnerTenantRequest`

NewAddPartnerTenantRequestWithDefaults instantiates a new AddPartnerTenantRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AddPartnerTenantRequest) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AddPartnerTenantRequest) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AddPartnerTenantRequest) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLanguage

`func (o *AddPartnerTenantRequest) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *AddPartnerTenantRequest) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *AddPartnerTenantRequest) SetLanguage(v string)`

SetLanguage sets Language field to given value.


### GetLogin

`func (o *AddPartnerTenantRequest) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AddPartnerTenantRequest) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AddPartnerTenantRequest) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetName

`func (o *AddPartnerTenantRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AddPartnerTenantRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AddPartnerTenantRequest) SetName(v string)`

SetName sets Name field to given value.


### GetOfferMappingId

`func (o *AddPartnerTenantRequest) GetOfferMappingId() string`

GetOfferMappingId returns the OfferMappingId field if non-nil, zero value otherwise.

### GetOfferMappingIdOk

`func (o *AddPartnerTenantRequest) GetOfferMappingIdOk() (*string, bool)`

GetOfferMappingIdOk returns a tuple with the OfferMappingId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferMappingId

`func (o *AddPartnerTenantRequest) SetOfferMappingId(v string)`

SetOfferMappingId sets OfferMappingId field to given value.


### GetOptions

`func (o *AddPartnerTenantRequest) GetOptions() []OptionPartnerTenant`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *AddPartnerTenantRequest) GetOptionsOk() (*[]OptionPartnerTenant, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *AddPartnerTenantRequest) SetOptions(v []OptionPartnerTenant)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *AddPartnerTenantRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetProperties

`func (o *AddPartnerTenantRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AddPartnerTenantRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AddPartnerTenantRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AddPartnerTenantRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


