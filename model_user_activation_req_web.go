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
// UserActivationReqWeb body of the user activation request
type UserActivationReqWeb struct {
	// identifier of tenant account this user belongs to. Expected identifier (max 24 characters)
	TenantId string `json:"tenantId"`
	// user identifier in tenant account. Expected identifier (max 24 characters)
	UserId string `json:"userId"`
}
