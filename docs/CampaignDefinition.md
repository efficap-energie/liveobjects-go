# CampaignDefinition

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | Campaign description | [optional] 
**Name** | **string** | Campaign name | 
**Operations** | [**[]CampaignOperation**](CampaignOperation.md) | List of operations to be executed | 
**Options** | Pointer to [**CampaignOptions**](CampaignOptions.md) |  | [optional] 
**Planning** | [**CampaignPlanning**](CampaignPlanning.md) |  | 
**Targets** | [**DeviceSelector**](DeviceSelector.md) |  | 

## Methods

### NewCampaignDefinition

`func NewCampaignDefinition(name string, operations []CampaignOperation, planning CampaignPlanning, targets DeviceSelector, ) *CampaignDefinition`

NewCampaignDefinition instantiates a new CampaignDefinition object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignDefinitionWithDefaults

`func NewCampaignDefinitionWithDefaults() *CampaignDefinition`

NewCampaignDefinitionWithDefaults instantiates a new CampaignDefinition object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *CampaignDefinition) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CampaignDefinition) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CampaignDefinition) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CampaignDefinition) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetName

`func (o *CampaignDefinition) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CampaignDefinition) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CampaignDefinition) SetName(v string)`

SetName sets Name field to given value.


### GetOperations

`func (o *CampaignDefinition) GetOperations() []CampaignOperation`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CampaignDefinition) GetOperationsOk() (*[]CampaignOperation, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CampaignDefinition) SetOperations(v []CampaignOperation)`

SetOperations sets Operations field to given value.


### GetOptions

`func (o *CampaignDefinition) GetOptions() CampaignOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CampaignDefinition) GetOptionsOk() (*CampaignOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CampaignDefinition) SetOptions(v CampaignOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CampaignDefinition) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetPlanning

`func (o *CampaignDefinition) GetPlanning() CampaignPlanning`

GetPlanning returns the Planning field if non-nil, zero value otherwise.

### GetPlanningOk

`func (o *CampaignDefinition) GetPlanningOk() (*CampaignPlanning, bool)`

GetPlanningOk returns a tuple with the Planning field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanning

`func (o *CampaignDefinition) SetPlanning(v CampaignPlanning)`

SetPlanning sets Planning field to given value.


### GetTargets

`func (o *CampaignDefinition) GetTargets() DeviceSelector`

GetTargets returns the Targets field if non-nil, zero value otherwise.

### GetTargetsOk

`func (o *CampaignDefinition) GetTargetsOk() (*DeviceSelector, bool)`

GetTargetsOk returns a tuple with the Targets field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargets

`func (o *CampaignDefinition) SetTargets(v DeviceSelector)`

SetTargets sets Targets field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


