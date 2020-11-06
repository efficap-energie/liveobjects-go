# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Device URN | 
**Name** | Pointer to **string** | Human readable name | [optional] 
**Description** | Pointer to **string** | Device description | [optional] 
**Tags** | Pointer to **[]string** | Device tags | [optional] 
**Properties** | Pointer to **map[string]string** | Device properties (from device provisioning) | [optional] 
**Group** | Pointer to [**DeviceGroup**](DeviceGroup.md) |  | [optional] 
**Interfaces** | Pointer to [**[]DeviceInterface**](DeviceInterface.md) | List of this device&#39;s interfaces (i.e. &#39;connectivity nodes&#39;) | [optional] 
**Config** | Pointer to [**map[string]DeviceParameterValue**](DeviceParameterValue.md) | Device configuration (last reported parameter values) | [optional] 
**Firmwares** | Pointer to **map[string]string** | Device firmware versions | [optional] 
**Resources** | Pointer to **map[string]string** | Device resource versions | [optional] 
**DefaultDataStreamId** | Pointer to **string** | default data streamId | [optional] 
**Created** | **string** | Date/time when device was first registered | 
**Updated** | Pointer to **string** | Date/time when device status has been lastly updated | [optional] 
**ActivityState** | Pointer to **string** | Activity state of the device according to the activity rules set for this device | [optional] 

## Methods

### NewDevice

`func NewDevice(id string, created string, ) *Device`

NewDevice instantiates a new Device object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceWithDefaults

`func NewDeviceWithDefaults() *Device`

NewDeviceWithDefaults instantiates a new Device object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Device) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Device) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Device) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Device) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Device) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Device) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Device) HasName() bool`

HasName returns a boolean if a field has been set.

### GetDescription

`func (o *Device) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Device) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Device) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Device) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetTags

`func (o *Device) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *Device) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *Device) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *Device) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetProperties

`func (o *Device) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *Device) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *Device) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *Device) HasProperties() bool`

HasProperties returns a boolean if a field has been set.

### GetGroup

`func (o *Device) GetGroup() DeviceGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *Device) GetGroupOk() (*DeviceGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *Device) SetGroup(v DeviceGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *Device) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetInterfaces

`func (o *Device) GetInterfaces() []DeviceInterface`

GetInterfaces returns the Interfaces field if non-nil, zero value otherwise.

### GetInterfacesOk

`func (o *Device) GetInterfacesOk() (*[]DeviceInterface, bool)`

GetInterfacesOk returns a tuple with the Interfaces field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterfaces

`func (o *Device) SetInterfaces(v []DeviceInterface)`

SetInterfaces sets Interfaces field to given value.

### HasInterfaces

`func (o *Device) HasInterfaces() bool`

HasInterfaces returns a boolean if a field has been set.

### GetConfig

`func (o *Device) GetConfig() map[string]DeviceParameterValue`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *Device) GetConfigOk() (*map[string]DeviceParameterValue, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *Device) SetConfig(v map[string]DeviceParameterValue)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *Device) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetFirmwares

`func (o *Device) GetFirmwares() map[string]string`

GetFirmwares returns the Firmwares field if non-nil, zero value otherwise.

### GetFirmwaresOk

`func (o *Device) GetFirmwaresOk() (*map[string]string, bool)`

GetFirmwaresOk returns a tuple with the Firmwares field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirmwares

`func (o *Device) SetFirmwares(v map[string]string)`

SetFirmwares sets Firmwares field to given value.

### HasFirmwares

`func (o *Device) HasFirmwares() bool`

HasFirmwares returns a boolean if a field has been set.

### GetResources

`func (o *Device) GetResources() map[string]string`

GetResources returns the Resources field if non-nil, zero value otherwise.

### GetResourcesOk

`func (o *Device) GetResourcesOk() (*map[string]string, bool)`

GetResourcesOk returns a tuple with the Resources field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResources

`func (o *Device) SetResources(v map[string]string)`

SetResources sets Resources field to given value.

### HasResources

`func (o *Device) HasResources() bool`

HasResources returns a boolean if a field has been set.

### GetDefaultDataStreamId

`func (o *Device) GetDefaultDataStreamId() string`

GetDefaultDataStreamId returns the DefaultDataStreamId field if non-nil, zero value otherwise.

### GetDefaultDataStreamIdOk

`func (o *Device) GetDefaultDataStreamIdOk() (*string, bool)`

GetDefaultDataStreamIdOk returns a tuple with the DefaultDataStreamId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultDataStreamId

`func (o *Device) SetDefaultDataStreamId(v string)`

SetDefaultDataStreamId sets DefaultDataStreamId field to given value.

### HasDefaultDataStreamId

`func (o *Device) HasDefaultDataStreamId() bool`

HasDefaultDataStreamId returns a boolean if a field has been set.

### GetCreated

`func (o *Device) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *Device) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *Device) SetCreated(v string)`

SetCreated sets Created field to given value.


### GetUpdated

`func (o *Device) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *Device) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *Device) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *Device) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.

### GetActivityState

`func (o *Device) GetActivityState() string`

GetActivityState returns the ActivityState field if non-nil, zero value otherwise.

### GetActivityStateOk

`func (o *Device) GetActivityStateOk() (*string, bool)`

GetActivityStateOk returns a tuple with the ActivityState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivityState

`func (o *Device) SetActivityState(v string)`

SetActivityState sets ActivityState field to given value.

### HasActivityState

`func (o *Device) HasActivityState() bool`

HasActivityState returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


