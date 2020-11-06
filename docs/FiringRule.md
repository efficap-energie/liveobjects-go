# FiringRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AggregationKeys** | Pointer to **[]string** | the list of jsonPath in the Data that will define on which group of data this FiringRule should be set. For instance &#39;streamId&#39;, &#39;metadata.source&#39;, &#39;value.type&#39;. | [optional] 
**Enabled** | Pointer to **bool** | activate or not the FiringRule. Default is false. | [optional] 
**FiringType** | **string** | define the type of firing mechanism : ONCE, SLEEP, or ALWAYS | 
**Id** | Pointer to **string** | id of the FiringRule. Should be null when used for POST. | [optional] 
**MatchingRuleIds** | Pointer to **[]string** | the list of MatchingRule ids that will be handeld by this FiringRule. | [optional] 
**Name** | **string** | name of the FiringRule. Must be unique. | 
**SleepDuration** | Pointer to **string** | sleep duration of the FiringRule. Is defined as a ISO-8601 Period string, restricted to days, hours, minutes and seconds. 1 day will always be equivalent to 24H, regardless of daylight saving time. eg. : &#39;P1D&#39;, &#39;PT30M&#39;. Must be set only for &#39;SLEEP&#39; FiringType. | [optional] 

## Methods

### NewFiringRule

`func NewFiringRule(firingType string, name string, ) *FiringRule`

NewFiringRule instantiates a new FiringRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiringRuleWithDefaults

`func NewFiringRuleWithDefaults() *FiringRule`

NewFiringRuleWithDefaults instantiates a new FiringRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAggregationKeys

`func (o *FiringRule) GetAggregationKeys() []string`

GetAggregationKeys returns the AggregationKeys field if non-nil, zero value otherwise.

### GetAggregationKeysOk

`func (o *FiringRule) GetAggregationKeysOk() (*[]string, bool)`

GetAggregationKeysOk returns a tuple with the AggregationKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAggregationKeys

`func (o *FiringRule) SetAggregationKeys(v []string)`

SetAggregationKeys sets AggregationKeys field to given value.

### HasAggregationKeys

`func (o *FiringRule) HasAggregationKeys() bool`

HasAggregationKeys returns a boolean if a field has been set.

### GetEnabled

`func (o *FiringRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *FiringRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *FiringRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *FiringRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFiringType

`func (o *FiringRule) GetFiringType() string`

GetFiringType returns the FiringType field if non-nil, zero value otherwise.

### GetFiringTypeOk

`func (o *FiringRule) GetFiringTypeOk() (*string, bool)`

GetFiringTypeOk returns a tuple with the FiringType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiringType

`func (o *FiringRule) SetFiringType(v string)`

SetFiringType sets FiringType field to given value.


### GetId

`func (o *FiringRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiringRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiringRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FiringRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMatchingRuleIds

`func (o *FiringRule) GetMatchingRuleIds() []string`

GetMatchingRuleIds returns the MatchingRuleIds field if non-nil, zero value otherwise.

### GetMatchingRuleIdsOk

`func (o *FiringRule) GetMatchingRuleIdsOk() (*[]string, bool)`

GetMatchingRuleIdsOk returns a tuple with the MatchingRuleIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRuleIds

`func (o *FiringRule) SetMatchingRuleIds(v []string)`

SetMatchingRuleIds sets MatchingRuleIds field to given value.

### HasMatchingRuleIds

`func (o *FiringRule) HasMatchingRuleIds() bool`

HasMatchingRuleIds returns a boolean if a field has been set.

### GetName

`func (o *FiringRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FiringRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FiringRule) SetName(v string)`

SetName sets Name field to given value.


### GetSleepDuration

`func (o *FiringRule) GetSleepDuration() string`

GetSleepDuration returns the SleepDuration field if non-nil, zero value otherwise.

### GetSleepDurationOk

`func (o *FiringRule) GetSleepDurationOk() (*string, bool)`

GetSleepDurationOk returns a tuple with the SleepDuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSleepDuration

`func (o *FiringRule) SetSleepDuration(v string)`

SetSleepDuration sets SleepDuration field to given value.

### HasSleepDuration

`func (o *FiringRule) HasSleepDuration() bool`

HasSleepDuration returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


