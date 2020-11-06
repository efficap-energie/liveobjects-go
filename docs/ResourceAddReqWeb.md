# ResourceAddReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** | the resource&#39;s connector. A connector must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**Description** | Pointer to **string** | the resource&#39;s description. Expected string (max 255 characters) | [optional] 
**Label** | **string** | the resource&#39;s label. Expected string (max 255 characters) | 
**Metadata** | Pointer to **map[string]string** | the resource metadata. Max number of metadata is 100. Metadata name max length is 255. Metadata value max length is 255. | [optional] 
**ResourceId** | **string** | the resource&#39;s id. Expected string (max 255 characters) | 

## Methods

### NewResourceAddReqWeb

`func NewResourceAddReqWeb(connector string, label string, resourceId string, ) *ResourceAddReqWeb`

NewResourceAddReqWeb instantiates a new ResourceAddReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceAddReqWebWithDefaults

`func NewResourceAddReqWebWithDefaults() *ResourceAddReqWeb`

NewResourceAddReqWebWithDefaults instantiates a new ResourceAddReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ResourceAddReqWeb) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ResourceAddReqWeb) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ResourceAddReqWeb) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetDescription

`func (o *ResourceAddReqWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceAddReqWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceAddReqWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceAddReqWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *ResourceAddReqWeb) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ResourceAddReqWeb) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ResourceAddReqWeb) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetadata

`func (o *ResourceAddReqWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceAddReqWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceAddReqWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceAddReqWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetResourceId

`func (o *ResourceAddReqWeb) GetResourceId() string`

GetResourceId returns the ResourceId field if non-nil, zero value otherwise.

### GetResourceIdOk

`func (o *ResourceAddReqWeb) GetResourceIdOk() (*string, bool)`

GetResourceIdOk returns a tuple with the ResourceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceId

`func (o *ResourceAddReqWeb) SetResourceId(v string)`

SetResourceId sets ResourceId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


