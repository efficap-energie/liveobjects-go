# Command

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | command unique identifier | [optional] 
**TargetDeviceId** | Pointer to **string** | targeted device identifier (URN) | [optional] 
**Request** | Pointer to [**CommandRequest**](CommandRequest.md) |  | [optional] 
**Response** | Pointer to [**CommandResponse**](CommandResponse.md) |  | [optional] 
**Status** | Pointer to **string** | command current status | [optional] 
**DeliveryStatus** | Pointer to **string** | command current delivery status | [optional] 
**ErrorCode** | Pointer to **string** | error code in case of ERROR status | [optional] 
**Policy** | Pointer to [**CommandPolicy**](CommandPolicy.md) |  | [optional] 
**History** | Pointer to [**[]CommandHistory**](CommandHistory.md) | command history | [optional] 
**NotifyTo** | Pointer to **string** | topic where command status change notification are published | [optional] 
**Created** | Pointer to **string** | command creation date/time | [optional] 
**Updated** | Pointer to **string** | command last status update date/time | [optional] 

## Methods

### NewCommand

`func NewCommand() *Command`

NewCommand instantiates a new Command object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandWithDefaults

`func NewCommandWithDefaults() *Command`

NewCommandWithDefaults instantiates a new Command object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Command) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Command) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Command) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Command) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTargetDeviceId

`func (o *Command) GetTargetDeviceId() string`

GetTargetDeviceId returns the TargetDeviceId field if non-nil, zero value otherwise.

### GetTargetDeviceIdOk

`func (o *Command) GetTargetDeviceIdOk() (*string, bool)`

GetTargetDeviceIdOk returns a tuple with the TargetDeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetDeviceId

`func (o *Command) SetTargetDeviceId(v string)`

SetTargetDeviceId sets TargetDeviceId field to given value.

### HasTargetDeviceId

`func (o *Command) HasTargetDeviceId() bool`

HasTargetDeviceId returns a boolean if a field has been set.

### GetRequest

`func (o *Command) GetRequest() CommandRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *Command) GetRequestOk() (*CommandRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *Command) SetRequest(v CommandRequest)`

SetRequest sets Request field to given value.

### HasRequest

`func (o *Command) HasRequest() bool`

HasRequest returns a boolean if a field has been set.

### GetResponse

`func (o *Command) GetResponse() CommandResponse`

GetResponse returns the Response field if non-nil, zero value otherwise.

### GetResponseOk

`func (o *Command) GetResponseOk() (*CommandResponse, bool)`

GetResponseOk returns a tuple with the Response field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResponse

`func (o *Command) SetResponse(v CommandResponse)`

SetResponse sets Response field to given value.

### HasResponse

`func (o *Command) HasResponse() bool`

HasResponse returns a boolean if a field has been set.

### GetStatus

`func (o *Command) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Command) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Command) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Command) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeliveryStatus

`func (o *Command) GetDeliveryStatus() string`

GetDeliveryStatus returns the DeliveryStatus field if non-nil, zero value otherwise.

### GetDeliveryStatusOk

`func (o *Command) GetDeliveryStatusOk() (*string, bool)`

GetDeliveryStatusOk returns a tuple with the DeliveryStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeliveryStatus

`func (o *Command) SetDeliveryStatus(v string)`

SetDeliveryStatus sets DeliveryStatus field to given value.

### HasDeliveryStatus

`func (o *Command) HasDeliveryStatus() bool`

HasDeliveryStatus returns a boolean if a field has been set.

### GetErrorCode

`func (o *Command) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *Command) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *Command) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *Command) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetPolicy

`func (o *Command) GetPolicy() CommandPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *Command) GetPolicyOk() (*CommandPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *Command) SetPolicy(v CommandPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *Command) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.

### GetHistory

`func (o *Command) GetHistory() []CommandHistory`

GetHistory returns the History field if non-nil, zero value otherwise.

### GetHistoryOk

`func (o *Command) GetHistoryOk() (*[]CommandHistory, bool)`

GetHistoryOk returns a tuple with the History field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHistory

`func (o *Command) SetHistory(v []CommandHistory)`

SetHistory sets History field to given value.

### HasHistory

`func (o *Command) HasHistory() bool`

HasHistory returns a boolean if a field has been set.

### GetNotifyTo

`func (o *Command) GetNotifyTo() string`

GetNotifyTo returns the NotifyTo field if non-nil, zero value otherwise.

### GetNotifyToOk

`func (o *Command) GetNotifyToOk() (*string, bool)`

GetNotifyToOk returns a tuple with the NotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTo

`func (o *Command) SetNotifyTo(v string)`

SetNotifyTo sets NotifyTo field to given value.

### HasNotifyTo

`func (o *Command) HasNotifyTo() bool`

HasNotifyTo returns a boolean if a field has been set.

### GetCreated

`func (o *Command) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Command) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Command) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *Command) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetUpdated

`func (o *Command) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Command) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Command) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Command) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


