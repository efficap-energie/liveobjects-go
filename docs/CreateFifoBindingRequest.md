# CreateFifoBindingRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FifoId** | **string** | FIFO identifier. Fifo name must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_~]+&lt;/code&gt; (max 255 characters) | 
**RoutingKeyFilter** | **string** | routing Key filter. Routing key must respect the following regular expression &lt;code&gt;[A-Za-z\\d-_:~#\\.\\*]+&lt;/code&gt; (max 1000 characters) | 

## Methods

### NewCreateFifoBindingRequest

`func NewCreateFifoBindingRequest(fifoId string, routingKeyFilter string, ) *CreateFifoBindingRequest`

NewCreateFifoBindingRequest instantiates a new CreateFifoBindingRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateFifoBindingRequestWithDefaults

`func NewCreateFifoBindingRequestWithDefaults() *CreateFifoBindingRequest`

NewCreateFifoBindingRequestWithDefaults instantiates a new CreateFifoBindingRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFifoId

`func (o *CreateFifoBindingRequest) GetFifoId() string`

GetFifoId returns the FifoId field if non-nil, zero value otherwise.

### GetFifoIdOk

`func (o *CreateFifoBindingRequest) GetFifoIdOk() (*string, bool)`

GetFifoIdOk returns a tuple with the FifoId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifoId

`func (o *CreateFifoBindingRequest) SetFifoId(v string)`

SetFifoId sets FifoId field to given value.


### GetRoutingKeyFilter

`func (o *CreateFifoBindingRequest) GetRoutingKeyFilter() string`

GetRoutingKeyFilter returns the RoutingKeyFilter field if non-nil, zero value otherwise.

### GetRoutingKeyFilterOk

`func (o *CreateFifoBindingRequest) GetRoutingKeyFilterOk() (*string, bool)`

GetRoutingKeyFilterOk returns a tuple with the RoutingKeyFilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoutingKeyFilter

`func (o *CreateFifoBindingRequest) SetRoutingKeyFilter(v string)`

SetRoutingKeyFilter sets RoutingKeyFilter field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


