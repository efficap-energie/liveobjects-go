# InterfaceCapability

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Available** | Pointer to **bool** | Indicates whether the capability is available for the interface | [optional] [readonly] 
**Version** | Pointer to **int32** | Capability version | [optional] [readonly] 

## Methods

### NewInterfaceCapability

`func NewInterfaceCapability() *InterfaceCapability`

NewInterfaceCapability instantiates a new InterfaceCapability object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceCapabilityWithDefaults

`func NewInterfaceCapabilityWithDefaults() *InterfaceCapability`

NewInterfaceCapabilityWithDefaults instantiates a new InterfaceCapability object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAvailable

`func (o *InterfaceCapability) GetAvailable() bool`

GetAvailable returns the Available field if non-nil, zero value otherwise.

### GetAvailableOk

`func (o *InterfaceCapability) GetAvailableOk() (*bool, bool)`

GetAvailableOk returns a tuple with the Available field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAvailable

`func (o *InterfaceCapability) SetAvailable(v bool)`

SetAvailable sets Available field to given value.

### HasAvailable

`func (o *InterfaceCapability) HasAvailable() bool`

HasAvailable returns a boolean if a field has been set.

### GetVersion

`func (o *InterfaceCapability) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *InterfaceCapability) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *InterfaceCapability) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *InterfaceCapability) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


