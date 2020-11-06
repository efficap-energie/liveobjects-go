# LoraDeviceCreateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationType** | **string** | LoRa device activation type | 
**AppEUI** | **string** | Application identifier (AppEUI) | 
**AppKey** | **string** | Application Key (AppKey) | 
**ConnectivityOptions** | Pointer to [**LoraConnectivityOptionsWeb**](LoraConnectivityOptionsWeb.md) |  | [optional] 
**ConnectivityPlan** | Pointer to **string** | Device connectivity plan. | [optional] 
**DevEUI** | **string** | LoRa End-device identifier (DevEUI) | 
**DeviceStatus** | **string** | LoraDeviceUpdateReqWeb status | 
**Encoding** | Pointer to **string** | LoRa device encoding | [optional] 
**Name** | **string** | LoRa device name | 
**Profile** | **string** | LoRa device profile | 
**Tags** | Pointer to **[]string** | List of tags, used to tag uplink from this device | [optional] 

## Methods

### NewLoraDeviceCreateReqWeb

`func NewLoraDeviceCreateReqWeb(activationType string, appEUI string, appKey string, devEUI string, deviceStatus string, name string, profile string, ) *LoraDeviceCreateReqWeb`

NewLoraDeviceCreateReqWeb instantiates a new LoraDeviceCreateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraDeviceCreateReqWebWithDefaults

`func NewLoraDeviceCreateReqWebWithDefaults() *LoraDeviceCreateReqWeb`

NewLoraDeviceCreateReqWebWithDefaults instantiates a new LoraDeviceCreateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetActivationType

`func (o *LoraDeviceCreateReqWeb) GetActivationType() string`

GetActivationType returns the ActivationType field if non-nil, zero value otherwise.

### GetActivationTypeOk

`func (o *LoraDeviceCreateReqWeb) GetActivationTypeOk() (*string, bool)`

GetActivationTypeOk returns a tuple with the ActivationType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActivationType

`func (o *LoraDeviceCreateReqWeb) SetActivationType(v string)`

SetActivationType sets ActivationType field to given value.


### GetAppEUI

`func (o *LoraDeviceCreateReqWeb) GetAppEUI() string`

GetAppEUI returns the AppEUI field if non-nil, zero value otherwise.

### GetAppEUIOk

`func (o *LoraDeviceCreateReqWeb) GetAppEUIOk() (*string, bool)`

GetAppEUIOk returns a tuple with the AppEUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEUI

`func (o *LoraDeviceCreateReqWeb) SetAppEUI(v string)`

SetAppEUI sets AppEUI field to given value.


### GetAppKey

`func (o *LoraDeviceCreateReqWeb) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *LoraDeviceCreateReqWeb) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *LoraDeviceCreateReqWeb) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.


### GetConnectivityOptions

`func (o *LoraDeviceCreateReqWeb) GetConnectivityOptions() LoraConnectivityOptionsWeb`

GetConnectivityOptions returns the ConnectivityOptions field if non-nil, zero value otherwise.

### GetConnectivityOptionsOk

`func (o *LoraDeviceCreateReqWeb) GetConnectivityOptionsOk() (*LoraConnectivityOptionsWeb, bool)`

GetConnectivityOptionsOk returns a tuple with the ConnectivityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityOptions

`func (o *LoraDeviceCreateReqWeb) SetConnectivityOptions(v LoraConnectivityOptionsWeb)`

SetConnectivityOptions sets ConnectivityOptions field to given value.

### HasConnectivityOptions

`func (o *LoraDeviceCreateReqWeb) HasConnectivityOptions() bool`

HasConnectivityOptions returns a boolean if a field has been set.

### GetConnectivityPlan

`func (o *LoraDeviceCreateReqWeb) GetConnectivityPlan() string`

GetConnectivityPlan returns the ConnectivityPlan field if non-nil, zero value otherwise.

### GetConnectivityPlanOk

`func (o *LoraDeviceCreateReqWeb) GetConnectivityPlanOk() (*string, bool)`

GetConnectivityPlanOk returns a tuple with the ConnectivityPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityPlan

`func (o *LoraDeviceCreateReqWeb) SetConnectivityPlan(v string)`

SetConnectivityPlan sets ConnectivityPlan field to given value.

### HasConnectivityPlan

`func (o *LoraDeviceCreateReqWeb) HasConnectivityPlan() bool`

HasConnectivityPlan returns a boolean if a field has been set.

### GetDevEUI

`func (o *LoraDeviceCreateReqWeb) GetDevEUI() string`

GetDevEUI returns the DevEUI field if non-nil, zero value otherwise.

### GetDevEUIOk

`func (o *LoraDeviceCreateReqWeb) GetDevEUIOk() (*string, bool)`

GetDevEUIOk returns a tuple with the DevEUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDevEUI

`func (o *LoraDeviceCreateReqWeb) SetDevEUI(v string)`

SetDevEUI sets DevEUI field to given value.


### GetDeviceStatus

`func (o *LoraDeviceCreateReqWeb) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *LoraDeviceCreateReqWeb) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *LoraDeviceCreateReqWeb) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.


### GetEncoding

`func (o *LoraDeviceCreateReqWeb) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *LoraDeviceCreateReqWeb) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *LoraDeviceCreateReqWeb) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *LoraDeviceCreateReqWeb) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetName

`func (o *LoraDeviceCreateReqWeb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoraDeviceCreateReqWeb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoraDeviceCreateReqWeb) SetName(v string)`

SetName sets Name field to given value.


### GetProfile

`func (o *LoraDeviceCreateReqWeb) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *LoraDeviceCreateReqWeb) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *LoraDeviceCreateReqWeb) SetProfile(v string)`

SetProfile sets Profile field to given value.


### GetTags

`func (o *LoraDeviceCreateReqWeb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoraDeviceCreateReqWeb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoraDeviceCreateReqWeb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoraDeviceCreateReqWeb) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


