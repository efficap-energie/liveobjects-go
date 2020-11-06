# ResourceVersionAddReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Checksum** | **string** | The Base64-encoded MD5 checksum of the version&#39;s raw file (non Base64-encoded)Expected string (max 255 characters) | 
**CompatibleVersions** | Pointer to **[]string** | set of the versions from which a resource update to the version can be done. Max number of versions is 100. A version max length is 255.  | [optional] 
**Description** | Pointer to **string** | the version&#39;s description. Expected string (max 255 characters) | [optional] 
**File** | **string** | the version&#39;s Base64-encoded binary. File max size is 10MB | 
**ResourceVersionId** | **string** | the resource version id. Expected string (max 255 characters) | 

## Methods

### NewResourceVersionAddReqWeb

`func NewResourceVersionAddReqWeb(checksum string, file string, resourceVersionId string, ) *ResourceVersionAddReqWeb`

NewResourceVersionAddReqWeb instantiates a new ResourceVersionAddReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceVersionAddReqWebWithDefaults

`func NewResourceVersionAddReqWebWithDefaults() *ResourceVersionAddReqWeb`

NewResourceVersionAddReqWebWithDefaults instantiates a new ResourceVersionAddReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetChecksum

`func (o *ResourceVersionAddReqWeb) GetChecksum() string`

GetChecksum returns the Checksum field if non-nil, zero value otherwise.

### GetChecksumOk

`func (o *ResourceVersionAddReqWeb) GetChecksumOk() (*string, bool)`

GetChecksumOk returns a tuple with the Checksum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetChecksum

`func (o *ResourceVersionAddReqWeb) SetChecksum(v string)`

SetChecksum sets Checksum field to given value.


### GetCompatibleVersions

`func (o *ResourceVersionAddReqWeb) GetCompatibleVersions() []string`

GetCompatibleVersions returns the CompatibleVersions field if non-nil, zero value otherwise.

### GetCompatibleVersionsOk

`func (o *ResourceVersionAddReqWeb) GetCompatibleVersionsOk() (*[]string, bool)`

GetCompatibleVersionsOk returns a tuple with the CompatibleVersions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompatibleVersions

`func (o *ResourceVersionAddReqWeb) SetCompatibleVersions(v []string)`

SetCompatibleVersions sets CompatibleVersions field to given value.

### HasCompatibleVersions

`func (o *ResourceVersionAddReqWeb) HasCompatibleVersions() bool`

HasCompatibleVersions returns a boolean if a field has been set.

### GetDescription

`func (o *ResourceVersionAddReqWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *ResourceVersionAddReqWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *ResourceVersionAddReqWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *ResourceVersionAddReqWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetFile

`func (o *ResourceVersionAddReqWeb) GetFile() string`

GetFile returns the File field if non-nil, zero value otherwise.

### GetFileOk

`func (o *ResourceVersionAddReqWeb) GetFileOk() (*string, bool)`

GetFileOk returns a tuple with the File field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFile

`func (o *ResourceVersionAddReqWeb) SetFile(v string)`

SetFile sets File field to given value.


### GetResourceVersionId

`func (o *ResourceVersionAddReqWeb) GetResourceVersionId() string`

GetResourceVersionId returns the ResourceVersionId field if non-nil, zero value otherwise.

### GetResourceVersionIdOk

`func (o *ResourceVersionAddReqWeb) GetResourceVersionIdOk() (*string, bool)`

GetResourceVersionIdOk returns a tuple with the ResourceVersionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResourceVersionId

`func (o *ResourceVersionAddReqWeb) SetResourceVersionId(v string)`

SetResourceVersionId sets ResourceVersionId field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


