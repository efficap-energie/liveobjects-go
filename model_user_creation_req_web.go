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
// UserCreationReqWeb body for create a new user
type UserCreationReqWeb struct {
	// the user email. Expected a valid email (max 254 characters)
	Email string `json:"email"`
	// user language. Expected language ISO 639-1 (example: \"en\", \"fr\", \"sk\", \"ro\", \"es\") (max 2 characters)
	Language string `json:"language,omitempty"`
	// the user login. If no external identity provider is used, then login must respect the following regular expression <code>[a-zA-Z0-9_-]{3,254}</code>
	Login string `json:"login"`
	// list of user associated roles. Basic roles are \"USER_R\", \"USER_W\", \"API_KEY_R\", \"API_KEY_W\" or any role string supplied at tenant account creation time. Expected array of role name (max all roles, role value max 255 characters)
	Roles []string `json:"roles"`
	// identifier of tenant account this user will belong to. Expected identifier (max 24 characters)
	TenantId string `json:"tenantId"`
}
