# UpdatePasswordReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Captcha** | **string** | security captcha got from user | 
**CaptchaToken** | **string** | captcha token, that will be checked for authorizing the update | 
**NewPassword** | **string** | new password string | 
**OldPassword** | **string** | old password string | 
**TokenId** | **string** | token identifier | 

## Methods

### NewUpdatePasswordReqWeb

`func NewUpdatePasswordReqWeb(captcha string, captchaToken string, newPassword string, oldPassword string, tokenId string, ) *UpdatePasswordReqWeb`

NewUpdatePasswordReqWeb instantiates a new UpdatePasswordReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordReqWebWithDefaults

`func NewUpdatePasswordReqWebWithDefaults() *UpdatePasswordReqWeb`

NewUpdatePasswordReqWebWithDefaults instantiates a new UpdatePasswordReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptcha

`func (o *UpdatePasswordReqWeb) GetCaptcha() string`

GetCaptcha returns the Captcha field if non-nil, zero value otherwise.

### GetCaptchaOk

`func (o *UpdatePasswordReqWeb) GetCaptchaOk() (*string, bool)`

GetCaptchaOk returns a tuple with the Captcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptcha

`func (o *UpdatePasswordReqWeb) SetCaptcha(v string)`

SetCaptcha sets Captcha field to given value.


### GetCaptchaToken

`func (o *UpdatePasswordReqWeb) GetCaptchaToken() string`

GetCaptchaToken returns the CaptchaToken field if non-nil, zero value otherwise.

### GetCaptchaTokenOk

`func (o *UpdatePasswordReqWeb) GetCaptchaTokenOk() (*string, bool)`

GetCaptchaTokenOk returns a tuple with the CaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaToken

`func (o *UpdatePasswordReqWeb) SetCaptchaToken(v string)`

SetCaptchaToken sets CaptchaToken field to given value.


### GetNewPassword

`func (o *UpdatePasswordReqWeb) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *UpdatePasswordReqWeb) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *UpdatePasswordReqWeb) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.


### GetOldPassword

`func (o *UpdatePasswordReqWeb) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *UpdatePasswordReqWeb) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *UpdatePasswordReqWeb) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.


### GetTokenId

`func (o *UpdatePasswordReqWeb) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UpdatePasswordReqWeb) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UpdatePasswordReqWeb) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


