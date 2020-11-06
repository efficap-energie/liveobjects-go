# FirmwareUpdateWeb

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
**NotifyTo** | Pointer to **string** |  | [optional] 
**Progress** | Pointer to **float64** |  | [optional] 
**RequestedVersion** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **string** |  | [optional] 
**Updated** | Pointer to **string** |  | [optional] 

## Methods

### NewFirmwareUpdateWeb

`func NewFirmwareUpdateWeb() *FirmwareUpdateWeb`

NewFirmwareUpdateWeb instantiates a new FirmwareUpdateWeb object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFirmwareUpdateWebWithDefaults

`func NewFirmwareUpdateWebWithDefaults() *FirmwareUpdateWeb`

NewFirmwareUpdateWebWithDefaults instantiates a new FirmwareUpdateWeb object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreated

`func (o *FirmwareUpdateWeb) GetCreated() string`

GetCreated returns the Created field if non-nil, zero value otherwise.

### GetCreatedOk

`func (o *FirmwareUpdateWeb) GetCreatedOk() (*string, bool)`

GetCreatedOk returns a tuple with the Created field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreated

`func (o *FirmwareUpdateWeb) SetCreated(v string)`

SetCreated sets Created field to given value.

### HasCreated

`func (o *FirmwareUpdateWeb) HasCreated() bool`

HasCreated returns a boolean if a field has been set.

### GetDeviceErrorCode

`func (o *FirmwareUpdateWeb) GetDeviceErrorCode() string`

GetDeviceErrorCode returns the DeviceErrorCode field if non-nil, zero value otherwise.

### GetDeviceErrorCodeOk

`func (o *FirmwareUpdateWeb) GetDeviceErrorCodeOk() (*string, bool)`

GetDeviceErrorCodeOk returns a tuple with the DeviceErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceErrorCode

`func (o *FirmwareUpdateWeb) SetDeviceErrorCode(v string)`

SetDeviceErrorCode sets DeviceErrorCode field to given value.

### HasDeviceErrorCode

`func (o *FirmwareUpdateWeb) HasDeviceErrorCode() bool`

HasDeviceErrorCode returns a boolean if a field has been set.

### GetErrorCode

`func (o *FirmwareUpdateWeb) GetErrorCode() string`

GetErrorCode returns the ErrorCode field if non-nil, zero value otherwise.

### GetErrorCodeOk

`func (o *FirmwareUpdateWeb) GetErrorCodeOk() (*string, bool)`

GetErrorCodeOk returns a tuple with the ErrorCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCode

`func (o *FirmwareUpdateWeb) SetErrorCode(v string)`

SetErrorCode sets ErrorCode field to given value.

### HasErrorCode

`func (o *FirmwareUpdateWeb) HasErrorCode() bool`

HasErrorCode returns a boolean if a field has been set.

### GetErrorDetails

`func (o *FirmwareUpdateWeb) GetErrorDetails() string`

GetErrorDetails returns the ErrorDetails field if non-nil, zero value otherwise.

### GetErrorDetailsOk

`func (o *FirmwareUpdateWeb) GetErrorDetailsOk() (*string, bool)`

GetErrorDetailsOk returns a tuple with the ErrorDetails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorDetails

`func (o *FirmwareUpdateWeb) SetErrorDetails(v string)`

SetErrorDetails sets ErrorDetails field to given value.

### HasErrorDetails

`func (o *FirmwareUpdateWeb) HasErrorDetails() bool`

HasErrorDetails returns a boolean if a field has been set.

### GetErrorOrigin

`func (o *FirmwareUpdateWeb) GetErrorOrigin() string`

GetErrorOrigin returns the ErrorOrigin field if non-nil, zero value otherwise.

### GetErrorOriginOk

`func (o *FirmwareUpdateWeb) GetErrorOriginOk() (*string, bool)`

GetErrorOriginOk returns a tuple with the ErrorOrigin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorOrigin

`func (o *FirmwareUpdateWeb) SetErrorOrigin(v string)`

SetErrorOrigin sets ErrorOrigin field to given value.

