# ResetPasswordReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Captcha** | **string** | Security captcha got from user | 
**CaptchaToken** | **string** | Security captcha waited token | 
**UserLogin** | **string** | the user login | 

## Methods

### NewResetPasswordReqWeb

`func NewResetPasswordReqWeb(captcha string, captchaToken string, userLogin string, ) *ResetPasswordReqWeb`

NewResetPasswordReqWeb instantiates a new ResetPasswordReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResetPasswordReqWebWithDefaults

`func NewResetPasswordReqWebWithDefaults() *ResetPasswordReqWeb`

NewResetPasswordReqWebWithDefaults instantiates a new ResetPasswordReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptcha

`func (o *ResetPasswordReqWeb) GetCaptcha() string`

GetCaptcha returns the Captcha field if non-nil, zero value otherwise.

### GetCaptchaOk

`func (o *ResetPasswordReqWeb) GetCaptchaOk() (*string, bool)`

GetCaptchaOk returns a tuple with the Captcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptcha

`func (o *ResetPasswordReqWeb) SetCaptcha(v string)`

SetCaptcha sets Captcha field to given value.


### GetCaptchaToken

`func (o *ResetPasswordReqWeb) GetCaptchaToken() string`

GetCaptchaToken returns the CaptchaToken field if non-nil, zero value otherwise.

### GetCaptchaTokenOk

`func (o *ResetPasswordReqWeb) GetCaptchaTokenOk() (*string, bool)`

GetCaptchaTokenOk returns a tuple with the CaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaToken

`func (o *ResetPasswordReqWeb) SetCaptchaToken(v string)`

SetCaptchaToken sets CaptchaToken field to given value.


### GetUserLogin

`func (o *ResetPasswordReqWeb) GetUserLogin() string`

GetUserLogin returns the UserLogin field if non-nil, zero value otherwise.

### GetUserLoginOk

`func (o *ResetPasswordReqWeb) GetUserLoginOk() (*string, bool)`

GetUserLoginOk returns a tuple with the UserLogin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserLogin

`func (o *ResetPasswordReqWeb) SetUserLogin(v string)`

SetUserLogin sets UserLogin field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


