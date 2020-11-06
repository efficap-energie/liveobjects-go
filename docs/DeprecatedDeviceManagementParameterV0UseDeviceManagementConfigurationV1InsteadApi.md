# \DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetAssetParamUsingGET**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#GetAssetParamUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId}/params/{paramKey} | Get a specific asset parameter
[**GetAssetParamsUsingGET**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#GetAssetParamsUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId}/params | Get a specific asset list of parameters
[**SetAssetParamsUsingPOST**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#SetAssetParamsUsingPOST) | **Post** /api/v0/assets/{assetNamespace}/{assetId}/params | Update a specific asset list of parameters
[**SetDeviceParamUpdateStatusUsingPUT**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#SetDeviceParamUpdateStatusUsingPUT) | **Put** /api/v0/assets/{assetNamespace}/{assetId}/params/{paramKey}/status | Update the status of a specific asset parameter update
[**SetDeviceParamsUpdateStatusUsingPUT**](DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.md#SetDeviceParamsUpdateStatusUsingPUT) | **Put** /api/v0/assets/{assetNamespace}/{assetId}/params/status | Update the status of a specific asset parameters update



## GetAssetParamUsingGET

> AssetParameter GetAssetParamUsingGET(ctx, assetNamespace, assetId, paramKey).XAPIKEY(xAPIKEY).Execute()

Get a specific asset parameter



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
    assetNamespace := "assetNamespace_example" // string | asset namespace ex : myNode1. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    paramKey := "paramKey_example" // string | key identifying the targeted asset parameter. Expected string (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.GetAssetParamUsingGET(context.Background(), assetNamespace, assetId, paramKey).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.GetAssetParamUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetParamUsingGET`: AssetParameter
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.GetAssetParamUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**paramKey** | **string** | key identifying the targeted asset parameter. Expected string (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetParamUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**AssetParameter**](AssetParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetAssetParamsUsingGET

> map[string]AssetParameter GetAssetParamsUsingGET(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).Execute()

Get a specific asset list of parameters



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
    assetNamespace := "assetNamespace_example" // string | asset namespace ex : myNode1. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.GetAssetParamsUsingGET(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.GetAssetParamsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetParamsUsingGET`: map[string]AssetParameter
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.GetAssetParamsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetParamsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**map[string]AssetParameter**](AssetParameter.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetAssetParamsUsingPOST

> SetAssetParamsUsingPOST(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).NotifyTo(notifyTo).NewParamValues(newParamValues).Execute()

Update a specific asset list of parameters



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
    assetNamespace := "assetNamespace_example" // string | asset namespace ex : myNode1. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    notifyTo := "notifyTo_example" // string | fifo used to relay status notification ex: fifo/~notif. A valid 'notify to' starts with one of ['fifo/', 'pubsub/', 'router/'] and max length is 255 (optional)
    newParamValues := map[string]AssetParameterValue{ "key": *openapiclient.NewAssetParameterValue()} // map[string]AssetParameterValue | new param value ( type ex :  INT32, UINT32 , STRING, FLOAT, BINARY ). Max number of parameters is 100. Parameter name max length is 128. Parameter value must be valid according to the type and max length is 2000. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.SetAssetParamsUsingPOST(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).NotifyTo(notifyTo).NewParamValues(newParamValues).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.SetAssetParamsUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetAssetParamsUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **notifyTo** | **string** | fifo used to relay status notification ex: fifo/~notif. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;, &#39;router/&#39;] and max length is 255 | 
 **newParamValues** | [**map[string]AssetParameterValue**](AssetParameterValue.md) | new param value ( type ex :  INT32, UINT32 , STRING, FLOAT, BINARY ). Max number of parameters is 100. Parameter name max length is 128. Parameter value must be valid according to the type and max length is 2000. | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceParamUpdateStatusUsingPUT

> SetDeviceParamUpdateStatusUsingPUT(ctx, assetNamespace, assetId, paramKey).XAPIKEY(xAPIKEY).Force(force).NewStatus(newStatus).Execute()

Update the status of a specific asset parameter update



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
    assetNamespace := "assetNamespace_example" // string | asset namespace ex : myNode1. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    paramKey := "paramKey_example" // string | key identifying the targeted asset parameter. Expected string (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | force the status update (optional) (default to false)
    newStatus := "newStatus_example" // string | future state of the parameter --> CANCELED (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.SetDeviceParamUpdateStatusUsingPUT(context.Background(), assetNamespace, assetId, paramKey).XAPIKEY(xAPIKEY).Force(force).NewStatus(newStatus).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.SetDeviceParamUpdateStatusUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 
**paramKey** | **string** | key identifying the targeted asset parameter. Expected string (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceParamUpdateStatusUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | force the status update | [default to false]
 **newStatus** | **string** | future state of the parameter --&gt; CANCELED | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SetDeviceParamsUpdateStatusUsingPUT

> SetDeviceParamsUpdateStatusUsingPUT(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).Force(force).ReqWeb(reqWeb).Execute()

Update the status of a specific asset parameters update



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
    assetNamespace := "assetNamespace_example" // string | asset namespace ex : myNode1. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    force := true // bool | force the status update (optional) (default to false)
    reqWeb := *openapiclient.NewAssetParamsStatusUpdateReqWeb("NewStatus_example", []string{"ParamKeys_example")) // AssetParamsStatusUpdateReqWeb | parameters keys and their future status. Maximum 100 parameters, and parameter name max length is 128 (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.SetDeviceParamsUpdateStatusUsingPUT(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).Force(force).ReqWeb(reqWeb).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementParameterV0UseDeviceManagementConfigurationV1InsteadApi.SetDeviceParamsUpdateStatusUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | asset namespace ex : myNode1. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | asset identifier ex: assetInteg. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiSetDeviceParamsUpdateStatusUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **force** | **bool** | force the status update | [default to false]
 **reqWeb** | [**AssetParamsStatusUpdateReqWeb**](AssetParamsStatusUpdateReqWeb.md) | parameters keys and their future status. Maximum 100 parameters, and parameter name max length is 128 | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

