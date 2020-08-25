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
// LoraConnectivityOptions struct for LoraConnectivityOptions
type LoraConnectivityOptions struct {
	// Enable or disable ack UL option
	AckUl bool `json:"ackUl"`
	// Enable or disable TDOA option
	Tdoa bool `json:"tdoa,omitempty"`
}
