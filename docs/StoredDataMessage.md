# StoredDataMessage

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

### NewStoredDataMessage

`func NewStoredDataMessage() *StoredDataMessage`

NewStoredDataMessage instantiates a new StoredDataMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStoredDataMessageWithDefaults

`func NewStoredDataMessageWithDefaults() *StoredDataMessage`

NewStoredDataMessageWithDefaults instantiates a new StoredDataMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *StoredDataMessage) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *StoredDataMessage) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *StoredDataMessage) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *StoredDataMessage) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetExtra

`func (o *StoredDataMessage) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *StoredDataMessage) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *StoredDataMessage) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *StoredDataMessage) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetId

`func (o *StoredDataMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *StoredDataMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *StoredDataMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *StoredDataMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLocation

`func (o *StoredDataMessage) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *StoredDataMessage) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *StoredDataMessage) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *StoredDataMessage) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *StoredDataMessage) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *StoredDataMessage) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *StoredDataMessage) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *StoredDataMessage) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *StoredDataMessage) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *StoredDataMessage) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *StoredDataMessage) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *StoredDataMessage) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStreamId

`func (o *StoredDataMessage) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *StoredDataMessage) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *StoredDataMessage) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *StoredDataMessage) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTags

`func (o *StoredDataMessage) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *StoredDataMessage) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *StoredDataMessage) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *StoredDataMessage) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *StoredDataMessage) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *StoredDataMessage) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *StoredDataMessage) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *StoredDataMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *StoredDataMessage) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *StoredDataMessage) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *StoredDataMessage) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *StoredDataMessage) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


