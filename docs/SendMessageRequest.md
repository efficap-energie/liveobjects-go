# SendMessageRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Contact** | [**ContactMessage**](ContactMessage.md) |  | 
**Message** | **map[string]string** | message | 

## Methods

### NewSendMessageRequest

`func NewSendMessageRequest(contact ContactMessage, message map[string]string, ) *SendMessageRequest`

NewSendMessageRequest instantiates a new SendMessageRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSendMessageRequestWithDefaults

`func NewSendMessageRequestWithDefaults() *SendMessageRequest`

NewSendMessageRequestWithDefaults instantiates a new SendMessageRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContact

`func (o *SendMessageRequest) GetContact() ContactMessage`

GetContact returns the Contact field if non-nil, zero value otherwise.

### GetContactOk

`func (o *SendMessageRequest) GetContactOk() (*ContactMessage, bool)`

GetContactOk returns a tuple with the Contact field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContact

`func (o *SendMessageRequest) SetContact(v ContactMessage)`

SetContact sets Contact field to given value.


### GetMessage

`func (o *SendMessageRequest) GetMessage() map[string]string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *SendMessageRequest) GetMessageOk() (*map[string]string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *SendMessageRequest) SetMessage(v map[string]string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


