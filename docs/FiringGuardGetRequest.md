# FiringGuardGetRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FiringRuleId** | **string** | firingRuleId linked to the FiringGuard to reset | 
**SelectionCriteria** | Pointer to [**[]SelectionCriteria**](SelectionCriteria.md) | criteria that should match the FiringGuard to reset | [optional] 

## Methods

### NewFiringGuardGetRequest

`func NewFiringGuardGetRequest(firingRuleId string, ) *FiringGuardGetRequest`

NewFiringGuardGetRequest instantiates a new FiringGuardGetRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiringGuardGetRequestWithDefaults

`func NewFiringGuardGetRequestWithDefaults() *FiringGuardGetRequest`

NewFiringGuardGetRequestWithDefaults instantiates a new FiringGuardGetRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiringRuleId

`func (o *FiringGuardGetRequest) GetFiringRuleId() string`

GetFiringRuleId returns the FiringRuleId field if non-nil, zero value otherwise.

### GetFiringRuleIdOk

`func (o *FiringGuardGetRequest) GetFiringRuleIdOk() (*string, bool)`

GetFiringRuleIdOk returns a tuple with the FiringRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiringRuleId

`func (o *FiringGuardGetRequest) SetFiringRuleId(v string)`

SetFiringRuleId sets FiringRuleId field to given value.


### GetSelectionCriteria

`func (o *FiringGuardGetRequest) GetSelectionCriteria() []SelectionCriteria`

GetSelectionCriteria returns the SelectionCriteria field if non-nil, zero value otherwise.

### GetSelectionCriteriaOk

`func (o *FiringGuardGetRequest) GetSelectionCriteriaOk() (*[]SelectionCriteria, bool)`

GetSelectionCriteriaOk returns a tuple with the SelectionCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSelectionCriteria

`func (o *FiringGuardGetRequest) SetSelectionCriteria(v []SelectionCriteria)`

SetSelectionCriteria sets SelectionCriteria field to given value.

### HasSelectionCriteria

`func (o *FiringGuardGetRequest) HasSelectionCriteria() bool`

HasSelectionCriteria returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


