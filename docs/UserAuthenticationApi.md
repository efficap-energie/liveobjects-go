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

> AuthResWeb AuthenticateUserUsingPOST(ctx).Cookie(cookie).AuthenticationRequest(authenticationRequest).Execute()

Authenticate a user



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    cookie := true // bool | if true, send the API key value in a secure cookie (optional) (default to false)
    authenticationRequest := *openapiclient.NewAuthReqWeb("Password_example") // AuthReqWeb | body of authentication request (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAuthenticationApi.AuthenticateUserUsingPOST(context.Background()).Cookie(cookie).AuthenticationRequest(authenticationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticationApi.AuthenticateUserUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AuthenticateUserUsingPOST`: AuthResWeb
    fmt.Fprintf(os.Stdout, "Response from `UserAuthenticationApi.AuthenticateUserUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAuthenticateUserUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **cookie** | **bool** | if true, send the API key value in a secure cookie | [default to false]
 **authenticationRequest** | [**AuthReqWeb**](AuthReqWeb.md) | body of authentication request | 

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

> CookiesDeleteUsingDELETE(ctx).XAPIKEY(xAPIKEY).Execute()

cookiesDelete

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAuthenticationApi.CookiesDeleteUsingDELETE(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticationApi.CookiesDeleteUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCookiesDeleteUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

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

> SimpleStringWeb GetTenantIdUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

Get your tenant id

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAuthenticationApi.GetTenantIdUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticationApi.GetTenantIdUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetTenantIdUsingGET`: SimpleStringWeb
    fmt.Fprintf(os.Stdout, "Response from `UserAuthenticationApi.GetTenantIdUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantIdUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

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

> LogoutUsingPOST(ctx).XAPIKEY(xAPIKEY).Cookie(cookie).Execute()

Log out of the current session

### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    cookie := true // bool | if true, send the API key value in a secure cookie (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAuthenticationApi.LogoutUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Cookie(cookie).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticationApi.LogoutUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiLogoutUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **cookie** | **bool** | if true, send the API key value in a secure cookie | [default to false]

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

> ResetUserPasswordUsingPOST(ctx).UserResetPasswordRequest(userResetPasswordRequest).Execute()

Reset user's password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userResetPasswordRequest := *openapiclient.NewResetPasswordReqWeb("Captcha_example", "CaptchaToken_example", "UserLogin_example") // ResetPasswordReqWeb | body of user reset password  (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAuthenticationApi.ResetUserPasswordUsingPOST(context.Background()).UserResetPasswordRequest(userResetPasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticationApi.ResetUserPasswordUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiResetUserPasswordUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userResetPasswordRequest** | [**ResetPasswordReqWeb**](ResetPasswordReqWeb.md) | body of user reset password  | 

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

> UpdateUserEmailWithTokenUsingPOST(ctx).XAPIKEY(xAPIKEY).UserUpdateEmailRequest(userUpdateEmailRequest).Execute()

Update user's email



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    userUpdateEmailRequest := *openapiclient.NewUpdateEmailReqWeb("Captcha_example", "CaptchaToken_example", "TokenId_example") // UpdateEmailReqWeb | body of update User Email (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAuthenticationApi.UpdateUserEmailWithTokenUsingPOST(context.Background()).XAPIKEY(xAPIKEY).UserUpdateEmailRequest(userUpdateEmailRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticationApi.UpdateUserEmailWithTokenUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserEmailWithTokenUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **userUpdateEmailRequest** | [**UpdateEmailReqWeb**](UpdateEmailReqWeb.md) | body of update User Email | 

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

> UpdateUserPasswordWithTokenUsingPOST(ctx).UserUpdatePasswordRequest(userUpdatePasswordRequest).Execute()

Update user's password



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {
    userUpdatePasswordRequest := *openapiclient.NewUpdatePasswordWithTokenReqWeb("Captcha_example", "CaptchaToken_example", "Company_example", "FirstName_example", "LastName_example", "Login_example", "Password_example", "TokenId_example") // UpdatePasswordWithTokenReqWeb | body of user update password with token (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UserAuthenticationApi.UpdateUserPasswordWithTokenUsingPOST(context.Background()).UserUpdatePasswordRequest(userUpdatePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserAuthenticationApi.UpdateUserPasswordWithTokenUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordWithTokenUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **userUpdatePasswordRequest** | [**UpdatePasswordWithTokenReqWeb**](UpdatePasswordWithTokenReqWeb.md) | body of user update password with token | 

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

