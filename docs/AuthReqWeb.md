# AuthReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Email** | Pointer to **string** | the user email | [optional] 
**Login** | Pointer to **string** | the user login | [optional] 
**Password** | **string** | the user password | 

## Methods

### NewAuthReqWeb

`func NewAuthReqWeb(password string, ) *AuthReqWeb`

NewAuthReqWeb instantiates a new AuthReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthReqWebWithDefaults

`func NewAuthReqWebWithDefaults() *AuthReqWeb`

NewAuthReqWebWithDefaults instantiates a new AuthReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmail

`func (o *AuthReqWeb) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *AuthReqWeb) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *AuthReqWeb) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *AuthReqWeb) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetLogin

`func (o *AuthReqWeb) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *AuthReqWeb) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *AuthReqWeb) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *AuthReqWeb) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetPassword

`func (o *AuthReqWeb) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *AuthReqWeb) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *AuthReqWeb) SetPassword(v string)`

SetPassword sets Password field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


