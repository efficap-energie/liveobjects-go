# \DeviceManagementGroupsV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateGroupUsingPOST1**](DeviceManagementGroupsV1Api.md#CreateGroupUsingPOST1) | **Post** /api/v1/deviceMgt/groups | Create a group
[**DeleteGroupUsingDELETE1**](DeviceManagementGroupsV1Api.md#DeleteGroupUsingDELETE1) | **Delete** /api/v1/deviceMgt/groups/{id} | Delete a group
[**GetGroupUsingGET1**](DeviceManagementGroupsV1Api.md#GetGroupUsingGET1) | **Get** /api/v1/deviceMgt/groups/{id} | Get a group
[**ListGroupsUsingGET1**](DeviceManagementGroupsV1Api.md#ListGroupsUsingGET1) | **Get** /api/v1/deviceMgt/groups | List registered groups
[**UpdateGroupUsingPUT1**](DeviceManagementGroupsV1Api.md#UpdateGroupUsingPUT1) | **Put** /api/v1/deviceMgt/groups/{id} | Update a group



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
    resp, r, err := api_client.DeviceManagementGroupsV1Api.CreateGroupUsingPOST1(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementGroupsV1Api.CreateGroupUsingPOST1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateGroupUsingPOST1`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementGroupsV1Api.CreateGroupUsingPOST1`: %v\n", resp)
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

> DeleteGroupUsingDELETE1(ctx, id).XAPIKEY(xAPIKEY).Execute()

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
    id := "id_example" // string | targeted for deletion group identifier. Expected string (max 6 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementGroupsV1Api.DeleteGroupUsingDELETE1(context.Background(), id).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementGroupsV1Api.DeleteGroupUsingDELETE1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | targeted for deletion group identifier. Expected string (max 6 characters) | 

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

> Group GetGroupUsingGET1(ctx, id).XAPIKEY(xAPIKEY).Execute()

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
    id := "id_example" // string | the Group Identifier. Expected string (max 6 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementGroupsV1Api.GetGroupUsingGET1(context.Background(), id).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementGroupsV1Api.GetGroupUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGroupUsingGET1`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementGroupsV1Api.GetGroupUsingGET1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | the Group Identifier. Expected string (max 6 characters) | 

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

> []Group ListGroupsUsingGET1(ctx).XAPIKEY(xAPIKEY).Limit(limit).Offset(offset).ParentId(parentId).XTotalCount(xTotalCount).Execute()

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
    limit := 987 // int32 | the maximum number of items per page (optional) (optional) (default to 20)
    offset := 987 // int32 | the requested page number (optional) (optional) (default to 0)
    parentId := "parentId_example" // string | filter list by group's parent. Expected string (max 6 characters) (optional)
    xTotalCount := true // bool | true if a total count must be returned in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementGroupsV1Api.ListGroupsUsingGET1(context.Background()).XAPIKEY(xAPIKEY).Limit(limit).Offset(offset).ParentId(parentId).XTotalCount(xTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementGroupsV1Api.ListGroupsUsingGET1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListGroupsUsingGET1`: []Group
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementGroupsV1Api.ListGroupsUsingGET1`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListGroupsUsingGET1Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **limit** | **int32** | the maximum number of items per page (optional) | [default to 20]
 **offset** | **int32** | the requested page number (optional) | [default to 0]
 **parentId** | **string** | filter list by group&#39;s parent. Expected string (max 6 characters) | 
 **xTotalCount** | **bool** | true if a total count must be returned in response | [default to false]

### Return type

[**[]Group**](Group.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateGroupUsingPUT1

> Group UpdateGroupUsingPUT1(ctx, id).XAPIKEY(xAPIKEY).Body(body).Execute()

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
    id := "id_example" // string | identifier to update group. Expected string (max 6 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    body := *openapiclient.NewGroupUpdateRequest("PathNode_example") // GroupUpdateRequest | The group to register (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementGroupsV1Api.UpdateGroupUsingPUT1(context.Background(), id).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementGroupsV1Api.UpdateGroupUsingPUT1``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGroupUsingPUT1`: Group
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementGroupsV1Api.UpdateGroupUsingPUT1`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | identifier to update group. Expected string (max 6 characters) | 

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

