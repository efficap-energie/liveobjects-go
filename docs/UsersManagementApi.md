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

> ActivateUserUsingPOST3(ctx).XAPIKEY(xAPIKEY).UserActivationRequest(userActivationRequest).Execute()

Activate a user in a tenant account



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
    userActivationRequest := *openapiclient.NewUserActivationReqWeb("TenantId_example", "UserId_example") // UserActivationReqWeb | body of user activation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.ActivateUserUsingPOST3(context.Background()).XAPIKEY(xAPIKEY).UserActivationRequest(userActivationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.ActivateUserUsingPOST3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserUsingPOST3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **userActivationRequest** | [**UserActivationReqWeb**](UserActivationReqWeb.md) | body of user activation | 

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

> User CreateUserAccountUsingPOST3(ctx).XAPIKEY(xAPIKEY).UserCreationRequest(userCreationRequest).Execute()

Create a user



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
    userCreationRequest := *openapiclient.NewUserCreationReqWeb("Email_example", "Login_example", []string{"Roles_example"), "TenantId_example") // UserCreationReqWeb | body of user creation (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.CreateUserAccountUsingPOST3(context.Background()).XAPIKEY(xAPIKEY).UserCreationRequest(userCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.CreateUserAccountUsingPOST3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUserAccountUsingPOST3`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersManagementApi.CreateUserAccountUsingPOST3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserAccountUsingPOST3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **userCreationRequest** | [**UserCreationReqWeb**](UserCreationReqWeb.md) | body of user creation | 

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

