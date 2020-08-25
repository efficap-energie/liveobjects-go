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
// InterfaceCapability struct for InterfaceCapability
type InterfaceCapability struct {
	// Indicates whether the capability is available for the interface
	Available bool `json:"available,omitempty"`
	// Capability version
	Version int32 `json:"version,omitempty"`
}