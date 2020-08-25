# ApiKey

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActingTenantId** | **string** | identifier of real tenant in case of identity usurpation | [optional] 
**ActingUserId** | **string** | identifier of real user in case of identity usurpation | [optional] 
**Active** | **bool** | Switch to activate/deactivate the API Key | [optional] 
**ClientCert** | [**ClientCertificatesConfiguration**](ClientCertificatesConfiguration.md) |  | [optional] 
**CreationTs** | **int64** | Date/time of creation (in ms) | [optional] 
**DebugModeEndTs** | **int64** | Timestamp indicating the end date for the debug mode | [optional] 
**Description** | **string** | Short description of the key | [optional] 
**Expired** | **bool** |  | [optional] 
**From** | **int64** | Date/time of start of validity (in ms) | [optional] 
**Id** | **string** | API key unique identifier, randomly generated at creation | [optional] 
**Label** | **string** | Title of the key | [optional] 
**LastActivity** | **int64** | Date/time of last activity (in ms) | [optional] 
**MasterKey** | **bool** |  | [optional] 
**Nonce** | **string** | Nonce | [optional] 
**ParentId** | **string** | identifier of the parent API key, can be null if API key is a master API key | [optional] 
**RateLimit** | [**RateLimit**](RateLimit.md) |  | [optional] 
**Roles** | **[]string** | list of API key associated roles. | 
**Scope** | [**ScopeApplication**](ScopeApplication.md) |  | [optional] 
**SessionKey** | **bool** |  | [optional] 
**SessionTTL** | **int64** | Duration of validity since the last activity (in ms) | [optional] 
**TenantId** | **string** | identifier of tenant account this API key belongs to | [optional] 
**To** | **int64** |  Date/time of end of validity (in ms) | [optional] 
**UserId** | **string** | identifier of the user account this API key belongs to (or null if not a user session API key) | [optional] 
**UsurpedKey** | **bool** |  | [optional] 
**Value** | **string** | API key value (&#x3D; the secret!) | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


