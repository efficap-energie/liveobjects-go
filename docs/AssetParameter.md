# AssetParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**LastKnownValueTs** | Pointer to **int64** |  | [optional] 
**TargetValue** | Pointer to [**AssetParameterValue**](AssetParameterValue.md) |  | [optional] 
**UpdateReqCorrelationId** | Pointer to **int32** |  | [optional] 
**UpdateStatus** | Pointer to **string** |  | [optional] 
**UpdateStatusTs** | Pointer to **int64** |  | [optional] 
**Value** | Pointer to [**AssetParameterValue**](AssetParameterValue.md) |  | [optional] 

## Methods

### NewAssetParameter

`func NewAssetParameter() *AssetParameter`

NewAssetParameter instantiates a new AssetParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetParameterWithDefaults

`func NewAssetParameterWithDefaults() *AssetParameter`

NewAssetParameterWithDefaults instantiates a new AssetParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetLastKnownValueTs

`func (o *AssetParameter) GetLastKnownValueTs() int64`

GetLastKnownValueTs returns the LastKnownValueTs field if non-nil, zero value otherwise.

### GetLastKnownValueTsOk

`func (o *AssetParameter) GetLastKnownValueTsOk() (*int64, bool)`

GetLastKnownValueTsOk returns a tuple with the LastKnownValueTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastKnownValueTs

`func (o *AssetParameter) SetLastKnownValueTs(v int64)`

SetLastKnownValueTs sets LastKnownValueTs field to given value.

### HasLastKnownValueTs

`func (o *AssetParameter) HasLastKnownValueTs() bool`

HasLastKnownValueTs returns a boolean if a field has been set.

### GetTargetValue

`func (o *AssetParameter) GetTargetValue() AssetParameterValue`

GetTargetValue returns the TargetValue field if non-nil, zero value otherwise.

### GetTargetValueOk

`func (o *AssetParameter) GetTargetValueOk() (*AssetParameterValue, bool)`

GetTargetValueOk returns a tuple with the TargetValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTargetValue

`func (o *AssetParameter) SetTargetValue(v AssetParameterValue)`

SetTargetValue sets TargetValue field to given value.

### HasTargetValue

`func (o *AssetParameter) HasTargetValue() bool`

HasTargetValue returns a boolean if a field has been set.

### GetUpdateReqCorrelationId

`func (o *AssetParameter) GetUpdateReqCorrelationId() int32`

GetUpdateReqCorrelationId returns the UpdateReqCorrelationId field if non-nil, zero value otherwise.

### GetUpdateReqCorrelationIdOk

`func (o *AssetParameter) GetUpdateReqCorrelationIdOk() (*int32, bool)`

GetUpdateReqCorrelationIdOk returns a tuple with the UpdateReqCorrelationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateReqCorrelationId

`func (o *AssetParameter) SetUpdateReqCorrelationId(v int32)`

SetUpdateReqCorrelationId sets UpdateReqCorrelationId field to given value.

### HasUpdateReqCorrelationId

`func (o *AssetParameter) HasUpdateReqCorrelationId() bool`

HasUpdateReqCorrelationId returns a boolean if a field has been set.

### GetUpdateStatus

`func (o *AssetParameter) GetUpdateStatus() string`

GetUpdateStatus returns the UpdateStatus field if non-nil, zero value otherwise.

### GetUpdateStatusOk

`func (o *AssetParameter) GetUpdateStatusOk() (*string, bool)`

GetUpdateStatusOk returns a tuple with the UpdateStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatus

`func (o *AssetParameter) SetUpdateStatus(v string)`

SetUpdateStatus sets UpdateStatus field to given value.

### HasUpdateStatus

`func (o *AssetParameter) HasUpdateStatus() bool`

HasUpdateStatus returns a boolean if a field has been set.

### GetUpdateStatusTs

`func (o *AssetParameter) GetUpdateStatusTs() int64`

GetUpdateStatusTs returns the UpdateStatusTs field if non-nil, zero value otherwise.

### GetUpdateStatusTsOk

`func (o *AssetParameter) GetUpdateStatusTsOk() (*int64, bool)`

GetUpdateStatusTsOk returns a tuple with the UpdateStatusTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateStatusTs

`func (o *AssetParameter) SetUpdateStatusTs(v int64)`

SetUpdateStatusTs sets UpdateStatusTs field to given value.

### HasUpdateStatusTs

`func (o *AssetParameter) HasUpdateStatusTs() bool`

HasUpdateStatusTs returns a boolean if a field has been set.

### GetValue

`func (o *AssetParameter) GetValue() AssetParameterValue`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *AssetParameter) GetValueOk() (*AssetParameterValue, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *AssetParameter) SetValue(v AssetParameterValue)`

SetValue sets Value field to given value.

### HasValue

`func (o *AssetParameter) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


