# UserUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | user email. Expected a valid email (max 254 characters) | [optional] 
**Language** | **string** | user language. Expected language ISO 639-1 (example: \&quot;en\&quot;, \&quot;fr\&quot;, \&quot;sk\&quot;, \&quot;ro\&quot;, \&quot;es\&quot;) (max 2 characters) | [optional] 
**Login** | **string** | the user login. If no external identity provider is used, then login must respect the following regular expression &lt;code&gt;[a-zA-Z0-9_-]{3,254}&lt;/code&gt; | [optional] 
**MainUser** | **bool** | Set the user as main user. Expected boolean (true/false) | [optional] 
**Roles** | **[]string** | list of user associated roles. Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | [optional] 
**State** | **string** | user state : disabled, enabled or suspended | 
**TenantId** | **string** | identifier of tenant account this user belongs to. Expected identifier (max 24 characters) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


