# UpdateEmailReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Captcha** | **string** | Security captcha got from user | 
**CaptchaToken** | **string** | security captcha waited | 
**TokenId** | **string** | token to for mail change | 

## Methods

### NewUpdateEmailReqWeb

`func NewUpdateEmailReqWeb(captcha string, captchaToken string, tokenId string, ) *UpdateEmailReqWeb`

NewUpdateEmailReqWeb instantiates a new UpdateEmailReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateEmailReqWebWithDefaults

`func NewUpdateEmailReqWebWithDefaults() *UpdateEmailReqWeb`

NewUpdateEmailReqWebWithDefaults instantiates a new UpdateEmailReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptcha

`func (o *UpdateEmailReqWeb) GetCaptcha() string`

GetCaptcha returns the Captcha field if non-nil, zero value otherwise.

### GetCaptchaOk

`func (o *UpdateEmailReqWeb) GetCaptchaOk() (*string, bool)`

GetCaptchaOk returns a tuple with the Captcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptcha

`func (o *UpdateEmailReqWeb) SetCaptcha(v string)`

SetCaptcha sets Captcha field to given value.


### GetCaptchaToken

`func (o *UpdateEmailReqWeb) GetCaptchaToken() string`

GetCaptchaToken returns the CaptchaToken field if non-nil, zero value otherwise.

### GetCaptchaTokenOk

`func (o *UpdateEmailReqWeb) GetCaptchaTokenOk() (*string, bool)`

GetCaptchaTokenOk returns a tuple with the CaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaToken

`func (o *UpdateEmailReqWeb) SetCaptchaToken(v string)`

SetCaptchaToken sets CaptchaToken field to given value.


### GetTokenId

`func (o *UpdateEmailReqWeb) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UpdateEmailReqWeb) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UpdateEmailReqWeb) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


