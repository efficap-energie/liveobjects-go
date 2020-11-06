# \AuditLogApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SearchUsingGET**](AuditLogApi.md#SearchUsingGET) | **Get** /api/v0/auditlog/messages | Retrieve messages available in your AuditLog



## SearchUsingGET

> []AuditLogMessage SearchUsingGET(ctx).XAPIKEY(xAPIKEY).From(from).To(to).Offset(offset).Limit(limit).Sort(sort).Filters(filters).Any(any).Execute()

Retrieve messages available in your AuditLog



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
    from := "from_example" // string | Search for messages after this date. Use ISO-8601 normalization. (optional)
    to := "to_example" // string | Search for messages before this date. Use ISO-8601 normalization. (optional)
    offset := 987 // int32 | Offset from the first result you want to fetch. offset + limit should not exceed 10.000. (optional) (default to 0)
    limit := 987 // int32 | Maximum amount of messages to return. limit should not exceed 1.000. (optional) (default to 100)
    sort := "sort_example" // string | Sort order, based on timestamp field of the AuditLog message. (optional) (default to "desc")
    filters := "filters_example" // string | Filter query based on parameter name = field path and parameter value = value to search for : {fiedlName}={value}. You can put several filters that way, they will all be treated with an AND operator. Common filters field names you can use : level, category, subcategory, type, source.deviceId, source.nodeId, description, detailedDescription. e.g. : level=error&source.deviceId=urn:lora:0E5EAB0ABCD00000 (optional)
    any := []string{"Inner_example"} // []string | Search for AuditLog Messages where any of the fields contains all these values. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.AuditLogApi.SearchUsingGET(context.Background()).XAPIKEY(xAPIKEY).From(from).To(to).Offset(offset).Limit(limit).Sort(sort).Filters(filters).Any(any).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AuditLogApi.SearchUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SearchUsingGET`: []AuditLogMessage
    fmt.Fprintf(os.Stdout, "Response from `AuditLogApi.SearchUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSearchUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **from** | **string** | Search for messages after this date. Use ISO-8601 normalization. | 
 **to** | **string** | Search for messages before this date. Use ISO-8601 normalization. | 
 **offset** | **int32** | Offset from the first result you want to fetch. offset + limit should not exceed 10.000. | [default to 0]
 **limit** | **int32** | Maximum amount of messages to return. limit should not exceed 1.000. | [default to 100]
 **sort** | **string** | Sort order, based on timestamp field of the AuditLog message. | [default to &quot;desc&quot;]
 **filters** | **string** | Filter query based on parameter name &#x3D; field path and parameter value &#x3D; value to search for : {fiedlName}&#x3D;{value}. You can put several filters that way, they will all be treated with an AND operator. Common filters field names you can use : level, category, subcategory, type, source.deviceId, source.nodeId, description, detailedDescription. e.g. : level&#x3D;error&amp;source.deviceId&#x3D;urn:lora:0E5EAB0ABCD00000 | 
 **any** | [**[]string**](string.md) | Search for AuditLog Messages where any of the fields contains all these values. | 

### Return type

[**[]AuditLogMessage**](AuditLogMessage.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

