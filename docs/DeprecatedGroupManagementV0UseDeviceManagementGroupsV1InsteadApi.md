# \DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupUsingPOST**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#CreateGroupUsingPOST) | **Post** /api/v0/groups | Create a group
[**DeleteGroupUsingDELETE**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#DeleteGroupUsingDELETE) | **Delete** /api/v0/groups/{groupId} | Delete a group
[**GetGroupUsingGET**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#GetGroupUsingGET) | **Get** /api/v0/groups/{groupId} | Get a group
[**ListGroupsUsingGET**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#ListGroupsUsingGET) | **Get** /api/v0/groups | List registered groups
[**UpdateGroupUsingPUT**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#UpdateGroupUsingPUT) | **Put** /api/v0/groups/{groupId} | Update a group



## CreateGroupUsingPOST

> Group CreateGroupUsingPOST(ctx).XAPIKEY(xAPIKEY).Request(request).Execute()

Create a group



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
    request := *openapiclient.NewGroupCreateRequest("PathNode_example") // GroupCreateRequest | The group to register (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.CreateGroupUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.CreateGroupUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupUsingPOST`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.CreateGroupUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**GroupCreateRequest**](GroupCreateRequest.md) | The group to register | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteGroupUsingDELETE

> DeleteGroupUsingDELETE(ctx, groupId).XAPIKEY(xAPIKEY).Execute()

Delete a group



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
    groupId := "groupId_example" // string | targeted for deletion group identifier. Expected string (max 6 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.DeleteGroupUsingDELETE(context.Background(), groupId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.DeleteGroupUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | targeted for deletion group identifier. Expected string (max 6 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteGroupUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

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


## GetGroupUsingGET

> Group GetGroupUsingGET(ctx, groupId).XAPIKEY(xAPIKEY).Execute()

Get a group



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
    groupId := "groupId_example" // string | the Group Identifier. Expected string (max 6 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.GetGroupUsingGET(context.Background(), groupId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.GetGroupUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupUsingGET`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.GetGroupUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | the Group Identifier. Expected string (max 6 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListGroupsUsingGET

> PageableGroup ListGroupsUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Parent(parent).Execute()

List registered groups



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
    size := 987 // int64 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int64 | the requested page number (optional) (optional) (default to 0)
    parent := "parent_example" // string | filter list by group's parent. Expected string (max 6 characters) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.ListGroupsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Parent(parent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.ListGroupsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupsUsingGET`: PageableGroup
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.ListGroupsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int64** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int64** | the requested page number (optional) | [default to 0]
 **parent** | **string** | filter list by group&#39;s parent. Expected string (max 6 characters) | 

### Return type

[**PageableGroup**](Pageable«Group».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupUsingPUT

> Group UpdateGroupUsingPUT(ctx, groupId).XAPIKEY(xAPIKEY).Body(body).Execute()

Update a group



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
    groupId := "groupId_example" // string | identifier to update group. Expected string (max 6 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    body := *openapiclient.NewGroupUpdateRequest("PathNode_example") // GroupUpdateRequest | The group to register (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.UpdateGroupUsingPUT(context.Background(), groupId).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.UpdateGroupUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupUsingPUT`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.UpdateGroupUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | identifier to update group. Expected string (max 6 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **body** | [**GroupUpdateRequest**](GroupUpdateRequest.md) | The group to register | 

### Return type

[**Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

