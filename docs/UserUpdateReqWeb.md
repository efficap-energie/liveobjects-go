# UserUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | user email. Expected a valid email (max 254 characters) | [optional] 
**Language** | Pointer to **string** | user language. Expected language ISO 639-1 (example: \&quot;en\&quot;, \&quot;fr\&quot;, \&quot;sk\&quot;, \&quot;ro\&quot;, \&quot;es\&quot;) (max 2 characters) | [optional] 
**Login** | Pointer to **string** | the user login. If no external identity provider is used, then login must respect the following regular expression &lt;code&gt;[a-zA-Z0-9_-]{3,254}&lt;/code&gt; | [optional] 
**MainUser** | Pointer to **bool** | Set the user as main user. Expected boolean (true/false) | [optional] 
**Roles** | Pointer to **[]string** | list of user associated roles. Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | [optional] 
**State** | **string** | user state : disabled, enabled or suspended | 
**TenantId** | Pointer to **string** | identifier of tenant account this user belongs to. Expected identifier (max 24 characters) | [optional] 

## Methods

### NewUserUpdateReqWeb

`func NewUserUpdateReqWeb(state string, ) *UserUpdateReqWeb`

NewUserUpdateReqWeb instantiates a new UserUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserUpdateReqWebWithDefaults

`func NewUserUpdateReqWebWithDefaults() *UserUpdateReqWeb`

NewUserUpdateReqWebWithDefaults instantiates a new UserUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserUpdateReqWeb) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserUpdateReqWeb) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserUpdateReqWeb) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserUpdateReqWeb) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLanguage

`func (o *UserUpdateReqWeb) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserUpdateReqWeb) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserUpdateReqWeb) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserUpdateReqWeb) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLogin

`func (o *UserUpdateReqWeb) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserUpdateReqWeb) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserUpdateReqWeb) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *UserUpdateReqWeb) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetMainUser

`func (o *UserUpdateReqWeb) GetMainUser() bool`

GetMainUser returns the MainUser field if non-nil, zero value otherwise.

### GetMainUserOk

`func (o *UserUpdateReqWeb) GetMainUserOk() (*bool, bool)`

GetMainUserOk returns a tuple with the MainUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainUser

`func (o *UserUpdateReqWeb) SetMainUser(v bool)`

SetMainUser sets MainUser field to given value.

### HasMainUser

`func (o *UserUpdateReqWeb) HasMainUser() bool`

HasMainUser returns a boolean if a field has been set.

### GetRoles

`func (o *UserUpdateReqWeb) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserUpdateReqWeb) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserUpdateReqWeb) SetRoles(v []string)`

SetRoles sets Roles field to given value.

### HasRoles

`func (o *UserUpdateReqWeb) HasRoles() bool`

HasRoles returns a boolean if a field has been set.

### GetState

`func (o *UserUpdateReqWeb) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *UserUpdateReqWeb) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *UserUpdateReqWeb) SetState(v string)`

SetState sets State field to given value.


### GetTenantId

`func (o *UserUpdateReqWeb) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserUpdateReqWeb) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserUpdateReqWeb) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *UserUpdateReqWeb) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


