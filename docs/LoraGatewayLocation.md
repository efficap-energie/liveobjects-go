# LoraGatewayLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alt** | Pointer to **float64** | Last geolocation altitude (meter) | [optional] 
**LastUpdateTs** | Pointer to **string** | Date/time of the last location | [optional] 
**Lat** | Pointer to **float64** | Last geolocation latitude (GPS coordinate system) | [optional] 
**Lon** | Pointer to **float64** | Last geolocation longitude (GPS coordinate system) | [optional] 
**Provider** | Pointer to **string** | Computing geolocation method | [optional] 

## Methods

### NewLoraGatewayLocation

`func NewLoraGatewayLocation() *LoraGatewayLocation`

NewLoraGatewayLocation instantiates a new LoraGatewayLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraGatewayLocationWithDefaults

`func NewLoraGatewayLocationWithDefaults() *LoraGatewayLocation`

NewLoraGatewayLocationWithDefaults instantiates a new LoraGatewayLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlt

`func (o *LoraGatewayLocation) GetAlt() float64`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *LoraGatewayLocation) GetAltOk() (*float64, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *LoraGatewayLocation) SetAlt(v float64)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *LoraGatewayLocation) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### GetLastUpdateTs

`func (o *LoraGatewayLocation) GetLastUpdateTs() string`

GetLastUpdateTs returns the LastUpdateTs field if non-nil, zero value otherwise.

### GetLastUpdateTsOk

`func (o *LoraGatewayLocation) GetLastUpdateTsOk() (*string, bool)`

GetLastUpdateTsOk returns a tuple with the LastUpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTs

`func (o *LoraGatewayLocation) SetLastUpdateTs(v string)`

SetLastUpdateTs sets LastUpdateTs field to given value.

### HasLastUpdateTs

`func (o *LoraGatewayLocation) HasLastUpdateTs() bool`

HasLastUpdateTs returns a boolean if a field has been set.

### GetLat

`func (o *LoraGatewayLocation) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *LoraGatewayLocation) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *LoraGatewayLocation) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *LoraGatewayLocation) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *LoraGatewayLocation) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *LoraGatewayLocation) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *LoraGatewayLocation) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *LoraGatewayLocation) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetProvider

`func (o *LoraGatewayLocation) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LoraGatewayLocation) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LoraGatewayLocation) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LoraGatewayLocation) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


