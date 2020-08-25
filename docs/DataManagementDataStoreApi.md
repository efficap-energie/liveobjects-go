# \DataManagementDataStoreApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddDataMessageUsingPOST**](DataManagementDataStoreApi.md#AddDataMessageUsingPOST) | **Post** /api/v0/data/streams/{streamId} | Insert a new Data into the stream
[**RetrieveDataUsingGET**](DataManagementDataStoreApi.md#RetrieveDataUsingGET) | **Get** /api/v0/data/streams/{streamId} | Retrieve data from the streamId



## AddDataMessageUsingPOST

> string AddDataMessageUsingPOST(ctx, streamId, xAPIKEY, optional)

Insert a new Data into the stream

In order for the data to be enriched by decoders and pipelines functionalities, 'metadata.encoding' field must be set.<br><br>Restricted to API keys with at least one of the following roles : DATA_W.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string**| StreamId in which the data will be added. Should not contains following character or space : &#39; \\ \&quot; ; { } ( ) | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***AddDataMessageUsingPOSTOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a AddDataMessageUsingPOSTOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **data** | [**optional.Interface of DataWeb**](DataWeb.md)| DataMessage to add | 

### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RetrieveDataUsingGET

> []DataStoredWeb RetrieveDataUsingGET(ctx, streamId, xAPIKEY, optional)

Retrieve data from the streamId

return an array of StoredDataMessage matching the request parameters.<br><br>Restricted to API keys with at least one of the following roles : DATA_R.

### Required Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**streamId** | **string**| StreamId from which the data will be retrieved | 
**xAPIKEY** | **string**| a valid API key | 
 **optional** | ***RetrieveDataUsingGETOpts** | optional parameters | nil if no parameters

### Optional Parameters

Optional parameters are passed through a pointer to a RetrieveDataUsingGETOpts struct


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **limit** | **optional.Int32**| max number of data to return, value is limited to 1000 | [default to 100]
 **timeRange** | [**optional.Interface of []string**](string.md)| filter data where timestamp is in timeRange (must be ISO 8601 formatted) : ?timeRange&#x3D;[lowerbound],[upperbound]. Coma is mandatory, lowerbound and upperbound are optionals. Lowerbound is inclusiv, upperbound is exclusiv. | 
 **bookmarkId** | **optional.String**| id of the last document retrieved that can be used to paginate : first result will be the one following this document id | 

### Return type

[**[]DataStoredWeb**](DataStoredWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

