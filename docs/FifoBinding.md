# FifoBinding

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DestFifoName** | **string** | Destination Fifo Name | 
**RoutingKeyFilter** | **string** | routing Key filter | 

## Methods

### NewFifoBinding

`func NewFifoBinding(destFifoName string, routingKeyFilter string, ) *FifoBinding`

NewFifoBinding instantiates a new FifoBinding object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFifoBindingWithDefaults

`func NewFifoBindingWithDefaults() *FifoBinding`

NewFifoBindingWithDefaults instantiates a new FifoBinding object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDestFifoName

`func (o *FifoBinding) GetDestFifoName() string`

GetDestFifoName returns the DestFifoName field if non-nil, zero value otherwise.

### GetDestFifoNameOk

`func (o *FifoBinding) GetDestFifoNameOk() (*string, bool)`

GetDestFifoNameOk returns a tuple with the DestFifoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestFifoName

`func (o *FifoBinding) SetDestFifoName(v string)`

SetDestFifoName sets DestFifoName field to given value.


### GetRoutingKeyFilter

`func (o *FifoBinding) GetRoutingKeyFilter() string`

GetRoutingKeyFilter returns the RoutingKeyFilter field if non-nil, zero value otherwise.

### GetRoutingKeyFilterOk

`func (o *FifoBinding) GetRoutingKeyFilterOk() (*string, bool)`

GetRoutingKeyFilterOk returns a tuple with the RoutingKeyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKeyFilter

`func (o *FifoBinding) SetRoutingKeyFilter(v string)`

SetRoutingKeyFilter sets RoutingKeyFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


