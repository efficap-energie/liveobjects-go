# ErrorResponseWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | **string** | Error code | 
**Details** | Pointer to **string** | Detailed error description | [optional] 
**Id** | **string** | Unique identifier of this error instance | 
**Message** | **string** | Short error description | 

## Methods

### NewErrorResponseWeb

`func NewErrorResponseWeb(code string, id string, message string, ) *ErrorResponseWeb`

NewErrorResponseWeb instantiates a new ErrorResponseWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorResponseWebWithDefaults

`func NewErrorResponseWebWithDefaults() *ErrorResponseWeb`

NewErrorResponseWebWithDefaults instantiates a new ErrorResponseWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *ErrorResponseWeb) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *ErrorResponseWeb) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *ErrorResponseWeb) SetCode(v string)`

SetCode sets Code field to given value.


### GetDetails

`func (o *ErrorResponseWeb) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *ErrorResponseWeb) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *ErrorResponseWeb) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *ErrorResponseWeb) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetId

`func (o *ErrorResponseWeb) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ErrorResponseWeb) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ErrorResponseWeb) SetId(v string)`

SetId sets Id field to given value.


### GetMessage

`func (o *ErrorResponseWeb) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *ErrorResponseWeb) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *ErrorResponseWeb) SetMessage(v string)`

SetMessage sets Message field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


