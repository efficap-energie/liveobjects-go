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
// CampaignOperationState struct for CampaignOperationState
type CampaignOperationState struct {
	// Type of operation
	Action string `json:"action"`
	// Optional current number of retries made for the operation
	CurrentRetry int32 `json:"currentRetry,omitempty"`
	// Date when the operation was finished
	Ended string `json:"ended,omitempty"`
	Error CampaignOperationStateError `json:"error,omitempty"`
	// Optional identifier associated to the current underlying operation (e.g. commandId,...)
	OperationId string `json:"operationId,omitempty"`
	// Current status returned by the underlying operation
	OperationStatus string `json:"operationStatus,omitempty"`
	// Date when the operation was started
	Started string `json:"started,omitempty"`
	// Date when the operation status was updated for the last time
	Updated string `json:"updated,omitempty"`
}
