# StateProcessingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | activate or not the StateProcessingRule. Default is false. | [optional] 
**FilterPredicate** | Pointer to **map[string]interface{}** | the JsonLogic (http://jsonlogic.com/) filter predicate applied before state function | [optional] 
**Id** | Pointer to **string** | id of the StateProcessingRule. Should be null when used for POST. | [optional] 
**Name** | **string** | name of the StateProcessingRule. Must be unique. | 
**StateFunction** | Pointer to **map[string]interface{}** | the JsonLogic (http://jsonlogic.com/) state function | [optional] 
**StateKeyPath** | Pointer to **string** | the json path applied on dataMessage used to compute the stateKey | [optional] 

## Methods

### NewStateProcessingRule

`func NewStateProcessingRule(name string, ) *StateProcessingRule`

NewStateProcessingRule instantiates a new StateProcessingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateProcessingRuleWithDefaults

`func NewStateProcessingRuleWithDefaults() *StateProcessingRule`

NewStateProcessingRuleWithDefaults instantiates a new StateProcessingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *StateProcessingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *StateProcessingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *StateProcessingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *StateProcessingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFilterPredicate

`func (o *StateProcessingRule) GetFilterPredicate() map[string]interface{}`

GetFilterPredicate returns the FilterPredicate field if non-nil, zero value otherwise.

### GetFilterPredicateOk

`func (o *StateProcessingRule) GetFilterPredicateOk() (*map[string]interface{}, bool)`

GetFilterPredicateOk returns a tuple with the FilterPredicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterPredicate

`func (o *StateProcessingRule) SetFilterPredicate(v map[string]interface{})`

SetFilterPredicate sets FilterPredicate field to given value.

### HasFilterPredicate

`func (o *StateProcessingRule) HasFilterPredicate() bool`

HasFilterPredicate returns a boolean if a field has been set.

### GetId

`func (o *StateProcessingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StateProcessingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StateProcessingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StateProcessingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *StateProcessingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *StateProcessingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *StateProcessingRule) SetName(v string)`

SetName sets Name field to given value.


### GetStateFunction

`func (o *StateProcessingRule) GetStateFunction() map[string]interface{}`

GetStateFunction returns the StateFunction field if non-nil, zero value otherwise.

### GetStateFunctionOk

`func (o *StateProcessingRule) GetStateFunctionOk() (*map[string]interface{}, bool)`

GetStateFunctionOk returns a tuple with the StateFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateFunction

`func (o *StateProcessingRule) SetStateFunction(v map[string]interface{})`

SetStateFunction sets StateFunction field to given value.

### HasStateFunction

`func (o *StateProcessingRule) HasStateFunction() bool`

HasStateFunction returns a boolean if a field has been set.

### GetStateKeyPath

`func (o *StateProcessingRule) GetStateKeyPath() string`

GetStateKeyPath returns the StateKeyPath field if non-nil, zero value otherwise.

### GetStateKeyPathOk

`func (o *StateProcessingRule) GetStateKeyPathOk() (*string, bool)`

GetStateKeyPathOk returns a tuple with the StateKeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateKeyPath

`func (o *StateProcessingRule) SetStateKeyPath(v string)`

SetStateKeyPath sets StateKeyPath field to given value.

### HasStateKeyPath

`func (o *StateProcessingRule) HasStateKeyPath() bool`

HasStateKeyPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


