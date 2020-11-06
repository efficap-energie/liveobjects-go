# ConnectorStatusResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**StatusPerMsisdn** | Pointer to **map[string]string** | list all msisdn with status treated | [optional] 

## Methods

### NewConnectorStatusResponse

`func NewConnectorStatusResponse() *ConnectorStatusResponse`

NewConnectorStatusResponse instantiates a new ConnectorStatusResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorStatusResponseWithDefaults

`func NewConnectorStatusResponseWithDefaults() *ConnectorStatusResponse`

NewConnectorStatusResponseWithDefaults instantiates a new ConnectorStatusResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatusPerMsisdn

`func (o *ConnectorStatusResponse) GetStatusPerMsisdn() map[string]string`

GetStatusPerMsisdn returns the StatusPerMsisdn field if non-nil, zero value otherwise.

### GetStatusPerMsisdnOk

`func (o *ConnectorStatusResponse) GetStatusPerMsisdnOk() (*map[string]string, bool)`

GetStatusPerMsisdnOk returns a tuple with the StatusPerMsisdn field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusPerMsisdn

`func (o *ConnectorStatusResponse) SetStatusPerMsisdn(v map[string]string)`

SetStatusPerMsisdn sets StatusPerMsisdn field to given value.

### HasStatusPerMsisdn

`func (o *ConnectorStatusResponse) HasStatusPerMsisdn() bool`

HasStatusPerMsisdn returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


