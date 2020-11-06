# SearchDataMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**Extra** | Pointer to **map[string]string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** |  | [optional] 
**Model** | Pointer to **string** |  | [optional] 
**StreamId** | Pointer to **string** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 
**Value** | Pointer to **map[string]interface{}** |  | [optional] 

## Methods

### NewSearchDataMessage

`func NewSearchDataMessage() *SearchDataMessage`

NewSearchDataMessage instantiates a new SearchDataMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSearchDataMessageWithDefaults

`func NewSearchDataMessageWithDefaults() *SearchDataMessage`

NewSearchDataMessageWithDefaults instantiates a new SearchDataMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *SearchDataMessage) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *SearchDataMessage) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *SearchDataMessage) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *SearchDataMessage) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExtra

`func (o *SearchDataMessage) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *SearchDataMessage) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *SearchDataMessage) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *SearchDataMessage) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetId

`func (o *SearchDataMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SearchDataMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SearchDataMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *SearchDataMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *SearchDataMessage) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *SearchDataMessage) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *SearchDataMessage) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *SearchDataMessage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *SearchDataMessage) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *SearchDataMessage) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *SearchDataMessage) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *SearchDataMessage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *SearchDataMessage) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *SearchDataMessage) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *SearchDataMessage) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *SearchDataMessage) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStreamId

`func (o *SearchDataMessage) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *SearchDataMessage) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *SearchDataMessage) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *SearchDataMessage) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTags

`func (o *SearchDataMessage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *SearchDataMessage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *SearchDataMessage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *SearchDataMessage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *SearchDataMessage) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *SearchDataMessage) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *SearchDataMessage) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *SearchDataMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *SearchDataMessage) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SearchDataMessage) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SearchDataMessage) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *SearchDataMessage) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


