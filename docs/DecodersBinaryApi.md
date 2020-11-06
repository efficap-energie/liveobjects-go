# \DecodersBinaryApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUsingPUT8**](DecodersBinaryApi.md#ActivateUsingPUT8) | **Put** /api/v0/decoders/binary/{decoderId}/enabled | Activate or deactivate a decoder
[**DeleteUsingDELETE8**](DecodersBinaryApi.md#DeleteUsingDELETE8) | **Delete** /api/v0/decoders/binary/{decoderId} | Delete a binary decoder
[**GetUsingGET10**](DecodersBinaryApi.md#GetUsingGET10) | **Get** /api/v0/decoders/binary/{decoderId} | Retrieve a binary decoder
[**ListUsingGET10**](DecodersBinaryApi.md#ListUsingGET10) | **Get** /api/v0/decoders/binary | Retrieve the list of binary decoders
[**PostUsingPOST10**](DecodersBinaryApi.md#PostUsingPOST10) | **Post** /api/v0/decoders/binary | Create a binary decoder
[**PutUsingPUT2**](DecodersBinaryApi.md#PutUsingPUT2) | **Put** /api/v0/decoders/binary/{decoderId} | Update a binary decoder
[**TestUsingPOST2**](DecodersBinaryApi.md#TestUsingPOST2) | **Post** /api/v0/decoders/binary/test | Test a binary decoder format with an encoded payload



## ActivateUsingPUT8

> ActivateUsingPUT8(ctx, decoderId).XAPIKEY(xAPIKEY).Enabled(enabled).Execute()

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
    decoderId := "decoderId_example" // string | id of the binary decoder to activate or deactivate
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    enabled := true // bool | true to activate, false otherwise (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersBinaryApi.ActivateUsingPUT8(context.Background(), decoderId).XAPIKEY(xAPIKEY).Enabled(enabled).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersBinaryApi.ActivateUsingPUT8``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the binary decoder to activate or deactivate | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUsingPUT8Request struct via the builder pattern


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


## DeleteUsingDELETE8

> DeleteUsingDELETE8(ctx, decoderId).XAPIKEY(xAPIKEY).Execute()

Delete a binary decoder



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
    decoderId := "decoderId_example" // string | id of the binary decoder to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersBinaryApi.DeleteUsingDELETE8(context.Background(), decoderId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersBinaryApi.DeleteUsingDELETE8``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the binary decoder to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE8Request struct via the builder pattern


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


## GetUsingGET10

> BinaryPayloadDescription GetUsingGET10(ctx, decoderId).XAPIKEY(xAPIKEY).Execute()

Retrieve a binary decoder



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
    decoderId := "decoderId_example" // string | id of the binary decoder to retrieve
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersBinaryApi.GetUsingGET10(context.Background(), decoderId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersBinaryApi.GetUsingGET10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetUsingGET10`: BinaryPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersBinaryApi.GetUsingGET10`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the binary decoder to retrieve | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetUsingGET10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsingGET10

> []BinaryPayloadDescription ListUsingGET10(ctx).XAPIKEY(xAPIKEY).Tags(tags).Execute()

Retrieve the list of binary decoders



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
    resp, r, err := api_client.DecodersBinaryApi.ListUsingGET10(context.Background()).XAPIKEY(xAPIKEY).Tags(tags).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersBinaryApi.ListUsingGET10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET10`: []BinaryPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersBinaryApi.ListUsingGET10`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **tags** | [**[]string**](string.md) | target filtering tags | 

### Return type

[**[]BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostUsingPOST10

> BinaryPayloadDescription PostUsingPOST10(ctx).XAPIKEY(xAPIKEY).BinaryPayloadDescription(binaryPayloadDescription).Execute()

Create a binary decoder



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
    binaryPayloadDescription := *openapiclient.NewBinaryPayloadDescription(false, "Encoding_example", "Format_example") // BinaryPayloadDescription | Binary decoder description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersBinaryApi.PostUsingPOST10(context.Background()).XAPIKEY(xAPIKEY).BinaryPayloadDescription(binaryPayloadDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersBinaryApi.PostUsingPOST10``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PostUsingPOST10`: BinaryPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersBinaryApi.PostUsingPOST10`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostUsingPOST10Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **binaryPayloadDescription** | [**BinaryPayloadDescription**](BinaryPayloadDescription.md) | Binary decoder description | 

### Return type

[**BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutUsingPUT2

> BinaryPayloadDescription PutUsingPUT2(ctx, decoderId).XAPIKEY(xAPIKEY).BinaryPayloadDescription(binaryPayloadDescription).Execute()

Update a binary decoder



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
    decoderId := "decoderId_example" // string | id of the binary decoder to update
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    binaryPayloadDescription := *openapiclient.NewBinaryPayloadDescription(false, "Encoding_example", "Format_example") // BinaryPayloadDescription | Binary decoder description (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersBinaryApi.PutUsingPUT2(context.Background(), decoderId).XAPIKEY(xAPIKEY).BinaryPayloadDescription(binaryPayloadDescription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersBinaryApi.PutUsingPUT2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PutUsingPUT2`: BinaryPayloadDescription
    fmt.Fprintf(os.Stdout, "Response from `DecodersBinaryApi.PutUsingPUT2`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**decoderId** | **string** | id of the binary decoder to update | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutUsingPUT2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **binaryPayloadDescription** | [**BinaryPayloadDescription**](BinaryPayloadDescription.md) | Binary decoder description | 

### Return type

[**BinaryPayloadDescription**](BinaryPayloadDescription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TestUsingPOST2

> PayloadDescriptionTestResult TestUsingPOST2(ctx).XAPIKEY(xAPIKEY).DataDecodingTestRequest(dataDecodingTestRequest).Execute()

Test a binary decoder format with an encoded payload



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
    dataDecodingTestRequest := *openapiclient.NewBinaryPayloadDescriptionTestRequest() // BinaryPayloadDescriptionTestRequest | Binary format and encoded payload on which decoding shall be peformed (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DecodersBinaryApi.TestUsingPOST2(context.Background()).XAPIKEY(xAPIKEY).DataDecodingTestRequest(dataDecodingTestRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DecodersBinaryApi.TestUsingPOST2``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `TestUsingPOST2`: PayloadDescriptionTestResult
    fmt.Fprintf(os.Stdout, "Response from `DecodersBinaryApi.TestUsingPOST2`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTestUsingPOST2Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **dataDecodingTestRequest** | [**BinaryPayloadDescriptionTestRequest**](BinaryPayloadDescriptionTestRequest.md) | Binary format and encoded payload on which decoding shall be peformed | 

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

