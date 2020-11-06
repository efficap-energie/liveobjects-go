# NewDeviceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | Connector ID. A connector must respect the following regular expression &lt;code&gt;([\\w\\-_]{1,128})&lt;/code&gt; (max 128 characters) | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the interface is enabled | [optional] 
**Definition** | Pointer to **map[string]interface{}** | Base definition. Expected string (max 10000 characters) | [optional] 

## Methods

### NewNewDeviceInterface

`func NewNewDeviceInterface() *NewDeviceInterface`

NewNewDeviceInterface instantiates a new NewDeviceInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewNewDeviceInterfaceWithDefaults

`func NewNewDeviceInterfaceWithDefaults() *NewDeviceInterface`

NewNewDeviceInterfaceWithDefaults instantiates a new NewDeviceInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *NewDeviceInterface) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *NewDeviceInterface) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *NewDeviceInterface) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *NewDeviceInterface) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetEnabled

`func (o *NewDeviceInterface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *NewDeviceInterface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *NewDeviceInterface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *NewDeviceInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetDefinition

`func (o *NewDeviceInterface) GetDefinition() map[string]interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *NewDeviceInterface) GetDefinitionOk() (*map[string]interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *NewDeviceInterface) SetDefinition(v map[string]interface{})`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *NewDeviceInterface) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


