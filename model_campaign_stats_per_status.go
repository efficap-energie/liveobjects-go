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
// CampaignStatsPerStatus Statistics about devices with a given status
type CampaignStatsPerStatus struct {
	// Number of devices with status 'canceled'
	Canceled int32 `json:"canceled"`
	// Number of devices with status 'failure'
	Failure int32 `json:"failure"`
	// Number of devices with status 'in progress'
	InProgress int32 `json:"inProgress"`
	// Number of devices with status 'not started'
	NotStarted int32 `json:"notStarted"`
	// Number of devices with status 'pending'
	Pending int32 `json:"pending"`
	// Number of devices with status 'success'
	Success int32 `json:"success"`
}