# AzureEventHubsAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the webhook body. If empty, the raw event will be used. | [optional] 
**EventHubName** | **string** | The name of the targeted Event Hub. | 
**EventHubsNamespace** | **string** | the Event Hubs namespace where is located the targeted Event Hub. | 
**JsonPath** | Pointer to **string** | The json path to extract from the considered message (or event), it will be taken as the root datacontext object when combined with a mustache template in content | [optional] 
**RetryOnFailure** | Pointer to **bool** | Indicate if a retry policy should be set up in case of a Event Hub delivery failure | [optional] 
**SharedAccessKey** | **string** | The shared access key (also called shared access policy). | 
**SharedAccessKeyName** | **string** | The name of the shared access key (also called shared access policy) from the targeted Event Hub. | 

## Methods

### NewAzureEventHubsAction

`func NewAzureEventHubsAction(eventHubName string, eventHubsNamespace string, sharedAccessKey string, sharedAccessKeyName string, ) *AzureEventHubsAction`

NewAzureEventHubsAction instantiates a new AzureEventHubsAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAzureEventHubsActionWithDefaults

`func NewAzureEventHubsActionWithDefaults() *AzureEventHubsAction`

NewAzureEventHubsActionWithDefaults instantiates a new AzureEventHubsAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *AzureEventHubsAction) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *AzureEventHubsAction) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *AzureEventHubsAction) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *AzureEventHubsAction) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetEventHubName

`func (o *AzureEventHubsAction) GetEventHubName() string`

GetEventHubName returns the EventHubName field if non-nil, zero value otherwise.

### GetEventHubNameOk

`func (o *AzureEventHubsAction) GetEventHubNameOk() (*string, bool)`

GetEventHubNameOk returns a tuple with the EventHubName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubName

`func (o *AzureEventHubsAction) SetEventHubName(v string)`

SetEventHubName sets EventHubName field to given value.


### GetEventHubsNamespace

`func (o *AzureEventHubsAction) GetEventHubsNamespace() string`

GetEventHubsNamespace returns the EventHubsNamespace field if non-nil, zero value otherwise.

### GetEventHubsNamespaceOk

`func (o *AzureEventHubsAction) GetEventHubsNamespaceOk() (*string, bool)`

GetEventHubsNamespaceOk returns a tuple with the EventHubsNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEventHubsNamespace

`func (o *AzureEventHubsAction) SetEventHubsNamespace(v string)`

SetEventHubsNamespace sets EventHubsNamespace field to given value.


### GetJsonPath

`func (o *AzureEventHubsAction) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *AzureEventHubsAction) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *AzureEventHubsAction) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.

### HasJsonPath

`func (o *AzureEventHubsAction) HasJsonPath() bool`

HasJsonPath returns a boolean if a field has been set.

### GetRetryOnFailure

`func (o *AzureEventHubsAction) GetRetryOnFailure() bool`

GetRetryOnFailure returns the RetryOnFailure field if non-nil, zero value otherwise.

### GetRetryOnFailureOk

`func (o *AzureEventHubsAction) GetRetryOnFailureOk() (*bool, bool)`

GetRetryOnFailureOk returns a tuple with the RetryOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOnFailure

`func (o *AzureEventHubsAction) SetRetryOnFailure(v bool)`

SetRetryOnFailure sets RetryOnFailure field to given value.

### HasRetryOnFailure

`func (o *AzureEventHubsAction) HasRetryOnFailure() bool`

HasRetryOnFailure returns a boolean if a field has been set.

### GetSharedAccessKey

`func (o *AzureEventHubsAction) GetSharedAccessKey() string`

GetSharedAccessKey returns the SharedAccessKey field if non-nil, zero value otherwise.

### GetSharedAccessKeyOk

`func (o *AzureEventHubsAction) GetSharedAccessKeyOk() (*string, bool)`

GetSharedAccessKeyOk returns a tuple with the SharedAccessKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAccessKey

`func (o *AzureEventHubsAction) SetSharedAccessKey(v string)`

SetSharedAccessKey sets SharedAccessKey field to given value.


### GetSharedAccessKeyName

`func (o *AzureEventHubsAction) GetSharedAccessKeyName() string`

GetSharedAccessKeyName returns the SharedAccessKeyName field if non-nil, zero value otherwise.

### GetSharedAccessKeyNameOk

`func (o *AzureEventHubsAction) GetSharedAccessKeyNameOk() (*string, bool)`

GetSharedAccessKeyNameOk returns a tuple with the SharedAccessKeyName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSharedAccessKeyName

`func (o *AzureEventHubsAction) SetSharedAccessKeyName(v string)`

SetSharedAccessKeyName sets SharedAccessKeyName field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


