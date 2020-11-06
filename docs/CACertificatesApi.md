# \CACertificatesApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateUsingPOST**](CACertificatesApi.md#CreateUsingPOST) | **Post** /api/v0/certificates/ca | Upload CA certificate
[**DeleteUsingDELETE9**](CACertificatesApi.md#DeleteUsingDELETE9) | **Delete** /api/v0/certificates/ca/{certificateId} | Delete CA certificate
[**ListUsingGET11**](CACertificatesApi.md#ListUsingGET11) | **Get** /api/v0/certificates/ca | List CA certificates
[**RetrieveUsingGET**](CACertificatesApi.md#RetrieveUsingGET) | **Get** /api/v0/certificates/ca/{certificateId} | Retrieve CA certificate



## CreateUsingPOST

> CaCertificateCreateResWeb CreateUsingPOST(ctx).XAPIKEY(xAPIKEY).CaCertificate(caCertificate).Execute()

Upload CA certificate



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
    caCertificate := *openapiclient.NewCaCertificateCreateReqWeb("Certificate_example", "Comment_example") // CaCertificateCreateReqWeb | Root or intermediate Certification Authority (CA) certificate. Only PEM format is supported. Certificate chains are not allowed. In case that intermediates CA are used, it must be the intermediate CA certificate which signs the devices certificates. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CACertificatesApi.CreateUsingPOST(context.Background()).XAPIKEY(xAPIKEY).CaCertificate(caCertificate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CACertificatesApi.CreateUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateUsingPOST`: CaCertificateCreateResWeb
    fmt.Fprintf(os.Stdout, "Response from `CACertificatesApi.CreateUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **caCertificate** | [**CaCertificateCreateReqWeb**](CaCertificateCreateReqWeb.md) | Root or intermediate Certification Authority (CA) certificate. Only PEM format is supported. Certificate chains are not allowed. In case that intermediates CA are used, it must be the intermediate CA certificate which signs the devices certificates. | 

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

> DeleteUsingDELETE9(ctx, certificateId).XAPIKEY(xAPIKEY).Execute()

Delete CA certificate



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
    certificateId := "certificateId_example" // string | id of the CA certificate to delete
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CACertificatesApi.DeleteUsingDELETE9(context.Background(), certificateId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CACertificatesApi.DeleteUsingDELETE9``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string** | id of the CA certificate to delete | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUsingDELETE9Request struct via the builder pattern


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


## ListUsingGET11

> []CaCertificate ListUsingGET11(ctx).XAPIKEY(xAPIKEY).Execute()

List CA certificates



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

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CACertificatesApi.ListUsingGET11(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CACertificatesApi.ListUsingGET11``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsingGET11`: []CaCertificate
    fmt.Fprintf(os.Stdout, "Response from `CACertificatesApi.ListUsingGET11`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListUsingGET11Request struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

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

> CaCertificate RetrieveUsingGET(ctx, certificateId).XAPIKEY(xAPIKEY).Execute()

Retrieve CA certificate



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
    certificateId := "certificateId_example" // string | id of the CA certificate
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.CACertificatesApi.RetrieveUsingGET(context.Background(), certificateId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `CACertificatesApi.RetrieveUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `RetrieveUsingGET`: CaCertificate
    fmt.Fprintf(os.Stdout, "Response from `CACertificatesApi.RetrieveUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**certificateId** | **string** | id of the CA certificate | 

### Other Parameters

Other parameters are passed through a pointer to a apiRetrieveUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

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

