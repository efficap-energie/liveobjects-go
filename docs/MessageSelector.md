# MessageSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filter** | Pointer to [**MessageSelectorFilter**](MessageSelectorFilter.md) |  | [optional] 
**Origin** | **string** | Specifies the source of the message that will trigger an action | 

## Methods

### NewMessageSelector

`func NewMessageSelector(origin string, ) *MessageSelector`

NewMessageSelector instantiates a new MessageSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageSelectorWithDefaults

`func NewMessageSelectorWithDefaults() *MessageSelector`

NewMessageSelectorWithDefaults instantiates a new MessageSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilter

`func (o *MessageSelector) GetFilter() MessageSelectorFilter`

GetFilter returns the Filter field if non-nil, zero value otherwise.

### GetFilterOk

`func (o *MessageSelector) GetFilterOk() (*MessageSelectorFilter, bool)`

GetFilterOk returns a tuple with the Filter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilter

`func (o *MessageSelector) SetFilter(v MessageSelectorFilter)`

SetFilter sets Filter field to given value.

### HasFilter

`func (o *MessageSelector) HasFilter() bool`

HasFilter returns a boolean if a field has been set.

### GetOrigin

`func (o *MessageSelector) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *MessageSelector) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *MessageSelector) SetOrigin(v string)`

SetOrigin sets Origin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


