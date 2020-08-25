# UserCreationReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | the user email. Expected a valid email (max 254 characters) | 
**Language** | **string** | user language. Expected language ISO 639-1 (example: \&quot;en\&quot;, \&quot;fr\&quot;, \&quot;sk\&quot;, \&quot;ro\&quot;, \&quot;es\&quot;) (max 2 characters) | [optional] 
**Login** | **string** | the user login. If no external identity provider is used, then login must respect the following regular expression &lt;code&gt;[a-zA-Z0-9_-]{3,254}&lt;/code&gt; | 
**Roles** | **[]string** | list of user associated roles. Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | 
**TenantId** | **string** | identifier of tenant account this user will belong to. Expected identifier (max 24 characters) | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


