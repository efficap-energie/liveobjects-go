# LoraDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationType** | **string** | LoRa device activation type | 
**AppEUI** | **string** | Application identifier (AppEUI) | 
**ConnectivityOptions** | [**LoraConnectivityOptions**](LoraConnectivityOptions.md) |  | 
**ConnectivityPlan** | Pointer to **string** | Device connectivity plan. | [optional] 
**CreationTs** | Pointer to **string** | Date/time of the device creation | [optional] [readonly] 
**DevEUI** | **string** | LoRa End-device identifier (DevEUI) | 
**DeviceStatus** | **string** | Device status | 
**Encoding** | **string** | LoRa encoder | 
**LastActivationTs** | Pointer to **string** | Date/time of the last activation | [optional] [readonly] 
**LastBatteryLevel** | Pointer to **int32** | Last battery level returned by the device (0: external power source, 1..254: 1&#x3D;min / 254 &#x3D; max, 255 : NA) | [optional] [readonly] 
**LastCommunicationTs** | Pointer to **string** | Date/time of the last communication | [optional] [readonly] 
**LastDeactivationTs** | Pointer to **string** | Date/time of the last deactivation | [optional] [readonly] 
**LastDlFcnt** | Pointer to **int32** | Frame counter of the last downlink | [optional] [readonly] 
**LastSignalLevel** | Pointer to **int32** | Signal level of the last uplink (between 0 to 5) | [optional] [readonly] 
**LastUlFcnt** | Pointer to **int32** | Frame counter of the last uplink | [optional] [readonly] 
**Location** | Pointer to [**LoraDeviceLocation**](LoraDeviceLocation.md) |  | [optional] 
**Name** | **string** | LoRa device name | 
**Profile** | **string** | LoRa device profile | 
**Tags** | Pointer to **[]string** | List of tags, used to tag uplink from this device | [optional] 
**UpdateTs** | Pointer to **string** | Date/time of the last device modification | [optional] [readonly] 

## Methods

### NewLoraDevice

`func NewLoraDevice(activationType string, appEUI string, connectivityOptions LoraConnectivityOptions, devEUI string, deviceStatus string, encoding string, name string, profile string, ) *LoraDevice`

NewLoraDevice instantiates a new LoraDevice object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraDeviceWithDefaults

`func NewLoraDeviceWithDefaults() *LoraDevice`

NewLoraDeviceWithDefaults instantiates a new LoraDevice object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationType

`func (o *LoraDevice) GetActivationType() string`

GetActivationType returns the ActivationType field if non-nil, zero value otherwise.

### GetActivationTypeOk

`func (o *LoraDevice) GetActivationTypeOk() (*string, bool)`

GetActivationTypeOk returns a tuple with the ActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationType

`func (o *LoraDevice) SetActivationType(v string)`

SetActivationType sets ActivationType field to given value.


### GetAppEUI

`func (o *LoraDevice) GetAppEUI() string`

GetAppEUI returns the AppEUI field if non-nil, zero value otherwise.

### GetAppEUIOk

`func (o *LoraDevice) GetAppEUIOk() (*string, bool)`

GetAppEUIOk returns a tuple with the AppEUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEUI

`func (o *LoraDevice) SetAppEUI(v string)`

SetAppEUI sets AppEUI field to given value.


### GetConnectivityOptions

`func (o *LoraDevice) GetConnectivityOptions() LoraConnectivityOptions`

GetConnectivityOptions returns the ConnectivityOptions field if non-nil, zero value otherwise.

### GetConnectivityOptionsOk

`func (o *LoraDevice) GetConnectivityOptionsOk() (*LoraConnectivityOptions, bool)`

GetConnectivityOptionsOk returns a tuple with the ConnectivityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityOptions

`func (o *LoraDevice) SetConnectivityOptions(v LoraConnectivityOptions)`

SetConnectivityOptions sets ConnectivityOptions field to given value.


### GetConnectivityPlan

`func (o *LoraDevice) GetConnectivityPlan() string`

GetConnectivityPlan returns the ConnectivityPlan field if non-nil, zero value otherwise.

### GetConnectivityPlanOk

`func (o *LoraDevice) GetConnectivityPlanOk() (*string, bool)`

GetConnectivityPlanOk returns a tuple with the ConnectivityPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityPlan

`func (o *LoraDevice) SetConnectivityPlan(v string)`

SetConnectivityPlan sets ConnectivityPlan field to given value.

### HasConnectivityPlan

`func (o *LoraDevice) HasConnectivityPlan() bool`

HasConnectivityPlan returns a boolean if a field has been set.

### GetCreationTs

`func (o *LoraDevice) GetCreationTs() string`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *LoraDevice) GetCreationTsOk() (*string, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *LoraDevice) SetCreationTs(v string)`

SetCreationTs sets CreationTs field to given value.

### HasCreationTs

`func (o *LoraDevice) HasCreationTs() bool`

HasCreationTs returns a boolean if a field has been set.

### GetDevEUI

`func (o *LoraDevice) GetDevEUI() string`

GetDevEUI returns the DevEUI field if non-nil, zero value otherwise.

### GetDevEUIOk

`func (o *LoraDevice) GetDevEUIOk() (*string, bool)`

GetDevEUIOk returns a tuple with the DevEUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevEUI

`func (o *LoraDevice) SetDevEUI(v string)`

SetDevEUI sets DevEUI field to given value.


### GetDeviceStatus

`func (o *LoraDevice) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *LoraDevice) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *LoraDevice) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.


