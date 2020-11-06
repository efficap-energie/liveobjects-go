# DeviceParametersSetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Parameters** | Pointer to [**map[string]NewDeviceParameterValue**](NewDeviceParameterValue.md) | Device configuration update. Max number of parameters is 100. Parameter name max length is 128. Parameter value must be valid according to the type and max length is 2000. | [optional] 

## Methods

### NewDeviceParametersSetRequest

`func NewDeviceParametersSetRequest() *DeviceParametersSetRequest`

NewDeviceParametersSetRequest instantiates a new DeviceParametersSetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceParametersSetRequestWithDefaults

`func NewDeviceParametersSetRequestWithDefaults() *DeviceParametersSetRequest`

NewDeviceParametersSetRequestWithDefaults instantiates a new DeviceParametersSetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParameters

`func (o *DeviceParametersSetRequest) GetParameters() map[string]NewDeviceParameterValue`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DeviceParametersSetRequest) GetParametersOk() (*map[string]NewDeviceParameterValue, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DeviceParametersSetRequest) SetParameters(v map[string]NewDeviceParameterValue)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DeviceParametersSetRequest) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


