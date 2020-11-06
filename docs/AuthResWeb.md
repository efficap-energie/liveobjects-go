# AuthResWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** |  | [optional] 
**ApiKey** | Pointer to [**ApiKey**](ApiKey.md) |  | [optional] 

## Methods

### NewAuthResWeb

`func NewAuthResWeb() *AuthResWeb`

NewAuthResWeb instantiates a new AuthResWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthResWebWithDefaults

`func NewAuthResWebWithDefaults() *AuthResWeb`

NewAuthResWebWithDefaults instantiates a new AuthResWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *AuthResWeb) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *AuthResWeb) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *AuthResWeb) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *AuthResWeb) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetApiKey

`func (o *AuthResWeb) GetApiKey() ApiKey`

GetApiKey returns the ApiKey field if non-nil, zero value otherwise.

### GetApiKeyOk

`func (o *AuthResWeb) GetApiKeyOk() (*ApiKey, bool)`

GetApiKeyOk returns a tuple with the ApiKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApiKey

`func (o *AuthResWeb) SetApiKey(v ApiKey)`

SetApiKey sets ApiKey field to given value.

### HasApiKey

`func (o *AuthResWeb) HasApiKey() bool`

HasApiKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


