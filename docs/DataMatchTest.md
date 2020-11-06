# DataMatchTest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Data** | **map[string]interface{}** | data to test the dataPredicate with | 
**DataPredicate** | **map[string]interface{}** | the jsonLogic that will be used to evaluate the data | 

## Methods

### NewDataMatchTest

`func NewDataMatchTest(data map[string]interface{}, dataPredicate map[string]interface{}, ) *DataMatchTest`

NewDataMatchTest instantiates a new DataMatchTest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDataMatchTestWithDefaults

`func NewDataMatchTestWithDefaults() *DataMatchTest`

NewDataMatchTestWithDefaults instantiates a new DataMatchTest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetData

`func (o *DataMatchTest) GetData() map[string]interface{}`

GetData returns the Data field if non-nil, zero value otherwise.

### GetDataOk

`func (o *DataMatchTest) GetDataOk() (*map[string]interface{}, bool)`

GetDataOk returns a tuple with the Data field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetData

`func (o *DataMatchTest) SetData(v map[string]interface{})`

SetData sets Data field to given value.


### GetDataPredicate

`func (o *DataMatchTest) GetDataPredicate() map[string]interface{}`

GetDataPredicate returns the DataPredicate field if non-nil, zero value otherwise.

### GetDataPredicateOk

`func (o *DataMatchTest) GetDataPredicateOk() (*map[string]interface{}, bool)`

GetDataPredicateOk returns a tuple with the DataPredicate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDataPredicate

`func (o *DataMatchTest) SetDataPredicate(v map[string]interface{})`

SetDataPredicate sets DataPredicate field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


