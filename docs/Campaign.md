# Campaign

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CampaignStatus** | Pointer to **string** |  | [optional] [readonly] 
**Created** | Pointer to **string** |  | [optional] [readonly] 
**Description** | Pointer to **string** | Campaign description | [optional] 
**Id** | **string** | campaign id | 
**Name** | **string** | campaign name | 
**NumberOfTargets** | Pointer to **int32** | number of targets after deduplication | [optional] [readonly] 
**Operations** | [**[]CampaignOperation**](CampaignOperation.md) | requested operations | 
**Options** | Pointer to [**CampaignOptions**](CampaignOptions.md) |  | [optional] 
**Planning** | [**CampaignPlanning**](CampaignPlanning.md) |  | 
**Target** | [**DeviceSelector**](DeviceSelector.md) |  | 
**TotalTargetsPerStatus** | Pointer to [**CampaignStatsPerStatus**](CampaignStatsPerStatus.md) |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] [readonly] 

## Methods

### NewCampaign

`func NewCampaign(id string, name string, operations []CampaignOperation, planning CampaignPlanning, target DeviceSelector, ) *Campaign`

NewCampaign instantiates a new Campaign object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignWithDefaults

`func NewCampaignWithDefaults() *Campaign`

NewCampaignWithDefaults instantiates a new Campaign object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCampaignStatus

`func (o *Campaign) GetCampaignStatus() string`

GetCampaignStatus returns the CampaignStatus field if non-nil, zero value otherwise.

### GetCampaignStatusOk

`func (o *Campaign) GetCampaignStatusOk() (*string, bool)`

GetCampaignStatusOk returns a tuple with the CampaignStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCampaignStatus

`func (o *Campaign) SetCampaignStatus(v string)`

SetCampaignStatus sets CampaignStatus field to given value.

### HasCampaignStatus

`func (o *Campaign) HasCampaignStatus() bool`

HasCampaignStatus returns a boolean if a field has been set.

### GetCreated

`func (o *Campaign) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Campaign) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Campaign) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Campaign) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Campaign) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Campaign) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Campaign) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Campaign) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *Campaign) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Campaign) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Campaign) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Campaign) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Campaign) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Campaign) SetName(v string)`

SetName sets Name field to given value.


### GetNumberOfTargets

`func (o *Campaign) GetNumberOfTargets() int32`

GetNumberOfTargets returns the NumberOfTargets field if non-nil, zero value otherwise.

### GetNumberOfTargetsOk

`func (o *Campaign) GetNumberOfTargetsOk() (*int32, bool)`

GetNumberOfTargetsOk returns a tuple with the NumberOfTargets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNumberOfTargets

`func (o *Campaign) SetNumberOfTargets(v int32)`

SetNumberOfTargets sets NumberOfTargets field to given value.

### HasNumberOfTargets

`func (o *Campaign) HasNumberOfTargets() bool`

HasNumberOfTargets returns a boolean if a field has been set.

### GetOperations

`func (o *Campaign) GetOperations() []CampaignOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *Campaign) GetOperationsOk() (*[]CampaignOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *Campaign) SetOperations(v []CampaignOperation)`

SetOperations sets Operations field to given value.


### GetOptions

`func (o *Campaign) GetOptions() CampaignOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *Campaign) GetOptionsOk() (*CampaignOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *Campaign) SetOptions(v CampaignOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *Campaign) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPlanning

`func (o *Campaign) GetPlanning() CampaignPlanning`

GetPlanning returns the Planning field if non-nil, zero value otherwise.

### GetPlanningOk

`func (o *Campaign) GetPlanningOk() (*CampaignPlanning, bool)`

GetPlanningOk returns a tuple with the Planning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanning

`func (o *Campaign) SetPlanning(v CampaignPlanning)`

SetPlanning sets Planning field to given value.


### GetTarget

`func (o *Campaign) GetTarget() DeviceSelector`

GetTarget returns the Target field if non-nil, zero value otherwise.

### GetTargetOk

`func (o *Campaign) GetTargetOk() (*DeviceSelector, bool)`

GetTargetOk returns a tuple with the Target field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTarget

`func (o *Campaign) SetTarget(v DeviceSelector)`

SetTarget sets Target field to given value.


### GetTotalTargetsPerStatus

`func (o *Campaign) GetTotalTargetsPerStatus() CampaignStatsPerStatus`

GetTotalTargetsPerStatus returns the TotalTargetsPerStatus field if non-nil, zero value otherwise.

### GetTotalTargetsPerStatusOk

`func (o *Campaign) GetTotalTargetsPerStatusOk() (*CampaignStatsPerStatus, bool)`

GetTotalTargetsPerStatusOk returns a tuple with the TotalTargetsPerStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalTargetsPerStatus

`func (o *Campaign) SetTotalTargetsPerStatus(v CampaignStatsPerStatus)`

SetTotalTargetsPerStatus sets TotalTargetsPerStatus field to given value.

### HasTotalTargetsPerStatus

`func (o *Campaign) HasTotalTargetsPerStatus() bool`

HasTotalTargetsPerStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *Campaign) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Campaign) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Campaign) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Campaign) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


