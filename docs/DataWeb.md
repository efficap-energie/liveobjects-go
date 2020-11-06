# DataWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Metadata** | Pointer to [**Metadata**](Metadata.md) |  | [optional] 
**Model** | Pointer to **string** | Model of the injected data. Data with the same model must have coherent types in each value fields. This field is required to perform an advanced search on the data. Can not contains space or dot characters. | [optional] 
**Tags** | Pointer to **[]string** | Easy way to &#39;tag&#39; the data in order to ease up advanced search through all streams and models | [optional] 
**Timestamp** | Pointer to **string** | ISO 8601 date time | [optional] 
**Value** | **map[string]interface{}** | Must be a JsonObject (primitive objects like string or int are not allowed). It can not contain field names starting with &#39;$&#39; character or containing dot &#39;.&#39; character and its size should not exceed 1 MiB. To use decoding capability, &#39;value.payload&#39; field must be use to set encoded content as String (in case of binary content, HexBinary String representation of this content) | 

## Methods

### NewDataWeb

`func NewDataWeb(value map[string]interface{}, ) *DataWeb`

NewDataWeb instantiates a new DataWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataWebWithDefaults

`func NewDataWebWithDefaults() *DataWeb`

NewDataWebWithDefaults instantiates a new DataWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLocation

`func (o *DataWeb) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DataWeb) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DataWeb) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DataWeb) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *DataWeb) GetMetadata() Metadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataWeb) GetMetadataOk() (*Metadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataWeb) SetMetadata(v Metadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DataWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *DataWeb) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DataWeb) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DataWeb) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DataWeb) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *DataWeb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DataWeb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DataWeb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DataWeb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *DataWeb) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DataWeb) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DataWeb) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *DataWeb) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetValue

`func (o *DataWeb) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataWeb) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataWeb) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


