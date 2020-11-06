# LoraNetworkTrigger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**LoraNetworkFilter**](LoraNetworkFilter.md) |  | [optional] 
**Version** | **int32** | requested object version | 

## Methods

### NewLoraNetworkTrigger

`func NewLoraNetworkTrigger(version int32, ) *LoraNetworkTrigger`

NewLoraNetworkTrigger instantiates a new LoraNetworkTrigger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraNetworkTriggerWithDefaults

`func NewLoraNetworkTriggerWithDefaults() *LoraNetworkTrigger`

NewLoraNetworkTriggerWithDefaults instantiates a new LoraNetworkTrigger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *LoraNetworkTrigger) GetFilter() LoraNetworkFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *LoraNetworkTrigger) GetFilterOk() (*LoraNetworkFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *LoraNetworkTrigger) SetFilter(v LoraNetworkFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *LoraNetworkTrigger) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetVersion

`func (o *LoraNetworkTrigger) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LoraNetworkTrigger) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LoraNetworkTrigger) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


