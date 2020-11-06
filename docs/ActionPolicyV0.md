# ActionPolicyV0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Actions** | Pointer to [**Actions**](Actions.md) |  | [optional] 
**Enabled** | Pointer to **bool** | Enable or disable this action policy | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Name** | Pointer to **string** | The user-defined name of the action policy | [optional] 
**Triggers** | Pointer to [**ActionTriggersV0**](ActionTriggersV0.md) |  | [optional] 

## Methods

### NewActionPolicyV0

`func NewActionPolicyV0() *ActionPolicyV0`

NewActionPolicyV0 instantiates a new ActionPolicyV0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionPolicyV0WithDefaults

`func NewActionPolicyV0WithDefaults() *ActionPolicyV0`

NewActionPolicyV0WithDefaults instantiates a new ActionPolicyV0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActions

`func (o *ActionPolicyV0) GetActions() Actions`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *ActionPolicyV0) GetActionsOk() (*Actions, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *ActionPolicyV0) SetActions(v Actions)`

SetActions sets Actions field to given value.

### HasActions

`func (o *ActionPolicyV0) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetEnabled

`func (o *ActionPolicyV0) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActionPolicyV0) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActionPolicyV0) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActionPolicyV0) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ActionPolicyV0) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActionPolicyV0) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActionPolicyV0) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActionPolicyV0) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActionPolicyV0) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActionPolicyV0) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActionPolicyV0) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *ActionPolicyV0) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTriggers

`func (o *ActionPolicyV0) GetTriggers() ActionTriggersV0`

GetTriggers returns the Triggers field if non-nil, zero value otherwise.

### GetTriggersOk

`func (o *ActionPolicyV0) GetTriggersOk() (*ActionTriggersV0, bool)`

GetTriggersOk returns a tuple with the Triggers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTriggers

`func (o *ActionPolicyV0) SetTriggers(v ActionTriggersV0)`

SetTriggers sets Triggers field to given value.

### HasTriggers

`func (o *ActionPolicyV0) HasTriggers() bool`

HasTriggers returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


