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
// NewData original data that triggered the match
type NewData struct {
	Extra map[string]string `json:"extra,omitempty"`
	Location Location `json:"location,omitempty"`
	Metadata map[string]interface{} `json:"metadata,omitempty"`
	Model string `json:"model,omitempty"`
	StreamId string `json:"streamId,omitempty"`
	Tags []string `json:"tags,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Value map[string]interface{} `json:"value,omitempty"`
}
