# OAuth2ErrorWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Error** | Pointer to **string** | Error value | [optional] 
**ErrorDescription** | Pointer to **string** | Error description | [optional] 

## Methods

### NewOAuth2ErrorWeb

`func NewOAuth2ErrorWeb() *OAuth2ErrorWeb`

NewOAuth2ErrorWeb instantiates a new OAuth2ErrorWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOAuth2ErrorWebWithDefaults

`func NewOAuth2ErrorWebWithDefaults() *OAuth2ErrorWeb`

NewOAuth2ErrorWebWithDefaults instantiates a new OAuth2ErrorWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetError

`func (o *OAuth2ErrorWeb) GetError() string`

GetError returns the Error field if non-nil, zero value otherwise.

### GetErrorOk

`func (o *OAuth2ErrorWeb) GetErrorOk() (*string, bool)`

GetErrorOk returns a tuple with the Error field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetError

`func (o *OAuth2ErrorWeb) SetError(v string)`

SetError sets Error field to given value.

### HasError

`func (o *OAuth2ErrorWeb) HasError() bool`

HasError returns a boolean if a field has been set.

### GetErrorDescription

`func (o *OAuth2ErrorWeb) GetErrorDescription() string`

GetErrorDescription returns the ErrorDescription field if non-nil, zero value otherwise.

### GetErrorDescriptionOk

`func (o *OAuth2ErrorWeb) GetErrorDescriptionOk() (*string, bool)`

GetErrorDescriptionOk returns a tuple with the ErrorDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDescription

`func (o *OAuth2ErrorWeb) SetErrorDescription(v string)`

SetErrorDescription sets ErrorDescription field to given value.

### HasErrorDescription

`func (o *OAuth2ErrorWeb) HasErrorDescription() bool`

HasErrorDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


