# InterfaceCapabilities

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Command** | Pointer to [**InterfaceCapability**](InterfaceCapability.md) |  | [optional] 
**Configuration** | Pointer to [**InterfaceCapability**](InterfaceCapability.md) |  | [optional] 
**Resource** | Pointer to [**InterfaceCapability**](InterfaceCapability.md) |  | [optional] 

## Methods

### NewInterfaceCapabilities

`func NewInterfaceCapabilities() *InterfaceCapabilities`

NewInterfaceCapabilities instantiates a new InterfaceCapabilities object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInterfaceCapabilitiesWithDefaults

`func NewInterfaceCapabilitiesWithDefaults() *InterfaceCapabilities`

NewInterfaceCapabilitiesWithDefaults instantiates a new InterfaceCapabilities object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCommand

`func (o *InterfaceCapabilities) GetCommand() InterfaceCapability`

GetCommand returns the Command field if non-nil, zero value otherwise.

### GetCommandOk

`func (o *InterfaceCapabilities) GetCommandOk() (*InterfaceCapability, bool)`

GetCommandOk returns a tuple with the Command field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCommand

`func (o *InterfaceCapabilities) SetCommand(v InterfaceCapability)`

SetCommand sets Command field to given value.

### HasCommand

`func (o *InterfaceCapabilities) HasCommand() bool`

HasCommand returns a boolean if a field has been set.

### GetConfiguration

`func (o *InterfaceCapabilities) GetConfiguration() InterfaceCapability`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *InterfaceCapabilities) GetConfigurationOk() (*InterfaceCapability, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *InterfaceCapabilities) SetConfiguration(v InterfaceCapability)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *InterfaceCapabilities) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetResource

`func (o *InterfaceCapabilities) GetResource() InterfaceCapability`

GetResource returns the Resource field if non-nil, zero value otherwise.

### GetResourceOk

`func (o *InterfaceCapabilities) GetResourceOk() (*InterfaceCapability, bool)`

GetResourceOk returns a tuple with the Resource field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResource

`func (o *InterfaceCapabilities) SetResource(v InterfaceCapability)`

SetResource sets Resource field to given value.

### HasResource

`func (o *InterfaceCapabilities) HasResource() bool`

HasResource returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


