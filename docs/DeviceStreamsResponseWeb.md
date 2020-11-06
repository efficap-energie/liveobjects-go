# DeviceStreamsResponseWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StreamId** | Pointer to **string** | id of the stream | [optional] 
**Count** | Pointer to **int32** | number of messages sent by the device into this stream | [optional] 
**LastUpdate** | Pointer to **string** | date of the last message sent by the device into this stream | [optional] 

## Methods

### NewDeviceStreamsResponseWeb

`func NewDeviceStreamsResponseWeb() *DeviceStreamsResponseWeb`

NewDeviceStreamsResponseWeb instantiates a new DeviceStreamsResponseWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceStreamsResponseWebWithDefaults

`func NewDeviceStreamsResponseWebWithDefaults() *DeviceStreamsResponseWeb`

NewDeviceStreamsResponseWebWithDefaults instantiates a new DeviceStreamsResponseWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStreamId

`func (o *DeviceStreamsResponseWeb) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *DeviceStreamsResponseWeb) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *DeviceStreamsResponseWeb) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *DeviceStreamsResponseWeb) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetCount

`func (o *DeviceStreamsResponseWeb) GetCount() int32`

GetCount returns the Count field if non-nil, zero value otherwise.

### GetCountOk

`func (o *DeviceStreamsResponseWeb) GetCountOk() (*int32, bool)`

GetCountOk returns a tuple with the Count field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCount

`func (o *DeviceStreamsResponseWeb) SetCount(v int32)`

SetCount sets Count field to given value.

### HasCount

`func (o *DeviceStreamsResponseWeb) HasCount() bool`

HasCount returns a boolean if a field has been set.

### GetLastUpdate

`func (o *DeviceStreamsResponseWeb) GetLastUpdate() string`

GetLastUpdate returns the LastUpdate field if non-nil, zero value otherwise.

### GetLastUpdateOk

`func (o *DeviceStreamsResponseWeb) GetLastUpdateOk() (*string, bool)`

GetLastUpdateOk returns a tuple with the LastUpdate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdate

`func (o *DeviceStreamsResponseWeb) SetLastUpdate(v string)`

SetLastUpdate sets LastUpdate field to given value.

### HasLastUpdate

`func (o *DeviceStreamsResponseWeb) HasLastUpdate() bool`

HasLastUpdate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


