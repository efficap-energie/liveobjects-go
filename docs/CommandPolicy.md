# CommandPolicy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ExpirationInSeconds** | Pointer to **int64** | Expiration in seconds since command creation date. Default is no expiry. Min value is 5 seconds | [optional] 
**AckTimeoutInSeconds** | Pointer to **int64** | Ack timeout in seconds since command was sent. Default is no ack timeout. Min value is 10 seconds | [optional] 
**AckMode** | Pointer to **string** | Ack mode for this command. NONE (or AUTO) ack means that the command is automatically acknowledged (set to &#39;PROCESSED&#39; status) as the command is sent to the device. NETWORK ack requires a reception acknowledge. APPLICATIVE (or DEVICE) ack requires a command response from the device to change its status. Default ack mode is connectivity dependant. | [optional] 
**Attempts** | Pointer to **int32** | Number of attemps in case of ERROR. Default to 1 | [optional] 

## Methods

### NewCommandPolicy

`func NewCommandPolicy() *CommandPolicy`

NewCommandPolicy instantiates a new CommandPolicy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandPolicyWithDefaults

`func NewCommandPolicyWithDefaults() *CommandPolicy`

NewCommandPolicyWithDefaults instantiates a new CommandPolicy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExpirationInSeconds

`func (o *CommandPolicy) GetExpirationInSeconds() int64`

GetExpirationInSeconds returns the ExpirationInSeconds field if non-nil, zero value otherwise.

### GetExpirationInSecondsOk

`func (o *CommandPolicy) GetExpirationInSecondsOk() (*int64, bool)`

GetExpirationInSecondsOk returns a tuple with the ExpirationInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExpirationInSeconds

`func (o *CommandPolicy) SetExpirationInSeconds(v int64)`

SetExpirationInSeconds sets ExpirationInSeconds field to given value.

### HasExpirationInSeconds

`func (o *CommandPolicy) HasExpirationInSeconds() bool`

HasExpirationInSeconds returns a boolean if a field has been set.

### GetAckTimeoutInSeconds

`func (o *CommandPolicy) GetAckTimeoutInSeconds() int64`

GetAckTimeoutInSeconds returns the AckTimeoutInSeconds field if non-nil, zero value otherwise.

### GetAckTimeoutInSecondsOk

`func (o *CommandPolicy) GetAckTimeoutInSecondsOk() (*int64, bool)`

GetAckTimeoutInSecondsOk returns a tuple with the AckTimeoutInSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckTimeoutInSeconds

`func (o *CommandPolicy) SetAckTimeoutInSeconds(v int64)`

SetAckTimeoutInSeconds sets AckTimeoutInSeconds field to given value.

### HasAckTimeoutInSeconds

`func (o *CommandPolicy) HasAckTimeoutInSeconds() bool`

HasAckTimeoutInSeconds returns a boolean if a field has been set.

### GetAckMode

`func (o *CommandPolicy) GetAckMode() string`

GetAckMode returns the AckMode field if non-nil, zero value otherwise.

### GetAckModeOk

`func (o *CommandPolicy) GetAckModeOk() (*string, bool)`

GetAckModeOk returns a tuple with the AckMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAckMode

`func (o *CommandPolicy) SetAckMode(v string)`

SetAckMode sets AckMode field to given value.

### HasAckMode

`func (o *CommandPolicy) HasAckMode() bool`

HasAckMode returns a boolean if a field has been set.

### GetAttempts

`func (o *CommandPolicy) GetAttempts() int32`

GetAttempts returns the Attempts field if non-nil, zero value otherwise.

### GetAttemptsOk

`func (o *CommandPolicy) GetAttemptsOk() (*int32, bool)`

GetAttemptsOk returns a tuple with the Attempts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttempts

`func (o *CommandPolicy) SetAttempts(v int32)`

SetAttempts sets Attempts field to given value.

### HasAttempts

`func (o *CommandPolicy) HasAttempts() bool`

HasAttempts returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


