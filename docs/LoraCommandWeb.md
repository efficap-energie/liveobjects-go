# LoraCommandWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Confirmed** | **bool** | Network ack confirmation | 
**Data** | **string** | Hexadecimal raw data of the command | 
**Port** | **int32** | Port of the device on which the command was sent (cf. LoRaWan) | 

## Methods

### NewLoraCommandWeb

`func NewLoraCommandWeb(confirmed bool, data string, port int32, ) *LoraCommandWeb`

NewLoraCommandWeb instantiates a new LoraCommandWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraCommandWebWithDefaults

`func NewLoraCommandWebWithDefaults() *LoraCommandWeb`

NewLoraCommandWebWithDefaults instantiates a new LoraCommandWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfirmed

`func (o *LoraCommandWeb) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *LoraCommandWeb) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *LoraCommandWeb) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.


### GetData

`func (o *LoraCommandWeb) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoraCommandWeb) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoraCommandWeb) SetData(v string)`

SetData sets Data field to given value.


### GetPort

`func (o *LoraCommandWeb) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoraCommandWeb) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoraCommandWeb) SetPort(v int32)`

SetPort sets Port field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


