# GroupPath

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**IncludeSubPath** | **bool** | if true, all sub-paths will be targeted by this ActivityRule  | 
**Path** | **string** | path of the group | 

## Methods

### NewGroupPath

`func NewGroupPath(includeSubPath bool, path string, ) *GroupPath`

NewGroupPath instantiates a new GroupPath object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupPathWithDefaults

`func NewGroupPathWithDefaults() *GroupPath`

NewGroupPathWithDefaults instantiates a new GroupPath object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncludeSubPath

`func (o *GroupPath) GetIncludeSubPath() bool`

GetIncludeSubPath returns the IncludeSubPath field if non-nil, zero value otherwise.

### GetIncludeSubPathOk

`func (o *GroupPath) GetIncludeSubPathOk() (*bool, bool)`

GetIncludeSubPathOk returns a tuple with the IncludeSubPath field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncludeSubPath

`func (o *GroupPath) SetIncludeSubPath(v bool)`

SetIncludeSubPath sets IncludeSubPath field to given value.


### GetPath

`func (o *GroupPath) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *GroupPath) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *GroupPath) SetPath(v string)`

SetPath sets Path field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


