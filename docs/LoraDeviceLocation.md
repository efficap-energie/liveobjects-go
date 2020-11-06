# LoraDeviceLocation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to **float64** | Last global geolocation tolerance (meter) | [optional] [readonly] 
**Alt** | Pointer to **float64** | Last geolocation altitude (meter) | [optional] [readonly] 
**LastUpdateTs** | Pointer to **string** | Date/time of the last location | [optional] [readonly] 
**Lat** | Pointer to **float64** | Last geolocation latitude (GPS coordinate system) | [optional] [readonly] 
**Lon** | Pointer to **float64** | Last geolocation longitude (GPS coordinate system) | [optional] [readonly] 
**Provider** | Pointer to **string** | Computing geolocation method | [optional] [readonly] 

## Methods

### NewLoraDeviceLocation

`func NewLoraDeviceLocation() *LoraDeviceLocation`

NewLoraDeviceLocation instantiates a new LoraDeviceLocation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLoraDeviceLocationWithDefaults

`func NewLoraDeviceLocationWithDefaults() *LoraDeviceLocation`

NewLoraDeviceLocationWithDefaults instantiates a new LoraDeviceLocation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *LoraDeviceLocation) GetAccuracy() float64`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *LoraDeviceLocation) GetAccuracyOk() (*float64, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *LoraDeviceLocation) SetAccuracy(v float64)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *LoraDeviceLocation) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetAlt

`func (o *LoraDeviceLocation) GetAlt() float64`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *LoraDeviceLocation) GetAltOk() (*float64, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *LoraDeviceLocation) SetAlt(v float64)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *LoraDeviceLocation) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### GetLastUpdateTs

`func (o *LoraDeviceLocation) GetLastUpdateTs() string`

GetLastUpdateTs returns the LastUpdateTs field if non-nil, zero value otherwise.

### GetLastUpdateTsOk

`func (o *LoraDeviceLocation) GetLastUpdateTsOk() (*string, bool)`

GetLastUpdateTsOk returns a tuple with the LastUpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastUpdateTs

`func (o *LoraDeviceLocation) SetLastUpdateTs(v string)`

SetLastUpdateTs sets LastUpdateTs field to given value.

### HasLastUpdateTs

`func (o *LoraDeviceLocation) HasLastUpdateTs() bool`

HasLastUpdateTs returns a boolean if a field has been set.

### GetLat

`func (o *LoraDeviceLocation) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *LoraDeviceLocation) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *LoraDeviceLocation) SetLat(v float64)`

SetLat sets Lat field to given value.

### HasLat

`func (o *LoraDeviceLocation) HasLat() bool`

HasLat returns a boolean if a field has been set.

### GetLon

`func (o *LoraDeviceLocation) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *LoraDeviceLocation) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *LoraDeviceLocation) SetLon(v float64)`

SetLon sets Lon field to given value.

### HasLon

`func (o *LoraDeviceLocation) HasLon() bool`

HasLon returns a boolean if a field has been set.

### GetProvider

`func (o *LoraDeviceLocation) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *LoraDeviceLocation) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *LoraDeviceLocation) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *LoraDeviceLocation) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


