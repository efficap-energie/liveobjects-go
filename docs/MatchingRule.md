# MatchingRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DataPredicate** | Pointer to **map[string]interface{}** | the JsonLogic (http://jsonlogic.com/) pattern matching that will trigger an event for each new data that match this predicate. This JsonLogic predicate needs to return a boolean. | [optional] 
**Enabled** | Pointer to **bool** | activate or not the MatchingRule. Default is false. | [optional] 
**Id** | Pointer to **string** | id of the MatchingRule. Should be null when used for POST. | [optional] 
**Name** | **string** | name of the MatchingRule. Must be unique. | 

## Methods

### NewMatchingRule

`func NewMatchingRule(name string, ) *MatchingRule`

NewMatchingRule instantiates a new MatchingRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchingRuleWithDefaults

`func NewMatchingRuleWithDefaults() *MatchingRule`

NewMatchingRuleWithDefaults instantiates a new MatchingRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDataPredicate

`func (o *MatchingRule) GetDataPredicate() map[string]interface{}`

GetDataPredicate returns the DataPredicate field if non-nil, zero value otherwise.

### GetDataPredicateOk

`func (o *MatchingRule) GetDataPredicateOk() (*map[string]interface{}, bool)`

GetDataPredicateOk returns a tuple with the DataPredicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPredicate

`func (o *MatchingRule) SetDataPredicate(v map[string]interface{})`

SetDataPredicate sets DataPredicate field to given value.

### HasDataPredicate

`func (o *MatchingRule) HasDataPredicate() bool`

HasDataPredicate returns a boolean if a field has been set.

### GetEnabled

`func (o *MatchingRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *MatchingRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *MatchingRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *MatchingRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *MatchingRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *MatchingRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *MatchingRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *MatchingRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *MatchingRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MatchingRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MatchingRule) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


