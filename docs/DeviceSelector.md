# DeviceSelector

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FilterQuery** | Pointer to **string** | Filtering expression using RSQL notation. Supported device properties are &#39;groupPath&#39;, &#39;groupId&#39;, &#39;tags&#39;, &#39;properties&#39;. Supported RSQL operators are &#39;&#x3D;&#x3D;&#39;,&#39;!&#x3D;&#39;,&#39;&#x3D;in&#x3D;&#39;,&#39;&#x3D;out&#x3D;&#39;,&#39;&#x3D;re&#x3D;&#39;,&#39;&#x3D;lt&#x3D;&#39;,&#39;&#x3D;le&#x3D;&#39;,&#39;&#x3D;gt&#x3D;&#39;,&#39;&#x3D;ge&#x3D;&#39;,&#39;and&#39;,&#39;or&#39; | [optional] 
**IdList** | Pointer to **[]string** | List of device IDs in the form urn:lo:nsid:${assetNamespace}:${assetId} | [optional] 

## Methods

### NewDeviceSelector

`func NewDeviceSelector() *DeviceSelector`

NewDeviceSelector instantiates a new DeviceSelector object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDeviceSelectorWithDefaults

`func NewDeviceSelectorWithDefaults() *DeviceSelector`

NewDeviceSelectorWithDefaults instantiates a new DeviceSelector object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFilterQuery

`func (o *DeviceSelector) GetFilterQuery() string`

GetFilterQuery returns the FilterQuery field if non-nil, zero value otherwise.

### GetFilterQueryOk

`func (o *DeviceSelector) GetFilterQueryOk() (*string, bool)`

GetFilterQueryOk returns a tuple with the FilterQuery field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilterQuery

`func (o *DeviceSelector) SetFilterQuery(v string)`

SetFilterQuery sets FilterQuery field to given value.

### HasFilterQuery

`func (o *DeviceSelector) HasFilterQuery() bool`

HasFilterQuery returns a boolean if a field has been set.

### GetIdList

`func (o *DeviceSelector) GetIdList() []string`

GetIdList returns the IdList field if non-nil, zero value otherwise.

### GetIdListOk

`func (o *DeviceSelector) GetIdListOk() (*[]string, bool)`

GetIdListOk returns a tuple with the IdList field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIdList

`func (o *DeviceSelector) SetIdList(v []string)`

SetIdList sets IdList field to given value.

### HasIdList

`func (o *DeviceSelector) HasIdList() bool`

HasIdList returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


