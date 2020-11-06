# \DeviceManagementResourcesManagementApi

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddResourceUsingPOST**](DeviceManagementResourcesManagementApi.md#AddResourceUsingPOST) | **Post** /api/v0/rm | Add a resource in Resource Manager
[**AddResourceVersionUsingPOST**](DeviceManagementResourcesManagementApi.md#AddResourceVersionUsingPOST) | **Post** /api/v0/rm/{resourceId}/version | Add a new version to a resource
[**DeleteResourceUsingDELETE**](DeviceManagementResourcesManagementApi.md#DeleteResourceUsingDELETE) | **Delete** /api/v0/rm/{resourceId} | Delete a resource
[**DeleteResourceVersionUsingDELETE**](DeviceManagementResourcesManagementApi.md#DeleteResourceVersionUsingDELETE) | **Delete** /api/v0/rm/{resourceId}/version/{resourceVersionId} | Delete a resource&#39;s version
[**GetAllConnectorsUsingGET**](DeviceManagementResourcesManagementApi.md#GetAllConnectorsUsingGET) | **Get** /api/v0/rm/connectors | Get all available connectors
[**GetConnectorMandatoryAndOptionalMetadataUsingGET**](DeviceManagementResourcesManagementApi.md#GetConnectorMandatoryAndOptionalMetadataUsingGET) | **Get** /api/v0/rm/connectors/{connector}/metadata | Get mandatory and optional metadata of a connector
[**GetResourceCompatibleVersionsUsingGET**](DeviceManagementResourcesManagementApi.md#GetResourceCompatibleVersionsUsingGET) | **Get** /api/v0/rm/{resourceId}/compatibleVersion/{currentVersion} | List the versions from which a resource update to the version can be done
[**GetResourceUsingGET**](DeviceManagementResourcesManagementApi.md#GetResourceUsingGET) | **Get** /api/v0/rm/{resourceId} | Get a resource
[**GetResourceVersionUsingGET**](DeviceManagementResourcesManagementApi.md#GetResourceVersionUsingGET) | **Get** /api/v0/rm/{resourceId}/version/{resourceVersionId} | Get a resource&#39;s version
[**GetResourceVersionsUsingGET**](DeviceManagementResourcesManagementApi.md#GetResourceVersionsUsingGET) | **Get** /api/v0/rm/{resourceId}/version | List resource&#39;s versions
[**GetResourcesUsingGET**](DeviceManagementResourcesManagementApi.md#GetResourcesUsingGET) | **Get** /api/v0/rm | List all resources
[**UpdateResourceUsingPUT**](DeviceManagementResourcesManagementApi.md#UpdateResourceUsingPUT) | **Put** /api/v0/rm/{resourceId} | Update a resource
[**UpdateResourceVersionUsingPUT**](DeviceManagementResourcesManagementApi.md#UpdateResourceVersionUsingPUT) | **Put** /api/v0/rm/{resourceId}/version/{resourceVersionId} | Update a resource&#39;s version



## AddResourceUsingPOST

> Resource AddResourceUsingPOST(ctx).XAPIKEY(xAPIKEY).Request(request).Execute()

Add a resource in Resource Manager



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
    request := *openapiclient.NewResourceAddReqWeb("Connector_example", "Label_example", "ResourceId_example") // ResourceAddReqWeb | body of add resource (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.AddResourceUsingPOST(context.Background()).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.AddResourceUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddResourceUsingPOST`: Resource
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.AddResourceUsingPOST`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**ResourceAddReqWeb**](ResourceAddReqWeb.md) | body of add resource | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## AddResourceVersionUsingPOST

> ResourceVersion AddResourceVersionUsingPOST(ctx, resourceId).XAPIKEY(xAPIKEY).Request(request).Execute()

Add a new version to a resource



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    request := *openapiclient.NewResourceVersionAddReqWeb("Checksum_example", "File_example", "ResourceVersionId_example") // ResourceVersionAddReqWeb | body of add resource version (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.AddResourceVersionUsingPOST(context.Background(), resourceId).XAPIKEY(xAPIKEY).Request(request).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.AddResourceVersionUsingPOST``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AddResourceVersionUsingPOST`: ResourceVersion
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.AddResourceVersionUsingPOST`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiAddResourceVersionUsingPOSTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **request** | [**ResourceVersionAddReqWeb**](ResourceVersionAddReqWeb.md) | body of add resource version | 

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteResourceUsingDELETE

> DeleteResourceUsingDELETE(ctx, resourceId).XAPIKEY(xAPIKEY).Execute()

Delete a resource



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.DeleteResourceUsingDELETE(context.Background(), resourceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.DeleteResourceUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceUsingDELETERequest struct via the builder pattern


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


## DeleteResourceVersionUsingDELETE

> DeleteResourceVersionUsingDELETE(ctx, resourceId, resourceVersionId).XAPIKEY(xAPIKEY).Execute()

Delete a resource's version



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    resourceVersionId := "resourceVersionId_example" // string | the resource version id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.DeleteResourceVersionUsingDELETE(context.Background(), resourceId, resourceVersionId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.DeleteResourceVersionUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 
**resourceVersionId** | **string** | the resource version id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteResourceVersionUsingDELETERequest struct via the builder pattern


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


## GetAllConnectorsUsingGET

> []string GetAllConnectorsUsingGET(ctx).XAPIKEY(xAPIKEY).Execute()

Get all available connectors



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
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.GetAllConnectorsUsingGET(context.Background()).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.GetAllConnectorsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetAllConnectorsUsingGET`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.GetAllConnectorsUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetAllConnectorsUsingGETRequest struct via the builder pattern


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


## GetConnectorMandatoryAndOptionalMetadataUsingGET

> MetadataResourceConnectors GetConnectorMandatoryAndOptionalMetadataUsingGET(ctx, connector).XAPIKEY(xAPIKEY).Execute()

Get mandatory and optional metadata of a connector



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
    connector := "connector_example" // string | the resource's connector (ex. \"http-updater\"). A connector must respect the following regular expression <code>([\\w\\-_]{1,128})</code> (max 128 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.GetConnectorMandatoryAndOptionalMetadataUsingGET(context.Background(), connector).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.GetConnectorMandatoryAndOptionalMetadataUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetConnectorMandatoryAndOptionalMetadataUsingGET`: MetadataResourceConnectors
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.GetConnectorMandatoryAndOptionalMetadataUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** | the resource&#39;s connector (ex. \&quot;http-updater\&quot;). A connector must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConnectorMandatoryAndOptionalMetadataUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**MetadataResourceConnectors**](MetadataResourceConnectors.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceCompatibleVersionsUsingGET

> []string GetResourceCompatibleVersionsUsingGET(ctx, resourceId, currentVersion).XAPIKEY(xAPIKEY).Execute()

List the versions from which a resource update to the version can be done



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    currentVersion := "currentVersion_example" // string | current version. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.GetResourceCompatibleVersionsUsingGET(context.Background(), resourceId, currentVersion).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.GetResourceCompatibleVersionsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceCompatibleVersionsUsingGET`: []string
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.GetResourceCompatibleVersionsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 
**currentVersion** | **string** | current version. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceCompatibleVersionsUsingGETRequest struct via the builder pattern


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


## GetResourceUsingGET

> Resource GetResourceUsingGET(ctx, resourceId).XAPIKEY(xAPIKEY).Execute()

Get a resource



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.GetResourceUsingGET(context.Background(), resourceId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.GetResourceUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceUsingGET`: Resource
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.GetResourceUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceVersionUsingGET

> ResourceVersion GetResourceVersionUsingGET(ctx, resourceId, resourceVersionId).XAPIKEY(xAPIKEY).Execute()

Get a resource's version



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    resourceVersionId := "resourceVersionId_example" // string | the resource version id
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.GetResourceVersionUsingGET(context.Background(), resourceId, resourceVersionId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.GetResourceVersionUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceVersionUsingGET`: ResourceVersion
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.GetResourceVersionUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 
**resourceVersionId** | **string** | the resource version id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceVersionUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourceVersionsUsingGET

> PageableResourceVersion GetResourceVersionsUsingGET(ctx, resourceId).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()

List resource's versions



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    size := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    page := 987 // int32 | the requested page number (optional) (optional) (default to 0)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.GetResourceVersionsUsingGET(context.Background(), resourceId).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.GetResourceVersionsUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourceVersionsUsingGET`: PageableResourceVersion
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.GetResourceVersionsUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetResourceVersionsUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**PageableResourceVersion**](Pageable«ResourceVersion».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetResourcesUsingGET

> PageableResource GetResourcesUsingGET(ctx).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()

List all resources



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
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.GetResourcesUsingGET(context.Background()).XAPIKEY(xAPIKEY).Size(size).Page(page).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.GetResourcesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetResourcesUsingGET`: PageableResource
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.GetResourcesUsingGET`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetResourcesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **xAPIKEY** | **string** | a valid API key | 
 **size** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **page** | **int32** | the requested page number (optional) | [default to 0]

### Return type

[**PageableResource**](Pageable«Resource».md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceUsingPUT

> Resource UpdateResourceUsingPUT(ctx, resourceId).XAPIKEY(xAPIKEY).ResourceUpdate(resourceUpdate).Execute()

Update a resource



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
    resourceId := "resourceId_example" // string | the resource Id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    resourceUpdate := *openapiclient.NewResourceUpdateReqWeb("Connector_example", "Label_example") // ResourceUpdateReqWeb | body of resource update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.UpdateResourceUsingPUT(context.Background(), resourceId).XAPIKEY(xAPIKEY).ResourceUpdate(resourceUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.UpdateResourceUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResourceUsingPUT`: Resource
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.UpdateResourceUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource Id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **resourceUpdate** | [**ResourceUpdateReqWeb**](ResourceUpdateReqWeb.md) | body of resource update | 

### Return type

[**Resource**](Resource.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateResourceVersionUsingPUT

> ResourceVersion UpdateResourceVersionUsingPUT(ctx, resourceId, resourceVersionId).XAPIKEY(xAPIKEY).ResourceVersionUpdate(resourceVersionUpdate).Execute()

Update a resource's version



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
    resourceId := "resourceId_example" // string | the resource id. Expected string (max 255 characters)
    resourceVersionId := "resourceVersionId_example" // string | the resource version id. Expected string (max 255 characters)
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    resourceVersionUpdate := *openapiclient.NewResourceVersionUpdateReqWeb() // ResourceVersionUpdateReqWeb | body of resource version update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementResourcesManagementApi.UpdateResourceVersionUsingPUT(context.Background(), resourceId, resourceVersionId).XAPIKEY(xAPIKEY).ResourceVersionUpdate(resourceVersionUpdate).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementResourcesManagementApi.UpdateResourceVersionUsingPUT``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateResourceVersionUsingPUT`: ResourceVersion
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementResourcesManagementApi.UpdateResourceVersionUsingPUT`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**resourceId** | **string** | the resource id. Expected string (max 255 characters) | 
**resourceVersionId** | **string** | the resource version id. Expected string (max 255 characters) | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateResourceVersionUsingPUTRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **resourceVersionUpdate** | [**ResourceVersionUpdateReqWeb**](ResourceVersionUpdateReqWeb.md) | body of resource version update | 

### Return type

[**ResourceVersion**](ResourceVersion.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

