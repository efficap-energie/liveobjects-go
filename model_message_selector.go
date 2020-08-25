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
// MessageSelector struct for MessageSelector
type MessageSelector struct {
	Filter MessageSelectorFilter `json:"filter,omitempty"`
	// Specifies the source of the message that will trigger an action
	Origin string `json:"origin"`
}