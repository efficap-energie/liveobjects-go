# BinaryPayloadDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | decoder activation. Default is false. | 
**Encoding** | **string** | unique decoder name. Will be used to trigger the decoding function | 
**Format** | **string** | the JBBP binary payload format with LO additional tokens (float, utf8) | 
**Id** | Pointer to **string** | id of the decoder. Should be null when used for POST. Required for update. | [optional] 
**MathEvalEnabled** | Pointer to **bool** | applying math evaluation on templated decoded result | [optional] 
**Metadata** | Pointer to [**PayloadDescriptionMetadata**](PayloadDescriptionMetadata.md) |  | [optional] 
**Model** | Pointer to **string** | if empty, the decoded data will use the value of &#39;model&#39; field from encoded data. If set, this value will be used to override &#39;model&#39; field in decoded data. | [optional] 
**Tags** | Pointer to **[]string** | tags used for filtering | [optional] 
**Template** | Pointer to **string** | decoding result optional template | [optional] 
**Type** | Pointer to **string** | decoder type : csv, binary, javascript | [optional] 
**Updated** | Pointer to **string** | date of the last update. ISO 8601 date time. | [optional] [readonly] 

## Methods

### NewBinaryPayloadDescription

`func NewBinaryPayloadDescription(enabled bool, encoding string, format string, ) *BinaryPayloadDescription`

NewBinaryPayloadDescription instantiates a new BinaryPayloadDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBinaryPayloadDescriptionWithDefaults

`func NewBinaryPayloadDescriptionWithDefaults() *BinaryPayloadDescription`

NewBinaryPayloadDescriptionWithDefaults instantiates a new BinaryPayloadDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *BinaryPayloadDescription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BinaryPayloadDescription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BinaryPayloadDescription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEncoding

`func (o *BinaryPayloadDescription) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *BinaryPayloadDescription) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *BinaryPayloadDescription) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetFormat

`func (o *BinaryPayloadDescription) GetFormat() string`

GetFormat returns the Format field if non-nil, zero value otherwise.

### GetFormatOk

`func (o *BinaryPayloadDescription) GetFormatOk() (*string, bool)`

GetFormatOk returns a tuple with the Format field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFormat

`func (o *BinaryPayloadDescription) SetFormat(v string)`

SetFormat sets Format field to given value.


### GetId

`func (o *BinaryPayloadDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BinaryPayloadDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BinaryPayloadDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *BinaryPayloadDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMathEvalEnabled

`func (o *BinaryPayloadDescription) GetMathEvalEnabled() bool`

GetMathEvalEnabled returns the MathEvalEnabled field if non-nil, zero value otherwise.

### GetMathEvalEnabledOk

`func (o *BinaryPayloadDescription) GetMathEvalEnabledOk() (*bool, bool)`

GetMathEvalEnabledOk returns a tuple with the MathEvalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathEvalEnabled

`func (o *BinaryPayloadDescription) SetMathEvalEnabled(v bool)`

SetMathEvalEnabled sets MathEvalEnabled field to given value.

### HasMathEvalEnabled

`func (o *BinaryPayloadDescription) HasMathEvalEnabled() bool`

HasMathEvalEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *BinaryPayloadDescription) GetMetadata() PayloadDescriptionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *BinaryPayloadDescription) GetMetadataOk() (*PayloadDescriptionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *BinaryPayloadDescription) SetMetadata(v PayloadDescriptionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *BinaryPayloadDescription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *BinaryPayloadDescription) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *BinaryPayloadDescription) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *BinaryPayloadDescription) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *BinaryPayloadDescription) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *BinaryPayloadDescription) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *BinaryPayloadDescription) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *BinaryPayloadDescription) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *BinaryPayloadDescription) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *BinaryPayloadDescription) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *BinaryPayloadDescription) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *BinaryPayloadDescription) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *BinaryPayloadDescription) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *BinaryPayloadDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *BinaryPayloadDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *BinaryPayloadDescription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *BinaryPayloadDescription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *BinaryPayloadDescription) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *BinaryPayloadDescription) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *BinaryPayloadDescription) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *BinaryPayloadDescription) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


