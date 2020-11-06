# DeviceParameterValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | configuration parameter value type (INT32, UINT32, FLOAT, STRING or BINARY) | [optional] 
**Value** | Pointer to **map[string]interface{}** | configuration parameter value (number for INT32/UINT32 type, string for STRING type,float for FLOAT type, base64-encoded string for BINARY type)  | [optional] 
**Timestamp** | Pointer to **string** | configuration parameter value associated date/time (ISO 8601) | [optional] 

## Methods

### NewDeviceParameterValue

`func NewDeviceParameterValue() *DeviceParameterValue`

NewDeviceParameterValue instantiates a new DeviceParameterValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceParameterValueWithDefaults

`func NewDeviceParameterValueWithDefaults() *DeviceParameterValue`

NewDeviceParameterValueWithDefaults instantiates a new DeviceParameterValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DeviceParameterValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DeviceParameterValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DeviceParameterValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *DeviceParameterValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *DeviceParameterValue) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DeviceParameterValue) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DeviceParameterValue) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *DeviceParameterValue) HasValue() bool`

HasValue returns a boolean if a field has been set.

### GetTimestamp

`func (o *DeviceParameterValue) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DeviceParameterValue) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DeviceParameterValue) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DeviceParameterValue) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


