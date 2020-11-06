# ExternalIdentity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | external identity identifier | [optional] 
**Login** | Pointer to **string** | external identity login | [optional] 
**Provider** | Pointer to **string** | external identity provider | [optional] 

## Methods

### NewExternalIdentity

`func NewExternalIdentity() *ExternalIdentity`

NewExternalIdentity instantiates a new ExternalIdentity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalIdentityWithDefaults

`func NewExternalIdentityWithDefaults() *ExternalIdentity`

NewExternalIdentityWithDefaults instantiates a new ExternalIdentity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ExternalIdentity) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ExternalIdentity) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ExternalIdentity) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ExternalIdentity) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLogin

`func (o *ExternalIdentity) GetLogin() string`

GetLogin returns the Login field if non-nil, zero value otherwise.

### GetLoginOk

`func (o *ExternalIdentity) GetLoginOk() (*string, bool)`

GetLoginOk returns a tuple with the Login field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLogin

`func (o *ExternalIdentity) SetLogin(v string)`

SetLogin sets Login field to given value.

### HasLogin

`func (o *ExternalIdentity) HasLogin() bool`

HasLogin returns a boolean if a field has been set.

### GetProvider

`func (o *ExternalIdentity) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *ExternalIdentity) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *ExternalIdentity) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *ExternalIdentity) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