### GetEncoding

`func (o *LoraDevice) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *LoraDevice) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *LoraDevice) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.


### GetLastActivationTs

`func (o *LoraDevice) GetLastActivationTs() string`

GetLastActivationTs returns the LastActivationTs field if non-nil, zero value otherwise.

### GetLastActivationTsOk

`func (o *LoraDevice) GetLastActivationTsOk() (*string, bool)`

GetLastActivationTsOk returns a tuple with the LastActivationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastActivationTs

`func (o *LoraDevice) SetLastActivationTs(v string)`

SetLastActivationTs sets LastActivationTs field to given value.

### HasLastActivationTs

`func (o *LoraDevice) HasLastActivationTs() bool`

HasLastActivationTs returns a boolean if a field has been set.

### GetLastBatteryLevel

`func (o *LoraDevice) GetLastBatteryLevel() int32`

GetLastBatteryLevel returns the LastBatteryLevel field if non-nil, zero value otherwise.

### GetLastBatteryLevelOk

`func (o *LoraDevice) GetLastBatteryLevelOk() (*int32, bool)`

GetLastBatteryLevelOk returns a tuple with the LastBatteryLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastBatteryLevel

`func (o *LoraDevice) SetLastBatteryLevel(v int32)`

SetLastBatteryLevel sets LastBatteryLevel field to given value.

### HasLastBatteryLevel

`func (o *LoraDevice) HasLastBatteryLevel() bool`

HasLastBatteryLevel returns a boolean if a field has been set.

### GetLastCommunicationTs

`func (o *LoraDevice) GetLastCommunicationTs() string`

GetLastCommunicationTs returns the LastCommunicationTs field if non-nil, zero value otherwise.

### GetLastCommunicationTsOk

`func (o *LoraDevice) GetLastCommunicationTsOk() (*string, bool)`

GetLastCommunicationTsOk returns a tuple with the LastCommunicationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastCommunicationTs

`func (o *LoraDevice) SetLastCommunicationTs(v string)`

SetLastCommunicationTs sets LastCommunicationTs field to given value.

### HasLastCommunicationTs

`func (o *LoraDevice) HasLastCommunicationTs() bool`

HasLastCommunicationTs returns a boolean if a field has been set.

### GetLastDeactivationTs

`func (o *LoraDevice) GetLastDeactivationTs() string`

GetLastDeactivationTs returns the LastDeactivationTs field if non-nil, zero value otherwise.

### GetLastDeactivationTsOk

`func (o *LoraDevice) GetLastDeactivationTsOk() (*string, bool)`

GetLastDeactivationTsOk returns a tuple with the LastDeactivationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDeactivationTs

`func (o *LoraDevice) SetLastDeactivationTs(v string)`

SetLastDeactivationTs sets LastDeactivationTs field to given value.

### HasLastDeactivationTs

`func (o *LoraDevice) HasLastDeactivationTs() bool`

HasLastDeactivationTs returns a boolean if a field has been set.

### GetLastDlFcnt

`func (o *LoraDevice) GetLastDlFcnt() int32`

GetLastDlFcnt returns the LastDlFcnt field if non-nil, zero value otherwise.

### GetLastDlFcntOk

`func (o *LoraDevice) GetLastDlFcntOk() (*int32, bool)`

GetLastDlFcntOk returns a tuple with the LastDlFcnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastDlFcnt

`func (o *LoraDevice) SetLastDlFcnt(v int32)`

SetLastDlFcnt sets LastDlFcnt field to given value.

### HasLastDlFcnt

`func (o *LoraDevice) HasLastDlFcnt() bool`

HasLastDlFcnt returns a boolean if a field has been set.

### GetLastSignalLevel

`func (o *LoraDevice) GetLastSignalLevel() int32`

GetLastSignalLevel returns the LastSignalLevel field if non-nil, zero value otherwise.

### GetLastSignalLevelOk

`func (o *LoraDevice) GetLastSignalLevelOk() (*int32, bool)`

GetLastSignalLevelOk returns a tuple with the LastSignalLevel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastSignalLevel

`func (o *LoraDevice) SetLastSignalLevel(v int32)`

SetLastSignalLevel sets LastSignalLevel field to given value.

### HasLastSignalLevel

`func (o *LoraDevice) HasLastSignalLevel() bool`

HasLastSignalLevel returns a boolean if a field has been set.

### GetLastUlFcnt

`func (o *LoraDevice) GetLastUlFcnt() int32`

GetLastUlFcnt returns the LastUlFcnt field if non-nil, zero value otherwise.

### GetLastUlFcntOk

`func (o *LoraDevice) GetLastUlFcntOk() (*int32, bool)`

GetLastUlFcntOk returns a tuple with the LastUlFcnt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUlFcnt

`func (o *LoraDevice) SetLastUlFcnt(v int32)`

SetLastUlFcnt sets LastUlFcnt field to given value.

### HasLastUlFcnt

`func (o *LoraDevice) HasLastUlFcnt() bool`

HasLastUlFcnt returns a boolean if a field has been set.

### GetLocation

`func (o *LoraDevice) GetLocation() LoraDeviceLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoraDevice) GetLocationOk() (*LoraDeviceLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoraDevice) SetLocation(v LoraDeviceLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LoraDevice) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetName

`func (o *LoraDevice) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoraDevice) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoraDevice) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *LoraDevice) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *LoraDevice) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *LoraDevice) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTags

`func (o *LoraDevice) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoraDevice) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoraDevice) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoraDevice) HasTags() bool`

HasTags returns a boolean if a field has been set.

### GetUpdateTs

`func (o *LoraDevice) GetUpdateTs() string`

GetUpdateTs returns the UpdateTs field if non-nil, zero value otherwise.

### GetUpdateTsOk

`func (o *LoraDevice) GetUpdateTsOk() (*string, bool)`

GetUpdateTsOk returns a tuple with the UpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTs

`func (o *LoraDevice) SetUpdateTs(v string)`

SetUpdateTs sets UpdateTs field to given value.

### HasUpdateTs

`func (o *LoraDevice) HasUpdateTs() bool`

HasUpdateTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


