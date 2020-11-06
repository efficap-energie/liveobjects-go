# FifoPublishAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FifoName** | Pointer to **string** | Target fifo to publish into. Refer to Topics API for managing these fifos | [optional] 

## Methods

### NewFifoPublishAction

`func NewFifoPublishAction() *FifoPublishAction`

NewFifoPublishAction instantiates a new FifoPublishAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFifoPublishActionWithDefaults

`func NewFifoPublishActionWithDefaults() *FifoPublishAction`

NewFifoPublishActionWithDefaults instantiates a new FifoPublishAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFifoName

`func (o *FifoPublishAction) GetFifoName() string`

GetFifoName returns the FifoName field if non-nil, zero value otherwise.

### GetFifoNameOk

`func (o *FifoPublishAction) GetFifoNameOk() (*string, bool)`

GetFifoNameOk returns a tuple with the FifoName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifoName

`func (o *FifoPublishAction) SetFifoName(v string)`

SetFifoName sets FifoName field to given value.

### HasFifoName

`func (o *FifoPublishAction) HasFifoName() bool`

HasFifoName returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


