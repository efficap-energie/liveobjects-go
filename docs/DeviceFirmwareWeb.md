# DeviceFirmwareWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | metadata associated with this device firmware | [optional] 
**Reported** | Pointer to [**DeviceFirmwareVersionValueWeb**](DeviceFirmwareVersionValueWeb.md) |  | [optional] 
**Requested** | Pointer to [**DeviceFirmwareVersionValueWeb**](DeviceFirmwareVersionValueWeb.md) |  | [optional] 

## Methods

### NewDeviceFirmwareWeb

`func NewDeviceFirmwareWeb() *DeviceFirmwareWeb`

NewDeviceFirmwareWeb instantiates a new DeviceFirmwareWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceFirmwareWebWithDefaults

`func NewDeviceFirmwareWebWithDefaults() *DeviceFirmwareWeb`

NewDeviceFirmwareWebWithDefaults instantiates a new DeviceFirmwareWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DeviceFirmwareWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceFirmwareWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceFirmwareWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeviceFirmwareWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReported

`func (o *DeviceFirmwareWeb) GetReported() DeviceFirmwareVersionValueWeb`

GetReported returns the Reported field if non-nil, zero value otherwise.

### GetReportedOk

`func (o *DeviceFirmwareWeb) GetReportedOk() (*DeviceFirmwareVersionValueWeb, bool)`

GetReportedOk returns a tuple with the Reported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReported

`func (o *DeviceFirmwareWeb) SetReported(v DeviceFirmwareVersionValueWeb)`

SetReported sets Reported field to given value.

### HasReported

`func (o *DeviceFirmwareWeb) HasReported() bool`

HasReported returns a boolean if a field has been set.

### GetRequested

`func (o *DeviceFirmwareWeb) GetRequested() DeviceFirmwareVersionValueWeb`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *DeviceFirmwareWeb) GetRequestedOk() (*DeviceFirmwareVersionValueWeb, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *DeviceFirmwareWeb) SetRequested(v DeviceFirmwareVersionValueWeb)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *DeviceFirmwareWeb) HasRequested() bool`

HasRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


