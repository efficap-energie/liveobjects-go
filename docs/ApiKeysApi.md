# \ApiKeysApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApiKeyUsingPOST**](ApiKeysApi.md#CreateApiKeyUsingPOST) | **Post** /api/v0/apiKeys | Create an API key
[**DeleteApiKeyUsingDELETE**](ApiKeysApi.md#DeleteApiKeyUsingDELETE) | **Delete** /api/v0/apiKeys/{apiKeyId} | Delete an API key
[**GetApiKeyFromAuthenticationUsingGET3**](ApiKeysApi.md#GetApiKeyFromAuthenticationUsingGET3) | **Get** /api/v0/apiKeys/current_key | getApiKeyFromAuthentication
[**GetApiKeyUsingGET3**](ApiKeysApi.md#GetApiKeyUsingGET3) | **Get** /api/v0/apiKeys/{apiKeyId} | Get an API key
[**GetApiKeysUsingGET3**](ApiKeysApi.md#GetApiKeysUsingGET3) | **Get** /api/v0/apiKeys | List API keys
[**SetApiKeyDebugModeUsingPUT3**](ApiKeysApi.md#SetApiKeyDebugModeUsingPUT3) | **Put** /api/v0/apiKeys/{apiKeyId}/debugMode | Activate/Deactivate the debug mode on an API key
[**UpdateApiKeyUsingPOST1**](ApiKeysApi.md#UpdateApiKeyUsingPOST1) | **Post** /api/v0/apiKeys/{apiKeyId} | Update an API key



## CreateApiKeyUsingPOST

> ApiKey CreateApiKeyUsingPOST(ctx, xAPIKEY, optional)

Create an API key

Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateApiKeyUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateApiKeyUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **apiKeyCreationRequest** | [**optional.Interface of ApiKeyCreationReqWeb**](ApiKeyCreationReqWeb.md)| body for create API key | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteApiKeyUsingDELETE

> DeleteApiKeyUsingDELETE(ctx, apiKeyId, xAPIKEY, optional)

Delete an API key

Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string**| identifier of your API key. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***DeleteApiKeyUsingDELETEOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteApiKeyUsingDELETEOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **optional.String**| identifier of tenant account (optionnal) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyFromAuthenticationUsingGET3

> ApiKey GetApiKeyFromAuthenticationUsingGET3(ctx, xAPIKEY)

getApiKeyFromAuthentication

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeyUsingGET3

> ApiKey GetApiKeyUsingGET3(ctx, apiKeyId, xAPIKEY, optional)

Get an API key

Restricted to API keys with at least one of the following roles : API_KEY_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string**| the id of your API key. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetApiKeyUsingGET3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApiKeyUsingGET3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **optional.String**| identifier of tenant account (optionnal) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApiKeysUsingGET3

> PageableApiKey GetApiKeysUsingGET3(ctx, xAPIKEY, optional)

List API keys

Restricted to API keys with at least one of the following roles : API_KEY_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetApiKeysUsingGET3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetApiKeysUsingGET3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]
 **tenantId** | **optional.String**| identifier of tenant account (optionnal) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 
 **parentId** | **optional.String**| the id of your parent (optional)  ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 
 **showSessionKeys** | **optional.Bool**| include the session Keys (optional) | [default to false]
 **roles** | [**optional.Interface of []string**](string.md)| list of API Key associated roles (optional). Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | 
 **showMasterKey** | **optional.Bool**| Boolean to show or not the master api key | [default to true]

### Return type

[**PageableApiKey**](Pageable«ApiKey».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetApiKeyDebugModeUsingPUT3

> ApiKey SetApiKeyDebugModeUsingPUT3(ctx, apiKeyId, xAPIKEY, optional)

Activate/Deactivate the debug mode on an API key

Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string**| identifier of your API key. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***SetApiKeyDebugModeUsingPUT3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a SetApiKeyDebugModeUsingPUT3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **setApiKeyDebugModeRequest** | [**optional.Interface of ApiKeySetDebugModeReqWeb**](ApiKeySetDebugModeReqWeb.md)| body for API key debug mode | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateApiKeyUsingPOST1

> ApiKey UpdateApiKeyUsingPOST1(ctx, apiKeyId, xAPIKEY, optional)

Update an API key

Update a set of properties of the selected API key<br><br>Usage of this API will be reported in your access log under 'api_key' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string**| the id of your API key. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateApiKeyUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateApiKeyUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **apiKeyUpdateRequest** | [**optional.Interface of ApiKeyUpdateReqWeb**](ApiKeyUpdateReqWeb.md)| body for update API key | 

### Return type

[**ApiKey**](ApiKey.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

