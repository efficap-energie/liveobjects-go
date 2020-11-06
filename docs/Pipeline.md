# Pipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** | pipeline&#39;s creation date (ISO-8601). Should be null for post and update action. | [optional] 
**Description** | Pointer to **string** | pipeline description (length limit : 2000). | [optional] 
**Enabled** | **bool** | pipeline activation. Default is false. | 
**Filter** | Pointer to [**PipelineFilter**](PipelineFilter.md) |  | [optional] 
**Id** | Pointer to **string** | id of the pipeline. Should be null when used for POST. Required for update. | [optional] 
**Name** | **string** | pipeline name (length limit : 1000). | 
**PriorityLevel** | **int64** | pipeline&#39;s priority level. When several pipelines match the input filter, then the priority is given to the pipeline with lowest priorityLevel. If several pipelines match filter and the share the same priorityLevel, then the oldest pipeline is chosen (based on &#39;created&#39; field). | 
**Steps** | [**[]PipelineStep**](PipelineStep.md) | list of steps of the pipelines. Should contain at least 1 step. Available steps are defined in developer&#39;s guide. | 

## Methods

### NewPipeline

`func NewPipeline(enabled bool, name string, priorityLevel int64, steps []PipelineStep, ) *Pipeline`

NewPipeline instantiates a new Pipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPipelineWithDefaults

`func NewPipelineWithDefaults() *Pipeline`

NewPipelineWithDefaults instantiates a new Pipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *Pipeline) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Pipeline) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Pipeline) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Pipeline) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDescription

`func (o *Pipeline) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Pipeline) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Pipeline) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Pipeline) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Pipeline) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Pipeline) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Pipeline) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetFilter

`func (o *Pipeline) GetFilter() PipelineFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *Pipeline) GetFilterOk() (*PipelineFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *Pipeline) SetFilter(v PipelineFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *Pipeline) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetId

`func (o *Pipeline) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Pipeline) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Pipeline) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Pipeline) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Pipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Pipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Pipeline) SetName(v string)`

SetName sets Name field to given value.


### GetPriorityLevel

`func (o *Pipeline) GetPriorityLevel() int64`

GetPriorityLevel returns the PriorityLevel field if non-nil, zero value otherwise.

### GetPriorityLevelOk

`func (o *Pipeline) GetPriorityLevelOk() (*int64, bool)`

GetPriorityLevelOk returns a tuple with the PriorityLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriorityLevel

`func (o *Pipeline) SetPriorityLevel(v int64)`

SetPriorityLevel sets PriorityLevel field to given value.


### GetSteps

`func (o *Pipeline) GetSteps() []PipelineStep`

GetSteps returns the Steps field if non-nil, zero value otherwise.

### GetStepsOk

`func (o *Pipeline) GetStepsOk() (*[]PipelineStep, bool)`

GetStepsOk returns a tuple with the Steps field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSteps

`func (o *Pipeline) SetSteps(v []PipelineStep)`

SetSteps sets Steps field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


