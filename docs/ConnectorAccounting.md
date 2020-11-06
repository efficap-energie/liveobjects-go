# ConnectorAccounting

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Inventory** | Pointer to **map[string]int64** |  | [optional] 
**Traffic** | Pointer to [**Traffic**](Traffic.md) |  | [optional] 

## Methods

### NewConnectorAccounting

`func NewConnectorAccounting() *ConnectorAccounting`

NewConnectorAccounting instantiates a new ConnectorAccounting object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectorAccountingWithDefaults

`func NewConnectorAccountingWithDefaults() *ConnectorAccounting`

NewConnectorAccountingWithDefaults instantiates a new ConnectorAccounting object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetInventory

`func (o *ConnectorAccounting) GetInventory() map[string]int64`

GetInventory returns the Inventory field if non-nil, zero value otherwise.

### GetInventoryOk

`func (o *ConnectorAccounting) GetInventoryOk() (*map[string]int64, bool)`

GetInventoryOk returns a tuple with the Inventory field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInventory

`func (o *ConnectorAccounting) SetInventory(v map[string]int64)`

SetInventory sets Inventory field to given value.

### HasInventory

`func (o *ConnectorAccounting) HasInventory() bool`

HasInventory returns a boolean if a field has been set.

### GetTraffic

`func (o *ConnectorAccounting) GetTraffic() Traffic`

GetTraffic returns the Traffic field if non-nil, zero value otherwise.

### GetTrafficOk

`func (o *ConnectorAccounting) GetTrafficOk() (*Traffic, bool)`

GetTrafficOk returns a tuple with the Traffic field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTraffic

`func (o *ConnectorAccounting) SetTraffic(v Traffic)`

SetTraffic sets Traffic field to given value.

### HasTraffic

`func (o *ConnectorAccounting) HasTraffic() bool`

HasTraffic returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


