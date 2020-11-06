# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingTenantId** | Pointer to **string** | identifier of real tenant in case of identity usurpation | [optional] 
**ActingUserId** | Pointer to **string** | identifier of real user in case of identity usurpation | [optional] 
**Active** | Pointer to **bool** | Switch to activate/deactivate the API Key | [optional] 
**ClientCert** | Pointer to [**ClientCertificatesConfiguration**](ClientCertificatesConfiguration.md) |  | [optional] 
**CreationTs** | Pointer to **int64** | Date/time of creation (in ms) | [optional] 
**DebugModeEndTs** | Pointer to **int64** | Timestamp indicating the end date for the debug mode | [optional] 
**Description** | Pointer to **string** | Short description of the key | [optional] 
**Expired** | Pointer to **bool** |  | [optional] 
**From** | Pointer to **int64** | Date/time of start of validity (in ms) | [optional] 
**Id** | Pointer to **string** | API key unique identifier, randomly generated at creation | [optional] 
**Label** | Pointer to **string** | Title of the key | [optional] 
**LastActivity** | Pointer to **int64** | Date/time of last activity (in ms) | [optional] 
**MasterKey** | Pointer to **bool** |  | [optional] 
**Nonce** | Pointer to **string** | Nonce | [optional] 
**ParentId** | Pointer to **string** | identifier of the parent API key, can be null if API key is a master API key | [optional] 
**RateLimit** | Pointer to [**RateLimit**](RateLimit.md) |  | [optional] 
**Roles** | **[]string** | list of API key associated roles. | 
**Scope** | Pointer to [**ScopeApplication**](ScopeApplication.md) |  | [optional] 
**SessionKey** | Pointer to **bool** |  | [optional] 
**SessionTTL** | Pointer to **int64** | Duration of validity since the last activity (in ms) | [optional] 
**TenantId** | Pointer to **string** | identifier of tenant account this API key belongs to | [optional] 
**To** | Pointer to **int64** |  Date/time of end of validity (in ms) | [optional] 
**UserId** | Pointer to **string** | identifier of the user account this API key belongs to (or null if not a user session API key) | [optional] 
**UsurpedKey** | Pointer to **bool** |  | [optional] 
**Value** | Pointer to **string** | API key value (&#x3D; the secret!) | [optional] 

## Methods

### NewApiKey

`func NewApiKey(roles []string, ) *ApiKey`

NewApiKey instantiates a new ApiKey object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyWithDefaults

`func NewApiKeyWithDefaults() *ApiKey`

NewApiKeyWithDefaults instantiates a new ApiKey object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActingTenantId

`func (o *ApiKey) GetActingTenantId() string`

GetActingTenantId returns the ActingTenantId field if non-nil, zero value otherwise.

### GetActingTenantIdOk

`func (o *ApiKey) GetActingTenantIdOk() (*string, bool)`

GetActingTenantIdOk returns a tuple with the ActingTenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingTenantId

`func (o *ApiKey) SetActingTenantId(v string)`

SetActingTenantId sets ActingTenantId field to given value.

### HasActingTenantId

`func (o *ApiKey) HasActingTenantId() bool`

HasActingTenantId returns a boolean if a field has been set.

### GetActingUserId

`func (o *ApiKey) GetActingUserId() string`

GetActingUserId returns the ActingUserId field if non-nil, zero value otherwise.

### GetActingUserIdOk

`func (o *ApiKey) GetActingUserIdOk() (*string, bool)`

GetActingUserIdOk returns a tuple with the ActingUserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActingUserId

`func (o *ApiKey) SetActingUserId(v string)`

SetActingUserId sets ActingUserId field to given value.

### HasActingUserId

`func (o *ApiKey) HasActingUserId() bool`

HasActingUserId returns a boolean if a field has been set.

### GetActive

`func (o *ApiKey) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKey) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKey) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKey) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetClientCert

`func (o *ApiKey) GetClientCert() ClientCertificatesConfiguration`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ApiKey) GetClientCertOk() (*ClientCertificatesConfiguration, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ApiKey) SetClientCert(v ClientCertificatesConfiguration)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ApiKey) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetCreationTs

`func (o *ApiKey) GetCreationTs() int64`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *ApiKey) GetCreationTsOk() (*int64, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *ApiKey) SetCreationTs(v int64)`

SetCreationTs sets CreationTs field to given value.

### HasCreationTs

`func (o *ApiKey) HasCreationTs() bool`

HasCreationTs returns a boolean if a field has been set.

### GetDebugModeEndTs

`func (o *ApiKey) GetDebugModeEndTs() int64`

GetDebugModeEndTs returns the DebugModeEndTs field if non-nil, zero value otherwise.

### GetDebugModeEndTsOk

`func (o *ApiKey) GetDebugModeEndTsOk() (*int64, bool)`

GetDebugModeEndTsOk returns a tuple with the DebugModeEndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugModeEndTs

`func (o *ApiKey) SetDebugModeEndTs(v int64)`

SetDebugModeEndTs sets DebugModeEndTs field to given value.

### HasDebugModeEndTs

`func (o *ApiKey) HasDebugModeEndTs() bool`

HasDebugModeEndTs returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKey) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKey) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKey) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKey) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetExpired

`func (o *ApiKey) GetExpired() bool`

GetExpired returns the Expired field if non-nil, zero value otherwise.

### GetExpiredOk

`func (o *ApiKey) GetExpiredOk() (*bool, bool)`

GetExpiredOk returns a tuple with the Expired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpired

