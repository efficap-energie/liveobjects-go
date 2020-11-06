# CommandStatusFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **[]string** | list of command status devices. Null or empty to accept all status | [optional] 

## Methods

### NewCommandStatusFilter

`func NewCommandStatusFilter() *CommandStatusFilter`

NewCommandStatusFilter instantiates a new CommandStatusFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandStatusFilterWithDefaults

`func NewCommandStatusFilterWithDefaults() *CommandStatusFilter`

NewCommandStatusFilterWithDefaults instantiates a new CommandStatusFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *CommandStatusFilter) GetStatus() []string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommandStatusFilter) GetStatusOk() (*[]string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommandStatusFilter) SetStatus(v []string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommandStatusFilter) HasStatus() bool`

HasStatus returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


