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
// DeviceFirmwareVersionValueWeb a firmware version
type DeviceFirmwareVersionValueWeb struct {
	// firmware version associated date/time
	Timestamp string `json:"timestamp,omitempty"`
	// firmware version. Expected string (max 255 characters)
	Version string `json:"version,omitempty"`
}
