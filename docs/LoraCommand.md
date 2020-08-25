# LoraCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandStatus** | **string** | Status of the command. SENT: The command was injected into LPWA network core. ERROR: The command could injected into LPWA network core. | 
**Confirmed** | **bool** | Network ack confirmation | 
**CreationTs** | **string** | Date/time of the command creation | [optional] [readonly] 
**Data** | **string** | Hexadecimal raw data of the command | 
**Id** | **string** | Unique id of the command | 
**Port** | **int32** | Port of the device on which the command was sent (cf. LoRaWan) | 
**UpdateTs** | **string** | Date/time of the last command modification | [optional] [readonly] 

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


