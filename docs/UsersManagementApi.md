# \UsersManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUserUsingPOST3**](UsersManagementApi.md#ActivateUserUsingPOST3) | **Post** /api/v0/users/activateUser | Activate a user in a tenant account
[**CreateUserAccountUsingPOST3**](UsersManagementApi.md#CreateUserAccountUsingPOST3) | **Post** /api/v0/users | Create a user
[**DeleteUserUsingDELETE3**](UsersManagementApi.md#DeleteUserUsingDELETE3) | **Delete** /api/v0/users/{userId} | Delete a user in a tenant account
[**GetCurrentUserUsingGET**](UsersManagementApi.md#GetCurrentUserUsingGET) | **Get** /api/v0/users/me | Get a \&quot;myself\&quot; user data of Tenant
[**GetUserPortalDataUsingGET**](UsersManagementApi.md#GetUserPortalDataUsingGET) | **Get** /api/v0/users/me/portaldata | Get the portal data of me
[**GetUserUsingGET3**](UsersManagementApi.md#GetUserUsingGET3) | **Get** /api/v0/users/{userId} | Get details of a user in a tenant account
[**ListUsersUsingGET1**](UsersManagementApi.md#ListUsersUsingGET1) | **Get** /api/v0/users | List all users in a tenant account
[**UpdateUserPasswordUsingPOST**](UsersManagementApi.md#UpdateUserPasswordUsingPOST) | **Post** /api/v0/users/{userId}/password | Update user password
[**UpdateUserPasswordWithTokenWithoutCaptchaUsingPOST**](UsersManagementApi.md#UpdateUserPasswordWithTokenWithoutCaptchaUsingPOST) | **Post** /api/v0/users/activateTrial | Update user password with a token
[**UpdateUserPortalDataUsingPUT**](UsersManagementApi.md#UpdateUserPortalDataUsingPUT) | **Put** /api/v0/users/me/portaldata | Update the portal data of me
[**UpdateUserUsingPOST1**](UsersManagementApi.md#UpdateUserUsingPOST1) | **Post** /api/v0/users/{userId} | Update a user



## ActivateUserUsingPOST3

> ActivateUserUsingPOST3(ctx, xAPIKEY, optional)

Activate a user in a tenant account

Usage of this API will be reported in your access log under 'user' category.<br><br>Restricted to API keys with at least one of the following roles : USER_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ActivateUserUsingPOST3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ActivateUserUsingPOST3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userActivationRequest** | [**optional.Interface of UserActivationReqWeb**](UserActivationReqWeb.md)| body of user activation | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateUserAccountUsingPOST3

> User CreateUserAccountUsingPOST3(ctx, xAPIKEY, optional)

Create a user

Usage of this API will be reported in your access log under 'user' category.<br><br>Restricted to API keys with at least one of the following roles : USER_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateUserAccountUsingPOST3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUserAccountUsingPOST3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userCreationRequest** | [**optional.Interface of UserCreationReqWeb**](UserCreationReqWeb.md)| body of user creation | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUserUsingDELETE3

> DeleteUserUsingDELETE3(ctx, userId, xAPIKEY, optional)

Delete a user in a tenant account

Usage of this API will be reported in your access log under 'user' category.<br><br>Restricted to API keys with at least one of the following roles : USER_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| the User Identifier. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***DeleteUserUsingDELETE3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteUserUsingDELETE3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **optional.String**| identifier of tenant account ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

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


## GetCurrentUserUsingGET

> User GetCurrentUserUsingGET(ctx, xAPIKEY)

Get a \"myself\" user data of Tenant

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserPortalDataUsingGET

> map[string]interface{} GetUserPortalDataUsingGET(ctx, xAPIKEY)

Get the portal data of me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

**map[string]interface{}**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetUserUsingGET3

> User GetUserUsingGET3(ctx, userId, xAPIKEY, optional)

Get details of a user in a tenant account

Restricted to API keys with at least one of the following roles : USER_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| the User Identifier | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetUserUsingGET3Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetUserUsingGET3Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **tenantId** | **optional.String**| identifier of tenant account ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersUsingGET1

> PageableUser ListUsersUsingGET1(ctx, xAPIKEY, optional)

List all users in a tenant account

Restricted to API keys with at least one of the following roles : USER_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListUsersUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListUsersUsingGET1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]
 **tenantId** | **optional.String**| identifier of tenant account (optional) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot; | 
 **login** | **optional.String**| login of tenant account (optional) ex: \&quot;login_name\&quot; | 
 **email** | **optional.String**| email of tenant account (optional) ex: \&quot;myAccount@mail.com\&quot; | 
 **externalProvider** | **optional.String**| external (IDP) provider of tenant account (optional) ex: \&quot;SIU\&quot; | 
 **externalId** | **optional.String**| external (IDP) identifier of tenant account (optional) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot; | 
 **externalLogin** | **optional.String**| external (IDP) login of tenant account (optional) ex: \&quot;login_name\&quot; | 
 **roles** | [**optional.Interface of []string**](string.md)| roles of tenant account (optional) ex: \&quot;[TENANT_ADMIN, USER_KEY]\&quot; | 
 **rolesNotIn** | [**optional.Interface of []string**](string.md)| roles not in tenant account (optional) ex: \&quot;[LPWA_ORANGE_ADMIN]\&quot; | 
 **userState** | [**optional.Interface of []string**](string.md)| userState of tenant account (optional) ex: \&quot;[enabled]\&quot; | 

### Return type

[**PageableUser**](Pageable«User».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPasswordUsingPOST

> UpdateUserPasswordUsingPOST(ctx, userId, xAPIKEY, optional)

Update user password

Usage of this API will be reported in your access log under 'user' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| the User Identifier. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUserPasswordUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserPasswordUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userUpdatePasswordRequest** | [**optional.Interface of UpdatePasswordReqWeb**](UpdatePasswordReqWeb.md)| body of user update password | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPasswordWithTokenWithoutCaptchaUsingPOST

> UpdateUserPasswordWithTokenWithoutCaptchaUsingPOST(ctx, xAPIKEY, userUpdatePasswordRequest)

Update user password with a token

Usage of this API will be reported in your access log under 'user' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
**userUpdatePasswordRequest** | [**UpdatePasswordWithTokenReqWeb**](UpdatePasswordWithTokenReqWeb.md)| userUpdatePasswordRequest | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPortalDataUsingPUT

> UpdateUserPortalDataUsingPUT(ctx, xAPIKEY, optional)

Update the portal data of me

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUserPortalDataUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserPortalDataUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **portalData** | **optional.Map[string]interface{}**| body of my user portal data Update | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserUsingPOST1

> User UpdateUserUsingPOST1(ctx, userId, xAPIKEY, optional)

Update a user

Usage of this API will be reported in your access log under 'user' category.<br><br>Restricted to API keys with at least one of the following roles : USER_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string**| the User Identifier. Expected identifier (max 24 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUserUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **userUpdateRequest** | [**optional.Interface of UserUpdateReqWeb**](UserUpdateReqWeb.md)| body of user update | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

