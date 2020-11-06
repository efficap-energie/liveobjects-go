# AdminInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusinessUnit** | Pointer to [**BusinessUnit**](BusinessUnit.md) |  | [optional] 
**Country** | Pointer to **string** |  | [optional] 
**SupportVendorIds** | Pointer to **[]string** |  | [optional] 
**Vendor** | Pointer to [**Vendor**](Vendor.md) |  | [optional] 

## Methods

### NewAdminInfo

`func NewAdminInfo() *AdminInfo`

NewAdminInfo instantiates a new AdminInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminInfoWithDefaults

`func NewAdminInfoWithDefaults() *AdminInfo`

NewAdminInfoWithDefaults instantiates a new AdminInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusinessUnit

`func (o *AdminInfo) GetBusinessUnit() BusinessUnit`

GetBusinessUnit returns the BusinessUnit field if non-nil, zero value otherwise.

### GetBusinessUnitOk

`func (o *AdminInfo) GetBusinessUnitOk() (*BusinessUnit, bool)`

GetBusinessUnitOk returns a tuple with the BusinessUnit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusinessUnit

`func (o *AdminInfo) SetBusinessUnit(v BusinessUnit)`

SetBusinessUnit sets BusinessUnit field to given value.

### HasBusinessUnit

`func (o *AdminInfo) HasBusinessUnit() bool`

HasBusinessUnit returns a boolean if a field has been set.

### GetCountry

`func (o *AdminInfo) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *AdminInfo) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *AdminInfo) SetCountry(v string)`

SetCountry sets Country field to given value.

### HasCountry

`func (o *AdminInfo) HasCountry() bool`

HasCountry returns a boolean if a field has been set.

### GetSupportVendorIds

`func (o *AdminInfo) GetSupportVendorIds() []string`

GetSupportVendorIds returns the SupportVendorIds field if non-nil, zero value otherwise.

### GetSupportVendorIdsOk

`func (o *AdminInfo) GetSupportVendorIdsOk() (*[]string, bool)`

GetSupportVendorIdsOk returns a tuple with the SupportVendorIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportVendorIds

`func (o *AdminInfo) SetSupportVendorIds(v []string)`

SetSupportVendorIds sets SupportVendorIds field to given value.

### HasSupportVendorIds

`func (o *AdminInfo) HasSupportVendorIds() bool`

HasSupportVendorIds returns a boolean if a field has been set.

### GetVendor

`func (o *AdminInfo) GetVendor() Vendor`

GetVendor returns the Vendor field if non-nil, zero value otherwise.

### GetVendorOk

`func (o *AdminInfo) GetVendorOk() (*Vendor, bool)`

GetVendorOk returns a tuple with the Vendor field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVendor

`func (o *AdminInfo) SetVendor(v Vendor)`

SetVendor sets Vendor field to given value.

### HasVendor

`func (o *AdminInfo) HasVendor() bool`

HasVendor returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


