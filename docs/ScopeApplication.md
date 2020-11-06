# ScopeApplication

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Fifos** | Pointer to **[]string** | List of allowed FIFOs to publish or subscribe. Expected array of string (max 100 elements, value max 255 characters) | [optional] 

## Methods

### NewScopeApplication

`func NewScopeApplication() *ScopeApplication`

NewScopeApplication instantiates a new ScopeApplication object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScopeApplicationWithDefaults

`func NewScopeApplicationWithDefaults() *ScopeApplication`

NewScopeApplicationWithDefaults instantiates a new ScopeApplication object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFifos

`func (o *ScopeApplication) GetFifos() []string`

GetFifos returns the Fifos field if non-nil, zero value otherwise.

### GetFifosOk

`func (o *ScopeApplication) GetFifosOk() (*[]string, bool)`

GetFifosOk returns a tuple with the Fifos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFifos

`func (o *ScopeApplication) SetFifos(v []string)`

SetFifos sets Fifos field to given value.

### HasFifos

`func (o *ScopeApplication) HasFifos() bool`

HasFifos returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


