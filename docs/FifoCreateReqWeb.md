# FifoCreateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MaxLengthBytes** | **int64** | Maximum number of bytes that can be stored (0 means no storage, else min value is 524288 and max value depends of your account settings) | 
**MessageTtl** | Pointer to **int32** | Time in seconds before messages are dropped if they are not consumed | [optional] 
**Name** | **string** | Topic name of the FIFO. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 

## Methods

### NewFifoCreateReqWeb

`func NewFifoCreateReqWeb(maxLengthBytes int64, name string, ) *FifoCreateReqWeb`

NewFifoCreateReqWeb instantiates a new FifoCreateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFifoCreateReqWebWithDefaults

`func NewFifoCreateReqWebWithDefaults() *FifoCreateReqWeb`

NewFifoCreateReqWebWithDefaults instantiates a new FifoCreateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMaxLengthBytes

`func (o *FifoCreateReqWeb) GetMaxLengthBytes() int64`

GetMaxLengthBytes returns the MaxLengthBytes field if non-nil, zero value otherwise.

### GetMaxLengthBytesOk

`func (o *FifoCreateReqWeb) GetMaxLengthBytesOk() (*int64, bool)`

GetMaxLengthBytesOk returns a tuple with the MaxLengthBytes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxLengthBytes

`func (o *FifoCreateReqWeb) SetMaxLengthBytes(v int64)`

SetMaxLengthBytes sets MaxLengthBytes field to given value.


### GetMessageTtl

`func (o *FifoCreateReqWeb) GetMessageTtl() int32`

GetMessageTtl returns the MessageTtl field if non-nil, zero value otherwise.

### GetMessageTtlOk

`func (o *FifoCreateReqWeb) GetMessageTtlOk() (*int32, bool)`

GetMessageTtlOk returns a tuple with the MessageTtl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessageTtl

`func (o *FifoCreateReqWeb) SetMessageTtl(v int32)`

SetMessageTtl sets MessageTtl field to given value.

### HasMessageTtl

`func (o *FifoCreateReqWeb) HasMessageTtl() bool`

HasMessageTtl returns a boolean if a field has been set.

### GetName

`func (o *FifoCreateReqWeb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FifoCreateReqWeb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FifoCreateReqWeb) SetName(v string)`

SetName sets Name field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


