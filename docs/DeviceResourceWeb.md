# DeviceResourceWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | Pointer to **map[string]string** | metadata associated with this device resource | [optional] 
**Reported** | Pointer to [**DeviceResourceVersionValueWeb**](DeviceResourceVersionValueWeb.md) |  | [optional] 
**Requested** | Pointer to [**DeviceResourceVersionValueWeb**](DeviceResourceVersionValueWeb.md) |  | [optional] 

## Methods

### NewDeviceResourceWeb

`func NewDeviceResourceWeb() *DeviceResourceWeb`

NewDeviceResourceWeb instantiates a new DeviceResourceWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceResourceWebWithDefaults

`func NewDeviceResourceWebWithDefaults() *DeviceResourceWeb`

NewDeviceResourceWebWithDefaults instantiates a new DeviceResourceWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *DeviceResourceWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DeviceResourceWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DeviceResourceWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DeviceResourceWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetReported

`func (o *DeviceResourceWeb) GetReported() DeviceResourceVersionValueWeb`

GetReported returns the Reported field if non-nil, zero value otherwise.

### GetReportedOk

`func (o *DeviceResourceWeb) GetReportedOk() (*DeviceResourceVersionValueWeb, bool)`

GetReportedOk returns a tuple with the Reported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReported

`func (o *DeviceResourceWeb) SetReported(v DeviceResourceVersionValueWeb)`

SetReported sets Reported field to given value.

### HasReported

`func (o *DeviceResourceWeb) HasReported() bool`

HasReported returns a boolean if a field has been set.

### GetRequested

`func (o *DeviceResourceWeb) GetRequested() DeviceResourceVersionValueWeb`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *DeviceResourceWeb) GetRequestedOk() (*DeviceResourceVersionValueWeb, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *DeviceResourceWeb) SetRequested(v DeviceResourceVersionValueWeb)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *DeviceResourceWeb) HasRequested() bool`

HasRequested returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


