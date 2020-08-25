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

> Group CreateGroupUsingPOST1(ctx, xAPIKEY, optional)

Create a group

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateGroupUsingPOST1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateGroupUsingPOST1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **request** | [**optional.Interface of GroupCreateRequest**](GroupCreateRequest.md)| The group to register | 

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

> DeleteGroupUsingDELETE1(ctx, id, xAPIKEY)

Delete a group

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| targeted for deletion group identifier. Expected string (max 6 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> Group GetGroupUsingGET1(ctx, id, xAPIKEY)

Get a group

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| the Group Identifier. Expected string (max 6 characters) | 
**xAPIKEY** | **string**| a valid API key | 

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

> []Group ListGroupsUsingGET1(ctx, xAPIKEY, optional)

List registered groups

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListGroupsUsingGET1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGroupsUsingGET1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **limit** | **optional.Int32**| the maximum number of items per page (optional) | [default to 20]
 **offset** | **optional.Int32**| the requested page number (optional) | [default to 0]
 **parentId** | **optional.String**| filter list by group&#39;s parent. Expected string (max 6 characters) | 
 **xTotalCount** | **optional.Bool**| true if a total count must be returned in response | [default to false]

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

> Group UpdateGroupUsingPUT1(ctx, id, xAPIKEY, optional)

Update a group

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string**| identifier to update group. Expected string (max 6 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateGroupUsingPUT1Opts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateGroupUsingPUT1Opts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**optional.Interface of GroupUpdateRequest**](GroupUpdateRequest.md)| The group to register | 

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

