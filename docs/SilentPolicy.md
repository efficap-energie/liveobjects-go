# SilentPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Duration** | **string** | the duration (ISO 8601) after which the device will be considered as silent if it had no contact with the platform. If less than 10 minutes, this duration will be handled as a 10 minutes delay. A silent event/alarm will be sent after expiration. | 
**RepeatInterval** | Pointer to **string** | if the device previously triggered a silent alarm and still had no contact with the platform; duration (ISO 8601) after which a new silent event/alarm will be sent. If less than 10 minutes, this duration will be handled as a 10 minutes delay. If null, no further alarm will be sent until the new active/silent cycle. | [optional] 

## Methods

### NewSilentPolicy

`func NewSilentPolicy(duration string, ) *SilentPolicy`

NewSilentPolicy instantiates a new SilentPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSilentPolicyWithDefaults

`func NewSilentPolicyWithDefaults() *SilentPolicy`

NewSilentPolicyWithDefaults instantiates a new SilentPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDuration

`func (o *SilentPolicy) GetDuration() string`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *SilentPolicy) GetDurationOk() (*string, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *SilentPolicy) SetDuration(v string)`

SetDuration sets Duration field to given value.


### GetRepeatInterval

`func (o *SilentPolicy) GetRepeatInterval() string`

GetRepeatInterval returns the RepeatInterval field if non-nil, zero value otherwise.

### GetRepeatIntervalOk

`func (o *SilentPolicy) GetRepeatIntervalOk() (*string, bool)`

GetRepeatIntervalOk returns a tuple with the RepeatInterval field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRepeatInterval

`func (o *SilentPolicy) SetRepeatInterval(v string)`

SetRepeatInterval sets RepeatInterval field to given value.

### HasRepeatInterval

`func (o *SilentPolicy) HasRepeatInterval() bool`

HasRepeatInterval returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


