# PayloadDescriptionMetadata

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Certified** | Pointer to **bool** | the decoder is certified for the specified encoding. Only a public decoder can be certified. The default value is false. | [optional] 
**CertifiedSince** | Pointer to **string** | date since when the decoder is certified. ISO 8601 date time. | [optional] [readonly] 

## Methods

### NewPayloadDescriptionMetadata

`func NewPayloadDescriptionMetadata() *PayloadDescriptionMetadata`

NewPayloadDescriptionMetadata instantiates a new PayloadDescriptionMetadata object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPayloadDescriptionMetadataWithDefaults

`func NewPayloadDescriptionMetadataWithDefaults() *PayloadDescriptionMetadata`

NewPayloadDescriptionMetadataWithDefaults instantiates a new PayloadDescriptionMetadata object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCertified

`func (o *PayloadDescriptionMetadata) GetCertified() bool`

GetCertified returns the Certified field if non-nil, zero value otherwise.

### GetCertifiedOk

`func (o *PayloadDescriptionMetadata) GetCertifiedOk() (*bool, bool)`

GetCertifiedOk returns a tuple with the Certified field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertified

`func (o *PayloadDescriptionMetadata) SetCertified(v bool)`

SetCertified sets Certified field to given value.

### HasCertified

`func (o *PayloadDescriptionMetadata) HasCertified() bool`

HasCertified returns a boolean if a field has been set.

### GetCertifiedSince

`func (o *PayloadDescriptionMetadata) GetCertifiedSince() string`

GetCertifiedSince returns the CertifiedSince field if non-nil, zero value otherwise.

### GetCertifiedSinceOk

`func (o *PayloadDescriptionMetadata) GetCertifiedSinceOk() (*string, bool)`

GetCertifiedSinceOk returns a tuple with the CertifiedSince field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCertifiedSince

`func (o *PayloadDescriptionMetadata) SetCertifiedSince(v string)`

SetCertifiedSince sets CertifiedSince field to given value.

### HasCertifiedSince

`func (o *PayloadDescriptionMetadata) HasCertifiedSince() bool`

HasCertifiedSince returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


