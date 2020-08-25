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

> Group CreateGroupUsingPOST(ctx, xAPIKEY, optional)

Create a group

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateGroupUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateGroupUsingPOSTOpts struct


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


## DeleteGroupUsingDELETE

> DeleteGroupUsingDELETE(ctx, groupId, xAPIKEY)

Delete a group

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| targeted for deletion group identifier. Expected string (max 6 characters) | 
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


## GetGroupUsingGET

> Group GetGroupUsingGET(ctx, groupId, xAPIKEY)

Get a group

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| the Group Identifier. Expected string (max 6 characters) | 
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


## ListGroupsUsingGET

> PageableGroup ListGroupsUsingGET(ctx, xAPIKEY, optional)

List registered groups

Restricted to API keys with at least one of the following roles : DEVICE_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListGroupsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListGroupsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int64**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int64**| the requested page number (optional) | [default to 0]
 **parent** | **optional.String**| filter list by group&#39;s parent. Expected string (max 6 characters) | 

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

> Group UpdateGroupUsingPUT(ctx, groupId, xAPIKEY, optional)

Update a group

Usage of this API will be reported in your access log under 'device_inventory' category.<br><br>Restricted to API keys with at least one of the following roles : DEVICE_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupId** | **string**| identifier to update group. Expected string (max 6 characters) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***UpdateGroupUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a UpdateGroupUsingPUTOpts struct


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

