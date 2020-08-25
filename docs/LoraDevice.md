# LoraDevice

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ActivationType** | **string** | LoRa device activation type | 
**AppEUI** | **string** | Application identifier (AppEUI) | 
**ConnectivityOptions** | [**LoraConnectivityOptions**](LoraConnectivityOptions.md) |  | 
**ConnectivityPlan** | **string** | Device connectivity plan. | [optional] 
**CreationTs** | **string** | Date/time of the device creation | [optional] [readonly] 
**DevEUI** | **string** | LoRa End-device identifier (DevEUI) | 
**DeviceStatus** | **string** | Device status | 
**Encoding** | **string** | LoRa encoder | 
**LastActivationTs** | **string** | Date/time of the last activation | [optional] [readonly] 
**LastBatteryLevel** | **int32** | Last battery level returned by the device (0: external power source, 1..254: 1&#x3D;min / 254 &#x3D; max, 255 : NA) | [optional] [readonly] 
**LastCommunicationTs** | **string** | Date/time of the last communication | [optional] [readonly] 
**LastDeactivationTs** | **string** | Date/time of the last deactivation | [optional] [readonly] 
**LastDlFcnt** | **int32** | Frame counter of the last downlink | [optional] [readonly] 
**LastSignalLevel** | **int32** | Signal level of the last uplink (between 0 to 5) | [optional] [readonly] 
**LastUlFcnt** | **int32** | Frame counter of the last uplink | [optional] [readonly] 
**Location** | [**LoraDeviceLocation**](LoraDeviceLocation.md) |  | [optional] 
**Name** | **string** | LoRa device name | 
**Profile** | **string** | LoRa device profile | 
**Tags** | **[]string** | List of tags, used to tag uplink from this device | [optional] 
**UpdateTs** | **string** | Date/time of the last device modification | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


