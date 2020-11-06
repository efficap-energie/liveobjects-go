# OfferAndOptions

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Offer** | Pointer to [**Rule**](Rule.md) |  | [optional] 
**Options** | Pointer to [**[]Rule**](Rule.md) | displayName | [optional] 

## Methods

### NewOfferAndOptions

`func NewOfferAndOptions() *OfferAndOptions`

NewOfferAndOptions instantiates a new OfferAndOptions object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOfferAndOptionsWithDefaults

`func NewOfferAndOptionsWithDefaults() *OfferAndOptions`

NewOfferAndOptionsWithDefaults instantiates a new OfferAndOptions object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOffer

`func (o *OfferAndOptions) GetOffer() Rule`

GetOffer returns the Offer field if non-nil, zero value otherwise.

### GetOfferOk

`func (o *OfferAndOptions) GetOfferOk() (*Rule, bool)`

GetOfferOk returns a tuple with the Offer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOffer

`func (o *OfferAndOptions) SetOffer(v Rule)`

SetOffer sets Offer field to given value.

### HasOffer

`func (o *OfferAndOptions) HasOffer() bool`

HasOffer returns a boolean if a field has been set.

### GetOptions

`func (o *OfferAndOptions) GetOptions() []Rule`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *OfferAndOptions) GetOptionsOk() (*[]Rule, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *OfferAndOptions) SetOptions(v []Rule)`

SetOptions sets Options field to given value.

### HasOptions

`func (o *OfferAndOptions) HasOptions() bool`

HasOptions returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


