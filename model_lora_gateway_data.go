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
// LoraGatewayData struct for LoraGatewayData
type LoraGatewayData struct {
	// Description
	Description string `json:"description,omitempty"`
	Location LoraGatewayLocation `json:"location,omitempty"`
	// Name of the gateway
	Name string `json:"name,omitempty"`
}