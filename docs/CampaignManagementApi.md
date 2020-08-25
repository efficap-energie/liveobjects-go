# \CampaignManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CancelCampaignUsingPUT**](CampaignManagementApi.md#CancelCampaignUsingPUT) | **Put** /api/v0/deviceMgt/campaigns/{campaignId}/cancel | Cancel a campaign
[**CreateCampaignUsingPOST**](CampaignManagementApi.md#CreateCampaignUsingPOST) | **Post** /api/v0/deviceMgt/campaigns | Create a campaign
[**DeleteCampaignUsingDELETE**](CampaignManagementApi.md#DeleteCampaignUsingDELETE) | **Delete** /api/v0/deviceMgt/campaigns/{campaignId} | Delete a campaign
[**GetCampaignDetailsUsingGET**](CampaignManagementApi.md#GetCampaignDetailsUsingGET) | **Get** /api/v0/deviceMgt/campaigns/{campaignId}/targets | Get the campaign status per device
[**GetCampaignUsingGET**](CampaignManagementApi.md#GetCampaignUsingGET) | **Get** /api/v0/deviceMgt/campaigns/{campaignId} | Get a campaign
[**ListCampaignsUsingGET**](CampaignManagementApi.md#ListCampaignsUsingGET) | **Get** /api/v0/deviceMgt/campaigns | List all campaigns
[**UpdateCampaignUsingPATCH**](CampaignManagementApi.md#UpdateCampaignUsingPATCH) | **Patch** /api/v0/deviceMgt/campaigns/{campaignId} | Update a campaign



## CancelCampaignUsingPUT

> Campaign CancelCampaignUsingPUT(ctx, campaignId, xAPIKEY, optional)

Cancel a campaign

Usage of this API will be reported in your access log under 'device_campaign' category.<br><br>Restricted to API keys with at least one of the following roles : CAMPAIGN_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string**| identifier of a campaign | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CancelCampaignUsingPUTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CancelCampaignUsingPUTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| When force is false, the campaign manager prevents devices with status not started from starting, cancel devices with status pending and let devices with status in progress finish properly. When force is true, the campaign manager aborts all pending and in progress sequences immediately. | [default to false]

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreateCampaignUsingPOST

> Campaign CreateCampaignUsingPOST(ctx, xAPIKEY, campaign)

Create a campaign

Usage of this API will be reported in your access log under 'device_campaign' category.<br><br>Restricted to API keys with at least one of the following roles : CAMPAIGN_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
**campaign** | [**CampaignDefinition**](CampaignDefinition.md)| campaign | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCampaignUsingDELETE

> ResponseEntity DeleteCampaignUsingDELETE(ctx, campaignId, xAPIKEY, optional)

Delete a campaign

Usage of this API will be reported in your access log under 'device_campaign' category.<br><br>Restricted to API keys with at least one of the following roles : CAMPAIGN_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string**| identifier of a campaign | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***DeleteCampaignUsingDELETEOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a DeleteCampaignUsingDELETEOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **force** | **optional.Bool**| cancelling mode. When force is false, the campaign manager prevents devices with status not started from starting, cancel devices with status pending and let devices with status in progress finish properly. When force is true, the campaign manager aborts all pending and in progress sequences immediately. | [default to false]

### Return type

[**ResponseEntity**](ResponseEntity.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignDetailsUsingGET

> CampaignPerTargetList GetCampaignDetailsUsingGET(ctx, campaignId, xAPIKEY, optional)

Get the campaign status per device

Restricted to API keys with at least one of the following roles : CAMPAIGN_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string**| identifier of a campaign | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***GetCampaignDetailsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a GetCampaignDetailsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **status** | **optional.String**| Status of the device | 
 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]

### Return type

[**CampaignPerTargetList**](CampaignPerTargetList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCampaignUsingGET

> Campaign GetCampaignUsingGET(ctx, campaignId, xAPIKEY)

Get a campaign

Restricted to API keys with at least one of the following roles : CAMPAIGN_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string**| identifier of a campaign | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListCampaignsUsingGET

> CampaignList ListCampaignsUsingGET(ctx, xAPIKEY, optional)

List all campaigns

Restricted to API keys with at least one of the following roles : CAMPAIGN_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***ListCampaignsUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a ListCampaignsUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **size** | **optional.Int32**| the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **optional.Int32**| the requested page number (optional) | [default to 0]
 **status** | [**optional.Interface of []string**](string.md)| status list filter | 
 **sort** | [**optional.Interface of []string**](string.md)| sorting list by attributes (supported columns:  id, name, status, created, updated). DefaultValue : -created | 

### Return type

[**CampaignList**](CampaignList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateCampaignUsingPATCH

> Campaign UpdateCampaignUsingPATCH(ctx, campaignId, xAPIKEY, campaign)

Update a campaign

Usage of this API will be reported in your access log under 'device_campaign' category.<br><br>Restricted to API keys with at least one of the following roles : CAMPAIGN_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**campaignId** | **string**| identifier of a campaign | 
**xAPIKEY** | **string**| a valid API key | 
**campaign** | [**CampaignUpdating**](CampaignUpdating.md)| campaign | 

### Return type

[**Campaign**](Campaign.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

