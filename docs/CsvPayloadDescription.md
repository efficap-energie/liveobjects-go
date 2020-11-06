# CsvPayloadDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]CsvColumn**](CsvColumn.md) | colums description for the csv format | 
**Enabled** | **bool** | decoder activation. Default is false. | 
**Encoding** | **string** | unique decoder name. Will be used to trigger the decoding function | 
**Id** | Pointer to **string** | id of the decoder. Should be null when used for POST. Required for update. | [optional] 
**MathEvalEnabled** | Pointer to **bool** | applying math evaluation on templated decoded result | [optional] 
**Metadata** | Pointer to [**PayloadDescriptionMetadata**](PayloadDescriptionMetadata.md) |  | [optional] 
**Model** | Pointer to **string** | if empty, the decoded data will use the value of &#39;model&#39; field from encoded data. If set, this value will be used to override &#39;model&#39; field in decoded data. | [optional] 
**Options** | Pointer to [**CsvOptions**](CsvOptions.md) |  | [optional] 
**Tags** | Pointer to **[]string** | tags used for filtering | [optional] 
**Template** | Pointer to **string** | decoding result optional template | [optional] 
**Type** | Pointer to **string** | decoder type : csv, binary, javascript | [optional] 
**Updated** | Pointer to **string** | date of the last update. ISO 8601 date time. | [optional] [readonly] 

## Methods

### NewCsvPayloadDescription

`func NewCsvPayloadDescription(columns []CsvColumn, enabled bool, encoding string, ) *CsvPayloadDescription`

NewCsvPayloadDescription instantiates a new CsvPayloadDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsvPayloadDescriptionWithDefaults

`func NewCsvPayloadDescriptionWithDefaults() *CsvPayloadDescription`

NewCsvPayloadDescriptionWithDefaults instantiates a new CsvPayloadDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *CsvPayloadDescription) GetColumns() []CsvColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CsvPayloadDescription) GetColumnsOk() (*[]CsvColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CsvPayloadDescription) SetColumns(v []CsvColumn)`

SetColumns sets Columns field to given value.


### GetEnabled

`func (o *CsvPayloadDescription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CsvPayloadDescription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CsvPayloadDescription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEncoding

`func (o *CsvPayloadDescription) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *CsvPayloadDescription) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *CsvPayloadDescription) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetId

`func (o *CsvPayloadDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CsvPayloadDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CsvPayloadDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CsvPayloadDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMathEvalEnabled

`func (o *CsvPayloadDescription) GetMathEvalEnabled() bool`

GetMathEvalEnabled returns the MathEvalEnabled field if non-nil, zero value otherwise.

### GetMathEvalEnabledOk

`func (o *CsvPayloadDescription) GetMathEvalEnabledOk() (*bool, bool)`

GetMathEvalEnabledOk returns a tuple with the MathEvalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathEvalEnabled

`func (o *CsvPayloadDescription) SetMathEvalEnabled(v bool)`

SetMathEvalEnabled sets MathEvalEnabled field to given value.

### HasMathEvalEnabled

`func (o *CsvPayloadDescription) HasMathEvalEnabled() bool`

HasMathEvalEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *CsvPayloadDescription) GetMetadata() PayloadDescriptionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *CsvPayloadDescription) GetMetadataOk() (*PayloadDescriptionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *CsvPayloadDescription) SetMetadata(v PayloadDescriptionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *CsvPayloadDescription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *CsvPayloadDescription) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *CsvPayloadDescription) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *CsvPayloadDescription) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *CsvPayloadDescription) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetOptions

`func (o *CsvPayloadDescription) GetOptions() CsvOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CsvPayloadDescription) GetOptionsOk() (*CsvOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CsvPayloadDescription) SetOptions(v CsvOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CsvPayloadDescription) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetTags

`func (o *CsvPayloadDescription) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *CsvPayloadDescription) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *CsvPayloadDescription) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *CsvPayloadDescription) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *CsvPayloadDescription) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CsvPayloadDescription) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CsvPayloadDescription) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CsvPayloadDescription) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *CsvPayloadDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *CsvPayloadDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *CsvPayloadDescription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *CsvPayloadDescription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *CsvPayloadDescription) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *CsvPayloadDescription) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *CsvPayloadDescription) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *CsvPayloadDescription) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


