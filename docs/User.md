# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | **string** | user email | 
**ExternalIdentity** | [**ExternalIdentity**](ExternalIdentity.md) |  | [optional] 
**Id** | **string** | identifier of user | 
**LastAuthentication** | **int64** | user last authentication timestamp in ms | [optional] 
**Login** | **string** | user login | 
**MainUser** | **bool** | user is the tenant&#39;s main user | [optional] 
**PortalData** | **map[string]interface{}** | user portal data | [optional] 
**Roles** | **[]string** | list of user associated roles. | 
**State** | **string** | user state : disabled, enabled or suspended | 
**TenantId** | **string** | identifier of tenant account this user will belong to | 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


