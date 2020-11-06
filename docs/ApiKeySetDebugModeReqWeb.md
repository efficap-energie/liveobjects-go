# ApiKeySetDebugModeReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Activated** | **bool** | The state of the debug mode | 
**DurationSeconds** | Pointer to **int64** | The duration during which the debug mode will be activated (default &#x3D; 15 min) | [optional] 

## Methods

### NewApiKeySetDebugModeReqWeb

`func NewApiKeySetDebugModeReqWeb(activated bool, ) *ApiKeySetDebugModeReqWeb`

NewApiKeySetDebugModeReqWeb instantiates a new ApiKeySetDebugModeReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApiKeySetDebugModeReqWebWithDefaults

`func NewApiKeySetDebugModeReqWebWithDefaults() *ApiKeySetDebugModeReqWeb`

NewApiKeySetDebugModeReqWebWithDefaults instantiates a new ApiKeySetDebugModeReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivated

`func (o *ApiKeySetDebugModeReqWeb) GetActivated() bool`

GetActivated returns the Activated field if non-nil, zero value otherwise.

### GetActivatedOk

`func (o *ApiKeySetDebugModeReqWeb) GetActivatedOk() (*bool, bool)`

GetActivatedOk returns a tuple with the Activated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivated

`func (o *ApiKeySetDebugModeReqWeb) SetActivated(v bool)`

SetActivated sets Activated field to given value.


### GetDurationSeconds

`func (o *ApiKeySetDebugModeReqWeb) GetDurationSeconds() int64`

GetDurationSeconds returns the DurationSeconds field if non-nil, zero value otherwise.

### GetDurationSecondsOk

`func (o *ApiKeySetDebugModeReqWeb) GetDurationSecondsOk() (*int64, bool)`

GetDurationSecondsOk returns a tuple with the DurationSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDurationSeconds

`func (o *ApiKeySetDebugModeReqWeb) SetDurationSeconds(v int64)`

SetDurationSeconds sets DurationSeconds field to given value.

### HasDurationSeconds

`func (o *ApiKeySetDebugModeReqWeb) HasDurationSeconds() bool`

HasDurationSeconds returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


