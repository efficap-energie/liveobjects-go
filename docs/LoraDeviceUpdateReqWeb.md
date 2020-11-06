# LoraDeviceUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AppEUI** | Pointer to **string** | Application identifier (AppEUI) | [optional] 
**AppKey** | Pointer to **string** | Application Key (AppKey) | [optional] 
**ConnectivityOptions** | Pointer to [**LoraConnectivityOptionsWeb**](LoraConnectivityOptionsWeb.md) |  | [optional] 
**ConnectivityPlan** | Pointer to **string** | Device connectivity plan. | [optional] 
**DeviceStatus** | Pointer to **string** | LoraDeviceUpdateReqWeb status | [optional] 
**Encoding** | Pointer to **string** | LoRa device encoding | [optional] 
**Name** | Pointer to **string** | LoRa device name | [optional] 
**Profile** | Pointer to **string** | LoRa device profile | [optional] 
**Tags** | Pointer to **[]string** | List of tags, used to tag uplink from this device | [optional] 

## Methods

### NewLoraDeviceUpdateReqWeb

`func NewLoraDeviceUpdateReqWeb() *LoraDeviceUpdateReqWeb`

NewLoraDeviceUpdateReqWeb instantiates a new LoraDeviceUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraDeviceUpdateReqWebWithDefaults

`func NewLoraDeviceUpdateReqWebWithDefaults() *LoraDeviceUpdateReqWeb`

NewLoraDeviceUpdateReqWebWithDefaults instantiates a new LoraDeviceUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAppEUI

`func (o *LoraDeviceUpdateReqWeb) GetAppEUI() string`

GetAppEUI returns the AppEUI field if non-nil, zero value otherwise.

### GetAppEUIOk

`func (o *LoraDeviceUpdateReqWeb) GetAppEUIOk() (*string, bool)`

GetAppEUIOk returns a tuple with the AppEUI field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppEUI

`func (o *LoraDeviceUpdateReqWeb) SetAppEUI(v string)`

SetAppEUI sets AppEUI field to given value.

### HasAppEUI

`func (o *LoraDeviceUpdateReqWeb) HasAppEUI() bool`

HasAppEUI returns a boolean if a field has been set.

### GetAppKey

`func (o *LoraDeviceUpdateReqWeb) GetAppKey() string`

GetAppKey returns the AppKey field if non-nil, zero value otherwise.

### GetAppKeyOk

`func (o *LoraDeviceUpdateReqWeb) GetAppKeyOk() (*string, bool)`

GetAppKeyOk returns a tuple with the AppKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAppKey

`func (o *LoraDeviceUpdateReqWeb) SetAppKey(v string)`

SetAppKey sets AppKey field to given value.

### HasAppKey

`func (o *LoraDeviceUpdateReqWeb) HasAppKey() bool`

HasAppKey returns a boolean if a field has been set.

### GetConnectivityOptions

`func (o *LoraDeviceUpdateReqWeb) GetConnectivityOptions() LoraConnectivityOptionsWeb`

GetConnectivityOptions returns the ConnectivityOptions field if non-nil, zero value otherwise.

### GetConnectivityOptionsOk

`func (o *LoraDeviceUpdateReqWeb) GetConnectivityOptionsOk() (*LoraConnectivityOptionsWeb, bool)`

GetConnectivityOptionsOk returns a tuple with the ConnectivityOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityOptions

`func (o *LoraDeviceUpdateReqWeb) SetConnectivityOptions(v LoraConnectivityOptionsWeb)`

SetConnectivityOptions sets ConnectivityOptions field to given value.

### HasConnectivityOptions

`func (o *LoraDeviceUpdateReqWeb) HasConnectivityOptions() bool`

HasConnectivityOptions returns a boolean if a field has been set.

### GetConnectivityPlan

`func (o *LoraDeviceUpdateReqWeb) GetConnectivityPlan() string`

GetConnectivityPlan returns the ConnectivityPlan field if non-nil, zero value otherwise.

### GetConnectivityPlanOk

`func (o *LoraDeviceUpdateReqWeb) GetConnectivityPlanOk() (*string, bool)`

GetConnectivityPlanOk returns a tuple with the ConnectivityPlan field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConnectivityPlan

`func (o *LoraDeviceUpdateReqWeb) SetConnectivityPlan(v string)`

SetConnectivityPlan sets ConnectivityPlan field to given value.

### HasConnectivityPlan

`func (o *LoraDeviceUpdateReqWeb) HasConnectivityPlan() bool`

HasConnectivityPlan returns a boolean if a field has been set.

### GetDeviceStatus

`func (o *LoraDeviceUpdateReqWeb) GetDeviceStatus() string`

GetDeviceStatus returns the DeviceStatus field if non-nil, zero value otherwise.

### GetDeviceStatusOk

`func (o *LoraDeviceUpdateReqWeb) GetDeviceStatusOk() (*string, bool)`

GetDeviceStatusOk returns a tuple with the DeviceStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceStatus

`func (o *LoraDeviceUpdateReqWeb) SetDeviceStatus(v string)`

SetDeviceStatus sets DeviceStatus field to given value.

### HasDeviceStatus

`func (o *LoraDeviceUpdateReqWeb) HasDeviceStatus() bool`

HasDeviceStatus returns a boolean if a field has been set.

### GetEncoding

`func (o *LoraDeviceUpdateReqWeb) GetEncoding() string`

GetEncoding returns the Encoding field if non-nil, zero value otherwise.

### GetEncodingOk

`func (o *LoraDeviceUpdateReqWeb) GetEncodingOk() (*string, bool)`

GetEncodingOk returns a tuple with the Encoding field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEncoding

`func (o *LoraDeviceUpdateReqWeb) SetEncoding(v string)`

SetEncoding sets Encoding field to given value.

### HasEncoding

`func (o *LoraDeviceUpdateReqWeb) HasEncoding() bool`

HasEncoding returns a boolean if a field has been set.

### GetName

`func (o *LoraDeviceUpdateReqWeb) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoraDeviceUpdateReqWeb) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoraDeviceUpdateReqWeb) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoraDeviceUpdateReqWeb) HasName() bool`

HasName returns a boolean if a field has been set.

### GetProfile

`func (o *LoraDeviceUpdateReqWeb) GetProfile() string`

GetProfile returns the Profile field if non-nil, zero value otherwise.

### GetProfileOk

`func (o *LoraDeviceUpdateReqWeb) GetProfileOk() (*string, bool)`

GetProfileOk returns a tuple with the Profile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProfile

`func (o *LoraDeviceUpdateReqWeb) SetProfile(v string)`

SetProfile sets Profile field to given value.

### HasProfile

`func (o *LoraDeviceUpdateReqWeb) HasProfile() bool`

HasProfile returns a boolean if a field has been set.

### GetTags

`func (o *LoraDeviceUpdateReqWeb) GetTags() []string`

GetTags returns the Tags field if non-nil, zero value otherwise.

### GetTagsOk

`func (o *LoraDeviceUpdateReqWeb) GetTagsOk() (*[]string, bool)`

GetTagsOk returns a tuple with the Tags field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTags

`func (o *LoraDeviceUpdateReqWeb) SetTags(v []string)`

SetTags sets Tags field to given value.

### HasTags

`func (o *LoraDeviceUpdateReqWeb) HasTags() bool`

HasTags returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


