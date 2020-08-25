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
// ActionTriggersV0 struct for ActionTriggersV0
type ActionTriggersV0 struct {
	// A list of state processing, firing or activity processing rule ids that will trigger an action - incompatible with messageSelector
	EventRuleIds []string `json:"eventRuleIds,omitempty"`
	MessageSelector MessageSelector `json:"messageSelector,omitempty"`
}