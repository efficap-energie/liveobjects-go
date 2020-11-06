# PartnerDataItemSwagger

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metadata** | [**Metadata**](Metadata.md) |  | 
**Model** | Pointer to **string** | Model of the injected data. Data with the same model must have coherent types in each value fields. This field is required to perform an advanced search on the data. Can not contains space or dot characters. Should not be empty or blank. | [optional] 
**StreamId** | Pointer to **string** | StreamId in which the data will be added. Should not contains following character or space : &#39; \\ \&quot; ; { } ( ) | [optional] 
**Tags** | Pointer to **[]string** | Easy way to &#39;tag&#39; the data in order to ease up advanced search through all streams and models | [optional] 
**Timestamp** | Pointer to **string** | ISO 8601 date time | [optional] 
**Value** | **map[string]interface{}** | Must be a JsonObject (primitive objects like string or int are not allowed) | 

## Methods

### NewPartnerDataItemSwagger

`func NewPartnerDataItemSwagger(metadata Metadata, value map[string]interface{}, ) *PartnerDataItemSwagger`

NewPartnerDataItemSwagger instantiates a new PartnerDataItemSwagger object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPartnerDataItemSwaggerWithDefaults

`func NewPartnerDataItemSwaggerWithDefaults() *PartnerDataItemSwagger`

NewPartnerDataItemSwaggerWithDefaults instantiates a new PartnerDataItemSwagger object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetadata

`func (o *PartnerDataItemSwagger) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PartnerDataItemSwagger) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PartnerDataItemSwagger) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.


### GetModel

`func (o *PartnerDataItemSwagger) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PartnerDataItemSwagger) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PartnerDataItemSwagger) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PartnerDataItemSwagger) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStreamId

`func (o *PartnerDataItemSwagger) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *PartnerDataItemSwagger) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *PartnerDataItemSwagger) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *PartnerDataItemSwagger) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTags

`func (o *PartnerDataItemSwagger) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PartnerDataItemSwagger) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PartnerDataItemSwagger) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PartnerDataItemSwagger) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *PartnerDataItemSwagger) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *PartnerDataItemSwagger) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *PartnerDataItemSwagger) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *PartnerDataItemSwagger) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *PartnerDataItemSwagger) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *PartnerDataItemSwagger) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *PartnerDataItemSwagger) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


