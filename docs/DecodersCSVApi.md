# \DecodersCSVApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUsingPUT9**](DecodersCSVApi.md#ActivateUsingPUT9) | **Put** /api/v0/decoders/csv/{decoderId}/enabled | Activate or deactivate a decoder
[**DeleteUsingDELETE10**](DecodersCSVApi.md#DeleteUsingDELETE10) | **Delete** /api/v0/decoders/csv/{decoderId} | Delete a csv decoder
[**GetUsingGET11**](DecodersCSVApi.md#GetUsingGET11) | **Get** /api/v0/decoders/csv/{decoderId} | Retrieve a csv decoder
[**ListUsingGET12**](DecodersCSVApi.md#ListUsingGET12) | **Get** /api/v0/decoders/csv | Retrieve the list of csv decoders filtered by tags
[**PostUsingPOST11**](DecodersCSVApi.md#PostUsingPOST11) | **Post** /api/v0/decoders/csv | Create a csv decoder
[**PutUsingPUT3**](DecodersCSVApi.md#PutUsingPUT3) | **Put** /api/v0/decoders/csv/{decoderId} | Update a csv decoder
[**TestUsingPOST3**](DecodersCSVApi.md#TestUsingPOST3) | **Post** /api/v0/decoders/csv/test | Test a csv decoder format with an encoded payload



## ActivateUsingPUT9

> ActivateUsingPUT9(ctx, decoderId).XAPIKEY(xAPIKEY).Enabled(enabled).Execute()

Activate or deactivate a decoder



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
    decoderId := "decoderId_example" // string | id of the csv decoder to activate or deactivate
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    enabled := true // bool | true to activate, false otherwise (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersCSVApi.ActivateUsingPUT9(context.Background(), decoderId).XAPIKEY(xAPIKEY).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersCSVApi.ActivateUsingPUT9``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the csv decoder to activate or deactivate | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUsingPUT9Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **enabled** | **bool** | true to activate, false otherwise | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteUsingDELETE10

> DeleteUsingDELETE10(ctx, decoderId).XAPIKEY(xAPIKEY).Execute()

Delete a csv decoder



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
    decoderId := "decoderId_example" // string | id of the csv decoder to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersCSVApi.DeleteUsingDELETE10(context.Background(), decoderId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersCSVApi.DeleteUsingDELETE10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the csv decoder to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

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


## GetUsingGET11

> CsvPayloadDescription GetUsingGET11(ctx, decoderId).XAPIKEY(xAPIKEY).Execute()

Retrieve a csv decoder



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
    decoderId := "decoderId_example" // string | id of the csv decoder to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersCSVApi.GetUsingGET11(context.Background(), decoderId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersCSVApi.GetUsingGET11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET11`: CsvPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersCSVApi.GetUsingGET11`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the csv decoder to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET12

> []CsvPayloadDescription ListUsingGET12(ctx).XAPIKEY(xAPIKEY).Tags(tags).Execute()

Retrieve the list of csv decoders filtered by tags



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
    tags := []string{"Inner_example"} // []string | target filtering tags (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersCSVApi.ListUsingGET12(context.Background()).XAPIKEY(xAPIKEY).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersCSVApi.ListUsingGET12``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET12`: []CsvPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersCSVApi.ListUsingGET12`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET12Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **tags** | [**[]string**](string.md) | target filtering tags | 

### Return type

[**[]CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST11

> CsvPayloadDescription PostUsingPOST11(ctx).XAPIKEY(xAPIKEY).CsvPayloadDescription(csvPayloadDescription).Execute()

Create a csv decoder



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
    csvPayloadDescription := *openapiclient.NewCsvPayloadDescription([]CsvColumn{*openapiclient.NewCsvColumn("JsonType_example", "Name_example")), false, "Encoding_example") // CsvPayloadDescription | Csv decoder description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersCSVApi.PostUsingPOST11(context.Background()).XAPIKEY(xAPIKEY).CsvPayloadDescription(csvPayloadDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersCSVApi.PostUsingPOST11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsingPOST11`: CsvPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersCSVApi.PostUsingPOST11`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsingPOST11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **csvPayloadDescription** | [**CsvPayloadDescription**](CsvPayloadDescription.md) | Csv decoder description | 

### Return type

[**CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUsingPUT3

> CsvPayloadDescription PutUsingPUT3(ctx, decoderId).XAPIKEY(xAPIKEY).CsvPayloadDescription(csvPayloadDescription).Execute()

Update a csv decoder



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
    decoderId := "decoderId_example" // string | id of the csv decoder to update
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    csvPayloadDescription := *openapiclient.NewCsvPayloadDescription([]CsvColumn{*openapiclient.NewCsvColumn("JsonType_example", "Name_example")), false, "Encoding_example") // CsvPayloadDescription | Csv decoder description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersCSVApi.PutUsingPUT3(context.Background(), decoderId).XAPIKEY(xAPIKEY).CsvPayloadDescription(csvPayloadDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersCSVApi.PutUsingPUT3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUsingPUT3`: CsvPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersCSVApi.PutUsingPUT3`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the csv decoder to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUsingPUT3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **csvPayloadDescription** | [**CsvPayloadDescription**](CsvPayloadDescription.md) | Csv decoder description | 

### Return type

[**CsvPayloadDescription**](CsvPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST3

> PayloadDescriptionTestResult TestUsingPOST3(ctx).XAPIKEY(xAPIKEY).DataDecodingTestRequest(dataDecodingTestRequest).Execute()

Test a csv decoder format with an encoded payload



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
    dataDecodingTestRequest := *openapiclient.NewCsvPayloadDescriptionTestRequest([]CsvColumn{)) // CsvPayloadDescriptionTestRequest | Csv format and encoded payload on which decoding shall be peformed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersCSVApi.TestUsingPOST3(context.Background()).XAPIKEY(xAPIKEY).DataDecodingTestRequest(dataDecodingTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersCSVApi.TestUsingPOST3``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestUsingPOST3`: PayloadDescriptionTestResult
    fmt.Fprintf(os.Stdout, "Response from `DecodersCSVApi.TestUsingPOST3`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestUsingPOST3Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **dataDecodingTestRequest** | [**CsvPayloadDescriptionTestRequest**](CsvPayloadDescriptionTestRequest.md) | Csv format and encoded payload on which decoding shall be peformed | 

### Return type

[**PayloadDescriptionTestResult**](PayloadDescriptionTestResult.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

