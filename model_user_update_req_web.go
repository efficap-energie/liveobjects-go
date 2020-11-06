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

// UserUpdateReqWeb body of user update request
type UserUpdateReqWeb struct {
	// user email. Expected a valid email (max 254 characters)
	Email *string `json:"email,omitempty"`
	// user language. Expected language ISO 639-1 (example: \"en\", \"fr\", \"sk\", \"ro\", \"es\") (max 2 characters)
	Language *string `json:"language,omitempty"`
	// the user login. If no external identity provider is used, then login must respect the following regular expression <code>[a-zA-Z0-9_-]{3,254}</code>
	Login *string `json:"login,omitempty"`
	// Set the user as main user. Expected boolean (true/false)
	MainUser *bool `json:"mainUser,omitempty"`
	// list of user associated roles. Basic roles are \"USER_R\", \"USER_W\", \"API_KEY_R\", \"API_KEY_W\" or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters)
	Roles *[]string `json:"roles,omitempty"`
	// user state : disabled, enabled or suspended
	State string `json:"state"`
	// identifier of tenant account this user belongs to. Expected identifier (max 24 characters)
	TenantId *string `json:"tenantId,omitempty"`
}

// NewUserUpdateReqWeb instantiates a new UserUpdateReqWeb object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserUpdateReqWeb(state string, ) *UserUpdateReqWeb {
	this := UserUpdateReqWeb{}
	this.State = state
	return &this
}

// NewUserUpdateReqWebWithDefaults instantiates a new UserUpdateReqWeb object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserUpdateReqWebWithDefaults() *UserUpdateReqWeb {
	this := UserUpdateReqWeb{}
	return &this
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *UserUpdateReqWeb) GetEmail() string {
	if o == nil || o.Email == nil {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateReqWeb) GetEmailOk() (*string, bool) {
	if o == nil || o.Email == nil {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *UserUpdateReqWeb) HasEmail() bool {
	if o != nil && o.Email != nil {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *UserUpdateReqWeb) SetEmail(v string) {
	o.Email = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *UserUpdateReqWeb) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateReqWeb) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *UserUpdateReqWeb) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *UserUpdateReqWeb) SetLanguage(v string) {
	o.Language = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *UserUpdateReqWeb) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateReqWeb) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *UserUpdateReqWeb) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *UserUpdateReqWeb) SetLogin(v string) {
	o.Login = &v
}

// GetMainUser returns the MainUser field value if set, zero value otherwise.
func (o *UserUpdateReqWeb) GetMainUser() bool {
	if o == nil || o.MainUser == nil {
		var ret bool
		return ret
	}
	return *o.MainUser
}

// GetMainUserOk returns a tuple with the MainUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateReqWeb) GetMainUserOk() (*bool, bool) {
	if o == nil || o.MainUser == nil {
		return nil, false
	}
	return o.MainUser, true
}

// HasMainUser returns a boolean if a field has been set.
func (o *UserUpdateReqWeb) HasMainUser() bool {
	if o != nil && o.MainUser != nil {
		return true
	}

	return false
}

// SetMainUser gets a reference to the given bool and assigns it to the MainUser field.
func (o *UserUpdateReqWeb) SetMainUser(v bool) {
	o.MainUser = &v
}

// GetRoles returns the Roles field value if set, zero value otherwise.
func (o *UserUpdateReqWeb) GetRoles() []string {
	if o == nil || o.Roles == nil {
		var ret []string
		return ret
	}
	return *o.Roles
}

// GetRolesOk returns a tuple with the Roles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateReqWeb) GetRolesOk() (*[]string, bool) {
	if o == nil || o.Roles == nil {
		return nil, false
	}
	return o.Roles, true
}

// HasRoles returns a boolean if a field has been set.
func (o *UserUpdateReqWeb) HasRoles() bool {
	if o != nil && o.Roles != nil {
		return true
	}

	return false
}

// SetRoles gets a reference to the given []string and assigns it to the Roles field.
func (o *UserUpdateReqWeb) SetRoles(v []string) {
	o.Roles = &v
}

// GetState returns the State field value
func (o *UserUpdateReqWeb) GetState() string {
	if o == nil  {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *UserUpdateReqWeb) GetStateOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *UserUpdateReqWeb) SetState(v string) {
	o.State = v
}

// GetTenantId returns the TenantId field value if set, zero value otherwise.
func (o *UserUpdateReqWeb) GetTenantId() string {
	if o == nil || o.TenantId == nil {
		var ret string
		return ret
	}
	return *o.TenantId
}

// GetTenantIdOk returns a tuple with the TenantId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserUpdateReqWeb) GetTenantIdOk() (*string, bool) {
	if o == nil || o.TenantId == nil {
		return nil, false
	}
	return o.TenantId, true
}

// HasTenantId returns a boolean if a field has been set.
func (o *UserUpdateReqWeb) HasTenantId() bool {
	if o != nil && o.TenantId != nil {
		return true
	}

	return false
}

// SetTenantId gets a reference to the given string and assigns it to the TenantId field.
func (o *UserUpdateReqWeb) SetTenantId(v string) {
	o.TenantId = &v
}

func (o UserUpdateReqWeb) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Email != nil {
		toSerialize["email"] = o.Email
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.MainUser != nil {
		toSerialize["mainUser"] = o.MainUser
	}
	if o.Roles != nil {
		toSerialize["roles"] = o.Roles
	}
	if true {
		toSerialize["state"] = o.State
	}
	if o.TenantId != nil {
		toSerialize["tenantId"] = o.TenantId
	}
	return json.Marshal(toSerialize)
}

type NullableUserUpdateReqWeb struct {
	value *UserUpdateReqWeb
	isSet bool
}

func (v NullableUserUpdateReqWeb) Get() *UserUpdateReqWeb {
	return v.value
}

func (v *NullableUserUpdateReqWeb) Set(val *UserUpdateReqWeb) {
	v.value = val
	v.isSet = true
}

func (v NullableUserUpdateReqWeb) IsSet() bool {
	return v.isSet
}

func (v *NullableUserUpdateReqWeb) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserUpdateReqWeb(val *UserUpdateReqWeb) *NullableUserUpdateReqWeb {
	return &NullableUserUpdateReqWeb{value: val, isSet: true}
}

func (v NullableUserUpdateReqWeb) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserUpdateReqWeb) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


