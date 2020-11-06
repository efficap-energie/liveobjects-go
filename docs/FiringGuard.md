# FiringGuard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **string** | creation of this FiringGuard | 
**ExpireAt** | Pointer to **string** | expiration of this FiringGuard | [optional] 
**FiringRuleId** | **string** | id of the FiringRule linked to this FiringGuard. | 
**GuardCriteria** | [**[]SelectionCriteria**](SelectionCriteria.md) | pairs of aggregationKey/value that are hold by this FiringGuard | 
**Id** | **string** | id of the FiringGuard. | 
**MatchingContext** | [**MatchingContext**](MatchingContext.md) |  | 
**TenantId** | **string** | identifier of tenant account this FiringRule belongs to . | 

## Methods

### NewFiringGuard

`func NewFiringGuard(created string, firingRuleId string, guardCriteria []SelectionCriteria, id string, matchingContext MatchingContext, tenantId string, ) *FiringGuard`

NewFiringGuard instantiates a new FiringGuard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFiringGuardWithDefaults

`func NewFiringGuardWithDefaults() *FiringGuard`

NewFiringGuardWithDefaults instantiates a new FiringGuard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *FiringGuard) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FiringGuard) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FiringGuard) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetExpireAt

`func (o *FiringGuard) GetExpireAt() string`

GetExpireAt returns the ExpireAt field if non-nil, zero value otherwise.

### GetExpireAtOk

`func (o *FiringGuard) GetExpireAtOk() (*string, bool)`

GetExpireAtOk returns a tuple with the ExpireAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpireAt

`func (o *FiringGuard) SetExpireAt(v string)`

SetExpireAt sets ExpireAt field to given value.

### HasExpireAt

`func (o *FiringGuard) HasExpireAt() bool`

HasExpireAt returns a boolean if a field has been set.

### GetFiringRuleId

`func (o *FiringGuard) GetFiringRuleId() string`

GetFiringRuleId returns the FiringRuleId field if non-nil, zero value otherwise.

### GetFiringRuleIdOk

`func (o *FiringGuard) GetFiringRuleIdOk() (*string, bool)`

GetFiringRuleIdOk returns a tuple with the FiringRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiringRuleId

`func (o *FiringGuard) SetFiringRuleId(v string)`

SetFiringRuleId sets FiringRuleId field to given value.


### GetGuardCriteria

`func (o *FiringGuard) GetGuardCriteria() []SelectionCriteria`

GetGuardCriteria returns the GuardCriteria field if non-nil, zero value otherwise.

### GetGuardCriteriaOk

`func (o *FiringGuard) GetGuardCriteriaOk() (*[]SelectionCriteria, bool)`

GetGuardCriteriaOk returns a tuple with the GuardCriteria field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuardCriteria

`func (o *FiringGuard) SetGuardCriteria(v []SelectionCriteria)`

SetGuardCriteria sets GuardCriteria field to given value.


### GetId

`func (o *FiringGuard) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FiringGuard) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FiringGuard) SetId(v string)`

SetId sets Id field to given value.


### GetMatchingContext

`func (o *FiringGuard) GetMatchingContext() MatchingContext`

GetMatchingContext returns the MatchingContext field if non-nil, zero value otherwise.

### GetMatchingContextOk

`func (o *FiringGuard) GetMatchingContextOk() (*MatchingContext, bool)`

GetMatchingContextOk returns a tuple with the MatchingContext field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingContext

`func (o *FiringGuard) SetMatchingContext(v MatchingContext)`

SetMatchingContext sets MatchingContext field to given value.


### GetTenantId

`func (o *FiringGuard) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *FiringGuard) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *FiringGuard) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


