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
// ConnectorStatusResponse struct for ConnectorStatusResponse
type ConnectorStatusResponse struct {
	// list all msisdn with status treated
	StatusPerMsisdn map[string]string `json:"statusPerMsisdn,omitempty"`
}
