# CommandRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connector** | **string** | connector/protocol to use to forward the command | 
**Value** | **map[string]interface{}** | command value (protocol/connector-dependant). The length is limited | 

## Methods

### NewCommandRequest

`func NewCommandRequest(connector string, value map[string]interface{}, ) *CommandRequest`

NewCommandRequest instantiates a new CommandRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandRequestWithDefaults

`func NewCommandRequestWithDefaults() *CommandRequest`

NewCommandRequestWithDefaults instantiates a new CommandRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnector

`func (o *CommandRequest) GetConnector() string`

GetConnector returns the Connector field if non-nil, zero value otherwise.

### GetConnectorOk

`func (o *CommandRequest) GetConnectorOk() (*string, bool)`

GetConnectorOk returns a tuple with the Connector field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnector

`func (o *CommandRequest) SetConnector(v string)`

SetConnector sets Connector field to given value.


### GetValue

`func (o *CommandRequest) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *CommandRequest) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *CommandRequest) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


