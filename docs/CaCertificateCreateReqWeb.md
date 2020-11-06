# CaCertificateCreateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | **string** | PEM X.509 public certificate | 
**Comment** | **string** | Certificate description. | 

## Methods

### NewCaCertificateCreateReqWeb

`func NewCaCertificateCreateReqWeb(certificate string, comment string, ) *CaCertificateCreateReqWeb`

NewCaCertificateCreateReqWeb instantiates a new CaCertificateCreateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaCertificateCreateReqWebWithDefaults

`func NewCaCertificateCreateReqWebWithDefaults() *CaCertificateCreateReqWeb`

NewCaCertificateCreateReqWebWithDefaults instantiates a new CaCertificateCreateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CaCertificateCreateReqWeb) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CaCertificateCreateReqWeb) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CaCertificateCreateReqWeb) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.


### GetComment

`func (o *CaCertificateCreateReqWeb) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CaCertificateCreateReqWeb) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CaCertificateCreateReqWeb) SetComment(v string)`

SetComment sets Comment field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


