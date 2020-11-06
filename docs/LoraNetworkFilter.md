# LoraNetworkFilter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MessageTypes** | Pointer to **[]string** | list of filtered all messageTypes. Null or empty to accept all message types | [optional] 

## Methods

### NewLoraNetworkFilter

`func NewLoraNetworkFilter() *LoraNetworkFilter`

NewLoraNetworkFilter instantiates a new LoraNetworkFilter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraNetworkFilterWithDefaults

`func NewLoraNetworkFilterWithDefaults() *LoraNetworkFilter`

NewLoraNetworkFilterWithDefaults instantiates a new LoraNetworkFilter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMessageTypes

`func (o *LoraNetworkFilter) GetMessageTypes() []string`

GetMessageTypes returns the MessageTypes field if non-nil, zero value otherwise.

### GetMessageTypesOk

`func (o *LoraNetworkFilter) GetMessageTypesOk() (*[]string, bool)`

GetMessageTypesOk returns a tuple with the MessageTypes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTypes

`func (o *LoraNetworkFilter) SetMessageTypes(v []string)`

SetMessageTypes sets MessageTypes field to given value.

### HasMessageTypes

`func (o *LoraNetworkFilter) HasMessageTypes() bool`

HasMessageTypes returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


