# DeviceCreateRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | device unique identifier. A Live Objects URN &#39;urn:lo:nsid:{namespace}:{id}&#39; must respect the following regular expression &lt;code&gt;^urn:lo:nsid:([\\w\\-_]{1,128}):([\\w\\-_:]{1,128})$&lt;/code&gt; (max 269 characters) | 
**Name** | Pointer to **string** | human-readable device name. Expected string (max 255 characters) | [optional] 
**Description** | Pointer to **string** | human-readable device description. Expected string (max 500 characters) | [optional] 
**DefaultDataStreamId** | Pointer to **string** | default data streamId. Expected not empty string. Following character are forbidden &lt;code&gt;\&quot;&#39;\\\&quot;\\\\;{}() \&quot;&lt;/code&gt; (max 255 characters) | [optional] 
**Tags** | Pointer to **[]string** | set of tags associated with the new device. Max number of tags depends of your offer settings. Tag value max length is 32. | [optional] 
**Properties** | Pointer to **map[string]string** | map of key/value string pairs detailing properties of the device. Max number of properties depends of your offer settings. A property name must not include following characters &lt;code&gt;$.&lt;/code&gt; and max length is 128. Invalid property names are : &#39;class&#39;, &#39;_class&#39;. Property value max length is 256. | [optional] 
**Interfaces** | Pointer to [**[]NewDeviceInterface**](NewDeviceInterface.md) | list of device network interfaces | [optional] 
**Group** | Pointer to [**DeviceGroup**](DeviceGroup.md) |  | [optional] 

## Methods

### NewDeviceCreateRequest

`func NewDeviceCreateRequest(id string, ) *DeviceCreateRequest`

NewDeviceCreateRequest instantiates a new DeviceCreateRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceCreateRequestWithDefaults

`func NewDeviceCreateRequestWithDefaults() *DeviceCreateRequest`

NewDeviceCreateRequestWithDefaults instantiates a new DeviceCreateRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *DeviceCreateRequest) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *DeviceCreateRequest) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *DeviceCreateRequest) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *DeviceCreateRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DeviceCreateRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DeviceCreateRequest) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DeviceCreateRequest) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *DeviceCreateRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *DeviceCreateRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *DeviceCreateRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *DeviceCreateRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDefaultDataStreamId

`func (o *DeviceCreateRequest) GetDefaultDataStreamId() string`

GetDefaultDataStreamId returns the DefaultDataStreamId field if non-nil, zero value otherwise.

### GetDefaultDataStreamIdOk

`func (o *DeviceCreateRequest) GetDefaultDataStreamIdOk() (*string, bool)`

GetDefaultDataStreamIdOk returns a tuple with the DefaultDataStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDataStreamId

`func (o *DeviceCreateRequest) SetDefaultDataStreamId(v string)`

SetDefaultDataStreamId sets DefaultDataStreamId field to given value.

### HasDefaultDataStreamId

`func (o *DeviceCreateRequest) HasDefaultDataStreamId() bool`

HasDefaultDataStreamId returns a boolean if a field has been set.

### GetTags

`func (o *DeviceCreateRequest) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *DeviceCreateRequest) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *DeviceCreateRequest) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *DeviceCreateRequest) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProperties

`func (o *DeviceCreateRequest) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *DeviceCreateRequest) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *DeviceCreateRequest) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *DeviceCreateRequest) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetInterfaces

`func (o *DeviceCreateRequest) GetInterfaces() []NewDeviceInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *DeviceCreateRequest) GetInterfacesOk() (*[]NewDeviceInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *DeviceCreateRequest) SetInterfaces(v []NewDeviceInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *DeviceCreateRequest) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetGroup

`func (o *DeviceCreateRequest) GetGroup() DeviceGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DeviceCreateRequest) GetGroupOk() (*DeviceGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DeviceCreateRequest) SetGroup(v DeviceGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DeviceCreateRequest) HasGroup() bool`

HasGroup returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


