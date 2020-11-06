# ResourceUpdateWeb

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Created** | Pointer to **string** |  | [optional] 
**DeviceErrorCode** | Pointer to **string** |  | [optional] 
**ErrorCode** | Pointer to **string** |  | [optional] 
**ErrorDetails** | Pointer to **string** |  | [optional] 
**ErrorOrigin** | Pointer to **string** |  | [optional] 
**Id** | Pointer to **string** |  | [optional] 
**InitialVersion** | Pointer to **string** |  | [optional] 
**Metadata** | Pointer to **map[string]string** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**RequestedVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewResourceUpdateWeb

`func NewResourceUpdateWeb() *ResourceUpdateWeb`

NewResourceUpdateWeb instantiates a new ResourceUpdateWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResourceUpdateWebWithDefaults

`func NewResourceUpdateWebWithDefaults() *ResourceUpdateWeb`

NewResourceUpdateWebWithDefaults instantiates a new ResourceUpdateWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *ResourceUpdateWeb) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *ResourceUpdateWeb) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *ResourceUpdateWeb) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *ResourceUpdateWeb) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeviceErrorCode

`func (o *ResourceUpdateWeb) GetDeviceErrorCode() string`

GetDeviceErrorCode returns the DeviceErrorCode field if non-nil, zero value otherwise.

### GetDeviceErrorCodeOk

`func (o *ResourceUpdateWeb) GetDeviceErrorCodeOk() (*string, bool)`

GetDeviceErrorCodeOk returns a tuple with the DeviceErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceErrorCode

`func (o *ResourceUpdateWeb) SetDeviceErrorCode(v string)`

SetDeviceErrorCode sets DeviceErrorCode field to given value.

### HasDeviceErrorCode

`func (o *ResourceUpdateWeb) HasDeviceErrorCode() bool`

HasDeviceErrorCode returns a boolean if a field has been set.

### GetErrorCode

`func (o *ResourceUpdateWeb) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *ResourceUpdateWeb) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *ResourceUpdateWeb) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *ResourceUpdateWeb) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetails

`func (o *ResourceUpdateWeb) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *ResourceUpdateWeb) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *ResourceUpdateWeb) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *ResourceUpdateWeb) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetErrorOrigin

`func (o *ResourceUpdateWeb) GetErrorOrigin() string`

GetErrorOrigin returns the ErrorOrigin field if non-nil, zero value otherwise.

### GetErrorOriginOk

`func (o *ResourceUpdateWeb) GetErrorOriginOk() (*string, bool)`

GetErrorOriginOk returns a tuple with the ErrorOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOrigin

`func (o *ResourceUpdateWeb) SetErrorOrigin(v string)`

SetErrorOrigin sets ErrorOrigin field to given value.

### HasErrorOrigin

`func (o *ResourceUpdateWeb) HasErrorOrigin() bool`

HasErrorOrigin returns a boolean if a field has been set.

### GetId

`func (o *ResourceUpdateWeb) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ResourceUpdateWeb) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ResourceUpdateWeb) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *ResourceUpdateWeb) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialVersion

`func (o *ResourceUpdateWeb) GetInitialVersion() string`

GetInitialVersion returns the InitialVersion field if non-nil, zero value otherwise.

### GetInitialVersionOk

`func (o *ResourceUpdateWeb) GetInitialVersionOk() (*string, bool)`

GetInitialVersionOk returns a tuple with the InitialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialVersion

`func (o *ResourceUpdateWeb) SetInitialVersion(v string)`

SetInitialVersion sets InitialVersion field to given value.

### HasInitialVersion

`func (o *ResourceUpdateWeb) HasInitialVersion() bool`

HasInitialVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *ResourceUpdateWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *ResourceUpdateWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *ResourceUpdateWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *ResourceUpdateWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetProgress

`func (o *ResourceUpdateWeb) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *ResourceUpdateWeb) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *ResourceUpdateWeb) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *ResourceUpdateWeb) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRequestedVersion

`func (o *ResourceUpdateWeb) GetRequestedVersion() string`

GetRequestedVersion returns the RequestedVersion field if non-nil, zero value otherwise.

### GetRequestedVersionOk

`func (o *ResourceUpdateWeb) GetRequestedVersionOk() (*string, bool)`

GetRequestedVersionOk returns a tuple with the RequestedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedVersion

`func (o *ResourceUpdateWeb) SetRequestedVersion(v string)`

SetRequestedVersion sets RequestedVersion field to given value.

### HasRequestedVersion

`func (o *ResourceUpdateWeb) HasRequestedVersion() bool`

HasRequestedVersion returns a boolean if a field has been set.

### GetStatus

`func (o *ResourceUpdateWeb) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ResourceUpdateWeb) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ResourceUpdateWeb) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ResourceUpdateWeb) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *ResourceUpdateWeb) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *ResourceUpdateWeb) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *ResourceUpdateWeb) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *ResourceUpdateWeb) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


