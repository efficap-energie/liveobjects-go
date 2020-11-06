# DeviceConfigWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**map[string]DeviceParameterWeb**](DeviceParameterWeb.md) | device configuration | [optional] 
**NotifyTo** | Pointer to **string** | topic where configuration change/sync events must be published to | [optional] 

## Methods

### NewDeviceConfigWeb

`func NewDeviceConfigWeb() *DeviceConfigWeb`

NewDeviceConfigWeb instantiates a new DeviceConfigWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceConfigWebWithDefaults

`func NewDeviceConfigWebWithDefaults() *DeviceConfigWeb`

NewDeviceConfigWebWithDefaults instantiates a new DeviceConfigWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *DeviceConfigWeb) GetParameters() map[string]DeviceParameterWeb`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DeviceConfigWeb) GetParametersOk() (*map[string]DeviceParameterWeb, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DeviceConfigWeb) SetParameters(v map[string]DeviceParameterWeb)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DeviceConfigWeb) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetNotifyTo

`func (o *DeviceConfigWeb) GetNotifyTo() string`

GetNotifyTo returns the NotifyTo field if non-nil, zero value otherwise.

### GetNotifyToOk

`func (o *DeviceConfigWeb) GetNotifyToOk() (*string, bool)`

GetNotifyToOk returns a tuple with the NotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTo

`func (o *DeviceConfigWeb) SetNotifyTo(v string)`

SetNotifyTo sets NotifyTo field to given value.

### HasNotifyTo

`func (o *DeviceConfigWeb) HasNotifyTo() bool`

HasNotifyTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


