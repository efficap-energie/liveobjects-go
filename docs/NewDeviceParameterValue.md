# NewDeviceParameterValue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | Pointer to **string** | parameter value type (INT32, UINT32, FLOAT, STRING or BINARY) | [optional] 
**Value** | Pointer to **map[string]interface{}** | configuration parameter value (number for INT32/UINT32 type, string for STRING type,float for FLOAT type, base64-encoded string for BINARY type)  | [optional] 

## Methods

### NewNewDeviceParameterValue

`func NewNewDeviceParameterValue() *NewDeviceParameterValue`

NewNewDeviceParameterValue instantiates a new NewDeviceParameterValue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDeviceParameterValueWithDefaults

`func NewNewDeviceParameterValueWithDefaults() *NewDeviceParameterValue`

NewNewDeviceParameterValueWithDefaults instantiates a new NewDeviceParameterValue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *NewDeviceParameterValue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *NewDeviceParameterValue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *NewDeviceParameterValue) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *NewDeviceParameterValue) HasType() bool`

HasType returns a boolean if a field has been set.

### GetValue

`func (o *NewDeviceParameterValue) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *NewDeviceParameterValue) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *NewDeviceParameterValue) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *NewDeviceParameterValue) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


