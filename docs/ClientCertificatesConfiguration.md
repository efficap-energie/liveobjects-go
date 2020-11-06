# ClientCertificatesConfiguration

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CaCertIds** | Pointer to **[]string** | List of Ca Certificate Ids used to authenticate devices. Expected array of string (max 100 elements, value max 255 characters) | [optional] 
**Required** | Pointer to **bool** | Indicates if the client must use TLS client cert authentication | [optional] 

## Methods

### NewClientCertificatesConfiguration

`func NewClientCertificatesConfiguration() *ClientCertificatesConfiguration`

NewClientCertificatesConfiguration instantiates a new ClientCertificatesConfiguration object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewClientCertificatesConfigurationWithDefaults

`func NewClientCertificatesConfigurationWithDefaults() *ClientCertificatesConfiguration`

NewClientCertificatesConfigurationWithDefaults instantiates a new ClientCertificatesConfiguration object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCaCertIds

`func (o *ClientCertificatesConfiguration) GetCaCertIds() []string`

GetCaCertIds returns the CaCertIds field if non-nil, zero value otherwise.

### GetCaCertIdsOk

`func (o *ClientCertificatesConfiguration) GetCaCertIdsOk() (*[]string, bool)`

GetCaCertIdsOk returns a tuple with the CaCertIds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCaCertIds

`func (o *ClientCertificatesConfiguration) SetCaCertIds(v []string)`

SetCaCertIds sets CaCertIds field to given value.

### HasCaCertIds

`func (o *ClientCertificatesConfiguration) HasCaCertIds() bool`

HasCaCertIds returns a boolean if a field has been set.

### GetRequired

`func (o *ClientCertificatesConfiguration) GetRequired() bool`

GetRequired returns the Required field if non-nil, zero value otherwise.

### GetRequiredOk

`func (o *ClientCertificatesConfiguration) GetRequiredOk() (*bool, bool)`

GetRequiredOk returns a tuple with the Required field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRequired

`func (o *ClientCertificatesConfiguration) SetRequired(v bool)`

SetRequired sets Required field to given value.

### HasRequired

`func (o *ClientCertificatesConfiguration) HasRequired() bool`

HasRequired returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


