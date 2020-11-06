# UpdateDeviceResourceReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | the update metadata. Max number of metadata is 100. Metadata name max length is 255. Metadata value max length is 255. | [optional] 
**Version** | Pointer to **string** | requested resource version. Expected string (max 255 characters) | [optional] 

## Methods

### NewUpdateDeviceResourceReqWeb

`func NewUpdateDeviceResourceReqWeb() *UpdateDeviceResourceReqWeb`

NewUpdateDeviceResourceReqWeb instantiates a new UpdateDeviceResourceReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceResourceReqWebWithDefaults

`func NewUpdateDeviceResourceReqWebWithDefaults() *UpdateDeviceResourceReqWeb`

NewUpdateDeviceResourceReqWebWithDefaults instantiates a new UpdateDeviceResourceReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *UpdateDeviceResourceReqWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateDeviceResourceReqWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateDeviceResourceReqWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateDeviceResourceReqWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateDeviceResourceReqWeb) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateDeviceResourceReqWeb) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateDeviceResourceReqWeb) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateDeviceResourceReqWeb) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


