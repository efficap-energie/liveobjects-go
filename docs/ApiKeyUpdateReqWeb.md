# ApiKeyUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Status | [optional] 
**ClientCert** | Pointer to [**ClientCertificatesConfiguration**](ClientCertificatesConfiguration.md) |  | [optional] 
**DebugModeEndTs** | Pointer to **int64** | Timestamp until which the debug mode will be activated | [optional] 
**Description** | Pointer to **string** | Short description of the key. Expected string (limited to 140 first characters) | [optional] 
**From** | Pointer to **int64** | Timestamp from which the key is valid | [optional] 
**Label** | Pointer to **string** | Title of the key. Expected string (limited to 24 first characters) | [optional] 
**MasterKey** | Pointer to **bool** | Specify the request is for a MasterKey (only update available for admin) | [optional] 
**RegenerateValue** | Pointer to **bool** | Request a random value regeneration | [optional] 
**Roles** | Pointer to **[]string** | list of API key associated Roles. Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | [optional] 
**Scope** | Pointer to [**ScopeApplication**](ScopeApplication.md) |  | [optional] 
**SessionTtl** | Pointer to **int64** | Maximum inactivity period (ms) | [optional] 
**TenantId** | Pointer to **string** | identifier of tenant account this API key belongs to. Expected identifier (max 24 characters) | [optional] 
**To** | Pointer to **int64** | Timestamp until which the key is valid | [optional] 

## Methods

### NewApiKeyUpdateReqWeb

`func NewApiKeyUpdateReqWeb() *ApiKeyUpdateReqWeb`

NewApiKeyUpdateReqWeb instantiates a new ApiKeyUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyUpdateReqWebWithDefaults

`func NewApiKeyUpdateReqWebWithDefaults() *ApiKeyUpdateReqWeb`

NewApiKeyUpdateReqWebWithDefaults instantiates a new ApiKeyUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiKeyUpdateReqWeb) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyUpdateReqWeb) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyUpdateReqWeb) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKeyUpdateReqWeb) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetClientCert

`func (o *ApiKeyUpdateReqWeb) GetClientCert() ClientCertificatesConfiguration`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ApiKeyUpdateReqWeb) GetClientCertOk() (*ClientCertificatesConfiguration, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ApiKeyUpdateReqWeb) SetClientCert(v ClientCertificatesConfiguration)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ApiKeyUpdateReqWeb) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetDebugModeEndTs

`func (o *ApiKeyUpdateReqWeb) GetDebugModeEndTs() int64`

GetDebugModeEndTs returns the DebugModeEndTs field if non-nil, zero value otherwise.

### GetDebugModeEndTsOk

`func (o *ApiKeyUpdateReqWeb) GetDebugModeEndTsOk() (*int64, bool)`

GetDebugModeEndTsOk returns a tuple with the DebugModeEndTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDebugModeEndTs

`func (o *ApiKeyUpdateReqWeb) SetDebugModeEndTs(v int64)`

SetDebugModeEndTs sets DebugModeEndTs field to given value.

### HasDebugModeEndTs

`func (o *ApiKeyUpdateReqWeb) HasDebugModeEndTs() bool`

HasDebugModeEndTs returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyUpdateReqWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyUpdateReqWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyUpdateReqWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyUpdateReqWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrom

`func (o *ApiKeyUpdateReqWeb) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ApiKeyUpdateReqWeb) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ApiKeyUpdateReqWeb) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ApiKeyUpdateReqWeb) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLabel

`func (o *ApiKeyUpdateReqWeb) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiKeyUpdateReqWeb) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiKeyUpdateReqWeb) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApiKeyUpdateReqWeb) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMasterKey

`func (o *ApiKeyUpdateReqWeb) GetMasterKey() bool`

GetMasterKey returns the MasterKey field if non-nil, zero value otherwise.

### GetMasterKeyOk

`func (o *ApiKeyUpdateReqWeb) GetMasterKeyOk() (*bool, bool)`

GetMasterKeyOk returns a tuple with the MasterKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterKey

`func (o *ApiKeyUpdateReqWeb) SetMasterKey(v bool)`

SetMasterKey sets MasterKey field to given value.

### HasMasterKey

`func (o *ApiKeyUpdateReqWeb) HasMasterKey() bool`

HasMasterKey returns a boolean if a field has been set.

### GetRegenerateValue

`func (o *ApiKeyUpdateReqWeb) GetRegenerateValue() bool`

GetRegenerateValue returns the RegenerateValue field if non-nil, zero value otherwise.

### GetRegenerateValueOk

`func (o *ApiKeyUpdateReqWeb) GetRegenerateValueOk() (*bool, bool)`

GetRegenerateValueOk returns a tuple with the RegenerateValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegenerateValue

`func (o *ApiKeyUpdateReqWeb) SetRegenerateValue(v bool)`

SetRegenerateValue sets RegenerateValue field to given value.

### HasRegenerateValue

`func (o *ApiKeyUpdateReqWeb) HasRegenerateValue() bool`

HasRegenerateValue returns a boolean if a field has been set.

### GetRoles

`func (o *ApiKeyUpdateReqWeb) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiKeyUpdateReqWeb) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiKeyUpdateReqWeb) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *ApiKeyUpdateReqWeb) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetScope

`func (o *ApiKeyUpdateReqWeb) GetScope() ScopeApplication`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiKeyUpdateReqWeb) GetScopeOk() (*ScopeApplication, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiKeyUpdateReqWeb) SetScope(v ScopeApplication)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ApiKeyUpdateReqWeb) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSessionTtl

`func (o *ApiKeyUpdateReqWeb) GetSessionTtl() int64`

GetSessionTtl returns the SessionTtl field if non-nil, zero value otherwise.

### GetSessionTtlOk

`func (o *ApiKeyUpdateReqWeb) GetSessionTtlOk() (*int64, bool)`

GetSessionTtlOk returns a tuple with the SessionTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTtl

`func (o *ApiKeyUpdateReqWeb) SetSessionTtl(v int64)`

SetSessionTtl sets SessionTtl field to given value.

### HasSessionTtl

`func (o *ApiKeyUpdateReqWeb) HasSessionTtl() bool`

HasSessionTtl returns a boolean if a field has been set.

### GetTenantId

`func (o *ApiKeyUpdateReqWeb) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *ApiKeyUpdateReqWeb) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *ApiKeyUpdateReqWeb) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *ApiKeyUpdateReqWeb) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTo

`func (o *ApiKeyUpdateReqWeb) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ApiKeyUpdateReqWeb) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ApiKeyUpdateReqWeb) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *ApiKeyUpdateReqWeb) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


