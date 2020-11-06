# SelectionCriteria

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**KeyPath** | **string** | json path of field to select | 
**Value** | **map[string]interface{}** | value to match | 

## Methods

### NewSelectionCriteria

`func NewSelectionCriteria(keyPath string, value map[string]interface{}, ) *SelectionCriteria`

NewSelectionCriteria instantiates a new SelectionCriteria object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSelectionCriteriaWithDefaults

`func NewSelectionCriteriaWithDefaults() *SelectionCriteria`

NewSelectionCriteriaWithDefaults instantiates a new SelectionCriteria object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetKeyPath

`func (o *SelectionCriteria) GetKeyPath() string`

GetKeyPath returns the KeyPath field if non-nil, zero value otherwise.

### GetKeyPathOk

`func (o *SelectionCriteria) GetKeyPathOk() (*string, bool)`

GetKeyPathOk returns a tuple with the KeyPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeyPath

`func (o *SelectionCriteria) SetKeyPath(v string)`

SetKeyPath sets KeyPath field to given value.


### GetValue

`func (o *SelectionCriteria) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *SelectionCriteria) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *SelectionCriteria) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


