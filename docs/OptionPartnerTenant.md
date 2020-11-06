# OptionPartnerTenant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** | option identifier | 
**Parameters** | Pointer to **map[string]int32** | Parameters | [optional] 

## Methods

### NewOptionPartnerTenant

`func NewOptionPartnerTenant(id string, ) *OptionPartnerTenant`

NewOptionPartnerTenant instantiates a new OptionPartnerTenant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOptionPartnerTenantWithDefaults

`func NewOptionPartnerTenantWithDefaults() *OptionPartnerTenant`

NewOptionPartnerTenantWithDefaults instantiates a new OptionPartnerTenant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OptionPartnerTenant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OptionPartnerTenant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OptionPartnerTenant) SetId(v string)`

SetId sets Id field to given value.


### GetParameters

`func (o *OptionPartnerTenant) GetParameters() map[string]int32`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *OptionPartnerTenant) GetParametersOk() (*map[string]int32, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *OptionPartnerTenant) SetParameters(v map[string]int32)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *OptionPartnerTenant) HasParameters() bool`

HasParameters returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