`func (o *ApiKey) SetExpired(v bool)`

SetExpired sets Expired field to given value.

### HasExpired

`func (o *ApiKey) HasExpired() bool`

HasExpired returns a boolean if a field has been set.

### GetFrom

`func (o *ApiKey) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ApiKey) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ApiKey) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ApiKey) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetId

`func (o *ApiKey) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ApiKey) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ApiKey) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ApiKey) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *ApiKey) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiKey) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiKey) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApiKey) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetLastActivity

`func (o *ApiKey) GetLastActivity() int64`

GetLastActivity returns the LastActivity field if non-nil, zero value otherwise.

### GetLastActivityOk

`func (o *ApiKey) GetLastActivityOk() (*int64, bool)`

GetLastActivityOk returns a tuple with the LastActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivity

`func (o *ApiKey) SetLastActivity(v int64)`

SetLastActivity sets LastActivity field to given value.

### HasLastActivity

`func (o *ApiKey) HasLastActivity() bool`

HasLastActivity returns a boolean if a field has been set.

### GetMasterKey

`func (o *ApiKey) GetMasterKey() bool`

GetMasterKey returns the MasterKey field if non-nil, zero value otherwise.

### GetMasterKeyOk

`func (o *ApiKey) GetMasterKeyOk() (*bool, bool)`

GetMasterKeyOk returns a tuple with the MasterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKey

`func (o *ApiKey) SetMasterKey(v bool)`

SetMasterKey sets MasterKey field to given value.

### HasMasterKey

`func (o *ApiKey) HasMasterKey() bool`

HasMasterKey returns a boolean if a field has been set.

### GetNonce

`func (o *ApiKey) GetNonce() string`

GetNonce returns the Nonce field if non-nil, zero value otherwise.

### GetNonceOk

`func (o *ApiKey) GetNonceOk() (*string, bool)`

GetNonceOk returns a tuple with the Nonce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNonce

`func (o *ApiKey) SetNonce(v string)`

SetNonce sets Nonce field to given value.

### HasNonce

`func (o *ApiKey) HasNonce() bool`

HasNonce returns a boolean if a field has been set.

### GetParentId

`func (o *ApiKey) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ApiKey) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ApiKey) SetParentId(v string)`

SetParentId sets ParentId field to given value.

### HasParentId

`func (o *ApiKey) HasParentId() bool`

HasParentId returns a boolean if a field has been set.

### GetRateLimit

`func (o *ApiKey) GetRateLimit() RateLimit`

GetRateLimit returns the RateLimit field if non-nil, zero value otherwise.

### GetRateLimitOk

`func (o *ApiKey) GetRateLimitOk() (*RateLimit, bool)`

GetRateLimitOk returns a tuple with the RateLimit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRateLimit

`func (o *ApiKey) SetRateLimit(v RateLimit)`

SetRateLimit sets RateLimit field to given value.

### HasRateLimit

`func (o *ApiKey) HasRateLimit() bool`

HasRateLimit returns a boolean if a field has been set.

### GetRoles

`func (o *ApiKey) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiKey) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiKey) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetScope

`func (o *ApiKey) GetScope() ScopeApplication`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiKey) GetScopeOk() (*ScopeApplication, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiKey) SetScope(v ScopeApplication)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ApiKey) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSessionKey

`func (o *ApiKey) GetSessionKey() bool`

GetSessionKey returns the SessionKey field if non-nil, zero value otherwise.

### GetSessionKeyOk

`func (o *ApiKey) GetSessionKeyOk() (*bool, bool)`

GetSessionKeyOk returns a tuple with the SessionKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionKey

`func (o *ApiKey) SetSessionKey(v bool)`

SetSessionKey sets SessionKey field to given value.

### HasSessionKey

`func (o *ApiKey) HasSessionKey() bool`

HasSessionKey returns a boolean if a field has been set.

### GetSessionTTL

`func (o *ApiKey) GetSessionTTL() int64`

GetSessionTTL returns the SessionTTL field if non-nil, zero value otherwise.

### GetSessionTTLOk

`func (o *ApiKey) GetSessionTTLOk() (*int64, bool)`

GetSessionTTLOk returns a tuple with the SessionTTL field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTTL

`func (o *ApiKey) SetSessionTTL(v int64)`

SetSessionTTL sets SessionTTL field to given value.

### HasSessionTTL

`func (o *ApiKey) HasSessionTTL() bool`

HasSessionTTL returns a boolean if a field has been set.

### GetTenantId

`func (o *ApiKey) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ApiKey) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ApiKey) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ApiKey) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTo

`func (o *ApiKey) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ApiKey) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ApiKey) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *ApiKey) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetUserId

`func (o *ApiKey) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKey) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKey) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKey) HasUserId() bool`

HasUserId returns a boolean if a field has been set.

### GetUsurpedKey

`func (o *ApiKey) GetUsurpedKey() bool`

GetUsurpedKey returns the UsurpedKey field if non-nil, zero value otherwise.

### GetUsurpedKeyOk

`func (o *ApiKey) GetUsurpedKeyOk() (*bool, bool)`

GetUsurpedKeyOk returns a tuple with the UsurpedKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsurpedKey

`func (o *ApiKey) SetUsurpedKey(v bool)`

SetUsurpedKey sets UsurpedKey field to given value.

### HasUsurpedKey

`func (o *ApiKey) HasUsurpedKey() bool`

HasUsurpedKey returns a boolean if a field has been set.

### GetValue

`func (o *ApiKey) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ApiKey) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ApiKey) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *ApiKey) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


