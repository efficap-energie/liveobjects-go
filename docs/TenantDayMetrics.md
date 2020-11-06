# TenantDayMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to [**TenantExternalView**](TenantExternalView.md) |  | [optional] 
**Days** | Pointer to [**[]DayMetrics**](DayMetrics.md) | Statistics per day | [optional] 

## Methods

### NewTenantDayMetrics

`func NewTenantDayMetrics() *TenantDayMetrics`

NewTenantDayMetrics instantiates a new TenantDayMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantDayMetricsWithDefaults

`func NewTenantDayMetricsWithDefaults() *TenantDayMetrics`

NewTenantDayMetricsWithDefaults instantiates a new TenantDayMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *TenantDayMetrics) GetTenant() TenantExternalView`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TenantDayMetrics) GetTenantOk() (*TenantExternalView, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TenantDayMetrics) SetTenant(v TenantExternalView)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TenantDayMetrics) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetDays

`func (o *TenantDayMetrics) GetDays() []DayMetrics`

GetDays returns the Days field if non-nil, zero value otherwise.

### GetDaysOk

`func (o *TenantDayMetrics) GetDaysOk() (*[]DayMetrics, bool)`

GetDaysOk returns a tuple with the Days field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDays

`func (o *TenantDayMetrics) SetDays(v []DayMetrics)`

SetDays sets Days field to given value.

### HasDays

`func (o *TenantDayMetrics) HasDays() bool`

HasDays returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


