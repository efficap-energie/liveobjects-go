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
// LoraConnectivityOptionsWeb struct for LoraConnectivityOptionsWeb
type LoraConnectivityOptionsWeb struct {
	// Options for Ack uplink
	AckUl bool `json:"ackUl,omitempty"`
	// Options for TDOA geolocation
	Tdoa bool `json:"tdoa,omitempty"`
}