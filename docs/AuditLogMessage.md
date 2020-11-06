# AuditLogMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Category** | **string** | category of the log message. | 
**Content** | Pointer to **map[string]interface{}** | log message&#39;s specific content | [optional] 
**Created** | **string** | date of storage of the log message. ISO 8601 date time. | 
**Description** | **string** | high level description of the log message | 
**DetailedDescription** | Pointer to **string** | detailed description of log the message | [optional] 
**Level** | **string** | level of the log message. | 
**Source** | Pointer to [**Source**](Source.md) |  | [optional] 
**Subcategory** | **string** | subcategory of the log message. | 
**Timestamp** | **string** | timestamp of the log message. ISO 8601 date time. | 
**Type** | **string** | specific type of the log message. | 

## Methods

### NewAuditLogMessage

`func NewAuditLogMessage(category string, created string, description string, level string, subcategory string, timestamp string, type_ string, ) *AuditLogMessage`

NewAuditLogMessage instantiates a new AuditLogMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuditLogMessageWithDefaults

`func NewAuditLogMessageWithDefaults() *AuditLogMessage`

NewAuditLogMessageWithDefaults instantiates a new AuditLogMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCategory

`func (o *AuditLogMessage) GetCategory() string`

GetCategory returns the Category field if non-nil, zero value otherwise.

### GetCategoryOk

`func (o *AuditLogMessage) GetCategoryOk() (*string, bool)`

GetCategoryOk returns a tuple with the Category field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCategory

`func (o *AuditLogMessage) SetCategory(v string)`

SetCategory sets Category field to given value.


### GetContent

`func (o *AuditLogMessage) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AuditLogMessage) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AuditLogMessage) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *AuditLogMessage) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreated

`func (o *AuditLogMessage) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *AuditLogMessage) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *AuditLogMessage) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetDescription

`func (o *AuditLogMessage) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *AuditLogMessage) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *AuditLogMessage) SetDescription(v string)`

SetDescription sets Description field to given value.


### GetDetailedDescription

`func (o *AuditLogMessage) GetDetailedDescription() string`

GetDetailedDescription returns the DetailedDescription field if non-nil, zero value otherwise.

### GetDetailedDescriptionOk

`func (o *AuditLogMessage) GetDetailedDescriptionOk() (*string, bool)`

GetDetailedDescriptionOk returns a tuple with the DetailedDescription field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetailedDescription

`func (o *AuditLogMessage) SetDetailedDescription(v string)`

SetDetailedDescription sets DetailedDescription field to given value.

### HasDetailedDescription

`func (o *AuditLogMessage) HasDetailedDescription() bool`

HasDetailedDescription returns a boolean if a field has been set.

### GetLevel

`func (o *AuditLogMessage) GetLevel() string`

GetLevel returns the Level field if non-nil, zero value otherwise.

### GetLevelOk

`func (o *AuditLogMessage) GetLevelOk() (*string, bool)`

GetLevelOk returns a tuple with the Level field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLevel

`func (o *AuditLogMessage) SetLevel(v string)`

SetLevel sets Level field to given value.


### GetSource

`func (o *AuditLogMessage) GetSource() Source`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *AuditLogMessage) GetSourceOk() (*Source, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *AuditLogMessage) SetSource(v Source)`

SetSource sets Source field to given value.

### HasSource

`func (o *AuditLogMessage) HasSource() bool`

HasSource returns a boolean if a field has been set.

### GetSubcategory

`func (o *AuditLogMessage) GetSubcategory() string`

GetSubcategory returns the Subcategory field if non-nil, zero value otherwise.

### GetSubcategoryOk

`func (o *AuditLogMessage) GetSubcategoryOk() (*string, bool)`

GetSubcategoryOk returns a tuple with the Subcategory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubcategory

`func (o *AuditLogMessage) SetSubcategory(v string)`

SetSubcategory sets Subcategory field to given value.


### GetTimestamp

`func (o *AuditLogMessage) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *AuditLogMessage) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *AuditLogMessage) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.


### GetType

`func (o *AuditLogMessage) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *AuditLogMessage) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *AuditLogMessage) SetType(v string)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


