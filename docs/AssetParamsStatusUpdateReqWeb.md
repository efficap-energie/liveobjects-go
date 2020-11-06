# AssetParamsStatusUpdateReqWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewStatus** | **string** | future state of the parameter --&gt; CANCELED | 
**ParamKeys** | **[]string** | set of keys identifying the targeted asset parameters. Maximum 100 parameters, and parameter name max length is 128 | 

## Methods

### NewAssetParamsStatusUpdateReqWeb

`func NewAssetParamsStatusUpdateReqWeb(newStatus string, paramKeys []string, ) *AssetParamsStatusUpdateReqWeb`

NewAssetParamsStatusUpdateReqWeb instantiates a new AssetParamsStatusUpdateReqWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAssetParamsStatusUpdateReqWebWithDefaults

`func NewAssetParamsStatusUpdateReqWebWithDefaults() *AssetParamsStatusUpdateReqWeb`

NewAssetParamsStatusUpdateReqWebWithDefaults instantiates a new AssetParamsStatusUpdateReqWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewStatus

`func (o *AssetParamsStatusUpdateReqWeb) GetNewStatus() string`

GetNewStatus returns the NewStatus field if non-nil, zero value otherwise.

### GetNewStatusOk

`func (o *AssetParamsStatusUpdateReqWeb) GetNewStatusOk() (*string, bool)`

GetNewStatusOk returns a tuple with the NewStatus field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewStatus

`func (o *AssetParamsStatusUpdateReqWeb) SetNewStatus(v string)`

SetNewStatus sets NewStatus field to given value.


### GetParamKeys

`func (o *AssetParamsStatusUpdateReqWeb) GetParamKeys() []string`

GetParamKeys returns the ParamKeys field if non-nil, zero value otherwise.

### GetParamKeysOk

`func (o *AssetParamsStatusUpdateReqWeb) GetParamKeysOk() (*[]string, bool)`

GetParamKeysOk returns a tuple with the ParamKeys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParamKeys

`func (o *AssetParamsStatusUpdateReqWeb) SetParamKeys(v []string)`

SetParamKeys sets ParamKeys field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


