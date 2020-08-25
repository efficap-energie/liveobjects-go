# \UserAuthenticationApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AuthenticateUserUsingPOST**](UserAuthenticationApi.md#AuthenticateUserUsingPOST) | **Post** /api/v0/auth | Authenticate a user
[**CookiesDeleteUsingDELETE**](UserAuthenticationApi.md#CookiesDeleteUsingDELETE) | **Delete** /api/v0/cookies | cookiesDelete
[**GetTenantIdUsingGET**](UserAuthenticationApi.md#GetTenantIdUsingGET) | **Get** /api/v0/whoami | Get your tenant id
[**LogoutUsingPOST**](UserAuthenticationApi.md#LogoutUsingPOST) | **Post** /api/v0/logout | Log out of the current session
[**ResetUserPasswordUsingPOST**](UserAuthenticationApi.md#ResetUserPasswordUsingPOST) | **Post** /api/v0/resetpwd | Reset user&#39;s password
[**UpdateUserEmailWithTokenUsingPOST**](UserAuthenticationApi.md#UpdateUserEmailWithTokenUsingPOST) | **Post** /api/v0/updateEmail | Update user&#39;s email
[**UpdateUserPasswordWithTokenUsingPOST**](UserAuthenticationApi.md#UpdateUserPasswordWithTokenUsingPOST) | **Post** /api/v0/setpwd | Update user&#39;s password



## AuthenticateUserUsingPOST

> AuthResWeb AuthenticateUserUsingPOST(ctx, optional)

Authenticate a user

Usage of this API will be reported in your access log under 'authentication' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***AuthenticateUserUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AuthenticateUserUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **optional.Bool**| if true, send the API key value in a secure cookie | [default to false]
 **authenticationRequest** | [**optional.Interface of AuthReqWeb**](AuthReqWeb.md)| body of authentication request | 

### Return type

[**AuthResWeb**](AuthResWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CookiesDeleteUsingDELETE

> CookiesDeleteUsingDELETE(ctx, xAPIKEY)

cookiesDelete

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetTenantIdUsingGET

> SimpleStringWeb GetTenantIdUsingGET(ctx, xAPIKEY)

Get your tenant id

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**SimpleStringWeb**](SimpleStringWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## LogoutUsingPOST

> LogoutUsingPOST(ctx, xAPIKEY, optional)

Log out of the current session

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***LogoutUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a LogoutUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **cookie** | **optional.Bool**| if true, send the API key value in a secure cookie | [default to false]

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ResetUserPasswordUsingPOST

> ResetUserPasswordUsingPOST(ctx, optional)

Reset user's password

Usage of this API will be reported in your access log under 'authentication' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***ResetUserPasswordUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ResetUserPasswordUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userResetPasswordRequest** | [**optional.Interface of ResetPasswordReqWeb**](ResetPasswordReqWeb.md)| body of user reset password  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserEmailWithTokenUsingPOST

> UpdateUserEmailWithTokenUsingPOST(ctx, xAPIKEY, optional)

Update user's email

Usage of this API will be reported in your access log under 'authentication' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateUserEmailWithTokenUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserEmailWithTokenUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **userUpdateEmailRequest** | [**optional.Interface of UpdateEmailReqWeb**](UpdateEmailReqWeb.md)| body of update User Email | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateUserPasswordWithTokenUsingPOST

> UpdateUserPasswordWithTokenUsingPOST(ctx, optional)

Update user's password

Usage of this API will be reported in your access log under 'authentication' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
 **optional** | ***UpdateUserPasswordWithTokenUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateUserPasswordWithTokenUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUpdatePasswordRequest** | [**optional.Interface of UpdatePasswordWithTokenReqWeb**](UpdatePasswordWithTokenReqWeb.md)| body of user update password with token | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

