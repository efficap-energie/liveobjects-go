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
// SmsConnectorBusinessSettingsCreationReqWeb SMS Connector Business Settings  Data in Creation Request
type SmsConnectorBusinessSettingsCreationReqWeb struct {
	// decoder name ex: \"decoderName\"
	DecoderName string `json:"decoderName,omitempty"`
	// device msisdn
	Msisdns []Msisdns `json:"msisdns"`
	// server phone number ex: \"20258\"
	ServerPhoneNumber string `json:"serverPhoneNumber"`
}
