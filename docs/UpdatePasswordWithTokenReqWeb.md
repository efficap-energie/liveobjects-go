# UpdatePasswordWithTokenReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Captcha** | **string** | security captcha got from user | 
**CaptchaToken** | **string** | token security captcha waited | 
**Company** | **string** | tenant company | 
**FirstName** | **string** | tenant main user first name | 
**LastName** | **string** | tenant main user last name | 
**Login** | **string** | user email | 
**Password** | **string** | new user password | 
**PhoneNumber** | Pointer to **string** | tenant main user phone number | [optional] 
**TokenId** | **string** | token identifier to for mail change | 

## Methods

### NewUpdatePasswordWithTokenReqWeb

`func NewUpdatePasswordWithTokenReqWeb(captcha string, captchaToken string, company string, firstName string, lastName string, login string, password string, tokenId string, ) *UpdatePasswordWithTokenReqWeb`

NewUpdatePasswordWithTokenReqWeb instantiates a new UpdatePasswordWithTokenReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdatePasswordWithTokenReqWebWithDefaults

`func NewUpdatePasswordWithTokenReqWebWithDefaults() *UpdatePasswordWithTokenReqWeb`

NewUpdatePasswordWithTokenReqWebWithDefaults instantiates a new UpdatePasswordWithTokenReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaptcha

`func (o *UpdatePasswordWithTokenReqWeb) GetCaptcha() string`

GetCaptcha returns the Captcha field if non-nil, zero value otherwise.

### GetCaptchaOk

`func (o *UpdatePasswordWithTokenReqWeb) GetCaptchaOk() (*string, bool)`

GetCaptchaOk returns a tuple with the Captcha field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptcha

`func (o *UpdatePasswordWithTokenReqWeb) SetCaptcha(v string)`

SetCaptcha sets Captcha field to given value.


### GetCaptchaToken

`func (o *UpdatePasswordWithTokenReqWeb) GetCaptchaToken() string`

GetCaptchaToken returns the CaptchaToken field if non-nil, zero value otherwise.

### GetCaptchaTokenOk

`func (o *UpdatePasswordWithTokenReqWeb) GetCaptchaTokenOk() (*string, bool)`

GetCaptchaTokenOk returns a tuple with the CaptchaToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaptchaToken

`func (o *UpdatePasswordWithTokenReqWeb) SetCaptchaToken(v string)`

SetCaptchaToken sets CaptchaToken field to given value.


### GetCompany

`func (o *UpdatePasswordWithTokenReqWeb) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *UpdatePasswordWithTokenReqWeb) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *UpdatePasswordWithTokenReqWeb) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetFirstName

`func (o *UpdatePasswordWithTokenReqWeb) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *UpdatePasswordWithTokenReqWeb) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *UpdatePasswordWithTokenReqWeb) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *UpdatePasswordWithTokenReqWeb) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *UpdatePasswordWithTokenReqWeb) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *UpdatePasswordWithTokenReqWeb) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetLogin

`func (o *UpdatePasswordWithTokenReqWeb) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *UpdatePasswordWithTokenReqWeb) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *UpdatePasswordWithTokenReqWeb) SetLogin(v string)`

SetLogin sets Login field to given value.


### GetPassword

`func (o *UpdatePasswordWithTokenReqWeb) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UpdatePasswordWithTokenReqWeb) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UpdatePasswordWithTokenReqWeb) SetPassword(v string)`

SetPassword sets Password field to given value.


### GetPhoneNumber

`func (o *UpdatePasswordWithTokenReqWeb) GetPhoneNumber() string`

GetPhoneNumber returns the PhoneNumber field if non-nil, zero value otherwise.

### GetPhoneNumberOk

`func (o *UpdatePasswordWithTokenReqWeb) GetPhoneNumberOk() (*string, bool)`

GetPhoneNumberOk returns a tuple with the PhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumber

`func (o *UpdatePasswordWithTokenReqWeb) SetPhoneNumber(v string)`

SetPhoneNumber sets PhoneNumber field to given value.

### HasPhoneNumber

`func (o *UpdatePasswordWithTokenReqWeb) HasPhoneNumber() bool`

HasPhoneNumber returns a boolean if a field has been set.

### GetTokenId

`func (o *UpdatePasswordWithTokenReqWeb) GetTokenId() string`

GetTokenId returns the TokenId field if non-nil, zero value otherwise.

### GetTokenIdOk

`func (o *UpdatePasswordWithTokenReqWeb) GetTokenIdOk() (*string, bool)`

GetTokenIdOk returns a tuple with the TokenId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTokenId

`func (o *UpdatePasswordWithTokenReqWeb) SetTokenId(v string)`

SetTokenId sets TokenId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


