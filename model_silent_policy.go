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
// SilentPolicy defines the triggering behavior for silent devices
type SilentPolicy struct {
	// the duration (ISO 8601) after which the device will be considered as silent if it had no contact with the platform. If less than 10 minutes, this duration will be handled as a 10 minutes delay. A silent event/alarm will be sent after expiration.
	Duration string `json:"duration"`
	// if the device previously triggered a silent alarm and still had no contact with the platform; duration (ISO 8601) after which a new silent event/alarm will be sent. If less than 10 minutes, this duration will be handled as a 10 minutes delay. If null, no further alarm will be sent until the new active/silent cycle.
	RepeatInterval string `json:"repeatInterval,omitempty"`
}