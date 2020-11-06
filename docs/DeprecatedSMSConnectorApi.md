# \DeprecatedSMSConnectorApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendSMSUsingPOST**](DeprecatedSMSConnectorApi.md#SendSMSUsingPOST) | **Post** /api/v0/sms-connector/sms | send SMS by SMS Connector for a list of MSISDN



## SendSMSUsingPOST

> ConnectorStatusResponse SendSMSUsingPOST(ctx).XAPIKEY(xAPIKEY).Request(request).Execute()

send SMS by SMS Connector for a list of MSISDN



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
    request := *openapiclient.NewSMSConnectorSendSMSReqWeb([]string{"Msisdns_example"), "ServerPhoneNumber_example") // SMSConnectorSendSMSReqWeb | body to send SMS via SMS Connector (only one Payload field must be set) (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedSMSConnectorApi.SendSMSUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedSMSConnectorApi.SendSMSUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `SendSMSUsingPOST`: ConnectorStatusResponse
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedSMSConnectorApi.SendSMSUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiSendSMSUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**SMSConnectorSendSMSReqWeb**](SMSConnectorSendSMSReqWeb.md) | body to send SMS via SMS Connector (only one Payload field must be set) | 

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

