# ResponseEntity

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Body** | Pointer to **map[string]interface{}** |  | [optional] 
**StatusCode** | Pointer to **string** |  | [optional] 
**StatusCodeValue** | Pointer to **int32** |  | [optional] 

## Methods

### NewResponseEntity

`func NewResponseEntity() *ResponseEntity`

NewResponseEntity instantiates a new ResponseEntity object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewResponseEntityWithDefaults

`func NewResponseEntityWithDefaults() *ResponseEntity`

NewResponseEntityWithDefaults instantiates a new ResponseEntity object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBody

`func (o *ResponseEntity) GetBody() map[string]interface{}`

GetBody returns the Body field if non-nil, zero value otherwise.

### GetBodyOk

`func (o *ResponseEntity) GetBodyOk() (*map[string]interface{}, bool)`

GetBodyOk returns a tuple with the Body field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBody

`func (o *ResponseEntity) SetBody(v map[string]interface{})`

SetBody sets Body field to given value.

### HasBody

`func (o *ResponseEntity) HasBody() bool`

HasBody returns a boolean if a field has been set.

### GetStatusCode

`func (o *ResponseEntity) GetStatusCode() string`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *ResponseEntity) GetStatusCodeOk() (*string, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *ResponseEntity) SetStatusCode(v string)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *ResponseEntity) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetStatusCodeValue

`func (o *ResponseEntity) GetStatusCodeValue() int32`

GetStatusCodeValue returns the StatusCodeValue field if non-nil, zero value otherwise.

### GetStatusCodeValueOk

`func (o *ResponseEntity) GetStatusCodeValueOk() (*int32, bool)`

GetStatusCodeValueOk returns a tuple with the StatusCodeValue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCodeValue

`func (o *ResponseEntity) SetStatusCodeValue(v int32)`

SetStatusCodeValue sets StatusCodeValue field to given value.

### HasStatusCodeValue

`func (o *ResponseEntity) HasStatusCodeValue() bool`

HasStatusCodeValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


