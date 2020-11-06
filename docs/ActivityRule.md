# ActivityRule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** | activate or not the ActivityRule. Default is false. | [optional] 
**Id** | Pointer to **string** | id of the ActivityRule. Should be null when used for POST. | [optional] 
**Name** | **string** | user-defined name for the rule, must be unique and not empty. | 
**SilentPolicy** | [**SilentPolicy**](SilentPolicy.md) |  | 
**Tags** | Pointer to **[]string** | list of tags that will be set on alarms triggered by this rule. e.g. &#39;level:HIGH&#39; | [optional] 
**Targets** | [**Targets**](Targets.md) |  | 

## Methods

### NewActivityRule

`func NewActivityRule(name string, silentPolicy SilentPolicy, targets Targets, ) *ActivityRule`

NewActivityRule instantiates a new ActivityRule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActivityRuleWithDefaults

`func NewActivityRuleWithDefaults() *ActivityRule`

NewActivityRuleWithDefaults instantiates a new ActivityRule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *ActivityRule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *ActivityRule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *ActivityRule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *ActivityRule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *ActivityRule) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ActivityRule) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ActivityRule) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ActivityRule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *ActivityRule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *ActivityRule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *ActivityRule) SetName(v string)`

SetName sets Name field to given value.


### GetSilentPolicy

`func (o *ActivityRule) GetSilentPolicy() SilentPolicy`

GetSilentPolicy returns the SilentPolicy field if non-nil, zero value otherwise.

### GetSilentPolicyOk

`func (o *ActivityRule) GetSilentPolicyOk() (*SilentPolicy, bool)`

GetSilentPolicyOk returns a tuple with the SilentPolicy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSilentPolicy

`func (o *ActivityRule) SetSilentPolicy(v SilentPolicy)`

SetSilentPolicy sets SilentPolicy field to given value.


### GetTags

`func (o *ActivityRule) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ActivityRule) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ActivityRule) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ActivityRule) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTargets

`func (o *ActivityRule) GetTargets() Targets`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *ActivityRule) GetTargetsOk() (*Targets, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *ActivityRule) SetTargets(v Targets)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


