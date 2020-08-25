/*
 * Live Objects REST API Guide v2.12.2
 *
 * API description for Live Objects service
 *
 * API version: 2.12.2
 * Contact: liveobjects.support@orange.com
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package liveobjects
// User body of user data
type User struct {
	// user email
	Email string `json:"email"`
	ExternalIdentity ExternalIdentity `json:"externalIdentity,omitempty"`
	// identifier of user
	Id string `json:"id"`
	// user last authentication timestamp in ms
	LastAuthentication int64 `json:"lastAuthentication,omitempty"`
	// user login
	Login string `json:"login"`
	// user is the tenant's main user
	MainUser bool `json:"mainUser,omitempty"`
	// user portal data
	PortalData map[string]interface{} `json:"portalData,omitempty"`
	// list of user associated roles.
	Roles []string `json:"roles"`
	// user state : disabled, enabled or suspended
	State string `json:"state"`
	// identifier of tenant account this user will belong to
	TenantId string `json:"tenantId"`
}