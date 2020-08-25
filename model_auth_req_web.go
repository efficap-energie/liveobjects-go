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
// AuthReqWeb body of authentication request
type AuthReqWeb struct {
	// the user email
	Email string `json:"email,omitempty"`
	// the user login
	Login string `json:"login,omitempty"`
	// the user password
	Password string `json:"password"`
}