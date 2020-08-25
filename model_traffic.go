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
// Traffic struct for Traffic
type Traffic struct {
	In map[string]int64 `json:"in,omitempty"`
	Out map[string]int64 `json:"out,omitempty"`
}
