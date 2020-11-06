# ContextContainerList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int64** | number of the current page: starts at 0. | 
**Size** | **int64** | number of data per page (&#x3D; maximum number of data in the associated list of data:the last page can have less data) | 
**TotalCount** | **int64** | total count of data in the complete list. | 
**Data** | [**[]ContextContainer**](ContextContainer.md) | list of data in this page. | 

## Methods

### NewContextContainerList

`func NewContextContainerList(page int64, size int64, totalCount int64, data []ContextContainer, ) *ContextContainerList`

NewContextContainerList instantiates a new ContextContainerList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextContainerListWithDefaults

`func NewContextContainerListWithDefaults() *ContextContainerList`

NewContextContainerListWithDefaults instantiates a new ContextContainerList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *ContextContainerList) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *ContextContainerList) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *ContextContainerList) SetPage(v int64)`

SetPage sets Page field to given value.


### GetSize

`func (o *ContextContainerList) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *ContextContainerList) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *ContextContainerList) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTotalCount

`func (o *ContextContainerList) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *ContextContainerList) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *ContextContainerList) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetData

`func (o *ContextContainerList) GetData() []ContextContainer`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *ContextContainerList) GetDataOk() (*[]ContextContainer, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *ContextContainerList) SetData(v []ContextContainer)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


