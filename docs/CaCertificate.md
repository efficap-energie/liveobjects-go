# CaCertificate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certificate** | Pointer to **string** |  | [optional] 
**Comment** | Pointer to **string** |  | [optional] 
**Expires** | Pointer to [**time.Time**](time.Time.md) |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 

## Methods

### NewCaCertificate

`func NewCaCertificate() *CaCertificate`

NewCaCertificate instantiates a new CaCertificate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCaCertificateWithDefaults

`func NewCaCertificateWithDefaults() *CaCertificate`

NewCaCertificateWithDefaults instantiates a new CaCertificate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertificate

`func (o *CaCertificate) GetCertificate() string`

GetCertificate returns the Certificate field if non-nil, zero value otherwise.

### GetCertificateOk

`func (o *CaCertificate) GetCertificateOk() (*string, bool)`

GetCertificateOk returns a tuple with the Certificate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertificate

`func (o *CaCertificate) SetCertificate(v string)`

SetCertificate sets Certificate field to given value.

### HasCertificate

`func (o *CaCertificate) HasCertificate() bool`

HasCertificate returns a boolean if a field has been set.

### GetComment

`func (o *CaCertificate) GetComment() string`

GetComment returns the Comment field if non-nil, zero value otherwise.

### GetCommentOk

`func (o *CaCertificate) GetCommentOk() (*string, bool)`

GetCommentOk returns a tuple with the Comment field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetComment

`func (o *CaCertificate) SetComment(v string)`

SetComment sets Comment field to given value.

### HasComment

`func (o *CaCertificate) HasComment() bool`

HasComment returns a boolean if a field has been set.

### GetExpires

`func (o *CaCertificate) GetExpires() time.Time`

GetExpires returns the Expires field if non-nil, zero value otherwise.

### GetExpiresOk

`func (o *CaCertificate) GetExpiresOk() (*time.Time, bool)`

GetExpiresOk returns a tuple with the Expires field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpires

`func (o *CaCertificate) SetExpires(v time.Time)`

SetExpires sets Expires field to given value.

### HasExpires

`func (o *CaCertificate) HasExpires() bool`

HasExpires returns a boolean if a field has been set.

### GetId

`func (o *CaCertificate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CaCertificate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CaCertificate) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CaCertificate) HasId() bool`

HasId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


