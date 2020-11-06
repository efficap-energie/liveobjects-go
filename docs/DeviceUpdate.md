# DeviceUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DefaultDataStreamId** | Pointer to **string** | default data streamId. Expected not empty string. Following character are forbidden &lt;code&gt;\&quot;&#39;\\\&quot;\\\\;{}() \&quot;&lt;/code&gt; (max 255 characters) | [optional] 
**Description** | Pointer to **string** | new device description. Expected string (max 500 characters) | [optional] 
**Group** | Pointer to [**DeviceGroup**](DeviceGroup.md) |  | [optional] 
**Id** | **string** | new device identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**Name** | Pointer to **string** | new device human-readable name. Expected string (max 255 characters) | [optional] 
**Properties** | Pointer to **map[string]string** | map of key/value string pairs detailing device properties to update. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | [optional] 
**Tags** | Pointer to **[]string** | new device set of tags. Max number of tags depends of your offer settings. Tag value max length is 32. | [optional] 

## Methods

### NewDeviceUpdate

`func NewDeviceUpdate(id string, ) *DeviceUpdate`

NewDeviceUpdate instantiates a new DeviceUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceUpdateWithDefaults

`func NewDeviceUpdateWithDefaults() *DeviceUpdate`

NewDeviceUpdateWithDefaults instantiates a new DeviceUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDefaultDataStreamId

`func (o *DeviceUpdate) GetDefaultDataStreamId() string`

GetDefaultDataStreamId returns the DefaultDataStreamId field if non-nil, zero value otherwise.

### GetDefaultDataStreamIdOk

`func (o *DeviceUpdate) GetDefaultDataStreamIdOk() (*string, bool)`

GetDefaultDataStreamIdOk returns a tuple with the DefaultDataStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDataStreamId

`func (o *DeviceUpdate) SetDefaultDataStreamId(v string)`

SetDefaultDataStreamId sets DefaultDataStreamId field to given value.

### HasDefaultDataStreamId

`func (o *DeviceUpdate) HasDefaultDataStreamId() bool`

HasDefaultDataStreamId returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetGroup

`func (o *DeviceUpdate) GetGroup() DeviceGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DeviceUpdate) GetGroupOk() (*DeviceGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DeviceUpdate) SetGroup(v DeviceGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DeviceUpdate) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetId

`func (o *DeviceUpdate) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceUpdate) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceUpdate) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeviceUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceUpdate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceUpdate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProperties

`func (o *DeviceUpdate) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeviceUpdate) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeviceUpdate) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeviceUpdate) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetTags

`func (o *DeviceUpdate) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceUpdate) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceUpdate) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceUpdate) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


