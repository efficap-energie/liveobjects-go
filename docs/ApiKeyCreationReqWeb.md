# ApiKeyCreationReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | Pointer to **bool** | Status | [optional] 
**ClientCert** | Pointer to [**ClientCertificatesConfiguration**](ClientCertificatesConfiguration.md) |  | [optional] 
**Description** | Pointer to **string** | Short description of the key. Expected string (limited to 140 first characters) | [optional] 
**From** | Pointer to **int64** | Timestamp from which the key is valid | [optional] 
**Label** | Pointer to **string** | Title of the key. Expected string (limited to 24 first characters) | [optional] 
**ParentId** | **string** | identifier of the parent key to set. Expected identifier (max 24 characters) | 
**Roles** | **[]string** | list of API key associated roles. Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | 
**Scope** | Pointer to [**ScopeApplication**](ScopeApplication.md) |  | [optional] 
**SessionTtl** | Pointer to **int64** | Maximum inactivity period (ms) | [optional] 
**To** | Pointer to **int64** | Timestamp until which the key is valid | [optional] 
**Ttl** | Pointer to **int64** | Time to live (ms) | [optional] 
**UserId** | Pointer to **string** | identifier of the user owner of this key. Expected identifier (max 24 characters) | [optional] 

## Methods

### NewApiKeyCreationReqWeb

`func NewApiKeyCreationReqWeb(parentId string, roles []string, ) *ApiKeyCreationReqWeb`

NewApiKeyCreationReqWeb instantiates a new ApiKeyCreationReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeyCreationReqWebWithDefaults

`func NewApiKeyCreationReqWebWithDefaults() *ApiKeyCreationReqWeb`

NewApiKeyCreationReqWebWithDefaults instantiates a new ApiKeyCreationReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActive

`func (o *ApiKeyCreationReqWeb) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *ApiKeyCreationReqWeb) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *ApiKeyCreationReqWeb) SetActive(v bool)`

SetActive sets Active field to given value.

### HasActive

`func (o *ApiKeyCreationReqWeb) HasActive() bool`

HasActive returns a boolean if a field has been set.

### GetClientCert

`func (o *ApiKeyCreationReqWeb) GetClientCert() ClientCertificatesConfiguration`

GetClientCert returns the ClientCert field if non-nil, zero value otherwise.

### GetClientCertOk

`func (o *ApiKeyCreationReqWeb) GetClientCertOk() (*ClientCertificatesConfiguration, bool)`

GetClientCertOk returns a tuple with the ClientCert field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientCert

`func (o *ApiKeyCreationReqWeb) SetClientCert(v ClientCertificatesConfiguration)`

SetClientCert sets ClientCert field to given value.

### HasClientCert

`func (o *ApiKeyCreationReqWeb) HasClientCert() bool`

HasClientCert returns a boolean if a field has been set.

### GetDescription

`func (o *ApiKeyCreationReqWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ApiKeyCreationReqWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ApiKeyCreationReqWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ApiKeyCreationReqWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFrom

`func (o *ApiKeyCreationReqWeb) GetFrom() int64`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *ApiKeyCreationReqWeb) GetFromOk() (*int64, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *ApiKeyCreationReqWeb) SetFrom(v int64)`

SetFrom sets From field to given value.

### HasFrom

`func (o *ApiKeyCreationReqWeb) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLabel

`func (o *ApiKeyCreationReqWeb) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ApiKeyCreationReqWeb) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ApiKeyCreationReqWeb) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *ApiKeyCreationReqWeb) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetParentId

`func (o *ApiKeyCreationReqWeb) GetParentId() string`

GetParentId returns the ParentId field if non-nil, zero value otherwise.

### GetParentIdOk

`func (o *ApiKeyCreationReqWeb) GetParentIdOk() (*string, bool)`

GetParentIdOk returns a tuple with the ParentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParentId

`func (o *ApiKeyCreationReqWeb) SetParentId(v string)`

SetParentId sets ParentId field to given value.


### GetRoles

`func (o *ApiKeyCreationReqWeb) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *ApiKeyCreationReqWeb) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *ApiKeyCreationReqWeb) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetScope

`func (o *ApiKeyCreationReqWeb) GetScope() ScopeApplication`

GetScope returns the Scope field if non-nil, zero value otherwise.

### GetScopeOk

`func (o *ApiKeyCreationReqWeb) GetScopeOk() (*ScopeApplication, bool)`

GetScopeOk returns a tuple with the Scope field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetScope

`func (o *ApiKeyCreationReqWeb) SetScope(v ScopeApplication)`

SetScope sets Scope field to given value.

### HasScope

`func (o *ApiKeyCreationReqWeb) HasScope() bool`

HasScope returns a boolean if a field has been set.

### GetSessionTtl

`func (o *ApiKeyCreationReqWeb) GetSessionTtl() int64`

GetSessionTtl returns the SessionTtl field if non-nil, zero value otherwise.

### GetSessionTtlOk

`func (o *ApiKeyCreationReqWeb) GetSessionTtlOk() (*int64, bool)`

GetSessionTtlOk returns a tuple with the SessionTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSessionTtl

`func (o *ApiKeyCreationReqWeb) SetSessionTtl(v int64)`

SetSessionTtl sets SessionTtl field to given value.

### HasSessionTtl

`func (o *ApiKeyCreationReqWeb) HasSessionTtl() bool`

HasSessionTtl returns a boolean if a field has been set.

### GetTo

`func (o *ApiKeyCreationReqWeb) GetTo() int64`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ApiKeyCreationReqWeb) GetToOk() (*int64, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ApiKeyCreationReqWeb) SetTo(v int64)`

SetTo sets To field to given value.

### HasTo

`func (o *ApiKeyCreationReqWeb) HasTo() bool`

HasTo returns a boolean if a field has been set.

### GetTtl

`func (o *ApiKeyCreationReqWeb) GetTtl() int64`

GetTtl returns the Ttl field if non-nil, zero value otherwise.

### GetTtlOk

`func (o *ApiKeyCreationReqWeb) GetTtlOk() (*int64, bool)`

GetTtlOk returns a tuple with the Ttl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTtl

`func (o *ApiKeyCreationReqWeb) SetTtl(v int64)`

SetTtl sets Ttl field to given value.

### HasTtl

`func (o *ApiKeyCreationReqWeb) HasTtl() bool`

HasTtl returns a boolean if a field has been set.

### GetUserId

`func (o *ApiKeyCreationReqWeb) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *ApiKeyCreationReqWeb) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *ApiKeyCreationReqWeb) SetUserId(v string)`

SetUserId sets UserId field to given value.

### HasUserId

`func (o *ApiKeyCreationReqWeb) HasUserId() bool`

HasUserId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


