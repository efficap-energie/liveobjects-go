# StateChangeFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RuleIds** | Pointer to **[]string** | list of filtered rule Ids. Null or empty to accept all rule Ids | [optional] 

## Methods

### NewStateChangeFilter

`func NewStateChangeFilter() *StateChangeFilter`

NewStateChangeFilter instantiates a new StateChangeFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateChangeFilterWithDefaults

`func NewStateChangeFilterWithDefaults() *StateChangeFilter`

NewStateChangeFilterWithDefaults instantiates a new StateChangeFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRuleIds

`func (o *StateChangeFilter) GetRuleIds() []string`

GetRuleIds returns the RuleIds field if non-nil, zero value otherwise.

### GetRuleIdsOk

`func (o *StateChangeFilter) GetRuleIdsOk() (*[]string, bool)`

GetRuleIdsOk returns a tuple with the RuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRuleIds

`func (o *StateChangeFilter) SetRuleIds(v []string)`

SetRuleIds sets RuleIds field to given value.

### HasRuleIds

`func (o *StateChangeFilter) HasRuleIds() bool`

HasRuleIds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


