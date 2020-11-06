# UpdateDeviceFirmwareReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | the update metadata. Max number of metadata is 100. Metadata name max length is 255. Metadata value max length is 255. | [optional] 
**NotifyTo** | Pointer to **string** | (optional) topic where firmware update status change events must be published. A valid &#39;notify to&#39; starts with one of [&#39;fifo/&#39;, &#39;pubsub/&#39;,&#39;router/&#39;]and max length is 255 | [optional] 
**Version** | Pointer to **string** | requested firmware version. Expected string (max 255 characters) | [optional] 

## Methods

### NewUpdateDeviceFirmwareReqWeb

`func NewUpdateDeviceFirmwareReqWeb() *UpdateDeviceFirmwareReqWeb`

NewUpdateDeviceFirmwareReqWeb instantiates a new UpdateDeviceFirmwareReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateDeviceFirmwareReqWebWithDefaults

`func NewUpdateDeviceFirmwareReqWebWithDefaults() *UpdateDeviceFirmwareReqWeb`

NewUpdateDeviceFirmwareReqWebWithDefaults instantiates a new UpdateDeviceFirmwareReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *UpdateDeviceFirmwareReqWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *UpdateDeviceFirmwareReqWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *UpdateDeviceFirmwareReqWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *UpdateDeviceFirmwareReqWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotifyTo

`func (o *UpdateDeviceFirmwareReqWeb) GetNotifyTo() string`

GetNotifyTo returns the NotifyTo field if non-nil, zero value otherwise.

### GetNotifyToOk

`func (o *UpdateDeviceFirmwareReqWeb) GetNotifyToOk() (*string, bool)`

GetNotifyToOk returns a tuple with the NotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTo

`func (o *UpdateDeviceFirmwareReqWeb) SetNotifyTo(v string)`

SetNotifyTo sets NotifyTo field to given value.

### HasNotifyTo

`func (o *UpdateDeviceFirmwareReqWeb) HasNotifyTo() bool`

HasNotifyTo returns a boolean if a field has been set.

### GetVersion

`func (o *UpdateDeviceFirmwareReqWeb) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateDeviceFirmwareReqWeb) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateDeviceFirmwareReqWeb) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *UpdateDeviceFirmwareReqWeb) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


