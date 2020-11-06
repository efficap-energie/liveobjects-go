# \DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupUsingPOST1**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#CreateGroupUsingPOST1) | **Post** /api/v0/groups | Create a group
[**DeleteGroupUsingDELETE1**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#DeleteGroupUsingDELETE1) | **Delete** /api/v0/groups/{groupId} | Delete a group
[**GetGroupUsingGET1**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#GetGroupUsingGET1) | **Get** /api/v0/groups/{groupId} | Get a group
[**ListGroupsUsingGET1**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#ListGroupsUsingGET1) | **Get** /api/v0/groups | List registered groups
[**UpdateGroupUsingPUT1**](DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.md#UpdateGroupUsingPUT1) | **Put** /api/v0/groups/{groupId} | Update a group



## CreateGroupUsingPOST1

> Group CreateGroupUsingPOST1(ctx).XAPIKEY(xAPIKEY).Request(request).Execute()

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
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.CreateGroupUsingPOST1(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.CreateGroupUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupUsingPOST1`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.CreateGroupUsingPOST1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateGroupUsingPOST1Request struct via the builder pattern


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


## DeleteGroupUsingDELETE1

> DeleteGroupUsingDELETE1(ctx, groupId).XAPIKEY(xAPIKEY).Execute()

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
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.DeleteGroupUsingDELETE1(context.Background(), groupId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.DeleteGroupUsingDELETE1``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteGroupUsingDELETE1Request struct via the builder pattern


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


## GetGroupUsingGET1

> Group GetGroupUsingGET1(ctx, groupId).XAPIKEY(xAPIKEY).Execute()

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
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.GetGroupUsingGET1(context.Background(), groupId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.GetGroupUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupUsingGET1`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.GetGroupUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | the Group Identifier. Expected string (max 6 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGroupUsingGET1Request struct via the builder pattern


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


## ListGroupsUsingGET1

> PageableGroup ListGroupsUsingGET1(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Parent(parent).Execute()

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
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.ListGroupsUsingGET1(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Parent(parent).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.ListGroupsUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupsUsingGET1`: PageableGroup
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.ListGroupsUsingGET1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsUsingGET1Request struct via the builder pattern


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


## UpdateGroupUsingPUT1

> Group UpdateGroupUsingPUT1(ctx, groupId).XAPIKEY(xAPIKEY).Body(body).Execute()

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
    resp, r, err := api_client.DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.UpdateGroupUsingPUT1(context.Background(), groupId).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.UpdateGroupUsingPUT1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupUsingPUT1`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedGroupManagementV0UseDeviceManagementGroupsV1InsteadApi.UpdateGroupUsingPUT1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string** | identifier to update group. Expected string (max 6 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGroupUsingPUT1Request struct via the builder pattern


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

