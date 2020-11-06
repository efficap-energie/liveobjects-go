# StateProcessingFunctionTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentState** | Pointer to **map[string]interface{}** | current state that can be used by the state function | [optional] 
**Data** | [**NewData**](NewData.md) |  | 
**StateProcessingFunction** | **map[string]interface{}** | the state function in jsonLogic syntax | 

## Methods

### NewStateProcessingFunctionTest

`func NewStateProcessingFunctionTest(data NewData, stateProcessingFunction map[string]interface{}, ) *StateProcessingFunctionTest`

NewStateProcessingFunctionTest instantiates a new StateProcessingFunctionTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStateProcessingFunctionTestWithDefaults

`func NewStateProcessingFunctionTestWithDefaults() *StateProcessingFunctionTest`

NewStateProcessingFunctionTestWithDefaults instantiates a new StateProcessingFunctionTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentState

`func (o *StateProcessingFunctionTest) GetCurrentState() map[string]interface{}`

GetCurrentState returns the CurrentState field if non-nil, zero value otherwise.

### GetCurrentStateOk

`func (o *StateProcessingFunctionTest) GetCurrentStateOk() (*map[string]interface{}, bool)`

GetCurrentStateOk returns a tuple with the CurrentState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentState

`func (o *StateProcessingFunctionTest) SetCurrentState(v map[string]interface{})`

SetCurrentState sets CurrentState field to given value.

### HasCurrentState

`func (o *StateProcessingFunctionTest) HasCurrentState() bool`

HasCurrentState returns a boolean if a field has been set.

### GetData

`func (o *StateProcessingFunctionTest) GetData() NewData`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *StateProcessingFunctionTest) GetDataOk() (*NewData, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *StateProcessingFunctionTest) SetData(v NewData)`

SetData sets Data field to given value.


### GetStateProcessingFunction

`func (o *StateProcessingFunctionTest) GetStateProcessingFunction() map[string]interface{}`

GetStateProcessingFunction returns the StateProcessingFunction field if non-nil, zero value otherwise.

### GetStateProcessingFunctionOk

`func (o *StateProcessingFunctionTest) GetStateProcessingFunctionOk() (*map[string]interface{}, bool)`

GetStateProcessingFunctionOk returns a tuple with the StateProcessingFunction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStateProcessingFunction

`func (o *StateProcessingFunctionTest) SetStateProcessingFunction(v map[string]interface{})`

SetStateProcessingFunction sets StateProcessingFunction field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


