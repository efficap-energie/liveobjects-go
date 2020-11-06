# TenantStatsWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Properties** | Pointer to **map[string]string** |  | [optional] 
**StatisticsPerDay** | Pointer to [**map[string]map[string]ConnectorStatistics**](map.md) | all statistics per day | [optional] 
**StatisticsPerMonth** | Pointer to [**map[string]map[string]ConnectorStatistics**](map.md) | Statistics per month : number of distinct sources (lora, mqtt, datazone) | [optional] 
**TenantId** | Pointer to **string** | The tenant Id | [optional] 
**TenantName** | Pointer to **string** | The tenant name | [optional] 
**Total** | Pointer to [**map[string]ConnectorStatistics**](ConnectorStatistics.md) | Aggregation of all statistics | [optional] 

## Methods

### NewTenantStatsWeb

`func NewTenantStatsWeb() *TenantStatsWeb`

NewTenantStatsWeb instantiates a new TenantStatsWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantStatsWebWithDefaults

`func NewTenantStatsWebWithDefaults() *TenantStatsWeb`

NewTenantStatsWebWithDefaults instantiates a new TenantStatsWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetProperties

`func (o *TenantStatsWeb) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TenantStatsWeb) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TenantStatsWeb) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TenantStatsWeb) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetStatisticsPerDay

`func (o *TenantStatsWeb) GetStatisticsPerDay() map[string]map[string]ConnectorStatistics`

GetStatisticsPerDay returns the StatisticsPerDay field if non-nil, zero value otherwise.

### GetStatisticsPerDayOk

`func (o *TenantStatsWeb) GetStatisticsPerDayOk() (*map[string]map[string]ConnectorStatistics, bool)`

GetStatisticsPerDayOk returns a tuple with the StatisticsPerDay field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsPerDay

`func (o *TenantStatsWeb) SetStatisticsPerDay(v map[string]map[string]ConnectorStatistics)`

SetStatisticsPerDay sets StatisticsPerDay field to given value.

### HasStatisticsPerDay

`func (o *TenantStatsWeb) HasStatisticsPerDay() bool`

HasStatisticsPerDay returns a boolean if a field has been set.

### GetStatisticsPerMonth

`func (o *TenantStatsWeb) GetStatisticsPerMonth() map[string]map[string]ConnectorStatistics`

GetStatisticsPerMonth returns the StatisticsPerMonth field if non-nil, zero value otherwise.

### GetStatisticsPerMonthOk

`func (o *TenantStatsWeb) GetStatisticsPerMonthOk() (*map[string]map[string]ConnectorStatistics, bool)`

GetStatisticsPerMonthOk returns a tuple with the StatisticsPerMonth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatisticsPerMonth

`func (o *TenantStatsWeb) SetStatisticsPerMonth(v map[string]map[string]ConnectorStatistics)`

SetStatisticsPerMonth sets StatisticsPerMonth field to given value.

### HasStatisticsPerMonth

`func (o *TenantStatsWeb) HasStatisticsPerMonth() bool`

HasStatisticsPerMonth returns a boolean if a field has been set.

### GetTenantId

`func (o *TenantStatsWeb) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *TenantStatsWeb) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *TenantStatsWeb) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *TenantStatsWeb) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTenantName

`func (o *TenantStatsWeb) GetTenantName() string`

GetTenantName returns the TenantName field if non-nil, zero value otherwise.

### GetTenantNameOk

`func (o *TenantStatsWeb) GetTenantNameOk() (*string, bool)`

GetTenantNameOk returns a tuple with the TenantName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantName

`func (o *TenantStatsWeb) SetTenantName(v string)`

SetTenantName sets TenantName field to given value.

### HasTenantName

`func (o *TenantStatsWeb) HasTenantName() bool`

HasTenantName returns a boolean if a field has been set.

### GetTotal

`func (o *TenantStatsWeb) GetTotal() map[string]ConnectorStatistics`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *TenantStatsWeb) GetTotalOk() (*map[string]ConnectorStatistics, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *TenantStatsWeb) SetTotal(v map[string]ConnectorStatistics)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *TenantStatsWeb) HasTotal() bool`

HasTotal returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


