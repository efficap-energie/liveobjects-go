# DeviceGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | group identifier. Expected string (max 6 characters) | [optional] 
**Path** | Pointer to **string** | group path. Authorized: letter (lowercase and uppercase), accented characters, number, space, dash, underscore and simple quote. A valid path must respect the following regular expression &lt;code&gt;^[\\wÀ-ÖØ-öø-ÿ&#39; -]{1,255}&lt;/code&gt;.Expected string (max 255 characters) | [optional] 

## Methods

### NewDeviceGroup

`func NewDeviceGroup() *DeviceGroup`

NewDeviceGroup instantiates a new DeviceGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceGroupWithDefaults

`func NewDeviceGroupWithDefaults() *DeviceGroup`

NewDeviceGroupWithDefaults instantiates a new DeviceGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceGroup) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceGroup) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceGroup) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *DeviceGroup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPath

`func (o *DeviceGroup) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *DeviceGroup) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *DeviceGroup) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *DeviceGroup) HasPath() bool`

HasPath returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


