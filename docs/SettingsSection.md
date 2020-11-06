# SettingsSection

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Content** | Pointer to **map[string]interface{}** |  | [optional] 
**Revision** | Pointer to **int32** |  | [optional] 
**Version** | Pointer to **string** |  | [optional] 

## Methods

### NewSettingsSection

`func NewSettingsSection() *SettingsSection`

NewSettingsSection instantiates a new SettingsSection object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSettingsSectionWithDefaults

`func NewSettingsSectionWithDefaults() *SettingsSection`

NewSettingsSectionWithDefaults instantiates a new SettingsSection object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContent

`func (o *SettingsSection) GetContent() map[string]interface{}`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *SettingsSection) GetContentOk() (*map[string]interface{}, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *SettingsSection) SetContent(v map[string]interface{})`

SetContent sets Content field to given value.

### HasContent

`func (o *SettingsSection) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetRevision

`func (o *SettingsSection) GetRevision() int32`

GetRevision returns the Revision field if non-nil, zero value otherwise.

### GetRevisionOk

`func (o *SettingsSection) GetRevisionOk() (*int32, bool)`

GetRevisionOk returns a tuple with the Revision field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRevision

`func (o *SettingsSection) SetRevision(v int32)`

SetRevision sets Revision field to given value.

### HasRevision

`func (o *SettingsSection) HasRevision() bool`

HasRevision returns a boolean if a field has been set.

### GetVersion

`func (o *SettingsSection) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SettingsSection) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SettingsSection) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *SettingsSection) HasVersion() bool`

HasVersion returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


