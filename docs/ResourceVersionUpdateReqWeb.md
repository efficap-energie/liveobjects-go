# ResourceVersionUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CompatibleVersions** | Pointer to **[]string** | set of the versions from which a resource update to the version can be done. Max number of versions is 100. A version max length is 255.  | [optional] 
**Description** | Pointer to **string** | the version&#39;s description. Expected string (max 255 characters) | [optional] 

## Methods

### NewResourceVersionUpdateReqWeb

`func NewResourceVersionUpdateReqWeb() *ResourceVersionUpdateReqWeb`

NewResourceVersionUpdateReqWeb instantiates a new ResourceVersionUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVersionUpdateReqWebWithDefaults

`func NewResourceVersionUpdateReqWebWithDefaults() *ResourceVersionUpdateReqWeb`

NewResourceVersionUpdateReqWebWithDefaults instantiates a new ResourceVersionUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompatibleVersions

`func (o *ResourceVersionUpdateReqWeb) GetCompatibleVersions() []string`

GetCompatibleVersions returns the CompatibleVersions field if non-nil, zero value otherwise.

### GetCompatibleVersionsOk

`func (o *ResourceVersionUpdateReqWeb) GetCompatibleVersionsOk() (*[]string, bool)`

GetCompatibleVersionsOk returns a tuple with the CompatibleVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleVersions

`func (o *ResourceVersionUpdateReqWeb) SetCompatibleVersions(v []string)`

SetCompatibleVersions sets CompatibleVersions field to given value.

### HasCompatibleVersions

`func (o *ResourceVersionUpdateReqWeb) HasCompatibleVersions() bool`

HasCompatibleVersions returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceVersionUpdateReqWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceVersionUpdateReqWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceVersionUpdateReqWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceVersionUpdateReqWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


