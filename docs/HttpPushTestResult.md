# HttpPushTestResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ErrorMessage** | Pointer to **string** |  | [optional] 
**HttpResponseBody** | Pointer to **string** |  | [optional] 
**HttpResponseStatusCode** | Pointer to **int32** |  | [optional] 
**Success** | Pointer to **bool** |  | [optional] 
**UrlBlacklisted** | Pointer to **bool** |  | [optional] 

## Methods

### NewHttpPushTestResult

`func NewHttpPushTestResult() *HttpPushTestResult`

NewHttpPushTestResult instantiates a new HttpPushTestResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHttpPushTestResultWithDefaults

`func NewHttpPushTestResultWithDefaults() *HttpPushTestResult`

NewHttpPushTestResultWithDefaults instantiates a new HttpPushTestResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrorMessage

`func (o *HttpPushTestResult) GetErrorMessage() string`

GetErrorMessage returns the ErrorMessage field if non-nil, zero value otherwise.

### GetErrorMessageOk

`func (o *HttpPushTestResult) GetErrorMessageOk() (*string, bool)`

GetErrorMessageOk returns a tuple with the ErrorMessage field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorMessage

`func (o *HttpPushTestResult) SetErrorMessage(v string)`

SetErrorMessage sets ErrorMessage field to given value.

### HasErrorMessage

`func (o *HttpPushTestResult) HasErrorMessage() bool`

HasErrorMessage returns a boolean if a field has been set.

### GetHttpResponseBody

`func (o *HttpPushTestResult) GetHttpResponseBody() string`

GetHttpResponseBody returns the HttpResponseBody field if non-nil, zero value otherwise.

### GetHttpResponseBodyOk

`func (o *HttpPushTestResult) GetHttpResponseBodyOk() (*string, bool)`

GetHttpResponseBodyOk returns a tuple with the HttpResponseBody field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseBody

`func (o *HttpPushTestResult) SetHttpResponseBody(v string)`

SetHttpResponseBody sets HttpResponseBody field to given value.

### HasHttpResponseBody

`func (o *HttpPushTestResult) HasHttpResponseBody() bool`

HasHttpResponseBody returns a boolean if a field has been set.

### GetHttpResponseStatusCode

`func (o *HttpPushTestResult) GetHttpResponseStatusCode() int32`

GetHttpResponseStatusCode returns the HttpResponseStatusCode field if non-nil, zero value otherwise.

### GetHttpResponseStatusCodeOk

`func (o *HttpPushTestResult) GetHttpResponseStatusCodeOk() (*int32, bool)`

GetHttpResponseStatusCodeOk returns a tuple with the HttpResponseStatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpResponseStatusCode

`func (o *HttpPushTestResult) SetHttpResponseStatusCode(v int32)`

SetHttpResponseStatusCode sets HttpResponseStatusCode field to given value.

### HasHttpResponseStatusCode

`func (o *HttpPushTestResult) HasHttpResponseStatusCode() bool`

HasHttpResponseStatusCode returns a boolean if a field has been set.

### GetSuccess

`func (o *HttpPushTestResult) GetSuccess() bool`

GetSuccess returns the Success field if non-nil, zero value otherwise.

### GetSuccessOk

`func (o *HttpPushTestResult) GetSuccessOk() (*bool, bool)`

GetSuccessOk returns a tuple with the Success field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSuccess

`func (o *HttpPushTestResult) SetSuccess(v bool)`

SetSuccess sets Success field to given value.

### HasSuccess

`func (o *HttpPushTestResult) HasSuccess() bool`

HasSuccess returns a boolean if a field has been set.

### GetUrlBlacklisted

`func (o *HttpPushTestResult) GetUrlBlacklisted() bool`

GetUrlBlacklisted returns the UrlBlacklisted field if non-nil, zero value otherwise.

### GetUrlBlacklistedOk

`func (o *HttpPushTestResult) GetUrlBlacklistedOk() (*bool, bool)`

GetUrlBlacklistedOk returns a tuple with the UrlBlacklisted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrlBlacklisted

`func (o *HttpPushTestResult) SetUrlBlacklisted(v bool)`

SetUrlBlacklisted sets UrlBlacklisted field to given value.

### HasUrlBlacklisted

`func (o *HttpPushTestResult) HasUrlBlacklisted() bool`

HasUrlBlacklisted returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


