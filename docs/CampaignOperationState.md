# CampaignOperationState

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Type of operation | 
**CurrentRetry** | Pointer to **int32** | Optional current number of retries made for the operation | [optional] 
**Ended** | Pointer to **string** | Date when the operation was finished | [optional] 
**Error** | Pointer to [**CampaignOperationStateError**](CampaignOperationStateError.md) |  | [optional] 
**OperationId** | Pointer to **string** | Optional identifier associated to the current underlying operation (e.g. commandId,...) | [optional] 
**OperationStatus** | Pointer to **string** | Current status returned by the underlying operation | [optional] 
**Started** | Pointer to **string** | Date when the operation was started | [optional] 
**Updated** | Pointer to **string** | Date when the operation status was updated for the last time | [optional] 

## Methods

### NewCampaignOperationState

`func NewCampaignOperationState(action string, ) *CampaignOperationState`

NewCampaignOperationState instantiates a new CampaignOperationState object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignOperationStateWithDefaults

`func NewCampaignOperationStateWithDefaults() *CampaignOperationState`

NewCampaignOperationStateWithDefaults instantiates a new CampaignOperationState object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CampaignOperationState) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CampaignOperationState) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CampaignOperationState) SetAction(v string)`

SetAction sets Action field to given value.


### GetCurrentRetry

`func (o *CampaignOperationState) GetCurrentRetry() int32`

GetCurrentRetry returns the CurrentRetry field if non-nil, zero value otherwise.

### GetCurrentRetryOk

`func (o *CampaignOperationState) GetCurrentRetryOk() (*int32, bool)`

GetCurrentRetryOk returns a tuple with the CurrentRetry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentRetry

`func (o *CampaignOperationState) SetCurrentRetry(v int32)`

SetCurrentRetry sets CurrentRetry field to given value.

### HasCurrentRetry

`func (o *CampaignOperationState) HasCurrentRetry() bool`

HasCurrentRetry returns a boolean if a field has been set.

### GetEnded

`func (o *CampaignOperationState) GetEnded() string`

GetEnded returns the Ended field if non-nil, zero value otherwise.

### GetEndedOk

`func (o *CampaignOperationState) GetEndedOk() (*string, bool)`

GetEndedOk returns a tuple with the Ended field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnded

`func (o *CampaignOperationState) SetEnded(v string)`

SetEnded sets Ended field to given value.

### HasEnded

`func (o *CampaignOperationState) HasEnded() bool`

HasEnded returns a boolean if a field has been set.

### GetError

`func (o *CampaignOperationState) GetError() CampaignOperationStateError`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *CampaignOperationState) GetErrorOk() (*CampaignOperationStateError, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *CampaignOperationState) SetError(v CampaignOperationStateError)`

SetError sets Error field to given value.

### HasError

`func (o *CampaignOperationState) HasError() bool`

HasError returns a boolean if a field has been set.

### GetOperationId

`func (o *CampaignOperationState) GetOperationId() string`

GetOperationId returns the OperationId field if non-nil, zero value otherwise.

### GetOperationIdOk

`func (o *CampaignOperationState) GetOperationIdOk() (*string, bool)`

GetOperationIdOk returns a tuple with the OperationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationId

`func (o *CampaignOperationState) SetOperationId(v string)`

SetOperationId sets OperationId field to given value.

### HasOperationId

`func (o *CampaignOperationState) HasOperationId() bool`

HasOperationId returns a boolean if a field has been set.

### GetOperationStatus

`func (o *CampaignOperationState) GetOperationStatus() string`

GetOperationStatus returns the OperationStatus field if non-nil, zero value otherwise.

### GetOperationStatusOk

`func (o *CampaignOperationState) GetOperationStatusOk() (*string, bool)`

GetOperationStatusOk returns a tuple with the OperationStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOperationStatus

`func (o *CampaignOperationState) SetOperationStatus(v string)`

SetOperationStatus sets OperationStatus field to given value.

### HasOperationStatus

`func (o *CampaignOperationState) HasOperationStatus() bool`

HasOperationStatus returns a boolean if a field has been set.

### GetStarted

`func (o *CampaignOperationState) GetStarted() string`

GetStarted returns the Started field if non-nil, zero value otherwise.

### GetStartedOk

`func (o *CampaignOperationState) GetStartedOk() (*string, bool)`

GetStartedOk returns a tuple with the Started field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStarted

`func (o *CampaignOperationState) SetStarted(v string)`

SetStarted sets Started field to given value.

### HasStarted

`func (o *CampaignOperationState) HasStarted() bool`

HasStarted returns a boolean if a field has been set.

### GetUpdated

`func (o *CampaignOperationState) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CampaignOperationState) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CampaignOperationState) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CampaignOperationState) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


