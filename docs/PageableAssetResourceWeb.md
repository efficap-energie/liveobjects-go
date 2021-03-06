# PageableAssetResourceWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int64** | number of the current page: starts at 0. | 
**Size** | **int64** | number of data per page (&#x3D; maximum number of data in the associated list of data:the last page can have less data) | 
**TotalCount** | **int64** | total count of data in the complete list. | 
**Data** | [**[]AssetResourceWeb**](AssetResourceWeb.md) | list of data in this page. | 

## Methods

### NewPageableAssetResourceWeb

`func NewPageableAssetResourceWeb(page int64, size int64, totalCount int64, data []AssetResourceWeb, ) *PageableAssetResourceWeb`

NewPageableAssetResourceWeb instantiates a new PageableAssetResourceWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableAssetResourceWebWithDefaults

`func NewPageableAssetResourceWebWithDefaults() *PageableAssetResourceWeb`

NewPageableAssetResourceWebWithDefaults instantiates a new PageableAssetResourceWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageableAssetResourceWeb) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageableAssetResourceWeb) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageableAssetResourceWeb) SetPage(v int64)`

SetPage sets Page field to given value.


### GetSize

`func (o *PageableAssetResourceWeb) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageableAssetResourceWeb) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageableAssetResourceWeb) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTotalCount

`func (o *PageableAssetResourceWeb) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PageableAssetResourceWeb) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PageableAssetResourceWeb) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetData

`func (o *PageableAssetResourceWeb) GetData() []AssetResourceWeb`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageableAssetResourceWeb) GetDataOk() (*[]AssetResourceWeb, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageableAssetResourceWeb) SetData(v []AssetResourceWeb)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


