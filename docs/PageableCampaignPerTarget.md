# PageableCampaignPerTarget

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int64** | number of the current page: starts at 0. | 
**Size** | **int64** | number of data per page (&#x3D; maximum number of data in the associated list of data:the last page can have less data) | 
**TotalCount** | **int64** | total count of data in the complete list. | 
**Data** | [**[]CampaignPerTarget**](CampaignPerTarget.md) | list of data in this page. | 

## Methods

### NewPageableCampaignPerTarget

`func NewPageableCampaignPerTarget(page int64, size int64, totalCount int64, data []CampaignPerTarget, ) *PageableCampaignPerTarget`

NewPageableCampaignPerTarget instantiates a new PageableCampaignPerTarget object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableCampaignPerTargetWithDefaults

`func NewPageableCampaignPerTargetWithDefaults() *PageableCampaignPerTarget`

NewPageableCampaignPerTargetWithDefaults instantiates a new PageableCampaignPerTarget object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageableCampaignPerTarget) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageableCampaignPerTarget) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageableCampaignPerTarget) SetPage(v int64)`

SetPage sets Page field to given value.


### GetSize

`func (o *PageableCampaignPerTarget) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageableCampaignPerTarget) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageableCampaignPerTarget) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTotalCount

`func (o *PageableCampaignPerTarget) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PageableCampaignPerTarget) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PageableCampaignPerTarget) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetData

`func (o *PageableCampaignPerTarget) GetData() []CampaignPerTarget`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageableCampaignPerTarget) GetDataOk() (*[]CampaignPerTarget, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageableCampaignPerTarget) SetData(v []CampaignPerTarget)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


