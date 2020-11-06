# AssetUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** | user-readable description of the asset. Expected string (max 500 characters) | [optional] 
**GroupId** | Pointer to **string** | identifier of the group this asset is associated with. Expected string (max 6 characters) | [optional] 
**GroupPath** | Pointer to **string** | path of the group this asset is associated with. Authorized: letter (lowercase and uppercase), accented characters, number, space, dash, underscore and simple quote. A valid path must respect the following regular expression &lt;code&gt;^[\\wÀ-ÖØ-öø-ÿ&#39; -]{1,255}&lt;/code&gt;.Expected string (max 255 characters) | [optional] 
**Name** | Pointer to **string** | asset user-readable name (for display in web portal). Expected string (max 255 characters) | [optional] 
**Properties** | Pointer to **map[string]string** | map of key/value string pairs detailed properties of the device. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | [optional] 
**Tags** | Pointer to **[]string** | series of tags associated with the asset. Max number of tags depends of your offer settings. Tag value max length is 32. | [optional] 

## Methods

### NewAssetUpdateReqWeb

`func NewAssetUpdateReqWeb() *AssetUpdateReqWeb`

NewAssetUpdateReqWeb instantiates a new AssetUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetUpdateReqWebWithDefaults

`func NewAssetUpdateReqWebWithDefaults() *AssetUpdateReqWeb`

NewAssetUpdateReqWebWithDefaults instantiates a new AssetUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *AssetUpdateReqWeb) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AssetUpdateReqWeb) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AssetUpdateReqWeb) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *AssetUpdateReqWeb) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroupId

`func (o *AssetUpdateReqWeb) GetGroupId() string`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *AssetUpdateReqWeb) GetGroupIdOk() (*string, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *AssetUpdateReqWeb) SetGroupId(v string)`

SetGroupId sets GroupId field to given value.

### HasGroupId

`func (o *AssetUpdateReqWeb) HasGroupId() bool`

HasGroupId returns a boolean if a field has been set.

### GetGroupPath

`func (o *AssetUpdateReqWeb) GetGroupPath() string`

GetGroupPath returns the GroupPath field if non-nil, zero value otherwise.

### GetGroupPathOk

`func (o *AssetUpdateReqWeb) GetGroupPathOk() (*string, bool)`

GetGroupPathOk returns a tuple with the GroupPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupPath

`func (o *AssetUpdateReqWeb) SetGroupPath(v string)`

SetGroupPath sets GroupPath field to given value.

### HasGroupPath

`func (o *AssetUpdateReqWeb) HasGroupPath() bool`

HasGroupPath returns a boolean if a field has been set.

### GetName

`func (o *AssetUpdateReqWeb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AssetUpdateReqWeb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AssetUpdateReqWeb) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AssetUpdateReqWeb) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *AssetUpdateReqWeb) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *AssetUpdateReqWeb) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *AssetUpdateReqWeb) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *AssetUpdateReqWeb) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *AssetUpdateReqWeb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *AssetUpdateReqWeb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *AssetUpdateReqWeb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *AssetUpdateReqWeb) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


