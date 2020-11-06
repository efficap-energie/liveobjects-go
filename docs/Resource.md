# Resource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | the resource&#39;s connector | [optional] 
**CreationTs** | Pointer to **int64** | creation timestamp | [optional] 
**Description** | Pointer to **string** | the resource&#39;s description | [optional] 
**Label** | Pointer to **string** | the resource&#39;s label | [optional] 
**Metadata** | Pointer to **map[string]string** | the resource metadata | [optional] 
**ResourceId** | Pointer to **string** | the resource&#39;s identifier | [optional] 
**TenantId** | Pointer to **string** | identifier of tenant account | [optional] 
**UpdateTs** | Pointer to **int64** | last update timestamp | [optional] 
**VersionAliases** | Pointer to **map[string]string** | the resource version aliases | [optional] 

## Methods

### NewResource

`func NewResource() *Resource`

NewResource instantiates a new Resource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceWithDefaults

`func NewResourceWithDefaults() *Resource`

NewResourceWithDefaults instantiates a new Resource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *Resource) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *Resource) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *Resource) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *Resource) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetCreationTs

`func (o *Resource) GetCreationTs() int64`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *Resource) GetCreationTsOk() (*int64, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *Resource) SetCreationTs(v int64)`

SetCreationTs sets CreationTs field to given value.

### HasCreationTs

`func (o *Resource) HasCreationTs() bool`

HasCreationTs returns a boolean if a field has been set.

### GetDescription

`func (o *Resource) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Resource) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Resource) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Resource) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *Resource) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Resource) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Resource) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Resource) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetMetadata

`func (o *Resource) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *Resource) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *Resource) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *Resource) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourceId

`func (o *Resource) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *Resource) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *Resource) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.

### HasResourceId

`func (o *Resource) HasResourceId() bool`

HasResourceId returns a boolean if a field has been set.

### GetTenantId

`func (o *Resource) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *Resource) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *Resource) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *Resource) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetUpdateTs

`func (o *Resource) GetUpdateTs() int64`

GetUpdateTs returns the UpdateTs field if non-nil, zero value otherwise.

### GetUpdateTsOk

`func (o *Resource) GetUpdateTsOk() (*int64, bool)`

GetUpdateTsOk returns a tuple with the UpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTs

`func (o *Resource) SetUpdateTs(v int64)`

SetUpdateTs sets UpdateTs field to given value.

### HasUpdateTs

`func (o *Resource) HasUpdateTs() bool`

HasUpdateTs returns a boolean if a field has been set.

### GetVersionAliases

`func (o *Resource) GetVersionAliases() map[string]string`

GetVersionAliases returns the VersionAliases field if non-nil, zero value otherwise.

### GetVersionAliasesOk

`func (o *Resource) GetVersionAliasesOk() (*map[string]string, bool)`

GetVersionAliasesOk returns a tuple with the VersionAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionAliases

`func (o *Resource) SetVersionAliases(v map[string]string)`

SetVersionAliases sets VersionAliases field to given value.

### HasVersionAliases

`func (o *Resource) HasVersionAliases() bool`

HasVersionAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


