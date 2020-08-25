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
// CampaignOperation struct for CampaignOperation
type CampaignOperation struct {
	// Type of operation
	Action string `json:"action"`
	// Version of operation (default:0)
	Version int32 `json:"version,omitempty"`
}
