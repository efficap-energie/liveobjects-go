# DataBulkItemWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Model** | Pointer to **string** | Model of the injected data. Data with the same model must have coherent types in each value fields. This field is required to perform an advanced search on the data. Can not contains space or dot characters. | [optional] 
**StreamId** | Pointer to **string** | StreamId in which the data will be added. Should not contains following character or space : &#39; \\ \&quot; ; { } ( ) | [optional] 
**Tags** | Pointer to **[]string** | Easy way to &#39;tag&#39; the data in order to ease up advanced search through all streams and models | [optional] 
**Timestamp** | Pointer to **string** | ISO 8601 date time | [optional] 
**Value** | **map[string]interface{}** | Must be a JsonObject (primitive objects like string or int are not allowed). It can not contain field names starting with &#39;$&#39; character or containing dot &#39;.&#39; character and its size should not exceed 1 MiB. To use decoding capability, &#39;value.payload&#39; field must be use to set encoded content as String (in case of binary content, HexBinary String representation of this content) | 

## Methods

### NewDataBulkItemWeb

`func NewDataBulkItemWeb(value map[string]interface{}, ) *DataBulkItemWeb`

NewDataBulkItemWeb instantiates a new DataBulkItemWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataBulkItemWebWithDefaults

`func NewDataBulkItemWebWithDefaults() *DataBulkItemWeb`

NewDataBulkItemWebWithDefaults instantiates a new DataBulkItemWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *DataBulkItemWeb) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DataBulkItemWeb) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DataBulkItemWeb) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DataBulkItemWeb) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *DataBulkItemWeb) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataBulkItemWeb) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataBulkItemWeb) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DataBulkItemWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *DataBulkItemWeb) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DataBulkItemWeb) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DataBulkItemWeb) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DataBulkItemWeb) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStreamId

`func (o *DataBulkItemWeb) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *DataBulkItemWeb) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *DataBulkItemWeb) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *DataBulkItemWeb) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTags

`func (o *DataBulkItemWeb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DataBulkItemWeb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DataBulkItemWeb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DataBulkItemWeb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *DataBulkItemWeb) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DataBulkItemWeb) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DataBulkItemWeb) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DataBulkItemWeb) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *DataBulkItemWeb) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataBulkItemWeb) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataBulkItemWeb) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


