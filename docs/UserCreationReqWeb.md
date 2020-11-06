# UserCreationReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | the user email. Expected a valid email (max 254 characters) | 
**Language** | Pointer to **string** | user language. Expected language ISO 639-1 (example: \&quot;en\&quot;, \&quot;fr\&quot;, \&quot;sk\&quot;, \&quot;ro\&quot;, \&quot;es\&quot;) (max 2 characters) | [optional] 
**Login** | **string** | the user login. If no external identity provider is used, then login must respect the following regular expression &lt;code&gt;[a-zA-Z0-9_-]{3,254}&lt;/code&gt; | 
**Roles** | **[]string** | list of user associated roles. Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | 
**TenantId** | **string** | identifier of tenant account this user will belong to. Expected identifier (max 24 characters) | 

## Methods

### NewUserCreationReqWeb

`func NewUserCreationReqWeb(email string, login string, roles []string, tenantId string, ) *UserCreationReqWeb`

NewUserCreationReqWeb instantiates a new UserCreationReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreationReqWebWithDefaults

`func NewUserCreationReqWebWithDefaults() *UserCreationReqWeb`

NewUserCreationReqWebWithDefaults instantiates a new UserCreationReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *UserCreationReqWeb) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserCreationReqWeb) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserCreationReqWeb) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetLanguage

`func (o *UserCreationReqWeb) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserCreationReqWeb) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserCreationReqWeb) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserCreationReqWeb) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLogin

`func (o *UserCreationReqWeb) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UserCreationReqWeb) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UserCreationReqWeb) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetRoles

`func (o *UserCreationReqWeb) GetRoles() []string`

GetRoles returns the Roles field if non-nil, zero value otherwise.

### GetRolesOk

`func (o *UserCreationReqWeb) GetRolesOk() (*[]string, bool)`

GetRolesOk returns a tuple with the Roles field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoles

`func (o *UserCreationReqWeb) SetRoles(v []string)`

SetRoles sets Roles field to given value.


### GetTenantId

`func (o *UserCreationReqWeb) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserCreationReqWeb) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserCreationReqWeb) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


