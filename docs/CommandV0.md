# CommandV0

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AssetId** | Pointer to **string** |  | [optional] 
**AssetIdNamespace** | Pointer to **string** |  | [optional] 
**CreationTs** | Pointer to **int64** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**NotifyTo** | Pointer to **string** |  | [optional] 
**ReqData** | Pointer to **map[string]string** |  | [optional] 
**ReqEvent** | Pointer to **string** |  | [optional] 
**ReqPayload** | Pointer to **[]string** |  | [optional] 
**ResData** | Pointer to **map[string]string** |  | [optional] 
**ResPayload** | Pointer to **[]string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**UpdateTs** | Pointer to **int64** |  | [optional] 

## Methods

### NewCommandV0

`func NewCommandV0() *CommandV0`

NewCommandV0 instantiates a new CommandV0 object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCommandV0WithDefaults

`func NewCommandV0WithDefaults() *CommandV0`

NewCommandV0WithDefaults instantiates a new CommandV0 object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAssetId

`func (o *CommandV0) GetAssetId() string`

GetAssetId returns the AssetId field if non-nil, zero value otherwise.

### GetAssetIdOk

`func (o *CommandV0) GetAssetIdOk() (*string, bool)`

GetAssetIdOk returns a tuple with the AssetId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetId

`func (o *CommandV0) SetAssetId(v string)`

SetAssetId sets AssetId field to given value.

### HasAssetId

`func (o *CommandV0) HasAssetId() bool`

HasAssetId returns a boolean if a field has been set.

### GetAssetIdNamespace

`func (o *CommandV0) GetAssetIdNamespace() string`

GetAssetIdNamespace returns the AssetIdNamespace field if non-nil, zero value otherwise.

### GetAssetIdNamespaceOk

`func (o *CommandV0) GetAssetIdNamespaceOk() (*string, bool)`

GetAssetIdNamespaceOk returns a tuple with the AssetIdNamespace field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAssetIdNamespace

`func (o *CommandV0) SetAssetIdNamespace(v string)`

SetAssetIdNamespace sets AssetIdNamespace field to given value.

### HasAssetIdNamespace

`func (o *CommandV0) HasAssetIdNamespace() bool`

HasAssetIdNamespace returns a boolean if a field has been set.

### GetCreationTs

`func (o *CommandV0) GetCreationTs() int64`

GetCreationTs returns the CreationTs field if non-nil, zero value otherwise.

### GetCreationTsOk

`func (o *CommandV0) GetCreationTsOk() (*int64, bool)`

GetCreationTsOk returns a tuple with the CreationTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTs

`func (o *CommandV0) SetCreationTs(v int64)`

SetCreationTs sets CreationTs field to given value.

### HasCreationTs

`func (o *CommandV0) HasCreationTs() bool`

HasCreationTs returns a boolean if a field has been set.

### GetId

`func (o *CommandV0) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CommandV0) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CommandV0) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *CommandV0) HasId() bool`

HasId returns a boolean if a field has been set.

### GetNotifyTo

`func (o *CommandV0) GetNotifyTo() string`

GetNotifyTo returns the NotifyTo field if non-nil, zero value otherwise.

### GetNotifyToOk

`func (o *CommandV0) GetNotifyToOk() (*string, bool)`

GetNotifyToOk returns a tuple with the NotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTo

`func (o *CommandV0) SetNotifyTo(v string)`

SetNotifyTo sets NotifyTo field to given value.

### HasNotifyTo

`func (o *CommandV0) HasNotifyTo() bool`

HasNotifyTo returns a boolean if a field has been set.

### GetReqData

`func (o *CommandV0) GetReqData() map[string]string`

GetReqData returns the ReqData field if non-nil, zero value otherwise.

### GetReqDataOk

`func (o *CommandV0) GetReqDataOk() (*map[string]string, bool)`

GetReqDataOk returns a tuple with the ReqData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqData

`func (o *CommandV0) SetReqData(v map[string]string)`

SetReqData sets ReqData field to given value.

### HasReqData

`func (o *CommandV0) HasReqData() bool`

HasReqData returns a boolean if a field has been set.

### GetReqEvent

`func (o *CommandV0) GetReqEvent() string`

GetReqEvent returns the ReqEvent field if non-nil, zero value otherwise.

### GetReqEventOk

`func (o *CommandV0) GetReqEventOk() (*string, bool)`

GetReqEventOk returns a tuple with the ReqEvent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqEvent

`func (o *CommandV0) SetReqEvent(v string)`

SetReqEvent sets ReqEvent field to given value.

### HasReqEvent

`func (o *CommandV0) HasReqEvent() bool`

HasReqEvent returns a boolean if a field has been set.

### GetReqPayload

`func (o *CommandV0) GetReqPayload() []string`

GetReqPayload returns the ReqPayload field if non-nil, zero value otherwise.

### GetReqPayloadOk

`func (o *CommandV0) GetReqPayloadOk() (*[]string, bool)`

GetReqPayloadOk returns a tuple with the ReqPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReqPayload

`func (o *CommandV0) SetReqPayload(v []string)`

SetReqPayload sets ReqPayload field to given value.

### HasReqPayload

`func (o *CommandV0) HasReqPayload() bool`

HasReqPayload returns a boolean if a field has been set.

### GetResData

`func (o *CommandV0) GetResData() map[string]string`

GetResData returns the ResData field if non-nil, zero value otherwise.

### GetResDataOk

`func (o *CommandV0) GetResDataOk() (*map[string]string, bool)`

GetResDataOk returns a tuple with the ResData field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResData

`func (o *CommandV0) SetResData(v map[string]string)`

SetResData sets ResData field to given value.

### HasResData

`func (o *CommandV0) HasResData() bool`

HasResData returns a boolean if a field has been set.

### GetResPayload

`func (o *CommandV0) GetResPayload() []string`

GetResPayload returns the ResPayload field if non-nil, zero value otherwise.

### GetResPayloadOk

`func (o *CommandV0) GetResPayloadOk() (*[]string, bool)`

GetResPayloadOk returns a tuple with the ResPayload field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetResPayload

`func (o *CommandV0) SetResPayload(v []string)`

SetResPayload sets ResPayload field to given value.

### HasResPayload

`func (o *CommandV0) HasResPayload() bool`

HasResPayload returns a boolean if a field has been set.

### GetStatus

`func (o *CommandV0) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *CommandV0) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *CommandV0) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *CommandV0) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdateTs

`func (o *CommandV0) GetUpdateTs() int64`

GetUpdateTs returns the UpdateTs field if non-nil, zero value otherwise.

### GetUpdateTsOk

`func (o *CommandV0) GetUpdateTsOk() (*int64, bool)`

GetUpdateTsOk returns a tuple with the UpdateTs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateTs

`func (o *CommandV0) SetUpdateTs(v int64)`

SetUpdateTs sets UpdateTs field to given value.

### HasUpdateTs

`func (o *CommandV0) HasUpdateTs() bool`

HasUpdateTs returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


