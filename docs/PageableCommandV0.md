# PageableCommandV0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int64** | number of the current page: starts at 0. | 
**Size** | **int64** | number of data per page (&#x3D; maximum number of data in the associated list of data:the last page can have less data) | 
**TotalCount** | **int64** | total count of data in the complete list. | 
**Data** | [**[]CommandV0**](CommandV0.md) | list of data in this page. | 

## Methods

### NewPageableCommandV0

`func NewPageableCommandV0(page int64, size int64, totalCount int64, data []CommandV0, ) *PageableCommandV0`

NewPageableCommandV0 instantiates a new PageableCommandV0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableCommandV0WithDefaults

`func NewPageableCommandV0WithDefaults() *PageableCommandV0`

NewPageableCommandV0WithDefaults instantiates a new PageableCommandV0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageableCommandV0) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageableCommandV0) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageableCommandV0) SetPage(v int64)`

SetPage sets Page field to given value.


### GetSize

`func (o *PageableCommandV0) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageableCommandV0) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageableCommandV0) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTotalCount

`func (o *PageableCommandV0) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PageableCommandV0) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PageableCommandV0) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetData

`func (o *PageableCommandV0) GetData() []CommandV0`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageableCommandV0) GetDataOk() (*[]CommandV0, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageableCommandV0) SetData(v []CommandV0)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


