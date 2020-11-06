# ActionTriggers

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandStatus** | Pointer to [**CommandStatusTrigger**](CommandStatusTrigger.md) |  | [optional] 
**DataMessage** | Pointer to [**DataMessageTrigger**](DataMessageTrigger.md) |  | [optional] 
**DeviceActivity** | Pointer to [**DeviceActivityTrigger**](DeviceActivityTrigger.md) |  | [optional] 
**DeviceCreated** | Pointer to [**DeviceCreatedTrigger**](DeviceCreatedTrigger.md) |  | [optional] 
**DeviceDeleted** | Pointer to [**DeviceDeletedTrigger**](DeviceDeletedTrigger.md) |  | [optional] 
**DeviceStatus** | Pointer to [**DeviceStatusTrigger**](DeviceStatusTrigger.md) |  | [optional] 
**LoraNetwork** | Pointer to [**LoraNetworkTrigger**](LoraNetworkTrigger.md) |  | [optional] 
**MatchingFired** | Pointer to [**MatchingFiredTrigger**](MatchingFiredTrigger.md) |  | [optional] 
**StateChange** | Pointer to [**StateChangeTrigger**](StateChangeTrigger.md) |  | [optional] 

## Methods

### NewActionTriggers

`func NewActionTriggers() *ActionTriggers`

NewActionTriggers instantiates a new ActionTriggers object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTriggersWithDefaults

`func NewActionTriggersWithDefaults() *ActionTriggers`

NewActionTriggersWithDefaults instantiates a new ActionTriggers object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandStatus

`func (o *ActionTriggers) GetCommandStatus() CommandStatusTrigger`

GetCommandStatus returns the CommandStatus field if non-nil, zero value otherwise.

### GetCommandStatusOk

`func (o *ActionTriggers) GetCommandStatusOk() (*CommandStatusTrigger, bool)`

GetCommandStatusOk returns a tuple with the CommandStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandStatus

`func (o *ActionTriggers) SetCommandStatus(v CommandStatusTrigger)`

SetCommandStatus sets CommandStatus field to given value.

### HasCommandStatus

`func (o *ActionTriggers) HasCommandStatus() bool`

HasCommandStatus returns a boolean if a field has been set.

### GetDataMessage

`func (o *ActionTriggers) GetDataMessage() DataMessageTrigger`

GetDataMessage returns the DataMessage field if non-nil, zero value otherwise.

### GetDataMessageOk

`func (o *ActionTriggers) GetDataMessageOk() (*DataMessageTrigger, bool)`

GetDataMessageOk returns a tuple with the DataMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataMessage

`func (o *ActionTriggers) SetDataMessage(v DataMessageTrigger)`

SetDataMessage sets DataMessage field to given value.

### HasDataMessage

`func (o *ActionTriggers) HasDataMessage() bool`

HasDataMessage returns a boolean if a field has been set.

### GetDeviceActivity

`func (o *ActionTriggers) GetDeviceActivity() DeviceActivityTrigger`

GetDeviceActivity returns the DeviceActivity field if non-nil, zero value otherwise.

### GetDeviceActivityOk

`func (o *ActionTriggers) GetDeviceActivityOk() (*DeviceActivityTrigger, bool)`

GetDeviceActivityOk returns a tuple with the DeviceActivity field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceActivity

`func (o *ActionTriggers) SetDeviceActivity(v DeviceActivityTrigger)`

SetDeviceActivity sets DeviceActivity field to given value.

### HasDeviceActivity

`func (o *ActionTriggers) HasDeviceActivity() bool`

HasDeviceActivity returns a boolean if a field has been set.

### GetDeviceCreated

`func (o *ActionTriggers) GetDeviceCreated() DeviceCreatedTrigger`

GetDeviceCreated returns the DeviceCreated field if non-nil, zero value otherwise.

### GetDeviceCreatedOk

`func (o *ActionTriggers) GetDeviceCreatedOk() (*DeviceCreatedTrigger, bool)`

GetDeviceCreatedOk returns a tuple with the DeviceCreated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCreated

`func (o *ActionTriggers) SetDeviceCreated(v DeviceCreatedTrigger)`

SetDeviceCreated sets DeviceCreated field to given value.

### HasDeviceCreated

`func (o *ActionTriggers) HasDeviceCreated() bool`

HasDeviceCreated returns a boolean if a field has been set.

### GetDeviceDeleted

`func (o *ActionTriggers) GetDeviceDeleted() DeviceDeletedTrigger`

GetDeviceDeleted returns the DeviceDeleted field if non-nil, zero value otherwise.

### GetDeviceDeletedOk

`func (o *ActionTriggers) GetDeviceDeletedOk() (*DeviceDeletedTrigger, bool)`

GetDeviceDeletedOk returns a tuple with the DeviceDeleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceDeleted

`func (o *ActionTriggers) SetDeviceDeleted(v DeviceDeletedTrigger)`

SetDeviceDeleted sets DeviceDeleted field to given value.

### HasDeviceDeleted

`func (o *ActionTriggers) HasDeviceDeleted() bool`

HasDeviceDeleted returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *ActionTriggers) GetDeviceStatus() DeviceStatusTrigger`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *ActionTriggers) GetDeviceStatusOk() (*DeviceStatusTrigger, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *ActionTriggers) SetDeviceStatus(v DeviceStatusTrigger)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *ActionTriggers) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetLoraNetwork

`func (o *ActionTriggers) GetLoraNetwork() LoraNetworkTrigger`

GetLoraNetwork returns the LoraNetwork field if non-nil, zero value otherwise.

### GetLoraNetworkOk

`func (o *ActionTriggers) GetLoraNetworkOk() (*LoraNetworkTrigger, bool)`

GetLoraNetworkOk returns a tuple with the LoraNetwork field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoraNetwork

`func (o *ActionTriggers) SetLoraNetwork(v LoraNetworkTrigger)`

SetLoraNetwork sets LoraNetwork field to given value.

### HasLoraNetwork

`func (o *ActionTriggers) HasLoraNetwork() bool`

HasLoraNetwork returns a boolean if a field has been set.

### GetMatchingFired

`func (o *ActionTriggers) GetMatchingFired() MatchingFiredTrigger`

GetMatchingFired returns the MatchingFired field if non-nil, zero value otherwise.

### GetMatchingFiredOk

`func (o *ActionTriggers) GetMatchingFiredOk() (*MatchingFiredTrigger, bool)`

GetMatchingFiredOk returns a tuple with the MatchingFired field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingFired

`func (o *ActionTriggers) SetMatchingFired(v MatchingFiredTrigger)`

SetMatchingFired sets MatchingFired field to given value.

### HasMatchingFired

`func (o *ActionTriggers) HasMatchingFired() bool`

HasMatchingFired returns a boolean if a field has been set.

### GetStateChange

`func (o *ActionTriggers) GetStateChange() StateChangeTrigger`

GetStateChange returns the StateChange field if non-nil, zero value otherwise.

### GetStateChangeOk

`func (o *ActionTriggers) GetStateChangeOk() (*StateChangeTrigger, bool)`

GetStateChangeOk returns a tuple with the StateChange field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateChange

`func (o *ActionTriggers) SetStateChange(v StateChangeTrigger)`

SetStateChange sets StateChange field to given value.

### HasStateChange

`func (o *ActionTriggers) HasStateChange() bool`

HasStateChange returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


