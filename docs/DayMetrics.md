# DayMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Connectors** | Pointer to [**map[string]ConnectorAccounting**](ConnectorAccounting.md) | connector (lora, mqtt, http) statistics | [optional] 
**Day** | Pointer to **string** | day in \&quot;YYYY-MM-DD\&quot; format. | [optional] 
**Services** | Pointer to [**map[string]map[string]int64**](map.md) | service statistics | [optional] 

## Methods

### NewDayMetrics

`func NewDayMetrics() *DayMetrics`

NewDayMetrics instantiates a new DayMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDayMetricsWithDefaults

`func NewDayMetricsWithDefaults() *DayMetrics`

NewDayMetricsWithDefaults instantiates a new DayMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConnectors

`func (o *DayMetrics) GetConnectors() map[string]ConnectorAccounting`

GetConnectors returns the Connectors field if non-nil, zero value otherwise.

### GetConnectorsOk

`func (o *DayMetrics) GetConnectorsOk() (*map[string]ConnectorAccounting, bool)`

GetConnectorsOk returns a tuple with the Connectors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectors

`func (o *DayMetrics) SetConnectors(v map[string]ConnectorAccounting)`

SetConnectors sets Connectors field to given value.

### HasConnectors

`func (o *DayMetrics) HasConnectors() bool`

HasConnectors returns a boolean if a field has been set.

### GetDay

`func (o *DayMetrics) GetDay() string`

GetDay returns the Day field if non-nil, zero value otherwise.

### GetDayOk

`func (o *DayMetrics) GetDayOk() (*string, bool)`

GetDayOk returns a tuple with the Day field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDay

`func (o *DayMetrics) SetDay(v string)`

SetDay sets Day field to given value.

### HasDay

`func (o *DayMetrics) HasDay() bool`

HasDay returns a boolean if a field has been set.

### GetServices

`func (o *DayMetrics) GetServices() map[string]map[string]int64`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *DayMetrics) GetServicesOk() (*map[string]map[string]int64, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *DayMetrics) SetServices(v map[string]map[string]int64)`

SetServices sets Services field to given value.

### HasServices

`func (o *DayMetrics) HasServices() bool`

HasServices returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


