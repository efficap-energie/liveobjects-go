# \CACertificatesApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST**](CACertificatesApi.md#CreateUsingPOST) | **Post** /api/v0/certificates/ca | Upload CA certificate
[**DeleteUsingDELETE9**](CACertificatesApi.md#DeleteUsingDELETE9) | **Delete** /api/v0/certificates/ca/{certificateId} | Delete CA certificate
[**ListUsingGET11**](CACertificatesApi.md#ListUsingGET11) | **Get** /api/v0/certificates/ca | List CA certificates
[**RetrieveUsingGET**](CACertificatesApi.md#RetrieveUsingGET) | **Get** /api/v0/certificates/ca/{certificateId} | Retrieve CA certificate



## CreateUsingPOST

> CaCertificateCreateResWeb CreateUsingPOST(ctx, xAPIKEY, optional)

Upload CA certificate

Usage of this API will be reported in your access log under 'security' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***CreateUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a CreateUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **caCertificate** | [**optional.Interface of CaCertificateCreateReqWeb**](CaCertificateCreateReqWeb.md)| Root or intermediate Certification Authority (CA) certificate. Only PEM format is supported. Certificate chains are not allowed. In case that intermediates CA are used, it must be the intermediate CA certificate which signs the devices certificates. | 

### Return type

[**CaCertificateCreateResWeb**](CaCertificateCreateResWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE9

> DeleteUsingDELETE9(ctx, certificateId, xAPIKEY)

Delete CA certificate

Usage of this API will be reported in your access log under 'security' category.<br><br>Restricted to API keys with at least one of the following roles : API_KEY_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string**| id of the CA certificate to delete | 
**xAPIKEY** | **string**| a valid API key | 

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


## ListUsingGET11

> []CaCertificate ListUsingGET11(ctx, xAPIKEY)

List CA certificates

Restricted to API keys with at least one of the following roles : API_KEY_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**[]CaCertificate**](CaCertificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveUsingGET

> CaCertificate RetrieveUsingGET(ctx, certificateId, xAPIKEY)

Retrieve CA certificate

Restricted to API keys with at least one of the following roles : API_KEY_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string**| id of the CA certificate | 
**xAPIKEY** | **string**| a valid API key | 

### Return type

[**CaCertificate**](CaCertificate.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

