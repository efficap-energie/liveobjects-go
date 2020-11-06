# HttpPushAction

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **string** | A raw string or a mustache (https://mustache.github.io/mustache.5.html) template describing the webhook body. If empty, the raw event will be used. | [optional] 
**Headers** | Pointer to [**map[string][]string**](array.md) | A map of header to be passed in the http POST headers | [optional] 
**JsonPath** | Pointer to **string** | The json path to extract from the considered message (or event), it will be taken as the root datacontext object when combined with a mustache template in content | [optional] 
**RetryOnFailure** | Pointer to **bool** | Indicate if a retry policy should be set up in case of a http push delivery failure | [optional] 
**WebhookUrl** | **string** | The url of the target service webhook (only the ports 80, 443, 8080, 8443 and 9243 are allowed). | 

## Methods

### NewHttpPushAction

`func NewHttpPushAction(webhookUrl string, ) *HttpPushAction`

NewHttpPushAction instantiates a new HttpPushAction object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpPushActionWithDefaults

`func NewHttpPushActionWithDefaults() *HttpPushAction`

NewHttpPushActionWithDefaults instantiates a new HttpPushAction object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *HttpPushAction) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *HttpPushAction) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *HttpPushAction) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *HttpPushAction) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetHeaders

`func (o *HttpPushAction) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpPushAction) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpPushAction) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpPushAction) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetJsonPath

`func (o *HttpPushAction) GetJsonPath() string`

GetJsonPath returns the JsonPath field if non-nil, zero value otherwise.

### GetJsonPathOk

`func (o *HttpPushAction) GetJsonPathOk() (*string, bool)`

GetJsonPathOk returns a tuple with the JsonPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJsonPath

`func (o *HttpPushAction) SetJsonPath(v string)`

SetJsonPath sets JsonPath field to given value.

### HasJsonPath

`func (o *HttpPushAction) HasJsonPath() bool`

HasJsonPath returns a boolean if a field has been set.

### GetRetryOnFailure

`func (o *HttpPushAction) GetRetryOnFailure() bool`

GetRetryOnFailure returns the RetryOnFailure field if non-nil, zero value otherwise.

### GetRetryOnFailureOk

`func (o *HttpPushAction) GetRetryOnFailureOk() (*bool, bool)`

GetRetryOnFailureOk returns a tuple with the RetryOnFailure field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRetryOnFailure

`func (o *HttpPushAction) SetRetryOnFailure(v bool)`

SetRetryOnFailure sets RetryOnFailure field to given value.

### HasRetryOnFailure

`func (o *HttpPushAction) HasRetryOnFailure() bool`

HasRetryOnFailure returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *HttpPushAction) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *HttpPushAction) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *HttpPushAction) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


