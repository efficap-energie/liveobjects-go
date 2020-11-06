# FifoTopic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Consumers** | Pointer to **int64** | Number of active consumers | [optional] 
**MaxLengthBytes** | Pointer to **int64** | Maximum number of bytes that can be stored | [optional] 
**MessageBytes** | Pointer to **int64** | Total number of bytes stored | [optional] 
**MessageTtl** | Pointer to **int32** | Time in seconds before a message is dropped if not consumed | [optional] 
**MessageUnacknowledged** | Pointer to **int64** | Total number of bytes of unacknowledged messages | [optional] 
**MessageUnacknowledgedBytes** | Pointer to **int64** | Total number of bytes of unacknowledged messages | [optional] 
**Messages** | Pointer to **int64** | Number of message stored | [optional] 
**Name** | **string** | Topic name of the FIFO | 

## Methods

### NewFifoTopic

`func NewFifoTopic(name string, ) *FifoTopic`

NewFifoTopic instantiates a new FifoTopic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFifoTopicWithDefaults

`func NewFifoTopicWithDefaults() *FifoTopic`

NewFifoTopicWithDefaults instantiates a new FifoTopic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConsumers

`func (o *FifoTopic) GetConsumers() int64`

GetConsumers returns the Consumers field if non-nil, zero value otherwise.

### GetConsumersOk

`func (o *FifoTopic) GetConsumersOk() (*int64, bool)`

GetConsumersOk returns a tuple with the Consumers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConsumers

`func (o *FifoTopic) SetConsumers(v int64)`

SetConsumers sets Consumers field to given value.

### HasConsumers

`func (o *FifoTopic) HasConsumers() bool`

HasConsumers returns a boolean if a field has been set.

### GetMaxLengthBytes

`func (o *FifoTopic) GetMaxLengthBytes() int64`

GetMaxLengthBytes returns the MaxLengthBytes field if non-nil, zero value otherwise.

### GetMaxLengthBytesOk

`func (o *FifoTopic) GetMaxLengthBytesOk() (*int64, bool)`

GetMaxLengthBytesOk returns a tuple with the MaxLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLengthBytes

`func (o *FifoTopic) SetMaxLengthBytes(v int64)`

SetMaxLengthBytes sets MaxLengthBytes field to given value.

### HasMaxLengthBytes

`func (o *FifoTopic) HasMaxLengthBytes() bool`

HasMaxLengthBytes returns a boolean if a field has been set.

### GetMessageBytes

`func (o *FifoTopic) GetMessageBytes() int64`

GetMessageBytes returns the MessageBytes field if non-nil, zero value otherwise.

### GetMessageBytesOk

`func (o *FifoTopic) GetMessageBytesOk() (*int64, bool)`

GetMessageBytesOk returns a tuple with the MessageBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageBytes

`func (o *FifoTopic) SetMessageBytes(v int64)`

SetMessageBytes sets MessageBytes field to given value.

### HasMessageBytes

`func (o *FifoTopic) HasMessageBytes() bool`

HasMessageBytes returns a boolean if a field has been set.

### GetMessageTtl

`func (o *FifoTopic) GetMessageTtl() int32`

GetMessageTtl returns the MessageTtl field if non-nil, zero value otherwise.

### GetMessageTtlOk

`func (o *FifoTopic) GetMessageTtlOk() (*int32, bool)`

GetMessageTtlOk returns a tuple with the MessageTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtl

`func (o *FifoTopic) SetMessageTtl(v int32)`

SetMessageTtl sets MessageTtl field to given value.

### HasMessageTtl

`func (o *FifoTopic) HasMessageTtl() bool`

HasMessageTtl returns a boolean if a field has been set.

### GetMessageUnacknowledged

`func (o *FifoTopic) GetMessageUnacknowledged() int64`

GetMessageUnacknowledged returns the MessageUnacknowledged field if non-nil, zero value otherwise.

### GetMessageUnacknowledgedOk

`func (o *FifoTopic) GetMessageUnacknowledgedOk() (*int64, bool)`

GetMessageUnacknowledgedOk returns a tuple with the MessageUnacknowledged field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageUnacknowledged

`func (o *FifoTopic) SetMessageUnacknowledged(v int64)`

SetMessageUnacknowledged sets MessageUnacknowledged field to given value.

### HasMessageUnacknowledged

`func (o *FifoTopic) HasMessageUnacknowledged() bool`

HasMessageUnacknowledged returns a boolean if a field has been set.

### GetMessageUnacknowledgedBytes

`func (o *FifoTopic) GetMessageUnacknowledgedBytes() int64`

GetMessageUnacknowledgedBytes returns the MessageUnacknowledgedBytes field if non-nil, zero value otherwise.

### GetMessageUnacknowledgedBytesOk

`func (o *FifoTopic) GetMessageUnacknowledgedBytesOk() (*int64, bool)`

GetMessageUnacknowledgedBytesOk returns a tuple with the MessageUnacknowledgedBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageUnacknowledgedBytes

`func (o *FifoTopic) SetMessageUnacknowledgedBytes(v int64)`

SetMessageUnacknowledgedBytes sets MessageUnacknowledgedBytes field to given value.

### HasMessageUnacknowledgedBytes

`func (o *FifoTopic) HasMessageUnacknowledgedBytes() bool`

HasMessageUnacknowledgedBytes returns a boolean if a field has been set.

### GetMessages

`func (o *FifoTopic) GetMessages() int64`

GetMessages returns the Messages field if non-nil, zero value otherwise.

### GetMessagesOk

`func (o *FifoTopic) GetMessagesOk() (*int64, bool)`

GetMessagesOk returns a tuple with the Messages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessages

`func (o *FifoTopic) SetMessages(v int64)`

SetMessages sets Messages field to given value.

### HasMessages

`func (o *FifoTopic) HasMessages() bool`

HasMessages returns a boolean if a field has been set.

### GetName

`func (o *FifoTopic) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FifoTopic) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FifoTopic) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


