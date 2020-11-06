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

> ApiKey CreateApiKeyUsingPOST(ctx).XAPIKEY(xAPIKEY).ApiKeyCreationRequest(apiKeyCreationRequest).Execute()

Create an API key



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
    apiKeyCreationRequest := *openapiclient.NewApiKeyCreationReqWeb("ParentId_example", []string{"Roles_example")) // ApiKeyCreationReqWeb | body for create API key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysApi.CreateApiKeyUsingPOST(context.Background()).XAPIKEY(xAPIKEY).ApiKeyCreationRequest(apiKeyCreationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.CreateApiKeyUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateApiKeyUsingPOST`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.CreateApiKeyUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApiKeyUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **apiKeyCreationRequest** | [**ApiKeyCreationReqWeb**](ApiKeyCreationReqWeb.md) | body for create API key | 

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

> DeleteApiKeyUsingDELETE(ctx, apiKeyId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()

Delete an API key



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
    apiKeyId := "apiKeyId_example" // string | identifier of your API key. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    tenantId := "tenantId_example" // string | identifier of tenant account (optionnal) ex: \"57xxxxxxxxxxxxxxxxxxxxxx\". Expected identifier (max 24 characters) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysApi.DeleteApiKeyUsingDELETE(context.Background(), apiKeyId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.DeleteApiKeyUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | identifier of your API key. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteApiKeyUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **tenantId** | **string** | identifier of tenant account (optionnal) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

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

> ApiKey GetApiKeyFromAuthenticationUsingGET3(ctx).XAPIKEY(xAPIKEY).Execute()

getApiKeyFromAuthentication

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
    resp, r, err := api_client.ApiKeysApi.GetApiKeyFromAuthenticationUsingGET3(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.GetApiKeyFromAuthenticationUsingGET3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeyFromAuthenticationUsingGET3`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.GetApiKeyFromAuthenticationUsingGET3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyFromAuthenticationUsingGET3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

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

> ApiKey GetApiKeyUsingGET3(ctx, apiKeyId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()

Get an API key



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
    apiKeyId := "apiKeyId_example" // string | the id of your API key. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    tenantId := "tenantId_example" // string | identifier of tenant account (optionnal) ex: \"57xxxxxxxxxxxxxxxxxxxxxx\". Expected identifier (max 24 characters) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysApi.GetApiKeyUsingGET3(context.Background(), apiKeyId).XAPIKEY(xAPIKEY).TenantId(tenantId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.GetApiKeyUsingGET3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeyUsingGET3`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.GetApiKeyUsingGET3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | the id of your API key. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeyUsingGET3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **tenantId** | **string** | identifier of tenant account (optionnal) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 

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

> PageableApiKey GetApiKeysUsingGET3(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).TenantId(tenantId).ParentId(parentId).ShowSessionKeys(showSessionKeys).Roles(roles).ShowMasterKey(showMasterKey).Execute()

List API keys



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
    size := 987 // int32 | the maximum number of items per page (optional) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)
    tenantId := "tenantId_example" // string | identifier of tenant account (optionnal) ex: \"57xxxxxxxxxxxxxxxxxxxxxx\". Expected identifier (max 24 characters) (optional)
    parentId := "parentId_example" // string | the id of your parent (optional)  ex: \"57xxxxxxxxxxxxxxxxxxxxxx\". Expected identifier (max 24 characters) (optional)
    showSessionKeys := true // bool | include the session Keys (optional) (optional) (default to false)
    roles := []string{"Inner_example"} // []string | list of API Key associated roles (optional). Basic roles are \"USER_R\", \"USER_W\", \"API_KEY_R\", \"API_KEY_W\" or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) (optional)
    showMasterKey := true // bool | Boolean to show or not the master api key (optional) (default to true)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysApi.GetApiKeysUsingGET3(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).TenantId(tenantId).ParentId(parentId).ShowSessionKeys(showSessionKeys).Roles(roles).ShowMasterKey(showMasterKey).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.GetApiKeysUsingGET3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetApiKeysUsingGET3`: PageableApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.GetApiKeysUsingGET3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApiKeysUsingGET3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]
 **tenantId** | **string** | identifier of tenant account (optionnal) ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 
 **parentId** | **string** | the id of your parent (optional)  ex: \&quot;57xxxxxxxxxxxxxxxxxxxxxx\&quot;. Expected identifier (max 24 characters) | 
 **showSessionKeys** | **bool** | include the session Keys (optional) | [default to false]
 **roles** | [**[]string**](string.md) | list of API Key associated roles (optional). Basic roles are \&quot;USER_R\&quot;, \&quot;USER_W\&quot;, \&quot;API_KEY_R\&quot;, \&quot;API_KEY_W\&quot; or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters) | 
 **showMasterKey** | **bool** | Boolean to show or not the master api key | [default to true]

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

> ApiKey SetApiKeyDebugModeUsingPUT3(ctx, apiKeyId).XAPIKEY(xAPIKEY).SetApiKeyDebugModeRequest(setApiKeyDebugModeRequest).Execute()

Activate/Deactivate the debug mode on an API key



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
    apiKeyId := "apiKeyId_example" // string | identifier of your API key. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    setApiKeyDebugModeRequest := *openapiclient.NewApiKeySetDebugModeReqWeb(false) // ApiKeySetDebugModeReqWeb | body for API key debug mode (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysApi.SetApiKeyDebugModeUsingPUT3(context.Background(), apiKeyId).XAPIKEY(xAPIKEY).SetApiKeyDebugModeRequest(setApiKeyDebugModeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.SetApiKeyDebugModeUsingPUT3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SetApiKeyDebugModeUsingPUT3`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.SetApiKeyDebugModeUsingPUT3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | identifier of your API key. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetApiKeyDebugModeUsingPUT3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **setApiKeyDebugModeRequest** | [**ApiKeySetDebugModeReqWeb**](ApiKeySetDebugModeReqWeb.md) | body for API key debug mode | 

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

> ApiKey UpdateApiKeyUsingPOST1(ctx, apiKeyId).XAPIKEY(xAPIKEY).ApiKeyUpdateRequest(apiKeyUpdateRequest).Execute()

Update an API key



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
    apiKeyId := "apiKeyId_example" // string | the id of your API key. Expected identifier (max 24 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    apiKeyUpdateRequest := *openapiclient.NewApiKeyUpdateReqWeb() // ApiKeyUpdateReqWeb | body for update API key (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.ApiKeysApi.UpdateApiKeyUsingPOST1(context.Background(), apiKeyId).XAPIKEY(xAPIKEY).ApiKeyUpdateRequest(apiKeyUpdateRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ApiKeysApi.UpdateApiKeyUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateApiKeyUsingPOST1`: ApiKey
    fmt.Fprintf(os.Stdout, "Response from `ApiKeysApi.UpdateApiKeyUsingPOST1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**apiKeyId** | **string** | the id of your API key. Expected identifier (max 24 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateApiKeyUsingPOST1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **apiKeyUpdateRequest** | [**ApiKeyUpdateReqWeb**](ApiKeyUpdateReqWeb.md) | body for update API key | 

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

