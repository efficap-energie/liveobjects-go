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
// LoraGatewayInfo struct for LoraGatewayInfo
type LoraGatewayInfo struct {
	Config GwConfig `json:"config,omitempty"`
	// Description
	Description string `json:"description,omitempty"`
	// Unique id of the gateway
	Id string `json:"id,omitempty"`
	// Date/time of the last report
	LastReportTs string `json:"lastReportTs,omitempty"`
	Location LoraGatewayLocation `json:"location,omitempty"`
	// Manufacturer of the gateway
	Manufacturer string `json:"manufacturer,omitempty"`
	// Name of the gateway
	Name string `json:"name,omitempty"`
	// Status of the gateway
	Status string `json:"status,omitempty"`
	System GwSystem `json:"system,omitempty"`
	// Type of the gateway
	Type string `json:"type,omitempty"`
}
