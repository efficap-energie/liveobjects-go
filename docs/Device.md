# Device

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | Device URN | 
**Name** | **string** | Human readable name | [optional] 
**Description** | **string** | Device description | [optional] 
**Tags** | **[]string** | Device tags | [optional] 
**Properties** | **map[string]string** | Device properties (from device provisioning) | [optional] 
**Group** | [**DeviceGroup**](DeviceGroup.md) |  | [optional] 
**Interfaces** | [**[]DeviceInterface**](DeviceInterface.md) | List of this device&#39;s interfaces (i.e. &#39;connectivity nodes&#39;) | [optional] 
**Config** | [**map[string]DeviceParameterValue**](DeviceParameterValue.md) | Device configuration (last reported parameter values) | [optional] 
**Firmwares** | **map[string]string** | Device firmware versions | [optional] 
**DefaultDataStreamId** | **string** | default data streamId | [optional] 
**Created** | **string** | Date/time when device was first registered | 
**Updated** | **string** | Date/time when device status has been lastly updated | [optional] 
**ActivityState** | **string** | Activity state of the device according to the activity rules set for this device | [optional] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


