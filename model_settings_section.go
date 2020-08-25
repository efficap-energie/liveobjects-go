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
// SettingsSection struct for SettingsSection
type SettingsSection struct {
	Content map[string]interface{} `json:"content,omitempty"`
	Revision int32 `json:"revision,omitempty"`
	Version string `json:"version,omitempty"`
}
