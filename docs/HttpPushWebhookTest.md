# HttpPushWebhookTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Headers** | Pointer to [**map[string][]string**](array.md) | A map of header to be passed in the http POST headers | [optional] 
**RequestBody** | Pointer to **string** | The body of the http POST test request | [optional] 
**RequestTimeout** | Pointer to **int64** | The maximum delay allowed to complete the http POST test request | [optional] 
**WebhookUrl** | Pointer to **string** | The url of the target service webhook (only the ports 80, 443, 8080, 8443 and 9243 are allowed). | [optional] 

## Methods

### NewHttpPushWebhookTest

`func NewHttpPushWebhookTest() *HttpPushWebhookTest`

NewHttpPushWebhookTest instantiates a new HttpPushWebhookTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpPushWebhookTestWithDefaults

`func NewHttpPushWebhookTestWithDefaults() *HttpPushWebhookTest`

NewHttpPushWebhookTestWithDefaults instantiates a new HttpPushWebhookTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHeaders

`func (o *HttpPushWebhookTest) GetHeaders() map[string][]string`

GetHeaders returns the Headers field if non-nil, zero value otherwise.

### GetHeadersOk

`func (o *HttpPushWebhookTest) GetHeadersOk() (*map[string][]string, bool)`

GetHeadersOk returns a tuple with the Headers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHeaders

`func (o *HttpPushWebhookTest) SetHeaders(v map[string][]string)`

SetHeaders sets Headers field to given value.

### HasHeaders

`func (o *HttpPushWebhookTest) HasHeaders() bool`

HasHeaders returns a boolean if a field has been set.

### GetRequestBody

`func (o *HttpPushWebhookTest) GetRequestBody() string`

GetRequestBody returns the RequestBody field if non-nil, zero value otherwise.

### GetRequestBodyOk

`func (o *HttpPushWebhookTest) GetRequestBodyOk() (*string, bool)`

GetRequestBodyOk returns a tuple with the RequestBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestBody

`func (o *HttpPushWebhookTest) SetRequestBody(v string)`

SetRequestBody sets RequestBody field to given value.

### HasRequestBody

`func (o *HttpPushWebhookTest) HasRequestBody() bool`

HasRequestBody returns a boolean if a field has been set.

### GetRequestTimeout

`func (o *HttpPushWebhookTest) GetRequestTimeout() int64`

GetRequestTimeout returns the RequestTimeout field if non-nil, zero value otherwise.

### GetRequestTimeoutOk

`func (o *HttpPushWebhookTest) GetRequestTimeoutOk() (*int64, bool)`

GetRequestTimeoutOk returns a tuple with the RequestTimeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestTimeout

`func (o *HttpPushWebhookTest) SetRequestTimeout(v int64)`

SetRequestTimeout sets RequestTimeout field to given value.

### HasRequestTimeout

`func (o *HttpPushWebhookTest) HasRequestTimeout() bool`

HasRequestTimeout returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *HttpPushWebhookTest) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *HttpPushWebhookTest) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *HttpPushWebhookTest) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *HttpPushWebhookTest) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


