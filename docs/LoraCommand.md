# LoraCommand

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CommandStatus** | **string** | Status of the command. SENT: The command was injected into LPWA network core. ERROR: The command could injected into LPWA network core. | 
**Confirmed** | **bool** | Network ack confirmation | 
**CreationTs** | Pointer to **string** | Date/time of the command creation | [optional] [readonly] 
**Data** | **string** | Hexadecimal raw data of the command | 
**Id** | **string** | Unique id of the command | 
**Port** | **int32** | Port of the device on which the command was sent (cf. LoRaWan) | 
**UpdateTs** | Pointer to **string** | Date/time of the last command modification | [optional] [readonly] 

## Methods

### NewLoraCommand

`func NewLoraCommand(commandStatus string, confirmed bool, data string, id string, port int32, ) *LoraCommand`

NewLoraCommand instantiates a new LoraCommand object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraCommandWithDefaults

`func NewLoraCommandWithDefaults() *LoraCommand`

NewLoraCommandWithDefaults instantiates a new LoraCommand object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommandStatus

`func (o *LoraCommand) GetCommandStatus() string`

GetCommandStatus returns the CommandStatus field if non-nil, zero value otherwise.

### GetCommandStatusOk

`func (o *LoraCommand) GetCommandStatusOk() (*string, bool)`

GetCommandStatusOk returns a tuple with the CommandStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommandStatus

`func (o *LoraCommand) SetCommandStatus(v string)`

SetCommandStatus sets CommandStatus field to given value.


### GetConfirmed

`func (o *LoraCommand) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *LoraCommand) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *LoraCommand) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.


### GetCreationTs

`func (o *LoraCommand) GetCreationTs() string`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *LoraCommand) GetCreationTsOk() (*string, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *LoraCommand) SetCreationTs(v string)`

SetCreationTs sets CreationTs field to given value.

### HasCreationTs

`func (o *LoraCommand) HasCreationTs() bool`

HasCreationTs returns a boolean if a field has been set.

### GetData

`func (o *LoraCommand) GetData() string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *LoraCommand) GetDataOk() (*string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *LoraCommand) SetData(v string)`

SetData sets Data field to given value.


### GetId

`func (o *LoraCommand) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoraCommand) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoraCommand) SetId(v string)`

SetId sets Id field to given value.


### GetPort

`func (o *LoraCommand) GetPort() int32`

GetPort returns the Port field if non-nil, zero value otherwise.

### GetPortOk

`func (o *LoraCommand) GetPortOk() (*int32, bool)`

GetPortOk returns a tuple with the Port field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPort

`func (o *LoraCommand) SetPort(v int32)`

SetPort sets Port field to given value.


### GetUpdateTs

`func (o *LoraCommand) GetUpdateTs() string`

GetUpdateTs returns the UpdateTs field if non-nil, zero value otherwise.

### GetUpdateTsOk

`func (o *LoraCommand) GetUpdateTsOk() (*string, bool)`

GetUpdateTsOk returns a tuple with the UpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTs

`func (o *LoraCommand) SetUpdateTs(v string)`

SetUpdateTs sets UpdateTs field to given value.

### HasUpdateTs

`func (o *LoraCommand) HasUpdateTs() bool`

HasUpdateTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


