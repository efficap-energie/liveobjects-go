# \PartnersTokensManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartnerTokenUsingPOST**](PartnersTokensManagementApi.md#CreatePartnerTokenUsingPOST) | **Post** /api/v0/oauth2/token | Request a new Token



## CreatePartnerTokenUsingPOST

> PartnerToken CreatePartnerTokenUsingPOST(ctx, authorization, optional)

Request a new Token

Usage of this API will be reported in your access log under 'authentication' category.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authorization** | **string**| Basic Base64(clientId:clientSecret) | 
 **optional** | ***CreatePartnerTokenUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreatePartnerTokenUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **grantType** | **optional.String**| the authorization grant type. Must be set to \&quot;client_credentials\&quot; | 

### Return type

[**PartnerToken**](PartnerToken.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

