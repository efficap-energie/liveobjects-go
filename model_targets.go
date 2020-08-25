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
// Targets defines the targeted device of an ActivityRule
type Targets struct {
	// List of targeted device ids. This field should not be empty if groupPaths is empty.
	DeviceIds []string `json:"deviceIds,omitempty"`
	// list of targeted group paths. This field should not be empty if deviceIds is empty.
	GroupPaths []GroupPath `json:"groupPaths,omitempty"`
}
