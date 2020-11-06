# CampaignOperation

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Action** | **string** | Type of operation | 
**Version** | Pointer to **int32** | Version of operation (default:0) | [optional] 

## Methods

### NewCampaignOperation

`func NewCampaignOperation(action string, ) *CampaignOperation`

NewCampaignOperation instantiates a new CampaignOperation object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignOperationWithDefaults

`func NewCampaignOperationWithDefaults() *CampaignOperation`

NewCampaignOperationWithDefaults instantiates a new CampaignOperation object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAction

`func (o *CampaignOperation) GetAction() string`

GetAction returns the Action field if non-nil, zero value otherwise.

### GetActionOk

`func (o *CampaignOperation) GetActionOk() (*string, bool)`

GetActionOk returns a tuple with the Action field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAction

`func (o *CampaignOperation) SetAction(v string)`

SetAction sets Action field to given value.


### GetVersion

`func (o *CampaignOperation) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *CampaignOperation) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *CampaignOperation) SetVersion(v int32)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *CampaignOperation) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