> DeleteUserUsingDELETE3(ctx, userId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()

Delete a user in a tenant account



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
    userId := "userId_example" // string | the User Identifier. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    tenantId := "tenantId_example" // string | identifier of tenant account ex: \"57xxxxxxxxxxxxxxxxxxxxxx\". Expected identifier (max 24 characters) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.DeleteUserUsingDELETE3(context.Background(), userId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.DeleteUserUsingDELETE3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the User Identifier. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserUsingDELETE3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **tenantId** | **string** | identifier of tenant account ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

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

> User GetCurrentUserUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

Get a \"myself\" user data of Tenant

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
    resp, r, err := api_client.UsersManagementApi.GetCurrentUserUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.GetCurrentUserUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUserUsingGET`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersManagementApi.GetCurrentUserUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

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

> map[string]interface{} GetUserPortalDataUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

Get the portal data of me

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
    resp, r, err := api_client.UsersManagementApi.GetUserPortalDataUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.GetUserPortalDataUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserPortalDataUsingGET`: map[string]interface{}
    fmt.Fprintf(os.Stdout, "Response from `UsersManagementApi.GetUserPortalDataUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPortalDataUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

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

> User GetUserUsingGET3(ctx, userId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()

Get details of a user in a tenant account



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
    userId := "userId_example" // string | the User Identifier
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    tenantId := "tenantId_example" // string | identifier of tenant account ex: \"57xxxxxxxxxxxxxxxxxxxxxx\". Expected identifier (max 24 characters) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.GetUserUsingGET3(context.Background(), userId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.GetUserUsingGET3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUserUsingGET3`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersManagementApi.GetUserUsingGET3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the User Identifier | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserUsingGET3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **tenantId** | **string** | identifier of tenant account ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

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

> PageableUser ListUsersUsingGET1(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).TenantId(tenantId).Login(login).Email(email).ExternalProvider(externalProvider).ExternalId(externalId).ExternalLogin(externalLogin).Roles(roles).RolesNotIn(rolesNotIn).UserState(userState).Execute()

List all users in a tenant account



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
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)
    tenantId := "tenantId_example" // string | identifier of tenant account (optional) ex: \"57xxxxxxxxxxxxxxxxxxxxxx\" (optional)
    login := "login_example" // string | login of tenant account (optional) ex: \"login_name\" (optional)
    email := "email_example" // string | email of tenant account (optional) ex: \"myAccount@mail.com\" (optional)
    externalProvider := "externalProvider_example" // string | external (IDP) provider of tenant account (optional) ex: \"SIU\" (optional)
    externalId := "externalId_example" // string | external (IDP) identifier of tenant account (optional) ex: \"57xxxxxxxxxxxxxxxxxxxxxx\" (optional)
    externalLogin := "externalLogin_example" // string | external (IDP) login of tenant account (optional) ex: \"login_name\" (optional)
    roles := []string{"Inner_example"} // []string | roles of tenant account (optional) ex: \"[TENANT_ADMIN, USER_KEY]\" (optional)
    rolesNotIn := []string{"Inner_example"} // []string | roles not in tenant account (optional) ex: \"[LPWA_ORANGE_ADMIN]\" (optional)
    userState := []string{"Inner_example"} // []string | userState of tenant account (optional) ex: \"[enabled]\" (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.ListUsersUsingGET1(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).TenantId(tenantId).Login(login).Email(email).ExternalProvider(externalProvider).ExternalId(externalId).ExternalLogin(externalLogin).Roles(roles).RolesNotIn(rolesNotIn).UserState(userState).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.ListUsersUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsersUsingGET1`: PageableUser
    fmt.Fprintf(os.Stdout, "Response from `UsersManagementApi.ListUsersUsingGET1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsersUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]
 **tenantId** | **string** | identifier of tenant account (optional) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot; | 
 **login** | **string** | login of tenant account (optional) ex: \&quot;login_name\&quot; | 
 **email** | **string** | email of tenant account (optional) ex: \&quot;myAccount@mail.com\&quot; | 
 **externalProvider** | **string** | external (IDP) provider of tenant account (optional) ex: \&quot;SIU\&quot; | 
 **externalId** | **string** | external (IDP) identifier of tenant account (optional) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot; | 
 **externalLogin** | **string** | external (IDP) login of tenant account (optional) ex: \&quot;login_name\&quot; | 
 **roles** | [**[]string**](string.md) | roles of tenant account (optional) ex: \&quot;[TENANT_ADMIN, USER_KEY]\&quot; | 
 **rolesNotIn** | [**[]string**](string.md) | roles not in tenant account (optional) ex: \&quot;[LPWA_ORANGE_ADMIN]\&quot; | 
 **userState** | [**[]string**](string.md) | userState of tenant account (optional) ex: \&quot;[enabled]\&quot; | 

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

> UpdateUserPasswordUsingPOST(ctx, userId).XAPIKEY(xAPIKEY).UserUpdatePasswordRequest(userUpdatePasswordRequest).Execute()

Update user password



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
    userId := "userId_example" // string | the User Identifier. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    userUpdatePasswordRequest := *openapiclient.NewUpdatePasswordReqWeb("Captcha_example", "CaptchaToken_example", "NewPassword_example", "OldPassword_example", "TokenId_example") // UpdatePasswordReqWeb | body of user update password (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.UpdateUserPasswordUsingPOST(context.Background(), userId).XAPIKEY(xAPIKEY).UserUpdatePasswordRequest(userUpdatePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.UpdateUserPasswordUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the User Identifier. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **userUpdatePasswordRequest** | [**UpdatePasswordReqWeb**](UpdatePasswordReqWeb.md) | body of user update password | 

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

> UpdateUserPasswordWithTokenWithoutCaptchaUsingPOST(ctx).XAPIKEY(xAPIKEY).UserUpdatePasswordRequest(userUpdatePasswordRequest).Execute()

Update user password with a token



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
    userUpdatePasswordRequest := *openapiclient.NewUpdatePasswordWithTokenReqWeb("Captcha_example", "CaptchaToken_example", "Company_example", "FirstName_example", "LastName_example", "Login_example", "Password_example", "TokenId_example") // UpdatePasswordWithTokenReqWeb | userUpdatePasswordRequest

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.UpdateUserPasswordWithTokenWithoutCaptchaUsingPOST(context.Background()).XAPIKEY(xAPIKEY).UserUpdatePasswordRequest(userUpdatePasswordRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.UpdateUserPasswordWithTokenWithoutCaptchaUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPasswordWithTokenWithoutCaptchaUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **userUpdatePasswordRequest** | [**UpdatePasswordWithTokenReqWeb**](UpdatePasswordWithTokenReqWeb.md) | userUpdatePasswordRequest | 

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

> UpdateUserPortalDataUsingPUT(ctx).XAPIKEY(xAPIKEY).PortalData(portalData).Execute()

Update the portal data of me

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
    portalData := 987 // map[string]interface{} | body of my user portal data Update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.UpdateUserPortalDataUsingPUT(context.Background()).XAPIKEY(xAPIKEY).PortalData(portalData).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.UpdateUserPortalDataUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPortalDataUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **portalData** | **map[string]interface{}** | body of my user portal data Update | 

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

> User UpdateUserUsingPOST1(ctx, userId).XAPIKEY(xAPIKEY).UserUpdateRequest(userUpdateRequest).Execute()

Update a user



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
    userId := "userId_example" // string | the User Identifier. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    userUpdateRequest := *openapiclient.NewUserUpdateReqWeb("State_example") // UserUpdateReqWeb | body of user update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.UsersManagementApi.UpdateUserUsingPOST1(context.Background(), userId).XAPIKEY(xAPIKEY).UserUpdateRequest(userUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UsersManagementApi.UpdateUserUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateUserUsingPOST1`: User
    fmt.Fprintf(os.Stdout, "Response from `UsersManagementApi.UpdateUserUsingPOST1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the User Identifier. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **userUpdateRequest** | [**UserUpdateReqWeb**](UserUpdateReqWeb.md) | body of user update | 

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

