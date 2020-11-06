# SMSAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ContentTemplate** | **string** | A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the sms content | 
**DestinationPhoneNumbers** | **[]string** | A list of valid msisdn | 

## Methods

### NewSMSAction

`func NewSMSAction(contentTemplate string, destinationPhoneNumbers []string, ) *SMSAction`

NewSMSAction instantiates a new SMSAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSMSActionWithDefaults

`func NewSMSActionWithDefaults() *SMSAction`

NewSMSActionWithDefaults instantiates a new SMSAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContentTemplate

`func (o *SMSAction) GetContentTemplate() string`

GetContentTemplate returns the ContentTemplate field if non-nil, zero value otherwise.

### GetContentTemplateOk

`func (o *SMSAction) GetContentTemplateOk() (*string, bool)`

GetContentTemplateOk returns a tuple with the ContentTemplate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContentTemplate

`func (o *SMSAction) SetContentTemplate(v string)`

SetContentTemplate sets ContentTemplate field to given value.


### GetDestinationPhoneNumbers

`func (o *SMSAction) GetDestinationPhoneNumbers() []string`

GetDestinationPhoneNumbers returns the DestinationPhoneNumbers field if non-nil, zero value otherwise.

### GetDestinationPhoneNumbersOk

`func (o *SMSAction) GetDestinationPhoneNumbersOk() (*[]string, bool)`

GetDestinationPhoneNumbersOk returns a tuple with the DestinationPhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationPhoneNumbers

`func (o *SMSAction) SetDestinationPhoneNumbers(v []string)`

SetDestinationPhoneNumbers sets DestinationPhoneNumbers field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


