# UpdateConnectorNodeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | New device id this connector node must be associated with (or null) | [optional] 
**Enabled** | Pointer to **bool** | New enabled state for this connector node (or null) | [optional] 
**Definition** | Pointer to **map[string]interface{}** | New interface definition | [optional] 

## Methods

### NewUpdateConnectorNodeRequest

`func NewUpdateConnectorNodeRequest() *UpdateConnectorNodeRequest`

NewUpdateConnectorNodeRequest instantiates a new UpdateConnectorNodeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateConnectorNodeRequestWithDefaults

`func NewUpdateConnectorNodeRequestWithDefaults() *UpdateConnectorNodeRequest`

NewUpdateConnectorNodeRequestWithDefaults instantiates a new UpdateConnectorNodeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *UpdateConnectorNodeRequest) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UpdateConnectorNodeRequest) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UpdateConnectorNodeRequest) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *UpdateConnectorNodeRequest) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEnabled

`func (o *UpdateConnectorNodeRequest) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UpdateConnectorNodeRequest) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UpdateConnectorNodeRequest) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UpdateConnectorNodeRequest) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefinition

`func (o *UpdateConnectorNodeRequest) GetDefinition() map[string]interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *UpdateConnectorNodeRequest) GetDefinitionOk() (*map[string]interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *UpdateConnectorNodeRequest) SetDefinition(v map[string]interface{})`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *UpdateConnectorNodeRequest) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


