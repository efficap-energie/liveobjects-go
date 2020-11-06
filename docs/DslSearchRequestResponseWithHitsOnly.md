# DslSearchRequestResponseWithHitsOnly

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Hits** | Pointer to [**[]SearchDataMessage**](SearchDataMessage.md) |  | [optional] 
**TookMillis** | Pointer to **int64** |  | [optional] 
**TotalHits** | Pointer to **int64** |  | [optional] 

## Methods

### NewDslSearchRequestResponseWithHitsOnly

`func NewDslSearchRequestResponseWithHitsOnly() *DslSearchRequestResponseWithHitsOnly`

NewDslSearchRequestResponseWithHitsOnly instantiates a new DslSearchRequestResponseWithHitsOnly object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDslSearchRequestResponseWithHitsOnlyWithDefaults

`func NewDslSearchRequestResponseWithHitsOnlyWithDefaults() *DslSearchRequestResponseWithHitsOnly`

NewDslSearchRequestResponseWithHitsOnlyWithDefaults instantiates a new DslSearchRequestResponseWithHitsOnly object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHits

`func (o *DslSearchRequestResponseWithHitsOnly) GetHits() []SearchDataMessage`

GetHits returns the Hits field if non-nil, zero value otherwise.

### GetHitsOk

`func (o *DslSearchRequestResponseWithHitsOnly) GetHitsOk() (*[]SearchDataMessage, bool)`

GetHitsOk returns a tuple with the Hits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHits

`func (o *DslSearchRequestResponseWithHitsOnly) SetHits(v []SearchDataMessage)`

SetHits sets Hits field to given value.

### HasHits

`func (o *DslSearchRequestResponseWithHitsOnly) HasHits() bool`

HasHits returns a boolean if a field has been set.

### GetTookMillis

`func (o *DslSearchRequestResponseWithHitsOnly) GetTookMillis() int64`

GetTookMillis returns the TookMillis field if non-nil, zero value otherwise.

### GetTookMillisOk

`func (o *DslSearchRequestResponseWithHitsOnly) GetTookMillisOk() (*int64, bool)`

GetTookMillisOk returns a tuple with the TookMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTookMillis

`func (o *DslSearchRequestResponseWithHitsOnly) SetTookMillis(v int64)`

SetTookMillis sets TookMillis field to given value.

### HasTookMillis

`func (o *DslSearchRequestResponseWithHitsOnly) HasTookMillis() bool`

HasTookMillis returns a boolean if a field has been set.

### GetTotalHits

`func (o *DslSearchRequestResponseWithHitsOnly) GetTotalHits() int64`

GetTotalHits returns the TotalHits field if non-nil, zero value otherwise.

### GetTotalHitsOk

`func (o *DslSearchRequestResponseWithHitsOnly) GetTotalHitsOk() (*int64, bool)`

GetTotalHitsOk returns a tuple with the TotalHits field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalHits

`func (o *DslSearchRequestResponseWithHitsOnly) SetTotalHits(v int64)`

SetTotalHits sets TotalHits field to given value.

### HasTotalHits

`func (o *DslSearchRequestResponseWithHitsOnly) HasTotalHits() bool`

HasTotalHits returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


