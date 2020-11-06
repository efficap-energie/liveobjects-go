# GeozoneContainerBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Geometry** | Pointer to [**Polygon**](Polygon.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 

## Methods

### NewGeozoneContainerBase

`func NewGeozoneContainerBase() *GeozoneContainerBase`

NewGeozoneContainerBase instantiates a new GeozoneContainerBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeozoneContainerBaseWithDefaults

`func NewGeozoneContainerBaseWithDefaults() *GeozoneContainerBase`

NewGeozoneContainerBaseWithDefaults instantiates a new GeozoneContainerBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GeozoneContainerBase) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeozoneContainerBase) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeozoneContainerBase) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeozoneContainerBase) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGeometry

`func (o *GeozoneContainerBase) GetGeometry() Polygon`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *GeozoneContainerBase) GetGeometryOk() (*Polygon, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *GeozoneContainerBase) SetGeometry(v Polygon)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *GeozoneContainerBase) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### GetProperties

`func (o *GeozoneContainerBase) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GeozoneContainerBase) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GeozoneContainerBase) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GeozoneContainerBase) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *GeozoneContainerBase) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GeozoneContainerBase) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GeozoneContainerBase) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GeozoneContainerBase) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


