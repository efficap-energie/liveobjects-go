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
// SmsConnectorSettings body of sms connector business settings
type SmsConnectorSettings struct {
	// decoders use to translate messages received
	Decoders []string `json:"decoders"`
	// indicate if setting is enabled or not
	Enabled bool `json:"enabled"`
	Limits Thresholds `json:"limits,omitempty"`
	// platform phone number use as source of message
	ServerPhoneNumber string `json:"serverPhoneNumber"`
}
