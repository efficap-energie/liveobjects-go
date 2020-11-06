# \DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateBusinessSettingsUsingPOST**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#CreateBusinessSettingsUsingPOST) | **Post** /api/v0/sms-connector/settings/business | Create a new business settings of SMS Connector (use Device management - Interfaces - V1 instead)
[**DeleteBusinessSettingsByMsiSDNUsingDELETE**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#DeleteBusinessSettingsByMsiSDNUsingDELETE) | **Delete** /api/v0/sms-connector/settings/business/msisdn | Delete msisdn in business settings of SMS Connector
[**DeleteBusinessSettingsForOneMsiSDNUsingDELETE**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#DeleteBusinessSettingsForOneMsiSDNUsingDELETE) | **Delete** /api/v0/sms-connector/settings/business/{serverPhoneNumber}/msisdn/{msisdnNumber} | Delete one msisdn in business settings of SMS Connector
[**GetBusinessSettingsUsingGET**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#GetBusinessSettingsUsingGET) | **Get** /api/v0/sms-connector/settings/business | Get a business settings of SMS Connector
[**ListMsisdnUsingGET**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#ListMsisdnUsingGET) | **Get** /api/v0/sms-connector/settings/business/msisdn | List msisdn of business settings of SMS Connector
[**ListSettingsUsingGET**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#ListSettingsUsingGET) | **Get** /api/v0/sms-connector/settings | List all the business settings of the SMSConnector for a tenant
[**UpdateBusinessSettingsUsingPOST**](DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.md#UpdateBusinessSettingsUsingPOST) | **Post** /api/v0/sms-connector/settings/business/{serverPhoneNumber} | Update business settings of the SMSConnector



## CreateBusinessSettingsUsingPOST

> ConnectorStatusResponse CreateBusinessSettingsUsingPOST(ctx).XAPIKEY(xAPIKEY).Request(request).Execute()

Create a new business settings of SMS Connector (use Device management - Interfaces - V1 instead)



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
    request := *openapiclient.NewSMSConnectorBusinessSettingsCreationReqWeb([]Msisdns{*openapiclient.NewMsisdns("Msisdn_example")), "ServerPhoneNumber_example") // SMSConnectorBusinessSettingsCreationReqWeb | body to create business settings for SMS Connector (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.CreateBusinessSettingsUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.CreateBusinessSettingsUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateBusinessSettingsUsingPOST`: ConnectorStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.CreateBusinessSettingsUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateBusinessSettingsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**SMSConnectorBusinessSettingsCreationReqWeb**](SMSConnectorBusinessSettingsCreationReqWeb.md) | body to create business settings for SMS Connector | 

### Return type

[**ConnectorStatusResponse**](ConnectorStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessSettingsByMsiSDNUsingDELETE

> ConnectorStatusResponse DeleteBusinessSettingsByMsiSDNUsingDELETE(ctx).XAPIKEY(xAPIKEY).Request(request).Execute()

Delete msisdn in business settings of SMS Connector



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
    request := *openapiclient.NewSMSConnectorBusinessSettingsDeleteMsisdnReqWeb([]string{"Msisdns_example"), "ServerPhoneNumber_example") // SMSConnectorBusinessSettingsDeleteMsisdnReqWeb | body to delete business settings per msisdn for SMS Connector (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.DeleteBusinessSettingsByMsiSDNUsingDELETE(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.DeleteBusinessSettingsByMsiSDNUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeleteBusinessSettingsByMsiSDNUsingDELETE`: ConnectorStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.DeleteBusinessSettingsByMsiSDNUsingDELETE`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessSettingsByMsiSDNUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**SMSConnectorBusinessSettingsDeleteMsisdnReqWeb**](SMSConnectorBusinessSettingsDeleteMsisdnReqWeb.md) | body to delete business settings per msisdn for SMS Connector | 

### Return type

[**ConnectorStatusResponse**](ConnectorStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteBusinessSettingsForOneMsiSDNUsingDELETE

> DeleteBusinessSettingsForOneMsiSDNUsingDELETE(ctx, serverPhoneNumber, msisdnNumber).XAPIKEY(xAPIKEY).Execute()

Delete one msisdn in business settings of SMS Connector



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
    serverPhoneNumber := "serverPhoneNumber_example" // string | server phone number ex: \"20259\", // Must be defined in OfferSettings
    msisdnNumber := "msisdnNumber_example" // string | msisdn number ex: \"0606060606\"
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.DeleteBusinessSettingsForOneMsiSDNUsingDELETE(context.Background(), serverPhoneNumber, msisdnNumber).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.DeleteBusinessSettingsForOneMsiSDNUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverPhoneNumber** | **string** | server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 
**msisdnNumber** | **string** | msisdn number ex: \&quot;0606060606\&quot; | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteBusinessSettingsForOneMsiSDNUsingDELETERequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

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


## GetBusinessSettingsUsingGET

> BusinessSettings GetBusinessSettingsUsingGET(ctx).ServerPhoneNumber(serverPhoneNumber).XAPIKEY(xAPIKEY).DecoderName(decoderName).Execute()

Get a business settings of SMS Connector



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
    serverPhoneNumber := "serverPhoneNumber_example" // string | the server phone number ex: \"20259\", // Must be defined in OfferSettings
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    decoderName := "decoderName_example" // string | the decoder Name ex: \"decoderName\" (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.GetBusinessSettingsUsingGET(context.Background()).ServerPhoneNumber(serverPhoneNumber).XAPIKEY(xAPIKEY).DecoderName(decoderName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.GetBusinessSettingsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetBusinessSettingsUsingGET`: BusinessSettings
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.GetBusinessSettingsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetBusinessSettingsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverPhoneNumber** | **string** | the server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 
 **xAPIKEY** | **string** | a valid API key | 
 **decoderName** | **string** | the decoder Name ex: \&quot;decoderName\&quot; | 

### Return type

[**BusinessSettings**](BusinessSettings.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListMsisdnUsingGET

> SMSConnectorMsisdnPageWeb ListMsisdnUsingGET(ctx).ServerPhoneNumber(serverPhoneNumber).XAPIKEY(xAPIKEY).DecoderName(decoderName).Size(size).Page(page).Execute()

List msisdn of business settings of SMS Connector



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
    serverPhoneNumber := "serverPhoneNumber_example" // string | the server phone number ex: \"20259\", // Must be defined in OfferSettings
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    decoderName := "decoderName_example" // string | the decoder Name ex: \"decoderName\" (optional)
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.ListMsisdnUsingGET(context.Background()).ServerPhoneNumber(serverPhoneNumber).XAPIKEY(xAPIKEY).DecoderName(decoderName).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.ListMsisdnUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListMsisdnUsingGET`: SMSConnectorMsisdnPageWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.ListMsisdnUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListMsisdnUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **serverPhoneNumber** | **string** | the server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 
 **xAPIKEY** | **string** | a valid API key | 
 **decoderName** | **string** | the decoder Name ex: \&quot;decoderName\&quot; | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**SMSConnectorMsisdnPageWeb**](SMSConnectorMsisdnPageWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListSettingsUsingGET

> SMSConnectorBusinessSettingsPageWeb ListSettingsUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()

List all the business settings of the SMSConnector for a tenant



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
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.ListSettingsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.ListSettingsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListSettingsUsingGET`: SMSConnectorBusinessSettingsPageWeb
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.ListSettingsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListSettingsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**SMSConnectorBusinessSettingsPageWeb**](SMSConnectorBusinessSettingsPageWeb.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateBusinessSettingsUsingPOST

> ConnectorStatusResponse UpdateBusinessSettingsUsingPOST(ctx, serverPhoneNumber).XAPIKEY(xAPIKEY).Request(request).Execute()

Update business settings of the SMSConnector



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
    serverPhoneNumber := "serverPhoneNumber_example" // string | server phone number ex: \"20259\", // Must be defined in OfferSettings
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    request := *openapiclient.NewSMSConnectorBusinessSettingsUpdateReqWeb() // SMSConnectorBusinessSettingsUpdateReqWeb | body to update business settings for SMS Connector (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.UpdateBusinessSettingsUsingPOST(context.Background(), serverPhoneNumber).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.UpdateBusinessSettingsUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateBusinessSettingsUsingPOST`: ConnectorStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedSMSConnectorBusinessSettingsUseDeviceManagementConnectorNodesV1InsteadApi.UpdateBusinessSettingsUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**serverPhoneNumber** | **string** | server phone number ex: \&quot;20259\&quot;, // Must be defined in OfferSettings | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateBusinessSettingsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**SMSConnectorBusinessSettingsUpdateReqWeb**](SMSConnectorBusinessSettingsUpdateReqWeb.md) | body to update business settings for SMS Connector | 

### Return type

[**ConnectorStatusResponse**](ConnectorStatusResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

