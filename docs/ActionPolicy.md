# ActionPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | [**Actions**](Actions.md) |  | 
**Enabled** | Pointer to **bool** | Enable or disable this action policy | [optional] 
**Id** | Pointer to **string** | Id of the action policy. Set by LiveObjects, should be null on POST. | [optional] 
**Name** | **string** | The user-defined name of the action policy | 
**Triggers** | [**ActionTriggers**](ActionTriggers.md) |  | 

## Methods

### NewActionPolicy

`func NewActionPolicy(actions Actions, name string, triggers ActionTriggers, ) *ActionPolicy`

NewActionPolicy instantiates a new ActionPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPolicyWithDefaults

`func NewActionPolicyWithDefaults() *ActionPolicy`

NewActionPolicyWithDefaults instantiates a new ActionPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ActionPolicy) GetActions() Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionPolicy) GetActionsOk() (*Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionPolicy) SetActions(v Actions)`

SetActions sets Actions field to given value.


### GetEnabled

`func (o *ActionPolicy) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionPolicy) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionPolicy) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActionPolicy) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ActionPolicy) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPolicy) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPolicy) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionPolicy) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActionPolicy) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionPolicy) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionPolicy) SetName(v string)`

SetName sets Name field to given value.


### GetTriggers

`func (o *ActionPolicy) GetTriggers() ActionTriggers`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ActionPolicy) GetTriggersOk() (*ActionTriggers, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ActionPolicy) SetTriggers(v ActionTriggers)`

SetTriggers sets Triggers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


