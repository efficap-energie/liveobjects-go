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
// FiringGuardResetRequest struct for FiringGuardResetRequest
type FiringGuardResetRequest struct {
	// firingRuleId linked to the FiringGuard to reset
	FiringRuleId string `json:"firingRuleId"`
	// criteria that should match the FiringGuard to reset
	SelectionCriteria []SelectionCriteria `json:"selectionCriteria,omitempty"`
}
