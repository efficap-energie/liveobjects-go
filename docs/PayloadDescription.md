# PayloadDescription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | **bool** | decoder activation. Default is false. | 
**Encoding** | **string** | unique decoder name. Will be used to trigger the decoding function | 
**Id** | Pointer to **string** | id of the decoder. Should be null when used for POST. Required for update. | [optional] 
**MathEvalEnabled** | Pointer to **bool** | applying math evaluation on templated decoded result | [optional] 
**Metadata** | Pointer to [**PayloadDescriptionMetadata**](PayloadDescriptionMetadata.md) |  | [optional] 
**Model** | Pointer to **string** | if empty, the decoded data will use the value of &#39;model&#39; field from encoded data. If set, this value will be used to override &#39;model&#39; field in decoded data. | [optional] 
**Tags** | Pointer to **[]string** | tags used for filtering | [optional] 
**Template** | Pointer to **string** | decoding result optional template | [optional] 
**Type** | Pointer to **string** | decoder type : csv, binary, javascript | [optional] 
**Updated** | Pointer to **string** | date of the last update. ISO 8601 date time. | [optional] [readonly] 

## Methods

### NewPayloadDescription

`func NewPayloadDescription(enabled bool, encoding string, ) *PayloadDescription`

NewPayloadDescription instantiates a new PayloadDescription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadDescriptionWithDefaults

`func NewPayloadDescriptionWithDefaults() *PayloadDescription`

NewPayloadDescriptionWithDefaults instantiates a new PayloadDescription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *PayloadDescription) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *PayloadDescription) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *PayloadDescription) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.


### GetEncoding

`func (o *PayloadDescription) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *PayloadDescription) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *PayloadDescription) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetId

`func (o *PayloadDescription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *PayloadDescription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *PayloadDescription) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *PayloadDescription) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMathEvalEnabled

`func (o *PayloadDescription) GetMathEvalEnabled() bool`

GetMathEvalEnabled returns the MathEvalEnabled field if non-nil, zero value otherwise.

### GetMathEvalEnabledOk

`func (o *PayloadDescription) GetMathEvalEnabledOk() (*bool, bool)`

GetMathEvalEnabledOk returns a tuple with the MathEvalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMathEvalEnabled

`func (o *PayloadDescription) SetMathEvalEnabled(v bool)`

SetMathEvalEnabled sets MathEvalEnabled field to given value.

### HasMathEvalEnabled

`func (o *PayloadDescription) HasMathEvalEnabled() bool`

HasMathEvalEnabled returns a boolean if a field has been set.

### GetMetadata

`func (o *PayloadDescription) GetMetadata() PayloadDescriptionMetadata`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *PayloadDescription) GetMetadataOk() (*PayloadDescriptionMetadata, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *PayloadDescription) SetMetadata(v PayloadDescriptionMetadata)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *PayloadDescription) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetModel

`func (o *PayloadDescription) GetModel() string`

GetModel returns the Model field if non-nil, zero value otherwise.

### GetModelOk

`func (o *PayloadDescription) GetModelOk() (*string, bool)`

GetModelOk returns a tuple with the Model field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetModel

`func (o *PayloadDescription) SetModel(v string)`

SetModel sets Model field to given value.

### HasModel

`func (o *PayloadDescription) HasModel() bool`

HasModel returns a boolean if a field has been set.

### GetTags

`func (o *PayloadDescription) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *PayloadDescription) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *PayloadDescription) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *PayloadDescription) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetTemplate

`func (o *PayloadDescription) GetTemplate() string`

GetTemplate returns the Template field if non-nil, zero value otherwise.

### GetTemplateOk

`func (o *PayloadDescription) GetTemplateOk() (*string, bool)`

GetTemplateOk returns a tuple with the Template field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTemplate

`func (o *PayloadDescription) SetTemplate(v string)`

SetTemplate sets Template field to given value.

### HasTemplate

`func (o *PayloadDescription) HasTemplate() bool`

HasTemplate returns a boolean if a field has been set.

### GetType

`func (o *PayloadDescription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *PayloadDescription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *PayloadDescription) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *PayloadDescription) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUpdated

`func (o *PayloadDescription) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *PayloadDescription) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *PayloadDescription) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *PayloadDescription) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


