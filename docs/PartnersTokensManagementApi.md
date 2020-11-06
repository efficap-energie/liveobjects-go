# \PartnersTokensManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePartnerTokenUsingPOST**](PartnersTokensManagementApi.md#CreatePartnerTokenUsingPOST) | **Post** /api/v0/oauth2/token | Request a new Token



## CreatePartnerTokenUsingPOST

> PartnerToken CreatePartnerTokenUsingPOST(ctx).Authorization(authorization).GrantType(grantType).Execute()

Request a new Token



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
    authorization := "authorization_example" // string | Basic Base64(clientId:clientSecret)
    grantType := "grantType_example" // string | the authorization grant type. Must be set to \"client_credentials\" (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.PartnersTokensManagementApi.CreatePartnerTokenUsingPOST(context.Background()).Authorization(authorization).GrantType(grantType).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PartnersTokensManagementApi.CreatePartnerTokenUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePartnerTokenUsingPOST`: PartnerToken
    fmt.Fprintf(os.Stdout, "Response from `PartnersTokensManagementApi.CreatePartnerTokenUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePartnerTokenUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **authorization** | **string** | Basic Base64(clientId:clientSecret) | 
 **grantType** | **string** | the authorization grant type. Must be set to \&quot;client_credentials\&quot; | 

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

