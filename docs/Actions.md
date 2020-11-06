# Actions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AzureEventHubs** | Pointer to [**[]AzureEventHubsAction**](AzureEventHubsAction.md) | A collection of Azure Event Hubs actions | [optional] 
**Emails** | Pointer to [**[]EmailAction**](EmailAction.md) | A collection of email actions | [optional] 
**FifoPublish** | Pointer to [**[]FifoPublishAction**](FifoPublishAction.md) | A collection of fifo publish actions | [optional] 
**HttpPush** | Pointer to [**[]HttpPushAction**](HttpPushAction.md) | A collection of webhook actions | [optional] 
**Sms** | Pointer to [**[]SMSAction**](SMSAction.md) | A collection of sms actions | [optional] 

## Methods

### NewActions

`func NewActions() *Actions`

NewActions instantiates a new Actions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewActionsWithDefaults

`func NewActionsWithDefaults() *Actions`

NewActionsWithDefaults instantiates a new Actions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAzureEventHubs

`func (o *Actions) GetAzureEventHubs() []AzureEventHubsAction`

GetAzureEventHubs returns the AzureEventHubs field if non-nil, zero value otherwise.

### GetAzureEventHubsOk

`func (o *Actions) GetAzureEventHubsOk() (*[]AzureEventHubsAction, bool)`

GetAzureEventHubsOk returns a tuple with the AzureEventHubs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAzureEventHubs

`func (o *Actions) SetAzureEventHubs(v []AzureEventHubsAction)`

SetAzureEventHubs sets AzureEventHubs field to given value.

### HasAzureEventHubs

`func (o *Actions) HasAzureEventHubs() bool`

HasAzureEventHubs returns a boolean if a field has been set.

### GetEmails

`func (o *Actions) GetEmails() []EmailAction`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *Actions) GetEmailsOk() (*[]EmailAction, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *Actions) SetEmails(v []EmailAction)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *Actions) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetFifoPublish

`func (o *Actions) GetFifoPublish() []FifoPublishAction`

GetFifoPublish returns the FifoPublish field if non-nil, zero value otherwise.

### GetFifoPublishOk

`func (o *Actions) GetFifoPublishOk() (*[]FifoPublishAction, bool)`

GetFifoPublishOk returns a tuple with the FifoPublish field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifoPublish

`func (o *Actions) SetFifoPublish(v []FifoPublishAction)`

SetFifoPublish sets FifoPublish field to given value.

### HasFifoPublish

`func (o *Actions) HasFifoPublish() bool`

HasFifoPublish returns a boolean if a field has been set.

### GetHttpPush

`func (o *Actions) GetHttpPush() []HttpPushAction`

GetHttpPush returns the HttpPush field if non-nil, zero value otherwise.

### GetHttpPushOk

`func (o *Actions) GetHttpPushOk() (*[]HttpPushAction, bool)`

GetHttpPushOk returns a tuple with the HttpPush field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpPush

`func (o *Actions) SetHttpPush(v []HttpPushAction)`

SetHttpPush sets HttpPush field to given value.

### HasHttpPush

`func (o *Actions) HasHttpPush() bool`

HasHttpPush returns a boolean if a field has been set.

### GetSms

`func (o *Actions) GetSms() []SMSAction`

GetSms returns the Sms field if non-nil, zero value otherwise.

### GetSmsOk

`func (o *Actions) GetSmsOk() (*[]SMSAction, bool)`

GetSmsOk returns a tuple with the Sms field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSms

`func (o *Actions) SetSms(v []SMSAction)`

SetSms sets Sms field to given value.

### HasSms

`func (o *Actions) HasSms() bool`

HasSms returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


