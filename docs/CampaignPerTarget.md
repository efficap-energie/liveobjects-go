# CampaignPerTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **string** |  | 
**Device** | **string** | Device identifier | 
**Operations** | [**[]CampaignOperationState**](CampaignOperationState.md) |  | 
**Status** | **string** | Global status of the sequence of operations | 
**Updated** | **string** |  | 

## Methods

### NewCampaignPerTarget

`func NewCampaignPerTarget(created string, device string, operations []CampaignOperationState, status string, updated string, ) *CampaignPerTarget`

NewCampaignPerTarget instantiates a new CampaignPerTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignPerTargetWithDefaults

`func NewCampaignPerTargetWithDefaults() *CampaignPerTarget`

NewCampaignPerTargetWithDefaults instantiates a new CampaignPerTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *CampaignPerTarget) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *CampaignPerTarget) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *CampaignPerTarget) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetDevice

`func (o *CampaignPerTarget) GetDevice() string`

GetDevice returns the Device field if non-nil, zero value otherwise.

### GetDeviceOk

`func (o *CampaignPerTarget) GetDeviceOk() (*string, bool)`

GetDeviceOk returns a tuple with the Device field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevice

`func (o *CampaignPerTarget) SetDevice(v string)`

SetDevice sets Device field to given value.


### GetOperations

`func (o *CampaignPerTarget) GetOperations() []CampaignOperationState`

GetOperations returns the Operations field if non-nil, zero value otherwise.

### GetOperationsOk

`func (o *CampaignPerTarget) GetOperationsOk() (*[]CampaignOperationState, bool)`

GetOperationsOk returns a tuple with the Operations field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperations

`func (o *CampaignPerTarget) SetOperations(v []CampaignOperationState)`

SetOperations sets Operations field to given value.


### GetStatus

`func (o *CampaignPerTarget) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CampaignPerTarget) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CampaignPerTarget) SetStatus(v string)`

SetStatus sets Status field to given value.


### GetUpdated

`func (o *CampaignPerTarget) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CampaignPerTarget) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CampaignPerTarget) SetUpdated(v string)`

SetUpdated sets Updated field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


