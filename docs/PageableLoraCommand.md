# PageableLoraCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Page** | **int64** | number of the current page: starts at 0. | 
**Size** | **int64** | number of data per page (&#x3D; maximum number of data in the associated list of data:the last page can have less data) | 
**TotalCount** | **int64** | total count of data in the complete list. | 
**Data** | [**[]LoraCommand**](LoraCommand.md) | list of data in this page. | 

## Methods

### NewPageableLoraCommand

`func NewPageableLoraCommand(page int64, size int64, totalCount int64, data []LoraCommand, ) *PageableLoraCommand`

NewPageableLoraCommand instantiates a new PageableLoraCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPageableLoraCommandWithDefaults

`func NewPageableLoraCommandWithDefaults() *PageableLoraCommand`

NewPageableLoraCommandWithDefaults instantiates a new PageableLoraCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPage

`func (o *PageableLoraCommand) GetPage() int64`

GetPage returns the Page field if non-nil, zero value otherwise.

### GetPageOk

`func (o *PageableLoraCommand) GetPageOk() (*int64, bool)`

GetPageOk returns a tuple with the Page field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPage

`func (o *PageableLoraCommand) SetPage(v int64)`

SetPage sets Page field to given value.


### GetSize

`func (o *PageableLoraCommand) GetSize() int64`

GetSize returns the Size field if non-nil, zero value otherwise.

### GetSizeOk

`func (o *PageableLoraCommand) GetSizeOk() (*int64, bool)`

GetSizeOk returns a tuple with the Size field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSize

`func (o *PageableLoraCommand) SetSize(v int64)`

SetSize sets Size field to given value.


### GetTotalCount

`func (o *PageableLoraCommand) GetTotalCount() int64`

GetTotalCount returns the TotalCount field if non-nil, zero value otherwise.

### GetTotalCountOk

`func (o *PageableLoraCommand) GetTotalCountOk() (*int64, bool)`

GetTotalCountOk returns a tuple with the TotalCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotalCount

`func (o *PageableLoraCommand) SetTotalCount(v int64)`

SetTotalCount sets TotalCount field to given value.


### GetData

`func (o *PageableLoraCommand) GetData() []LoraCommand`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *PageableLoraCommand) GetDataOk() (*[]LoraCommand, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *PageableLoraCommand) SetData(v []LoraCommand)`

SetData sets Data field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


