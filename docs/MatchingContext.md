# MatchingContext

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | [**NewData**](NewData.md) |  | 
**MatchingRule** | [**MatchingRule**](MatchingRule.md) |  | 
**TenantId** | Pointer to **string** |  | [optional] 
**Timestamp** | Pointer to **string** |  | [optional] 

## Methods

### NewMatchingContext

`func NewMatchingContext(data NewData, matchingRule MatchingRule, ) *MatchingContext`

NewMatchingContext instantiates a new MatchingContext object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMatchingContextWithDefaults

`func NewMatchingContextWithDefaults() *MatchingContext`

NewMatchingContextWithDefaults instantiates a new MatchingContext object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *MatchingContext) GetData() NewData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *MatchingContext) GetDataOk() (*NewData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *MatchingContext) SetData(v NewData)`

SetData sets Data field to given value.


### GetMatchingRule

`func (o *MatchingContext) GetMatchingRule() MatchingRule`

GetMatchingRule returns the MatchingRule field if non-nil, zero value otherwise.

### GetMatchingRuleOk

`func (o *MatchingContext) GetMatchingRuleOk() (*MatchingRule, bool)`

GetMatchingRuleOk returns a tuple with the MatchingRule field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMatchingRule

`func (o *MatchingContext) SetMatchingRule(v MatchingRule)`

SetMatchingRule sets MatchingRule field to given value.


### GetTenantId

`func (o *MatchingContext) GetTenantId() string`

GetTenantId returns the TenantId field if non-nil, zero value otherwise.

### GetTenantIdOk

`func (o *MatchingContext) GetTenantIdOk() (*string, bool)`

GetTenantIdOk returns a tuple with the TenantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantId

`func (o *MatchingContext) SetTenantId(v string)`

SetTenantId sets TenantId field to given value.

### HasTenantId

`func (o *MatchingContext) HasTenantId() bool`

HasTenantId returns a boolean if a field has been set.

### GetTimestamp

`func (o *MatchingContext) GetTimestamp() string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *MatchingContext) GetTimestampOk() (*string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *MatchingContext) SetTimestamp(v string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *MatchingContext) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


