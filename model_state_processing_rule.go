/*
 * Live Objects REST API Guide v2.13.3
 *
 * API description for Live Objects service
 *
 * API version: 2.13.3
 * Contact: liveobjects.support@orange.com
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package liveobjects

import (
	"encoding/json"
)

// StateProcessingRule defines the JsonLogic Function that will be applied to all new data. This function will compute a state
type StateProcessingRule struct {
	// activate or not the StateProcessingRule. Default is false.
	Enabled *bool `json:"enabled,omitempty"`
	// the JsonLogic (http://jsonlogic.com/) filter predicate applied before state function
	FilterPredicate *map[string]interface{} `json:"filterPredicate,omitempty"`
	// id of the StateProcessingRule. Should be null when used for POST.
	Id *string `json:"id,omitempty"`
	// name of the StateProcessingRule. Must be unique.
	Name string `json:"name"`
	// the JsonLogic (http://jsonlogic.com/) state function
	StateFunction *map[string]interface{} `json:"stateFunction,omitempty"`
	// the json path applied on dataMessage used to compute the stateKey
	StateKeyPath *string `json:"stateKeyPath,omitempty"`
}

// NewStateProcessingRule instantiates a new StateProcessingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStateProcessingRule(name string, ) *StateProcessingRule {
	this := StateProcessingRule{}
	this.Name = name
	return &this
}

// NewStateProcessingRuleWithDefaults instantiates a new StateProcessingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStateProcessingRuleWithDefaults() *StateProcessingRule {
	this := StateProcessingRule{}
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *StateProcessingRule) GetEnabled() bool {
	if o == nil || o.Enabled == nil {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateProcessingRule) GetEnabledOk() (*bool, bool) {
	if o == nil || o.Enabled == nil {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *StateProcessingRule) HasEnabled() bool {
	if o != nil && o.Enabled != nil {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *StateProcessingRule) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFilterPredicate returns the FilterPredicate field value if set, zero value otherwise.
func (o *StateProcessingRule) GetFilterPredicate() map[string]interface{} {
	if o == nil || o.FilterPredicate == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FilterPredicate
}

// GetFilterPredicateOk returns a tuple with the FilterPredicate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateProcessingRule) GetFilterPredicateOk() (*map[string]interface{}, bool) {
	if o == nil || o.FilterPredicate == nil {
		return nil, false
	}
	return o.FilterPredicate, true
}

// HasFilterPredicate returns a boolean if a field has been set.
func (o *StateProcessingRule) HasFilterPredicate() bool {
	if o != nil && o.FilterPredicate != nil {
		return true
	}

	return false
}

// SetFilterPredicate gets a reference to the given map[string]interface{} and assigns it to the FilterPredicate field.
func (o *StateProcessingRule) SetFilterPredicate(v map[string]interface{}) {
	o.FilterPredicate = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *StateProcessingRule) GetId() string {
	if o == nil || o.Id == nil {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateProcessingRule) GetIdOk() (*string, bool) {
	if o == nil || o.Id == nil {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *StateProcessingRule) HasId() bool {
	if o != nil && o.Id != nil {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *StateProcessingRule) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *StateProcessingRule) GetName() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *StateProcessingRule) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *StateProcessingRule) SetName(v string) {
	o.Name = v
}

// GetStateFunction returns the StateFunction field value if set, zero value otherwise.
func (o *StateProcessingRule) GetStateFunction() map[string]interface{} {
	if o == nil || o.StateFunction == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.StateFunction
}

// GetStateFunctionOk returns a tuple with the StateFunction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateProcessingRule) GetStateFunctionOk() (*map[string]interface{}, bool) {
	if o == nil || o.StateFunction == nil {
		return nil, false
	}
	return o.StateFunction, true
}

// HasStateFunction returns a boolean if a field has been set.
func (o *StateProcessingRule) HasStateFunction() bool {
	if o != nil && o.StateFunction != nil {
		return true
	}

	return false
}

// SetStateFunction gets a reference to the given map[string]interface{} and assigns it to the StateFunction field.
func (o *StateProcessingRule) SetStateFunction(v map[string]interface{}) {
	o.StateFunction = &v
}

// GetStateKeyPath returns the StateKeyPath field value if set, zero value otherwise.
func (o *StateProcessingRule) GetStateKeyPath() string {
	if o == nil || o.StateKeyPath == nil {
		var ret string
		return ret
	}
	return *o.StateKeyPath
}

// GetStateKeyPathOk returns a tuple with the StateKeyPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StateProcessingRule) GetStateKeyPathOk() (*string, bool) {
	if o == nil || o.StateKeyPath == nil {
		return nil, false
	}
	return o.StateKeyPath, true
}

// HasStateKeyPath returns a boolean if a field has been set.
func (o *StateProcessingRule) HasStateKeyPath() bool {
	if o != nil && o.StateKeyPath != nil {
		return true
	}

	return false
}

// SetStateKeyPath gets a reference to the given string and assigns it to the StateKeyPath field.
func (o *StateProcessingRule) SetStateKeyPath(v string) {
	o.StateKeyPath = &v
}

func (o StateProcessingRule) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Enabled != nil {
		toSerialize["enabled"] = o.Enabled
	}
	if o.FilterPredicate != nil {
		toSerialize["filterPredicate"] = o.FilterPredicate
	}
	if o.Id != nil {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.StateFunction != nil {
		toSerialize["stateFunction"] = o.StateFunction
	}
	if o.StateKeyPath != nil {
		toSerialize["stateKeyPath"] = o.StateKeyPath
	}
	return json.Marshal(toSerialize)
}

type NullableStateProcessingRule struct {
	value *StateProcessingRule
	isSet bool
}

func (v NullableStateProcessingRule) Get() *StateProcessingRule {
	return v.value
}

func (v *NullableStateProcessingRule) Set(val *StateProcessingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableStateProcessingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableStateProcessingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStateProcessingRule(val *StateProcessingRule) *NullableStateProcessingRule {
	return &NullableStateProcessingRule{value: val, isSet: true}
}

func (v NullableStateProcessingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStateProcessingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


