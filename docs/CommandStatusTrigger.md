# CommandStatusTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**CommandStatusFilter**](CommandStatusFilter.md) |  | [optional] 
**Version** | **int32** | requested object version | 

## Methods

### NewCommandStatusTrigger

`func NewCommandStatusTrigger(version int32, ) *CommandStatusTrigger`

NewCommandStatusTrigger instantiates a new CommandStatusTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandStatusTriggerWithDefaults

`func NewCommandStatusTriggerWithDefaults() *CommandStatusTrigger`

NewCommandStatusTriggerWithDefaults instantiates a new CommandStatusTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *CommandStatusTrigger) GetFilter() CommandStatusFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *CommandStatusTrigger) GetFilterOk() (*CommandStatusFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *CommandStatusTrigger) SetFilter(v CommandStatusFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *CommandStatusTrigger) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetVersion

`func (o *CommandStatusTrigger) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CommandStatusTrigger) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CommandStatusTrigger) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


