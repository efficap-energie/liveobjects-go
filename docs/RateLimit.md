# RateLimit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HttpMaxCalls** | Pointer to **int64** | http rate limit: maximum api calls allowed per time window | [optional] 
**HttpWindowSize** | Pointer to **int64** | http rate limit: window size in seconds | [optional] 
**MqttBridgeMaxMessages** | Pointer to **int64** | mqtt bridge rate limit: maximum messages allowed per time window | [optional] 
**MqttBridgeWindowSize** | Pointer to **int64** | mqtt bridge rate limit: window size in seconds | [optional] 
**MqttDeviceMaxMessages** | Pointer to **int64** | mqtt device rate limit: maximum messages allowed per time window | [optional] 
**MqttDeviceWindowSize** | Pointer to **int64** | mqtt device rate limit: window size in seconds | [optional] 

## Methods

### NewRateLimit

`func NewRateLimit() *RateLimit`

NewRateLimit instantiates a new RateLimit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRateLimitWithDefaults

`func NewRateLimitWithDefaults() *RateLimit`

NewRateLimitWithDefaults instantiates a new RateLimit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHttpMaxCalls

`func (o *RateLimit) GetHttpMaxCalls() int64`

GetHttpMaxCalls returns the HttpMaxCalls field if non-nil, zero value otherwise.

### GetHttpMaxCallsOk

`func (o *RateLimit) GetHttpMaxCallsOk() (*int64, bool)`

GetHttpMaxCallsOk returns a tuple with the HttpMaxCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpMaxCalls

`func (o *RateLimit) SetHttpMaxCalls(v int64)`

SetHttpMaxCalls sets HttpMaxCalls field to given value.

### HasHttpMaxCalls

`func (o *RateLimit) HasHttpMaxCalls() bool`

HasHttpMaxCalls returns a boolean if a field has been set.

### GetHttpWindowSize

`func (o *RateLimit) GetHttpWindowSize() int64`

GetHttpWindowSize returns the HttpWindowSize field if non-nil, zero value otherwise.

### GetHttpWindowSizeOk

`func (o *RateLimit) GetHttpWindowSizeOk() (*int64, bool)`

GetHttpWindowSizeOk returns a tuple with the HttpWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHttpWindowSize

`func (o *RateLimit) SetHttpWindowSize(v int64)`

SetHttpWindowSize sets HttpWindowSize field to given value.

### HasHttpWindowSize

`func (o *RateLimit) HasHttpWindowSize() bool`

HasHttpWindowSize returns a boolean if a field has been set.

### GetMqttBridgeMaxMessages

`func (o *RateLimit) GetMqttBridgeMaxMessages() int64`

GetMqttBridgeMaxMessages returns the MqttBridgeMaxMessages field if non-nil, zero value otherwise.

### GetMqttBridgeMaxMessagesOk

`func (o *RateLimit) GetMqttBridgeMaxMessagesOk() (*int64, bool)`

GetMqttBridgeMaxMessagesOk returns a tuple with the MqttBridgeMaxMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttBridgeMaxMessages

`func (o *RateLimit) SetMqttBridgeMaxMessages(v int64)`

SetMqttBridgeMaxMessages sets MqttBridgeMaxMessages field to given value.

### HasMqttBridgeMaxMessages

`func (o *RateLimit) HasMqttBridgeMaxMessages() bool`

HasMqttBridgeMaxMessages returns a boolean if a field has been set.

### GetMqttBridgeWindowSize

`func (o *RateLimit) GetMqttBridgeWindowSize() int64`

GetMqttBridgeWindowSize returns the MqttBridgeWindowSize field if non-nil, zero value otherwise.

### GetMqttBridgeWindowSizeOk

`func (o *RateLimit) GetMqttBridgeWindowSizeOk() (*int64, bool)`

GetMqttBridgeWindowSizeOk returns a tuple with the MqttBridgeWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttBridgeWindowSize

`func (o *RateLimit) SetMqttBridgeWindowSize(v int64)`

SetMqttBridgeWindowSize sets MqttBridgeWindowSize field to given value.

### HasMqttBridgeWindowSize

`func (o *RateLimit) HasMqttBridgeWindowSize() bool`

HasMqttBridgeWindowSize returns a boolean if a field has been set.

### GetMqttDeviceMaxMessages

`func (o *RateLimit) GetMqttDeviceMaxMessages() int64`

GetMqttDeviceMaxMessages returns the MqttDeviceMaxMessages field if non-nil, zero value otherwise.

### GetMqttDeviceMaxMessagesOk

`func (o *RateLimit) GetMqttDeviceMaxMessagesOk() (*int64, bool)`

GetMqttDeviceMaxMessagesOk returns a tuple with the MqttDeviceMaxMessages field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttDeviceMaxMessages

`func (o *RateLimit) SetMqttDeviceMaxMessages(v int64)`

SetMqttDeviceMaxMessages sets MqttDeviceMaxMessages field to given value.

### HasMqttDeviceMaxMessages

`func (o *RateLimit) HasMqttDeviceMaxMessages() bool`

HasMqttDeviceMaxMessages returns a boolean if a field has been set.

### GetMqttDeviceWindowSize

`func (o *RateLimit) GetMqttDeviceWindowSize() int64`

GetMqttDeviceWindowSize returns the MqttDeviceWindowSize field if non-nil, zero value otherwise.

### GetMqttDeviceWindowSizeOk

`func (o *RateLimit) GetMqttDeviceWindowSizeOk() (*int64, bool)`

GetMqttDeviceWindowSizeOk returns a tuple with the MqttDeviceWindowSize field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMqttDeviceWindowSize

`func (o *RateLimit) SetMqttDeviceWindowSize(v int64)`

SetMqttDeviceWindowSize sets MqttDeviceWindowSize field to given value.

### HasMqttDeviceWindowSize

`func (o *RateLimit) HasMqttDeviceWindowSize() bool`

HasMqttDeviceWindowSize returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


