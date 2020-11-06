# MonthMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to [**map[string]ConnectorAccounting**](ConnectorAccounting.md) | connector (lora, mqtt, http) statistics | [optional] 
**Month** | Pointer to **string** | month in \&quot;YYYY-MM\&quot; format. | [optional] 
**Services** | Pointer to [**map[string]map[string]int64**](map.md) | service statistics | [optional] 

## Methods

### NewMonthMetrics

`func NewMonthMetrics() *MonthMetrics`

NewMonthMetrics instantiates a new MonthMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMonthMetricsWithDefaults

`func NewMonthMetricsWithDefaults() *MonthMetrics`

NewMonthMetricsWithDefaults instantiates a new MonthMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *MonthMetrics) GetConnectors() map[string]ConnectorAccounting`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *MonthMetrics) GetConnectorsOk() (*map[string]ConnectorAccounting, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *MonthMetrics) SetConnectors(v map[string]ConnectorAccounting)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *MonthMetrics) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetMonth

`func (o *MonthMetrics) GetMonth() string`

GetMonth returns the Month field if non-nil, zero value otherwise.

### GetMonthOk

`func (o *MonthMetrics) GetMonthOk() (*string, bool)`

GetMonthOk returns a tuple with the Month field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonth

`func (o *MonthMetrics) SetMonth(v string)`

SetMonth sets Month field to given value.

### HasMonth

`func (o *MonthMetrics) HasMonth() bool`

HasMonth returns a boolean if a field has been set.

### GetServices

`func (o *MonthMetrics) GetServices() map[string]map[string]int64`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *MonthMetrics) GetServicesOk() (*map[string]map[string]int64, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *MonthMetrics) SetServices(v map[string]map[string]int64)`

SetServices sets Services field to given value.

### HasServices

`func (o *MonthMetrics) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


