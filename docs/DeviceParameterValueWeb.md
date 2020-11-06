# DeviceParameterValueWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | configuration parameter value type (INT32, UINT32, FLOAT, STRING or BINARY) | [optional] 
**Value** | Pointer to **map[string]interface{}** | configuration parameter value (number for INT32/UINT32 type, string for STRING type,float for FLOAT type, base64-encoded string for BINARY type)  | [optional] 
**Timestamp** | Pointer to **string** | configuration parameter value associated date/time (ISO 8601) | [optional] 

## Methods

### NewDeviceParameterValueWeb

`func NewDeviceParameterValueWeb() *DeviceParameterValueWeb`

NewDeviceParameterValueWeb instantiates a new DeviceParameterValueWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceParameterValueWebWithDefaults

`func NewDeviceParameterValueWebWithDefaults() *DeviceParameterValueWeb`

NewDeviceParameterValueWebWithDefaults instantiates a new DeviceParameterValueWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeviceParameterValueWeb) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceParameterValueWeb) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceParameterValueWeb) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceParameterValueWeb) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DeviceParameterValueWeb) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeviceParameterValueWeb) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeviceParameterValueWeb) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DeviceParameterValueWeb) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *DeviceParameterValueWeb) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeviceParameterValueWeb) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeviceParameterValueWeb) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DeviceParameterValueWeb) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


