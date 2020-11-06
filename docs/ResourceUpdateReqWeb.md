# ResourceUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** | the resource&#39;s connector. A connector must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | 
**Description** | Pointer to **string** | the resource&#39;s description. Expected string (max 255 characters) | [optional] 
**Label** | **string** | the resource&#39;s label. Expected string (max 255 characters) | 
**Metadata** | Pointer to **map[string]string** | the operation metadata. Max number of metadata is 100. Metadata name max length is 255. Metadata value max length is 255. | [optional] 
**VersionAliases** | Pointer to **map[string]string** | the resource version aliases. Max number of alias is 5. Alias name max length is 128. Alias value max length is 255. | [optional] 

## Methods

### NewResourceUpdateReqWeb

`func NewResourceUpdateReqWeb(connector string, label string, ) *ResourceUpdateReqWeb`

NewResourceUpdateReqWeb instantiates a new ResourceUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUpdateReqWebWithDefaults

`func NewResourceUpdateReqWebWithDefaults() *ResourceUpdateReqWeb`

NewResourceUpdateReqWebWithDefaults instantiates a new ResourceUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *ResourceUpdateReqWeb) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *ResourceUpdateReqWeb) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *ResourceUpdateReqWeb) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetDescription

`func (o *ResourceUpdateReqWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceUpdateReqWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceUpdateReqWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceUpdateReqWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetLabel

`func (o *ResourceUpdateReqWeb) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *ResourceUpdateReqWeb) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *ResourceUpdateReqWeb) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetMetadata

`func (o *ResourceUpdateReqWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceUpdateReqWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceUpdateReqWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceUpdateReqWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetVersionAliases

`func (o *ResourceUpdateReqWeb) GetVersionAliases() map[string]string`

GetVersionAliases returns the VersionAliases field if non-nil, zero value otherwise.

### GetVersionAliasesOk

`func (o *ResourceUpdateReqWeb) GetVersionAliasesOk() (*map[string]string, bool)`

GetVersionAliasesOk returns a tuple with the VersionAliases field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersionAliases

`func (o *ResourceUpdateReqWeb) SetVersionAliases(v map[string]string)`

SetVersionAliases sets VersionAliases field to given value.

### HasVersionAliases

`func (o *ResourceUpdateReqWeb) HasVersionAliases() bool`

HasVersionAliases returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


