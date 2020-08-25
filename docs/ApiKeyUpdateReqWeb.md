# ApiKeyUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Active** | **bool** | Status | [optional] 
**ClientCert** | [**ClientCertificatesConfiguration**](ClientCertificatesConfiguration.md) |  | [optional] 
**DebugModeEndTs** | **int64** | Timestamp until which the debug mode will be activated | [optional] 
**Description** | **string** | Short description of the key. Expected string (limited to 140 first characters) | [optional] 
**From** | **int64** | Timestamp from which the key is valid | [optional] 
**Label** | **string** | Title of the key. Expected string (limited to 24 first characters) | [optional] 
**MasterKey** | **bool** | Specify the request is for a MasterKey (only update available for admin) | [optional] 
**RegenerateValue** | **bool** | Request a random value regeneration | [optional] 
**Roles** | **[]string** | list of API key associated Roles. Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | [optional] 
**Scope** | [**ScopeApplication**](ScopeApplication.md) |  | [optional] 
**SessionTtl** | **int64** | Maximum inactivity period (ms) | [optional] 
**TenantId** | **string** | identifier of tenant account this API key belongs to. Expected identifier (max 24 characters) | [optional] 
**To** | **int64** | Timestamp until which the key is valid | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


