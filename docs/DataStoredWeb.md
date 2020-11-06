# DataStoredWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | **string** | ISO 8601 date time, when the data has been stored in the system | 
**Extra** | Pointer to **map[string]string** | Extra info of the data: extra information enriched on the collected data (for example from the device inventory) that can be used for cross model or stream queries | [optional] 
**Id** | **string** | System id of the data. Can be use as bookmark to paginate results. | 
**Location** | Pointer to [**Location**](Location.md) |  | [optional] 
**Metadata** | Pointer to **map[string]interface{}** | Metadata of the stored value : may include user metadata (source) as well as additional information (protocol, server, user...) | [optional] 
**Model** | Pointer to **string** | Model of the injected data. Data with the same model have coherent types in each value fields. | [optional] 
**StreamId** | Pointer to **string** | streamId of the data | [optional] 
**Tags** | Pointer to **[]string** | Tags of the data. Can be used to ease up advanced search through all streams and models | [optional] 
**Timestamp** | **string** | ISO 8601 date time, timestamp of the value | 
**Value** | **map[string]interface{}** | JsonObject | 

## Methods

### NewDataStoredWeb

`func NewDataStoredWeb(created string, id string, timestamp string, value map[string]interface{}, ) *DataStoredWeb`

NewDataStoredWeb instantiates a new DataStoredWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataStoredWebWithDefaults

`func NewDataStoredWebWithDefaults() *DataStoredWeb`

NewDataStoredWebWithDefaults instantiates a new DataStoredWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *DataStoredWeb) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *DataStoredWeb) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *DataStoredWeb) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetExtra

`func (o *DataStoredWeb) GetExtra() map[string]string`

GetExtra returns the Extra field if non-nil, zero value otherwise.

### GetExtraOk

`func (o *DataStoredWeb) GetExtraOk() (*map[string]string, bool)`

GetExtraOk returns a tuple with the Extra field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtra

`func (o *DataStoredWeb) SetExtra(v map[string]string)`

SetExtra sets Extra field to given value.

### HasExtra

`func (o *DataStoredWeb) HasExtra() bool`

HasExtra returns a boolean if a field has been set.

### GetId

`func (o *DataStoredWeb) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DataStoredWeb) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DataStoredWeb) SetId(v string)`

SetId sets Id field to given value.


### GetLocation

`func (o *DataStoredWeb) GetLocation() Location`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *DataStoredWeb) GetLocationOk() (*Location, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *DataStoredWeb) SetLocation(v Location)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *DataStoredWeb) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetMetadata

`func (o *DataStoredWeb) GetMetadata() map[string]interface{}`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *DataStoredWeb) GetMetadataOk() (*map[string]interface{}, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *DataStoredWeb) SetMetadata(v map[string]interface{})`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *DataStoredWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *DataStoredWeb) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *DataStoredWeb) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *DataStoredWeb) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *DataStoredWeb) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetStreamId

`func (o *DataStoredWeb) GetStreamId() string`

GetStreamId returns the StreamId field if non-nil, zero value otherwise.

### GetStreamIdOk

`func (o *DataStoredWeb) GetStreamIdOk() (*string, bool)`

GetStreamIdOk returns a tuple with the StreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStreamId

`func (o *DataStoredWeb) SetStreamId(v string)`

SetStreamId sets StreamId field to given value.

### HasStreamId

`func (o *DataStoredWeb) HasStreamId() bool`

HasStreamId returns a boolean if a field has been set.

### GetTags

`func (o *DataStoredWeb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DataStoredWeb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DataStoredWeb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DataStoredWeb) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTimestamp

`func (o *DataStoredWeb) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *DataStoredWeb) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *DataStoredWeb) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetValue

`func (o *DataStoredWeb) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *DataStoredWeb) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *DataStoredWeb) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


