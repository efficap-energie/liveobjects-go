# TenantExternalView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **string** | tenant Id | [optional] 
**Name** | Pointer to **string** | tenant name | [optional] 
**OfferAndOptions** | Pointer to [**OfferAndOptions**](OfferAndOptions.md) |  | [optional] 
**Properties** | Pointer to **map[string]string** | tenant properties | [optional] 

## Methods

### NewTenantExternalView

`func NewTenantExternalView() *TenantExternalView`

NewTenantExternalView instantiates a new TenantExternalView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTenantExternalViewWithDefaults

`func NewTenantExternalViewWithDefaults() *TenantExternalView`

NewTenantExternalViewWithDefaults instantiates a new TenantExternalView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TenantExternalView) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TenantExternalView) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TenantExternalView) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *TenantExternalView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *TenantExternalView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *TenantExternalView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *TenantExternalView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *TenantExternalView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOfferAndOptions

`func (o *TenantExternalView) GetOfferAndOptions() OfferAndOptions`

GetOfferAndOptions returns the OfferAndOptions field if non-nil, zero value otherwise.

### GetOfferAndOptionsOk

`func (o *TenantExternalView) GetOfferAndOptionsOk() (*OfferAndOptions, bool)`

GetOfferAndOptionsOk returns a tuple with the OfferAndOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOfferAndOptions

`func (o *TenantExternalView) SetOfferAndOptions(v OfferAndOptions)`

SetOfferAndOptions sets OfferAndOptions field to given value.

### HasOfferAndOptions

`func (o *TenantExternalView) HasOfferAndOptions() bool`

HasOfferAndOptions returns a boolean if a field has been set.

### GetProperties

`func (o *TenantExternalView) GetProperties() map[string]string`

GetProperties returns the Properties field if non-nil, zero value otherwise.

### GetPropertiesOk

`func (o *TenantExternalView) GetPropertiesOk() (*map[string]string, bool)`

GetPropertiesOk returns a tuple with the Properties field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProperties

`func (o *TenantExternalView) SetProperties(v map[string]string)`

SetProperties sets Properties field to given value.

### HasProperties

`func (o *TenantExternalView) HasProperties() bool`

HasProperties returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


