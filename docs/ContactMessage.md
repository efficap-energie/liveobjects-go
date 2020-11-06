# ContactMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cc** | Pointer to **[]string** | cc | [optional] 
**Cci** | Pointer to **[]string** | cci | [optional] 
**PhoneNumbers** | Pointer to **[]string** | phoneNumbers | [optional] 
**To** | Pointer to **[]string** | to | [optional] 

## Methods

### NewContactMessage

`func NewContactMessage() *ContactMessage`

NewContactMessage instantiates a new ContactMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContactMessageWithDefaults

`func NewContactMessageWithDefaults() *ContactMessage`

NewContactMessageWithDefaults instantiates a new ContactMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCc

`func (o *ContactMessage) GetCc() []string`

GetCc returns the Cc field if non-nil, zero value otherwise.

### GetCcOk

`func (o *ContactMessage) GetCcOk() (*[]string, bool)`

GetCcOk returns a tuple with the Cc field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCc

`func (o *ContactMessage) SetCc(v []string)`

SetCc sets Cc field to given value.

### HasCc

`func (o *ContactMessage) HasCc() bool`

HasCc returns a boolean if a field has been set.

### GetCci

`func (o *ContactMessage) GetCci() []string`

GetCci returns the Cci field if non-nil, zero value otherwise.

### GetCciOk

`func (o *ContactMessage) GetCciOk() (*[]string, bool)`

GetCciOk returns a tuple with the Cci field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCci

`func (o *ContactMessage) SetCci(v []string)`

SetCci sets Cci field to given value.

### HasCci

`func (o *ContactMessage) HasCci() bool`

HasCci returns a boolean if a field has been set.

### GetPhoneNumbers

`func (o *ContactMessage) GetPhoneNumbers() []string`

GetPhoneNumbers returns the PhoneNumbers field if non-nil, zero value otherwise.

### GetPhoneNumbersOk

`func (o *ContactMessage) GetPhoneNumbersOk() (*[]string, bool)`

GetPhoneNumbersOk returns a tuple with the PhoneNumbers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPhoneNumbers

`func (o *ContactMessage) SetPhoneNumbers(v []string)`

SetPhoneNumbers sets PhoneNumbers field to given value.

### HasPhoneNumbers

`func (o *ContactMessage) HasPhoneNumbers() bool`

HasPhoneNumbers returns a boolean if a field has been set.

### GetTo

`func (o *ContactMessage) GetTo() []string`

GetTo returns the To field if non-nil, zero value otherwise.

### GetToOk

`func (o *ContactMessage) GetToOk() (*[]string, bool)`

GetToOk returns a tuple with the To field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTo

`func (o *ContactMessage) SetTo(v []string)`

SetTo sets To field to given value.

### HasTo

`func (o *ContactMessage) HasTo() bool`

HasTo returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


