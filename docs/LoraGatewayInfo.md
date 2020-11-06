# LoraGatewayInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Config** | Pointer to [**GwConfig**](GwConfig.md) |  | [optional] 
**Description** | Pointer to **string** | Description | [optional] 
**Id** | Pointer to **string** | Unique id of the gateway | [optional] 
**LastReportTs** | Pointer to **string** | Date/time of the last report | [optional] 
**Location** | Pointer to [**LoraGatewayLocation**](LoraGatewayLocation.md) |  | [optional] 
**Manufacturer** | Pointer to **string** | Manufacturer of the gateway | [optional] 
**Name** | Pointer to **string** | Name of the gateway | [optional] 
**Status** | Pointer to **string** | Status of the gateway | [optional] 
**System** | Pointer to [**GwSystem**](GwSystem.md) |  | [optional] 
**Type** | Pointer to **string** | Type of the gateway | [optional] 

## Methods

### NewLoraGatewayInfo

`func NewLoraGatewayInfo() *LoraGatewayInfo`

NewLoraGatewayInfo instantiates a new LoraGatewayInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraGatewayInfoWithDefaults

`func NewLoraGatewayInfoWithDefaults() *LoraGatewayInfo`

NewLoraGatewayInfoWithDefaults instantiates a new LoraGatewayInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetConfig

`func (o *LoraGatewayInfo) GetConfig() GwConfig`

GetConfig returns the Config field if non-nil, zero value otherwise.

### GetConfigOk

`func (o *LoraGatewayInfo) GetConfigOk() (*GwConfig, bool)`

GetConfigOk returns a tuple with the Config field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfig

`func (o *LoraGatewayInfo) SetConfig(v GwConfig)`

SetConfig sets Config field to given value.

### HasConfig

`func (o *LoraGatewayInfo) HasConfig() bool`

HasConfig returns a boolean if a field has been set.

### GetDescription

`func (o *LoraGatewayInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *LoraGatewayInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *LoraGatewayInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *LoraGatewayInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *LoraGatewayInfo) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LoraGatewayInfo) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LoraGatewayInfo) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *LoraGatewayInfo) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLastReportTs

`func (o *LoraGatewayInfo) GetLastReportTs() string`

GetLastReportTs returns the LastReportTs field if non-nil, zero value otherwise.

### GetLastReportTsOk

`func (o *LoraGatewayInfo) GetLastReportTsOk() (*string, bool)`

GetLastReportTsOk returns a tuple with the LastReportTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastReportTs

`func (o *LoraGatewayInfo) SetLastReportTs(v string)`

SetLastReportTs sets LastReportTs field to given value.

### HasLastReportTs

`func (o *LoraGatewayInfo) HasLastReportTs() bool`

HasLastReportTs returns a boolean if a field has been set.

### GetLocation

`func (o *LoraGatewayInfo) GetLocation() LoraGatewayLocation`

GetLocation returns the Location field if non-nil, zero value otherwise.

### GetLocationOk

`func (o *LoraGatewayInfo) GetLocationOk() (*LoraGatewayLocation, bool)`

GetLocationOk returns a tuple with the Location field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLocation

`func (o *LoraGatewayInfo) SetLocation(v LoraGatewayLocation)`

SetLocation sets Location field to given value.

### HasLocation

`func (o *LoraGatewayInfo) HasLocation() bool`

HasLocation returns a boolean if a field has been set.

### GetManufacturer

`func (o *LoraGatewayInfo) GetManufacturer() string`

GetManufacturer returns the Manufacturer field if non-nil, zero value otherwise.

### GetManufacturerOk

`func (o *LoraGatewayInfo) GetManufacturerOk() (*string, bool)`

GetManufacturerOk returns a tuple with the Manufacturer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManufacturer

`func (o *LoraGatewayInfo) SetManufacturer(v string)`

SetManufacturer sets Manufacturer field to given value.

### HasManufacturer

`func (o *LoraGatewayInfo) HasManufacturer() bool`

HasManufacturer returns a boolean if a field has been set.

### GetName

`func (o *LoraGatewayInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LoraGatewayInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LoraGatewayInfo) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LoraGatewayInfo) HasName() bool`

HasName returns a boolean if a field has been set.

### GetStatus

`func (o *LoraGatewayInfo) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *LoraGatewayInfo) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *LoraGatewayInfo) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *LoraGatewayInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetSystem

`func (o *LoraGatewayInfo) GetSystem() GwSystem`

GetSystem returns the System field if non-nil, zero value otherwise.

### GetSystemOk

`func (o *LoraGatewayInfo) GetSystemOk() (*GwSystem, bool)`

GetSystemOk returns a tuple with the System field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSystem

`func (o *LoraGatewayInfo) SetSystem(v GwSystem)`

SetSystem sets System field to given value.

### HasSystem

`func (o *LoraGatewayInfo) HasSystem() bool`

HasSystem returns a boolean if a field has been set.

### GetType

`func (o *LoraGatewayInfo) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *LoraGatewayInfo) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *LoraGatewayInfo) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *LoraGatewayInfo) HasType() bool`

HasType returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


