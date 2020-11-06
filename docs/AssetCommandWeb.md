# AssetCommandWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | Pointer to **map[string]string** | Data fields / arguments of the command sent to the device. Expected map of string,string (max 100 elements, key max 255 characters, value is limited) | [optional] 
**Event** | Pointer to **string** | &#39;event&#39; of the command message to send to the device (usually used to indicate called function). The length is limited | [optional] 
**Payload** | Pointer to **string** | payload of the command message to send to the device. The length is limited | [optional] 

## Methods

### NewAssetCommandWeb

`func NewAssetCommandWeb() *AssetCommandWeb`

NewAssetCommandWeb instantiates a new AssetCommandWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetCommandWebWithDefaults

`func NewAssetCommandWebWithDefaults() *AssetCommandWeb`

NewAssetCommandWebWithDefaults instantiates a new AssetCommandWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *AssetCommandWeb) GetData() map[string]string`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *AssetCommandWeb) GetDataOk() (*map[string]string, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *AssetCommandWeb) SetData(v map[string]string)`

SetData sets Data field to given value.

### HasData

`func (o *AssetCommandWeb) HasData() bool`

HasData returns a boolean if a field has been set.

### GetEvent

`func (o *AssetCommandWeb) GetEvent() string`

GetEvent returns the Event field if non-nil, zero value otherwise.

### GetEventOk

`func (o *AssetCommandWeb) GetEventOk() (*string, bool)`

GetEventOk returns a tuple with the Event field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEvent

`func (o *AssetCommandWeb) SetEvent(v string)`

SetEvent sets Event field to given value.

### HasEvent

`func (o *AssetCommandWeb) HasEvent() bool`

HasEvent returns a boolean if a field has been set.

### GetPayload

`func (o *AssetCommandWeb) GetPayload() string`

GetPayload returns the Payload field if non-nil, zero value otherwise.

### GetPayloadOk

`func (o *AssetCommandWeb) GetPayloadOk() (*string, bool)`

GetPayloadOk returns a tuple with the Payload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPayload

`func (o *AssetCommandWeb) SetPayload(v string)`

SetPayload sets Payload field to given value.

### HasPayload

`func (o *AssetCommandWeb) HasPayload() bool`

HasPayload returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


