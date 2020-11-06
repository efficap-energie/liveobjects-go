# DeviceParameterWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reported** | Pointer to [**DeviceParameterValueWeb**](DeviceParameterValueWeb.md) |  | [optional] 
**Requested** | Pointer to [**DeviceParameterValueWeb**](DeviceParameterValueWeb.md) |  | [optional] 
**SyncStatus** | Pointer to **string** | configuration parameter synchronization status: NONE, PENDING, OK or FAILED | [optional] 

## Methods

### NewDeviceParameterWeb

`func NewDeviceParameterWeb() *DeviceParameterWeb`

NewDeviceParameterWeb instantiates a new DeviceParameterWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceParameterWebWithDefaults

`func NewDeviceParameterWebWithDefaults() *DeviceParameterWeb`

NewDeviceParameterWebWithDefaults instantiates a new DeviceParameterWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReported

`func (o *DeviceParameterWeb) GetReported() DeviceParameterValueWeb`

GetReported returns the Reported field if non-nil, zero value otherwise.

### GetReportedOk

`func (o *DeviceParameterWeb) GetReportedOk() (*DeviceParameterValueWeb, bool)`

GetReportedOk returns a tuple with the Reported field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReported

`func (o *DeviceParameterWeb) SetReported(v DeviceParameterValueWeb)`

SetReported sets Reported field to given value.

### HasReported

`func (o *DeviceParameterWeb) HasReported() bool`

HasReported returns a boolean if a field has been set.

### GetRequested

`func (o *DeviceParameterWeb) GetRequested() DeviceParameterValueWeb`

GetRequested returns the Requested field if non-nil, zero value otherwise.

### GetRequestedOk

`func (o *DeviceParameterWeb) GetRequestedOk() (*DeviceParameterValueWeb, bool)`

GetRequestedOk returns a tuple with the Requested field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequested

`func (o *DeviceParameterWeb) SetRequested(v DeviceParameterValueWeb)`

SetRequested sets Requested field to given value.

### HasRequested

`func (o *DeviceParameterWeb) HasRequested() bool`

HasRequested returns a boolean if a field has been set.

### GetSyncStatus

`func (o *DeviceParameterWeb) GetSyncStatus() string`

GetSyncStatus returns the SyncStatus field if non-nil, zero value otherwise.

### GetSyncStatusOk

`func (o *DeviceParameterWeb) GetSyncStatusOk() (*string, bool)`

GetSyncStatusOk returns a tuple with the SyncStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSyncStatus

`func (o *DeviceParameterWeb) SetSyncStatus(v string)`

SetSyncStatus sets SyncStatus field to given value.

### HasSyncStatus

`func (o *DeviceParameterWeb) HasSyncStatus() bool`

HasSyncStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


