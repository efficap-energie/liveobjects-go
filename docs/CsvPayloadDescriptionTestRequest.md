# CsvPayloadDescriptionTestRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Columns** | [**[]CsvColumn**](CsvColumn.md) | the colums description of the csv format | 
**CsvPayload** | Pointer to **string** |  | [optional] 
**MathEvalEnabled** | Pointer to **bool** |  | [optional] 
**Options** | Pointer to [**CsvOptions**](CsvOptions.md) |  | [optional] 
**Template** | Pointer to **string** |  | [optional] 

## Methods

### NewCsvPayloadDescriptionTestRequest

`func NewCsvPayloadDescriptionTestRequest(columns []CsvColumn, ) *CsvPayloadDescriptionTestRequest`

NewCsvPayloadDescriptionTestRequest instantiates a new CsvPayloadDescriptionTestRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCsvPayloadDescriptionTestRequestWithDefaults

`func NewCsvPayloadDescriptionTestRequestWithDefaults() *CsvPayloadDescriptionTestRequest`

NewCsvPayloadDescriptionTestRequestWithDefaults instantiates a new CsvPayloadDescriptionTestRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetColumns

`func (o *CsvPayloadDescriptionTestRequest) GetColumns() []CsvColumn`

GetColumns returns the Columns field if non-nil, zero value otherwise.

### GetColumnsOk

`func (o *CsvPayloadDescriptionTestRequest) GetColumnsOk() (*[]CsvColumn, bool)`

GetColumnsOk returns a tuple with the Columns field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetColumns

`func (o *CsvPayloadDescriptionTestRequest) SetColumns(v []CsvColumn)`

SetColumns sets Columns field to given value.


### GetCsvPayload

`func (o *CsvPayloadDescriptionTestRequest) GetCsvPayload() string`

GetCsvPayload returns the CsvPayload field if non-nil, zero value otherwise.

### GetCsvPayloadOk

`func (o *CsvPayloadDescriptionTestRequest) GetCsvPayloadOk() (*string, bool)`

GetCsvPayloadOk returns a tuple with the CsvPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCsvPayload

`func (o *CsvPayloadDescriptionTestRequest) SetCsvPayload(v string)`

SetCsvPayload sets CsvPayload field to given value.

### HasCsvPayload

`func (o *CsvPayloadDescriptionTestRequest) HasCsvPayload() bool`

HasCsvPayload returns a boolean if a field has been set.

### GetMathEvalEnabled

`func (o *CsvPayloadDescriptionTestRequest) GetMathEvalEnabled() bool`

GetMathEvalEnabled returns the MathEvalEnabled field if non-nil, zero value otherwise.

### GetMathEvalEnabledOk

`func (o *CsvPayloadDescriptionTestRequest) GetMathEvalEnabledOk() (*bool, bool)`

GetMathEvalEnabledOk returns a tuple with the MathEvalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathEvalEnabled

`func (o *CsvPayloadDescriptionTestRequest) SetMathEvalEnabled(v bool)`

SetMathEvalEnabled sets MathEvalEnabled field to given value.

### HasMathEvalEnabled

`func (o *CsvPayloadDescriptionTestRequest) HasMathEvalEnabled() bool`

HasMathEvalEnabled returns a boolean if a field has been set.

### GetOptions

`func (o *CsvPayloadDescriptionTestRequest) GetOptions() CsvOptions`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *CsvPayloadDescriptionTestRequest) GetOptionsOk() (*CsvOptions, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *CsvPayloadDescriptionTestRequest) SetOptions(v CsvOptions)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *CsvPayloadDescriptionTestRequest) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

### GetTemplate

`func (o *CsvPayloadDescriptionTestRequest) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *CsvPayloadDescriptionTestRequest) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *CsvPayloadDescriptionTestRequest) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *CsvPayloadDescriptionTestRequest) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


