# AssetResourceWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **string** | the asset id | [optional] 
**AssetIdNamespace** | Pointer to **string** | the asset namespace | [optional] 
**ConnectorMetadata** | Pointer to **map[string]string** | the connector metadata | [optional] 
**CurrentVersion** | Pointer to **string** | the current resource version | [optional] 
**CurrentVersionChangeTs** | Pointer to **int64** | the current resource version timestamp | [optional] 
**CurrentVersionTs** | Pointer to **int64** | the current resource version timestamp | [optional] 
**ResourceId** | Pointer to **string** | the resource&#39;s identifier | [optional] 
**TargetVersion** | Pointer to **string** | the target resource version timestamp | [optional] 
**TargetVersionTs** | Pointer to **int64** | the target resource version timestamp | [optional] 
**TenantId** | Pointer to **string** | identifier of tenant account | [optional] 

## Methods

### NewAssetResourceWeb

`func NewAssetResourceWeb() *AssetResourceWeb`

NewAssetResourceWeb instantiates a new AssetResourceWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetResourceWebWithDefaults

`func NewAssetResourceWebWithDefaults() *AssetResourceWeb`

NewAssetResourceWebWithDefaults instantiates a new AssetResourceWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *AssetResourceWeb) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *AssetResourceWeb) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *AssetResourceWeb) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *AssetResourceWeb) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetIdNamespace

`func (o *AssetResourceWeb) GetAssetIdNamespace() string`

GetAssetIdNamespace returns the AssetIdNamespace field if non-nil, zero value otherwise.

### GetAssetIdNamespaceOk

`func (o *AssetResourceWeb) GetAssetIdNamespaceOk() (*string, bool)`

GetAssetIdNamespaceOk returns a tuple with the AssetIdNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdNamespace

`func (o *AssetResourceWeb) SetAssetIdNamespace(v string)`

SetAssetIdNamespace sets AssetIdNamespace field to given value.

### HasAssetIdNamespace

`func (o *AssetResourceWeb) HasAssetIdNamespace() bool`

HasAssetIdNamespace returns a boolean if a field has been set.

### GetConnectorMetadata

`func (o *AssetResourceWeb) GetConnectorMetadata() map[string]string`

GetConnectorMetadata returns the ConnectorMetadata field if non-nil, zero value otherwise.

### GetConnectorMetadataOk

`func (o *AssetResourceWeb) GetConnectorMetadataOk() (*map[string]string, bool)`

GetConnectorMetadataOk returns a tuple with the ConnectorMetadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectorMetadata

`func (o *AssetResourceWeb) SetConnectorMetadata(v map[string]string)`

SetConnectorMetadata sets ConnectorMetadata field to given value.

### HasConnectorMetadata

`func (o *AssetResourceWeb) HasConnectorMetadata() bool`

HasConnectorMetadata returns a boolean if a field has been set.

### GetCurrentVersion

`func (o *AssetResourceWeb) GetCurrentVersion() string`

GetCurrentVersion returns the CurrentVersion field if non-nil, zero value otherwise.

### GetCurrentVersionOk

`func (o *AssetResourceWeb) GetCurrentVersionOk() (*string, bool)`

GetCurrentVersionOk returns a tuple with the CurrentVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersion

`func (o *AssetResourceWeb) SetCurrentVersion(v string)`

SetCurrentVersion sets CurrentVersion field to given value.

### HasCurrentVersion

`func (o *AssetResourceWeb) HasCurrentVersion() bool`

HasCurrentVersion returns a boolean if a field has been set.

### GetCurrentVersionChangeTs

`func (o *AssetResourceWeb) GetCurrentVersionChangeTs() int64`

GetCurrentVersionChangeTs returns the CurrentVersionChangeTs field if non-nil, zero value otherwise.

### GetCurrentVersionChangeTsOk

`func (o *AssetResourceWeb) GetCurrentVersionChangeTsOk() (*int64, bool)`

GetCurrentVersionChangeTsOk returns a tuple with the CurrentVersionChangeTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersionChangeTs

`func (o *AssetResourceWeb) SetCurrentVersionChangeTs(v int64)`

SetCurrentVersionChangeTs sets CurrentVersionChangeTs field to given value.

### HasCurrentVersionChangeTs

`func (o *AssetResourceWeb) HasCurrentVersionChangeTs() bool`

HasCurrentVersionChangeTs returns a boolean if a field has been set.

### GetCurrentVersionTs

`func (o *AssetResourceWeb) GetCurrentVersionTs() int64`

GetCurrentVersionTs returns the CurrentVersionTs field if non-nil, zero value otherwise.

### GetCurrentVersionTsOk

`func (o *AssetResourceWeb) GetCurrentVersionTsOk() (*int64, bool)`

GetCurrentVersionTsOk returns a tuple with the CurrentVersionTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentVersionTs

`func (o *AssetResourceWeb) SetCurrentVersionTs(v int64)`

SetCurrentVersionTs sets CurrentVersionTs field to given value.

### HasCurrentVersionTs

`func (o *AssetResourceWeb) HasCurrentVersionTs() bool`

HasCurrentVersionTs returns a boolean if a field has been set.

### GetResourceId

`func (o *AssetResourceWeb) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *AssetResourceWeb) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *AssetResourceWeb) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *AssetResourceWeb) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetTargetVersion

`func (o *AssetResourceWeb) GetTargetVersion() string`

GetTargetVersion returns the TargetVersion field if non-nil, zero value otherwise.

### GetTargetVersionOk

`func (o *AssetResourceWeb) GetTargetVersionOk() (*string, bool)`

GetTargetVersionOk returns a tuple with the TargetVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersion

`func (o *AssetResourceWeb) SetTargetVersion(v string)`

SetTargetVersion sets TargetVersion field to given value.

### HasTargetVersion

`func (o *AssetResourceWeb) HasTargetVersion() bool`

HasTargetVersion returns a boolean if a field has been set.

### GetTargetVersionTs

`func (o *AssetResourceWeb) GetTargetVersionTs() int64`

GetTargetVersionTs returns the TargetVersionTs field if non-nil, zero value otherwise.

### GetTargetVersionTsOk

`func (o *AssetResourceWeb) GetTargetVersionTsOk() (*int64, bool)`

GetTargetVersionTsOk returns a tuple with the TargetVersionTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetVersionTs

`func (o *AssetResourceWeb) SetTargetVersionTs(v int64)`

SetTargetVersionTs sets TargetVersionTs field to given value.

### HasTargetVersionTs

`func (o *AssetResourceWeb) HasTargetVersionTs() bool`

HasTargetVersionTs returns a boolean if a field has been set.

### GetTenantId

`func (o *AssetResourceWeb) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *AssetResourceWeb) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *AssetResourceWeb) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *AssetResourceWeb) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


