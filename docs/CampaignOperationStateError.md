# CampaignOperationStateError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Code** | Pointer to **string** | Optional error code associated to the current underlying operation | [optional] 
**Details** | Pointer to **string** | Optional device originated error details associated to the current underlying operation | [optional] 
**DeviceCode** | Pointer to **string** | Optional device originated error code associated to the current underlying operation | [optional] 
**Origin** | Pointer to **string** | Optional error origin associated to the current underlying operation | [optional] 

## Methods

### NewCampaignOperationStateError

`func NewCampaignOperationStateError() *CampaignOperationStateError`

NewCampaignOperationStateError instantiates a new CampaignOperationStateError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCampaignOperationStateErrorWithDefaults

`func NewCampaignOperationStateErrorWithDefaults() *CampaignOperationStateError`

NewCampaignOperationStateErrorWithDefaults instantiates a new CampaignOperationStateError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCode

`func (o *CampaignOperationStateError) GetCode() string`

GetCode returns the Code field if non-nil, zero value otherwise.

### GetCodeOk

`func (o *CampaignOperationStateError) GetCodeOk() (*string, bool)`

GetCodeOk returns a tuple with the Code field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCode

`func (o *CampaignOperationStateError) SetCode(v string)`

SetCode sets Code field to given value.

### HasCode

`func (o *CampaignOperationStateError) HasCode() bool`

HasCode returns a boolean if a field has been set.

### GetDetails

`func (o *CampaignOperationStateError) GetDetails() string`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *CampaignOperationStateError) GetDetailsOk() (*string, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *CampaignOperationStateError) SetDetails(v string)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *CampaignOperationStateError) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetDeviceCode

`func (o *CampaignOperationStateError) GetDeviceCode() string`

GetDeviceCode returns the DeviceCode field if non-nil, zero value otherwise.

### GetDeviceCodeOk

`func (o *CampaignOperationStateError) GetDeviceCodeOk() (*string, bool)`

GetDeviceCodeOk returns a tuple with the DeviceCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceCode

`func (o *CampaignOperationStateError) SetDeviceCode(v string)`

SetDeviceCode sets DeviceCode field to given value.

### HasDeviceCode

`func (o *CampaignOperationStateError) HasDeviceCode() bool`

HasDeviceCode returns a boolean if a field has been set.

### GetOrigin

`func (o *CampaignOperationStateError) GetOrigin() string`

GetOrigin returns the Origin field if non-nil, zero value otherwise.

### GetOriginOk

`func (o *CampaignOperationStateError) GetOriginOk() (*string, bool)`

GetOriginOk returns a tuple with the Origin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrigin

`func (o *CampaignOperationStateError) SetOrigin(v string)`

SetOrigin sets Origin field to given value.

### HasOrigin

`func (o *CampaignOperationStateError) HasOrigin() bool`

HasOrigin returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


