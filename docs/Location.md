# Location

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Accuracy** | Pointer to **float64** | Accuracy value | [optional] 
**Alt** | Pointer to **float64** | Altitude value | [optional] 
**Lat** | **float64** | Latitude value | 
**Lon** | **float64** | Longitude value | 
**Provider** | Pointer to **string** | Provider of the location | [optional] 

## Methods

### NewLocation

`func NewLocation(lat float64, lon float64, ) *Location`

NewLocation instantiates a new Location object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLocationWithDefaults

`func NewLocationWithDefaults() *Location`

NewLocationWithDefaults instantiates a new Location object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAccuracy

`func (o *Location) GetAccuracy() float64`

GetAccuracy returns the Accuracy field if non-nil, zero value otherwise.

### GetAccuracyOk

`func (o *Location) GetAccuracyOk() (*float64, bool)`

GetAccuracyOk returns a tuple with the Accuracy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccuracy

`func (o *Location) SetAccuracy(v float64)`

SetAccuracy sets Accuracy field to given value.

### HasAccuracy

`func (o *Location) HasAccuracy() bool`

HasAccuracy returns a boolean if a field has been set.

### GetAlt

`func (o *Location) GetAlt() float64`

GetAlt returns the Alt field if non-nil, zero value otherwise.

### GetAltOk

`func (o *Location) GetAltOk() (*float64, bool)`

GetAltOk returns a tuple with the Alt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlt

`func (o *Location) SetAlt(v float64)`

SetAlt sets Alt field to given value.

### HasAlt

`func (o *Location) HasAlt() bool`

HasAlt returns a boolean if a field has been set.

### GetLat

`func (o *Location) GetLat() float64`

GetLat returns the Lat field if non-nil, zero value otherwise.

### GetLatOk

`func (o *Location) GetLatOk() (*float64, bool)`

GetLatOk returns a tuple with the Lat field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLat

`func (o *Location) SetLat(v float64)`

SetLat sets Lat field to given value.


### GetLon

`func (o *Location) GetLon() float64`

GetLon returns the Lon field if non-nil, zero value otherwise.

### GetLonOk

`func (o *Location) GetLonOk() (*float64, bool)`

GetLonOk returns a tuple with the Lon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLon

`func (o *Location) SetLon(v float64)`

SetLon sets Lon field to given value.


### GetProvider

`func (o *Location) GetProvider() string`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *Location) GetProviderOk() (*string, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *Location) SetProvider(v string)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *Location) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


