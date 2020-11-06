# CampaignUpdating

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Campaign description | [optional] 
**Name** | Pointer to **string** | Campaign name | [optional] 
**Operations** | Pointer to [**[]CampaignOperation**](CampaignOperation.md) | List of operations to be executed | [optional] 
**Options** | Pointer to [**CampaignOptions**](CampaignOptions.md) |  | [optional] 
**Planning** | Pointer to [**CampaignPlanning**](CampaignPlanning.md) |  | [optional] 
**Targets** | Pointer to [**DeviceSelector**](DeviceSelector.md) |  | [optional] 

## Methods

### NewCampaignUpdating

`func NewCampaignUpdating() *CampaignUpdating`

NewCampaignUpdating instantiates a new CampaignUpdating object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignUpdatingWithDefaults

`func NewCampaignUpdatingWithDefaults() *CampaignUpdating`

NewCampaignUpdatingWithDefaults instantiates a new CampaignUpdating object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CampaignUpdating) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignUpdating) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignUpdating) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignUpdating) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CampaignUpdating) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignUpdating) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignUpdating) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *CampaignUpdating) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOperations

`func (o *CampaignUpdating) GetOperations() []CampaignOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CampaignUpdating) GetOperationsOk() (*[]CampaignOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CampaignUpdating) SetOperations(v []CampaignOperation)`

SetOperations sets Operations field to given value.

### HasOperations

`func (o *CampaignUpdating) HasOperations() bool`

HasOperations returns a boolean if a field has been set.

### GetOptions

`func (o *CampaignUpdating) GetOptions() CampaignOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CampaignUpdating) GetOptionsOk() (*CampaignOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CampaignUpdating) SetOptions(v CampaignOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CampaignUpdating) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPlanning

`func (o *CampaignUpdating) GetPlanning() CampaignPlanning`

GetPlanning returns the Planning field if non-nil, zero value otherwise.

### GetPlanningOk

`func (o *CampaignUpdating) GetPlanningOk() (*CampaignPlanning, bool)`

GetPlanningOk returns a tuple with the Planning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanning

`func (o *CampaignUpdating) SetPlanning(v CampaignPlanning)`

SetPlanning sets Planning field to given value.

### HasPlanning

`func (o *CampaignUpdating) HasPlanning() bool`

HasPlanning returns a boolean if a field has been set.

### GetTargets

`func (o *CampaignUpdating) GetTargets() DeviceSelector`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CampaignUpdating) GetTargetsOk() (*DeviceSelector, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CampaignUpdating) SetTargets(v DeviceSelector)`

SetTargets sets Targets field to given value.

### HasTargets

`func (o *CampaignUpdating) HasTargets() bool`

HasTargets returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


