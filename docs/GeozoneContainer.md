# GeozoneContainer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional] 
**Geometry** | Pointer to [**Polygon**](Polygon.md) |  | [optional] 
**Properties** | Pointer to **map[string]interface{}** |  | [optional] 
**Tags** | Pointer to **[]string** |  | [optional] 
**ZoneId** | Pointer to **string** |  | [optional] 

## Methods

### NewGeozoneContainer

`func NewGeozoneContainer() *GeozoneContainer`

NewGeozoneContainer instantiates a new GeozoneContainer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeozoneContainerWithDefaults

`func NewGeozoneContainerWithDefaults() *GeozoneContainer`

NewGeozoneContainerWithDefaults instantiates a new GeozoneContainer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GeozoneContainer) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeozoneContainer) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeozoneContainer) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeozoneContainer) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGeometry

`func (o *GeozoneContainer) GetGeometry() Polygon`

GetGeometry returns the Geometry field if non-nil, zero value otherwise.

### GetGeometryOk

`func (o *GeozoneContainer) GetGeometryOk() (*Polygon, bool)`

GetGeometryOk returns a tuple with the Geometry field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGeometry

`func (o *GeozoneContainer) SetGeometry(v Polygon)`

SetGeometry sets Geometry field to given value.

### HasGeometry

`func (o *GeozoneContainer) HasGeometry() bool`

HasGeometry returns a boolean if a field has been set.

### GetProperties

`func (o *GeozoneContainer) GetProperties() map[string]interface{}`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *GeozoneContainer) GetPropertiesOk() (*map[string]interface{}, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *GeozoneContainer) SetProperties(v map[string]interface{})`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *GeozoneContainer) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *GeozoneContainer) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *GeozoneContainer) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *GeozoneContainer) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *GeozoneContainer) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetZoneId

`func (o *GeozoneContainer) GetZoneId() string`

GetZoneId returns the ZoneId field if non-nil, zero value otherwise.

### GetZoneIdOk

`func (o *GeozoneContainer) GetZoneIdOk() (*string, bool)`

GetZoneIdOk returns a tuple with the ZoneId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetZoneId

`func (o *GeozoneContainer) SetZoneId(v string)`

SetZoneId sets ZoneId field to given value.

### HasZoneId

`func (o *GeozoneContainer) HasZoneId() bool`

HasZoneId returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


