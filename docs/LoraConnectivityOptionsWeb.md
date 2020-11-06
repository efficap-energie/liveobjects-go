# LoraConnectivityOptionsWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AckUl** | Pointer to **bool** | Options for Ack uplink | [optional] 
**Tdoa** | Pointer to **bool** | Options for TDOA geolocation | [optional] 

## Methods

### NewLoraConnectivityOptionsWeb

`func NewLoraConnectivityOptionsWeb() *LoraConnectivityOptionsWeb`

NewLoraConnectivityOptionsWeb instantiates a new LoraConnectivityOptionsWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraConnectivityOptionsWebWithDefaults

`func NewLoraConnectivityOptionsWebWithDefaults() *LoraConnectivityOptionsWeb`

NewLoraConnectivityOptionsWebWithDefaults instantiates a new LoraConnectivityOptionsWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAckUl

`func (o *LoraConnectivityOptionsWeb) GetAckUl() bool`

GetAckUl returns the AckUl field if non-nil, zero value otherwise.

### GetAckUlOk

`func (o *LoraConnectivityOptionsWeb) GetAckUlOk() (*bool, bool)`

GetAckUlOk returns a tuple with the AckUl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckUl

`func (o *LoraConnectivityOptionsWeb) SetAckUl(v bool)`

SetAckUl sets AckUl field to given value.

### HasAckUl

`func (o *LoraConnectivityOptionsWeb) HasAckUl() bool`

HasAckUl returns a boolean if a field has been set.

### GetTdoa

`func (o *LoraConnectivityOptionsWeb) GetTdoa() bool`

GetTdoa returns the Tdoa field if non-nil, zero value otherwise.

### GetTdoaOk

`func (o *LoraConnectivityOptionsWeb) GetTdoaOk() (*bool, bool)`

GetTdoaOk returns a tuple with the Tdoa field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTdoa

`func (o *LoraConnectivityOptionsWeb) SetTdoa(v bool)`

SetTdoa sets Tdoa field to given value.

### HasTdoa

`func (o *LoraConnectivityOptionsWeb) HasTdoa() bool`

HasTdoa returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


