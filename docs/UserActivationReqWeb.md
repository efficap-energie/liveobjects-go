# UserActivationReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantId** | **string** | identifier of tenant account this user belongs to. Expected identifier (max 24 characters) | 
**UserId** | **string** | user identifier in tenant account. Expected identifier (max 24 characters) | 

## Methods

### NewUserActivationReqWeb

`func NewUserActivationReqWeb(tenantId string, userId string, ) *UserActivationReqWeb`

NewUserActivationReqWeb instantiates a new UserActivationReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserActivationReqWebWithDefaults

`func NewUserActivationReqWebWithDefaults() *UserActivationReqWeb`

NewUserActivationReqWebWithDefaults instantiates a new UserActivationReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantId

`func (o *UserActivationReqWeb) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *UserActivationReqWeb) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *UserActivationReqWeb) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.


### GetUserId

`func (o *UserActivationReqWeb) GetUserId() string`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserActivationReqWeb) GetUserIdOk() (*string, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserActivationReqWeb) SetUserId(v string)`

SetUserId sets UserId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


