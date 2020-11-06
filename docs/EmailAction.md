# EmailAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cc** | Pointer to **[]string** |  | [optional] 
**Cci** | Pointer to **[]string** |  | [optional] 
**ContentTemplate** | **string** | A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the email content | 
**SubjectTemplate** | **string** | A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the email subject | 
**To** | **[]string** |  | 

## Methods

### NewEmailAction

`func NewEmailAction(contentTemplate string, subjectTemplate string, to []string, ) *EmailAction`

NewEmailAction instantiates a new EmailAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEmailActionWithDefaults

`func NewEmailActionWithDefaults() *EmailAction`

NewEmailActionWithDefaults instantiates a new EmailAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCc

`func (o *EmailAction) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *EmailAction) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *EmailAction) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *EmailAction) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetCci

`func (o *EmailAction) GetCci() []string`

GetCci returns the Cci field if non-nil, zero value otherwise.

### GetCciOk

`func (o *EmailAction) GetCciOk() (*[]string, bool)`

GetCciOk returns a tuple with the Cci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCci

`func (o *EmailAction) SetCci(v []string)`

SetCci sets Cci field to given value.

### HasCci

`func (o *EmailAction) HasCci() bool`

HasCci returns a boolean if a field has been set.

### GetContentTemplate

`func (o *EmailAction) GetContentTemplate() string`

GetContentTemplate returns the ContentTemplate field if non-nil, zero value otherwise.

### GetContentTemplateOk

`func (o *EmailAction) GetContentTemplateOk() (*string, bool)`

GetContentTemplateOk returns a tuple with the ContentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTemplate

`func (o *EmailAction) SetContentTemplate(v string)`

SetContentTemplate sets ContentTemplate field to given value.


### GetSubjectTemplate

`func (o *EmailAction) GetSubjectTemplate() string`

GetSubjectTemplate returns the SubjectTemplate field if non-nil, zero value otherwise.

### GetSubjectTemplateOk

`func (o *EmailAction) GetSubjectTemplateOk() (*string, bool)`

GetSubjectTemplateOk returns a tuple with the SubjectTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSubjectTemplate

`func (o *EmailAction) SetSubjectTemplate(v string)`

SetSubjectTemplate sets SubjectTemplate field to given value.


### GetTo

`func (o *EmailAction) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *EmailAction) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *EmailAction) SetTo(v []string)`

SetTo sets To field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


