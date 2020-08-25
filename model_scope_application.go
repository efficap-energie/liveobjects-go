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
// ScopeApplication struct for ScopeApplication
type ScopeApplication struct {
	// List of allowed FIFOs to publish or subscribe. Expected array of string (max 100 elements, value max 255 characters)
	Fifos []string `json:"fifos,omitempty"`
}
