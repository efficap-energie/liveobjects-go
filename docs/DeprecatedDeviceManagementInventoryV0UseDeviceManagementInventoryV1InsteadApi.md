# \DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAssetUsingPOST**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#CreateAssetUsingPOST) | **Post** /api/v0/assets | Create a device
[**DeleteDeviceStatusUsingDELETE**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#DeleteDeviceStatusUsingDELETE) | **Delete** /api/v0/assets/{assetNamespace}/{assetId} | Delete a device
[**GetAssetStatusUsingGET**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#GetAssetStatusUsingGET) | **Get** /api/v0/assets/{assetNamespace}/{assetId} | Get a device status
[**ListAssetNamespacesUsingGET**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#ListAssetNamespacesUsingGET) | **Get** /api/v0/inventory/namespaces | Enumerates the used asset namespaces
[**ListAssetsUsingGET**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#ListAssetsUsingGET) | **Get** /api/v0/assets | List registered assets status
[**UpdateAssetUsingPUT**](DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.md#UpdateAssetUsingPUT) | **Put** /api/v0/assets/{assetNamespace}/{assetId} | Update a device



## CreateAssetUsingPOST

> Asset CreateAssetUsingPOST(ctx).XAPIKEY(xAPIKEY).Body(body).Execute()

Create a device



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
    body := *openapiclient.NewAssetCreateReqWeb("Id_example", "Namespace_example") // AssetCreateReqWeb | The device to register (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.CreateAssetUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.CreateAssetUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateAssetUsingPOST`: Asset
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.CreateAssetUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateAssetUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **body** | [**AssetCreateReqWeb**](AssetCreateReqWeb.md) | The device to register | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteDeviceStatusUsingDELETE

> DeleteDeviceStatusUsingDELETE(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).Execute()

Delete a device



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
    assetNamespace := "assetNamespace_example" // string | targeted for deletion asset namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | targeted for deletion asset identifier. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.DeleteDeviceStatusUsingDELETE(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.DeleteDeviceStatusUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | targeted for deletion asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | targeted for deletion asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDeviceStatusUsingDELETERequest struct via the builder pattern


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


## GetAssetStatusUsingGET

> Asset GetAssetStatusUsingGET(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).Execute()

Get a device status



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
    assetNamespace := "assetNamespace_example" // string | requested asset namespace. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | requested asset identifier. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.GetAssetStatusUsingGET(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.GetAssetStatusUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAssetStatusUsingGET`: Asset
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.GetAssetStatusUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | requested asset namespace. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | requested asset identifier. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetAssetStatusUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetNamespacesUsingGET

> []string ListAssetNamespacesUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

Enumerates the used asset namespaces



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
    resp, r, err := api_client.DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.ListAssetNamespacesUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.ListAssetNamespacesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssetNamespacesUsingGET`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.ListAssetNamespacesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetNamespacesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 

### Return type

**[]string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListAssetsUsingGET

> PageableAsset ListAssetsUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Namespace(namespace).GroupPath(groupPath).GroupId(groupId).Id(id).Connected(connected).Name(name).Tags(tags).PropertyFilterName(propertyFilterName).Execute()

List registered assets status



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
    size := 987 // int64 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int64 | the requested page number (optional) (optional) (default to 0)
    sort := []string{"Inner_example"} // []string | sorting list by attributes.  Supported columns: namespace, id, creationTs, lastUpdateTs, connected, groupPath. (optional)
    namespace := "namespace_example" // string | filter list by asset namespace. Expected string (max 128 characters) (optional)
    groupPath := "groupPath_example" // string | filter list by asset groupPath.  (with optional use of wildcard '/_*' at the end of search term)Expected string (max 255 characters) (optional)
    groupId := "groupId_example" // string | filter list by asset groupId. Expected string (max 6 characters) (optional)
    id := "id_example" // string | filter list by asset id. (with optional use of wildcard '*' at the beginning or end of search term) Expected string (max 128 characters) (optional)
    connected := true // bool | filter list by asset connected state (optional)
    name := "name_example" // string | filter list by asset name.  (with optional use of wildcard '*' at the beginning or end of search term)Expected string (max 255 characters) (optional)
    tags := []string{"Inner_example"} // []string | filter list by asset tags. Max number of tags depends of your offer settings. Tag value max length is 32. (optional)
    propertyFilterName := "propertyFilterName_example" // string | Multiple filters, Example : assets?property.temperature=25&property.humidity=58...[cannot be tested in swagger-ui]. Max number of properties depends of your offer settings. A property name must not include following characters <code>$.</code> and max length is 128. Invalid property names are : 'class', '_class'. Property value max length is 256. (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.ListAssetsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Sort(sort).Namespace(namespace).GroupPath(groupPath).GroupId(groupId).Id(id).Connected(connected).Name(name).Tags(tags).PropertyFilterName(propertyFilterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.ListAssetsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListAssetsUsingGET`: PageableAsset
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.ListAssetsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListAssetsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int64** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int64** | the requested page number (optional) | [default to 0]
 **sort** | [**[]string**](string.md) | sorting list by attributes.  Supported columns: namespace, id, creationTs, lastUpdateTs, connected, groupPath. | 
 **namespace** | **string** | filter list by asset namespace. Expected string (max 128 characters) | 
 **groupPath** | **string** | filter list by asset groupPath.  (with optional use of wildcard &#39;/_*&#39; at the end of search term)Expected string (max 255 characters) | 
 **groupId** | **string** | filter list by asset groupId. Expected string (max 6 characters) | 
 **id** | **string** | filter list by asset id. (with optional use of wildcard &#39;*&#39; at the beginning or end of search term) Expected string (max 128 characters) | 
 **connected** | **bool** | filter list by asset connected state | 
 **name** | **string** | filter list by asset name.  (with optional use of wildcard &#39;*&#39; at the beginning or end of search term)Expected string (max 255 characters) | 
 **tags** | [**[]string**](string.md) | filter list by asset tags. Max number of tags depends of your offer settings. Tag value max length is 32. | 
 **propertyFilterName** | **string** | Multiple filters, Example : assets?property.temperature&#x3D;25&amp;property.humidity&#x3D;58...[cannot be tested in swagger-ui]. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | 

### Return type

[**PageableAsset**](Pageable«Asset».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateAssetUsingPUT

> Asset UpdateAssetUsingPUT(ctx, assetNamespace, assetId).XAPIKEY(xAPIKEY).Body(body).Execute()

Update a device



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
    assetNamespace := "assetNamespace_example" // string | namespace targeted to update assets. Asset namespace must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    assetId := "assetId_example" // string | asset identifier to update. Asset identifier must respect the following regular expression <code>([\\w\\-_:]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    body := *openapiclient.NewAssetUpdateReqWeb() // AssetUpdateReqWeb | the updated asset definition (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.UpdateAssetUsingPUT(context.Background(), assetNamespace, assetId).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.UpdateAssetUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateAssetUsingPUT`: Asset
    fmt.Fprintf(os.Stdout, "Response from `DeprecatedDeviceManagementInventoryV0UseDeviceManagementInventoryV1InsteadApi.UpdateAssetUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**assetNamespace** | **string** | namespace targeted to update assets. Asset namespace must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**assetId** | **string** | asset identifier to update. Asset identifier must respect the following regular expression &lt;code&gt;([\\w\\-_:]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAssetUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **body** | [**AssetUpdateReqWeb**](AssetUpdateReqWeb.md) | the updated asset definition | 

### Return type

[**Asset**](Asset.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

