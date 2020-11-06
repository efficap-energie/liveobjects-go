# \DeviceManagementConnectorNodesV1Api

All URIs are relative to *https://liveobjects.orange-business.com*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteNodeUsingDELETE**](DeviceManagementConnectorNodesV1Api.md#DeleteNodeUsingDELETE) | **Delete** /api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId} | Delete a connector node
[**GetNodeUsingGET**](DeviceManagementConnectorNodesV1Api.md#GetNodeUsingGET) | **Get** /api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId} | Get details of connector node
[**ListNodesUsingGET**](DeviceManagementConnectorNodesV1Api.md#ListNodesUsingGET) | **Get** /api/v1/deviceMgt/connectors/{connector}/nodes | List all connector nodes
[**UpdateNodeUsingPATCH**](DeviceManagementConnectorNodesV1Api.md#UpdateNodeUsingPATCH) | **Patch** /api/v1/deviceMgt/connectors/{connector}/nodes/{nodeId} | Update a connector node



## DeleteNodeUsingDELETE

> DeleteNodeUsingDELETE(ctx, connector, nodeId).XAPIKEY(xAPIKEY).Execute()

Delete a connector node



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
    connector := "connector_example" // string | the connector id
    nodeId := "nodeId_example" // string | the node id
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementConnectorNodesV1Api.DeleteNodeUsingDELETE(context.Background(), connector, nodeId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConnectorNodesV1Api.DeleteNodeUsingDELETE``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** | the connector id | 
**nodeId** | **string** | the node id | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteNodeUsingDELETERequest struct via the builder pattern


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


## GetNodeUsingGET

> ConnectorNode GetNodeUsingGET(ctx, connector, nodeId).XAPIKEY(xAPIKEY).Execute()

Get details of connector node



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
    connector := "connector_example" // string | the connector
    nodeId := "nodeId_example" // string | the node id
    xAPIKEY := "xAPIKEY_example" // string | a valid API key

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementConnectorNodesV1Api.GetNodeUsingGET(context.Background(), connector, nodeId).XAPIKEY(xAPIKEY).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConnectorNodesV1Api.GetNodeUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetNodeUsingGET`: ConnectorNode
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementConnectorNodesV1Api.GetNodeUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** | the connector | 
**nodeId** | **string** | the node id | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 

### Return type

[**ConnectorNode**](ConnectorNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListNodesUsingGET

> []ConnectorNode ListNodesUsingGET(ctx, connector).XAPIKEY(xAPIKEY).Limit(limit).Offset(offset).XTotalCount(xTotalCount).Execute()

List all connector nodes



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
    connector := "connector_example" // string | the connector
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    limit := 987 // int32 | the maximum number of items per page (optional, highest value is 1000) (optional) (default to 20)
    offset := 987 // int32 | number of items to skip (optional) (optional) (default to 0)
    xTotalCount := true // bool | true if a total count must be returned in response (optional) (default to false)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementConnectorNodesV1Api.ListNodesUsingGET(context.Background(), connector).XAPIKEY(xAPIKEY).Limit(limit).Offset(offset).XTotalCount(xTotalCount).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConnectorNodesV1Api.ListNodesUsingGET``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListNodesUsingGET`: []ConnectorNode
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementConnectorNodesV1Api.ListNodesUsingGET`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** | the connector | 

### Other Parameters

Other parameters are passed through a pointer to a apiListNodesUsingGETRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **xAPIKEY** | **string** | a valid API key | 
 **limit** | **int32** | the maximum number of items per page (optional, highest value is 1000) | [default to 20]
 **offset** | **int32** | number of items to skip (optional) | [default to 0]
 **xTotalCount** | **bool** | true if a total count must be returned in response | [default to false]

### Return type

[**[]ConnectorNode**](ConnectorNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateNodeUsingPATCH

> ConnectorNode UpdateNodeUsingPATCH(ctx, connector, nodeId).XAPIKEY(xAPIKEY).Body(body).Execute()

Update a connector node



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
    connector := "connector_example" // string | the connector id
    nodeId := "nodeId_example" // string | the node id
    xAPIKEY := "xAPIKEY_example" // string | a valid API key
    body := *openapiclient.NewUpdateConnectorNodeRequest() // UpdateConnectorNodeRequest | The fields to update (optional)

    configuration := openapiclient.NewConfiguration()
    api_client := openapiclient.NewAPIClient(configuration)
    resp, r, err := api_client.DeviceManagementConnectorNodesV1Api.UpdateNodeUsingPATCH(context.Background(), connector, nodeId).XAPIKEY(xAPIKEY).Body(body).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `DeviceManagementConnectorNodesV1Api.UpdateNodeUsingPATCH``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateNodeUsingPATCH`: ConnectorNode
    fmt.Fprintf(os.Stdout, "Response from `DeviceManagementConnectorNodesV1Api.UpdateNodeUsingPATCH`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**connector** | **string** | the connector id | 
**nodeId** | **string** | the node id | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateNodeUsingPATCHRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **xAPIKEY** | **string** | a valid API key | 
 **body** | [**UpdateConnectorNodeRequest**](UpdateConnectorNodeRequest.md) | The fields to update | 

### Return type

[**ConnectorNode**](ConnectorNode.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

