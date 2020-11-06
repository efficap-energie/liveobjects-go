# Thresholds

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxMsisdn** | **int32** | max of msisdn allow to register | 
**MaxSmsInPerDay** | **int32** | max of sms received allow | 
**MaxSmsOutPerDay** | **int32** | max of sms send allow | 
**MaxTags** | **int32** | max of tag allow to register | 

## Methods

### NewThresholds

`func NewThresholds(maxMsisdn int32, maxSmsInPerDay int32, maxSmsOutPerDay int32, maxTags int32, ) *Thresholds`

NewThresholds instantiates a new Thresholds object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewThresholdsWithDefaults

`func NewThresholdsWithDefaults() *Thresholds`

NewThresholdsWithDefaults instantiates a new Thresholds object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxMsisdn

`func (o *Thresholds) GetMaxMsisdn() int32`

GetMaxMsisdn returns the MaxMsisdn field if non-nil, zero value otherwise.

### GetMaxMsisdnOk

`func (o *Thresholds) GetMaxMsisdnOk() (*int32, bool)`

GetMaxMsisdnOk returns a tuple with the MaxMsisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxMsisdn

`func (o *Thresholds) SetMaxMsisdn(v int32)`

SetMaxMsisdn sets MaxMsisdn field to given value.


### GetMaxSmsInPerDay

`func (o *Thresholds) GetMaxSmsInPerDay() int32`

GetMaxSmsInPerDay returns the MaxSmsInPerDay field if non-nil, zero value otherwise.

### GetMaxSmsInPerDayOk

`func (o *Thresholds) GetMaxSmsInPerDayOk() (*int32, bool)`

GetMaxSmsInPerDayOk returns a tuple with the MaxSmsInPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSmsInPerDay

`func (o *Thresholds) SetMaxSmsInPerDay(v int32)`

SetMaxSmsInPerDay sets MaxSmsInPerDay field to given value.


### GetMaxSmsOutPerDay

`func (o *Thresholds) GetMaxSmsOutPerDay() int32`

GetMaxSmsOutPerDay returns the MaxSmsOutPerDay field if non-nil, zero value otherwise.

### GetMaxSmsOutPerDayOk

`func (o *Thresholds) GetMaxSmsOutPerDayOk() (*int32, bool)`

GetMaxSmsOutPerDayOk returns a tuple with the MaxSmsOutPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxSmsOutPerDay

`func (o *Thresholds) SetMaxSmsOutPerDay(v int32)`

SetMaxSmsOutPerDay sets MaxSmsOutPerDay field to given value.


### GetMaxTags

`func (o *Thresholds) GetMaxTags() int32`

GetMaxTags returns the MaxTags field if non-nil, zero value otherwise.

### GetMaxTagsOk

`func (o *Thresholds) GetMaxTagsOk() (*int32, bool)`

GetMaxTagsOk returns a tuple with the MaxTags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxTags

`func (o *Thresholds) SetMaxTags(v int32)`

SetMaxTags sets MaxTags field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


