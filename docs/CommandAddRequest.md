# CommandAddRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Request** | [**CommandRequest**](CommandRequest.md) |  | 
**Policy** | Pointer to [**CommandPolicy**](CommandPolicy.md) |  | [optional] 

## Methods

### NewCommandAddRequest

`func NewCommandAddRequest(request CommandRequest, ) *CommandAddRequest`

NewCommandAddRequest instantiates a new CommandAddRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandAddRequestWithDefaults

`func NewCommandAddRequestWithDefaults() *CommandAddRequest`

NewCommandAddRequestWithDefaults instantiates a new CommandAddRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRequest

`func (o *CommandAddRequest) GetRequest() CommandRequest`

GetRequest returns the Request field if non-nil, zero value otherwise.

### GetRequestOk

`func (o *CommandAddRequest) GetRequestOk() (*CommandRequest, bool)`

GetRequestOk returns a tuple with the Request field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequest

`func (o *CommandAddRequest) SetRequest(v CommandRequest)`

SetRequest sets Request field to given value.


### GetPolicy

`func (o *CommandAddRequest) GetPolicy() CommandPolicy`

GetPolicy returns the Policy field if non-nil, zero value otherwise.

### GetPolicyOk

`func (o *CommandAddRequest) GetPolicyOk() (*CommandPolicy, bool)`

GetPolicyOk returns a tuple with the Policy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPolicy

`func (o *CommandAddRequest) SetPolicy(v CommandPolicy)`

SetPolicy sets Policy field to given value.

### HasPolicy

`func (o *CommandAddRequest) HasPolicy() bool`

HasPolicy returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


