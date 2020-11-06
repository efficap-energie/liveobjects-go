# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | user email | 
**ExternalIdentity** | Pointer to [**ExternalIdentity**](ExternalIdentity.md) |  | [optional] 
**Id** | **string** | identifier of user | 
**LastAuthentication** | Pointer to **int64** | user last authentication timestamp in ms | [optional] 
**Login** | **string** | user login | 
**MainUser** | Pointer to **bool** | user is the tenant&#39;s main user | [optional] 
**PortalData** | Pointer to **map[string]interface{}** | user portal data | [optional] 
**Roles** | **[]string** | list of user associated roles. | 
**State** | **string** | user state : disabled, enabled or suspended | 
**TenantId** | **string** | identifier of tenant account this user will belong to | 

## Methods

### NewUser

`func NewUser(email string, id string, login string, roles []string, state string, tenantId string, ) *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetExternalIdentity

`func (o *User) GetExternalIdentity() ExternalIdentity`

GetExternalIdentity returns the ExternalIdentity field if non-nil, zero value otherwise.

### GetExternalIdentityOk

`func (o *User) GetExternalIdentityOk() (*ExternalIdentity, bool)`

GetExternalIdentityOk returns a tuple with the ExternalIdentity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExternalIdentity

`func (o *User) SetExternalIdentity(v ExternalIdentity)`

SetExternalIdentity sets ExternalIdentity field to given value.

### HasExternalIdentity

`func (o *User) HasExternalIdentity() bool`

HasExternalIdentity returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v string)`

SetId sets Id field to given value.


### GetLastAuthentication

`func (o *User) GetLastAuthentication() int64`

GetLastAuthentication returns the LastAuthentication field if non-nil, zero value otherwise.

### GetLastAuthenticationOk

`func (o *User) GetLastAuthenticationOk() (*int64, bool)`

GetLastAuthenticationOk returns a tuple with the LastAuthentication field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastAuthentication

`func (o *User) SetLastAuthentication(v int64)`

SetLastAuthentication sets LastAuthentication field to given value.

### HasLastAuthentication

`func (o *User) HasLastAuthentication() bool`

HasLastAuthentication returns a boolean if a field has been set.

### GetLogin

`func (o *User) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *User) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *User) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetMainUser

`func (o *User) GetMainUser() bool`

GetMainUser returns the MainUser field if non-nil, zero value otherwise.

### GetMainUserOk

`func (o *User) GetMainUserOk() (*bool, bool)`

GetMainUserOk returns a tuple with the MainUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUser

`func (o *User) SetMainUser(v bool)`

SetMainUser sets MainUser field to given value.

### HasMainUser

`func (o *User) HasMainUser() bool`

HasMainUser returns a boolean if a field has been set.

### GetPortalData

`func (o *User) GetPortalData() map[string]interface{}`

GetPortalData returns the PortalData field if non-nil, zero value otherwise.

### GetPortalDataOk

`func (o *User) GetPortalDataOk() (*map[string]interface{}, bool)`

GetPortalDataOk returns a tuple with the PortalData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPortalData

`func (o *User) SetPortalData(v map[string]interface{})`

SetPortalData sets PortalData field to given value.

### HasPortalData

`func (o *User) HasPortalData() bool`

HasPortalData returns a boolean if a field has been set.

### GetRoles

`func (o *User) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *User) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *User) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetState

`func (o *User) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *User) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *User) SetState(v string)`

SetState sets State field to given value.


### GetTenantId

`func (o *User) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *User) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *User) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


