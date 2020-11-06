# DataMessageTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**DataMessageFilter**](DataMessageFilter.md) |  | [optional] 
**Version** | **int32** | requested object version | 

## Methods

### NewDataMessageTrigger

`func NewDataMessageTrigger(version int32, ) *DataMessageTrigger`

NewDataMessageTrigger instantiates a new DataMessageTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataMessageTriggerWithDefaults

`func NewDataMessageTriggerWithDefaults() *DataMessageTrigger`

NewDataMessageTriggerWithDefaults instantiates a new DataMessageTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *DataMessageTrigger) GetFilter() DataMessageFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *DataMessageTrigger) GetFilterOk() (*DataMessageFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *DataMessageTrigger) SetFilter(v DataMessageFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *DataMessageTrigger) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetVersion

`func (o *DataMessageTrigger) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *DataMessageTrigger) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *DataMessageTrigger) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


