# ActionTriggersV0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EventRuleIds** | Pointer to **[]string** | A list of state processing, firing or activity processing rule ids that will trigger an action - incompatible with messageSelector | [optional] 
**MessageSelector** | Pointer to [**MessageSelector**](MessageSelector.md) |  | [optional] 

## Methods

### NewActionTriggersV0

`func NewActionTriggersV0() *ActionTriggersV0`

NewActionTriggersV0 instantiates a new ActionTriggersV0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionTriggersV0WithDefaults

`func NewActionTriggersV0WithDefaults() *ActionTriggersV0`

NewActionTriggersV0WithDefaults instantiates a new ActionTriggersV0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEventRuleIds

`func (o *ActionTriggersV0) GetEventRuleIds() []string`

GetEventRuleIds returns the EventRuleIds field if non-nil, zero value otherwise.

### GetEventRuleIdsOk

`func (o *ActionTriggersV0) GetEventRuleIdsOk() (*[]string, bool)`

GetEventRuleIdsOk returns a tuple with the EventRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventRuleIds

`func (o *ActionTriggersV0) SetEventRuleIds(v []string)`

SetEventRuleIds sets EventRuleIds field to given value.

### HasEventRuleIds

`func (o *ActionTriggersV0) HasEventRuleIds() bool`

HasEventRuleIds returns a boolean if a field has been set.

### GetMessageSelector

`func (o *ActionTriggersV0) GetMessageSelector() MessageSelector`

GetMessageSelector returns the MessageSelector field if non-nil, zero value otherwise.

### GetMessageSelectorOk

`func (o *ActionTriggersV0) GetMessageSelectorOk() (*MessageSelector, bool)`

GetMessageSelectorOk returns a tuple with the MessageSelector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageSelector

`func (o *ActionTriggersV0) SetMessageSelector(v MessageSelector)`

SetMessageSelector sets MessageSelector field to given value.

### HasMessageSelector

`func (o *ActionTriggersV0) HasMessageSelector() bool`

HasMessageSelector returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


