# TenantMonthMetrics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Tenant** | Pointer to [**TenantExternalView**](TenantExternalView.md) |  | [optional] 
**Months** | Pointer to [**[]MonthMetrics**](MonthMetrics.md) | Statistics per month | [optional] 

## Methods

### NewTenantMonthMetrics

`func NewTenantMonthMetrics() *TenantMonthMetrics`

NewTenantMonthMetrics instantiates a new TenantMonthMetrics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantMonthMetricsWithDefaults

`func NewTenantMonthMetricsWithDefaults() *TenantMonthMetrics`

NewTenantMonthMetricsWithDefaults instantiates a new TenantMonthMetrics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenant

`func (o *TenantMonthMetrics) GetTenant() TenantExternalView`

GetTenant returns the Tenant field if non-nil, zero value otherwise.

### GetTenantOk

`func (o *TenantMonthMetrics) GetTenantOk() (*TenantExternalView, bool)`

GetTenantOk returns a tuple with the Tenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenant

`func (o *TenantMonthMetrics) SetTenant(v TenantExternalView)`

SetTenant sets Tenant field to given value.

### HasTenant

`func (o *TenantMonthMetrics) HasTenant() bool`

HasTenant returns a boolean if a field has been set.

### GetMonths

`func (o *TenantMonthMetrics) GetMonths() []MonthMetrics`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *TenantMonthMetrics) GetMonthsOk() (*[]MonthMetrics, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *TenantMonthMetrics) SetMonths(v []MonthMetrics)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *TenantMonthMetrics) HasMonths() bool`

HasMonths returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


