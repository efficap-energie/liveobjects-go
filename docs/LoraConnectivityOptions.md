# LoraConnectivityOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckUl** | **bool** | Enable or disable ack UL option | 
**Tdoa** | Pointer to **bool** | Enable or disable TDOA option | [optional] 

## Methods

### NewLoraConnectivityOptions

`func NewLoraConnectivityOptions(ackUl bool, ) *LoraConnectivityOptions`

NewLoraConnectivityOptions instantiates a new LoraConnectivityOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraConnectivityOptionsWithDefaults

`func NewLoraConnectivityOptionsWithDefaults() *LoraConnectivityOptions`

NewLoraConnectivityOptionsWithDefaults instantiates a new LoraConnectivityOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckUl

`func (o *LoraConnectivityOptions) GetAckUl() bool`

GetAckUl returns the AckUl field if non-nil, zero value otherwise.

### GetAckUlOk

`func (o *LoraConnectivityOptions) GetAckUlOk() (*bool, bool)`

GetAckUlOk returns a tuple with the AckUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUl

`func (o *LoraConnectivityOptions) SetAckUl(v bool)`

SetAckUl sets AckUl field to given value.


### GetTdoa

`func (o *LoraConnectivityOptions) GetTdoa() bool`

GetTdoa returns the Tdoa field if non-nil, zero value otherwise.

### GetTdoaOk

`func (o *LoraConnectivityOptions) GetTdoaOk() (*bool, bool)`

GetTdoaOk returns a tuple with the Tdoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdoa

`func (o *LoraConnectivityOptions) SetTdoa(v bool)`

SetTdoa sets Tdoa field to given value.

### HasTdoa

`func (o *LoraConnectivityOptions) HasTdoa() bool`

HasTdoa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


