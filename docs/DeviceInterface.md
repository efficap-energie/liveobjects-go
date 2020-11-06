# DeviceInterface

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | Pointer to **string** | Interface connector ID | [optional] 
**NodeId** | Pointer to **string** | Interface node ID | [optional] 
**DeviceId** | Pointer to **string** | Device identifier (URN) | [optional] 
**Enabled** | Pointer to **bool** | Indicates whether the interface is enabled | [optional] 
**Status** | Pointer to **string** | Interface status | [optional] 
**Definition** | Pointer to **map[string]interface{}** | Base definition | [optional] 
**LastContact** | Pointer to **string** | Last contact date | [optional] 
**Capabilities** | Pointer to [**InterfaceCapabilities**](InterfaceCapabilities.md) |  | [optional] 
**Activity** | Pointer to **map[string]interface{}** | Interface activity | [optional] 
**Created** | Pointer to **string** | Date/time of the device creation | [optional] [readonly] 
**Updated** | Pointer to **string** | Date/time of the device update | [optional] [readonly] 

## Methods

### NewDeviceInterface

`func NewDeviceInterface() *DeviceInterface`

NewDeviceInterface instantiates a new DeviceInterface object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceInterfaceWithDefaults

`func NewDeviceInterfaceWithDefaults() *DeviceInterface`

NewDeviceInterfaceWithDefaults instantiates a new DeviceInterface object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *DeviceInterface) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *DeviceInterface) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *DeviceInterface) SetConnector(v string)`

SetConnector sets Connector field to given value.

### HasConnector

`func (o *DeviceInterface) HasConnector() bool`

HasConnector returns a boolean if a field has been set.

### GetNodeId

`func (o *DeviceInterface) GetNodeId() string`

GetNodeId returns the NodeId field if non-nil, zero value otherwise.

### GetNodeIdOk

`func (o *DeviceInterface) GetNodeIdOk() (*string, bool)`

GetNodeIdOk returns a tuple with the NodeId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeId

`func (o *DeviceInterface) SetNodeId(v string)`

SetNodeId sets NodeId field to given value.

### HasNodeId

`func (o *DeviceInterface) HasNodeId() bool`

HasNodeId returns a boolean if a field has been set.

### GetDeviceId

`func (o *DeviceInterface) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *DeviceInterface) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *DeviceInterface) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *DeviceInterface) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetEnabled

`func (o *DeviceInterface) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *DeviceInterface) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *DeviceInterface) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *DeviceInterface) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetStatus

`func (o *DeviceInterface) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *DeviceInterface) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *DeviceInterface) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *DeviceInterface) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDefinition

`func (o *DeviceInterface) GetDefinition() map[string]interface{}`

GetDefinition returns the Definition field if non-nil, zero value otherwise.

### GetDefinitionOk

`func (o *DeviceInterface) GetDefinitionOk() (*map[string]interface{}, bool)`

GetDefinitionOk returns a tuple with the Definition field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefinition

`func (o *DeviceInterface) SetDefinition(v map[string]interface{})`

SetDefinition sets Definition field to given value.

### HasDefinition

`func (o *DeviceInterface) HasDefinition() bool`

HasDefinition returns a boolean if a field has been set.

### GetLastContact

`func (o *DeviceInterface) GetLastContact() string`

GetLastContact returns the LastContact field if non-nil, zero value otherwise.

### GetLastContactOk

`func (o *DeviceInterface) GetLastContactOk() (*string, bool)`

GetLastContactOk returns a tuple with the LastContact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastContact

`func (o *DeviceInterface) SetLastContact(v string)`

SetLastContact sets LastContact field to given value.

### HasLastContact

`func (o *DeviceInterface) HasLastContact() bool`

HasLastContact returns a boolean if a field has been set.

### GetCapabilities

`func (o *DeviceInterface) GetCapabilities() InterfaceCapabilities`

GetCapabilities returns the Capabilities field if non-nil, zero value otherwise.

### GetCapabilitiesOk

`func (o *DeviceInterface) GetCapabilitiesOk() (*InterfaceCapabilities, bool)`

GetCapabilitiesOk returns a tuple with the Capabilities field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCapabilities

`func (o *DeviceInterface) SetCapabilities(v InterfaceCapabilities)`

SetCapabilities sets Capabilities field to given value.

### HasCapabilities

`func (o *DeviceInterface) HasCapabilities() bool`

HasCapabilities returns a boolean if a field has been set.

### GetActivity

`func (o *DeviceInterface) GetActivity() map[string]interface{}`

GetActivity returns the Activity field if non-nil, zero value otherwise.

### GetActivityOk

`func (o *DeviceInterface) GetActivityOk() (*map[string]interface{}, bool)`

GetActivityOk returns a tuple with the Activity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivity

`func (o *DeviceInterface) SetActivity(v map[string]interface{})`

SetActivity sets Activity field to given value.

### HasActivity

`func (o *DeviceInterface) HasActivity() bool`

HasActivity returns a boolean if a field has been set.

### GetCreated

`func (o *DeviceInterface) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DeviceInterface) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DeviceInterface) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *DeviceInterface) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *DeviceInterface) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *DeviceInterface) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *DeviceInterface) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *DeviceInterface) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