### HasErrorOrigin

`func (o *FirmwareUpdateWeb) HasErrorOrigin() bool`

HasErrorOrigin returns a boolean if a field has been set.

### GetId

`func (o *FirmwareUpdateWeb) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FirmwareUpdateWeb) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FirmwareUpdateWeb) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *FirmwareUpdateWeb) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitialVersion

`func (o *FirmwareUpdateWeb) GetInitialVersion() string`

GetInitialVersion returns the InitialVersion field if non-nil, zero value otherwise.

### GetInitialVersionOk

`func (o *FirmwareUpdateWeb) GetInitialVersionOk() (*string, bool)`

GetInitialVersionOk returns a tuple with the InitialVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitialVersion

`func (o *FirmwareUpdateWeb) SetInitialVersion(v string)`

SetInitialVersion sets InitialVersion field to given value.

### HasInitialVersion

`func (o *FirmwareUpdateWeb) HasInitialVersion() bool`

HasInitialVersion returns a boolean if a field has been set.

### GetMetadata

`func (o *FirmwareUpdateWeb) GetMetadata() map[string]string`

GetMetadata returns the Metadata field if non-nil, zero value otherwise.

### GetMetadataOk

`func (o *FirmwareUpdateWeb) GetMetadataOk() (*map[string]string, bool)`

GetMetadataOk returns a tuple with the Metadata field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetadata

`func (o *FirmwareUpdateWeb) SetMetadata(v map[string]string)`

SetMetadata sets Metadata field to given value.

### HasMetadata

`func (o *FirmwareUpdateWeb) HasMetadata() bool`

HasMetadata returns a boolean if a field has been set.

### GetNotifyTo

`func (o *FirmwareUpdateWeb) GetNotifyTo() string`

GetNotifyTo returns the NotifyTo field if non-nil, zero value otherwise.

### GetNotifyToOk

`func (o *FirmwareUpdateWeb) GetNotifyToOk() (*string, bool)`

GetNotifyToOk returns a tuple with the NotifyTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotifyTo

`func (o *FirmwareUpdateWeb) SetNotifyTo(v string)`

SetNotifyTo sets NotifyTo field to given value.

### HasNotifyTo

`func (o *FirmwareUpdateWeb) HasNotifyTo() bool`

HasNotifyTo returns a boolean if a field has been set.

### GetProgress

`func (o *FirmwareUpdateWeb) GetProgress() float64`

GetProgress returns the Progress field if non-nil, zero value otherwise.

### GetProgressOk

`func (o *FirmwareUpdateWeb) GetProgressOk() (*float64, bool)`

GetProgressOk returns a tuple with the Progress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProgress

`func (o *FirmwareUpdateWeb) SetProgress(v float64)`

SetProgress sets Progress field to given value.

### HasProgress

`func (o *FirmwareUpdateWeb) HasProgress() bool`

HasProgress returns a boolean if a field has been set.

### GetRequestedVersion

`func (o *FirmwareUpdateWeb) GetRequestedVersion() string`

GetRequestedVersion returns the RequestedVersion field if non-nil, zero value otherwise.

### GetRequestedVersionOk

`func (o *FirmwareUpdateWeb) GetRequestedVersionOk() (*string, bool)`

GetRequestedVersionOk returns a tuple with the RequestedVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequestedVersion

`func (o *FirmwareUpdateWeb) SetRequestedVersion(v string)`

SetRequestedVersion sets RequestedVersion field to given value.

### HasRequestedVersion

`func (o *FirmwareUpdateWeb) HasRequestedVersion() bool`

HasRequestedVersion returns a boolean if a field has been set.

### GetStatus

`func (o *FirmwareUpdateWeb) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *FirmwareUpdateWeb) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *FirmwareUpdateWeb) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *FirmwareUpdateWeb) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUpdated

`func (o *FirmwareUpdateWeb) GetUpdated() string`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *FirmwareUpdateWeb) GetUpdatedOk() (*string, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *FirmwareUpdateWeb) SetUpdated(v string)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *FirmwareUpdateWeb) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


