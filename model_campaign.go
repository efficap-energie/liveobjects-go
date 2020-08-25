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
// Campaign struct for Campaign
type Campaign struct {
	CampaignStatus string `json:"campaignStatus,omitempty"`
	Created string `json:"created,omitempty"`
	// Campaign description
	Description string `json:"description,omitempty"`
	// campaign id
	Id string `json:"id"`
	// campaign name
	Name string `json:"name"`
	// number of targets after deduplication
	NumberOfTargets int32 `json:"numberOfTargets,omitempty"`
	// requested operations
	Operations []CampaignOperation `json:"operations"`
	Options CampaignOptions `json:"options,omitempty"`
	Planning CampaignPlanning `json:"planning"`
	Target DeviceSelector `json:"target"`
	TotalTargetsPerStatus CampaignStatsPerStatus `json:"totalTargetsPerStatus,omitempty"`
	Updated string `json:"updated,omitempty"`
}
