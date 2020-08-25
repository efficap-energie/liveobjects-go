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
// ActionPolicyV0 Object describing the relationship between triggers (message or events)  and actions (email, sms, fifo push or http push)
type ActionPolicyV0 struct {
	Actions Actions `json:"actions,omitempty"`
	// Enable or disable this action policy
	Enabled bool `json:"enabled,omitempty"`
	Id string `json:"id,omitempty"`
	// The user-defined name of the action policy
	Name string `json:"name,omitempty"`
	Triggers ActionTriggersV0 `json:"triggers,omitempty"`
}