# ContextContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContextData** | **map[string]interface{}** | context data. | 
**ContextKey** | Pointer to **string** | key for accessing context data. Should be null when used for PUT. | [optional] 
**Tags** | Pointer to **[]string** | optional tags list for free use | [optional] 

## Methods

### NewContextContainer

`func NewContextContainer(contextData map[string]interface{}, ) *ContextContainer`

NewContextContainer instantiates a new ContextContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextContainerWithDefaults

`func NewContextContainerWithDefaults() *ContextContainer`

NewContextContainerWithDefaults instantiates a new ContextContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContextData

`func (o *ContextContainer) GetContextData() map[string]interface{}`

GetContextData returns the ContextData field if non-nil, zero value otherwise.

### GetContextDataOk

`func (o *ContextContainer) GetContextDataOk() (*map[string]interface{}, bool)`

GetContextDataOk returns a tuple with the ContextData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextData

`func (o *ContextContainer) SetContextData(v map[string]interface{})`

SetContextData sets ContextData field to given value.


### GetContextKey

`func (o *ContextContainer) GetContextKey() string`

GetContextKey returns the ContextKey field if non-nil, zero value otherwise.

### GetContextKeyOk

`func (o *ContextContainer) GetContextKeyOk() (*string, bool)`

GetContextKeyOk returns a tuple with the ContextKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContextKey

`func (o *ContextContainer) SetContextKey(v string)`

SetContextKey sets ContextKey field to given value.

### HasContextKey

`func (o *ContextContainer) HasContextKey() bool`

HasContextKey returns a boolean if a field has been set.

### GetTags

`func (o *ContextContainer) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *ContextContainer) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *ContextContainer) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *ContextContainer) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


